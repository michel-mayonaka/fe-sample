package status

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    scpopup "ui_sample/internal/game/scenes/common/popup"
    scenes "ui_sample/internal/game/scenes"
    inventory "ui_sample/internal/game/scenes/inventory"
    "ui_sample/internal/user"
)

// Status はステータス画面の Scene です。
type Status struct{ E *scenes.Env; pop bool; sw, sh int }
func NewStatus(e *scenes.Env) *Status { return &Status{E:e} }
func (s *Status) ShouldPop() bool { return s.pop }

func (s *Status) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    mx, my := ebiten.CursorPosition()
    // 直前にポップアップが開いていたか（クリックの二重処理防止）
    wasPopup := s.E != nil && s.E.PopupActive
    // 戻る
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    if scenes.PointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.pop = true; return nil, nil }
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.Cancel) { s.pop = true; return nil, nil }

    // レベルアップ
    lbx, lby, lbw, lbh := uiwidgets.LevelUpButtonRect(s.sw, s.sh)
    unit := s.E.Selected()
    if unit.Level < game.LevelCap && !s.E.PopupActive && scenes.PointIn(mx,my,lbx,lby,lbw,lbh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        gains := scpopup.RollLevelUp(unit, s.E.RNG.Float64)
        scpopup.ApplyGains(&unit, gains, game.LevelCap)
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

    // ポップアップを閉じるクリックと同フレームでは以降のクリック処理を無効化
    if wasPopup && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        return nil, nil
    }

    // 装備スロットクリックで在庫へ
    for i := 0; i < 5; i++ {
        x, y, w, h := scenes.EquipSlotRect(s.sw, s.sh, i)
        if scenes.PointIn(mx, my, x, y, w, h) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            s.E.CurrentSlot = i
            s.E.SelectingEquip = true
            s.E.HoverInv = -1
            // 既装備の種別からタブを初期選択（なければ武器）
            s.E.SelectingIsWeapon = true
            s.E.InvTab = 0
            if s.E.UserTable != nil {
                unit := s.E.Selected()
                if c, ok := s.E.UserTable.Find(unit.ID); ok {
                    if i < len(c.Equip) {
                        er := c.Equip[i]
                        if er.UserItemsID != "" { s.E.SelectingIsWeapon = false; s.E.InvTab = 1 }
                        if er.UserWeaponsID != "" { s.E.SelectingIsWeapon = true;  s.E.InvTab = 0 }
                    }
                }
            }
            return inventory.NewInventory(s.E), nil
        }
    }

    // 装備付け替え（ショートカット）
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.EquipToggle) {
        s.E.InvTab, s.E.HoverInv = 0, -1
        s.E.SelectingEquip, s.E.SelectingIsWeapon = true, true
        return inventory.NewInventory(s.E), nil
    }
    if ctx != nil && ctx.Input != nil && ctx.Input.Press(gamesvc.OpenItems) {
        s.E.InvTab, s.E.HoverInv = 1, -1
        s.E.SelectingEquip, s.E.SelectingIsWeapon = true, false
        return inventory.NewInventory(s.E), nil
    }
    return nil, nil
}

func (s *Status) Draw(dst *ebiten.Image) {
    // 本体（ステータス）
    unit := s.E.Selected()
    scenes.DrawStatus(dst, unit)
    // 戻るボタン
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    uiwidgets.DrawBackButton(dst, scenes.PointIn(mx, my, bx, by, bw, bh))
    // レベルアップボタン
    lvx, lvy, lvw, lvh := uiwidgets.LevelUpButtonRect(s.sw, s.sh)
    lvHovered := scenes.PointIn(mx, my, lvx, lvy, lvw, lvh)
    uiwidgets.DrawLevelUpButton(dst, lvHovered, unit.Level < game.LevelCap && !s.E.PopupActive)
    if s.E.PopupActive { scpopup.DrawLevelUpPopup(dst, unit, s.E.PopupGains) }
    ebitenutil.DebugPrintAt(dst, "装備: E/数字/DELETE で操作", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
}
