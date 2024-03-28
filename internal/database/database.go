package database

import (
	"github.com/go-pg/pg"
	"github.com/odysseymorphey/httpServer/internal/model"
)

type DB struct {
	DB *pg.DB
}

func NewDB() *DB {
	db := &DB{}
	db.Open()

	return db
}

func (db *DB) AddOrder(order model.Order) error {
	_, err := db.DB.Model(&order).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Open() {
	db.DB = pg.Connect(&pg.Options{
		User:     "wildberry",
		Password: "wildpass",
		Database: "wilddb",
	})
}

func (db *DB) Close() {
	db.DB.Close()
}
