package scenes

// InventoryPort は装備変更など“更新系”のユースケース境界です。
type InventoryPort interface {
    // EquipWeapon はユニットの武器スロットにユーザ武器を装備します。
    EquipWeapon(unitID string, slot int, userWeaponID string) error
    // EquipItem はユニットのアイテムスロットにユーザアイテムを装備します。
    EquipItem(unitID string, slot int, userItemID string) error
}
