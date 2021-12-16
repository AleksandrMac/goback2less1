package storage

import (
	"context"

	us "github.com/AleksandrMac/goback2less1/pkg/user/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type user struct {
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

func (x *Storage) Create(ctx context.Context, u *us.User) (id uuid.UUID, err error) {
	return
}

func (x *Storage) Read(ctx context.Context, id uuid.UUID) (*us.User, error) {
	return nil, nil
}

func (x *Storage) Update(ctx context.Context, id uuid.UUID, u *us.User, fields ...string) error {
	return nil
}

func (x *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (x *Storage) GetList(ctx context.Context, limit, offset int64, filter string) ([]us.User, error) {
	return nil, nil
}
