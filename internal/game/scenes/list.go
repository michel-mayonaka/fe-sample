package scenes

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    scpopup "ui_sample/internal/game/scenes/common/popup"
)

// List は一覧画面の Scene です。
type List struct {
    E *Env
    hoverIndex int
    // 模擬戦の選択フロー
    simSelecting bool
    simSelectStep int
    chooseHover int
    tmpAtk uicore.Unit
    sw, sh int
}

func NewList(e *Env) *List { return &List{E: e, hoverIndex: -1} }

func (s *List) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    mx, my := ebiten.CursorPosition()
    s.hoverIndex = -1
    for i := range s.E.Units {
        x, y, w, h := ListItemRect(s.sw, s.sh, i)
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
    bx, by, bw, bh := uiwidgets.SimBattleButtonRect(s.sw, s.sh)
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
            x, y, w, h := scpopup.ChooseUnitItemRect(s.sw, s.sh, i, len(s.E.Units))
            if pointIn(mx, my, x, y, w, h) { s.chooseHover = i }
        }
        if s.chooseHover >= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            if s.simSelectStep == 0 {
                s.tmpAtk = s.E.Units[s.chooseHover]
                s.simSelectStep = 1
                return nil, nil
            }
            def := s.E.Units[s.chooseHover]
            return NewSim(s.E, s.tmpAtk, def), nil
        }
    }
    return nil, nil
}

func (s *List) Draw(dst *ebiten.Image) {
    DrawCharacterList(dst, s.E.Units, s.hoverIndex)
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := uiwidgets.SimBattleButtonRect(s.sw, s.sh)
    hovered := pointIn(mx, my, bx, by, bw, bh)
    uiwidgets.DrawSimBattleButton(dst, hovered, len(s.E.Units)>1)
    ebitenutil.DebugPrintAt(dst, "W: 武器一覧 / I: アイテム一覧", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
    if s.simSelecting {
        title := "模擬戦: 攻撃側を選択"
        if s.simSelectStep == 1 { title = "模擬戦: 防御側を選択" }
        scpopup.DrawChooseUnitPopup(dst, title, s.E.Units, s.chooseHover)
    }
}
