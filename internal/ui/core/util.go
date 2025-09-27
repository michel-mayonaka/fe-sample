package uicore

import (
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font"
    "image/color"
    "strconv"
    "ui_sample/internal/assets"
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
        if img, err := assets.LoadImage(c.Portrait); err == nil {
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

// DrawWrapped は文字列を最大幅 maxW で折り返して描画します。
// 戻り値は最後に描画した行のY座標（次の行を書く起点）です。
func DrawWrapped(dst *ebiten.Image, face font.Face, s string, x, y int, col color.Color, maxW, lineH int) int {
    if maxW <= 0 {
        text.Draw(dst, s, face, x, y, col)
        return y + lineH
    }
    // ルーン単位で貪欲折り返し
    line := ""
    for _, r := range s {
        trial := line + string(r)
        if text.BoundString(face, trial).Dx() > maxW {
            text.Draw(dst, line, face, x, y, col)
            y += lineH
            line = string(r)
            continue
        }
        line = trial
    }
    if line != "" {
        text.Draw(dst, line, face, x, y, col)
        y += lineH
    }
    return y
}
