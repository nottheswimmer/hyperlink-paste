PKGS := $(shell go list ./internal/...)

# Go tools

GOPATH	  := $(HOME)/go
BIN_DIR   := $(GOPATH)/bin
COVER     := $(BIN_DIR)/cover
GOLINT  := $(BIN_DIR)/golint
LINTER    := $(BIN_DIR)/golangci-lint

$(COVER):
	@go get -u golang.org/x/tools/cmd/cover

$(LINTER):
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN_DIR) v1.23.7

$(GOLINT):
	@go get -u golang.org/x/lint/golint

## help: List available build targets and descriptions
.PHONY: help

help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Compile the application and output a binary in the bin folder
.PHONY: build

build:
	@go build -o bin/hyperlink_paste cmd/app/main.go

## lint: Check source code for style violations
.PHONY: lint

lint: $(GOLINT)
	@golint ./...

## run: Compile and run the application
.PHONY: run

run:
	@go run cmd/app/main.go

## test: Run the application's tests
.PHONY: test

test:
ifeq ($(with),cover)
	@go test $(PKGS) -coverprofile coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out
else
	@go test $(PKGS)
endif

## deps: Download dependencies to local cache
.PHONY: deps

deps:
	@-go mod download

## tidy: Add missing and/or remove unused modules
.PHONY: tidy

tidy:
	@-go mod tidy


## cover: Run tests and open a browser to an annotated coverage report
.PHONY: cover

cover: $(COVER)
	@go test $(PKGS) -coverprofile coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out
