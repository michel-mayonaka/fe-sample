# 02 実装: Makefile に MCP_OFFLINE 分岐を追加

ステータス: [完了]

## 目的
- `MCP_OFFLINE=1` 指定時、`go` コマンドがベンダーモードで依存解決しネットワークへ出ないようにする。

## 変更方針
- 共有変数 `GOFLAGS_COMMON` を導入し、通常は `-mod=readonly`、オフライン時は `-mod=vendor`。
- `GOPROXY=off`（保険）を `MCP_OFFLINE` 時にエクスポート。
- 既存の `GOMODCACHE`/`GOCACHE`/`GOWORK=off` は現状維持。

## 編集箇所（予定）
- `Makefile`: `check`/`check-all`/`test`/`test-all`/`test-all-ui` の `GOFLAGS` を `$(GOFLAGS_COMMON)` に置換。
- `Makefile`: 変数定義ブロックに `ifdef MCP_OFFLINE ... endif` を追加。

## DoD
- `MCP_OFFLINE=1 make mcp` が `vendor/` あり・`.gomodcache` 空でも成功。

## 進捗ログ
- 2025-09-30: Makefile に `GOENV`/`GOFLAGS_COMMON` を導入、主要ターゲットへ適用、`vendor-sync` 追加。

