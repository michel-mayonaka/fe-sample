# 01 設計: オフライン方針と適用範囲

ステータス: [未着手]

## 目的
- `make mcp` をオフラインで完走させるための方式を決定（`vendor/` 採用、`-mod=vendor` 適用範囲、環境変数の扱い）。

## やること
- 既存 `Makefile` の `check/check-all/test/test-all` での `GOFLAGS`/`GOMODCACHE` 指定を棚卸し。
- `MCP_OFFLINE` 時の共通 `GOFLAGS`（`-mod=vendor`）と `GOPROXY=off` を適用するポイントを定義。
- UI ビルド分岐（`check-ui`）の挙動を確認（既存スキップ条件の維持）。

## DoD
- 方針メモを本ファイルに確定記載。
- 実装タスクに必要な TODO を列挙。

## 進捗ログ
- 2025-09-30: 作成。

