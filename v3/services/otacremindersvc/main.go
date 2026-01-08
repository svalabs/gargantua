package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"

	config "github.com/hobbyfarm/gargantua/services/otacremindersvc/v3/internal/config"
	"github.com/hobbyfarm/gargantua/services/otacremindersvc/v3/internal/email"
	notifier "github.com/hobbyfarm/gargantua/services/otacremindersvc/v3/internal/notifier"
	"github.com/hobbyfarm/gargantua/v3/pkg/microservices"

	accesscodepb "github.com/hobbyfarm/gargantua/v3/protos/accesscode"
	userpb "github.com/hobbyfarm/gargantua/v3/protos/user"
)

var (
	serviceConfig *microservices.ServiceConfig
)

func init() {
	serviceConfig = microservices.BuildServiceConfig()
}

func main() {

	cfg, err := config.FromEnv()
	if err != nil {
		glog.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	services := []microservices.MicroService{
		microservices.AccessCode,
		microservices.User,
	}
	connections := microservices.EstablishConnections(services, serviceConfig.ClientCert)
	for _, conn := range connections {
		defer conn.Close()
	}

	acClient := accesscodepb.NewAccessCodeSvcClient(connections[microservices.AccessCode])
	userClient := userpb.NewUserSvcClient(connections[microservices.User])
	emailClient := &email.Client{
		Host:        cfg.SMTPHost,
		Port:        cfg.SMTPPort,
		User:        cfg.SMTPUser,
		Password:    cfg.SMTPPass,
		From:        cfg.SMTPFrom,
		RequireTLS:  cfg.RequireTLS,
		RequireAuth: cfg.RequireAuth,
		Timeout:     time.Duration(cfg.EmailTimeout) * time.Second,
		CC:          cfg.CC,
		Signature:   cfg.Signature,
		ReplyTo:     cfg.ReplyTo,
		MountPathCA: cfg.MountPathCA,
	}

	// graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.NotifierTimeout)*time.Second)
	defer cancel()

	if err := notifier.Run(ctx, cfg, emailClient, acClient, userClient); err != nil {
		glog.Error("notifier run failed", "err", err)
		os.Exit(1)
	}
}
