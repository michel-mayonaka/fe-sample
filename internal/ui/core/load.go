package uicore

import "github.com/hajimehoshi/ebiten/v2/ebitenutil"

func SampleUnit() Unit {
	if ut, err := LoadUnitsFromUser("db/user/usr_characters.json"); err == nil && len(ut) > 0 {
		for _, u := range ut {
			if u.ID == "iris" {
				return u
			}
		}
		return ut[0]
	}
	u := Unit{
		Name: "アイリス", Class: "ペガサスナイト", Level: 7, Exp: 56,
		HP: 22, HPMax: 26,
		Stats:  Stats{Str: 9, Mag: 0, Skl: 12, Spd: 14, Lck: 8, Def: 6, Res: 7, Mov: 7},
		Equip:  []Item{{Name: "アイアンランス", Uses: 35, Max: 45}, {Name: "ジャベリン", Uses: 12, Max: 20}, {Name: "傷薬", Uses: 3, Max: 3}},
		Weapon: WeaponRanks{Sword: "D", Lance: "B", Axe: "-", Bow: "-"},
		Magic:  MagicRanks{Anima: "-", Light: "-", Dark: "-", Staff: "-"},
		Growth: Growth{HP: 70, Str: 45, Mag: 10, Skl: 55, Spd: 65, Lck: 50, Def: 20, Res: 35, Mov: 0},
	}
	if img, _, err := ebitenutil.NewImageFromFile("assets/01_iris.png"); err == nil {
		u.Portrait = img
	}
	return u
}
