package migration

import (
	"github.com/go-pg/pg/orm"
	"github.com/odysseymorphey/httpServer/internal/database"
	"github.com/odysseymorphey/httpServer/internal/model"
)

func CreateSchema(db *database.DB) error {
	models := []interface{}{
		(*model.Order)(nil),
	}

	for _, modl := range models {
		op := orm.CreateTableOptions{}
		err := db.DB.Model(modl).CreateTable(&op)
		if err != nil {
			return err
		}
	}

	return nil
}
