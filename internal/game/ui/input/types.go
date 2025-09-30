// Package input は UI 層で扱う抽象入力の最小APIを提供します。
//
// 目的: Scenes からは「入力の意味（Action/Press/Down）」だけを参照し、
// 具象の取得/マッピング実装（Ebiten依存）は他所に隠蔽します。
package input

import (
    gamesvc "ui_sample/internal/game/service"
    ginput "ui_sample/pkg/game/input"
)

// Action はドメイン `pkg/game/input.Action` の別名（段階移行用の公開面）です。
type Action = ginput.Action

// 抽象アクション（ドメイン定義を再公開）
const (
    Up           = ginput.ActionUp
    Down         = ginput.ActionDown
    Left         = ginput.ActionLeft
    Right        = ginput.ActionRight
    Confirm      = ginput.ActionConfirm
    Cancel       = ginput.ActionCancel
    Menu         = ginput.ActionMenu
    Next         = ginput.ActionNext
    Prev         = ginput.ActionPrev
    OpenWeapons  = ginput.ActionOpenWeapons
    OpenItems    = ginput.ActionOpenItems
    EquipToggle  = ginput.ActionEquipToggle
    Slot1        = ginput.ActionSlot1
    Slot2        = ginput.ActionSlot2
    Slot3        = ginput.ActionSlot3
    Slot4        = ginput.ActionSlot4
    Slot5        = ginput.ActionSlot5
    Unassign     = ginput.ActionUnassign
    TerrainAtt1  = ginput.ActionTerrainAtt1
    TerrainAtt2  = ginput.ActionTerrainAtt2
    TerrainAtt3  = ginput.ActionTerrainAtt3
    TerrainDef1  = ginput.ActionTerrainDef1
    TerrainDef2  = ginput.ActionTerrainDef2
    TerrainDef3  = ginput.ActionTerrainDef3
)

// Reader はフレームスナップショットに対する読み取り最小APIです。
// Press: 立ち上がり検出 / Down: 押下継続。
type Reader interface {
    Press(Action) bool
    Down(Action) bool
}

// ServiceAdapter は既存の service.Input を Reader に適合させます。
type ServiceAdapter struct{ S *gamesvc.Input }

// WrapService は service.Input を Reader として扱うアダプタを返します。
func WrapService(s *gamesvc.Input) Reader { return ServiceAdapter{S: s} }

// DomainAdapter はドメインの Reader を UI の Reader として扱う薄いアダプタです。
type DomainAdapter struct{ R ginput.Reader }

// WrapDomain は ginput.Reader を uinput.Reader として扱うアダプタを返します。
func WrapDomain(r ginput.Reader) Reader { return DomainAdapter{R: r} }

// Press は uinput.Action を service.Action に変換して委譲します。
func (a ServiceAdapter) Press(act Action) bool {
    if a.S == nil {
        return false
    }
    if sa, ok := toServiceAction(act); ok {
        return a.S.Press(sa)
    }
    return false
}

// Down は uinput.Action を service.Action に変換して委譲します。
func (a ServiceAdapter) Down(act Action) bool {
    if a.S == nil {
        return false
    }
    if sa, ok := toServiceAction(act); ok {
        return a.S.Down(sa)
    }
    return false
}

// Press はそのまま委譲します（型は別名のため変換不要）。
func (a DomainAdapter) Press(act Action) bool {
    if a.R == nil {
        return false
    }
    return a.R.Press(act)
}

// Down はそのまま委譲します（型は別名のため変換不要）。
func (a DomainAdapter) Down(act Action) bool {
    if a.R == nil {
        return false
    }
    return a.R.Down(act)
}

// toServiceAction はドメインの Action を旧 service.Action へ変換します（段階移行用）。
func toServiceAction(a Action) (gamesvc.Action, bool) {
    if int(a) <= 0 {
        return 0, false
    }
    v := int(a) - 1 // Unknown=0 を詰めて互換列挙へ
    if v < 0 || v >= int(gamesvc.ActionCount) {
        return 0, false
    }
    return gamesvc.Action(v), true
}
