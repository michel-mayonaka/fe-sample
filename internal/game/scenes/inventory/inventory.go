package inventory

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    scenes "ui_sample/internal/game/scenes"
    "ui_sample/internal/user"
)

// Inventory は在庫画面の Scene です（character_list 準拠の構成）。
type Inventory struct{
    E *scenes.Env
    pop bool
    sw, sh int
    // ホバー状態（描画に使用）
    tabHover int // -1=なし, 0=武器, 1=アイテム
    backHovered bool
}

func NewInventory(e *scenes.Env) *Inventory { return &Inventory{E:e, tabHover: -1} }
func (s *Inventory) ShouldPop() bool { return s.pop }

// Intent 種別
type IntentKind int

const (
    intentNone IntentKind = iota
    intentBack
    intentTabWeapons
    intentTabItems
    intentChooseRow // Index に行番号
)

// Intent は入力の意味表現です。
type Intent struct{
    Kind IntentKind
    Index int
}
func (Intent) IsSceneIntent() {}

// 内部コントラクト（コンパイル保証）
type ivContract interface {
    ivHandleInput(ctx *game.Ctx) []scenes.Intent
    ivAdvance(intents []scenes.Intent)
    ivFlush(ctx *game.Ctx)
}
var _ ivContract = (*Inventory)(nil)

func (s *Inventory) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    intents := s.ivHandleInput(ctx)
    s.ivAdvance(intents)
    s.ivFlush(ctx)
    return nil, nil
}

func (s *Inventory) Draw(dst *ebiten.Image) {
    // 本体（タブに応じて）
    if s.E.InvTab == 0 {
        weapons := scenes.BuildWeaponRowsWithOwners(s.E.App.Inventory().Weapons(), s.E.App.WeaponsTable(), s.E.UserTable)
        scenes.DrawWeaponList(dst, weapons, s.E.HoverInv)
    } else {
        it, _ := model.LoadItemsJSON("db/master/mst_items.json")
        items := scenes.BuildItemRowsWithOwners(s.E.App.Inventory().Items(), it, s.E.UserTable)
        scenes.DrawItemList(dst, items, s.E.HoverInv)
    }
    // タブ描画
    tabW, tabH := uicore.S(160), uicore.S(44)
    lm := uicore.ListMarginPx()
    tx := lm + uicore.S(20)
    ty := lm + uicore.S(12)
    // 武器タブ
    uicore.DrawFramedRect(dst, float32(tx), float32(ty), float32(tabW), float32(tabH))
    baseW := color.RGBA{40, 60, 110, 255}
    if s.E.InvTab == 0 { baseW = color.RGBA{70, 100, 160, 255} }
    vector.DrawFilledRect(dst, float32(tx), float32(ty), float32(tabW), float32(tabH), baseW, false)
    uicore.TextDraw(dst, "武器", uicore.FaceMain, tx+uicore.S(56), ty+uicore.S(30), uicore.ColText)
    // アイテムタブ
    tx2 := tx + tabW + uicore.S(10)
    uicore.DrawFramedRect(dst, float32(tx2), float32(ty), float32(tabW), float32(tabH))
    baseI := color.RGBA{40, 60, 110, 255}
    if s.E.InvTab == 1 { baseI = color.RGBA{70, 100, 160, 255} }
    vector.DrawFilledRect(dst, float32(tx2), float32(ty), float32(tabW), float32(tabH), baseI, false)
    uicore.TextDraw(dst, "アイテム", uicore.FaceMain, tx2+uicore.S(34), ty+uicore.S(30), uicore.ColText)
    // 戻る
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    mx, my := ebiten.CursorPosition()
    s.backHovered = scenes.PointIn(mx, my, bx, by, bw, bh)
    uiwidgets.DrawBackButton(dst, s.backHovered)
    if s.E.SelectingEquip { ebitenutil.DebugPrintAt(dst, "クリックでスロットに装備", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10)) }
}

func (s *Inventory) equipWeapon(userWeaponID string) {
    if s.E.UserTable == nil { return }
    unit := s.E.Selected()
    if c, ok := s.E.UserTable.Find(unit.ID); ok {
        var prev user.EquipRef
        if s.E.CurrentSlot < len(c.Equip) { prev = c.Equip[s.E.CurrentSlot] }
        // 既装備のオーナーから外す
        ownerID := ""; ownerSlot := -1
        for _, oc := range s.E.UserTable.Slice() {
            for idx, er := range oc.Equip { if er.UserWeaponsID == userWeaponID { ownerID = oc.ID; ownerSlot = idx; break } }
            if ownerID != "" { break }
        }
        if ownerID != "" { if oc, ok2 := s.E.UserTable.Find(ownerID); ok2 {
            for len(oc.Equip) <= ownerSlot { oc.Equip = append(oc.Equip, user.EquipRef{}) }
            oc.Equip[ownerSlot] = prev
            s.E.UserTable.UpdateCharacter(oc)
        }}
        for len(c.Equip) <= s.E.CurrentSlot { c.Equip = append(c.Equip, user.EquipRef{}) }
        c.Equip[s.E.CurrentSlot] = user.EquipRef{UserWeaponsID: userWeaponID}
        s.E.UserTable.UpdateCharacter(c); _ = s.E.UserTable.Save(s.E.UserPath)
        s.refreshUnitByID(c.ID)
    }
}

func (s *Inventory) equipItem(userItemID string) {
    if s.E.UserTable == nil { return }
    unit := s.E.Selected()
    if c, ok := s.E.UserTable.Find(unit.ID); ok {
        var prev user.EquipRef
        if s.E.CurrentSlot < len(c.Equip) { prev = c.Equip[s.E.CurrentSlot] }
        ownerID := ""; ownerSlot := -1
        for _, oc := range s.E.UserTable.Slice() {
            for idx, er := range oc.Equip { if er.UserItemsID == userItemID { ownerID = oc.ID; ownerSlot = idx; break } }
            if ownerID != "" { break }
        }
        if ownerID != "" { if oc, ok2 := s.E.UserTable.Find(ownerID); ok2 {
            for len(oc.Equip) <= ownerSlot { oc.Equip = append(oc.Equip, user.EquipRef{}) }
            oc.Equip[ownerSlot] = prev
            s.E.UserTable.UpdateCharacter(oc)
        }}
        for len(c.Equip) <= s.E.CurrentSlot { c.Equip = append(c.Equip, user.EquipRef{}) }
        c.Equip[s.E.CurrentSlot] = user.EquipRef{UserItemsID: userItemID}
        s.E.UserTable.UpdateCharacter(c); _ = s.E.UserTable.Save(s.E.UserPath)
        s.refreshUnitByID(c.ID)
    }
}

func (s *Inventory) refreshUnitByID(id string) {
    if s.E == nil || s.E.UserTable == nil { return }
    c, ok := s.E.UserTable.Find(id); if !ok { return }
    u := uicore.UnitFromUser(c)
    for i := range s.E.Units { if s.E.Units[i].ID == id { s.E.Units[i] = u; if s.E.SelIndex==i { /* keep selected */ } } }
}

// --- 内部: handle → advance → flush -------------------------------------------------

func (s *Inventory) ivHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 3)
    mx, my := ebiten.CursorPosition()
    // ホバー更新
    s.tabHover = -1
    tabW, tabH := uicore.S(160), uicore.S(44)
    lm := uicore.ListMarginPx()
    tx := lm + uicore.S(20)
    ty := lm + uicore.S(12)
    if scenes.PointIn(mx, my, tx, ty, tabW, tabH) { s.tabHover = 0 }
    tx2 := tx + tabW + uicore.S(10)
    if scenes.PointIn(mx, my, tx2, ty, tabW, tabH) { s.tabHover = 1 }

    // リストホバー（件数はユーザ在庫から取得）
    s.E.HoverInv = -1
    count := 0
    if s.E.InvTab == 0 { count = len(s.E.App.Inventory().Weapons()) } else { count = len(s.E.App.Inventory().Items()) }
    for i := 0; i < count; i++ {
        x, y, w, h := scenes.ListItemRect(s.sw, s.sh, i)
        if scenes.PointIn(mx, my, x, y, w, h) { s.E.HoverInv = i }
    }

    // ボタンホバー
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = scenes.PointIn(mx, my, bx, by, bw, bh)

    if ctx != nil && ctx.Input != nil {
        if ctx.Input.Press(gamesvc.Cancel) { intents = append(intents, Intent{Kind: intentBack}) }
        if ctx.Input.Press(gamesvc.Confirm) {
            switch s.tabHover {
            case 0: intents = append(intents, Intent{Kind: intentTabWeapons})
            case 1: intents = append(intents, Intent{Kind: intentTabItems})
            }
            if s.backHovered { intents = append(intents, Intent{Kind: intentBack}) }
            if s.E.SelectingEquip && s.E.HoverInv >= 0 {
                intents = append(intents, Intent{Kind: intentChooseRow, Index: s.E.HoverInv})
            }
        }
    }
    return intents
}

func (s *Inventory) ivAdvance(intents []scenes.Intent) {
    for _, any := range intents {
        it, ok := any.(Intent); if !ok { continue }
        switch it.Kind {
        case intentBack:
            s.pop = true
        case intentTabWeapons:
            if s.E.App != nil && s.E.App.Inventory() != nil { s.E.InvTab, s.E.HoverInv = 0, -1 }
        case intentTabItems:
            if s.E.App != nil && s.E.App.Inventory() != nil { s.E.InvTab, s.E.HoverInv = 1, -1 }
        case intentChooseRow:
            if !s.E.SelectingEquip || it.Index < 0 { continue }
            if s.E.InvTab == 0 {
                owns := s.E.App.Inventory().Weapons()
                if it.Index < len(owns) { s.equipWeapon(owns[it.Index].ID); s.pop = true }
            } else {
                owns := s.E.App.Inventory().Items()
                if it.Index < len(owns) { s.equipItem(owns[it.Index].ID); s.pop = true }
            }
        }
    }
}

func (s *Inventory) ivFlush(_ *game.Ctx) { /* 今はなし */ }
