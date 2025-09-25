// Package main は Ebiten を用いた FE 風ステータスUIサンプルの
// エントリポイントを提供します。
package main

import (
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
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

    // 一覧/詳細の画面モード
    mode       screenMode
    units      []ui.Unit
    selIndex   int
    hoverIndex int
}

type screenMode int

const (
    modeList screenMode = iota
    modeStatus
)

func pointIn(px, py, x, y, w, h int) bool {
    return px >= x && py >= y && px < x+w && py < y+h
}

// NewGame は Game を初期化して返します。
func NewGame() *Game {
    g := &Game{}
    // ユーザテーブルから一覧を読み込む
    if us, err := ui.LoadUnitsFromUser("db/user/usr_characters.json"); err == nil && len(us) > 0 {
        g.units = us
        g.selIndex = 0
        g.unit = us[0]
    } else {
        // フォールバック
        g.unit = ui.SampleUnit()
        g.units = []ui.Unit{g.unit}
        g.selIndex = 0
    }
    g.mode = modeList
    g.hoverIndex = -1
    return g
}

// Update は毎フレームの更新処理を行います。
func (g *Game) Update() error {
    if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
        // データ再読み込み
        if us, err := ui.LoadUnitsFromUser("db/user/usr_characters.json"); err == nil && len(us) > 0 {
            g.units = us
            if g.selIndex >= len(us) { g.selIndex = 0 }
            g.unit = us[g.selIndex]
        } else {
            g.unit = ui.SampleUnit()
            g.units = []ui.Unit{g.unit}
            g.selIndex = 0
        }
    }
    if ebiten.IsKeyPressed(ebiten.KeyH) {
        g.showHelp = true
    } else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
        g.showHelp = false
    }

    // 入力（マウス）
    mx, my := ebiten.CursorPosition()
    switch g.mode {
    case modeList:
        g.hoverIndex = -1
        for i := range g.units {
            x, y, w, h := ui.ListItemRect(screenW, screenH, i)
            if pointIn(mx, my, x, y, w, h) { g.hoverIndex = i }
        }
        if g.hoverIndex >= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.selIndex = g.hoverIndex
            g.unit = g.units[g.selIndex]
            g.mode = modeStatus
        }
    case modeStatus:
        bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
        if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.mode = modeList
        }
    }
    return nil
}

// Draw は画面描画を行います。
func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{12, 18, 30, 255})
    switch g.mode {
    case modeList:
        ui.DrawCharacterList(screen, g.units, g.hoverIndex)
    case modeStatus:
        ui.DrawStatus(screen, g.unit)
        // 戻るボタン
        mx, my := ebiten.CursorPosition()
        bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
        hovered := pointIn(mx, my, bx, by, bw, bh)
        ui.DrawBackButton(screen, hovered)
    }
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
