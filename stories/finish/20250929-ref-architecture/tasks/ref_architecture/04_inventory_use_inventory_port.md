# 04_inventory_use_inventory_port — 在庫画面の依存切替

## 目的・背景
- 在庫参照・装備確定を `InventoryPort` に統一し、UI層からの永続化直書きを撤廃。

## 作業項目（変更点）
- `popup_weapon.go` / `popup_item.go` の装備確定処理を `EquipWeapon/EquipItem` 呼び出しに限定。
- `UserTable.UpdateCharacter/Save` 等のUI直更新を削除。
- 参照テーブルは `gdata.Provider()` で取得（Portへは載せない）。

## 完了条件
- ビルド成功、在庫→装備→一覧/ステータス反映の一連が従来通り動作。

## 影響範囲
- Inventory シーン配下のみ。

## 手順
1) `EquipWeapon/EquipItem` 呼び出しに統一（Env.Inv 経由ではなく Port 経由）。
2) 既オーナー巻き戻し/保存はユースケースに委譲（UIから削除）。
3) `make mcp`、装備移譲の手動確認（A→B装備→Bにあった装備がAへ戻る）。

## 品質ゲート
- `make mcp`

## 動作確認
- 在庫→武器/アイテム→スロット選択→装備→戻る→表示が整合。

## ロールバック
- 呼び出しを元のUI直更新に戻す（差分小）。

## Progress / Notes
- 2025-09-28: 着手
- 2025-09-28: 切替・mcp通過（在庫→Port）

## 関連
- `docs/ARCHITECTURE.md` 3章/6.1/7章/12.3
