package scenes

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/ui/core"
    "ui_sample/internal/ui"
    "ui_sample/internal/user"
)

// Status はステータス画面の Scene です。
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
    // 本体（ステータス）
    unit := s.E.Selected()
    ui.DrawStatus(dst, unit)
    // 戻るボタン
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(s.sw, s.sh)
    ui.DrawBackButton(dst, pointIn(mx, my, bx, by, bw, bh))
    // レベルアップボタン
    lvx, lvy, lvw, lvh := ui.LevelUpButtonRect(s.sw, s.sh)
    lvHovered := pointIn(mx, my, lvx, lvy, lvw, lvh)
    ui.DrawLevelUpButton(dst, lvHovered, unit.Level < game.LevelCap && !s.E.PopupActive)
    if s.E.PopupActive { ui.DrawLevelUpPopup(dst, unit, s.E.PopupGains) }
    ebitenutil.DebugPrintAt(dst, "装備: E/数字/DELETE で操作", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
}

