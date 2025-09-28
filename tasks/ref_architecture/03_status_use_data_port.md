# 03_status_use_data_port — Status画面の依存切替

## 目的・背景
- Status 画面の保存系操作（レベルアップ確定/明示保存）を `DataPort` へ委譲し、UI直書きを排除。

## 作業項目（変更点）
- `internal/game/scenes/status/status.go` の `PersistUnit` 呼び出しを `Env.Data` 経由に切替。
- `Env.App` 直接参照を残さない（コメント検索で確認）。

## 完了条件
- ビルド成功。
- レベルアップ後に保存が行われる（従来どおり）。

## 影響範囲
- Status シーンのみ。

## 手順
1) `status.go` 内の `s.E.App.PersistUnit(...)` を `s.E.Data.PersistUnit(...)` に置換。
2) 検索: `rg -n "PersistUnit\(" internal/game/scenes/status` で取りこぼしがないことを確認。
3) `make mcp`、手動でレベルアップ→戻るまで確認。

## 品質ゲート
- `make mcp`

## 動作確認
- レベルアップ→ポップアップ→確定→戻る→再度開いて値が保持されている。

## ロールバック
- 呼び出し箇所を `Env.App` に戻す。

## Progress / Notes
- YYYY-MM-DD: 着手
- YYYY-MM-DD: 切替完了・確認

## 関連
- `docs/ARCHITECTURE.md` 3章/6.2/12.3

