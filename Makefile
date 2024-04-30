build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

migrate_up:
	@go run cmd/goose/main.go up

migrate_status:
	@go run cmd/goose/main.go status

migrate_down:
	@go run cmd/goose/main.go down

migrate_fix:
	@go run cmd/goose/main.go fix

migrate_reset:
	@go run cmd/goose/main.go reset

migrate_version:
	@go run cmd/goose/main.go version

migrate_down_version:
	@go run cmd/goose/main.go version ${version}

generate_migrations:
	@goose -dir cmd/goose/migrations create ${name} ${type}
