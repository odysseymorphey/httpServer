.PHONY: build
build:
	go build -o server cmd/main.go

.PHONY: run
run:
	./server

.PHONY: pub
pub:
	go run cmd/publisher/main.go ./model.json

.PHONY: migrate
migrate:
	go run cmd/migrations/main.go

.PHONY: clean
clean:
	rm ./server

.DEFAULT_GOAL = build