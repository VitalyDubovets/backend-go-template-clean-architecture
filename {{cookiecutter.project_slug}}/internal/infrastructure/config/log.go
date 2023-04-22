package config

type LogConfig struct {
	LogLevel  string `env:"LOG_LEVEL" envDefault:"debug"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"json"`
}
