package main

import (
	"github.com/odysseymorphey/httpServer/internal/database"
	"github.com/odysseymorphey/httpServer/internal/migration"
	"log"
)

func main() {
	db := database.DB{}
	db.Open()
	defer db.Close()

	if err := migration.CreateSchema(&db); err != nil {
		log.Fatal(err)
	}
}
