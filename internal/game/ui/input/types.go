// Package input は UI 層で扱う抽象入力の最小APIを提供します。
//
// 目的: Scenes からは「入力の意味（Action/Press/Down）」だけを参照し、
// 具象の取得/マッピング実装（Ebiten依存）は他所に隠蔽します。
package input

import (
    gamesvc "ui_sample/internal/game/service"
)

// Action は抽象アクション種別です（ui/input 独自定義）。
type Action int

// 抽象アクション（順序は service.Action と同一）
const (
    Up Action = iota
    Down
    Left
    Right
    Confirm
    Cancel
    Menu
    Next
    Prev
    OpenWeapons
    OpenItems
    EquipToggle
    Slot1
    Slot2
    Slot3
    Slot4
    Slot5
    Unassign
    // 地形切替（攻撃側/防御側）
    TerrainAtt1
    TerrainAtt2
    TerrainAtt3
    TerrainDef1
    TerrainDef2
    TerrainDef3
    ActionCount
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

// Press は uinput.Action を service.Action に変換して委譲します。
func (a ServiceAdapter) Press(act Action) bool {
    if a.S == nil { return false }
    return a.S.Press(gamesvc.Action(act))
}

// Down は uinput.Action を service.Action に変換して委譲します。
func (a ServiceAdapter) Down(act Action) bool {
    if a.S == nil { return false }
    return a.S.Down(gamesvc.Action(act))
}
