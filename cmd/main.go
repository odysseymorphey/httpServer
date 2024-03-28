package main

import (
	"github.com/odysseymorphey/httpServer/internal/server"
	"log"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	srv.Run()
}
