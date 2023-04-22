package config

import "github.com/caarlos0/env/v8"

type AppConfig struct {
	LogConfig     *LogConfig
	TracingConfig *TracingConfig
}

func NewAppConfig() (*AppConfig, error) {
	config := &AppConfig{
		LogConfig:     &LogConfig{},
		TracingConfig: &TracingConfig{},
	}

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
