// Package actor は描画対象（Actor）の最小契約を提供します。
package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"ui_sample/internal/game"
)

// IActor は更新・描画・レイヤのインタフェースです。
// Update が false を返した場合、呼び出し側で破棄対象と見なします。
type IActor interface {
	Update(ctx *game.Ctx) bool
	Draw(dst *ebiten.Image)
	Layer() int // 描画順（例: 世界=100, FX=200, UI=300）
}
