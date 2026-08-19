.PHONY: build run migrate-up migrate-down

build:
	go build -o bin/api ./cmd/api

run:
	./bin/api

migrate-up:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down
