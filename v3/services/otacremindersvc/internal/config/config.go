package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AuthMechanism controls which SMTP AUTH mechanism is used.
type AuthMechanism string

const (
	AuthMechanismAuto  AuthMechanism = ""      // prefer PLAIN, then LOGIN
	AuthMechanismPlain AuthMechanism = "PLAIN" // force AUTH PLAIN
	AuthMechanismLogin AuthMechanism = "LOGIN" // force AUTH LOGIN
)

// EmailBodyContentType controls the content type of the email body.
type EmailBodyContentType string

const (
	EmailBodyContentTypePlain EmailBodyContentType = "text/plain" // plain text email body
	EmailBodyContentTypeHTML  EmailBodyContentType = "text/html"  // HTML email body
)

type Config struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPass     string
	SMTPFrom     string
	SMTPAuthType string
	CC           string
	ReplyTo      string
	Signature    string

	ScheduledEventId string

	// the exact template variables available in the NotificationSubjectTemplate and NotificationBodyTemplate are defined in the EmailTemplateData struct in internal/notifier/notifier.go
	NotificationSubjectTemplate string // optional custom subject template for notification emails, e.g. "Your access code expires in {{.RemainingDays}} days"
	NotificationBodyTemplate    string // optional custom body template for notification emails, e.g. "Hello @{{.RecipientEmail}}, your access code expires in {{.RemainingDays}} days on {{.ExpirationDateTimeUTC}}"
	// raw config for multi-window notifications
	NotificationWindowsRaw string // e.g. "6d,5d,4d,3d,2d,1d" or "144h,120h,96h,72h,48h,24h"

	RequireTLS      bool
	RequireAuth     bool
	EmailTimeout    int64 // timeout per email in seconds
	NotifierTimeout int64 // timeout for the notifier to send out all emails in seconds

	MountPathCA string // add CA cert path for email server

	ContentType EmailBodyContentType // content type for email body only supports "text/plain" or "text/html" and defaults to "text/plain"
}

// sensible defaults if env vars are not set:
const (
	defaultNotificationWindows = "6d,5d,4d,3d,2d,1d"
)

const (
	defaultNotificationSubjectTemplate = "Your one-time access expires in approximately {{.RemainingDays}} days and {{.RemainingHours}} hours"
	defaultNotificationBodyTemplate    = `Hello @{{.RecipientEmail}},

your one-time access was redeemed at:

    UTC:                     {{.ActivationDateTimeUTC}}
    Eastern (EST/EDT):       {{.ActivationDateTimeEastern}}
    Central Europe (CET/CEST): {{.ActivationDateTimeCentralEurope}}

with a maximum duration of {{.MaxDuration}}.
It will expire at:

    UTC:                     {{.ExpirationDateTimeUTC}}
    Eastern (EST/EDT):       {{.ExpirationDateTimeEastern}}
    Central Europe (CET/CEST): {{.ExpirationDateTimeCentralEurope}}

That is in approximately {{.RemainingHoursTotal}} hour(s).

If you still need access, please log in again or request a new token.
`
)

// check required env vars and set defaults for optional env vars
func FromEnv() (*Config, error) {
	// helper function to retrieve env vars
	get := func(key, def string, required bool) (string, error) {
		v := os.Getenv(key)
		if v == "" {
			v = def
		}
		if required && v == "" {
			return "", fmt.Errorf("missing required env var: %s", key)
		}
		return v, nil
	}

	smtpHost, err := get("SMTP_HOST", "", true)
	if err != nil {
		return nil, err
	}
	smtpPort, _ := get("SMTP_PORT", "587", false)
	smtpAuthType, _ := get("SMTP_AUTH_TYPE", "", false)

	cc, _ := get("SMTP_CC", "", false)
	replyTo, _ := get("SMTP_REPLY_TO", "", false)
	signature, _ := get("SMTP_SIGNATURE", "", false)

	eventId, err := get("SCHEDULED_EVENT_ID", "", true)
	if err != nil {
		return nil, err
	}

	notificationSubjectTemplate, _ := get("NOTIFICATION_SUBJECT_TEMPLATE", defaultNotificationSubjectTemplate, false)
	notificationBodyTemplate, _ := get("NOTIFICATION_BODY_TEMPLATE", defaultNotificationBodyTemplate, false)
	windowsRaw, _ := get("NOTIFICATION_WINDOWS", defaultNotificationWindows, false)

	requireTlsRaw, _ := get("SMTP_REQUIRE_TLS", "true", false)
	requireAuthRaw, _ := get("SMTP_REQUIRE_AUTH", "true", false)
	emailTimeoutRaw, _ := get("EMAIL_TIMEOUT", "5", false)
	notifierTimeoutRaw, _ := get("NOTIFIER_TIMEOUT", "60", false)

	requireTls, err := strconv.ParseBool(requireTlsRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env var: SMTP_REQUIRE_TLS")
	}

	requireAuth, err := strconv.ParseBool(requireAuthRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env var: SMTP_REQUIRE_AUTH")
	}
	smtpUser, err := get("SMTP_USER", "", true)
	if requireAuth && err != nil {
		return nil, err
	}
	smtpPass, _ := get("SMTP_PASSWORD", "", true)
	if requireAuth && err != nil {
		return nil, err
	}
	smtpFrom, _ := get("SMTP_FROM", smtpUser, true)
	if err != nil {
		return nil, err
	}

	emailTimeout, err := strconv.ParseInt(emailTimeoutRaw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env var: EMAIL_TIMEOUT")
	}

	notifierTimeout, err := strconv.ParseInt(notifierTimeoutRaw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env var: NOTIFIER_TIMEOUT")
	}

	mountPathCA, _ := get("MOUNT_PATH_CA", "", false)

	contentType, _ := get("CONTENT_TYPE", string(EmailBodyContentTypePlain), false)
	if contentType != string(EmailBodyContentTypePlain) && contentType != string(EmailBodyContentTypeHTML) {
		return nil, fmt.Errorf("invalid value for CONTENT_TYPE: %s. Only '%s' and '%s' are supported", contentType, EmailBodyContentTypePlain, EmailBodyContentTypeHTML)
	}

	return &Config{
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPUser:     smtpUser,
		SMTPPass:     smtpPass,
		SMTPFrom:     smtpFrom,
		SMTPAuthType: parseAuthType(smtpAuthType),
		CC:           cc,
		ReplyTo:      replyTo,
		Signature:    signature,

		ScheduledEventId: eventId,

		NotificationSubjectTemplate: notificationSubjectTemplate,
		NotificationBodyTemplate:    notificationBodyTemplate,
		NotificationWindowsRaw:      windowsRaw,
		RequireTLS:                  requireTls,
		RequireAuth:                 requireAuth,
		EmailTimeout:                emailTimeout,
		NotifierTimeout:             notifierTimeout,
		MountPathCA:                 mountPathCA,
		ContentType:                 EmailBodyContentType(contentType),
	}, nil
}

func parseAuthType(authType string) string {
	val := strings.ToUpper(strings.TrimSpace(authType))
	switch val {
	case "PLAIN":
		return string(AuthMechanismPlain)
	case "LOGIN":
		return string(AuthMechanismLogin)
	case "AUTO", "":
		return string(AuthMechanismAuto)
	default:
		return val // we don't handle other mechanisms gracefully to indicate to a user with an error message, that his auth type is not supported
	}
}
