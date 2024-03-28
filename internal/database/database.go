package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dsn = "postgres://wildberry:wildpass@localhost:5432/wilddb"

type DB struct {
	db     *pgxpool.Pool
	cancel context.CancelFunc
}

func NewDB() (*DB, error) {
	db := &DB{}

	ctx, cancel := context.WithCancel(context.Background())
	db.cancel = cancel

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	db.db = conn

	return db, nil
}
