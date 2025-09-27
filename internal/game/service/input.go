// Package service は UI 入出力・資産・音などの周辺サービスを提供します。
package service

import (
    "github.com/hajimehoshi/ebiten/v2"
)

// Action は抽象アクション種別です。
type Action int

// 抽象アクション（最小）
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

// Input はフレームごとのスナップショットを保持します。
type Input struct {
    curr, prev [ActionCount]bool
    mapKey     map[ebiten.Key]Action
}

// NewInput は既定のキー割り当てで初期化します。
func NewInput() *Input {
    m := map[ebiten.Key]Action{
        ebiten.KeyArrowUp:    Up,
        ebiten.KeyArrowDown:  Down,
        ebiten.KeyArrowLeft:  Left,
        ebiten.KeyArrowRight: Right,
        ebiten.KeyZ:          Confirm,
        ebiten.KeyEnter:      Confirm,
        ebiten.KeyX:          Cancel,
        ebiten.KeyEscape:     Cancel,
        ebiten.KeyA:          Prev,
        ebiten.KeyS:          Next,
        ebiten.KeyTab:        Menu,
        // 拡張ショートカット（UI便宜）
        ebiten.KeyW:          OpenWeapons,
        ebiten.KeyI:          OpenItems,
        ebiten.KeyE:          EquipToggle,
        ebiten.Key1:          Slot1,
        ebiten.Key2:          Slot2,
        ebiten.Key3:          Slot3,
        ebiten.Key4:          Slot4,
        ebiten.Key5:          Slot5,
        ebiten.KeyDelete:     Unassign,
    }
    return &Input{mapKey: m}
}

// BindKey はキーにアクションを割り当てます。
func (i *Input) BindKey(k ebiten.Key, a Action) {
    if i.mapKey == nil {
        i.mapKey = map[ebiten.Key]Action{}
    }
    i.mapKey[k] = a
}

// Snapshot は現在のキー状態を抽象アクションへ投影し、prev/curr を更新します。
func (i *Input) Snapshot() { i.SnapshotWith(ebiten.IsKeyPressed) }

// SnapshotWith はテスト用にキー状態の供給元を差し替え可能にした版です。
func (i *Input) SnapshotWith(isDown func(k ebiten.Key) bool) {
    i.prev = i.curr
    // クリア
    for a := Action(0); a < ActionCount; a++ {
        i.curr[a] = false
    }
    // 一般マッピング
    for k, a := range i.mapKey {
        if isDown(k) {
            i.curr[a] = true
        }
    }
    // 地形切替（1/2/3, Shift+1/2/3 を攻守に割当）
    shift := isDown(ebiten.KeyShift) || isDown(ebiten.KeyShiftLeft) || isDown(ebiten.KeyShiftRight)
    if isDown(ebiten.Key1) {
        if shift { i.curr[TerrainDef1] = true } else { i.curr[TerrainAtt1] = true }
    }
    if isDown(ebiten.Key2) {
        if shift { i.curr[TerrainDef2] = true } else { i.curr[TerrainAtt2] = true }
    }
    if isDown(ebiten.Key3) {
        if shift { i.curr[TerrainDef3] = true } else { i.curr[TerrainAtt3] = true }
    }
}

// Press はこのフレームで押された（立ち上がり）かを返します。
func (i *Input) Press(a Action) bool { return !i.prev[a] && i.curr[a] }

// Down は押下継続中かを返します。
func (i *Input) Down(a Action) bool { return i.curr[a] }
