package config

type TracingConfig struct {
	Environment string `env:"TRACING_ENVIRONMENT" envDefault:"dev"`
	ServiceID   int64  `env:"TRACING_SERVICE_ID" envDefault:"1"`
	ServiceName string `env:"TRACING_SERVICE_NAME" envDefault:"default"`
	ServiceURL  string `env:"TRACING_SERVICE_URL" envDefault:""`
}
