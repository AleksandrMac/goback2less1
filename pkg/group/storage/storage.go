package storage

import (
	us "github.com/AleksandrMac/goback2less1/pkg/user/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type corporationGroup struct {
	id            string
	name          string
	environmentID string
}

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		db: db,
	}
}

func (x *Storage) Create(u *us.User) (id uuid.UUID, err error) {
	return
}

func (x *Storage) Read(id uuid.UUID) (*us.User, error) {
	return nil, nil
}

func (x *Storage) Update(id uuid.UUID, u *us.User, fields ...string) error {
	return nil
}

func (x *Storage) Delete(id uuid.UUID) error {
	return nil
}

func (x *Storage) GetList(limit, offset int64, filter string) ([]us.User, error) {
	return nil, nil
}
