// Package main は Ebiten を用いた FE 風ステータスUIサンプルの
// エントリポイントを提供します。
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	gameapp "ui_sample/internal/game/app"
)

// main はウィンドウを作成しゲームループを開始します。
func main() {
	if err := ebiten.RunGame(gameapp.NewUIAppGame()); err != nil {
		panic(err)
	}
}
