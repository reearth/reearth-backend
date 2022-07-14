lint:
	golangci-lint run --fix

test:
	go test -race -v ./...

build:
	go build ./cmd/reearth

run-app:
	go run ./cmd/reearth

run-db:
	docker compose up -d reearth-mongo

gql:
	go generate ./internal/adapter/gql

.PHONY: lint test build run-app run-db gql
