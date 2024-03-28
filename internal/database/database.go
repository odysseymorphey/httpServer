package database

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/odysseymorphey/httpServer/internal/model"
)

type DB struct {
	DB     *pg.DB
	cancel context.CancelFunc
}

func NewDB() *DB {
	db := &DB{}

	db.DB = pg.Connect(&pg.Options{
		User:     "wildberry",
		Password: "wildpass",
		Database: "wilddb",
	})

	return db
}

func (db *DB) AddOrder(order model.Order) error {
	_, err := db.DB.Model(&order).Insert()
	if err != nil {
		return err
	}

	return nil
}
