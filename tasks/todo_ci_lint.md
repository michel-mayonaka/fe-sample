# TODO — ビルド/CI・Lint (2025-09-27)

- `.golangci.yml` と Makefile を `internal/...` も対象に（`revive/gofumpt/unused` 等を有効化）。
- CI（GitHub Actions）で `make mcp` 実行とキャッシュ、GUI不可環境のスキップ条件整備。
- README に Go 1.25 系の明示と `toolchain`/代替手順を追記。

