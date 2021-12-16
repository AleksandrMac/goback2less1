package environmentuser

import (
	"github.com/AleksandrMac/goback2less1/pkg/environment-user/http"
	"github.com/AleksandrMac/goback2less1/pkg/environment-user/service"
	"github.com/AleksandrMac/goback2less1/pkg/environment-user/storage"

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
