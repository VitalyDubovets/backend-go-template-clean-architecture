package config

type SentryConfig struct {
	Debug   bool   `env:"SENTRY_DEBUG" envDefault:"false"`
	DSN     string `env:"SENTRY_DSN" envDefault:""`
	Stage   string `env:"SENTRY_STAGE" envDefault:"dev"`
	Timeout uint8  `env:"SENTRY_TIMEOUT" envDefault:"5"`
}
