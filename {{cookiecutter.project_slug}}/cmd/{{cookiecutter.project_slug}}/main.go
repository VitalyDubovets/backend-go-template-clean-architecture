package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	appshutdown "{{ cookiecutter.project_slug }}/internal/infrastructure/app_shutdown"
	"{{ cookiecutter.project_slug }}/internal/infrastructure/config"
	"{{ cookiecutter.project_slug }}/internal/infrastructure/infraerrors"
	"{{ cookiecutter.project_slug }}/internal/infrastructure/logger"
	sentryapp "{{ cookiecutter.project_slug }}/internal/infrastructure/sentry_app"
	"{{ cookiecutter.project_slug }}/internal/infrastructure/tracing"

	"go.opentelemetry.io/otel"
)

func main() {
	log := logger.GetLog()

	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(infraerrors.AppConfigError.WrapWithError(err))
	}

	log, flushFn, err := logger.NewZapLog(appConfig.LogConfig)
	if err != nil {
		log.Fatal(infraerrors.LogInitError.WrapWithError(err))
	}

	logger.ReplaceGlobal(log)

	err = sentryapp.InitSentry(appConfig.SentryConfig)
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	tp, err := tracing.InitTracerProvider(appConfig.TracingConfig)
	if err != nil {
		log.Fatalf("tracing.Init: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	otel.SetTracerProvider(tp)

	tr := tp.Tracer("component-main")

	ctx, span := tr.Start(ctx, "main")
	defer span.End()

	applicationTimeout := time.Duration(appConfig.ShutDownApplicationTimeout) * time.Second
	waitCh := appshutdown.ShutDownApplication(ctx, applicationTimeout, map[string]appshutdown.Handler{
		"sentry": func(ctx context.Context) error {
			trs := tp.Tracer("component-sentry-shutdown")

			ctx, spanSentry := trs.Start(ctx, "sentryShutdown")
			defer spanSentry.End()

			sentry.Flush(time.Duration(appConfig.SentryConfig.Timeout) * time.Second)
			return nil
		},
		"tracing": func(ctx context.Context) error {
			trs := tp.Tracer("component-tracing-shutdown")

			ctx, spanSentry := trs.Start(ctx, "tracingShutdown")
			defer spanSentry.End()

			return tp.Shutdown(ctx)
		},
		"logger": func(ctx context.Context) error {
			trs := tp.Tracer("component-logger-shutdown")

			ctx, spanSentry := trs.Start(ctx, "loggerShutdown")
			defer spanSentry.End()

			return flushFn()
		},
	})

	<-waitCh
}
