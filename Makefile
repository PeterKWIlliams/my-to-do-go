PROJECT_ROOT := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
DB_DIR ?= $(PROJECT_ROOT)
DB_NAME ?= mydb
DBFILE ?= $(DB_DIR)/$(DB_NAME).db
MIGRATION_DIR := sql/schema
GOOSE := goose

-include .env

.PHONY: all build migrate clean run
all: build migrate

migrate:$(DBFILE)
	@echo "Running database migrations"
	cd $(MIGRATION_DIR) && $(GOOSE) sqlite3 $(DBFILE) up && sqlc generate

build:
	@echo "Building application"
	go build -o my-to-do .

run: build
	@echo "Starting app"
	./my-to-do
run: 
	@echo "Starting app"
	./my-to-do

test:
	go test ./... -v
