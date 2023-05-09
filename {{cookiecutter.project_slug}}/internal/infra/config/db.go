package config

import (
	"fmt"
)

type PostgreSQLConfig struct {
	Host               string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port               int    `env:"POSTGRES_PORT" envDefault:"5432"`
	User               string `env:"POSTGRES_USER" envDefault:"{{ cookiecutter.project_slug }}"`
	Password           string `env:"POSTGRES_PASSWORD" envDefault:"{{ cookiecutter.project_slug }}"`
	DBName             string `env:"POSTGRES_DB_NAME" envDefault:"{{ cookiecutter.project_slug }}"`
	MaxOpenConnections int    `env:"POSTGRES_MAX_OPEN_CONNECTIONS" envDefault:"10"`
}

func (c *PostgreSQLConfig) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
