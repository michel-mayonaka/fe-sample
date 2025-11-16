package app

import (
    "image/color"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/assets"
    "ui_sample/internal/config"
    cuim "ui_sample/internal/config/uimetrics"
    "ui_sample/internal/game"
    "ui_sample/internal/game/scenes"
    uicore "ui_sample/internal/game/service/ui"
    uiadapter "ui_sample/internal/game/ui/adapter"
    uinput "ui_sample/internal/game/ui/input"
    ginput "ui_sample/pkg/game/input"
)

const (
	screenW = 1920
	screenH = 1080
)

// Game は ebiten.Game を実装し、SceneStack とコンテキスト更新、ウィンドウ管理を行います。
type Game struct {
    Runner     Runner
    InputSrc   ginput.Source
    Edge       ginput.EdgeReader
    InputR     uinput.Reader
    Env        *scenes.Env
    prevTime   time.Time
    frame      uint64
    showHelp   bool
	reloadHold int
}

// Update は1フレーム更新します。
func (g *Game) Update() error {
    var cx, cy int
    if g.InputSrc != nil {
        g.Edge.Step(g.InputSrc.Poll())
        if p, ok := g.InputSrc.(interface{ Position() (int, int) }); ok {
            cx, cy = p.Position()
        }
    }
    g.updateGlobalToggles()

	now := time.Now()
	var dt float64
	if !g.prevTime.IsZero() {
		dt = now.Sub(g.prevTime).Seconds()
	}
	g.prevTime = now
	g.frame++
	// Reader（UI向けインターフェース）をコンテキストへ供給
    if g.InputR == nil {
        g.InputR = uinput.WrapDomain(&g.Edge)
    }
    ctx := &game.Ctx{ScreenW: screenW, ScreenH: screenH, CursorX: cx, CursorY: cy, Input: g.InputR, DT: dt, Frame: g.frame}
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
	if g.InputR != nil && g.InputR.Down(uinput.Menu) {
		g.reloadHold++
		if g.reloadHold == 30 { // 約0.5秒（60FPS時）
			if g.Env != nil && g.Env.Data != nil {
				_ = g.Env.Data.ReloadData()
				// 画像キャッシュはUI側の責務としてここでクリア
				assets.Clear()
			}
			// UIメトリクスの再読み込み（ユーザ→マスタ→既定）
			{
				m := cuim.LoadOrDefault(config.DefaultUserUIMetricsPath, config.DefaultUIMetricsPath)
				um := uicore.Metrics{}
				um.List.Margin = m.List.Margin
				um.List.ItemH = m.List.ItemH
				um.List.ItemGap = m.List.ItemGap
				um.List.PortraitSize = m.List.PortraitSize
				um.List.TitleOffset = m.List.TitleOffset
				um.List.HeaderTopGap = m.List.HeaderTopGap
				um.List.ItemsTopGap = m.List.ItemsTopGap
				um.List.PanelInnerPaddingX = m.List.PanelInnerPaddingX
				um.List.TitleXOffset = m.List.TitleXOffset
				um.List.HeaderBaseX = m.List.HeaderBaseX
				um.List.RowTextOffsetX = m.List.RowTextOffsetX
				um.List.RowTextOffsetY = m.List.RowTextOffsetY
				um.List.RowBorderPad = m.List.RowBorderPad
				um.List.RowRightIconSize = m.List.RowRightIconSize
				um.List.RowRightIconGap = m.List.RowRightIconGap
				um.List.HeaderColumnsItems = append([]int(nil), m.List.HeaderColumnsItems...)
				um.List.HeaderColumnsWeapons = append([]int(nil), m.List.HeaderColumnsWeapons...)
				um.List.RowColumnsItems = append([]int(nil), m.List.RowColumnsItems...)
				um.List.RowColumnsWeapons = append([]int(nil), m.List.RowColumnsWeapons...)
				um.Line.Main = m.Line.Main
				um.Line.Small = m.Line.Small
				// Status
				um.Status.PanelPad = m.Status.PanelPad
				um.Status.PortraitSize = m.Status.PortraitSize
				um.Status.TextGapX = m.Status.TextGapX
				um.Status.NameOffsetY = m.Status.NameOffsetY
				um.Status.ClassGapFromName = m.Status.ClassGapFromName
				um.Status.LevelGapFromName = m.Status.LevelGapFromName
				um.Status.HPGapFromName = m.Status.HPGapFromName
				um.Status.HPBarGapFromName = m.Status.HPBarGapFromName
				um.Status.HPBarW = m.Status.HPBarW
				um.Status.HPBarH = m.Status.HPBarH
				um.Status.StatsTopGap = m.Status.StatsTopGap
				um.Status.StatsLineH = m.Status.StatsLineH
				um.Status.StatsColGap = m.Status.StatsColGap
				um.Status.WeaponRanksXExtra = m.Status.WeaponRanksXExtra
				um.Status.RankLineH = m.Status.RankLineH
				um.Status.MagicRanksTopExtra = m.Status.MagicRanksTopExtra
				um.Status.EquipTitleGapY = m.Status.EquipTitleGapY
				um.Status.EquipLineH = m.Status.EquipLineH
				um.Status.EquipRectYOffset = m.Status.EquipRectYOffset
				um.Status.EquipRectW = m.Status.EquipRectW
				um.Status.EquipRectH = m.Status.EquipRectH
				um.Status.EquipLabelGapX = m.Status.EquipLabelGapX
				um.Status.EquipUsesX = m.Status.EquipUsesX
				// Sim
				um.Sim.StartBtnW = m.Sim.StartBtnW
				um.Sim.StartBtnH = m.Sim.StartBtnH
				um.Sim.AutoRunGap = m.Sim.AutoRunGap
				um.Sim.TitleYOffset = m.Sim.TitleYOffset
				um.Sim.TitleXOffsetFromCenter = m.Sim.TitleXOffsetFromCenter
				um.Sim.Terrain.ButtonW = m.Sim.Terrain.ButtonW
				um.Sim.Terrain.ButtonH = m.Sim.Terrain.ButtonH
				um.Sim.Terrain.BaseYFromBottom = m.Sim.Terrain.BaseYFromBottom
				um.Sim.Terrain.LeftBaseXOffset = m.Sim.Terrain.LeftBaseXOffset
				um.Sim.Terrain.RightBaseXInset = m.Sim.Terrain.RightBaseXInset
				um.Sim.Terrain.ButtonGap = m.Sim.Terrain.ButtonGap
				um.Sim.Terrain.LabelLeftXOffset = m.Sim.Terrain.LabelLeftXOffset
				um.Sim.Terrain.LabelYOffsetFromBottom = m.Sim.Terrain.LabelYOffsetFromBottom
				uicore.ApplyMetrics(um)
				if m.Base.W > 0 && m.Base.H > 0 {
					uicore.SetBaseResolution(m.Base.W, m.Base.H)
				}
			}
            // UIユニット再構築（Provider 経由）
            if g.Env != nil {
                if us := uiadapter.BuildUnitsFromProvider(uiadapter.AssetsPortraitLoader{}); len(us) > 0 {
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
	} else if g.InputR != nil && g.InputR.Down(uinput.Cancel) {
		g.showHelp = false
	}
}

// SetupWindow はTPS/ウィンドウ設定を行います。
func SetupWindow() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(60)
}
