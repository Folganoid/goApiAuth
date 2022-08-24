.PHONY: build
build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build ./cmd/api

.DEFAULT_GOAL := build

run:
	@echo "  >  Run app..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ./cmd/api/*.go

.DEFAULT_GOAL := run

.PHONY: test
test:
	@echo "  >  Test..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v -race -timeout 30s ./...