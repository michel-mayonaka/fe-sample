SHELL := /bin/bash

.PHONY: lint fmt

lint:
	golangci-lint run -c .golangci.yml

.PHONY: lint-list
lint-list:
	golangci-lint help linters || true

fmt:
	@command -v gofumpt >/dev/null 2>&1 && gofumpt -w . || true
	gofmt -s -w .
