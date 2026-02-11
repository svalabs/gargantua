package notifier

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/golang/glog"
	"google.golang.org/protobuf/types/known/timestamppb"

	config "github.com/hobbyfarm/gargantua/services/otacremindersvc/v3/internal/config"
	email "github.com/hobbyfarm/gargantua/services/otacremindersvc/v3/internal/email"
	"github.com/hobbyfarm/gargantua/v3/pkg/util"

	hferrors "github.com/hobbyfarm/gargantua/v3/pkg/errors"
	hflabels "github.com/hobbyfarm/gargantua/v3/pkg/labels"
	accesscodepb "github.com/hobbyfarm/gargantua/v3/protos/accesscode"
	generalpb "github.com/hobbyfarm/gargantua/v3/protos/general"
	userpb "github.com/hobbyfarm/gargantua/v3/protos/user"
)

type EmailTemplateData struct {
	RecipientEmail                  string
	RemainingDays                   int
	RemainingHours                  int
	RemainingHoursTotal             int
	ActivationDateTimeUTC           string
	ExpirationDateTimeUTC           string
	ActivationDateTimeEastern       string
	ExpirationDateTimeEastern       string
	ActivationDateTimeCentralEurope string
	ExpirationDateTimeCentralEurope string
	MaxDuration                     string
	WindowKey                       string
}

func Run(
	ctx context.Context,
	cfg *config.Config,
	emailClient *email.Client,
	acClient accesscodepb.AccessCodeSvcClient,
	userClient userpb.UserSvcClient,
) error {
	// parse windows & job interval from config
	windows, err := parseWindows(cfg.NotificationWindowsRaw)
	if err != nil {
		return fmt.Errorf("notifier: parse notification windows: %w", err)
	}

	now := time.Now().UTC()
	ns := util.GetReleaseNamespace()

	glog.Infof("notifier: starting OTAC expiry check for namespace %q with windows %q at %v", ns, cfg.NotificationWindowsRaw, now)

	// list OTACs
	listCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	otacList, err := acClient.ListOtac(listCtx, &generalpb.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", hflabels.ScheduledEventLabel, cfg.ScheduledEventId),
	})
	if err != nil {
		return fmt.Errorf("notifier: list otacs: %w", err)
	}

	glog.Infof("notifier: listed OTACs: %d", len(otacList.Otacs))

	for _, otac := range otacList.Otacs {
		otacId := otac.GetId()

		redeemedStr := strings.TrimSpace(otac.GetRedeemedTimestamp())
		maxDurationStr := strings.TrimSpace(otac.GetMaxDuration())
		if redeemedStr == "" || maxDurationStr == "" {
			// not redeemed or no duration
			continue
		}

		// parse redeemed timestamp (UnixDate)
		redeemed, err := time.Parse(time.UnixDate, redeemedStr)
		if err != nil {
			glog.Warningf("notifier: invalid RedeemedTimestamp %q on OTAC %q: %v", redeemedStr, otacId, err)
			continue
		}
		redeemed = redeemed.UTC()

		maxDur, err := parseDurationWithDays(maxDurationStr)
		if err != nil {
			glog.Warningf("notifier: invalid MaxDuration %q on OTAC %q: %v", maxDurationStr, otacId, err)
			continue
		}

		expiry := redeemed.Add(maxDur)
		remaining := expiry.Sub(now)
		if remaining <= 0 {
			// already expired
			continue
		}

		// ensure maps exist
		windowsMap := ensureWindows(otac)

		// retrieve user
		userID := strings.TrimSpace(otac.GetUser())
		if userID == "" {
			glog.Infof("notifier: OTAC %q has empty spec.User, skipping", otacId)
			continue
		}

		userCtx, cancelUser := context.WithTimeout(ctx, 10*time.Second)

		user, err := userClient.GetUserById(userCtx, &generalpb.GetRequest{Id: userID})
		cancelUser()

		if hferrors.IsGrpcNotFound(err) {
			glog.Infof("notifier: user %q for OTAC %q not found, skipping",
				userID,
				otacId,
			)
			continue
		}
		if err != nil {
			glog.Warningf("notifier: failed to get user %q for OTAC %q: %v", userID, otacId, err)
			continue
		}

		emailAddr := strings.TrimSpace(user.GetEmail())
		if emailAddr == "" {
			glog.Warningf("notifier: user %q for OTAC %q has no email, skipping notification", userID, otacId)
			continue
		}

		// per-window evaluation
		changed := false

		// Find the smallest already-sent configured window duration.
		// Once a smaller window is sent (e.g. 6d)...
		// larger windows (e.g. 7d, 8d) must never be sent afterwards.
		minSent := time.Duration(0)
		hasSent := false

		for key := range windowsMap {
			w, ok := windows[key]
			if !ok {
				// Ignore stale status keys that are not in current config.
				// This can happen if the configured windows changed since the last run.
				// Note: Changing windows config is not recommended and may lead to unexpected notifications. It however is handled gracefully to prevent crashes and to allow fixing config mistakes without manual OTAC status resets.
				// Example: If "3d" was sent, then changing config to "4d,1d" would cause "4d" to be sent afterwards, because although "3d" was already sent, it is not present in the current config anymore.
				glog.Warningf("notifier: found unexpected window key in OTAC status, ignoring: %q. This may happen if the configured windows changed since the last run.", key)
				continue
			}
			if !hasSent || w < minSent {
				minSent = w
				hasSent = true
			}
		}

		// Find the closest unsent window for which remaining <= w.
		// We need to do this, to prevent sending duplicate emails.
		bestKey := ""
		bestWindow := time.Duration(0)
		for key, w := range windows {
			// already sent for this window?
			if _, ok := windowsMap[key]; ok {
				continue
			}

			// Defensive guard to prevent older/larger windows triggering notifications after a smaller one was already sent.
			if hasSent && w > minSent {
				continue
			}

			if remaining <= w {
				// pick the smallest window that matches (closest to expiry)
				if bestKey == "" || w < bestWindow {
					bestKey = key
					bestWindow = w
				}
			}
		}

		if bestKey != "" {
			key := bestKey

			glog.Infof("notifier: sending expiry email for window %q on OTAC %q for user %q with email %q", key, otacId, userID, emailAddr)
			glog.Infof("notifier: remaining hours until expiry: %.2f (expiry at %s)", remaining.Hours(), expiry.UTC().Format(time.RFC1123))

			if err := sendExpiryMailForWindow(emailClient, emailAddr, key, remaining, expiry, redeemed, maxDurationStr, cfg.NotificationSubjectTemplate, cfg.NotificationBodyTemplate); err != nil {
				glog.Errorf("notifier: failed to send expiry email for window %q on OTAC %q for user %q with email %q: %v", key, otacId, userID, emailAddr, err)
			} else {
				windowsMap[key] = timestamppb.Now()
				changed = true
			}
		}

		// update status once if something changed
		if changed {
			updateCtx, cancelUpdate := context.WithTimeout(ctx, 10*time.Second)
			if _, err := acClient.UpdateOtacStatus(updateCtx, &accesscodepb.UpdateOneTimeAccessCodeStatusRequest{
				Id:     otac.GetId(),
				Status: otac.GetStatus(),
			}); err != nil {
				glog.Warningf("notifier: failed to update OTAC status for OTAC %q after sending notifications: %v", otacId, err)
			}
			cancelUpdate()
		}
	}

	glog.Info("notifier: finished OTAC expiry check")
	return nil
}

// parseWindows parses a comma-separated list like "3d,2d,1d" or "72h,48h,24h"
// into a map from original string key to parsed duration.
func parseWindows(raw string) (map[string]time.Duration, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, fmt.Errorf("notifier: NOTIFICATION_WINDOWS is empty")
	}

	parts := strings.Split(raw, ",")
	windows := make(map[string]time.Duration, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		// normalize "3d" -> "72h"
		norm, err := util.GetDurationWithDays(p)
		if err != nil {
			return nil, fmt.Errorf("notifier: invalid window %q: %w", p, err)
		}
		d, err := time.ParseDuration(norm)
		if err != nil {
			return nil, fmt.Errorf("notifier: invalid window %q after normalization (%q): %w", p, norm, err)
		}

		// keep original representation as key (e.g. "3d")
		windows[p] = d
	}

	if len(windows) == 0 {
		return nil, fmt.Errorf("notifier: no valid notification windows parsed from %q", raw)
	}

	return windows, nil
}

func parseDurationWithDays(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("notifier: duration is empty")
	}

	norm, err := util.GetDurationWithDays(s)
	if err != nil {
		return 0, fmt.Errorf("notifier: normalize duration %q: %w", s, err)
	}
	d, err := time.ParseDuration(norm)
	if err != nil {
		return 0, fmt.Errorf("notifier: failed to parse duration %q (normalized %q): %w", s, norm, err)
	}
	return d, nil
}

// sendExpiryMailForWindow builds a window-specific email and sends it via the email client.
func sendExpiryMailForWindow(
	emailClient *email.Client,
	to string,
	windowKey string, // e.g. "3d"
	remaining time.Duration,
	expiry time.Time,
	redeemed time.Time,
	maxDurationStr string,
	subjectTemplate string,
	bodyTemplate string,
) error {
	hoursLeft := int(remaining.Hours())
	if hoursLeft < 0 {
		hoursLeft = 0
	}

	// load locations; fall back to UTC if unavailable
	locNY, err := time.LoadLocation("America/New_York")
	if err != nil {
		return fmt.Errorf("notifier: failed to parse EST time: %w", err)
	}
	locBerlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return fmt.Errorf("notifier: failed to parse CET time: %w", err)
	}

	redeemedUTCStr := redeemed.UTC().Format(time.RFC1123)
	redeemedESTStr := redeemed.In(locNY).Format(time.RFC1123)
	redeemedCETStr := redeemed.In(locBerlin).Format(time.RFC1123)

	expiryUTCStr := expiry.UTC().Format(time.RFC1123)
	expiryESTStr := expiry.In(locNY).Format(time.RFC1123)
	expiryCETStr := expiry.In(locBerlin).Format(time.RFC1123)

	data := EmailTemplateData{
		RecipientEmail:                  to,
		RemainingDays:                   hoursLeft / 24,
		RemainingHours:                  hoursLeft % 24,
		RemainingHoursTotal:             hoursLeft,
		ActivationDateTimeUTC:           redeemedUTCStr,
		ExpirationDateTimeUTC:           expiryUTCStr,
		ActivationDateTimeEastern:       redeemedESTStr,
		ExpirationDateTimeEastern:       expiryESTStr,
		ActivationDateTimeCentralEurope: redeemedCETStr,
		ExpirationDateTimeCentralEurope: expiryCETStr,
		MaxDuration:                     maxDurationStr,
		WindowKey:                       windowKey,
	}

	subjTmpl, err := template.New("subject").Parse(subjectTemplate)
	if err != nil {
		return fmt.Errorf("notifier: parse subject template: %w", err)
	}

	bodyTmpl, err := template.New("body").Parse(bodyTemplate)
	if err != nil {
		return fmt.Errorf("notifier: parse body template: %w", err)
	}

	var subjectBuf, bodyBuf bytes.Buffer
	if err = subjTmpl.Execute(&subjectBuf, data); err != nil {
		return fmt.Errorf("notifier: execute subject template: %w", err)
	}
	if err = bodyTmpl.Execute(&bodyBuf, data); err != nil {
		return fmt.Errorf("notifier: execute body template: %w", err)
	}

	subject := subjectBuf.String()
	body := bodyBuf.String()

	return emailClient.Send(to, subject, body)
}

func ensureWindows(otac *accesscodepb.OneTimeAccessCode) map[string]*timestamppb.Timestamp {
	if otac.Status == nil {
		otac.Status = &accesscodepb.OneTimeAccessCodeStatus{}
	}
	if otac.Status.Notifications == nil {
		otac.Status.Notifications = &accesscodepb.OTACNotifications{}
	}
	if otac.Status.Notifications.Expiry == nil {
		otac.Status.Notifications.Expiry = &accesscodepb.ExpiryNotification{}
	}
	if otac.Status.Notifications.Expiry.Windows == nil {
		otac.Status.Notifications.Expiry.Windows = make(map[string]*timestamppb.Timestamp)
	}
	return otac.Status.Notifications.Expiry.Windows
}
