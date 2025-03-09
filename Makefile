SHELL := /bin/bash

# Determine the absolute directory of the Makefile
MAKEFILE_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the Commodore 64 project
	go build ./...

test: ## Run the tests
	go test ./...

vet: ## Run the go vet tool
	go vet ./...

staticcheck: ## Run the staticcheck tool
	staticcheck ./...
