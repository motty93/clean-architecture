package infrastructure

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewDatabaseConnection(dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	return conn, nil
}
