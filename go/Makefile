.PHONY: build
build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./cmd/api

.DEFAULT_GOAL := build

run:
	@echo "  >  Run app..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ./cmd/api/*.go

.DEFAULT_GOAL := run

migrations_up:
	@echo "  >  migrations up..."
	/home/fg/tmp/migrate/./migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

.DEFAULT_GOAL := migrations_up

migrations_down:
	@echo "  >  migrations down..."
	/home/fg/tmp/migrate/./migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

.DEFAULT_GOAL := migrations_up

.PHONY: test
test:
	@echo "  >  Test..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v -race -timeout 30s ./...