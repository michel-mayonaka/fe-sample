# 01_define_ports — Port骨子の追加

## 目的・背景
- Scene側に機能単位のPort（ユースケース境界）を定義し、Env肥大化を防ぐ土台を作る。
- Data/Battle/Inventory の最小契約を分離（実装は本工程では行わない）。

## 作業項目（変更点）
- `internal/game/scenes/ports_data.go` に `DataPort` を定義。
- `internal/game/scenes/ports_battle.go` に `BattlePort` を定義。
- `internal/game/scenes/ports_inventory.go` に `InventoryPort` を定義。
- （任意）`ports.go` に `UseCases`（合成）を定義し、段階移行の受け皿に。
- 既存の `UseCases`（Env内）の定義はこの工程では維持（互換目的）。

## 完了条件
- ビルド成功（Port追加のみで既存実装に影響しない）。
- 各PortのGoDocが1行説明を持つ。

## 影響範囲
- 型定義ファイルのみ。既存呼び出し・実装・配線なし。

## 手順
1) `ports_data.go`/`ports_battle.go`/`ports_inventory.go` を追加。
2) DataPort: `ReloadData() error`, `PersistUnit(u uicore.Unit) error`。
3) BattlePort: `RunBattleRound(units []uicore.Unit, selIndex int, attT, defT gcore.Terrain) (..., error)`。
4) InventoryPort: `Inventory() repo.InventoryRepo`, `EquipWeapon(...)`, `EquipItem(...)`。
5) それぞれ `package scenes` に置く（使う側に置く原則）。

## 品質ゲート
- `make mcp`（check-all + lint）。

## 動作確認
- なし（定義のみ）。`go build ./...` が通ればOK。

## ロールバック
- 追加ファイルを削除し、元の状態へ戻すのみ。

## Progress / Notes
- 2025-09-28: 着手
- 2025-09-28: Port定義追加・mcp通過

## 関連
- `docs/architecture/README.md` 3章/4章/12章
