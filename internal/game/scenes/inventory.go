package scenes

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/ui/core"
    "ui_sample/internal/ui"
    "ui_sample/internal/user"
)

// Inventory は在庫画面の Scene です。
type Inventory struct{ E *Env; pop bool; sw, sh int }
func NewInventory(e *Env) *Inventory { return &Inventory{E:e} }
func (s *Inventory) ShouldPop() bool { return s.pop }

func (s *Inventory) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    mx, my := ebiten.CursorPosition()
    // タブ切り替え
    tabW, tabH := uicore.S(160), uicore.S(44)
    lm := uicore.ListMarginPx()
    tx := lm + uicore.S(20)
    ty := lm + uicore.S(12)
    if pointIn(mx,my,tx,ty,tabW,tabH) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        if s.E.App != nil && s.E.App.Inv != nil { s.E.InvTab, s.E.HoverInv = 0, -1 }
    }
    tx2 := tx + tabW + uicore.S(10)
    if pointIn(mx,my,tx2,ty,tabW,tabH) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        if s.E.App != nil && s.E.App.Inv != nil { s.E.InvTab, s.E.HoverInv = 1, -1 }
    }
    // リスト操作（装備確定で戻る）
    if s.E.InvTab == 0 {
        weapons := ui.BuildWeaponRowsWithOwners(s.E.App.Inv.Weapons(), s.E.App.WeaponsTable(), s.E.UserTable)
        for i := range weapons {
            x, y, w, h := ui.ListItemRect(s.sw, s.sh, i)
            if pointIn(mx,my,x,y,w,h) { s.E.HoverInv = i }
            if s.E.SelectingEquip && s.E.SelectingIsWeapon && s.E.HoverInv==i && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
                chosen := weapons[i]
                s.equipWeapon(chosen.ID)
                s.pop = true
                return nil, nil
            }
        }
    } else {
        it, _ := model.LoadItemsJSON("db/master/mst_items.json")
        items := ui.BuildItemRowsWithOwners(s.E.App.Inv.Items(), it, s.E.UserTable)
        for i := range items {
            x, y, w, h := ui.ListItemRect(s.sw, s.sh, i)
            if pointIn(mx,my,x,y,w,h) { s.E.HoverInv = i }
            if s.E.SelectingEquip && !s.E.SelectingIsWeapon && s.E.HoverInv==i && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
                chosen := items[i]
                s.equipItem(chosen.ID)
                s.pop = true
                return nil, nil
            }
        }
    }
    // 戻る
    bx, by, bw, bh := ui.BackButtonRect(s.sw, s.sh)
    if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.pop = true }
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.Cancel) { s.pop = true }
    return nil, nil
}

func (s *Inventory) Draw(dst *ebiten.Image) {
    // 本体（タブに応じて）
    if s.E.InvTab == 0 {
        weapons := ui.BuildWeaponRowsWithOwners(s.E.App.Inv.Weapons(), s.E.App.WeaponsTable(), s.E.UserTable)
        ui.DrawWeaponList(dst, weapons, s.E.HoverInv)
    } else {
        it, _ := model.LoadItemsJSON("db/master/mst_items.json")
        items := ui.BuildItemRowsWithOwners(s.E.App.Inv.Items(), it, s.E.UserTable)
        ui.DrawItemList(dst, items, s.E.HoverInv)
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
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(s.sw, s.sh)
    ui.DrawBackButton(dst, pointIn(mx, my, bx, by, bw, bh))
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

