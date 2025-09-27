package uicore

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"strconv"
	"ui_sample/internal/user"
)

func Itoa(n int) string { return strconv.Itoa(n) }

func UnitFromUser(c user.Character) Unit {
	u := Unit{
		ID: c.ID, Name: c.Name, Class: c.Class, Level: c.Level, Exp: c.Exp,
		HPMax: c.HPMax, Stats: Stats(c.Stats),
		Weapon: WeaponRanks{Sword: c.Weapon.Sword, Lance: c.Weapon.Lance, Axe: c.Weapon.Axe, Bow: c.Weapon.Bow},
		Magic:  MagicRanks{Anima: c.Magic.Anima, Light: c.Magic.Light, Dark: c.Magic.Dark, Staff: c.Magic.Staff},
		Growth: Growth(c.Growth),
	}
	for _, it := range c.Equip {
		u.Equip = append(u.Equip, Item{Name: it.Name, Uses: it.Uses, Max: it.Max})
	}
	if c.Portrait != "" {
		if img, _, err := ebitenutil.NewImageFromFile(c.Portrait); err == nil {
			u.Portrait = img
		}
	}
	if u.HP == 0 && u.HPMax > 0 {
		u.HP = u.HPMax
	} else {
		u.HP = c.HP
	}
	return u
}

func LoadUnitsFromUser(path string) ([]Unit, error) {
	ut, err := user.LoadFromJSON(path)
	if err != nil {
		return nil, err
	}
	units := make([]Unit, 0, len(ut.ByID()))
	for _, c := range ut.Slice() {
		units = append(units, UnitFromUser(c))
	}
	return units, nil
}
