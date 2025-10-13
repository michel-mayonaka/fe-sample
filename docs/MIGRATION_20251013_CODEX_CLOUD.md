# 既存開発者向け移行ノート — codex-cloud 対応（2025-10-13）

## 変更サマリ
- 生成物出力: `out/` 配下に統一（`out/bin`, `out/logs`, `out/coverage`）。
- 追加ターゲット: `make smoke`, `make offline`, `make clean`, `make build-out`。
- Lint 対象の明確化: 既定はロジック層（`./pkg/... ./internal/usecase`）。
- CI: `smoke-offline` ジョブを追加し、`build-and-lint` → `ui-build-strict` の順で実行。

## 置き換え表（よく使うコマンド）
- 旧: `make check` + `make test` → 新: `make smoke`（オフライン強制で最短）
- 旧: `make mcp`（オンライン前提） → 新: `MCP_OFFLINE=1 make offline`（vendor 前提）
- 旧: 個別ビルド出力（`bin/`） → 新: `make build-out`（`out/bin/`へ）
- 旧: 生成物掃除手作業 → 新: `make clean`

## 注意点
- オフライン前提のため、`vendor/` の更新はオンライン環境で `make vendor-sync` を実行し、コミットしてください。
- UI の厳格ビルドが必要な場合は Linux に X11/GL 依存を導入し、`MCP_STRICT=1 make check-ui` を使用。

## 参考
- codex-cloud ガイド: `docs/CODEX_CLOUD.md`
- オフラインガイド: `docs/OFFLINE.md`
- CI 設定: `.github/workflows/ci.yml`
