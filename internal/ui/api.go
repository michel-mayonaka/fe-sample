package ui

import (
    "github.com/hajimehoshi/ebiten/v2"
    "math/rand"
    uicore "ui_sample/internal/ui/core"
    uipopup "ui_sample/internal/ui/popup"
    uiscreens "ui_sample/internal/ui/screens"
    uiwidgets "ui_sample/internal/ui/widgets"
    "ui_sample/internal/model"
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
func DrawBattleWithTerrain(dst *ebiten.Image, atk, def Unit, attT, defT gcore.Terrain, startEnabled bool) {
    uiscreens.DrawBattleWithTerrain(dst, atk, def, attT, defT, startEnabled)
}
func DrawBattleLogOverlay(dst *ebiten.Image, logs []string) { uiscreens.DrawBattleLogOverlay(dst, logs) }

// 模擬戦API
func SimulateBattleCopy(atk, def Unit, rng *rand.Rand) (Unit, Unit, []string) {
    a, d, l := uiscreens.SimulateBattleCopy(atk, def, rng)
    return a, d, l
}
func SimulateBattleCopyWithTerrain(atk, def Unit, attT, defT gcore.Terrain, rng *rand.Rand) (Unit, Unit, []string) {
    a, d, l := uiscreens.SimulateBattleCopyWithTerrain(atk, def, attT, defT, rng)
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

// 地形ボタンUI
func TerrainButtonRect(sw, sh int, left bool, idx int) (int, int, int, int) {
    return uiwidgets.TerrainButtonRect(sw, sh, left, idx)
}
func DrawTerrainButtons(dst *ebiten.Image, attSel, defSel int) {
    uiwidgets.DrawTerrainButtons(dst, attSel, defSel)
}

// レベルアップポップアップ
func RollLevelUp(u Unit, rnd func() float64) LevelUpGains        { return uipopup.RollLevelUp(u, rnd) }
func ApplyGains(u *Unit, g LevelUpGains, cap int)                { uipopup.ApplyGains(u, g, cap) }
func DrawLevelUpPopup(dst *ebiten.Image, u Unit, g LevelUpGains) { uipopup.DrawLevelUpPopup(dst, u, g) }

// 依存注入（武器テーブル共有）
func SetWeaponTable(wt *model.WeaponTable) { uiscreens.SetWeaponTable(wt) }

// 選択ポップアップ
func DrawChooseUnitPopup(dst *ebiten.Image, title string, units []Unit, hover int) {
    uipopup.DrawChooseUnitPopup(dst, title, units, hover)
}
func ChooseUnitItemRect(sw, sh, i, total int) (int, int, int, int) {
    return uipopup.ChooseUnitItemRect(sw, sh, i, total)
}
