package adapter

import (
    "github.com/hajimehoshi/ebiten/v2"
    "ui_sample/internal/model"
    usr "ui_sample/internal/model/user"
    uiview "ui_sample/internal/game/ui/view"
)

// BuildWeaponRows はユーザ所持武器と定義テーブル/ユーザテーブルから
// 表示用の行データを構築します。
func BuildWeaponRows(owns []usr.OwnWeapon, wt *model.WeaponTable, ut *usr.Table, pl PortraitLoader) []uiview.WeaponRow {
    rows := make([]uiview.WeaponRow, 0, len(owns))
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
        rows = append(rows, uiview.WeaponRow{ID: ow.ID, Name: name, Type: typ, Rank: rank, Might: mt, Hit: hit, Crit: crt, Weight: wtVal, RangeMin: rmin, RangeMax: rmax, Uses: ow.Uses, Max: ow.Max})
    }
    if ut == nil { return rows }
    own := map[string][]uiview.OwnerBadge{}
    for _, c := range ut.Slice() {
        for _, er := range c.Equip {
            if er.UserWeaponsID == "" { continue }
            var img *ebiten.Image
            if pl != nil && c.Portrait != "" {
                if im, err := pl.Load(c.Portrait); err == nil { img = im }
            }
            own[er.UserWeaponsID] = append(own[er.UserWeaponsID], uiview.OwnerBadge{Name: c.Name, Portrait: img})
        }
    }
    for i := range rows { rows[i].Owners = own[rows[i].ID] }
    return rows
}

