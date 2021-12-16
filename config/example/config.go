package config

import "github.com/caarlos0/env/v6"

type Config struct {
	DBURL string `env:"DB_URL" envDefault:"postgres://postgres:password@localhost:5432/testwfmt?sslmode=disable"`

	ServerPort      string `env:"SERVER_PORT" envDefault:"8000"`
	ServerTimeOut   int64  `env:"SERVER_TIMEOUT" envDefault:"30"`
	ServerGraceFull int64  `env:"SERVER_GRACEFULLTIME" envDefault:"5"` // second

	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func New() (*Config, error) {
	c := &Config{}
	err := env.Parse(c)
	return c, err
}
