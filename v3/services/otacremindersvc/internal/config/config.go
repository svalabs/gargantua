package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	SMTPHost  string
	SMTPPort  string
	SMTPUser  string
	SMTPPass  string
	SMTPFrom  string
	CC        string
	ReplyTo   string
	Signature string

	ScheduledEventId string

	// raw config for multi-window notifications
	NotificationWindowsRaw string // e.g. "6d,5d,4d,3d,2d,1d" or "144h,120h,96h,72h,48h,24h"

	RequireTLS      bool
	RequireAuth     bool
	EmailTimeout    int64 // timeout per email in seconds
	NotifierTimeout int64 // timeout for the notifier to send out all emails in seconds

	MountPathCA string // add CA cert path for email server
}

// sensible defaults if env vars are not set:
const (
	defaultNotificationWindows = "6d,5d,4d,3d,2d,1d"
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
	smtpUser, _ := get("SMTP_USER", "", true)
	smtpPass, _ := get("SMTP_PASSWORD", "", true)
	smtpFrom, _ := get("SMTP_FROM", smtpUser, true)

	cc, _ := get("SMTP_CC", "", false)
	replyTo, _ := get("SMTP_REPLY_TO", "", false)
	signature, _ := get("SMTP_SIGNATURE", "", false)

	eventId, _ := get("SCHEDULED_EVENT_ID", "", true)

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

	emailTimeout, err := strconv.ParseInt(emailTimeoutRaw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env var: EMAIL_TIMEOUT")
	}

	notifierTimeout, err := strconv.ParseInt(notifierTimeoutRaw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse env var: NOTIFIER_TIMEOUT")
	}

	mountPathCA, _ := get("MOUNT_PATH_CA", "", false)

	return &Config{
		SMTPHost:  smtpHost,
		SMTPPort:  smtpPort,
		SMTPUser:  smtpUser,
		SMTPPass:  smtpPass,
		SMTPFrom:  smtpFrom,
		CC:        cc,
		ReplyTo:   replyTo,
		Signature: signature,

		ScheduledEventId: eventId,

		NotificationWindowsRaw: windowsRaw,
		RequireTLS:             requireTls,
		RequireAuth:            requireAuth,
		EmailTimeout:           emailTimeout,
		NotifierTimeout:        notifierTimeout,
		MountPathCA:            mountPathCA,
	}, nil
}
