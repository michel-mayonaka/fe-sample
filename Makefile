SHELL := /bin/bash

# --- Offline-friendly Go environment --------------------------------------
# 共通キャッシュとワークスペース設定
GOENV_CACHE := GOMODCACHE=$(PWD)/.gomodcache GOCACHE=$(PWD)/.gocache GOWORK=off

# 生成物の標準出力先を統一
OUT_DIR ?= out
BIN_DIR ?= $(OUT_DIR)/bin
COVER_DIR ?= $(OUT_DIR)/coverage
LOG_DIR ?= $(OUT_DIR)/logs

# 既定は読み取り専用モード（ネットワークに出ない前提だが vendor は使わない）
GOFLAGS_COMMON ?= -mod=readonly

# MCP_OFFLINE=1 のときは vendor を強制し、GOPROXY を無効化
ifdef MCP_OFFLINE
GOFLAGS_COMMON := -mod=vendor
export GOPROXY := off
endif

# go/vet/build/test 実行時の共通環境
GOENV := GOFLAGS='$(GOFLAGS_COMMON)' $(GOENV_CACHE)

.PHONY: lint lint-ci fmt check check-all check-ui build build-out run mcp test test-all smoke offline clean

lint:
	@mkdir -p .gocache .gomodcache
	@if command -v golangci-lint >/dev/null 2>&1; then \
		$(GOENV) golangci-lint run -c .golangci.yml --build-tags "$(TEST_TAGS)" $(LINT_PKGS) || echo "[lint] エラーがあります（CIは現状非必須）" ; \
	else \
		echo "[lint] golangci-lint が見つからないためスキップ" ; \
	fi

# CI 用: 失敗をスキップしない厳格実行
LINT_PKGS ?= ./pkg/... ./internal/usecase
lint-ci:
	$(GOENV) golangci-lint run -c .golangci.yml $(EXTRA_LINTERS) $(LINT_PKGS)

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
	mkdir -p $(LOG_DIR); \
	$(GOENV) go build -o /dev/null ./cmd/ui_sample 2> $(LOG_DIR)/ui_build_err.log || FAILED=1; \
	if [ $$FAILED -eq 0 ]; then \
		echo "[check-ui] cmd/ui_sample build OK"; \
		rm -f $(LOG_DIR)/ui_build_err.log; \
		exit 0; \
	fi; \
	MSG=$$(cat $(LOG_DIR)/ui_build_err.log); \
	rm -f $(LOG_DIR)/ui_build_err.log; \
	# 典型的な環境依存: プロキシ/Java/権限 に加え X11/GL 開発ヘッダ欠如も検出 \
	# Linux(Ubuntu) での例: X11/Xlib.h や GL/gl.h が無い場合。 \
	if echo "$$MSG" | grep -Eqi 'proxy\.golang\.org|Unable to locate a Java Runtime|operation not permitted|X11/Xlib\.h|GL/gl\.h|libX11|xorg|wayland|xcb|GLX|EGL'; then \
		echo "[check-ui] 環境依存のためスキップ: $$MSG"; \
		if [ "$$MCP_STRICT" = "1" ]; then \
			echo "[check-ui] MCP_STRICT=1 のため失敗扱い"; \
			echo "[check-ui] Linuxでの依存例: sudo apt-get install -y xorg-dev libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev"; \
			exit 1; \
		else \
			echo "[check-ui] 提案: 別ジョブ(ui-build-strict)で依存導入後に厳格ビルドを実行"; \
			exit 0; \
		fi; \
	else \
		echo "[check-ui] ビルドエラー:"; \
		echo "$$MSG"; \
		exit 1; \
	fi

## --- Build & Run ----------------------------------------------------------
# 明示ビルド（バイナリを生成）
build:
	go build -o bin/ui_sample ./cmd/ui_sample

# 標準出力先にビルド（互換ターゲットを保ちつつ out/bin/ に集約）
build-out:
	@mkdir -p $(BIN_DIR)
	$(GOENV) go build -o $(BIN_DIR)/ui_sample ./cmd/ui_sample

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

# 最小スモーク（オフライン前提の最短経路）
smoke:
	@echo "[smoke] offline vet/build/test (pkg + usecase)"
	@mkdir -p .gocache .gomodcache $(OUT_DIR)
	MCP_OFFLINE=1 $(MAKE) check
	MCP_OFFLINE=1 $(MAKE) test
	@echo "[smoke] OK"

# オフライン開発の包括ターゲット（vendor 前提）
offline:
	@echo "[offline] vendor(-mod=vendor) + headless test"
	@mkdir -p .gocache .gomodcache $(OUT_DIR)
	MCP_OFFLINE=1 $(MAKE) mcp
	@echo "[offline] OK"

# 生成物クリーン
clean:
	rm -rf $(OUT_DIR) .gocache .gomodcache coverage.out coverage.* *.prof || true

# 依存を vendor に同期
.PHONY: vendor-sync
vendor-sync:
	go mod tidy
	go mod vendor

# --- WASM / Pages site ------------------------------------------------------
WASM_OUT ?= site
WASM_BIN := $(WASM_OUT)/ui_sample.wasm
WASM_EXEC_JS := $(WASM_OUT)/wasm_exec.js
WASM_INDEX := $(WASM_OUT)/index.html
# Go 1.21+ では misc から lib に移動しているため両方を考慮
WASM_EXEC_SRC := $(shell if [ -f "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ]; then echo "$$(go env GOROOT)/lib/wasm/wasm_exec.js"; else echo "$$(go env GOROOT)/misc/wasm/wasm_exec.js"; fi)

.PHONY: wasm site site-clean serve-site

# WebAssembly ビルド（Ebitengine WebGL）
wasm:
	@mkdir -p $(WASM_OUT)
	@echo "[wasm] building $(WASM_BIN)"
	GOOS=js GOARCH=wasm $(GOENV) go build -ldflags "-s -w" -o $(WASM_BIN) ./cmd/ui_sample
	@# wasm_exec.js を Go 付属から複製
	cp $(WASM_EXEC_SRC) $(WASM_EXEC_JS)

# サイト生成（ローダHTMLとアセットを含む）
site: wasm
	@# ローダHTML（テンプレ）を配置
	@[ -f web/index.html ] && cp web/index.html $(WASM_INDEX) || echo "[site] web/index.html が無いためスキップ"
	@# 画像等のアセットを公開用にコピー
	@rm -rf $(WASM_OUT)/assets && mkdir -p $(WASM_OUT)/assets
	@if [ -f assets/01_iris.png ]; then cp assets/01_iris.png $(WASM_OUT)/assets/; else cp -R assets/* $(WASM_OUT)/assets/ 2>/dev/null || true; fi
	@# GitHub Pages の Jekyll 無効化
	touch $(WASM_OUT)/.nojekyll
	@echo "[site] generated in $(WASM_OUT)"

# 生成物クリア
site-clean:
	rm -rf $(WASM_OUT)

# 簡易サーバ（ローカル確認用）
serve-site: site
	@echo "[serve-site] http://localhost:8000 で提供します (CTRL+C で停止)"
	python3 -m http.server -d $(WASM_OUT) 8000
# --- Stories ---
.PHONY: new-story
new-story:
	@./scripts/new_story.sh "$(SLUG)"

.PHONY: finish-story
finish-story:
	@./scripts/finish_story.sh "$(SLUG)" "$(STORY_PATH)"

.PHONY: story-index
story-index:
	@./scripts/gen_story_index.sh

# --- Discovery ---
.PHONY: new-discovery
new-discovery:
	@./scripts/new_discovery.sh

.PHONY: promote-discovery
promote-discovery:
	@./scripts/promote_discovery.sh

.PHONY: decline-discovery
decline-discovery:
	@./scripts/decline_discovery.sh

.PHONY: consume-discovery
consume-discovery:
	@./scripts/consume_discovery.sh

.PHONY: backlog-index
backlog-index:
	@./scripts/gen_backlog.sh

.PHONY: discovery-index
discovery-index:
	@./scripts/gen_discovery_index.sh

.PHONY: validate-meta
validate-meta:
	@./scripts/validate_meta.sh || true
