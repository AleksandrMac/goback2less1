package http

import (
	envHTTP "github.com/AleksandrMac/goback2less1/pkg/environment/http"
	usrHTTP "github.com/AleksandrMac/goback2less1/pkg/user/http"
	"github.com/rs/zerolog"
)

type EnvironmentUser struct {
	EnvironmentID string
	UserID        string
}

type Services interface {
	AddUser(userID, environmentID string) error
	DelUser(userID, environmentID string) error
	FindUserByEnvironment(name string) ([]usrHTTP.User, error)
	FindEnvironmentByUser(name string) ([]envHTTP.Environment, error)
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
