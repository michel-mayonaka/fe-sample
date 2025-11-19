# 06 codex-cloud 実行ガイド整備

## 目的
開発者が codex-cloud 環境で迷わず検証できるよう、手順と既知の注意点を記載する。

## 記載項目（ドラフト）
- 必要コマンド例: `make smoke`, `MCP_OFFLINE=1 make offline`
- 制約: ネットワーク不可/承認フロー/書込可能領域
- トラブルシュート: 依存解決に失敗、UI ビルドがスキップされる等

## 成功条件（DoD）
- README または `docs/KNOWLEDGE/ops/codex-cloud.md` に必要情報がまとまっている。

## 実装メモ（結果）
- 新規: `docs/KNOWLEDGE/ops/codex-cloud.md` を追加（環境変数、最短手順、出力/クリーン、トラブルシュート）。
- README に `smoke/offline` の使い方と CI フロー追記。
