package scenes

import (
    "github.com/hajimehoshi/ebiten/v2"
    "ui_sample/internal/assets"
    "ui_sample/internal/model"
    "ui_sample/internal/user"
)

// BuildWeaponRowsFromSnapshots は所持武器スナップショットと武器定義から行を構築します。
func BuildWeaponRowsFromSnapshots(owns []user.OwnWeapon, wt *model.WeaponTable) []WeaponRow {
    rows := make([]WeaponRow, 0, len(owns))
    for _, ow := range owns {
        name := ow.MstWeaponsID
        typ, rank := "", ""
        mt, hit, crt, wtVal, rmin, rmax := 0, 0, 0, 0, 1, 1
        if wt != nil {
            if w, ok := wt.FindByID(ow.MstWeaponsID); ok {
                name, typ, rank = w.Name, w.Type, w.Rank
                mt, hit, crt, wtVal = w.Might, w.Hit, w.Crit, w.Weight
                rmin, rmax = w.RangeMin, w.RangeMax
            }
        }
        rows = append(rows, WeaponRow{ID: ow.ID, Name: name, Type: typ, Rank: rank, Might: mt, Hit: hit, Crit: crt, Weight: wtVal, RangeMin: rmin, RangeMax: rmax, Uses: ow.Uses, Max: ow.Max})
    }
    return rows
}

// BuildItemRowsFromSnapshots は所持アイテムと定義から行を構築します。
func BuildItemRowsFromSnapshots(owns []user.OwnItem, it *model.ItemDefTable) []ItemRow {
    rows := make([]ItemRow, 0, len(owns))
    for _, oi := range owns {
        name := oi.MstItemsID
        typ, eff := "", ""
        pow := 0
        if it != nil {
            if d, ok := it.FindByID(oi.MstItemsID); ok {
                name, typ, eff, pow = d.Name, d.Type, d.Effect, d.Power
            }
        }
        rows = append(rows, ItemRow{ID: oi.ID, Name: name, Type: typ, Effect: eff, Power: pow, Uses: oi.Uses, Max: oi.Max})
    }
    return rows
}

// BuildWeaponRowsWithOwners は所有者バッジ情報付きの武器行を構築します。
func BuildWeaponRowsWithOwners(owns []user.OwnWeapon, wt *model.WeaponTable, ut *user.Table) []WeaponRow {
    rows := BuildWeaponRowsFromSnapshots(owns, wt)
    if ut == nil { return rows }
    own := map[string][]OwnerBadge{}
    for _, c := range ut.Slice() {
        for _, er := range c.Equip {
            if er.UserWeaponsID != "" {
                var img *ebiten.Image
                if c.Portrait != "" { if im, err := assets.LoadImage(c.Portrait); err == nil { img = im } }
                own[er.UserWeaponsID] = append(own[er.UserWeaponsID], OwnerBadge{Name: c.Name, Portrait: img})
            }
        }
    }
    for i := range rows { rows[i].Owners = own[rows[i].ID] }
    return rows
}

// BuildItemRowsWithOwners は所有者バッジ情報付きのアイテム行を構築します。
func BuildItemRowsWithOwners(owns []user.OwnItem, it *model.ItemDefTable, ut *user.Table) []ItemRow {
    rows := BuildItemRowsFromSnapshots(owns, it)
    if ut == nil { return rows }
    own := map[string][]OwnerBadge{}
    for _, c := range ut.Slice() {
        for _, er := range c.Equip {
            if er.UserItemsID != "" {
                var img *ebiten.Image
                if c.Portrait != "" { if im, err := assets.LoadImage(c.Portrait); err == nil { img = im } }
                own[er.UserItemsID] = append(own[er.UserItemsID], OwnerBadge{Name: c.Name, Portrait: img})
            }
        }
    }
    for i := range rows { rows[i].Owners = own[rows[i].ID] }
    return rows
}

