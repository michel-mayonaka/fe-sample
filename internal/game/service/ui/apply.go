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
	m.List.HeaderColumnsItems = append([]int(nil), ListHeaderColumnsItems...)
	m.List.HeaderColumnsWeapons = append([]int(nil), ListHeaderColumnsWeapons...)
	m.List.RowColumnsItems = append([]int(nil), ListRowColumnsItems...)
	m.List.RowColumnsWeapons = append([]int(nil), ListRowColumnsWeapons...)
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
	if m.List.Margin > 0 {
		ListMargin = m.List.Margin
	}
	if m.List.ItemH > 0 {
		ListItemH = m.List.ItemH
	}
	if m.List.ItemGap > 0 {
		ListItemGap = m.List.ItemGap
	}
	if m.List.PortraitSize > 0 {
		ListPortraitSz = m.List.PortraitSize
	}
	if m.List.TitleOffset > 0 {
		ListTitleOffset = m.List.TitleOffset
	}
	if m.List.HeaderTopGap > 0 {
		ListHeaderTopGap = m.List.HeaderTopGap
	}
	if m.List.ItemsTopGap > 0 {
		ListItemsTopGap = m.List.ItemsTopGap
	}
	if m.List.PanelInnerPaddingX > 0 {
		ListPanelInnerPaddingX = m.List.PanelInnerPaddingX
	}
	if m.List.TitleXOffset > 0 {
		ListTitleXOffset = m.List.TitleXOffset
	}
	if m.List.HeaderBaseX > 0 {
		ListHeaderBaseX = m.List.HeaderBaseX
	}
	if m.List.RowTextOffsetX > 0 {
		ListRowTextOffsetX = m.List.RowTextOffsetX
	}
	if m.List.RowTextOffsetY > 0 {
		ListRowTextOffsetY = m.List.RowTextOffsetY
	}
	if m.List.RowBorderPad > 0 {
		ListRowBorderPad = m.List.RowBorderPad
	}
	if m.List.RowRightIconSize > 0 {
		ListRowRightIconSize = m.List.RowRightIconSize
	}
	if m.List.RowRightIconGap > 0 {
		ListRowRightIconGap = m.List.RowRightIconGap
	}
	if len(m.List.HeaderColumnsItems) > 0 {
		ListHeaderColumnsItems = append([]int(nil), m.List.HeaderColumnsItems...)
	}
	if len(m.List.HeaderColumnsWeapons) > 0 {
		ListHeaderColumnsWeapons = append([]int(nil), m.List.HeaderColumnsWeapons...)
	}
	if len(m.List.RowColumnsItems) > 0 {
		ListRowColumnsItems = append([]int(nil), m.List.RowColumnsItems...)
	}
	if len(m.List.RowColumnsWeapons) > 0 {
		ListRowColumnsWeapons = append([]int(nil), m.List.RowColumnsWeapons...)
	}
	if m.Line.Main > 0 {
		LineHMain = m.Line.Main
	}
	if m.Line.Small > 0 {
		LineHSmall = m.Line.Small
	}
	// Status
	if m.Status.PanelPad > 0 {
		StatusPanelPad = m.Status.PanelPad
	}
	if m.Status.PortraitSize > 0 {
		StatusPortraitSize = m.Status.PortraitSize
	}
	if m.Status.TextGapX > 0 {
		StatusTextGapX = m.Status.TextGapX
	}
	if m.Status.NameOffsetY > 0 {
		StatusNameOffsetY = m.Status.NameOffsetY
	}
	if m.Status.ClassGapFromName > 0 {
		StatusClassGapFromName = m.Status.ClassGapFromName
	}
	if m.Status.LevelGapFromName > 0 {
		StatusLevelGapFromName = m.Status.LevelGapFromName
	}
	if m.Status.HPGapFromName > 0 {
		StatusHPGapFromName = m.Status.HPGapFromName
	}
	if m.Status.HPBarGapFromName > 0 {
		StatusHPBarGapFromName = m.Status.HPBarGapFromName
	}
	if m.Status.HPBarW > 0 {
		StatusHPBarW = m.Status.HPBarW
	}
	if m.Status.HPBarH > 0 {
		StatusHPBarH = m.Status.HPBarH
	}
	if m.Status.StatsTopGap > 0 {
		StatusStatsTopGap = m.Status.StatsTopGap
	}
	if m.Status.StatsLineH > 0 {
		StatusStatsLineH = m.Status.StatsLineH
	}
	if m.Status.StatsColGap > 0 {
		StatusStatsColGap = m.Status.StatsColGap
	}
	if m.Status.WeaponRanksXExtra > 0 {
		StatusWeaponRanksXExtra = m.Status.WeaponRanksXExtra
	}
	if m.Status.RankLineH > 0 {
		StatusRankLineH = m.Status.RankLineH
	}
	if m.Status.MagicRanksTopExtra > 0 {
		StatusMagicRanksTopExtra = m.Status.MagicRanksTopExtra
	}
	if m.Status.EquipTitleGapY > 0 {
		StatusEquipTitleGapY = m.Status.EquipTitleGapY
	}
	if m.Status.EquipLineH > 0 {
		StatusEquipLineH = m.Status.EquipLineH
	}
	if m.Status.EquipRectYOffset > 0 {
		StatusEquipRectYOffset = m.Status.EquipRectYOffset
	}
	if m.Status.EquipRectW > 0 {
		StatusEquipRectW = m.Status.EquipRectW
	}
	if m.Status.EquipRectH > 0 {
		StatusEquipRectH = m.Status.EquipRectH
	}
	if m.Status.EquipLabelGapX > 0 {
		StatusEquipLabelGapX = m.Status.EquipLabelGapX
	}
	if m.Status.EquipUsesX > 0 {
		StatusEquipUsesX = m.Status.EquipUsesX
	}
	// Sim
	if m.Sim.StartBtnW > 0 {
		SimStartBtnW = m.Sim.StartBtnW
	}
	if m.Sim.StartBtnH > 0 {
		SimStartBtnH = m.Sim.StartBtnH
	}
	if m.Sim.AutoRunGap > 0 {
		SimAutoRunGap = m.Sim.AutoRunGap
	}
	if m.Sim.TitleYOffset > 0 {
		SimTitleYOffset = m.Sim.TitleYOffset
	}
	if m.Sim.TitleXOffsetFromCenter > 0 {
		SimTitleXOffsetFromCenter = m.Sim.TitleXOffsetFromCenter
	}
	if m.Sim.Terrain.ButtonW > 0 {
		TerrainBtnW = m.Sim.Terrain.ButtonW
	}
	if m.Sim.Terrain.ButtonH > 0 {
		TerrainBtnH = m.Sim.Terrain.ButtonH
	}
	if m.Sim.Terrain.BaseYFromBottom > 0 {
		TerrainBaseYFromBottom = m.Sim.Terrain.BaseYFromBottom
	}
	if m.Sim.Terrain.LeftBaseXOffset > 0 {
		TerrainLeftBaseXOffset = m.Sim.Terrain.LeftBaseXOffset
	}
	if m.Sim.Terrain.RightBaseXInset > 0 {
		TerrainRightBaseXInset = m.Sim.Terrain.RightBaseXInset
	}
	if m.Sim.Terrain.ButtonGap > 0 {
		TerrainBtnGap = m.Sim.Terrain.ButtonGap
	}
	if m.Sim.Terrain.LabelLeftXOffset > 0 {
		TerrainLabelLeftXOffset = m.Sim.Terrain.LabelLeftXOffset
	}
	if m.Sim.Terrain.LabelYOffsetFromBottom > 0 {
		TerrainLabelYOffsetFromBottom = m.Sim.Terrain.LabelYOffsetFromBottom
	}
	// Sim Preview
	if m.Sim.Preview.LeftXPad > 0 {
		SimPreviewLeftXPad = m.Sim.Preview.LeftXPad
	}
	if m.Sim.Preview.RightXInset > 0 {
		SimPreviewRightXInset = m.Sim.Preview.RightXInset
	}
	if m.Sim.Preview.TopYFromMargin > 0 {
		SimPreviewTopYFromMargin = m.Sim.Preview.TopYFromMargin
	}
	if m.Sim.Preview.CardW > 0 {
		SimPreviewCardW = m.Sim.Preview.CardW
	}
	if m.Sim.Preview.CardH > 0 {
		SimPreviewCardH = m.Sim.Preview.CardH
	}
	if m.Sim.Preview.CardInnerPad > 0 {
		SimPreviewCardInnerPad = m.Sim.Preview.CardInnerPad
	}
	if m.Sim.Preview.PortraitSize > 0 {
		SimPreviewPortraitSize = m.Sim.Preview.PortraitSize
	}
	if m.Sim.Preview.ClassOffsetY > 0 {
		SimPreviewClassOffsetY = m.Sim.Preview.ClassOffsetY
	}
	if m.Sim.Preview.HPLabelX > 0 {
		SimPreviewHPLabelX = m.Sim.Preview.HPLabelX
	}
	if m.Sim.Preview.HPLabelY > 0 {
		SimPreviewHPLabelY = m.Sim.Preview.HPLabelY
	}
	if m.Sim.Preview.HPBarX > 0 {
		SimPreviewHPBarX = m.Sim.Preview.HPBarX
	}
	if m.Sim.Preview.HPBarY > 0 {
		SimPreviewHPBarY = m.Sim.Preview.HPBarY
	}
	if m.Sim.Preview.HPBarW > 0 {
		SimPreviewHPBarW = m.Sim.Preview.HPBarW
	}
	if m.Sim.Preview.HPBarH > 0 {
		SimPreviewHPBarH = m.Sim.Preview.HPBarH
	}
	if m.Sim.Preview.NameOffsetX > 0 {
		SimPreviewNameOffsetX = m.Sim.Preview.NameOffsetX
	}
	if m.Sim.Preview.NameOffsetY > 0 {
		SimPreviewNameOffsetY = m.Sim.Preview.NameOffsetY
	}
	if m.Sim.Preview.LineY > 0 {
		SimPreviewLineY = m.Sim.Preview.LineY
	}
	if m.Sim.Preview.BaseMin > 0 {
		SimPreviewBaseMin = m.Sim.Preview.BaseMin
	}
	if m.Sim.Preview.BaseMax > 0 {
		SimPreviewBaseMax = m.Sim.Preview.BaseMax
	}
	if m.Sim.Preview.WrapPad > 0 {
		SimPreviewWrapPad = m.Sim.Preview.WrapPad
	}
	if m.Sim.Preview.LogPadX > 0 {
		SimPreviewLogPadX = m.Sim.Preview.LogPadX
	}
	if m.Sim.Preview.LogPadY > 0 {
		SimPreviewLogPadY = m.Sim.Preview.LogPadY
	}
	// Popup
	if m.Popup.Cols4ThresholdW > 0 {
		PopupCols4ThresholdW = m.Popup.Cols4ThresholdW
	}
	if m.Popup.Cols3ThresholdW > 0 {
		PopupCols3ThresholdW = m.Popup.Cols3ThresholdW
	}
	if m.Popup.GridInnerXTotalPad > 0 {
		PopupGridInnerXTotalPad = m.Popup.GridInnerXTotalPad
	}
	if m.Popup.CellH > 0 {
		PopupCellH = m.Popup.CellH
	}
	if m.Popup.CellGap > 0 {
		PopupCellGap = m.Popup.CellGap
	}
	if m.Popup.GridXPad > 0 {
		PopupGridXPad = m.Popup.GridXPad
	}
	if m.Popup.GridYOff > 0 {
		PopupGridYOff = m.Popup.GridYOff
	}
	if m.Popup.MaxW > 0 {
		PopupMaxW = m.Popup.MaxW
	}
	if m.Popup.MaxH > 0 {
		PopupMaxH = m.Popup.MaxH
	}
	// Widgets
	if m.Widgets.Back.PanelRightInset > 0 {
		WidgetsBackPanelRightInset = m.Widgets.Back.PanelRightInset
	}
	if m.Widgets.Back.TopPad > 0 {
		WidgetsBackTopPad = m.Widgets.Back.TopPad
	}
	if m.Widgets.Back.W > 0 {
		WidgetsBackW = m.Widgets.Back.W
	}
	if m.Widgets.Back.H > 0 {
		WidgetsBackH = m.Widgets.Back.H
	}
	if m.Widgets.Back.LabelX > 0 {
		WidgetsBackLabelX = m.Widgets.Back.LabelX
	}
	if m.Widgets.Back.LabelY > 0 {
		WidgetsBackLabelY = m.Widgets.Back.LabelY
	}
	if m.Widgets.LevelUp.W > 0 {
		WidgetsLevelUpW = m.Widgets.LevelUp.W
	}
	if m.Widgets.LevelUp.H > 0 {
		WidgetsLevelUpH = m.Widgets.LevelUp.H
	}
	if m.Widgets.LevelUp.LabelX > 0 {
		WidgetsLevelUpLabelX = m.Widgets.LevelUp.LabelX
	}
	if m.Widgets.LevelUp.LabelY > 0 {
		WidgetsLevelUpLabelY = m.Widgets.LevelUp.LabelY
	}
	if m.Widgets.ToBattle.GapFromRightBtn > 0 {
		WidgetsToBattleGapFromRightBtn = m.Widgets.ToBattle.GapFromRightBtn
	}
	if m.Widgets.ToBattle.W > 0 {
		WidgetsToBattleW = m.Widgets.ToBattle.W
	}
	if m.Widgets.ToBattle.LabelX > 0 {
		WidgetsToBattleLabelX = m.Widgets.ToBattle.LabelX
	}
	if m.Widgets.ToBattle.LabelY > 0 {
		WidgetsToBattleLabelY = m.Widgets.ToBattle.LabelY
	}
	if m.Widgets.SimBattle.W > 0 {
		WidgetsSimBattleW = m.Widgets.SimBattle.W
	}
	if m.Widgets.SimBattle.H > 0 {
		WidgetsSimBattleH = m.Widgets.SimBattle.H
	}
	if m.Widgets.SimBattle.TopPad > 0 {
		WidgetsSimBattleTopPad = m.Widgets.SimBattle.TopPad
	}
	if m.Widgets.SimBattle.LabelX > 0 {
		WidgetsSimBattleLabelX = m.Widgets.SimBattle.LabelX
	}
	if m.Widgets.SimBattle.LabelY > 0 {
		WidgetsSimBattleLabelY = m.Widgets.SimBattle.LabelY
	}
}
