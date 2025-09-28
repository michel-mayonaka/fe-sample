package render

import (
    "github.com/hajimehoshi/ebiten/v2"
    "ui_sample/internal/game/actor"
)

// DrawLayer は指定したレイヤ範囲の Actor を描画します。
func DrawLayer(dst *ebiten.Image, actors []actor.IActor, minLayer, maxLayer int) {
    for _, a := range actors {
        if l := a.Layer(); l >= minLayer && l < maxLayer {
            a.Draw(dst)
        }
    }
}
