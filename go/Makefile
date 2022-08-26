MIGRATE=/home/fg/tmp/migrate/./migrate
PGSQL_URL=postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable

build:
	@echo "  >  Building binary..."
	go build ./cmd/api

run:
	@echo "  >  Run app..."
	go run ./cmd/api/*.go

migrations_up:
	@echo "  >  migrations up..."
	$(MIGRATE) -path migrations -database "$(PGSQL_URL)" up

migrations_down:
	@echo "  >  migrations down..."
	$(MIGRATE) -path migrations -database "$(PGSQL_URL)" down

test:
	@echo "  >  Test..."
	go test -v -race -timeout 30s ./...