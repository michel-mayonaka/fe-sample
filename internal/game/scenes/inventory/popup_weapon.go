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
    gdata "ui_sample/internal/game/data"
)

// WeaponRow はユーザ所持武器（耐久含む）+マスタ情報の結合行です（武器ビュー用）。
type WeaponRow struct {
    ID                  string
    Name, Type, Rank    string
    Might, Hit, Crit    int
    Weight              int
    RangeMin, RangeMax  int
    Uses, Max           int
    Owners              []OwnerBadge
}

// WeaponView は在庫（武器）サブビューです。
type WeaponView struct{
    E *scenes.Env
    Host *Inventory
    sw, sh int
    hover int
}

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
    var wt *model.WeaponTable
    if p := gdata.Provider(); p != nil { wt = p.WeaponsTable() }
    rows := BuildWeaponRowsWithOwners(v.E.Inv.Inventory().Weapons(), wt, v.E.UserTable)
    DrawWeaponListView(dst, rows, v.hover)
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
    count := len(v.E.Inv.Inventory().Weapons())
    for i := 0; i < count; i++ {
        x, y, w, h := scenes.ListItemRect(v.sw, v.sh, i)
        if scenes.PointIn(mx, my, x, y, w, h) { v.hover = i }
    }
    if ctx != nil && ctx.Input != nil {
        if v.Host.E.SelectingEquip && v.hover >= 0 && ctx.Input.Press(gamesvc.Confirm) {
            intents = append(intents, wvIntent{Kind: wvChooseRow, Index: v.hover})
        }
    }
    return intents
}

func (v *WeaponView) scAdvance(intents []scenes.Intent) {
    for _, any := range intents {
        it, ok := any.(wvIntent); if !ok { continue }
        switch it.Kind {
        case wvChooseRow:
            owns := v.E.Inv.Inventory().Weapons()
            if it.Index >= 0 && it.Index < len(owns) {
                _ = v.E.Inv.EquipWeapon(v.E.Selected().ID, v.E.CurrentSlot, owns[it.Index].ID)
                v.Host.refreshUnitByID(v.E.Selected().ID)
                v.Host.pop = true
            }
        }
    }
}

func (v *WeaponView) scFlush(_ *game.Ctx) { /* 今はなし */ }

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

// BuildWeaponRowsWithOwners は所有者バッジ情報付きの武器行を構築します。
func BuildWeaponRowsWithOwners(owns []user.OwnWeapon, wt *model.WeaponTable, ut *user.Table) []WeaponRow {
    rows := BuildWeaponRowsFromSnapshots(owns, wt)
    if ut == nil { return rows }
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
    for i := range rows { rows[i].Owners = own[rows[i].ID] }
    return rows
}

// toWidgetWeaponRows は描画ウィジェット用のVMへ変換します（最後の所有者のみ表示）。
// DrawWeaponListView は武器一覧（武器ビュー）を描画します。
func DrawWeaponListView(dst *ebiten.Image, rows []WeaponRow, hover int) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := uicore.ListMarginPx()
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    // タイトル
    uicore.TextDraw(dst, "武器一覧", uicore.FaceTitle, lm+uicore.S(20), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
    // ヘッダ
    startY := lm + uicore.ListTitleOffsetPx() + uicore.S(8)
    hx := lm + uicore.S(36)
    hy := startY
    uicore.TextDraw(dst, "名称", uicore.FaceSmall, hx+uicore.S(0), hy, uicore.ColText)
    uicore.TextDraw(dst, "種別", uicore.FaceSmall, hx+uicore.S(560), hy, uicore.ColText)
    uicore.TextDraw(dst, "ﾗﾝｸ", uicore.FaceSmall, hx+uicore.S(680), hy, uicore.ColText)
    uicore.TextDraw(dst, "威力", uicore.FaceSmall, hx+uicore.S(760), hy, uicore.ColText)
    uicore.TextDraw(dst, "命中", uicore.FaceSmall, hx+uicore.S(840), hy, uicore.ColText)
    uicore.TextDraw(dst, "必殺", uicore.FaceSmall, hx+uicore.S(920), hy, uicore.ColText)
    uicore.TextDraw(dst, "重さ", uicore.FaceSmall, hx+uicore.S(1000), hy, uicore.ColText)
    uicore.TextDraw(dst, "射程", uicore.FaceSmall, hx+uicore.S(1080), hy, uicore.ColText)
    uicore.TextDraw(dst, "耐久", uicore.FaceSmall, hx+uicore.S(1160), hy, uicore.ColText)

    for i, w := range rows {
        x, y, width, h := scenes.ListItemRect(sw, sh, i)
        bg := color.RGBA{30, 45, 78, 255}
        if i == hover { bg = color.RGBA{40, 60, 100, 255} }
        vector.DrawFilledRect(dst, float32(x), float32(y), float32(width), float32(h), bg, false)
        vector.DrawFilledRect(dst, float32(x-uicore.S(2)), float32(y-uicore.S(2)), float32(width+uicore.S(4)), float32(h+uicore.S(4)), uicore.ColBorder, false)

        tx := x + uicore.S(20)
        ty := y + uicore.S(36)
        uicore.TextDraw(dst, w.Name, uicore.FaceMain, tx, ty, uicore.ColText)
        uicore.TextDraw(dst, w.Type, uicore.FaceSmall, tx+uicore.S(540), ty, uicore.ColAccent)
        uicore.TextDraw(dst, w.Rank, uicore.FaceSmall, tx+uicore.S(660), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Might), uicore.FaceSmall, tx+uicore.S(750), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Hit), uicore.FaceSmall, tx+uicore.S(830), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Crit), uicore.FaceSmall, tx+uicore.S(910), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Weight), uicore.FaceSmall, tx+uicore.S(990), ty, uicore.ColAccent)
        rng := fmt.Sprintf("%d", w.RangeMin)
        if w.RangeMax != w.RangeMin { rng = fmt.Sprintf("%d-%d", w.RangeMin, w.RangeMax) }
        uicore.TextDraw(dst, rng, uicore.FaceSmall, tx+uicore.S(1070), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d/%d", w.Uses, w.Max), uicore.FaceSmall, tx+uicore.S(1150), ty, uicore.ColAccent)

        if n := len(w.Owners); n > 0 {
            ob := w.Owners[n-1]
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
