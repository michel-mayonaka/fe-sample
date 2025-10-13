// Package input は UI 層で扱う論理入力（イベント）を表す型群を提供します。
package input

// EventKind は素の入力イベント種別です。
type EventKind int

const (
    // EventUnknown は未定義のイベントです。
    EventUnknown EventKind = iota
    // EventKey はキーボード入力イベントです。
    EventKey
    // EventMouseButton はマウスボタンの押下/解放イベントです。
    EventMouseButton
    // EventMouseWheel はマウスホイールの回転イベントです。
    EventMouseWheel
    // EventGamepadButton はゲームパッドのボタン押下/解放イベントです。
    EventGamepadButton
    // EventGamepadAxis はゲームパッドのアナログ軸イベントです。
    EventGamepadAxis
)

// Modifier は修飾キーのビットフラグです。
type Modifier uint8

const (
    // ModNone は修飾なしを表します。
    ModNone  Modifier = 0
    // ModShift は Shift キー修飾を表します。
    ModShift Modifier = 1 << 0
    // ModCtrl は Ctrl キー修飾を表します。
    ModCtrl  Modifier = 1 << 1
    // ModAlt は Alt キー修飾を表します。
    ModAlt   Modifier = 1 << 2
    // ModMeta は Meta(Command/Windows) キー修飾を表します。
    ModMeta  Modifier = 1 << 3
)

// Event はフレーム内で発生した素の入力イベントです。
// Code はデバイス固有コードを抽象化した整数です（アダプタ側で割当）。
// Value はアナログ量（キー/ボタンは 0 or 1）。
type Event struct {
    Kind  EventKind
    Code  int
    Value float64
    Mods  Modifier
}
