package service

import (
	"fmt"

	envHTTP "github.com/AleksandrMac/goback2less1/pkg/environment/http"
	usrHTTP "github.com/AleksandrMac/goback2less1/pkg/user/http"

	"github.com/google/uuid"
)

type Environment struct {
	ID   uuid.UUID
	Name string
}

type User struct {
	ID   uuid.UUID
	Name string
}

type CRUDL interface {
	AddUser(userID, envID uuid.UUID) error
	DelUser(userID, envID uuid.UUID) error
	FindUserByEnvironment(name string) ([]User, error)
	FindEnvironmentByUser(name string) ([]Environment, error)
}

type Service struct {
	storage CRUDL
}

func New(storage CRUDL) *Service {
	return &Service{
		storage: storage,
	}
}

func (x *Service) AddUser(userID, environmentID string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("service-add-user: %w", err)
	}
	envid, err := uuid.Parse(environmentID)
	if err != nil {
		return fmt.Errorf("service-add-user: %w", err)
	}
	return x.storage.AddUser(uid, envid)
}

func (x *Service) DelUser(userID, environmentID string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("service-del-user: %w", err)
	}
	envid, err := uuid.Parse(environmentID)
	if err != nil {
		return fmt.Errorf("service-del-user: %w", err)
	}
	return x.storage.DelUser(uid, envid)
}

func (x *Service) FindUserByEnvironment(name string) ([]usrHTTP.User, error) {
	usr, err := x.storage.FindUserByEnvironment(name)
	if err != nil {
		return nil, fmt.Errorf("service-find-user-by-environment: %w", err)
	}
	outUser := make([]usrHTTP.User, len(usr))
	for i := range usr {
		outUser[i] = usrHTTP.User{
			ID:   usr[i].ID.String(),
			Name: usr[i].Name,
		}
	}
	return outUser, nil
}

func (x *Service) FindEnvironmentByUser(name string) ([]envHTTP.Environment, error) {
	env, err := x.storage.FindEnvironmentByUser(name)
	if err != nil {
		return nil, fmt.Errorf("service-find-environment-by-user: %w", err)
	}
	outEnv := make([]envHTTP.Environment, len(env))
	for i := range env {
		outEnv[i] = envHTTP.Environment{
			ID:   env[i].ID.String(),
			Name: env[i].Name,
		}
	}
	return outEnv, nil
}
