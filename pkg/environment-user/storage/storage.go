package storage

import (
	eus "github.com/AleksandrMac/goback2less1/pkg/environment-user/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Environment struct {
	id   string
	name string
}

type User struct {
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

func (x *Storage) AddUser(userID, envID uuid.UUID) error {
	return nil
}

func (x *Storage) DelUser(userID, envID uuid.UUID) error {
	return nil
}

func (x *Storage) FindUserByEnvironment(name string) ([]eus.User, error) {
	return nil, nil
}

func (x *Storage) FindEnvironmentByUser(name string) ([]eus.Environment, error) {
	return nil, nil
}
