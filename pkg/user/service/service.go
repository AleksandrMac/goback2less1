package service

import (
	"context"
	"fmt"

	uh "github.com/AleksandrMac/goback2less1/pkg/user/http"

	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
}

type CRUDL interface {
	Create(ctx context.Context, u *User) (id uuid.UUID, err error)
	Read(ctx context.Context, id uuid.UUID) (*User, error)
	// Update обновляет указанные в fields список полей, если fields = nil, происходит обновление всех полей.
	Update(ctx context.Context, id uuid.UUID, u *User, fields ...string) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetList(ctx context.Context, limit, offset int64, filter string) ([]User, error)
}

type Service struct {
	storage CRUDL
}

func New(storage CRUDL) *Service {
	return &Service{
		storage: storage,
	}
}

func (x *Service) CreateUser(ctx context.Context, name string) (id string, err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("service-create-user: %w", err)
	}
	uid, err = x.storage.Create(ctx, &User{
		ID:   uid,
		Name: name,
	})
	return uid.String(), err
}

func (x *Service) FindUserByName(ctx context.Context, name string, limit, offset int64) ([]uh.User, error) {
	u, err := x.storage.GetList(ctx, limit, offset, fmt.Sprintf(`name LIKE '%s%%'`, name))
	if err != nil {
		return nil, fmt.Errorf(`service-find-user-by-name: %w`, err)
	}
	outUser := make([]uh.User, len(u))
	for i := range u {
		outUser[i] = uh.User{
			ID:   u[i].ID.String(),
			Name: u[i].Name,
		}
	}
	return outUser, nil
}
