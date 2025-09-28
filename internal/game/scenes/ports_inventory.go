package scenes

import (
    "ui_sample/internal/repo"
)

// InventoryPort は在庫参照と装備変更のユースケース境界です。
type InventoryPort interface {
    // Inventory はユーザ在庫へのリポジトリを返します。
    Inventory() repo.InventoryRepo
    // EquipWeapon はユニットの武器スロットにユーザ武器を装備します。
    EquipWeapon(unitID string, slot int, userWeaponID string) error
    // EquipItem はユニットのアイテムスロットにユーザアイテムを装備します。
    EquipItem(unitID string, slot int, userItemID string) error
}
