# 09_usecase_tests_inventory — ユースケース最小テスト（在庫）

## 目的・背景
- `EquipWeapon/EquipItem` の所有者移動・巻き戻し・保存の一貫性をフェイクリポジトリで検証。

## 作業項目（変更点）
- `internal/usecase` にテスト追加（`*_test.go`）。
- フェイク `UserRepo/InventoryRepo` を最小実装で用意（インメモリ）。

## 完了条件
- `go test ./...`（少なくとも `internal/usecase`）が成功。
- ハッピーパスと代表的な境界（同一装備の移譲/空スロット/存在しないID）を網羅。

## 影響範囲
- テストコードのみ。

## 手順
1) フェイクリポジトリ定義（同パッケージ or `_test` パッケージ）。
2) ケース: AがX装備、BがY装備 → BへX装備 → Bの旧装備YはAに戻る。
3) `make mcp` 後に `go test ./internal/usecase -v` 実行。

## 品質ゲート
- `make mcp`
- `go test ./...`

## 動作確認
- テストのアサーションが全て緑。

## ロールバック
- テストファイルを削除。

## Progress / Notes
- YYYY-MM-DD: 着手
- YYYY-MM-DD: 追加・成功

## 関連
- `docs/ARCHITECTURE.md` 8章

