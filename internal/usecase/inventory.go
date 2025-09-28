package usecase

import (
    "ui_sample/internal/repo"
    "ui_sample/internal/user"
)

// Inventory は在庫リポジトリへのアクセサです（Scene層からの参照用）。
func (a *App) Inventory() repo.InventoryRepo { return a.Inv }

// EquipWeapon は指定のユーザ武器をスロットに装備し、既オーナーの装備を巻き戻します。
func (a *App) EquipWeapon(unitID string, slot int, userWeaponID string) error {
    if a == nil || a.Users == nil { return nil }
    t := a.Users.Table(); if t == nil { return nil }
    c, ok := t.Find(unitID); if !ok { return nil }
    var prev user.EquipRef
    if slot < len(c.Equip) { prev = c.Equip[slot] }
    // 既オーナーから外し、巻き戻す
    ownerID := ""; ownerSlot := -1
    for _, oc := range t.Slice() {
        for idx, er := range oc.Equip { if er.UserWeaponsID == userWeaponID { ownerID = oc.ID; ownerSlot = idx; break } }
        if ownerID != "" { break }
    }
    if ownerID != "" { if oc, ok2 := t.Find(ownerID); ok2 {
        for len(oc.Equip) <= ownerSlot { oc.Equip = append(oc.Equip, user.EquipRef{}) }
        oc.Equip[ownerSlot] = prev
        t.UpdateCharacter(oc)
    }}
    // 装備確定
    for len(c.Equip) <= slot { c.Equip = append(c.Equip, user.EquipRef{}) }
    c.Equip[slot] = user.EquipRef{UserWeaponsID: userWeaponID}
    t.UpdateCharacter(c)
    return a.Users.Save()
}

// EquipItem は指定のユーザアイテムをスロットに装備し、既オーナーの装備を巻き戻します。
func (a *App) EquipItem(unitID string, slot int, userItemID string) error {
    if a == nil || a.Users == nil { return nil }
    t := a.Users.Table(); if t == nil { return nil }
    c, ok := t.Find(unitID); if !ok { return nil }
    var prev user.EquipRef
    if slot < len(c.Equip) { prev = c.Equip[slot] }
    // 既オーナーから外し、巻き戻す
    ownerID := ""; ownerSlot := -1
    for _, oc := range t.Slice() {
        for idx, er := range oc.Equip { if er.UserItemsID == userItemID { ownerID = oc.ID; ownerSlot = idx; break } }
        if ownerID != "" { break }
    }
    if ownerID != "" { if oc, ok2 := t.Find(ownerID); ok2 {
        for len(oc.Equip) <= ownerSlot { oc.Equip = append(oc.Equip, user.EquipRef{}) }
        oc.Equip[ownerSlot] = prev
        t.UpdateCharacter(oc)
    }}
    // 装備確定
    for len(c.Equip) <= slot { c.Equip = append(c.Equip, user.EquipRef{}) }
    c.Equip[slot] = user.EquipRef{UserItemsID: userItemID}
    t.UpdateCharacter(c)
    return a.Users.Save()
}

