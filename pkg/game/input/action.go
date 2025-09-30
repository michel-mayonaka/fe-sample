package input

// Action は UI が扱う論理的な操作を表します。
// 列挙のゼロ値は Unknown とし、型名接頭辞 + Unknown を 0 に据えます。
type Action int

const (
    // ActionUnknown は未定義の操作を表します（ゼロ値安全）。
    ActionUnknown Action = iota
    ActionUp
    ActionDown
    ActionLeft
    ActionRight
    ActionConfirm
    ActionCancel
    ActionMenu
    ActionNext
    ActionPrev
    ActionOpenWeapons
    ActionOpenItems
    ActionEquipToggle
    ActionSlot1
    ActionSlot2
    ActionSlot3
    ActionSlot4
    ActionSlot5
    ActionUnassign
    // 地形切替（攻撃/防御）
    ActionTerrainAtt1
    ActionTerrainAtt2
    ActionTerrainAtt3
    ActionTerrainDef1
    ActionTerrainDef2
    ActionTerrainDef3
    // ActionCount は内部配列確保などで用いる上限です。
    ActionCount
)

