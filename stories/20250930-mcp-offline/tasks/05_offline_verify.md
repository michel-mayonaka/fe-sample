# 05 検証: オフライン再現テスト

ステータス: [進行中]

## 目的
- ベンダーモードと `MCP_OFFLINE` によりネットワーク無しで `make mcp` が完走することを確認。

## 手順（ローカル）
1. `.gomodcache` を削除または空にする。
2. `GOPROXY=off MCP_OFFLINE=1 make mcp` を実行。
3. 期待: `check-all`/`lint`（スキップ可）/`test-all` が成功し、終了コード 0。

## DoD
- 上記手順で成功することをスクリーンログに記録。

## 進捗ログ
- 2025-09-30: `GOPROXY=off MCP_OFFLINE=1 make mcp` 実行。オフラインで `vet/build/lint` は完走。`internal/usecase` のテストで既存失敗（例: `TestRunBattleRound_AttackerDouble_ConsumesTwo`）。機能不具合に起因するため本ストーリーでは未対応、別ストーリー/バグとして切り出し提案。
