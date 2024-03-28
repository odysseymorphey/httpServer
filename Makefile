.PHONY: build
build:
	go build -o server cmd/main.go

.PHONY: run
run:
	./server #--config ./config.yaml

.PHONY: clean
clean:
	rm ./server

.DEFAULT_GOAL = build