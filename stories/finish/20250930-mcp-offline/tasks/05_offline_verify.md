# 05 検証: オフライン再現テスト

ステータス: [完了]

## 目的
- ベンダーモードと `MCP_OFFLINE` によりネットワーク無しで `make mcp` が完走することを確認。

## 手順（ローカル）
1. `.gomodcache` を削除または空にする。
2. `GOPROXY=off MCP_OFFLINE=1 make mcp` を実行。
3. 期待: `check-all`/`lint`（スキップ可）/`test-all` が成功し、終了コード 0。

## DoD
- 上記手順で成功することをスクリーンログに記録。

## 進捗ログ
- 2025-09-30: `GOPROXY=off MCP_OFFLINE=1 make mcp` 実行。オフラインで `vet/build/lint` は完走。`internal/usecase` のテストに既存失敗あり → 別ストーリー化。
- 2025-09-30: 追撃テストを安定化し、`make test-all` 成功を確認。
- 2025-09-30: 最終確認として `GOPROXY=off MCP_OFFLINE=1 make mcp` でグリーン（lint の typecheck 警告は非致命）。
