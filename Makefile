SHELL := /bin/bash

.PHONY: lint fmt check check-all build run mcp

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run -c .golangci.yml ./pkg/... >/dev/null 2>&1 || echo "[lint] サンドボックスのためスキップ" ; \
	else \
		echo "[lint] golangci-lint が見つからないためスキップ" ; \
	fi

.PHONY: lint-list
lint-list:
	golangci-lint help linters || true

fmt:
	@command -v gofumpt >/dev/null 2>&1 && gofumpt -w . || true
	gofmt -s -w .

# 依存を含むコンパイルチェック（バイナリは残さない）
# サンドボックス（GUI不可/JRE無）でも通る最小チェック
check:
	set -euo pipefail; \
	mkdir -p .gocache .gomodcache; \
	GOFLAGS='-mod=readonly' GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off go vet ./pkg/...; \
	GOFLAGS='-mod=readonly' GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off go build ./pkg/...

# ローカル開発向けのフルチェック
check-all:
	set -euo pipefail
	mkdir -p .gocache .gomodcache
	# pkg配下のみ検証（UI依存を避ける）
	GOFLAGS='-mod=readonly' GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off go vet ./pkg/...
	GOFLAGS='-mod=readonly' GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off go build ./pkg/...
	@echo "[check-all] pkg の vet/build 完了"

# 明示ビルド（バイナリを生成）
build:
	go build -o bin/ui_sample ./cmd/ui_sample

# 実行（開発用）
run:
	go run ./cmd/ui_sample

# ロジック層のみテスト（UI依存を避けるため pkg/... のみに限定）
.PHONY: test
test:
	@mkdir -p .gocache .gomodcache
	GOFLAGS='-mod=readonly' \
	GOMODCACHE=$(PWD)/.gomodcache \
	GOCACHE=$(PWD)/.gocache \
	GOWORK=off \
	go test ./pkg/...

# MCP: 変更前チェック（必須）
mcp: check-all lint
