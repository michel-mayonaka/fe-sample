package adapter

import (
    gdata "ui_sample/internal/game/data"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/model"
    usr "ui_sample/internal/model/user"
)

// UnitFromUser はユーザキャラクタから表示用の uicore.Unit を構築します。
// 参照は gdata.Provider() 経由に限定し、ファイルI/Oへは依存しません。
// 画像読込は PortraitLoader に委譲します（テスト容易性のため）。
func UnitFromUser(c usr.Character, pl PortraitLoader) uicore.Unit {
    u := uicore.Unit{
        ID:    c.ID,
        Name:  c.Name,
        Class: c.Class,
        Level: c.Level,
        Exp:   c.Exp,
        HP:    c.HP,
        HPMax: c.HPMax,
        Stats: uicore.Stats(c.Stats),
        Weapon: uicore.WeaponRanks{
            Sword: c.Weapon.Sword,
            Lance: c.Weapon.Lance,
            Axe:   c.Weapon.Axe,
            Bow:   c.Weapon.Bow,
        },
        Magic: uicore.MagicRanks{
            Anima: c.Magic.Anima,
            Light: c.Magic.Light,
            Dark:  c.Magic.Dark,
            Staff: c.Magic.Staff,
        },
        Growth: uicore.Growth(c.Growth),
    }

    // HP 初期化（0 のときは Max を採用）
    if u.HP == 0 && u.HPMax > 0 {
        u.HP = u.HPMax
    }

    // テーブル取得（存在しない場合も考慮）
    var (
        wt *model.WeaponTable
        it *model.ItemDefTable
        uw []usr.OwnWeapon
        ui []usr.OwnItem
    )
    if p := gdata.Provider(); p != nil {
        // マスタ定義テーブル
        wt = p.WeaponsTable()
        it = p.ItemsTable()
        // ユーザ在庫スナップショット
        uw = p.UserWeapons()
        ui = p.UserItems()
    }

    // 装備はユーザ参照を優先して解決
    if len(c.Equip) > 0 {
        for _, er := range c.Equip {
            if er.UserWeaponsID != "" {
                for _, ow := range uw {
                    if ow.ID == er.UserWeaponsID {
                        name := ow.MstWeaponsID
                        if wt != nil {
                            if w, ok := wt.FindByID(ow.MstWeaponsID); ok { name = w.Name }
                        }
                        u.Equip = append(u.Equip, uicore.Item{ID: ow.ID, Name: name, Uses: ow.Uses, Max: ow.Max})
                        break
                    }
                }
            } else if er.UserItemsID != "" {
                for _, oi := range ui {
                    if oi.ID == er.UserItemsID {
                        name := oi.MstItemsID
                        if it != nil {
                            if d, ok := it.FindByID(oi.MstItemsID); ok { name = d.Name }
                        }
                        u.Equip = append(u.Equip, uicore.Item{ID: oi.ID, Name: name, Uses: oi.Uses, Max: oi.Max})
                        break
                    }
                }
            }
        }
    }

    // ポートレート
    if pl != nil && c.Portrait != "" {
        if img, err := pl.Load(c.Portrait); err == nil {
            u.Portrait = img
        }
    }
    return u
}
