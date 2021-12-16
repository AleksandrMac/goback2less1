package http

import "github.com/rs/zerolog"

type Community struct {
	ID            string
	Name          string
	environmentID string
}

type Services interface {
}

type HTTP struct {
	log     *zerolog.Logger
	service Services
}

func New(log *zerolog.Logger, service Services) *HTTP {
	return &HTTP{
		service: service,
	}
}
