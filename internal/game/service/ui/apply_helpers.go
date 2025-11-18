//go:build !headless

package uicore

// metricsTargets はレイアウト既定値への参照セットを保持し、テスト差し替えを容易にします。
type metricsTargets struct {
	List    listTargets
	Line    lineTargets
	Status  statusTargets
	Sim     simTargets
	Popup   popupTargets
	Widgets widgetsTargets
}

type listTargets struct {
	Margin               *int
	ItemH                *int
	ItemGap              *int
	PortraitSize         *int
	TitleOffset          *int
	HeaderTopGap         *int
	ItemsTopGap          *int
	PanelInnerPaddingX   *int
	TitleXOffset         *int
	HeaderBaseX          *int
	RowTextOffsetX       *int
	RowTextOffsetY       *int
	RowBorderPad         *int
	RowRightIconSize     *int
	RowRightIconGap      *int
	HeaderColumnsItems   *[]int
	HeaderColumnsWeapons *[]int
	RowColumnsItems      *[]int
	RowColumnsWeapons    *[]int
}

type lineTargets struct {
	Main  *int
	Small *int
}

type statusTargets struct {
	PanelPad           *int
	PortraitSize       *int
	TextGapX           *int
	NameOffsetY        *int
	ClassGapFromName   *int
	LevelGapFromName   *int
	HPGapFromName      *int
	HPBarGapFromName   *int
	HPBarW             *int
	HPBarH             *int
	StatsTopGap        *int
	StatsLineH         *int
	StatsColGap        *int
	WeaponRanksXExtra  *int
	RankLineH          *int
	MagicRanksTopExtra *int
	EquipTitleGapY     *int
	EquipLineH         *int
	EquipRectYOffset   *int
	EquipRectW         *int
	EquipRectH         *int
	EquipLabelGapX     *int
	EquipUsesX         *int
}

type simTargets struct {
	StartBtnW              *int
	StartBtnH              *int
	AutoRunGap             *int
	TitleYOffset           *int
	TitleXOffsetFromCenter *int
	Terrain                terrainTargets
	Preview                previewTargets
}

type terrainTargets struct {
	ButtonW                *int
	ButtonH                *int
	BaseYFromBottom        *int
	LeftBaseXOffset        *int
	RightBaseXInset        *int
	ButtonGap              *int
	LabelLeftXOffset       *int
	LabelYOffsetFromBottom *int
}

type previewTargets struct {
	LeftXPad       *int
	RightXInset    *int
	TopYFromMargin *int
	CardW          *int
	CardH          *int
	CardInnerPad   *int
	PortraitSize   *int
	ClassOffsetY   *int
	HPLabelX       *int
	HPLabelY       *int
	HPBarX         *int
	HPBarY         *int
	HPBarW         *int
	HPBarH         *int
	NameOffsetX    *int
	NameOffsetY    *int
	LineY          *int
	BaseMin        *int
	BaseMax        *int
	WrapPad        *int
	LogPadX        *int
	LogPadY        *int
}

type popupTargets struct {
	Cols4ThresholdW    *int
	Cols3ThresholdW    *int
	GridInnerXTotalPad *int
	CellH              *int
	CellGap            *int
	GridXPad           *int
	GridYOff           *int
	MaxW               *int
	MaxH               *int
}

type widgetsTargets struct {
	Back      widgetBackTargets
	LevelUp   widgetLevelUpTargets
	ToBattle  widgetToBattleTargets
	SimBattle widgetSimBattleTargets
}

type widgetBackTargets struct {
	PanelRightInset *int
	TopPad          *int
	W               *int
	H               *int
	LabelX          *int
	LabelY          *int
}

type widgetLevelUpTargets struct {
	W      *int
	H      *int
	LabelX *int
	LabelY *int
}

type widgetToBattleTargets struct {
	GapFromRightBtn *int
	LabelX          *int
	LabelY          *int
	W               *int
}

type widgetSimBattleTargets struct {
	W      *int
	H      *int
	TopPad *int
	LabelX *int
	LabelY *int
}

func newMetricsTargets() *metricsTargets {
	return &metricsTargets{
		List: listTargets{
			Margin:               &ListMargin,
			ItemH:                &ListItemH,
			ItemGap:              &ListItemGap,
			PortraitSize:         &ListPortraitSz,
			TitleOffset:          &ListTitleOffset,
			HeaderTopGap:         &ListHeaderTopGap,
			ItemsTopGap:          &ListItemsTopGap,
			PanelInnerPaddingX:   &ListPanelInnerPaddingX,
			TitleXOffset:         &ListTitleXOffset,
			HeaderBaseX:          &ListHeaderBaseX,
			RowTextOffsetX:       &ListRowTextOffsetX,
			RowTextOffsetY:       &ListRowTextOffsetY,
			RowBorderPad:         &ListRowBorderPad,
			RowRightIconSize:     &ListRowRightIconSize,
			RowRightIconGap:      &ListRowRightIconGap,
			HeaderColumnsItems:   &ListHeaderColumnsItems,
			HeaderColumnsWeapons: &ListHeaderColumnsWeapons,
			RowColumnsItems:      &ListRowColumnsItems,
			RowColumnsWeapons:    &ListRowColumnsWeapons,
		},
		Line: lineTargets{
			Main:  &LineHMain,
			Small: &LineHSmall,
		},
		Status: statusTargets{
			PanelPad:           &StatusPanelPad,
			PortraitSize:       &StatusPortraitSize,
			TextGapX:           &StatusTextGapX,
			NameOffsetY:        &StatusNameOffsetY,
			ClassGapFromName:   &StatusClassGapFromName,
			LevelGapFromName:   &StatusLevelGapFromName,
			HPGapFromName:      &StatusHPGapFromName,
			HPBarGapFromName:   &StatusHPBarGapFromName,
			HPBarW:             &StatusHPBarW,
			HPBarH:             &StatusHPBarH,
			StatsTopGap:        &StatusStatsTopGap,
			StatsLineH:         &StatusStatsLineH,
			StatsColGap:        &StatusStatsColGap,
			WeaponRanksXExtra:  &StatusWeaponRanksXExtra,
			RankLineH:          &StatusRankLineH,
			MagicRanksTopExtra: &StatusMagicRanksTopExtra,
			EquipTitleGapY:     &StatusEquipTitleGapY,
			EquipLineH:         &StatusEquipLineH,
			EquipRectYOffset:   &StatusEquipRectYOffset,
			EquipRectW:         &StatusEquipRectW,
			EquipRectH:         &StatusEquipRectH,
			EquipLabelGapX:     &StatusEquipLabelGapX,
			EquipUsesX:         &StatusEquipUsesX,
		},
		Sim: simTargets{
			StartBtnW:              &SimStartBtnW,
			StartBtnH:              &SimStartBtnH,
			AutoRunGap:             &SimAutoRunGap,
			TitleYOffset:           &SimTitleYOffset,
			TitleXOffsetFromCenter: &SimTitleXOffsetFromCenter,
			Terrain: terrainTargets{
				ButtonW:                &TerrainBtnW,
				ButtonH:                &TerrainBtnH,
				BaseYFromBottom:        &TerrainBaseYFromBottom,
				LeftBaseXOffset:        &TerrainLeftBaseXOffset,
				RightBaseXInset:        &TerrainRightBaseXInset,
				ButtonGap:              &TerrainBtnGap,
				LabelLeftXOffset:       &TerrainLabelLeftXOffset,
				LabelYOffsetFromBottom: &TerrainLabelYOffsetFromBottom,
			},
			Preview: previewTargets{
				LeftXPad:       &SimPreviewLeftXPad,
				RightXInset:    &SimPreviewRightXInset,
				TopYFromMargin: &SimPreviewTopYFromMargin,
				CardW:          &SimPreviewCardW,
				CardH:          &SimPreviewCardH,
				CardInnerPad:   &SimPreviewCardInnerPad,
				PortraitSize:   &SimPreviewPortraitSize,
				ClassOffsetY:   &SimPreviewClassOffsetY,
				HPLabelX:       &SimPreviewHPLabelX,
				HPLabelY:       &SimPreviewHPLabelY,
				HPBarX:         &SimPreviewHPBarX,
				HPBarY:         &SimPreviewHPBarY,
				HPBarW:         &SimPreviewHPBarW,
				HPBarH:         &SimPreviewHPBarH,
				NameOffsetX:    &SimPreviewNameOffsetX,
				NameOffsetY:    &SimPreviewNameOffsetY,
				LineY:          &SimPreviewLineY,
				BaseMin:        &SimPreviewBaseMin,
				BaseMax:        &SimPreviewBaseMax,
				WrapPad:        &SimPreviewWrapPad,
				LogPadX:        &SimPreviewLogPadX,
				LogPadY:        &SimPreviewLogPadY,
			},
		},
		Popup: popupTargets{
			Cols4ThresholdW:    &PopupCols4ThresholdW,
			Cols3ThresholdW:    &PopupCols3ThresholdW,
			GridInnerXTotalPad: &PopupGridInnerXTotalPad,
			CellH:              &PopupCellH,
			CellGap:            &PopupCellGap,
			GridXPad:           &PopupGridXPad,
			GridYOff:           &PopupGridYOff,
			MaxW:               &PopupMaxW,
			MaxH:               &PopupMaxH,
		},
		Widgets: widgetsTargets{
			Back: widgetBackTargets{
				PanelRightInset: &WidgetsBackPanelRightInset,
				TopPad:          &WidgetsBackTopPad,
				W:               &WidgetsBackW,
				H:               &WidgetsBackH,
				LabelX:          &WidgetsBackLabelX,
				LabelY:          &WidgetsBackLabelY,
			},
			LevelUp: widgetLevelUpTargets{
				W:      &WidgetsLevelUpW,
				H:      &WidgetsLevelUpH,
				LabelX: &WidgetsLevelUpLabelX,
				LabelY: &WidgetsLevelUpLabelY,
			},
			ToBattle: widgetToBattleTargets{
				GapFromRightBtn: &WidgetsToBattleGapFromRightBtn,
				LabelX:          &WidgetsToBattleLabelX,
				LabelY:          &WidgetsToBattleLabelY,
				W:               &WidgetsToBattleW,
			},
			SimBattle: widgetSimBattleTargets{
				W:      &WidgetsSimBattleW,
				H:      &WidgetsSimBattleH,
				TopPad: &WidgetsSimBattleTopPad,
				LabelX: &WidgetsSimBattleLabelX,
				LabelY: &WidgetsSimBattleLabelY,
			},
		},
	}
}

func assignPositive(dst *int, val int) {
	if dst == nil || val <= 0 {
		return
	}
	*dst = val
}

func assignSlice(dst *[]int, src []int) {
	if dst == nil || len(src) == 0 {
		return
	}
	*dst = copyInts(src)
}

func copyInts(src []int) []int {
	if src == nil {
		return nil
	}
	out := make([]int, len(src))
	copy(out, src)
	return out
}
