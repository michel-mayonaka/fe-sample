SHELL := /bin/bash

.PHONY: lint lint-ci fmt check check-all check-ui build run mcp

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run -c .golangci.yml ./... || echo "[lint] エラーがあります（CIは現状非必須）" ; \
	else \
		echo "[lint] golangci-lint が見つからないためスキップ" ; \
	fi

# CI 用: 失敗をスキップしない厳格実行
lint-ci:
	golangci-lint run -c .golangci.yml $(EXTRA_LINTERS) ./...

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
	# ロジック層（必須）
	GOFLAGS='-mod=readonly' GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off go vet ./pkg/...
	GOFLAGS='-mod=readonly' GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off go build ./pkg/...
	@echo "[check-all] pkg の vet/build 完了"
	# UI依存もビルド確認（環境依存で失敗する場合は非strict時にスキップ）
	$(MAKE) check-ui

check-ui:
	@set -e; \
	FAILED=0; \
	MSG=""; \
	GOFLAGS='-mod=readonly' GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off go build -o /dev/null ./cmd/ui_sample 2> ._ui_build_err.log || FAILED=1; \
	if [ $$FAILED -eq 0 ]; then \
		echo "[check-ui] cmd/ui_sample build OK"; \
		rm -f ._ui_build_err.log; \
		exit 0; \
	fi; \
	MSG=$$(cat ._ui_build_err.log); \
	rm -f ._ui_build_err.log; \
	if echo "$$MSG" | grep -Eqi 'proxy\.golang\.org|Unable to locate a Java Runtime|operation not permitted'; then \
		echo "[check-ui] 環境依存のためスキップ: $$MSG"; \
		if [ "$$MCP_STRICT" = "1" ]; then \
			echo "[check-ui] MCP_STRICT=1 のため失敗扱い"; \
			exit 1; \
		else \
			exit 0; \
		fi; \
	else \
		echo "[check-ui] ビルドエラー:"; \
		echo "$$MSG"; \
		exit 1; \
	fi

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
# 既定はローカル: lint / CI: lint-ci
MCP_LINT_TARGET ?= lint
ifdef CI
MCP_LINT_TARGET := lint-ci
endif

mcp: check-all $(MCP_LINT_TARGET)
