//go:build !headless

package uicore

import (
    "github.com/hajimehoshi/ebiten/v2"
    "golang.org/x/image/font"
    "image/color"
    "strconv"
    "ui_sample/internal/assets"
    usr "ui_sample/internal/model/user"
)

// Itoa は整数を10進文字列に変換します。
func Itoa(n int) string { return strconv.Itoa(n) }

// 以前の JSON 直読みキャッシュは廃止しました（Provider/Adapter 経由へ移行）。

// Bridge: adapter 層から登録されるフック。
// これにより uicore は adapter を直接 import せずに、adapter 実装へ委譲できます（循環参照回避）。
var (
    // UnitFromUserFunc は adapter.UnitFromUser を指す関数が登録されます。
    UnitFromUserFunc func(usr.Character) Unit
    // BuildUnitsFromProviderFunc は adapter.BuildUnitsFromProvider を指す関数が登録されます。
    BuildUnitsFromProviderFunc func() []Unit
)

// UnitFromUser はユーザのキャラクターレコードから表示用の Unit を構築します。
//
// Deprecated: Provider からの参照と UI 変換は
// internal/game/ui/adapter.UnitFromUser に集約してください。
// 本関数は当面の互換目的で残置します（JSON直読みの簡易キャッシュ）。
func UnitFromUser(c usr.Character) Unit {
    if UnitFromUserFunc != nil {
        return UnitFromUserFunc(c)
    }
    // Adapter のブリッジ未登録時は最小情報のみ（互換フォールバック）
    u := Unit{ID: c.ID, Name: c.Name, Class: c.Class, Level: c.Level, Exp: c.Exp, HP: c.HP, HPMax: c.HPMax,
        Stats: Stats(c.Stats), Weapon: WeaponRanks{Sword: c.Weapon.Sword, Lance: c.Weapon.Lance, Axe: c.Weapon.Axe, Bow: c.Weapon.Bow},
        Magic:  MagicRanks{Anima: c.Magic.Anima, Light: c.Magic.Light, Dark: c.Magic.Dark, Staff: c.Magic.Staff}, Growth: Growth(c.Growth)}
    if u.HP == 0 && u.HPMax > 0 { u.HP = u.HPMax }
    if c.Portrait != "" { if img, err := assets.LoadImage(c.Portrait); err == nil { u.Portrait = img } }
    return u
}

// LoadUnitsFromUser はユーザテーブルJSONから Unit の配列を構築します。
//
// Deprecated: 一覧生成は internal/game/ui/adapter.BuildUnitsFromProvider
// あるいは BuildUnitsFromUserTable を利用してください（Provider 経由）。
func LoadUnitsFromUser(string) ([]Unit, error) {
    if BuildUnitsFromProviderFunc != nil {
        return BuildUnitsFromProviderFunc(), nil
    }
    return nil, nil
}

// DrawWrapped は文字列を最大幅 maxW で折り返して描画します。
// 戻り値は最後に描画した行のY座標（次の行を書く起点）です。
func DrawWrapped(dst *ebiten.Image, face font.Face, s string, x, y int, col color.Color, maxW, lineH int) int {
	if maxW <= 0 {
		TextDraw(dst, s, face, x, y, col)
		return y + lineH
	}
	// ルーン単位で貪欲折り返し
	line := ""
	for _, r := range s {
		trial := line + string(r)
		if int(font.MeasureString(face, trial)>>6) > maxW {
			TextDraw(dst, line, face, x, y, col)
			y += lineH
			line = string(r)
			continue
		}
		line = trial
	}
	if line != "" {
		TextDraw(dst, line, face, x, y, col)
		y += lineH
	}
	return y
}
