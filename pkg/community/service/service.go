package service

import (
	"github.com/google/uuid"
)

type Community struct {
	ID            uuid.UUID
	Name          string
	EnvironmentID uuid.UUID
}

type CRUDL interface {
	// Create(u *Project) (id uuid.UUID, err error)
	// Read(id uuid.UUID) (*Project, error)
	// // Update обновляет указанные в fields список полей, если fields = nil, происходит обновление всех полей.
	// Update(id uuid.UUID, u *Project, fields ...string) error
	// Delete(id uuid.UUID) error
	// GetList(limit, offset int64, filter string) ([]Project, error)
}

type Service struct {
	storage CRUDL
}

func New(storage CRUDL) *Service {
	return &Service{
		storage: storage,
	}
}
