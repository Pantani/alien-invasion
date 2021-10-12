#! /usr/bin/make -f

# Project variables.
PACKAGE := github.com/Pantani/alien-invasion
VERSION := $(shell git describe --tags --always)
BUILD := $(shell git rev-parse --short HEAD)
DATETIME := $(shell date +"%Y.%m.%d-%H:%M:%S")
PROJECT_NAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOPKG := $(.)

# Go files
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=$(PACKAGE)/build.Version=$(VERSION) -X=$(PACKAGE)/build.Build=$(BUILD) -X=$(PACKAGE)/build.Date=$(DATETIME)"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## start: Start alien simulation from binary.
start: clean build
	@echo "  >  Starting $(PROJECT_NAME)"
	@-$(GOBIN)/$(PROJECT_NAME)

## build: Build the project.
build: go-get
	@echo "  >  Building $(PROJECT_NAME) binary..."
	GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECT_NAME) ./cmd

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECT_NAME) 2> /dev/null
	@echo "  >  Cleaning $(PROJECT_NAME) build cache"
	GOBIN=$(GOBIN) go clean

## test: Run all tests.
test: govet golint unit

## govet: Run go vet.
govet:
	@echo "  >  Running go vet"
	GOBIN=$(GOBIN) go vet ./...

## golint: Run golint.
golint:
	@echo "  >  Running golint"
	GOBIN=$(GOBIN) golint ./...

## test: Run unit tests.
unit:
	@echo "  >  Running unit tests"
	GOBIN=$(GOBIN) go test -race -tags=functional -v ./...

## fmt: Run `go fmt` for all go files.
fmt:
	@echo "  >  Format all go files"
	GOBIN=$(GOBIN) gofmt -w ${GOFMT_FILES}

## go-get: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	GOBIN=$(GOBIN) go get ./... $(get)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
