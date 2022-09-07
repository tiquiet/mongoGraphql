.SILENT:
.PHONY: build
build:
	go build -v ./cmd/main

.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: generate
generate:
	go run github.com/99designs/gqlgen generate

.PHONY: db
db:
	docker run --name mongodb -d -p 27017:27017 mongo

.DEFAULT_GOAL := run