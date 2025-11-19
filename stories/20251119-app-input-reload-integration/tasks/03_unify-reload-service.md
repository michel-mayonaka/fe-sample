# 03_統合 — ReloadService/メトリクス適用の一本化

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-19 02:06:50 +0900

## 目的
- `uicore.Metrics` への適用、`assets.Clear()`、ユニット再構築などホットリロード時の副作用を `internal/app/reload` に集約する。

## 完了条件（DoD）
- [ ] 共通の `reload.Trigger`（仮）が導入され、初期化時と Backspace 長押し時の処理が同じコードを通る。
- [ ] 副作用順序（Repo→Asset→Session→Metrics）がドキュメント化されている。
- [ ] docs/architecture/README.md にグローバルトグル/リロードの流れが追記されている。

## 進捗ログ
- 2025-11-19 02:06:50 +0900: タスク登録。
