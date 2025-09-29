package inventory

import (
    "github.com/hajimehoshi/ebiten/v2"
    "ui_sample/internal/game"
    "ui_sample/internal/game/scenes"
    gdata "ui_sample/internal/game/data"
    uidraw "ui_sample/internal/game/ui/draw"
    uilayout "ui_sample/internal/game/ui/layout"
    uiadapter "ui_sample/internal/game/ui/adapter"
    uinput "ui_sample/internal/game/ui/input"
    "ui_sample/pkg/game/geom"
)

// WeaponRow はユーザ所持武器（耐久含む）+マスタ情報の結合行です（武器ビュー用）。
// WeaponRow は `ui/view` へ移設しました。

// WeaponView は在庫（武器）サブビューです。
type WeaponView struct{
    E *scenes.Env
    Host *Inventory
    sw, sh int
    hover int
}
// NewWeaponView は武器一覧サブビューを生成します。
func NewWeaponView(e *scenes.Env, host *Inventory) *WeaponView { return &WeaponView{E:e, Host:host, hover:-1} }

// Update はサブビュー内の状態更新を行います。
func (v *WeaponView) Update(ctx *game.Ctx) (game.Scene, error) {
    v.sw, v.sh = ctx.ScreenW, ctx.ScreenH
    intents := v.scHandleInput(ctx)
    v.scAdvance(intents)
    v.scFlush(ctx)
    return nil, nil
}

// Draw は武器一覧の描画を行います。
func (v *WeaponView) Draw(dst *ebiten.Image) {
    // 参照は Provider 経由へ統一（未設定時は描画スキップ）
    if p := gdata.Provider(); p != nil {
        rows := uiadapter.BuildWeaponRows(p.UserWeapons(), p.WeaponsTable(), uiadapter.AssetsPortraitLoader{})
        uidraw.DrawWeaponListView(dst, rows, v.hover)
    }
}

// --- 内部: scHandleInput → scAdvance → scFlush --------------------------------------

type wvIntentKind int
const (
    wvNone wvIntentKind = iota
    wvChooseRow
)
type wvIntent struct{ Kind wvIntentKind; Index int }
func (wvIntent) IsSceneIntent() {}

func (v *WeaponView) scHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 1)
    mx, my := ebiten.CursorPosition()
    // 行ホバー更新
    v.hover = -1
    count := 0
    if p := gdata.Provider(); p != nil { count = len(p.UserWeapons()) }
    for i := 0; i < count; i++ {
        x, y, w, h := uilayout.ListItemRect(v.sw, v.sh, i)
        if geom.RectContains(mx, my, x, y, w, h) { v.hover = i }
    }
    if ctx != nil && ctx.Input != nil {
        if v.Host.E.SelectingEquip && v.hover >= 0 && ctx.Input.Press(uinput.Confirm) {
            intents = append(intents, wvIntent{Kind: wvChooseRow, Index: v.hover})
        }
    }
    return intents
}

func (v *WeaponView) scAdvance(intents []scenes.Intent) {
    for _, any := range intents {
        it, ok := any.(wvIntent); if !ok { continue }
        if it.Kind == wvChooseRow {
            if p := gdata.Provider(); p != nil {
                owns := p.UserWeapons()
                if it.Index >= 0 && it.Index < len(owns) {
                    _ = v.E.Inv.EquipWeapon(v.E.Selected().ID, v.E.CurrentSlot, owns[it.Index].ID)
                }
            }
            v.Host.refreshUnitByID(v.E.Selected().ID)
            v.Host.pop = true
        }
    }
}

func (v *WeaponView) scFlush(_ *game.Ctx) { /* 今はなし */ }

// BuildWeaponRowsFromSnapshots は所持武器スナップショットと武器定義から行を構築します。
// 行データ生成は ui/adapter へ移設しました。
