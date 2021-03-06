package community

import (
	"github.com/AleksandrMac/goback2less1/pkg/community/http"
	"github.com/AleksandrMac/goback2less1/pkg/community/service"
	"github.com/AleksandrMac/goback2less1/pkg/community/storage"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

func NewHTTP(log *zerolog.Logger, service http.Services) *http.HTTP {
	return http.New(log, service)
}

func NewService(storage service.CRUDL) *service.Service {
	return service.New(storage)
}

func NewStorage(db *pgxpool.Pool) *storage.Storage {
	return storage.New(db)
}
