package adapter

import (
    uicore "ui_sample/internal/ui/core"
    "ui_sample/internal/model"
    gcore "ui_sample/pkg/game"
)

// UIToGame は UI ユニットを戦闘ロジックの Unit に変換します（先頭装備のみ考慮）。
func UIToGame(wt *model.WeaponTable, u uicore.Unit) gcore.Unit {
    var w model.Weapon
    if len(u.Equip) > 0 && wt != nil {
        if ww, ok := wt.Find(u.Equip[0].Name); ok {
            w = ww
        }
    }
    return gcore.Unit{
        ID: u.ID, Name: u.Name, Class: u.Class, Lv: u.Level,
        S: gcore.Stats{HP: u.HP, Str: u.Stats.Str, Skl: u.Stats.Skl, Spd: u.Stats.Spd, Lck: u.Stats.Lck, Def: u.Stats.Def, Res: u.Stats.Res, Mov: u.Stats.Mov},
        W: gcore.Weapon{MT: w.Might, Hit: w.Hit, Crit: w.Crit, Wt: w.Weight, RMin: w.RangeMin, RMax: w.RangeMax, Type: w.Type},
    }
}

