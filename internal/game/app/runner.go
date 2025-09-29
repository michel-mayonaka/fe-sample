// Package app はゲームループ心臓部（SceneStack 管理）を提供します。
package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"ui_sample/internal/game"
)

// Runner は SceneStack の更新と描画を管理します。
// AfterUpdate が true を返した場合、Pop します（戻り条件の委譲用）。
type Runner struct {
	Stack       game.SceneStack
	AfterUpdate func(sc game.Scene) bool
}

// Update は現在の Scene を更新し、next があれば Push します。
func (r *Runner) Update(ctx *game.Ctx) error {
	sc := r.Stack.Current()
	if sc == nil {
		return nil
	}
	if next, _ := sc.Update(ctx); next != nil {
		r.Stack.Push(next)
	}
	if r.AfterUpdate != nil && r.AfterUpdate(sc) {
		r.Stack.Pop()
	}
	return nil
}

// Draw は現在の Scene を描画します。
func (r *Runner) Draw(dst *ebiten.Image) {
	if sc := r.Stack.Current(); sc != nil {
		sc.Draw(dst)
	}
}
