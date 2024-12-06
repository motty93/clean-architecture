package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDatabaseConnection(dsn string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), dsn)
}
