package storage

import (
	es "github.com/AleksandrMac/goback2less1/pkg/environment/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Environment struct {
	id   string
	name string
}

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		db: db,
	}
}

func (x *Storage) Create(u *es.Environment) (id uuid.UUID, err error) {
	return
}

func (x *Storage) Read(id uuid.UUID) (*es.Environment, error) {
	return nil, nil
}

func (x *Storage) Update(id uuid.UUID, u *es.Environment, fields ...string) error {
	return nil
}

func (x *Storage) Delete(id uuid.UUID) error {
	return nil
}

func (x *Storage) GetList(limit, offset int64, filter string) ([]es.Environment, error) {
	return nil, nil
}
