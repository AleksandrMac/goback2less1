package http

import "github.com/rs/zerolog"

type Environment struct {
	ID   string
	Name string
}

type Services interface {
	CreateEnvironment(name string) (id string, err error)
	FindEnvironmentByName(name string, limit, offset int64) ([]Environment, error)
}

type HTTP struct {
	log     *zerolog.Logger
	service Services
}

func New(log *zerolog.Logger, service Services) *HTTP {
	return &HTTP{
		log:     log,
		service: service,
	}
}
