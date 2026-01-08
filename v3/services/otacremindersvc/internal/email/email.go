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

	MountPathCA string // if not empty, add CA for email server
}

const defaultTimeout = 5 * time.Second

// Send sends a simple plain-text email to a single primary recipient.
// Optionally adds CC recipients and a Reply-To address configured on the client.
func (c *Client) Send(to, subject, body string) error {
	if c.Host == "" || c.Port == "" {
		return fmt.Errorf("email: Host and Port must be set")
	}

	addr := net.JoinHostPort(c.Host, c.Port)
	timeout := c.Timeout
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return fmt.Errorf("email: dial %s: %w", addr, err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, c.Host)
	if err != nil {
		return fmt.Errorf("email: new client: %w", err)
	}
	defer client.Close()

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		rootCAs = x509.NewCertPool()
	}

	if c.MountPathCA != "" {
		caCert, err := os.ReadFile(c.MountPathCA)
		if err != nil {
			return fmt.Errorf("load ca cert: %w", err)
		}
		if ok := rootCAs.AppendCertsFromPEM(caCert); !ok {
			return fmt.Errorf("append ca cert: no certs found")
		}
	}

	// STARTTLS (preferred, and required when RequireTLS is true)
	if ok, _ := client.Extension("STARTTLS"); ok {
		tlsConfig := &tls.Config{
			ServerName: c.Host,
			RootCAs:    rootCAs,
		}
		if err := client.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("email: starttls: %w", err)
		}
	} else if c.RequireTLS {
		return fmt.Errorf("email: server does not support STARTTLS and RequireTLS is true")
	}

	// AUTH (only after TLS)
	if c.RequireAuth && (c.User == "" || c.Password == "") {
		return fmt.Errorf("email: RequireAuth is true but no credentials provided")
	}

	if c.User != "" && c.Password != "" {
		auth := smtp.PlainAuth("", c.User, c.Password, c.Host)
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("email: auth: %w", err)
		}
	}

	from := sanitizeHeader(c.From)
	to = sanitizeHeader(to)
	subject = sanitizeHeader(subject)
	replyTo := sanitizeHeader(c.ReplyTo)

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
