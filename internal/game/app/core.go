package app

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"ui_sample/internal/config"
	cuim "ui_sample/internal/config/uimetrics"
	"ui_sample/internal/game"
	gdata "ui_sample/internal/game/data"
	"ui_sample/internal/game/scenes"
	characterlist "ui_sample/internal/game/scenes/character_list"
	gamesvc "ui_sample/internal/game/service"
	uicore "ui_sample/internal/game/service/ui"
	uinput "ui_sample/internal/game/ui/input"
	"ui_sample/internal/repo"
	"ui_sample/internal/usecase"
)

// NewUIAppGame は UI サンプル用にポートを注入し SceneStack を構築した ebiten.Game を返します。
func NewUIAppGame() *Game {
	// 乱数と入力
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	in := gamesvc.NewInput()
	in.BindKey(ebiten.KeyBackspace, gamesvc.Menu)

	// ユーザパス/テーブル
	userPath := config.DefaultUserPath
	// 一覧
	units, _ := uicore.LoadUnitsFromUser(userPath)
	if len(units) == 0 {
		units = []uicore.Unit{uicore.SampleUnit()}
	}

	// Ports（JSON）を注入して App を生成
	urepo, _ := repo.NewJSONUserRepo(userPath)
	wrepo, _ := repo.NewJSONWeaponsRepo(config.DefaultWeaponsPath)
	inv, _ := repo.NewJSONInventoryRepo(config.DefaultUserWeaponsPath, config.DefaultUserItemsPath, config.DefaultWeaponsPath, "db/master/mst_items.json")
	a := usecase.New(urepo, wrepo, inv, rng)
	// 推奨: プロバイダ経由でテーブルをDI
	gdata.SetProvider(a)
	// 互換呼び出しは不要になったため省略（Provider経由で参照）

	// 画面メトリクス初期化（基準解像度 + 外部メトリクス適用）
	uicore.SetBaseResolution(screenW, screenH)
	// 外部メトリクスの読み込み（ユーザ→マスタ→既定）
	if m := cuim.LoadOrDefault(config.DefaultUserUIMetricsPath, config.DefaultUIMetricsPath); true {
		// Apply to UI core
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
		// 基準解像度（任意上書き）
		if m.Base.W > 0 && m.Base.H > 0 {
			uicore.SetBaseResolution(m.Base.W, m.Base.H)
		}
	}

	// Env（共有状態）
	env := &scenes.Env{
		Data:     a,
		Battle:   a,
		Inv:      a,
		UserPath: userPath,
		RNG:      rng,
		Session:  &scenes.Session{Units: units, SelIndex: 0},
	}

	// Game（Runner + AfterUpdate）
	g := &Game{Runner: Runner{}, Input: in, InputR: uinput.WrapService(in), Env: env, prevTime: time.Now()}
	g.Runner.AfterUpdate = func(sc game.Scene) bool {
		if p, ok := sc.(interface{ ShouldPop() bool }); ok {
			return p.ShouldPop()
		}
		return false
	}
	g.Runner.Stack.Push(characterlist.NewList(env))

	// ウィンドウ・TPS
	SetupWindow()
	ebiten.SetWindowTitle("Ebiten UI サンプル - ステータス画面")
	return g
}
