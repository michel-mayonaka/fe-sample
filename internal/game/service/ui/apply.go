//go:build !headless

package uicore

// Metrics は UI レイアウトに関わるメトリクスのセットです。
// 単位はいずれも px（論理解像度基準の値）。
type Metrics struct {
	List struct {
		Margin               int
		ItemH                int
		ItemGap              int
		PortraitSize         int
		TitleOffset          int
		HeaderTopGap         int
		ItemsTopGap          int
		PanelInnerPaddingX   int
		TitleXOffset         int
		HeaderBaseX          int
		RowTextOffsetX       int
		RowTextOffsetY       int
		RowBorderPad         int
		RowRightIconSize     int
		RowRightIconGap      int
		HeaderColumnsItems   []int
		HeaderColumnsWeapons []int
		RowColumnsItems      []int
		RowColumnsWeapons    []int
	}
	Line struct {
		Main  int
		Small int
	}
	Status struct {
		PanelPad           int
		PortraitSize       int
		TextGapX           int
		NameOffsetY        int
		ClassGapFromName   int
		LevelGapFromName   int
		HPGapFromName      int
		HPBarGapFromName   int
		HPBarW             int
		HPBarH             int
		StatsTopGap        int
		StatsLineH         int
		StatsColGap        int
		WeaponRanksXExtra  int
		RankLineH          int
		MagicRanksTopExtra int
		EquipTitleGapY     int
		EquipLineH         int
		EquipRectYOffset   int
		EquipRectW         int
		EquipRectH         int
		EquipLabelGapX     int
		EquipUsesX         int
	}
	Sim struct {
		StartBtnW              int
		StartBtnH              int
		AutoRunGap             int
		TitleYOffset           int
		TitleXOffsetFromCenter int
		Terrain                struct {
			ButtonW                int
			ButtonH                int
			BaseYFromBottom        int
			LeftBaseXOffset        int
			RightBaseXInset        int
			ButtonGap              int
			LabelLeftXOffset       int
			LabelYOffsetFromBottom int
		}
		Preview struct {
			LeftXPad       int
			RightXInset    int
			TopYFromMargin int
			CardW          int
			CardH          int
			CardInnerPad   int
			PortraitSize   int
			ClassOffsetY   int
			HPLabelX       int
			HPLabelY       int
			HPBarX         int
			HPBarY         int
			HPBarW         int
			HPBarH         int
			NameOffsetX    int
			NameOffsetY    int
			LineY          int
			BaseMin        int
			BaseMax        int
			WrapPad        int
			LogPadX        int
			LogPadY        int
		}
	}
	Popup   struct{ Cols4ThresholdW, Cols3ThresholdW, GridInnerXTotalPad, CellH, CellGap, GridXPad, GridYOff, MaxW, MaxH int }
	Widgets struct {
		Back      struct{ PanelRightInset, TopPad, W, H, LabelX, LabelY int }
		LevelUp   struct{ W, H, LabelX, LabelY int }
		ToBattle  struct{ GapFromRightBtn, LabelX, LabelY, W int }
		SimBattle struct{ W, H, TopPad, LabelX, LabelY int }
	}
}

// DefaultMetrics は現在のビルトイン既定値を返します。
func DefaultMetrics() Metrics {
	var m Metrics
	m.List.Margin = ListMargin
	m.List.ItemH = ListItemH
	m.List.ItemGap = ListItemGap
	m.List.PortraitSize = ListPortraitSz
	m.List.TitleOffset = ListTitleOffset
	m.List.HeaderTopGap = ListHeaderTopGap
	m.List.ItemsTopGap = ListItemsTopGap
	m.List.PanelInnerPaddingX = ListPanelInnerPaddingX
	m.List.TitleXOffset = ListTitleXOffset
	m.List.HeaderBaseX = ListHeaderBaseX
	m.List.RowTextOffsetX = ListRowTextOffsetX
	m.List.RowTextOffsetY = ListRowTextOffsetY
	m.List.RowBorderPad = ListRowBorderPad
	m.List.RowRightIconSize = ListRowRightIconSize
	m.List.RowRightIconGap = ListRowRightIconGap
	m.List.HeaderColumnsItems = copyInts(ListHeaderColumnsItems)
	m.List.HeaderColumnsWeapons = copyInts(ListHeaderColumnsWeapons)
	m.List.RowColumnsItems = copyInts(ListRowColumnsItems)
	m.List.RowColumnsWeapons = copyInts(ListRowColumnsWeapons)
	m.Line.Main = LineHMain
	m.Line.Small = LineHSmall
	// Status
	m.Status.PanelPad = StatusPanelPad
	m.Status.PortraitSize = StatusPortraitSize
	m.Status.TextGapX = StatusTextGapX
	m.Status.NameOffsetY = StatusNameOffsetY
	m.Status.ClassGapFromName = StatusClassGapFromName
	m.Status.LevelGapFromName = StatusLevelGapFromName
	m.Status.HPGapFromName = StatusHPGapFromName
	m.Status.HPBarGapFromName = StatusHPBarGapFromName
	m.Status.HPBarW = StatusHPBarW
	m.Status.HPBarH = StatusHPBarH
	m.Status.StatsTopGap = StatusStatsTopGap
	m.Status.StatsLineH = StatusStatsLineH
	m.Status.StatsColGap = StatusStatsColGap
	m.Status.WeaponRanksXExtra = StatusWeaponRanksXExtra
	m.Status.RankLineH = StatusRankLineH
	m.Status.MagicRanksTopExtra = StatusMagicRanksTopExtra
	m.Status.EquipTitleGapY = StatusEquipTitleGapY
	m.Status.EquipLineH = StatusEquipLineH
	m.Status.EquipRectYOffset = StatusEquipRectYOffset
	m.Status.EquipRectW = StatusEquipRectW
	m.Status.EquipRectH = StatusEquipRectH
	m.Status.EquipLabelGapX = StatusEquipLabelGapX
	m.Status.EquipUsesX = StatusEquipUsesX
	// Sim
	m.Sim.StartBtnW = SimStartBtnW
	m.Sim.StartBtnH = SimStartBtnH
	m.Sim.AutoRunGap = SimAutoRunGap
	m.Sim.TitleYOffset = SimTitleYOffset
	m.Sim.TitleXOffsetFromCenter = SimTitleXOffsetFromCenter
	m.Sim.Terrain.ButtonW = TerrainBtnW
	m.Sim.Terrain.ButtonH = TerrainBtnH
	m.Sim.Terrain.BaseYFromBottom = TerrainBaseYFromBottom
	m.Sim.Terrain.LeftBaseXOffset = TerrainLeftBaseXOffset
	m.Sim.Terrain.RightBaseXInset = TerrainRightBaseXInset
	m.Sim.Terrain.ButtonGap = TerrainBtnGap
	m.Sim.Terrain.LabelLeftXOffset = TerrainLabelLeftXOffset
	m.Sim.Terrain.LabelYOffsetFromBottom = TerrainLabelYOffsetFromBottom
	// Sim Preview
	m.Sim.Preview.LeftXPad = SimPreviewLeftXPad
	m.Sim.Preview.RightXInset = SimPreviewRightXInset
	m.Sim.Preview.TopYFromMargin = SimPreviewTopYFromMargin
	m.Sim.Preview.CardW = SimPreviewCardW
	m.Sim.Preview.CardH = SimPreviewCardH
	m.Sim.Preview.CardInnerPad = SimPreviewCardInnerPad
	m.Sim.Preview.PortraitSize = SimPreviewPortraitSize
	m.Sim.Preview.ClassOffsetY = SimPreviewClassOffsetY
	m.Sim.Preview.HPLabelX = SimPreviewHPLabelX
	m.Sim.Preview.HPLabelY = SimPreviewHPLabelY
	m.Sim.Preview.HPBarX = SimPreviewHPBarX
	m.Sim.Preview.HPBarY = SimPreviewHPBarY
	m.Sim.Preview.HPBarW = SimPreviewHPBarW
	m.Sim.Preview.HPBarH = SimPreviewHPBarH
	m.Sim.Preview.NameOffsetX = SimPreviewNameOffsetX
	m.Sim.Preview.NameOffsetY = SimPreviewNameOffsetY
	m.Sim.Preview.LineY = SimPreviewLineY
	m.Sim.Preview.BaseMin = SimPreviewBaseMin
	m.Sim.Preview.BaseMax = SimPreviewBaseMax
	m.Sim.Preview.WrapPad = SimPreviewWrapPad
	m.Sim.Preview.LogPadX = SimPreviewLogPadX
	m.Sim.Preview.LogPadY = SimPreviewLogPadY
	// Popup defaults
	m.Popup.Cols4ThresholdW = PopupCols4ThresholdW
	m.Popup.Cols3ThresholdW = PopupCols3ThresholdW
	m.Popup.GridInnerXTotalPad = PopupGridInnerXTotalPad
	m.Popup.CellH = PopupCellH
	m.Popup.CellGap = PopupCellGap
	m.Popup.GridXPad = PopupGridXPad
	m.Popup.GridYOff = PopupGridYOff
	m.Popup.MaxW = PopupMaxW
	m.Popup.MaxH = PopupMaxH
	// Widgets defaults
	m.Widgets.Back.PanelRightInset = WidgetsBackPanelRightInset
	m.Widgets.Back.TopPad = WidgetsBackTopPad
	m.Widgets.Back.W = WidgetsBackW
	m.Widgets.Back.H = WidgetsBackH
	m.Widgets.Back.LabelX = WidgetsBackLabelX
	m.Widgets.Back.LabelY = WidgetsBackLabelY
	m.Widgets.LevelUp.W = WidgetsLevelUpW
	m.Widgets.LevelUp.H = WidgetsLevelUpH
	m.Widgets.LevelUp.LabelX = WidgetsLevelUpLabelX
	m.Widgets.LevelUp.LabelY = WidgetsLevelUpLabelY
	m.Widgets.ToBattle.GapFromRightBtn = WidgetsToBattleGapFromRightBtn
	m.Widgets.ToBattle.W = WidgetsToBattleW
	m.Widgets.ToBattle.LabelX = WidgetsToBattleLabelX
	m.Widgets.ToBattle.LabelY = WidgetsToBattleLabelY
	m.Widgets.SimBattle.W = WidgetsSimBattleW
	m.Widgets.SimBattle.H = WidgetsSimBattleH
	m.Widgets.SimBattle.TopPad = WidgetsSimBattleTopPad
	m.Widgets.SimBattle.LabelX = WidgetsSimBattleLabelX
	m.Widgets.SimBattle.LabelY = WidgetsSimBattleLabelY
	return m
}

// ApplyMetrics は与えられたメトリクスを UI 既定値へ適用します。
func ApplyMetrics(m Metrics) {
	targets := newMetricsTargets()
	applyList(targets, m)
	applyStatus(targets, m)
	applySim(targets, m)
	applyPopup(targets, m)
	applyWidgets(targets, m)
}
