module {{ cookiecutter.project_slug }}

go 1.20

require (
	github.com/caarlos0/env/v8 v8.0.0
	github.com/getsentry/sentry-go v0.20.0
	github.com/jackc/pgx/v4 v4.18.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/uptrace/opentelemetry-go-extra/otelsqlx v0.2.0
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.2.0 // indirect
	go.opentelemetry.io/otel/metric v0.38.1 // indirect
	go.opentelemetry.io/otel/trace v1.15.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)
