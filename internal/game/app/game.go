package app

import (
    "image/color"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    "ui_sample/internal/game/scenes"
    uicore "ui_sample/internal/game/service/ui"
)

const (
    screenW = 1920
    screenH = 1080
)

// Game は ebiten.Game を実装し、SceneStack とコンテキスト更新、ウィンドウ管理を行います。
type Game struct {
    Runner    Runner
    Input     *gamesvc.Input
    Env       *scenes.Env
    prevTime  time.Time
    frame     uint64
    showHelp  bool
    reloadHold int
}

// Update は1フレーム更新します。
func (g *Game) Update() error {
    if g.Input != nil { g.Input.Snapshot() }
    g.updateGlobalToggles()

    now := time.Now()
    var dt float64
    if !g.prevTime.IsZero() { dt = now.Sub(g.prevTime).Seconds() }
    g.prevTime = now
    g.frame++
    ctx := &game.Ctx{ScreenW: screenW, ScreenH: screenH, Input: g.Input, DT: dt, Frame: g.frame}
    return g.Runner.Update(ctx)
}

// Draw は現在の Scene を描画します。
func (g *Game) Draw(screen *ebiten.Image) {
    // メトリクスとフォント
    uicore.UpdateMetricsFromWindow()
    uicore.MaybeUpdateFontFaces()
    screen.Fill(color.RGBA{12, 18, 30, 255})
    g.Runner.Draw(screen)
    if g.showHelp {
        ebitenutil.DebugPrintAt(screen, "H: ヘルプ表示 / ESC: 閉じる\nBackspace長押し: 再読み込み", 16, screenH-64)
    }
}

// Layout は論理解像度を返します。
func (*Game) Layout(_, _ int) (int, int) { return screenW, screenH }

// updateGlobalToggles はヘルプ表示やデータ再読み込みなどのグローバル操作を処理します。
func (g *Game) updateGlobalToggles() {
    // Backspace(Menu) 長押しでデータ再読み込み
    if g.Input != nil && g.Input.Down(gamesvc.Menu) {
        g.reloadHold++
        if g.reloadHold == 30 { // 約0.5秒（60FPS時）
            if g.Env != nil && g.Env.App != nil {
                _ = g.Env.App.ReloadData()
                scenes.SetWeaponTable(g.Env.App.WeaponsTable())
            }
            // UIユニット再構築
            if g.Env != nil && g.Env.UserPath != "" {
                if us, err := uicore.LoadUnitsFromUser(g.Env.UserPath); err == nil && len(us) > 0 {
                    g.Env.Units = us
                    if g.Env.SelIndex >= len(us) { g.Env.SelIndex = 0 }
                }
            }
        }
    } else {
        g.reloadHold = 0
    }
    // ヘルプ
    if ebiten.IsKeyPressed(ebiten.KeyH) {
        g.showHelp = true
    } else if g.Input != nil && g.Input.Down(gamesvc.Cancel) {
        g.showHelp = false
    }
}

// SetupWindow はTPS/ウィンドウ設定を行います。
func SetupWindow() {
    ebiten.SetWindowSize(screenW, screenH)
    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    ebiten.SetTPS(60)
}
