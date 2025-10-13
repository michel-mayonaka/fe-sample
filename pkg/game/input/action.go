// Package input は UI 層で扱う論理入力（操作）を表す列挙と型を提供します。
package input

// Action は UI が扱う論理的な操作を表します。
// 列挙のゼロ値は Unknown とし、型名接頭辞 + Unknown を 0 に据えます。
type Action int

const (
    // ActionUnknown は未定義の操作を表します（ゼロ値安全）。
    ActionUnknown Action = iota
    // ActionUp はカーソルを上方向へ移動する操作です。
    ActionUp
    // ActionDown はカーソルを下方向へ移動する操作です。
    ActionDown
    // ActionLeft はカーソルを左方向へ移動する操作です。
    ActionLeft
    // ActionRight はカーソルを右方向へ移動する操作です。
    ActionRight
    // ActionConfirm は決定/実行の操作です。
    ActionConfirm
    // ActionCancel はキャンセル/戻るの操作です。
    ActionCancel
    // ActionMenu はメニューを開く操作です。
    ActionMenu
    // ActionNext は次候補/次ページへ進む操作です。
    ActionNext
    // ActionPrev は前候補/前ページへ戻る操作です。
    ActionPrev
    // ActionOpenWeapons は武器一覧を開く操作です。
    ActionOpenWeapons
    // ActionOpenItems はアイテム一覧を開く操作です。
    ActionOpenItems
    // ActionEquipToggle は装備/解除を切り替える操作です。
    ActionEquipToggle
    // ActionSlot1 はスロット1を選択する操作です。
    ActionSlot1
    // ActionSlot2 はスロット2を選択する操作です。
    ActionSlot2
    // ActionSlot3 はスロット3を選択する操作です。
    ActionSlot3
    // ActionSlot4 はスロット4を選択する操作です。
    ActionSlot4
    // ActionSlot5 はスロット5を選択する操作です。
    ActionSlot5
    // ActionUnassign は装備を外す操作です。
    ActionUnassign
    // ActionTerrainAtt1 は地形攻撃設定1へ切替える操作です。
    ActionTerrainAtt1
    // ActionTerrainAtt2 は地形攻撃設定2へ切替える操作です。
    ActionTerrainAtt2
    // ActionTerrainAtt3 は地形攻撃設定3へ切替える操作です。
    ActionTerrainAtt3
    // ActionTerrainDef1 は地形防御設定1へ切替える操作です。
    ActionTerrainDef1
    // ActionTerrainDef2 は地形防御設定2へ切替える操作です。
    ActionTerrainDef2
    // ActionTerrainDef3 は地形防御設定3へ切替える操作です。
    ActionTerrainDef3
    // ActionCount は内部配列確保などで用いる上限です。
    ActionCount
)
