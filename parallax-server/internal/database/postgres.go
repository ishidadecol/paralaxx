package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool() *pgxpool.Pool {
	databaseUrl :=
		"postgres://postgres:postgres@localhost:5432/parallax"

	pool, err :=
		pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	return pool
}
