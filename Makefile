SHELL := /bin/bash

# --- Offline-friendly Go environment --------------------------------------
# 共通キャッシュとワークスペース設定
GOENV_CACHE := GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off

# 既定は読み取り専用モード（ネットワークに出ない前提だが vendor は使わない）
GOFLAGS_COMMON ?= -mod=readonly

# MCP_OFFLINE=1 のときは vendor を強制し、GOPROXY を無効化
ifdef MCP_OFFLINE
GOFLAGS_COMMON := -mod=vendor
export GOPROXY := off
endif

# go/vet/build/test 実行時の共通環境
GOENV := GOFLAGS='$(GOFLAGS_COMMON)' $(GOENV_CACHE)

.PHONY: lint lint-ci fmt check check-all check-ui build run mcp test test-all

lint:
	@mkdir -p .gocache .gomodcache
	@if command -v golangci-lint >/dev/null 2>&1; then \
		$(GOENV) golangci-lint run -c .golangci.yml ./... || echo "[lint] エラーがあります（CIは現状非必須）" ; \
	else \
		echo "[lint] golangci-lint が見つからないためスキップ" ; \
	fi

# CI 用: 失敗をスキップしない厳格実行
lint-ci:
	$(GOENV) golangci-lint run -c .golangci.yml $(EXTRA_LINTERS) ./...

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
	$(GOENV) go vet ./pkg/...; \
	$(GOENV) go build ./pkg/...

# ローカル開発向けのフルチェック
check-all:
	set -euo pipefail
	mkdir -p .gocache .gomodcache
	# ロジック層（必須）
	$(GOENV) go vet ./pkg/...
	$(GOENV) go build ./pkg/...
	@echo "[check-all] pkg の vet/build 完了"
	# UI依存もビルド確認（環境依存で失敗する場合は非strict時にスキップ）
	$(MAKE) check-ui

check-ui:
	@set -e; \
	FAILED=0; \
	MSG=""; \
	$(GOENV) go build -o /dev/null ./cmd/ui_sample 2> ._ui_build_err.log || FAILED=1; \
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
	$(GOENV) go test ./pkg/...

# すべてのユニットテスト（UI 依存を避けるため、pkg と internal/usecase に限定）
# 既定は -race -cover。重い場合は TEST_FLAGS="" で上書き可能。
TEST_FLAGS ?= -race -cover
TEST_TAGS  ?=
ifdef MCP_OFFLINE
TEST_TAGS := headless
endif
TEST_PKGS  ?= ./pkg/... ./internal/usecase
.PHONY: test-all
test-all:
	@mkdir -p .gocache .gomodcache
	$(GOENV) go test $(TEST_FLAGS) -tags "$(TEST_TAGS)" $(TEST_PKGS)

# UI 関連パッケージも含めたテスト（CI 追加ジョブ向け）
.PHONY: test-all-ui
TEST_PKGS_UI ?= ./pkg/... ./internal/usecase ./internal/game/service/... ./internal/game/ui/...
test-all-ui:
	@mkdir -p .gocache .gomodcache
	$(GOENV) go test $(TEST_FLAGS) -tags "$(TEST_TAGS)" $(TEST_PKGS_UI)

# MCP: 変更前チェック（必須）
# 既定はローカル: lint / CI: lint-ci
MCP_LINT_TARGET ?= lint
ifdef CI
MCP_LINT_TARGET := lint-ci
endif

# MCP: 変更前チェック + Lint + ユニットテスト（pkg / usecase）
mcp: check-all $(MCP_LINT_TARGET) test-all

# 依存を vendor に同期
.PHONY: vendor-sync
vendor-sync:
	go mod tidy
	go mod vendor

# --- Stories ---
.PHONY: new-story
new-story:
	@./scripts/new_story.sh "$(SLUG)"

.PHONY: finish-story
finish-story:
	@./scripts/finish_story.sh "$(SLUG)" "$(PATH)"
