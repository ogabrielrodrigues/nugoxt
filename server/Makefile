.PHONY: default run test migration_up migration_down migration_fix

DB_URL=postgres://goshop:goshopdev@localhost:5432/goshop_dev?sslmode=disable

default: run

run:
	@go run cmd/main.go

test:
	@go test ./internal/model/*

migration_up: 
	migrate -path database/migration/ -database $(DB_URL) -verbose up

migration_down: 
	migrate -path database/migration/ -database $(DB_URL) -verbose down

migration_fix: 
	migrate -path database/migration/ -database $(DB_URL) force VERSION