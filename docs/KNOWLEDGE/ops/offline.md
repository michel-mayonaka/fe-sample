# オフライン実行ガイド（make mcp）

`make mcp`（vet/build/lint/test）をネットワーク無しで実行するための手順です。

## 事前準備（オンラインで一度だけ）
- 依存を `vendor/` に固定します。
```sh
make vendor-sync   # go mod tidy && go mod vendor
```
- 生成された `vendor/` をリポジトリにコミットしてください。

## 実行（オフライン環境）
- GOPROXY を無効化し、ベンダーモードで実行します。
```sh
# 最短スモーク（論理層のみ）
make smoke

# 包括（mcp をオフラインで実行）
MCP_OFFLINE=1 make offline

# 直接 mcp を実行する場合（同等）
GOPROXY=off MCP_OFFLINE=1 make mcp
```

## 挙動メモ
- `MCP_OFFLINE=1` のとき、`GOFLAGS='-mod=vendor'` が自動付与され、ネットワークに出ません。
- `golangci-lint` は未導入なら自動スキップします（導入済みでも解析中エラーは CI 非必須扱い）。
- `check-ui` は環境依存の既知エラー（JRE など）を検知した場合、既定ではスキップします。
- `.gomodcache` を空にする必要はありません（ベンダーモードが優先されます）。

## トラブルシューティング
- 「module lookup disabled by GOPROXY=off」が出る
  - `vendor/` が不足しています。オンラインで `make vendor-sync` を実行し、再度オフライン実行してください。
- Lint が import 解決に失敗する
  - 解析器の制約によるものです。`lint` は CI 非必須のため、`mcp` の成否には影響しません。

## 参考
- Makefile: `MCP_OFFLINE` 分岐と共通環境 `GOENV`
- コマンド: `make vendor-sync`, `make mcp`
