package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	sc.Publish("addNewOrder", b)
}
