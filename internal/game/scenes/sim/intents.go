package sim

import (
    "ui_sample/internal/game"
    scenes "ui_sample/internal/game/scenes"
)

// IntentKind は模擬戦シーンで扱う入力意図の種別です。
type IntentKind int

const (
    intentNone IntentKind = iota
    intentBack
    intentRunOne
    intentToggleAuto
    intentSetTerrainAtt // Index: 0..2
    intentSetTerrainDef // Index: 0..2
)

// Intent は入力を意味表現に変換したものです。
type Intent struct{ Kind IntentKind; Index int }
// IsSceneIntent は scenes.Intent 実装のマーカーです。
func (Intent) IsSceneIntent() {}

// scContract はパッケージ内コンパイル保証のためのインターフェースです。
// Sim が必要な sc* メソッドを実装していることを確認します。
type scContract interface{
    scHandleInput(ctx *game.Ctx) []scenes.Intent
    scAdvance([]scenes.Intent)
    scFlush(ctx *game.Ctx)
}
