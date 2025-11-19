# 15_usecase_split_files — Usecase のファイル分割

## 目的・背景
- 現在 `internal/usecase/app.go` に集約されている実装を、役割（Port）単位に分割し可読性・保守性を向上。
- 将来の拡張（Port追加/テーブルProvider拡張/テスト分離）に備え、変更局所性を高める。

## 作業項目（変更点）
- 追加: `internal/usecase/facade.go`（`type App` と `New(...)`、共通依存/ヘルパ）。
- 追加: `internal/usecase/data.go`（`DataPort` 実装: `ReloadData`, `PersistUnit`, `WeaponsTable`, `ItemsTable`）。
- 追加: `internal/usecase/battle.go`（`BattlePort` 実装: `RunBattleRound`）。
- 追加: `internal/usecase/inventory.go`（`InventoryPort` 実装: `Inventory`, `EquipWeapon`, `EquipItem`）。
- 既存: `internal/usecase/app.go` を段階的に空にし、最終的に削除（この工程では削除まで行う）。
- コンパイル保証: `var _ scenes.DataPort = (*App)(nil)` 等の型実装チェックを各ファイル先頭に置く（任意）。

## 完了条件
- `make mcp` 成功（`pkg`/`cmd/ui_sample` ビルドOK、lintは非必須）。
- 動作不変（起動・一覧/ステータス/在庫の挙動が従来通り）。

## 影響範囲
- `internal/usecase/` ディレクトリ配下のみ（呼び出し側の import/型は不変）。

## 手順
1) `facade.go` を作成し、`type App` と `New(...)` を移設。
2) `data.go` に `ReloadData/WeaponsTable/ItemsTable/PersistUnit` を移設。
3) `battle.go` に `RunBattleRound` を移設。
4) `inventory.go` に `Inventory/EquipWeapon/EquipItem` を移設。
5) `app.go` を削除（不要な import/循環がないことを確認）。
6) `make mcp` 実行。起動確認（`go run ./cmd/ui_sample`、環境によりスキップ可）。

## 品質ゲート
- `make mcp`

## 動作確認
- 起動し、一覧→ステータス→在庫→戻るが従来どおり機能。

## ロールバック
- 新規ファイルを削除し、`app.go` を元の内容で復元（コミット単位で巻き戻し可能）。

## Progress / Notes
- 2025-09-28: タスク追加
- 2025-09-28: 分割実施・mcp通過（facade/data/battle/inventoryへ分離）

## 関連
- `docs/architecture/README.md` 4章（Usecase実装の分割方針）
- `tasks/ref_architecture/index.md`
