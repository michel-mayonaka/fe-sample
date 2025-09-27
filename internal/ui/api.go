package ui

import (
    "github.com/hajimehoshi/ebiten/v2"
    "math/rand"
    uicore "ui_sample/internal/ui/core"
    uipopup "ui_sample/internal/ui/popup"
    uiscreens "ui_sample/internal/ui/screens"
    uiwidgets "ui_sample/internal/ui/widgets"
    gcore "ui_sample/pkg/game"
)

// 型エイリアス（外部API互換）
type (
	Unit        = uicore.Unit
	Stats       = uicore.Stats
	Item        = uicore.Item
	WeaponRanks = uicore.WeaponRanks
	MagicRanks  = uicore.MagicRanks
	Growth      = uicore.Growth

	LevelUpGains = uipopup.LevelUpGains
)

// データ読み込み
func SampleUnit() Unit                              { return uicore.SampleUnit() }
func LoadUnitsFromUser(path string) ([]Unit, error) { return uicore.LoadUnitsFromUser(path) }

// 画面描画
func DrawCharacterList(dst *ebiten.Image, units []Unit, hover int) {
	uiscreens.DrawCharacterList(dst, units, hover)
}
func DrawStatus(dst *ebiten.Image, u Unit)        { uiscreens.DrawStatus(dst, u) }
func DrawBattle(dst *ebiten.Image, atk, def Unit) { uiscreens.DrawBattle(dst, atk, def) }
func DrawBattleWithTerrain(dst *ebiten.Image, atk, def Unit, attT, defT gcore.Terrain) {
    uiscreens.DrawBattleWithTerrain(dst, atk, def, attT, defT)
}

// 模擬戦API
func DrawSimulationBattle(dst *ebiten.Image, atk, def Unit, logs []string) {
	uiscreens.DrawSimulationBattle(dst, atk, def, logs)
}
func SimulateBattleCopy(atk, def Unit, rng *rand.Rand) (Unit, Unit, []string) {
	a, d, l := uiscreens.SimulateBattleCopy(atk, def, rng)
	return a, d, l
}
func SimBattleButtonRect(sw, sh int) (int, int, int, int) {
	return uiwidgets.SimBattleButtonRect(sw, sh)
}
func DrawSimBattleButton(dst *ebiten.Image, hovered, enabled bool) {
	uiwidgets.DrawSimBattleButton(dst, hovered, enabled)
}
func ListItemRect(sw, sh, i int) (int, int, int, int) { return uiscreens.ListItemRect(sw, sh, i) }

// ボタン
func BackButtonRect(sw, sh int) (int, int, int, int)    { return uiwidgets.BackButtonRect(sw, sh) }
func DrawBackButton(dst *ebiten.Image, hovered bool)    { uiwidgets.DrawBackButton(dst, hovered) }
func LevelUpButtonRect(sw, sh int) (int, int, int, int) { return uiwidgets.LevelUpButtonRect(sw, sh) }
func DrawLevelUpButton(dst *ebiten.Image, hovered, enabled bool) {
	uiwidgets.DrawLevelUpButton(dst, hovered, enabled)
}
func ToBattleButtonRect(sw, sh int) (int, int, int, int) { return uiwidgets.ToBattleButtonRect(sw, sh) }
func DrawToBattleButton(dst *ebiten.Image, hovered, enabled bool) {
	uiwidgets.DrawToBattleButton(dst, hovered, enabled)
}
func BattleStartButtonRect(sw, sh int) (int, int, int, int) {
	return uiscreens.BattleStartButtonRect(sw, sh)
}

// レベルアップポップアップ
func RollLevelUp(u Unit, rnd func() float64) LevelUpGains        { return uipopup.RollLevelUp(u, rnd) }
func ApplyGains(u *Unit, g LevelUpGains, cap int)                { uipopup.ApplyGains(u, g, cap) }
func DrawLevelUpPopup(dst *ebiten.Image, u Unit, g LevelUpGains) { uipopup.DrawLevelUpPopup(dst, u, g) }
