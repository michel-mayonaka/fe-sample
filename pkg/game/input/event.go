package input

// EventKind は素の入力イベント種別です。
type EventKind int

const (
    // EventUnknown は未定義のイベントです。
    EventUnknown EventKind = iota
    EventKey
    EventMouseButton
    EventMouseWheel
    EventGamepadButton
    EventGamepadAxis
)

// Modifier は修飾キーのビットフラグです。
type Modifier uint8

const (
    ModNone  Modifier = 0
    ModShift Modifier = 1 << 0
    ModCtrl  Modifier = 1 << 1
    ModAlt   Modifier = 1 << 2
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

