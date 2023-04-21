package config

import "github.com/caarlos0/env/v8"

type AppConfig struct {
	LogLevel  string `env:"LOG_LEVEL" envDefault:"debug"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"json"`
}

func NewAppConfig() (*AppConfig, error) {
	config := new(AppConfig)

	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
