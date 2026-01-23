package auth

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

// loginAuth implements the AUTH LOGIN mechanism for SMTP.
// It is considered "basic auth" (username + password) just like PLAIN,
// but uses the challenge/response dance:
//
//	C: AUTH LOGIN
//	S: 334 VXNlcm5hbWU6  ("Username:")
//	C: base64(username)
//	S: 334 UGFzc3dvcmQ6  ("Password:")
//	C: base64(password)

// based on https://gist.github.com/andelf/5118732 (MIT licensed)
type loginAuth struct {
	username string
	password string
	host     string
}

// LoginAuth creates a new smtp.Auth for AUTH LOGIN.

// LoginAuth will only send the credentials if the connection is using TLS
// or is connected to localhost. Otherwise authentication will fail with an
// error, without sending the credentials (same safety as smtp.PlainAuth).
func LoginAuth(username, password string, host string) smtp.Auth {
	return &loginAuth{
		username: username,
		password: password,
		host:     host,
	}
}

func isLocalhost(name string) bool {
	return name == "localhost" || name == "127.0.0.1" || name == "::1"
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// Must have TLS, or else localhost server.
	// Note: If TLS is not true, then we can't trust ANYTHING in ServerInfo.
	// In particular, it doesn't matter if the server advertises PLAIN auth.
	// That might just be the attacker saying
	// "it's ok, you can trust me with your password."
	if !server.TLS && !isLocalhost(server.Name) {
		return "", nil, errors.New("unencrypted connection")
	}
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	// Initial response for LOGIN is usually empty;
	// we wait for "Username:" challenge.
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		// No more challenges; we're done.
		return nil, nil
	}

	// fromServer is the decoded challenge string, e.g. "Username:" or "Password:"
	prompt := strings.ToLower(strings.TrimSpace(string(fromServer)))

	switch {
	case strings.Contains(prompt, "username"):
		return []byte(a.username), nil
	case strings.Contains(prompt, "password"):
		return []byte(a.password), nil
	default:
		return nil, fmt.Errorf("unexpected server challenge: %q", fromServer)
	}
}
