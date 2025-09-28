package inventory

import (
    "fmt"
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    "ui_sample/internal/assets"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/model"
    "ui_sample/internal/user"
    scenes "ui_sample/internal/game/scenes"
)

// ItemRow はユーザ所持アイテム（耐久）+マスタ情報の結合行です（アイテムビュー用）。
type ItemRow struct {
    ID         string
    Name, Type string
    Effect     string
    Power      int
    Uses, Max  int
    Owners     []OwnerBadge
}

// ItemView は在庫（アイテム）サブビューです。
type ItemView struct{
    E *scenes.Env
    Host *Inventory
    sw, sh int
    hover int
}

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
    it, _ := model.LoadItemsJSON("db/master/mst_items.json")
    rows := BuildItemRowsWithOwners(v.E.App.Inventory().Items(), it, v.E.UserTable)
    DrawItemListView(dst, rows, v.hover)
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
    count := len(v.E.App.Inventory().Items())
    for i := 0; i < count; i++ {
        x, y, w, h := scenes.ListItemRect(v.sw, v.sh, i)
        if scenes.PointIn(mx, my, x, y, w, h) { v.hover = i }
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
        switch it.Kind {
        case ivChooseRow:
            owns := v.E.App.Inventory().Items()
            if it.Index >= 0 && it.Index < len(owns) { v.Host.equipItem(owns[it.Index].ID); v.Host.pop = true }
        }
    }
}

func (v *ItemView) scFlush(_ *game.Ctx) { /* 今はなし */ }

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

// toWidgetItemRows は描画ウィジェット用のVMへ変換します（最後の所有者のみ表示）。
// DrawItemListView はアイテム一覧（アイテムビュー）を描画します。
func DrawItemListView(dst *ebiten.Image, rows []ItemRow, hover int) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := uicore.ListMarginPx()
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    // タイトル
    uicore.TextDraw(dst, "アイテム一覧", uicore.FaceTitle, lm+uicore.S(20), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
    // ヘッダ
    startY := lm + uicore.ListTitleOffsetPx() + uicore.S(8)
    hx := lm + uicore.S(36)
    hy := startY
    uicore.TextDraw(dst, "名称", uicore.FaceSmall, hx+uicore.S(0), hy, uicore.ColText)
    uicore.TextDraw(dst, "種別", uicore.FaceSmall, hx+uicore.S(560), hy, uicore.ColText)
    uicore.TextDraw(dst, "効果", uicore.FaceSmall, hx+uicore.S(720), hy, uicore.ColText)
    uicore.TextDraw(dst, "数値", uicore.FaceSmall, hx+uicore.S(900), hy, uicore.ColText)
    uicore.TextDraw(dst, "耐久", uicore.FaceSmall, hx+uicore.S(1000), hy, uicore.ColText)

    for i, it := range rows {
        x, y, width, h := scenes.ListItemRect(sw, sh, i)
        bg := color.RGBA{30, 45, 78, 255}
        if i == hover { bg = color.RGBA{40, 60, 100, 255} }
        vector.DrawFilledRect(dst, float32(x), float32(y), float32(width), float32(h), bg, false)
        vector.DrawFilledRect(dst, float32(x-uicore.S(2)), float32(y-uicore.S(2)), float32(width+uicore.S(4)), float32(h+uicore.S(4)), uicore.ColBorder, false)

        tx := x + uicore.S(20)
        ty := y + uicore.S(36)
        uicore.TextDraw(dst, it.Name, uicore.FaceMain, tx, ty, uicore.ColText)
        uicore.TextDraw(dst, it.Type, uicore.FaceSmall, tx+uicore.S(540), ty, uicore.ColAccent)
        uicore.TextDraw(dst, it.Effect, uicore.FaceSmall, tx+uicore.S(700), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", it.Power), uicore.FaceSmall, tx+uicore.S(880), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d/%d", it.Uses, it.Max), uicore.FaceSmall, tx+uicore.S(980), ty, uicore.ColAccent)

        if n := len(it.Owners); n > 0 {
            ob := it.Owners[n-1]
            icon := uicore.S(24)
            ox := x + width - uicore.S(12) - icon
            oy := y + (h-icon)/2
            uicore.DrawFramedRect(dst, float32(ox), float32(oy), float32(icon), float32(icon))
            if ob.Portrait != nil {
                uicore.DrawPortrait(dst, ob.Portrait, float32(ox), float32(oy), float32(icon), float32(icon))
            } else {
                uicore.DrawPortraitPlaceholder(dst, float32(ox), float32(oy), float32(icon), float32(icon))
            }
        }
    }
}
