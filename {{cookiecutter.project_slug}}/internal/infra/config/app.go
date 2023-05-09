package config

type AppConfig struct {
	LogConfig                  *LogConfig
	SentryConfig               *SentryConfig
	TracingConfig              *TracingConfig
	ShutDownApplicationTimeout uint8 `env:"SHUTDOWN_APPLICATION_TIMEOUT" envDefault:"5"`
}

func NewAppConfig() (*AppConfig, error) {
	config := &AppConfig{
		LogConfig:     &LogConfig{},
		TracingConfig: &TracingConfig{},
		SentryConfig:  &SentryConfig{},
	}

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
