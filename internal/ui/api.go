// Package ui は UI レイヤの公開API（薄いファサード）を提供します。
package ui

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "math/rand"
    "encoding/json"
    "os"
    uicore "ui_sample/internal/ui/core"
    uipopup "ui_sample/internal/ui/popup"
    uiscreens "ui_sample/internal/ui/screens"
    uiwidgets "ui_sample/internal/ui/widgets"
    "ui_sample/internal/model"
    "ui_sample/internal/user"
    "ui_sample/internal/assets"
    gcore "ui_sample/pkg/game"
)

// Unit は UI 公開用の型エイリアスです。
type Unit = uicore.Unit
// Stats は UI 公開用の型エイリアスです。
type Stats = uicore.Stats
// Item は UI 公開用の型エイリアスです。
type Item = uicore.Item
// WeaponRanks は UI 公開用の型エイリアスです。
type WeaponRanks = uicore.WeaponRanks
// MagicRanks は UI 公開用の型エイリアスです。
type MagicRanks = uicore.MagicRanks
// Growth は UI 公開用の型エイリアスです。
type Growth = uicore.Growth

// LevelUpGains はレベルアップ結果の型エイリアスです。
type LevelUpGains = uipopup.LevelUpGains
// Weapon は武器定義の型エイリアスです。
type Weapon = model.Weapon
// ItemDef はアイテム定義の型エイリアスです。
type ItemDef = model.ItemDef
// WeaponRow はUIに表示する武器行の型エイリアスです。
type WeaponRow = uiscreens.WeaponRow
// ItemRow はUIに表示するアイテム行の型エイリアスです。
type ItemRow = uiscreens.ItemRow
// OwnerBadge は所有者バッジの型エイリアスです。
type OwnerBadge = uiscreens.OwnerBadge

// SampleUnit はサンプルユニットを返します。
func SampleUnit() Unit                              { return uicore.SampleUnit() }
// LoadUnitsFromUser はユーザJSONからユニット一覧を読み込みます。
func LoadUnitsFromUser(path string) ([]Unit, error) { return uicore.LoadUnitsFromUser(path) }

// DrawCharacterList はユニット一覧画面を描画します。
func DrawCharacterList(dst *ebiten.Image, units []Unit, hover int) {
	uiscreens.DrawCharacterList(dst, units, hover)
}
// DrawStatus はステータス画面を描画します。
func DrawStatus(dst *ebiten.Image, u Unit)        { uiscreens.DrawStatus(dst, u) }
// DrawBattle は旧版バトル画面（互換）を描画します。
func DrawBattle(dst *ebiten.Image, atk, def Unit) { uiscreens.DrawBattle(dst, atk, def) }
// DrawBattleWithTerrain は左右の地形を考慮したバトルプレビューを描画します。
func DrawBattleWithTerrain(dst *ebiten.Image, atk, def Unit, attT, defT gcore.Terrain, startEnabled bool) {
    uiscreens.DrawBattleWithTerrain(dst, atk, def, attT, defT, startEnabled)
}
// DrawBattleLogOverlay は全画面のログポップアップを描画します。
func DrawBattleLogOverlay(dst *ebiten.Image, logs []string) { uiscreens.DrawBattleLogOverlay(dst, logs) }
// DrawBattleLogs は画面下部のログオーバーレイを描画します。
func DrawBattleLogs(dst *ebiten.Image, logs []string)        { uiscreens.DrawBattleLogs(dst, logs) }
// DrawBattleLogOverlayScroll はログオーバーレイを末尾から offset 行スクロールして描画します。
func DrawBattleLogOverlayScroll(dst *ebiten.Image, logs []string, offset int) {
    uiscreens.DrawBattleLogOverlayScroll(dst, logs, offset)
}

// SimulateBattleCopy は地形なしの1ラウンド模擬戦を行います（永続化なし）。
func SimulateBattleCopy(atk, def Unit, rng *rand.Rand) (Unit, Unit, []string) {
    a, d, l := uiscreens.SimulateBattleCopy(atk, def, rng)
    return a, d, l
}
// SimulateBattleCopyWithTerrain は左右地形を考慮して1ラウンド模擬戦を行います（永続化なし）。
func SimulateBattleCopyWithTerrain(atk, def Unit, attT, defT gcore.Terrain, rng *rand.Rand) (Unit, Unit, []string) {
    a, d, l := uiscreens.SimulateBattleCopyWithTerrain(atk, def, attT, defT, rng)
    return a, d, l
}
// SimBattleButtonRect は一覧画面の「模擬戦」ボタン矩形を返します。
func SimBattleButtonRect(sw, sh int) (int, int, int, int) {
	return uiwidgets.SimBattleButtonRect(sw, sh)
}
// DrawSimBattleButton は一覧画面の「模擬戦」ボタンを描画します。
func DrawSimBattleButton(dst *ebiten.Image, hovered, enabled bool) {
	uiwidgets.DrawSimBattleButton(dst, hovered, enabled)
}
// ListItemRect は一覧画面の i 行目の矩形を返します。
func ListItemRect(sw, sh, i int) (int, int, int, int) { return uiscreens.ListItemRect(sw, sh, i) }

// BackButtonRect は画面右上の戻るボタン矩形を返します。
// BackButtonRect は戻るボタンの矩形を返します。
func BackButtonRect(sw, sh int) (int, int, int, int)    { return uiwidgets.BackButtonRect(sw, sh) }
// DrawBackButton は戻るボタンを描画します。
func DrawBackButton(dst *ebiten.Image, hovered bool)    { uiwidgets.DrawBackButton(dst, hovered) }
// LevelUpButtonRect はレベルアップボタンの矩形を返します。
func LevelUpButtonRect(sw, sh int) (int, int, int, int) { return uiwidgets.LevelUpButtonRect(sw, sh) }
// DrawLevelUpButton はレベルアップボタンを描画します。
func DrawLevelUpButton(dst *ebiten.Image, hovered, enabled bool) {
	uiwidgets.DrawLevelUpButton(dst, hovered, enabled)
}
// ToBattleButtonRect は「戦闘へ」ボタンの矩形を返します。
func ToBattleButtonRect(sw, sh int) (int, int, int, int) { return uiwidgets.ToBattleButtonRect(sw, sh) }
// DrawToBattleButton は「戦闘へ」ボタンを描画します。
func DrawToBattleButton(dst *ebiten.Image, hovered, enabled bool) {
    uiwidgets.DrawToBattleButton(dst, hovered, enabled)
}
// BattleStartButtonRect はバトル開始ボタンの矩形を返します。
func BattleStartButtonRect(sw, sh int) (int, int, int, int) {
    return uiscreens.BattleStartButtonRect(sw, sh)
}
// AutoRunButtonRect は開始ボタンの右隣（同サイズ/間隔S(20)）の矩形を返します。
// AutoRunButtonRect は開始ボタンの右隣に配置する（同サイズ/間隔S(20)）。
func AutoRunButtonRect(sw, sh int) (int, int, int, int) {
    bx, by, bw, bh := uiscreens.BattleStartButtonRect(sw, sh)
    gap := uicore.S(20)
    return bx + bw + gap, by, bw, bh
}

// DrawAutoRunButton は自動実行/停止のトグルボタンを描画します。
// DrawAutoRunButton は自動実行ボタンを描画します。
func DrawAutoRunButton(dst *ebiten.Image, hovered, running bool) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    bx, by, bw, bh := AutoRunButtonRect(sw, sh)
    uicore.DrawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
    base := color.RGBA{110, 90, 40, 255}
    if running {
        base = color.RGBA{150, 60, 60, 255}
    } else if hovered {
        base = color.RGBA{140, 110, 50, 255}
    }
    vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), base, false)
    label := "自動実行"
    if running { label = "停止" }
    uicore.TextDraw(dst, label, uicore.FaceMain, bx+uicore.S(70), by+uicore.S(38), uicore.ColText)
}

// TerrainButtonRect は地形ボタンの矩形を返します。DrawTerrainButtons は描画します。
// TerrainButtonRect は地形ボタン矩形を返します。
func TerrainButtonRect(sw, sh int, left bool, idx int) (int, int, int, int) {
    return uiwidgets.TerrainButtonRect(sw, sh, left, idx)
}
// DrawTerrainButtons は地形ボタンを描画します。
func DrawTerrainButtons(dst *ebiten.Image, attSel, defSel int) {
    uiwidgets.DrawTerrainButtons(dst, attSel, defSel)
}

// RollLevelUp はレベルアップの成長結果を生成します。ApplyGains はそれを適用します。DrawLevelUpPopup は結果を描画します。
// RollLevelUp はランダムに成長結果を生成します。
func RollLevelUp(u Unit, rnd func() float64) LevelUpGains        { return uipopup.RollLevelUp(u, rnd) }
// ApplyGains は成長結果をユニットに適用します。
func ApplyGains(u *Unit, g LevelUpGains, levelCap int)           { uipopup.ApplyGains(u, g, levelCap) }
// DrawLevelUpPopup はレベルアップ結果のポップアップを描画します。
func DrawLevelUpPopup(dst *ebiten.Image, u Unit, g LevelUpGains) { uipopup.DrawLevelUpPopup(dst, u, g) }

// SetWeaponTable は UI 層で共有する武器テーブルを設定します。
// SetWeaponTable は UI 層で共有する武器テーブルを設定します。
func SetWeaponTable(wt *model.WeaponTable) { uiscreens.SetWeaponTable(wt) }

// DrawChooseUnitPopup はユニット選択ポップアップを描画します。ChooseUnitItemRect は各項目の矩形を返します。
// DrawChooseUnitPopup はユニット選択ポップアップを描画します。
func DrawChooseUnitPopup(dst *ebiten.Image, title string, units []Unit, hover int) {
    uipopup.DrawChooseUnitPopup(dst, title, units, hover)
}
// ChooseUnitItemRect はユニット選択項目の矩形を返します。
func ChooseUnitItemRect(sw, sh, i, total int) (int, int, int, int) {
    return uipopup.ChooseUnitItemRect(sw, sh, i, total)
}
// EquipSlotRect はステータス画面の装備スロット行の矩形を返します。
func EquipSlotRect(sw, sh, i int) (int, int, int, int) { return uiscreens.EquipSlotRect(sw, sh, i) }

// DrawWeaponList は武器一覧を描画します。
func DrawWeaponList(dst *ebiten.Image, rows []WeaponRow, hover int) { uiscreens.DrawWeaponList(dst, rows, hover) }
// DrawItemList はアイテム一覧を描画します。
func DrawItemList(dst *ebiten.Image, rows []ItemRow, hover int)   { uiscreens.DrawItemList(dst, rows, hover) }

// LoadUserWeaponRows は usr_weapons と mst_weapons を結合した表示行を生成します。
func LoadUserWeaponRows(usrWeaponsPath, mstWeaponsPath string) ([]WeaponRow, error) {
    // usr
    ub, err := os.ReadFile(usrWeaponsPath)
    if err != nil { return nil, err }
    var owns []user.OwnWeapon
    if err := json.Unmarshal(ub, &owns); err != nil { return nil, err }
    // mst
    wt, err := model.LoadWeaponsJSON(mstWeaponsPath)
    if err != nil { return nil, err }
    rows := make([]WeaponRow, 0, len(owns))
    for _, ow := range owns {
        name := ow.MstWeaponsID
        typ, rank := "", ""
        mt, hit, crt, wtVal, rmin, rmax := 0, 0, 0, 0, 1, 1
        if w, ok := wt.FindByID(ow.MstWeaponsID); ok {
            name, typ, rank = w.Name, w.Type, w.Rank
            mt, hit, crt, wtVal = w.Might, w.Hit, w.Crit, w.Weight
            rmin, rmax = w.RangeMin, w.RangeMax
        }
        rows = append(rows, WeaponRow{ID: ow.ID, Name: name, Type: typ, Rank: rank, Might: mt, Hit: hit, Crit: crt, Weight: wtVal, RangeMin: rmin, RangeMax: rmax, Uses: ow.Uses, Max: ow.Max})
    }
    return rows, nil
}

// LoadUserItemRows は usr_items と mst_items を結合した表示行を生成します。
func LoadUserItemRows(usrItemsPath, mstItemsPath string) ([]ItemRow, error) {
    ub, err := os.ReadFile(usrItemsPath)
    if err != nil { return nil, err }
    var owns []user.OwnItem
    if err := json.Unmarshal(ub, &owns); err != nil { return nil, err }
    it, err := model.LoadItemsJSON(mstItemsPath)
    if err != nil { return nil, err }
    rows := make([]ItemRow, 0, len(owns))
    for _, oi := range owns {
        name := oi.MstItemsID
        typ, eff := "", ""
        pow := 0
        if d, ok := it.FindByID(oi.MstItemsID); ok {
            name, typ, eff, pow = d.Name, d.Type, d.Effect, d.Power
        }
        rows = append(rows, ItemRow{ID: oi.ID, Name: name, Type: typ, Effect: eff, Power: pow, Uses: oi.Uses, Max: oi.Max})
    }
    return rows, nil
}

// BuildWeaponRowsFromSnapshots は所持武器スナップショットと武器定義から行を構築します。
func BuildWeaponRowsFromSnapshots(owns []user.OwnWeapon, wt *model.WeaponTable) []WeaponRow {
    rows := make([]WeaponRow, 0, len(owns))
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
        rows = append(rows, WeaponRow{ID: ow.ID, Name: name, Type: typ, Rank: rank, Might: mt, Hit: hit, Crit: crt, Weight: wtVal, RangeMin: rmin, RangeMax: rmax, Uses: ow.Uses, Max: ow.Max})
    }
    return rows
}

// BuildItemRowsFromSnapshots は所持アイテムと定義から行を構築します。
func BuildItemRowsFromSnapshots(owns []user.OwnItem, it *model.ItemDefTable) []ItemRow {
    rows := make([]ItemRow, 0, len(owns))
    for _, oi := range owns {
        name := oi.MstItemsID
        typ, eff := "", ""
        pow := 0
        if it != nil {
            if d, ok := it.FindByID(oi.MstItemsID); ok {
                name, typ, eff, pow = d.Name, d.Type, d.Effect, d.Power
            }
        }
        rows = append(rows, ItemRow{ID: oi.ID, Name: name, Type: typ, Effect: eff, Power: pow, Uses: oi.Uses, Max: oi.Max})
    }
    return rows
}

// BuildWeaponRowsWithOwners は所有者バッジ情報付きの武器行を構築します。
func BuildWeaponRowsWithOwners(owns []user.OwnWeapon, wt *model.WeaponTable, ut *user.Table) []WeaponRow {
    rows := BuildWeaponRowsFromSnapshots(owns, wt)
    if ut == nil { return rows }
    // weaponID -> owners
    own := map[string][]OwnerBadge{}
    for _, c := range ut.Slice() {
        for _, er := range c.Equip {
            if er.UserWeaponsID != "" {
                var img *ebiten.Image
                if c.Portrait != "" { if im, err := assets.LoadImage(c.Portrait); err == nil { img = im } }
                own[er.UserWeaponsID] = append(own[er.UserWeaponsID], OwnerBadge{Name: c.Name, Portrait: img})
            }
        }
    }
    for i := range rows {
        rows[i].Owners = own[rows[i].ID]
    }
    return rows
}

// BuildItemRowsWithOwners は所有者バッジ情報付きのアイテム行を構築します。
func BuildItemRowsWithOwners(owns []user.OwnItem, it *model.ItemDefTable, ut *user.Table) []ItemRow {
    rows := BuildItemRowsFromSnapshots(owns, it)
    if ut == nil { return rows }
    own := map[string][]OwnerBadge{}
    for _, c := range ut.Slice() {
        for _, er := range c.Equip {
            if er.UserItemsID != "" {
                var img *ebiten.Image
                if c.Portrait != "" { if im, err := assets.LoadImage(c.Portrait); err == nil { img = im } }
                own[er.UserItemsID] = append(own[er.UserItemsID], OwnerBadge{Name: c.Name, Portrait: img})
            }
        }
    }
    for i := range rows { rows[i].Owners = own[rows[i].ID] }
    return rows
}
