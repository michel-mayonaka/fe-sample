SHELL := /bin/bash

.PHONY: lint fmt check build run mcp

lint:
	golangci-lint run -c .golangci.yml

.PHONY: lint-list
lint-list:
	golangci-lint help linters || true

fmt:
	@command -v gofumpt >/dev/null 2>&1 && gofumpt -w . || true
	gofmt -s -w .

# 依存を含むコンパイルチェック（バイナリは残さない）
check:
	set -euo pipefail; \
	go vet ./...; \
	go build ./...; \
	GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -o /dev/null ./cmd/ui_sample

# 明示ビルド（バイナリを生成）
build:
	go build -o bin/ui_sample ./cmd/ui_sample

# 実行（開発用）
run:
	go run ./cmd/ui_sample

# MCP: 変更前チェック（必須）
mcp: check lint
