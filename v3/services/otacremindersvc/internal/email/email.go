package email

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"

	auth "github.com/hobbyfarm/gargantua/services/otacremindersvc/v3/internal/auth"
	"github.com/hobbyfarm/gargantua/services/otacremindersvc/v3/internal/config"
)

type Client struct {
	Host        string        // SMTP host (without port)
	Port        string        // SMTP port, e.g. "587"
	User        string        // optional username for AUTH
	Password    string        // optional password for AUTH
	From        string        // from address (as it should appear in the email)
	RequireTLS  bool          // if true, fail when server does not support STARTTLS
	RequireAuth bool          // if true, fail if we cannot or do not authenticate
	Timeout     time.Duration // dial timeout, default if <= 0

	// ReplyTo is an optional reply address.
	// If non-empty, it will be used in the "Reply-To" header so that replies
	// go to this address instead of From.
	ReplyTo string

	// CC is an optional list of addresses that should receive a copy of every mail.
	// It can be a single address like "hobbyfarm@example.com" or a comma-separated list:
	// "hobby@example.com, farm@example.com".
	CC string

	// Signature is appended to the body if non-empty.
	// Can contain multiple lines, e.g.:
	// "Best regards,\nHobbyFarm OTAC Bot"
	Signature string

	// MountPathCA, if not empty, points to a CA bundle PEM file that is added
	// to the system cert pool for TLS verification.
	MountPathCA string

	// AuthMechanism selects which AUTH method to use.
	// - AuthMechanismAuto (default): Detect if PLAIN or LOGIN auth is available and prefer PLAIN auth over LOGIN auth.
	// - AuthMechanismPlain: force PLAIN (error if not supported)
	// - AuthMechanismLogin: force LOGIN (error if not supported)
	AuthMechanism config.AuthMechanism
}

const defaultTimeout = 5 * time.Second

// Send sends a simple plain-text email to a single primary recipient.
// Optionally adds CC recipients and a Reply-To address configured on the client.
func (c *Client) Send(to, subject, body string) error {
	host := strings.TrimSpace(c.Host)
	port := strings.TrimSpace(c.Port)

	if host == "" || port == "" {
		return fmt.Errorf("email: Host and Port must be set")
	}

	addr := net.JoinHostPort(host, port)
	timeout := c.Timeout
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return fmt.Errorf("email: dial %s: %w", addr, err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("email: new client: %w", err)
	}
	defer client.Close()

	// Build root CA pool (system + optional custom CA)
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		rootCAs = x509.NewCertPool()
	}
	if c.MountPathCA != "" {
		caCert, err := os.ReadFile(c.MountPathCA)
		if err != nil {
			return fmt.Errorf("email: load ca cert: %w", err)
		}
		if ok := rootCAs.AppendCertsFromPEM(caCert); !ok {
			return fmt.Errorf("email: append ca cert: no certs found")
		}
	}

	// STARTTLS (preferred, and required when RequireTLS is true)
	if ok, _ := client.Extension("STARTTLS"); ok {
		tlsConfig := &tls.Config{
			ServerName: host,
			RootCAs:    rootCAs,
		}
		if err := client.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("email: starttls: %w", err)
		}
	} else if c.RequireTLS {
		return fmt.Errorf("email: server does not support STARTTLS and RequireTLS is true")
	}

	// AUTH (only after TLS). Preferred, and required when RequireAuth is true.
	authn, authType, err := c.authenticate(client)
	if err != nil {
		if c.RequireAuth {
			return fmt.Errorf("email: RequireAuth is true but authentication failed: %w", err)
		} else {
			glog.Warningf("email: authentication failed... skipping authentication because it is not required: %v", err)
		}
	}
	// Actual authentication needs to be done here, because connection may fail if TLS is not enabled (even if RequireAuth is false).
	if err := client.Auth(authn); err != nil {
		return fmt.Errorf("email: auth (%s): %w", authType, err)
	}

	from := sanitizeHeader(c.From)
	to = sanitizeHeader(to)
	subject = sanitizeHeader(subject)
	replyTo := sanitizeHeader(c.ReplyTo)

	if from == "" {
		return fmt.Errorf("email: From must not be empty")
	}
	if to == "" {
		return fmt.Errorf("email: To must not be empty")
	}

	ccAddrs := splitAndCleanAddresses(c.CC) // []string, already sanitized for headers/rcpt
	var ccHeader string
	if len(ccAddrs) > 0 {
		ccHeader = strings.Join(ccAddrs, ", ")
	}

	msg := buildPlainTextMessage(from, to, ccHeader, replyTo, subject, body, c.Signature)

	if err := client.Mail(from); err != nil {
		return fmt.Errorf("email: mail from: %w", err)
	}

	// Primary recipient
	if err := client.Rcpt(to); err != nil {
		return fmt.Errorf("email: rcpt to: %w", err)
	}

	// CC recipients (if any)
	for _, cc := range ccAddrs {
		if err := client.Rcpt(cc); err != nil {
			return fmt.Errorf("email: rcpt cc %s: %w", cc, err)
		}
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("email: data: %w", err)
	}
	if _, err := w.Write([]byte(msg)); err != nil {
		_ = w.Close()
		return fmt.Errorf("email: write: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("email: close data: %w", err)
	}

	// Quit error is usually non-fatal once DATA succeeded; ignore it.
	_ = client.Quit()

	return nil
}

func (c *Client) authenticate(client *smtp.Client) (smtp.Auth, config.AuthMechanism, error) {

	user := strings.TrimSpace(c.User)
	pass := c.Password

	if user == "" || pass == "" {
		return nil, "", fmt.Errorf("email: no credentials provided")
	}

	// starting from here, username and password are defined

	// Check if server supports AUTH at all
	ok, mechList := client.Extension("AUTH")
	if !ok {
		return nil, "", fmt.Errorf("email: server does not advertise AUTH")
	}

	mechs := parseAuthMechanisms(mechList)
	chosenAuthMech, err := chooseAuthMechanism(c.AuthMechanism, mechs)
	if err != nil {
		return nil, "", fmt.Errorf("email: auth mechanism selection failed: %w", err)
	}

	switch chosenAuthMech {
	case "PLAIN":
		return smtp.PlainAuth("", user, pass, c.Host), chosenAuthMech, nil
	case "LOGIN":
		return auth.LoginAuth(user, pass, c.Host), chosenAuthMech, nil
	default:
		return nil, "", fmt.Errorf("email: unsupported chosen auth mechanism %q", chosenAuthMech)
	}
}

// parseAuthMechanisms parses an AUTH capability string like "PLAIN LOGIN XOAUTH2"
// into a slice of upper-case mechanism names.
func parseAuthMechanisms(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Fields(s)
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.ToUpper(strings.TrimSpace(p))
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}

// chooseAuthMechanism decides which mechanism to use, given the client's preference
// and the server's advertised mechanisms.
func chooseAuthMechanism(pref config.AuthMechanism, serverMechs []string) (config.AuthMechanism, error) {
	// Build a quick lookup map
	supported := make(map[string]bool, len(serverMechs))
	for _, m := range serverMechs {
		supported[m] = true
	}

	switch pref {
	case config.AuthMechanismPlain:
		if !supported["PLAIN"] {
			return "", fmt.Errorf("server does not support AUTH PLAIN")
		}
		return config.AuthMechanismPlain, nil

	case config.AuthMechanismLogin:
		if !supported["LOGIN"] {
			return "", fmt.Errorf("server does not support AUTH LOGIN")
		}
		return config.AuthMechanismLogin, nil

	case config.AuthMechanismAuto:
		// Auto: prefer PLAIN, fall back to LOGIN
		if supported["PLAIN"] {
			return config.AuthMechanismPlain, nil
		}
		if supported["LOGIN"] {
			return config.AuthMechanismLogin, nil
		}
		return "", fmt.Errorf("server supports no known auth mechanisms (have: %v)", serverMechs)

	default:
		return "", fmt.Errorf("unknown AuthMechanism %q", pref)
	}
}

// sanitizeHeader removes CR/LF to prevent header injection.
func sanitizeHeader(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}

// splitAndCleanAddresses splits a comma-separated list of addresses,
// trims spaces, and sanitizes each address.
func splitAndCleanAddresses(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		addr := strings.TrimSpace(p)
		if addr == "" {
			continue
		}
		addr = sanitizeHeader(addr)
		out = append(out, addr)
	}
	return out
}

// buildPlainTextMessage builds a simple UTF-8 plain-text MIME message,
// optionally including Cc and Reply-To headers and appending a signature after the body.
func buildPlainTextMessage(from, to, cc, replyTo, subject, body, signature string) string {
	var b strings.Builder

	b.WriteString("From: ")
	b.WriteString(from)
	b.WriteString("\r\n")

	b.WriteString("To: ")
	b.WriteString(to)
	b.WriteString("\r\n")

	if cc != "" {
		b.WriteString("Cc: ")
		b.WriteString(cc)
		b.WriteString("\r\n")
	}

	if replyTo != "" {
		b.WriteString("Reply-To: ")
		b.WriteString(replyTo)
		b.WriteString("\r\n")
	}

	b.WriteString("Subject: ")
	b.WriteString(subject)
	b.WriteString("\r\n")

	b.WriteString("MIME-Version: 1.0\r\n")
	b.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	b.WriteString("\r\n")

	// Body
	b.WriteString(body)

	// Optional signature
	if signature != "" {
		b.WriteString("\r\n\r\n")
		b.WriteString(signature)
	}

	return b.String()
}
