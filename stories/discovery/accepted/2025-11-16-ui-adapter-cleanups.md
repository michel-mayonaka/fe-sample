# 2025-11-16-ui-adapter-cleanups — UI adapter/bridge の後続クリーンアップ

- 作成: 2025-11-16 19:30:00 +0900
- 優先度: P2
- 関連ストーリー: 20250930-provider-ui-decouple

## 背景/課題
- `uicore` からの旧フォールバック（最小生成）を段階撤去し、`adapter` API へ完全移行する。
- 互換関数（`DrawAutoRunButton`/`BattleStartButtonRectCompat`）の整理とAPI一貫性の確保。
- 回帰テストの拡充（在庫タブ切替、装備確定、ログポップアップの確定操作）。

## 目的（完了の定義）
- [ ] `uicore.UnitFromUser` のフォールバック分岐（最小生成）を削除（bridge前提）。
- [ ] `uicore.LoadUnitsFromUser` を非推奨から削除（全呼び出しゼロを確認）。
- [ ] `ui/widgets` の旧互換API削除（`DrawAutoRunButton` と互換Rect廃止、`DrawAutoRunButtonAt` に一本化）。
- [ ] 回帰テスト追加（在庫タブ/装備確定/ログポップアップ確定）。

## メモ/方針
- 段階的にPRを分割（API削除→呼び出し置換→テスト追加）。
- 破壊的変更は minor バージョンアップ（tag運用がある場合）。

## 影響/リスク
- 互換APIの削除に伴う外部参照破断の可能性（現状サンプルのみで低）。

