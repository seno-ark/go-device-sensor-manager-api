#!make
include .env

dev:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

swagger:
	swag init --parseDependency --parseInternal --dir "cmd/api,internal/api/v1" --output cmd/api/docs

postgres:
	docker run --rm --name postgres-alpine-mertani -p ${DB_PORT}:5432 -e POSTGRES_DB=${DB_NAME} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASS} postgres:16.1-alpine

migrate-file:
	migrate create -ext sql --dir pkg/database/migration -seq $(name)

migrate-up:
	migrate -path pkg/database/migration -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path pkg/database/migration -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

.PHONY: dev build swagger postgres migrate-file migrate-up migrate-down