// Package main は Ebiten を用いた FE 風ステータスUIサンプルの
// エントリポイントを提供します。
package main

import (
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/ui"
)

const (
    // screenW は論理解像度の横幅（ピクセル）です。
    screenW = 1920
    // screenH は論理解像度の縦幅（ピクセル）です。
    screenH = 1080
)

// Game はゲーム状態（UI表示用）を保持します。
type Game struct {
    showHelp bool   // ヘルプ表示フラグ
    unit     ui.Unit // 表示対象ユニット
}

// NewGame は Game を初期化して返します。
func NewGame() *Game {
    g := &Game{}
    g.unit = ui.SampleUnit()
    return g
}

// Update は毎フレームの更新処理を行います。
func (g *Game) Update() error {
    if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
        g.unit = ui.SampleUnit()
    }
    if ebiten.IsKeyPressed(ebiten.KeyH) {
        g.showHelp = true
    } else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
        g.showHelp = false
    }
    return nil
}

// Draw は画面描画を行います。
func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{12, 18, 30, 255})
    ui.DrawStatus(screen, g.unit)
    if g.showHelp {
        ebitenutil.DebugPrintAt(screen, "H: ヘルプ表示切替 / ESC: 閉じる\nBackspace: サンプル値を再読み込み", 16, screenH-64)
    }
}

// Layout は論理解像度（内部解像度）を返します。
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenW, screenH
}

// main はウィンドウを作成しゲームループを開始します。
func main() {
    ebiten.SetWindowSize(screenW, screenH)
    ebiten.SetWindowTitle("Ebiten UI サンプル - ステータス画面")
    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    if err := ebiten.RunGame(NewGame()); err != nil {
        panic(err)
    }
}
