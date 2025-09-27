package scenes

import (
    "fmt"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/ui/core"
    "ui_sample/internal/ui"
    "ui_sample/internal/user"
    gcore "ui_sample/pkg/game"
)

// popAware: Runner.AfterUpdate からの Pop 判定に用いる（cmd側で type assert）。
type popAware interface{ ShouldPop() bool }

// List は一覧画面の Scene です。
type List struct{
    E *Env
    hoverIndex int
    // 模擬戦の選択フロー
    simSelecting bool
    simSelectStep int
    chooseHover int
    tmpAtk ui.Unit
    sw, sh int
}

func NewList(e *Env) *List { return &List{E: e, hoverIndex: -1} }

func (s *List) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    mx, my := ebiten.CursorPosition()
    s.hoverIndex = -1
    for i := range s.E.Units {
        x, y, w, h := ui.ListItemRect(s.sw, s.sh, i)
        if pointIn(mx, my, x, y, w, h) { s.hoverIndex = i }
    }
    if s.hoverIndex >= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        s.E.SelIndex = s.hoverIndex
        return NewStatus(s.E), nil
    }
    // ショートカット: 武器/アイテム
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.OpenWeapons) {
        s.E.InvTab, s.E.HoverInv = 0, -1
        s.E.SelectingEquip, s.E.SelectingIsWeapon = false, true
        return NewInventory(s.E), nil
    }
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.OpenItems) {
        s.E.InvTab, s.E.HoverInv = 1, -1
        s.E.SelectingEquip, s.E.SelectingIsWeapon = false, false
        return NewInventory(s.E), nil
    }
    // 模擬戦ボタン
    bx, by, bw, bh := ui.SimBattleButtonRect(s.sw, s.sh)
    if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && len(s.E.Units) > 1 {
        s.simSelecting = true
        s.simSelectStep = 0
        s.chooseHover = -1
    }
    if s.simSelecting {
        if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.Cancel) {
            s.simSelecting = false
            s.simSelectStep = 0
            s.chooseHover = -1
            return nil, nil
        }
        s.chooseHover = -1
        for i := range s.E.Units {
            x, y, w, h := ui.ChooseUnitItemRect(s.sw, s.sh, i, len(s.E.Units))
            if pointIn(mx, my, x, y, w, h) { s.chooseHover = i }
        }
        if s.chooseHover >= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            if s.simSelectStep == 0 {
                s.tmpAtk = s.E.Units[s.chooseHover]
                s.simSelectStep = 1
                return nil, nil
            } else {
                def := s.E.Units[s.chooseHover]
                return NewSim(s.E, s.tmpAtk, def), nil
            }
        }
    }
    return nil, nil
}

func (s *List) Draw(dst *ebiten.Image) {
    ui.DrawCharacterList(dst, s.E.Units, s.hoverIndex)
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.SimBattleButtonRect(s.sw, s.sh)
    hovered := pointIn(mx, my, bx, by, bw, bh)
    ui.DrawSimBattleButton(dst, hovered, len(s.E.Units)>1)
    ebitenutil.DebugPrintAt(dst, "W: 武器一覧 / I: アイテム一覧", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
    if s.simSelecting {
        title := "模擬戦: 攻撃側を選択"
        if s.simSelectStep == 1 { title = "模擬戦: 防御側を選択" }
        ui.DrawChooseUnitPopup(dst, title, s.E.Units, s.chooseHover)
    }
}

func pointIn(px, py, x, y, w, h int) bool { return px >= x && py >= y && px < x+w && py < y+h }

// Status シーン
type Status struct{ E *Env; pop bool; sw, sh int }
func NewStatus(e *Env) *Status { return &Status{E:e} }
func (s *Status) ShouldPop() bool { return s.pop }

func (s *Status) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    mx, my := ebiten.CursorPosition()
    // 戻る
    bx, by, bw, bh := ui.BackButtonRect(s.sw, s.sh)
    if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.pop = true; return nil, nil }
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.Cancel) { s.pop = true; return nil, nil }

    // レベルアップ
    lbx, lby, lbw, lbh := ui.LevelUpButtonRect(s.sw, s.sh)
    unit := s.E.Selected()
    if unit.Level < game.LevelCap && !s.E.PopupActive && pointIn(mx,my,lbx,lby,lbw,lbh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        gains := ui.RollLevelUp(unit, s.E.RNG.Float64)
        ui.ApplyGains(&unit, gains, game.LevelCap)
        s.E.SetSelected(unit)
        s.E.PopupGains, s.E.PopupActive, s.E.PopupJustOpened = gains, true, true
        if s.E.UserTable != nil {
            if c, ok := s.E.UserTable.Find(unit.ID); ok {
                c.Level = unit.Level; c.HPMax = unit.HPMax
                c.Stats = user.Stats{Str: unit.Stats.Str, Mag: unit.Stats.Mag, Skl: unit.Stats.Skl, Spd: unit.Stats.Spd, Lck: unit.Stats.Lck, Def: unit.Stats.Def, Res: unit.Stats.Res, Mov: unit.Stats.Mov}
                s.E.UserTable.UpdateCharacter(c)
            }
        }
        if s.E.App != nil { _ = s.E.App.PersistUnit(unit) }
    }
    if s.E.PopupActive {
        if s.E.PopupJustOpened { s.E.PopupJustOpened = false } else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.E.PopupActive = false }
    }

    // 装備付け替え（ショートカット）
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.EquipToggle) {
        s.E.InvTab, s.E.HoverInv = 0, -1
        s.E.SelectingEquip, s.E.SelectingIsWeapon = true, true
        return NewInventory(s.E), nil
    }
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.OpenItems) {
        s.E.InvTab, s.E.HoverInv = 1, -1
        s.E.SelectingEquip, s.E.SelectingIsWeapon = true, false
        return NewInventory(s.E), nil
    }
    return nil, nil
}

func (s *Status) Draw(dst *ebiten.Image) {
    ui.DrawStatus(dst, s.E.Selected())
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(s.sw, s.sh)
    ui.DrawBackButton(dst, pointIn(mx,my,bx,by,bw,bh))
    lvx, lvy, lvw, lvh := ui.LevelUpButtonRect(s.sw, s.sh)
    unit := s.E.Selected()
    hovered := pointIn(mx,my,lvx,lvy,lvw,lvh)
    ui.DrawLevelUpButton(dst, hovered, unit.Level < game.LevelCap && !s.E.PopupActive)
    if s.E.PopupActive { ui.DrawLevelUpPopup(dst, unit, s.E.PopupGains) }
}

// Inventory シーン
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
        if s.E.App != nil && s.E.App.Inv != nil {
            s.E.InvTab, s.E.HoverInv = 0, -1
        }
    }
    tx2 := tx + tabW + uicore.S(10)
    if pointIn(mx,my,tx2,ty,tabW,tabH) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        if s.E.App != nil && s.E.App.Inv != nil { s.E.InvTab, s.E.HoverInv = 1, -1 }
    }
    // リスト操作（装備確定で戻る）
    if s.E.InvTab == 0 {
        // 武器
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
        // アイテム
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
    if s.E.InvTab == 0 {
        weapons := ui.BuildWeaponRowsWithOwners(s.E.App.Inv.Weapons(), s.E.App.WeaponsTable(), s.E.UserTable)
        ui.DrawWeaponList(dst, weapons, s.E.HoverInv)
    } else {
        it, _ := model.LoadItemsJSON("db/master/mst_items.json")
        items := ui.BuildItemRowsWithOwners(s.E.App.Inv.Items(), it, s.E.UserTable)
        ui.DrawItemList(dst, items, s.E.HoverInv)
    }
    // タブ
    // 既存UI描画（省略: 既存描画は一覧側のまま）
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(s.sw, s.sh)
    ui.DrawBackButton(dst, pointIn(mx,my,bx,by,bw,bh))
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
    for i := range s.E.Units { if s.E.Units[i].ID == id { s.E.Units[i] = u; if s.E.SelIndex==i { /* keep */ } } }
}

// Sim シーン（模擬戦）
type Sim struct{
    E *Env
    simAtk ui.Unit
    simDef ui.Unit
    logs   []string
    logPopup bool
    auto bool
    autoCD int
    turn int
    attTerrain gcore.Terrain
    defTerrain gcore.Terrain
    attSel int
    defSel int
    pop bool
    sw, sh int
}

func (s *Sim) ShouldPop() bool { return s.pop }

func (s *Sim) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    // ログポップアップ中は閉じるのみ
    if s.logPopup {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.logPopup=false; return nil,nil }
        if ctx!=nil && ctx.Input!=nil && ctx.Input.Press(gamesvc.Confirm) { s.logPopup=false; return nil,nil }
        return nil,nil
    }
    // 戻る
    bx, by, bw, bh := ui.BackButtonRect(s.sw, s.sh)
    mx, my := ebiten.CursorPosition()
    if pointIn(mx,my,bx,by,bw,bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.pop=true; return nil,nil }
    if ctx!=nil && ctx.Input!=nil && ctx.Input.Press(gamesvc.Cancel) { s.pop=true; return nil,nil }
    // 実行
    canStart := s.simAtk.HP > 0 && s.simDef.HP > 0
    bx2, by2, bw2, bh2 := ui.BattleStartButtonRect(s.sw, s.sh)
    leftFirst := (s.turn%2==1)
    if canStart && pointIn(mx,my,bx2,by2,bw2,bh2) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        s.runOne(leftFirst)
    }
    if canStart && ctx!=nil && ctx.Input!=nil && ctx.Input.Press(gamesvc.Confirm) {
        s.runOne(leftFirst)
    }
    // 自動実行
    if s.auto && canStart && !s.logPopup {
        if s.autoCD>0 { s.autoCD-- } else { s.runOne(leftFirst); s.autoCD=10 }
    }
    // 自動実行トグル
    ax, ay, aw, ah := ui.AutoRunButtonRect(s.sw, s.sh)
    if pointIn(mx,my,ax,ay,aw,ah) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.auto=!s.auto; if s.auto { s.logPopup=false } }
    // 地形ボタン
    for i:=0;i<3;i++{
        ax,ay,aw,ah := ui.TerrainButtonRect(s.sw, s.sh, true, i)
        if pointIn(mx,my,ax,ay,aw,ah) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.attSel=i; switch i{case 0:s.attTerrain=gcore.Terrain{}; case 1:s.attTerrain=gcore.Terrain{Avoid:20,Def:1}; case 2:s.attTerrain=gcore.Terrain{Avoid:15,Def:2}} }
        dx,dy,dw,dh := ui.TerrainButtonRect(s.sw, s.sh, false, i)
        if pointIn(mx,my,dx,dy,dw,dh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.defSel=i; switch i{case 0:s.defTerrain=gcore.Terrain{}; case 1:s.defTerrain=gcore.Terrain{Avoid:20,Def:1}; case 2:s.defTerrain=gcore.Terrain{Avoid:15,Def:2}} }
    }
    return nil,nil
}

func (s *Sim) runOne(leftFirst bool){
    if leftFirst {
        a,d,lines := ui.SimulateBattleCopyWithTerrain(s.simAtk, s.simDef, s.attTerrain, s.defTerrain, s.E.RNG)
        s.simAtk, s.simDef = a,d; s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simAtk.Name)}, lines...)
    } else {
        a,d,lines := ui.SimulateBattleCopyWithTerrain(s.simDef, s.simAtk, s.defTerrain, s.attTerrain, s.E.RNG)
        s.simDef, s.simAtk = a,d; s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simDef.Name)}, lines...)
    }
    s.logPopup=true; s.turn++
}

func (s *Sim) Draw(dst *ebiten.Image){
    canStart := s.simAtk.HP>0 && s.simDef.HP>0 && !s.logPopup
    ui.DrawBattleWithTerrain(dst, s.simAtk, s.simDef, s.attTerrain, s.defTerrain, canStart)
    ui.DrawTerrainButtons(dst, s.attSel, s.defSel)
    mx,my := ebiten.CursorPosition(); ax,ay,aw,ah := ui.AutoRunButtonRect(s.sw, s.sh);
    ui.DrawAutoRunButton(dst, pointIn(mx,my,ax,ay,aw,ah), s.auto)
    if s.logPopup { ui.DrawBattleLogOverlay(dst, s.logs) }
    if s.turn<=0 { s.turn=1 }
    leftFirst := (s.turn%2==1); label := "先攻: "; if leftFirst { label+=s.simAtk.Name } else { label+=s.simDef.Name }
    ebitenutil.DebugPrintAt(dst, label, uicore.ListMarginPx()+uicore.S(40), uicore.ListMarginPx()+uicore.S(56))
    bx,by,bw,bh := ui.BackButtonRect(s.sw, s.sh); ui.DrawBackButton(dst, pointIn(mx,my,bx,by,bw,bh))
}

func NewSim(e *Env, atk, def ui.Unit) *Sim {
    return &Sim{E:e, simAtk:atk, simDef:def, turn:1}
}
