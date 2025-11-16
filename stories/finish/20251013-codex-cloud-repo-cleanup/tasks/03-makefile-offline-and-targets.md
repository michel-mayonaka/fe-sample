# 03 Make タスク再設計（オフライン/スモーク）

## 目的
ネットワークなしで再現性高く検証できるターゲット設計と、最小スモークチェックの定義を行う。

## 方針（ドラフト）
- 追加: `make smoke`（pkg と usecase の最小テストのみ）
- 追加: `make offline`（`-mod=vendor` でビルド/テスト）
- 既存: `make mcp` は非strict、UI 依存は自動スキップを維持

## 成功条件（DoD）
- オフライン状態でも `make smoke` が成功する設計案が提示される。

## 実装メモ（結果）
- 追加: `make smoke`, `make offline`, `make clean`, `make build-out`（出力 `out/` 集約）。
- `.gitignore` に `out/`/`site/` を追加。
- 検証: `make smoke` と `MCP_OFFLINE=1 make offline` が成功（2025-10-13 / 再検証 2025-11-16）。
