package main

import (
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/ui"
)

const (
    screenW = 1920
    screenH = 1080
)

type Game struct {
    showHelp bool
    unit     ui.Unit
}

func NewGame() *Game {
    g := &Game{}
    g.unit = ui.SampleUnit()
    return g
}

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

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{12, 18, 30, 255})
    ui.DrawStatus(screen, g.unit)
    if g.showHelp {
        ebitenutil.DebugPrintAt(screen, "H: ヘルプ表示切替 / ESC: 閉じる\nBackspace: サンプル値を再読み込み", 16, screenH-64)
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenW, screenH
}

func main() {
    ebiten.SetWindowSize(screenW, screenH)
    ebiten.SetWindowTitle("Ebiten UI Sample - FE Status Style")
    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    if err := ebiten.RunGame(NewGame()); err != nil {
        panic(err)
    }
}
