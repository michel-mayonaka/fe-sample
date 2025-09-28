package status

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    scpopup "ui_sample/internal/game/scenes/common/popup"
    scenes "ui_sample/internal/game/scenes"
    inventory "ui_sample/internal/game/scenes/inventory"
    "ui_sample/internal/user"
)

// Status はステータス画面の Scene です（character_list 準拠の構成）。
type Status struct{
    E *scenes.Env
    pop bool
    sw, sh int
    next game.Scene
    backHovered bool
    lvHovered bool
}

func NewStatus(e *scenes.Env) *Status { return &Status{E:e} }
func (s *Status) ShouldPop() bool { return s.pop }

// Intent 種別
type IntentKind int

const (
    intentNone IntentKind = iota
    intentBack
    intentClosePopup
    intentLevelUp
    intentOpenInvSlot // Index にスロット番号
    intentOpenInvWeapons
    intentOpenInvItems
)

type Intent struct{ Kind IntentKind; Index int }
func (Intent) IsSceneIntent() {}

type stContract interface{
    stHandleInput(ctx *game.Ctx) []scenes.Intent
    stAdvance([]scenes.Intent)
    stFlush(ctx *game.Ctx)
}
var _ stContract = (*Status)(nil)

func (s *Status) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    intents := s.stHandleInput(ctx)
    s.stAdvance(intents)
    s.stFlush(ctx)
    nxt := s.next
    s.next = nil
    return nxt, nil
}

func (s *Status) Draw(dst *ebiten.Image) {
    // 本体（ステータス）
    unit := s.E.Selected()
    scenes.DrawStatus(dst, unit)
    // 戻るボタン
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = scenes.PointIn(mx, my, bx, by, bw, bh)
    uiwidgets.DrawBackButton(dst, s.backHovered)
    // レベルアップボタン
    lvx, lvy, lvw, lvh := uiwidgets.LevelUpButtonRect(s.sw, s.sh)
    s.lvHovered = scenes.PointIn(mx, my, lvx, lvy, lvw, lvh)
    uiwidgets.DrawLevelUpButton(dst, s.lvHovered, unit.Level < game.LevelCap && !s.E.PopupActive)
    if s.E.PopupActive { scpopup.DrawLevelUpPopup(dst, unit, s.E.PopupGains) }
    ebitenutil.DebugPrintAt(dst, "装備: E/数字/DELETE で操作", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
}

// --- 内部: handle → advance → flush -------------------------------------------------

func (s *Status) stHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 4)
    mx, my := ebiten.CursorPosition()
    wasPopup := s.E != nil && s.E.PopupActive
    // 戻る/レベルアップ ホバー
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = scenes.PointIn(mx, my, bx, by, bw, bh)
    lvx, lvy, lvw, lvh := uiwidgets.LevelUpButtonRect(s.sw, s.sh)
    s.lvHovered = scenes.PointIn(mx, my, lvx, lvy, lvw, lvh)

    if ctx != nil && ctx.Input != nil {
        if ctx.Input.Press(gamesvc.Cancel) { intents = append(intents, Intent{Kind: intentBack}) }

        // ポップアップ中の操作
        if s.E.PopupActive {
            if s.E.PopupJustOpened { s.E.PopupJustOpened = false } else if ctx.Input.Press(gamesvc.Confirm) { intents = append(intents, Intent{Kind: intentClosePopup}) }
            // ポップアップを閉じる確定と同フレームでは以降を処理しない
            if wasPopup { return intents }
        }

        // レベルアップ
        unit := s.E.Selected()
        if unit.Level < game.LevelCap && !s.E.PopupActive && s.lvHovered && ctx.Input.Press(gamesvc.Confirm) {
            intents = append(intents, Intent{Kind: intentLevelUp})
        }

        // スロット選択→在庫
        for i := 0; i < 5; i++ {
            x, y, w, h := scenes.EquipSlotRect(s.sw, s.sh, i)
            if scenes.PointIn(mx, my, x, y, w, h) && ctx.Input.Press(gamesvc.Confirm) {
                intents = append(intents, Intent{Kind: intentOpenInvSlot, Index: i})
                break
            }
        }
        // ショートカット（E/I）
        if ctx.Input.Press(gamesvc.EquipToggle) { intents = append(intents, Intent{Kind: intentOpenInvWeapons}) }
        if ctx.Input.Press(gamesvc.OpenItems)   { intents = append(intents, Intent{Kind: intentOpenInvItems}) }
    }
    return intents
}

func (s *Status) stAdvance(intents []scenes.Intent) {
    for _, any := range intents {
        it, ok := any.(Intent); if !ok { continue }
        switch it.Kind {
        case intentBack:
            s.pop = true
        case intentClosePopup:
            s.E.PopupActive = false
        case intentLevelUp:
            unit := s.E.Selected()
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
        case intentOpenInvSlot:
            i := it.Index
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
            s.next = inventory.NewInventory(s.E)
        case intentOpenInvWeapons:
            s.E.InvTab, s.E.HoverInv = 0, -1
            s.E.SelectingEquip, s.E.SelectingIsWeapon = true, true
            s.next = inventory.NewInventory(s.E)
        case intentOpenInvItems:
            s.E.InvTab, s.E.HoverInv = 1, -1
            s.E.SelectingEquip, s.E.SelectingIsWeapon = true, false
            s.next = inventory.NewInventory(s.E)
        }
    }
}

func (s *Status) stFlush(_ *game.Ctx) { /* 今はなし */ }
