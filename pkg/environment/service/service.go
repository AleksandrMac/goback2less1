package service

import (
	"fmt"

	eh "github.com/AleksandrMac/goback2less1/pkg/environment/http"

	"github.com/google/uuid"
)

type Environment struct {
	ID   uuid.UUID
	Name string
}

type CRUDL interface {
	Create(u *Environment) (id uuid.UUID, err error)
	Read(id uuid.UUID) (*Environment, error)
	// Update обновляет указанные в fields список полей, если fields = nil, происходит обновление всех полей.
	Update(id uuid.UUID, u *Environment, fields ...string) error
	Delete(id uuid.UUID) error
	GetList(limit, offset int64, filter string) ([]Environment, error)
}

type Service struct {
	storage CRUDL
}

func New(storage CRUDL) *Service {
	return &Service{
		storage: storage,
	}
}

func (x *Service) CreateEnvironment(name string) (id string, err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("service-create-environment: %w", err)
	}
	uid, err = x.storage.Create(&Environment{
		ID:   uid,
		Name: name,
	})
	return uid.String(), err
}

func (x *Service) FindEnvironmentByName(name string, limit, offset int64) ([]eh.Environment, error) {
	u, err := x.storage.GetList(limit, offset, fmt.Sprintf(`name LIKE '%s%%'`, name))
	if err != nil {
		return nil, fmt.Errorf(`service-find-environment-by-name: %w`, err)
	}
	outUser := make([]eh.Environment, len(u))
	for i := range u {
		outUser[i] = eh.Environment{
			ID:   u[i].ID.String(),
			Name: u[i].Name,
		}
	}
	return outUser, nil
}
