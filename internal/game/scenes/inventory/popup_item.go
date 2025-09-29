package inventory

import (
    "github.com/hajimehoshi/ebiten/v2"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    "ui_sample/internal/game/scenes"
    gdata "ui_sample/internal/game/data"
    uidraw "ui_sample/internal/game/ui/draw"
    uilayout "ui_sample/internal/game/ui/layout"
    uiadapter "ui_sample/internal/game/ui/adapter"
    "ui_sample/internal/model"
    // usr はビルダー移設済みのため未使用
    "ui_sample/pkg/game/geom"
)

// ItemRow はユーザ所持アイテム（耐久）+マスタ情報の結合行です（アイテムビュー用）。
// ItemRow は `ui/view` へ移設しました。

// ItemView は在庫（アイテム）サブビューです。
type ItemView struct{
    E *scenes.Env
    Host *Inventory
    sw, sh int
    hover int
}
// NewItemView はアイテム一覧サブビューを生成します。
func NewItemView(e *scenes.Env, host *Inventory) *ItemView { return &ItemView{E:e, Host:host, hover:-1} }

// Update はサブビュー内の状態更新を行います。
func (v *ItemView) Update(ctx *game.Ctx) (game.Scene, error) {
    v.sw, v.sh = ctx.ScreenW, ctx.ScreenH
    intents := v.scHandleInput(ctx)
    v.scAdvance(intents)
    v.scFlush(ctx)
    return nil, nil
}

// Draw はアイテム一覧の描画を行います。
func (v *ItemView) Draw(dst *ebiten.Image) {
    var it *model.ItemDefTable
    if p := gdata.Provider(); p != nil { it = p.ItemsTable() }
    rows := uiadapter.BuildItemRows(v.E.Inv.Inventory().Items(), it, v.E.UserTable, uiadapter.AssetsPortraitLoader{})
    uidraw.DrawItemListView(dst, rows, v.hover)
}

// --- 内部: scHandleInput → scAdvance → scFlush --------------------------------------

type ivIntentKind int
const (
    ivNone ivIntentKind = iota
    ivChooseRow
)
type ivIntent struct{ Kind ivIntentKind; Index int }
func (ivIntent) IsSceneIntent() {}

func (v *ItemView) scHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 1)
    mx, my := ebiten.CursorPosition()
    // 行ホバー更新
    v.hover = -1
    count := len(v.E.Inv.Inventory().Items())
    for i := 0; i < count; i++ {
        x, y, w, h := uilayout.ListItemRect(v.sw, v.sh, i)
        if geom.RectContains(mx, my, x, y, w, h) { v.hover = i }
    }
    if ctx != nil && ctx.Input != nil {
        if v.Host.E.SelectingEquip && v.hover >= 0 && ctx.Input.Press(gamesvc.Confirm) {
            intents = append(intents, ivIntent{Kind: ivChooseRow, Index: v.hover})
        }
    }
    return intents
}

func (v *ItemView) scAdvance(intents []scenes.Intent) {
    for _, any := range intents {
        it, ok := any.(ivIntent); if !ok { continue }
        if it.Kind == ivChooseRow {
            owns := v.E.Inv.Inventory().Items()
            if it.Index >= 0 && it.Index < len(owns) {
                _ = v.E.Inv.EquipItem(v.E.Selected().ID, v.E.CurrentSlot, owns[it.Index].ID)
                v.Host.refreshUnitByID(v.E.Selected().ID)
                v.Host.pop = true
            }
        }
    }
}

func (v *ItemView) scFlush(_ *game.Ctx) { /* 今はなし */ }

// BuildItemRowsFromSnapshots は所持アイテムと定義から行を構築します。
// 行データ生成は ui/adapter へ移設しました。
