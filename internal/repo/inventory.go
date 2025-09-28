// Package repo は永続層（JSONバックエンド）の軽量リポジトリ実装を提供します。
package repo

import (
    "ui_sample/internal/model"
    "ui_sample/internal/user"
)

// InventoryRepo は「在庫」集約への最小アクセスです。
//
// 集約の考え方:
// - usr_weapons と usr_items を横断管理し、同一の保存境界として扱います。
// - 一覧UI/耐久消費など、武器とアイテムをまたぐ操作を一貫して扱うための設計です。
// - 将来のDB移行時も単一トランザクションへ収めやすくする狙いがあります。
// - 装備付け替えなどの所有者移譲ロジックは Usecase 側に置き、Repo は
//   スナップショット取得と保存・再読み込みに責務を限定します。
type InventoryRepo interface {
    // Consume はIDに応じて耐久を count 減らします（下限0）。
    Consume(id string, count int) error
    // Save は在庫（武器・アイテム）を保存します（JSON/将来はDBトランザクション）。
    Save() error
    // Reload は在庫スナップショットを再読み込みします。
    Reload() error
    // Snapshots
    Weapons() []user.OwnWeapon
    Items() []user.OwnItem
}

// JSONInventoryRepo は usr_weapons.json / usr_items.json を束ねて扱う軽量実装です。
// Weapons()/Items() はコピーを返し、呼び出し側での破壊的変更を防ぎます。
// Consume/Save/Reload は在庫集約として同一境界で動作します。
type JSONInventoryRepo struct {
    weaponsPath string
    itemsPath   string
    weapons     []user.OwnWeapon
    items       []user.OwnItem
    // 参照用（結合時に利用）
    wt *model.WeaponTable
    it *model.ItemDefTable
}

// NewJSONInventoryRepo はユーザ在庫（武器/アイテム）とマスタ定義を読み込みます。
func NewJSONInventoryRepo(usrWeaponsPath, usrItemsPath, mstWeaponsPath, mstItemsPath string) (*JSONInventoryRepo, error) {
    w, err := user.LoadUserWeaponsJSON(usrWeaponsPath)
    if err != nil { return nil, err }
    i, err := user.LoadUserItemsJSON(usrItemsPath)
    if err != nil { return nil, err }
    wt, _ := model.LoadWeaponsJSON(mstWeaponsPath)
    it, _ := model.LoadItemsJSON(mstItemsPath)
    return &JSONInventoryRepo{weaponsPath: usrWeaponsPath, itemsPath: usrItemsPath, weapons: w, items: i, wt: wt, it: it}, nil
}

// Weapons は所持武器のスナップショットを返します（コピー）。
func (r *JSONInventoryRepo) Weapons() []user.OwnWeapon { return append([]user.OwnWeapon(nil), r.weapons...) }
// Items は所持アイテムのスナップショットを返します（コピー）。
func (r *JSONInventoryRepo) Items() []user.OwnItem { return append([]user.OwnItem(nil), r.items...) }

// Consume は指定IDの耐久を減らします（下限0）。
func (r *JSONInventoryRepo) Consume(id string, count int) error {
    if count <= 0 { return nil }
    if len(id) >= 3 && id[:3] == "uw_" {
        for i := range r.weapons {
            if r.weapons[i].ID == id {
                if count > r.weapons[i].Uses { count = r.weapons[i].Uses }
                r.weapons[i].Uses -= count
                break
            }
        }
        return nil
    }
    if len(id) >= 3 && id[:3] == "ui_" {
        for i := range r.items {
            if r.items[i].ID == id {
                if count > r.items[i].Uses { count = r.items[i].Uses }
                r.items[i].Uses -= count
                break
            }
        }
        return nil
    }
    return nil
}

// Save は在庫JSONファイルへ保存します。
func (r *JSONInventoryRepo) Save() error {
    if err := user.SaveUserWeaponsJSON(r.weaponsPath, r.weapons); err != nil { return err }
    if err := user.SaveUserItemsJSON(r.itemsPath, r.items); err != nil { return err }
    return nil
}

// Reload は在庫とマスタを再読み込みします。
func (r *JSONInventoryRepo) Reload() error {
    w, err := user.LoadUserWeaponsJSON(r.weaponsPath)
    if err != nil { return err }
    i, err := user.LoadUserItemsJSON(r.itemsPath)
    if err != nil { return err }
    r.weapons, r.items = w, i
    wt, _ := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    it, _ := model.LoadItemsJSON("db/master/mst_items.json")
    r.wt, r.it = wt, it
    return nil
}
