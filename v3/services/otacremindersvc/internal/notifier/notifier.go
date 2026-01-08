package notifier

import (
	"context"
	"fmt"
	"strings"
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

func Run(
	ctx context.Context,
	cfg *config.Config,
	emailClient *email.Client,
	acClient accesscodepb.AccessCodeSvcClient,
	userClient userpb.UserSvcClient,
) error {
	// parse windows & job interval from config
	windows, windowKeys, err := parseWindows(cfg.NotificationWindowsRaw)
	if err != nil {
		return fmt.Errorf("parse notification windows: %w", err)
	}

	now := time.Now().UTC()
	ns := util.GetReleaseNamespace()

	glog.Info("starting OTAC expiry check",
		"now_utc", now,
		"namespace", ns,
		"windows_raw", cfg.NotificationWindowsRaw,
	)

	// list OTACs
	listCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	otacList, err := acClient.ListOtac(listCtx, &generalpb.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", hflabels.ScheduledEventLabel, cfg.ScheduledEventId),
	})
	if err != nil {
		return fmt.Errorf("list otacs: %w", err)
	}

	glog.Info("listed OTACs", "count", len(otacList.Otacs))

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
			glog.Warning("invalid RedeemedTimestamp on OTAC, skipping",
				"otac", otacId,
				"value", redeemedStr,
				"err", err,
			)
			continue
		}
		redeemed = redeemed.UTC()

		maxDur, err := parseDurationWithDays(maxDurationStr)
		if err != nil {
			glog.Warning("invalid MaxDuration on OTAC, skipping",
				"otac", otacId,
				"value", maxDurationStr,
				"err", err,
			)
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
			glog.Info("OTAC has empty spec.User, skipping",
				"otac", otacId,
			)
			continue
		}

		userCtx, cancelUser := context.WithTimeout(ctx, 10*time.Second)

		user, err := userClient.GetUserById(userCtx, &generalpb.GetRequest{Id: userID})
		cancelUser()

		if hferrors.IsGrpcNotFound(err) {
			glog.Info("user for OTAC not found, skipping",
				"otac", otacId,
				"userID", userID,
			)
			continue
		}
		if err != nil {
			glog.Warning("failed to get user for OTAC, skipping",
				"otac", otacId,
				"userID", userID,
				"err", err,
			)
			continue
		}

		emailAddr := strings.TrimSpace(user.GetEmail())
		if emailAddr == "" {
			glog.Warning("user has no email, skipping notification",
				"otac", otacId,
				"userID", userID,
			)
			continue
		}

		// per-window evaluation
		changed := false

		// Find the closest unsent window for which remaining <= w.
		// We need to do this, to prevent sending duplicate emails
		bestIdx := -1

		for i, w := range windows {
			key := windowKeys[i] // e.g. "3d", "2d", "1d"

			// already sent for this window?
			if _, ok := windowsMap[key]; ok {
				continue
			}

			if remaining <= w {
				// pick the smallest window that matches (closest to expiry)
				if bestIdx == -1 || w < windows[bestIdx] {
					bestIdx = i
				}
			}
		}

		if bestIdx != -1 {
			key := windowKeys[bestIdx]

			glog.Info("sending expiry email for window",
				"otac", otacId,
				"userID", userID,
				"email", emailAddr,
				"window", key,
				"remaining", remaining.String(),
				"expiry_utc", expiry,
			)

			if err := sendExpiryMailForWindow(emailClient, emailAddr, key, remaining, expiry, redeemed, maxDurationStr); err != nil {
				glog.Error("failed to send expiry email",
					"otac", otacId,
					"userID", userID,
					"email", emailAddr,
					"window", key,
					"err", err,
				)
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
				glog.Warning("failed to update OTAC status after notifications",
					"otac", otacId,
					"err", err,
				)
			}
			cancelUpdate()
		}
	}

	glog.Info("finished OTAC expiry check")
	return nil
}

// parseWindows parses a comma-separated list like "3d,2d,1d" or "72h,48h,24h"
// into durations + their original string keys.
func parseWindows(raw string) ([]time.Duration, []string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil, fmt.Errorf("NOTIFICATION_WINDOWS is empty")
	}

	parts := strings.Split(raw, ",")
	durations := make([]time.Duration, 0, len(parts))
	keys := make([]string, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		// normalize "3d" -> "72h"
		norm, err := util.GetDurationWithDays(p)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid window %q: %w", p, err)
		}
		d, err := time.ParseDuration(norm)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid window %q after normalization (%q): %w", p, norm, err)
		}

		durations = append(durations, d)
		keys = append(keys, p) // keep original representation as key (e.g. "3d")
	}

	if len(durations) == 0 {
		return nil, nil, fmt.Errorf("no valid notification windows parsed from %q", raw)
	}

	return durations, keys, nil
}

func parseDurationWithDays(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("duration is empty")
	}

	norm, err := util.GetDurationWithDays(s)
	if err != nil {
		return 0, fmt.Errorf("normalize duration %q: %w", s, err)
	}
	d, err := time.ParseDuration(norm)
	if err != nil {
		return 0, fmt.Errorf("parse duration %q (normalized %q): %w", s, norm, err)
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
) error {
	subject := fmt.Sprintf("Your one-time access expires in less than %s", windowKey)

	hoursLeft := int(remaining.Hours())
	if hoursLeft < 0 {
		hoursLeft = 0
	}

	body := fmt.Sprintf(`Hello @%s,

your one-time access was redeemed at:

    %s (UTC)

with a maximum duration of %s.
It will expire at:

    %s (UTC)

That is in approximately %d hour(s).

If you still need access, please log in again or request a new token.
`,
		to,
		redeemed.Format(time.RFC3339),
		maxDurationStr,
		expiry.Format(time.RFC3339),
		hoursLeft,
	)

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
