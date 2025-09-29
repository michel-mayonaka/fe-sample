package uicore

import "github.com/hajimehoshi/ebiten/v2"

var baseW, baseH int = 1920, 1080
var scale float32 = 1.0

// SetBaseResolution は論理解像度（Layoutの返すサイズ）を設定します。
func SetBaseResolution(w, h int) {
	if w > 0 && h > 0 {
		baseW, baseH = w, h
	}
}

// UpdateMetricsFromWindow は現在のウィンドウサイズからスケールを算出します。
func UpdateMetricsFromWindow() {
	w, h := ebiten.WindowSize()
	if w <= 0 || h <= 0 || baseW <= 0 || baseH <= 0 {
		scale = 1
		return
	}
	sx := float32(w) / float32(baseW)
	sy := float32(h) / float32(baseH)
	if sx < sy {
		scale = sx
	} else {
		scale = sy
	}
	if scale <= 0 {
		scale = 1
	}
}

// S は整数ピクセル値をスケールに応じて拡縮します。
func S(n int) int { return int(float32(n) * scale) }

// ListMarginPx は一覧パネルのマージン（スケール適用後）を返します。
func ListMarginPx() int { return S(ListMargin) }

// ListItemHPx は一覧行の高さ（スケール適用後）を返します。
func ListItemHPx() int { return S(ListItemH) }

// ListItemGapPx は一覧行の間隔（スケール適用後）を返します。
func ListItemGapPx() int { return S(ListItemGap) }

// ListPortraitSzPx はポートレート枠サイズ（スケール適用後）を返します。
func ListPortraitSzPx() int { return S(ListPortraitSz) }

// ListTitleOffsetPx はタイトルのYオフセット（スケール適用後）を返します。
func ListTitleOffsetPx() int { return S(ListTitleOffset) }

// LineHMainPx は本文行の高さ（スケール適用後）を返します。
func LineHMainPx() int { return S(LineHMain) }

// LineHSmallPx は小サイズ行の高さ（スケール適用後）を返します。
func LineHSmallPx() int { return S(LineHSmall) }

// CurrentScale は現在のスケール値を返します。
func CurrentScale() float32 { return scale }

// 追加メトリクス（スケール適用後）
func ListHeaderTopGapPx() int       { return S(ListHeaderTopGap) }
func ListItemsTopGapPx() int        { return S(ListItemsTopGap) }
func ListPanelInnerPaddingXPx() int { return S(ListPanelInnerPaddingX) }
func ListTitleXOffsetPx() int       { return S(ListTitleXOffset) }
func ListHeaderBaseXPx() int        { return S(ListHeaderBaseX) }
func ListRowTextOffsetXPx() int     { return S(ListRowTextOffsetX) }
func ListRowTextOffsetYPx() int     { return S(ListRowTextOffsetY) }
func ListRowBorderPadPx() int       { return S(ListRowBorderPad) }
func ListRowRightIconSizePx() int   { return S(ListRowRightIconSize) }
func ListRowRightIconGapPx() int    { return S(ListRowRightIconGap) }

func scaleSlice(xs []int) []int {
	if xs == nil {
		return nil
	}
	out := make([]int, len(xs))
	for i, v := range xs {
		out[i] = S(v)
	}
	return out
}

func ListHeaderColumnsItemsPx() []int   { return scaleSlice(ListHeaderColumnsItems) }
func ListHeaderColumnsWeaponsPx() []int { return scaleSlice(ListHeaderColumnsWeapons) }
func ListRowColumnsItemsPx() []int      { return scaleSlice(ListRowColumnsItems) }
func ListRowColumnsWeaponsPx() []int    { return scaleSlice(ListRowColumnsWeapons) }

// Status (scaled)
func StatusPanelPadPx() int           { return S(StatusPanelPad) }
func StatusPortraitSizePx() int       { return S(StatusPortraitSize) }
func StatusTextGapXPx() int           { return S(StatusTextGapX) }
func StatusNameOffsetYPx() int        { return S(StatusNameOffsetY) }
func StatusClassGapFromNamePx() int   { return S(StatusClassGapFromName) }
func StatusLevelGapFromNamePx() int   { return S(StatusLevelGapFromName) }
func StatusHPGapFromNamePx() int      { return S(StatusHPGapFromName) }
func StatusHPBarGapFromNamePx() int   { return S(StatusHPBarGapFromName) }
func StatusHPBarWPx() int             { return S(StatusHPBarW) }
func StatusHPBarHPx() int             { return S(StatusHPBarH) }
func StatusStatsTopGapPx() int        { return S(StatusStatsTopGap) }
func StatusStatsLineHPx() int         { return S(StatusStatsLineH) }
func StatusStatsColGapPx() int        { return S(StatusStatsColGap) }
func StatusWeaponRanksXExtraPx() int  { return S(StatusWeaponRanksXExtra) }
func StatusRankLineHPx() int          { return S(StatusRankLineH) }
func StatusMagicRanksTopExtraPx() int { return S(StatusMagicRanksTopExtra) }
func StatusEquipTitleGapYPx() int     { return S(StatusEquipTitleGapY) }
func StatusEquipLineHPx() int         { return S(StatusEquipLineH) }
func StatusEquipRectYOffsetPx() int   { return S(StatusEquipRectYOffset) }
func StatusEquipRectWPx() int         { return S(StatusEquipRectW) }
func StatusEquipRectHPx() int         { return S(StatusEquipRectH) }
func StatusEquipLabelGapXPx() int     { return S(StatusEquipLabelGapX) }
func StatusEquipUsesXPx() int         { return S(StatusEquipUsesX) }

// Sim (scaled)
func SimStartBtnWPx() int              { return S(SimStartBtnW) }
func SimStartBtnHPx() int              { return S(SimStartBtnH) }
func SimAutoRunGapPx() int             { return S(SimAutoRunGap) }
func SimTitleYOffsetPx() int           { return S(SimTitleYOffset) }
func SimTitleXOffsetFromCenterPx() int { return S(SimTitleXOffsetFromCenter) }

// Terrain (scaled)
func TerrainBtnWPx() int                   { return S(TerrainBtnW) }
func TerrainBtnHPx() int                   { return S(TerrainBtnH) }
func TerrainBaseYFromBottomPx() int        { return S(TerrainBaseYFromBottom) }
func TerrainLeftBaseXOffsetPx() int        { return S(TerrainLeftBaseXOffset) }
func TerrainRightBaseXInsetPx() int        { return S(TerrainRightBaseXInset) }
func TerrainBtnGapPx() int                 { return S(TerrainBtnGap) }
func TerrainLabelLeftXOffsetPx() int       { return S(TerrainLabelLeftXOffset) }
func TerrainLabelYOffsetFromBottomPx() int { return S(TerrainLabelYOffsetFromBottom) }

// Sim preview (scaled)
func SimPreviewLeftXPadPx() int       { return S(SimPreviewLeftXPad) }
func SimPreviewRightXInsetPx() int    { return S(SimPreviewRightXInset) }
func SimPreviewTopYFromMarginPx() int { return S(SimPreviewTopYFromMargin) }
func SimPreviewCardWPx() int          { return S(SimPreviewCardW) }
func SimPreviewCardHPx() int          { return S(SimPreviewCardH) }
func SimPreviewCardInnerPadPx() int   { return S(SimPreviewCardInnerPad) }
func SimPreviewPortraitSizePx() int   { return S(SimPreviewPortraitSize) }
func SimPreviewClassOffsetYPx() int   { return S(SimPreviewClassOffsetY) }
func SimPreviewHPLabelXPx() int       { return S(SimPreviewHPLabelX) }
func SimPreviewHPLabelYPx() int       { return S(SimPreviewHPLabelY) }
func SimPreviewHPBarXPx() int         { return S(SimPreviewHPBarX) }
func SimPreviewHPBarYPx() int         { return S(SimPreviewHPBarY) }
func SimPreviewHPBarWPx() int         { return S(SimPreviewHPBarW) }
func SimPreviewHPBarHPx() int         { return S(SimPreviewHPBarH) }
func SimPreviewNameOffsetXPx() int    { return S(SimPreviewNameOffsetX) }
func SimPreviewNameOffsetYPx() int    { return S(SimPreviewNameOffsetY) }
func SimPreviewLineYPx() int          { return S(SimPreviewLineY) }
func SimPreviewBaseMinPx() int        { return S(SimPreviewBaseMin) }
func SimPreviewBaseMaxPx() int        { return S(SimPreviewBaseMax) }
func SimPreviewWrapPadPx() int        { return S(SimPreviewWrapPad) }
func SimPreviewLogPadXPx() int        { return S(SimPreviewLogPadX) }
func SimPreviewLogPadYPx() int        { return S(SimPreviewLogPadY) }

// Popup (scaled)
func PopupCols4ThresholdWPx() int    { return S(PopupCols4ThresholdW) }
func PopupCols3ThresholdWPx() int    { return S(PopupCols3ThresholdW) }
func PopupGridInnerXTotalPadPx() int { return S(PopupGridInnerXTotalPad) }
func PopupCellHPx() int              { return S(PopupCellH) }
func PopupCellGapPx() int            { return S(PopupCellGap) }
func PopupGridXPadPx() int           { return S(PopupGridXPad) }
func PopupGridYOffPx() int           { return S(PopupGridYOff) }
func PopupMaxWPx() int               { return S(PopupMaxW) }
func PopupMaxHPx() int               { return S(PopupMaxH) }

// Widgets (scaled)
func WidgetsBackPanelRightInsetPx() int { return S(WidgetsBackPanelRightInset) }
func WidgetsBackTopPadPx() int          { return S(WidgetsBackTopPad) }
func WidgetsBackWPx() int               { return S(WidgetsBackW) }
func WidgetsBackHPx() int               { return S(WidgetsBackH) }
func WidgetsBackLabelXPx() int          { return S(WidgetsBackLabelX) }
func WidgetsBackLabelYPx() int          { return S(WidgetsBackLabelY) }

func WidgetsLevelUpWPx() int      { return S(WidgetsLevelUpW) }
func WidgetsLevelUpHPx() int      { return S(WidgetsLevelUpH) }
func WidgetsLevelUpLabelXPx() int { return S(WidgetsLevelUpLabelX) }
func WidgetsLevelUpLabelYPx() int { return S(WidgetsLevelUpLabelY) }

func WidgetsToBattleGapFromRightBtnPx() int { return S(WidgetsToBattleGapFromRightBtn) }
func WidgetsToBattleWPx() int               { return S(WidgetsToBattleW) }
func WidgetsToBattleLabelXPx() int          { return S(WidgetsToBattleLabelX) }
func WidgetsToBattleLabelYPx() int          { return S(WidgetsToBattleLabelY) }

func WidgetsSimBattleWPx() int      { return S(WidgetsSimBattleW) }
func WidgetsSimBattleHPx() int      { return S(WidgetsSimBattleH) }
func WidgetsSimBattleTopPadPx() int { return S(WidgetsSimBattleTopPad) }
func WidgetsSimBattleLabelXPx() int { return S(WidgetsSimBattleLabelX) }
func WidgetsSimBattleLabelYPx() int { return S(WidgetsSimBattleLabelY) }
