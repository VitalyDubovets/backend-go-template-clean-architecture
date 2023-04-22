package sentry

import (
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/log"

	"github.com/getsentry/sentry-go"
)

func InitSentry(config *config.SentryConfig, log log.BaseLogger) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://examplePublicKey@o0.ingest.sentry.io/0",
		Debug:       config.Debug,
		Environment: config.Stage,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	return nil
}
