package config

type TracingConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`
	ServiceID   int64  `env:"SERVICE_ID" envDefault:"dev"`
	ServiceName string `env:"SERVICE_NAME" envDefault:"default"`
	ServiceURL  string `env:"SERVICE_URL" envDefault:"dev"`
}
