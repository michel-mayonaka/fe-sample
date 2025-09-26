package ui

import (
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/user"
)

// SampleUnit はユーザテーブルから既定IDを読み、失敗時は内蔵値を返します。
func SampleUnit() Unit {
    if ut, err := user.LoadFromJSON("db/user/usr_characters.json"); err == nil {
        if uc, ok := ut.Find("iris"); ok { return unitFromUser(uc) }
    }
    u := Unit{
        Name:  "アイリス", Class: "ペガサスナイト", Level: 7, Exp: 56,
        HP: 22, HPMax: 26,
        Stats:  Stats{Str: 9, Mag: 0, Skl: 12, Spd: 14, Lck: 8, Def: 6, Res: 7, Mov: 7},
        Equip:  []Item{{Name: "アイアンランス", Uses: 35, Max: 45}, {Name: "ジャベリン", Uses: 12, Max: 20}, {Name: "傷薬", Uses: 3, Max: 3}},
        Weapon: WeaponRanks{Sword: "D", Lance: "B", Axe: "-", Bow: "-"},
        Magic:  MagicRanks{Anima: "-", Light: "-", Dark: "-", Staff: "-"},
        Growth: Growth{HP: 70, Str: 45, Mag: 10, Skl: 55, Spd: 65, Lck: 50, Def: 20, Res: 35, Mov: 0},
    }
    if img, _, err := ebitenutil.NewImageFromFile("assets/01_iris.png"); err == nil { u.Portrait = img }
    return u
}

func unitFromUser(c user.Character) Unit {
    u := Unit{
        ID: c.ID, Name: c.Name, Class: c.Class, Level: c.Level, Exp: c.Exp,
        HPMax: c.HPMax, Stats: Stats(c.Stats),
        Weapon: WeaponRanks{Sword: c.Weapon.Sword, Lance: c.Weapon.Lance, Axe: c.Weapon.Axe, Bow: c.Weapon.Bow},
        Magic:  MagicRanks{Anima: c.Magic.Anima, Light: c.Magic.Light, Dark: c.Magic.Dark, Staff: c.Magic.Staff},
        Growth: Growth(c.Growth),
    }
    for _, it := range c.Equip { u.Equip = append(u.Equip, Item{Name: it.Name, Uses: it.Uses, Max: it.Max}) }
    if c.Portrait != "" { if img, _, err := ebitenutil.NewImageFromFile(c.Portrait); err == nil { u.Portrait = img } }
    if u.HP == 0 && u.HPMax > 0 { u.HP = u.HPMax } else { u.HP = c.HP }
    return u
}

// LoadUnitsFromUser はユーザテーブル（usr_）から UI 用ユニット配列を生成します。
func LoadUnitsFromUser(path string) ([]Unit, error) {
    ut, err := user.LoadFromJSON(path)
    if err != nil { return nil, err }
    units := make([]Unit, 0, len(ut.ByID()))
    for _, c := range ut.Slice() { units = append(units, unitFromUser(c)) }
    return units, nil
}

