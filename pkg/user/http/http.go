package http

import (
	"context"

	"github.com/rs/zerolog"
)

type User struct {
	ID   string
	Name string
}

type Services interface {
	CreateUser(ctx context.Context, name string) (id string, err error)
	FindUserByName(ctx context.Context, name string, limit, offset int64) ([]User, error)
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
