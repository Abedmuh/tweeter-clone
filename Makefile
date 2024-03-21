IMAGE_NAME = crud-auth-go
PG_USERNAME := $(or $(PG_USERNAME), abdillah)
PG_PASSWORD := $(or $(PG_PASSWORD), pass)
PG_HOST := $(or $(PG_HOST), localhost)
PG_PORT := $(or $(PG_PORT), 5432)
PG_DATABASE := $(or $(PG_DATABASE), crud-auth-go)

build:
	go build -o main

migrate_up:
	migrate -path db/migrations -database "postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)?sslmode=disable" up

