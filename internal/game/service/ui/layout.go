//go:build !headless

package uicore

// レイアウト既定値（論理解像度に対する基準値）。
// 外部メトリクス適用により起動時に上書きされます。
var (
	// ListMargin は一覧パネルのマージン（px）。
	ListMargin = 24
	// ListItemH は一覧行の高さ（px）。
	ListItemH = 100
	// ListItemGap は一覧行の縦方向の間隔（px）。
	ListItemGap = 12
	// ListPortraitSz は一覧のポートレート枠サイズ（px）。
	ListPortraitSz = 80
	// ListTitleOffset はタイトルのYオフセット（px）。
	ListTitleOffset = 44
	// LineHMain は本文フォントの行間（px）。
	LineHMain = 26
	// LineHSmall は小サイズフォントの行間（px）。
	LineHSmall = 22
	// ListHeaderTopGap はタイトル直下のヘッダまでのY間隔（px）。
	ListHeaderTopGap = 8
	// ListItemsTopGap はタイトル直下からアイテム行開始までのY間隔（px）。
	ListItemsTopGap = 32
	// ListPanelInnerPaddingX はパネル内左右の余白（px）。
	ListPanelInnerPaddingX = 16
	// ListTitleXOffset はパネル左端からタイトル文字までのXオフセット（px）。
	ListTitleXOffset = 20
	// ListHeaderBaseX はパネル左端からヘッダ列0のXオフセット（px）。
	ListHeaderBaseX = 36
	// ListRowTextOffsetX は各行のテキスト開始Xオフセット（px）。
	ListRowTextOffsetX = 20
	// ListRowTextOffsetY は各行でのテキストのYオフセット（px）。
	ListRowTextOffsetY = 36
	// ListRowBorderPad は行の縁取り余白（px）。
	ListRowBorderPad = 2
	// ListRowRightIconSize は右端の所有者アイコンのサイズ（px）。
	ListRowRightIconSize = 24
	// ListRowRightIconGap は右端からアイコンまでの余白（px）。
	ListRowRightIconGap = 12
	// ListHeaderColumnsItems はヘッダ列（アイテム一覧）のXオフセット群（px, 先頭は0）。
	ListHeaderColumnsItems = []int{0, 560, 720, 900, 1000}
	// ListHeaderColumnsWeapons はヘッダ列（武器一覧）のXオフセット群（px, 先頭は0）。
	ListHeaderColumnsWeapons = []int{0, 560, 680, 760, 840, 920, 1000, 1080, 1160}
	// ListRowColumnsItems は行テキスト列（アイテム一覧）のXオフセット群（px, 先頭は0）。
	ListRowColumnsItems = []int{0, 540, 700, 880, 980}
	// ListRowColumnsWeapons は行テキスト列（武器一覧）のXオフセット群（px, 先頭は0）。
	ListRowColumnsWeapons = []int{0, 540, 660, 750, 830, 910, 990, 1070, 1150}
	// --- Status screen ---
	StatusPanelPad           = 24
	StatusPortraitSize       = 320
	StatusTextGapX           = 32
	StatusNameOffsetY        = 44
	StatusClassGapFromName   = 40
	StatusLevelGapFromName   = 70
	StatusHPGapFromName      = 110
	StatusHPBarGapFromName   = 116
	StatusHPBarW             = 300
	StatusHPBarH             = 14
	StatusStatsTopGap        = 160
	StatusStatsLineH         = 34
	StatusStatsColGap        = 180
	StatusWeaponRanksXExtra  = 64
	StatusRankLineH          = 32
	StatusMagicRanksTopExtra = 176
	StatusEquipTitleGapY     = 56
	StatusEquipLineH         = 30
	StatusEquipRectYOffset   = 20
	StatusEquipRectW         = 360
	StatusEquipRectH         = 26
	StatusEquipLabelGapX     = 14
	StatusEquipUsesX         = 300
	// --- Sim screen ---
	SimStartBtnW              = 240
	SimStartBtnH              = 60
	SimAutoRunGap             = 20
	SimTitleYOffset           = 56
	SimTitleXOffsetFromCenter = 120
	// --- Terrain widgets (Sim) ---
	TerrainBtnW                   = 120
	TerrainBtnH                   = 40
	TerrainBaseYFromBottom        = 200
	TerrainLeftBaseXOffset        = 40
	TerrainRightBaseXInset        = 560
	TerrainBtnGap                 = 16
	TerrainLabelLeftXOffset       = 40
	TerrainLabelYOffsetFromBottom = 230

	// --- Sim preview panel ---
	SimPreviewLeftXPad       = 40
	SimPreviewRightXInset    = 560
	SimPreviewTopYFromMargin = 80
	SimPreviewCardW          = 520
	SimPreviewCardH          = 420
	SimPreviewCardInnerPad   = 16
	SimPreviewPortraitSize   = 120
	SimPreviewClassOffsetY   = 48
	SimPreviewHPLabelX       = 160
	SimPreviewHPLabelY       = 90
	SimPreviewHPBarX         = 200
	SimPreviewHPBarY         = 86
	SimPreviewHPBarW         = 280
	SimPreviewHPBarH         = 14
	SimPreviewNameOffsetX    = 16
	SimPreviewNameOffsetY    = 16
	SimPreviewLineY          = 410
	SimPreviewBaseMin        = 420
	SimPreviewBaseMax        = 900
	SimPreviewWrapPad        = 80
	SimPreviewLogPadX        = 16
	SimPreviewLogPadY        = 0

	// --- Popup (choose unit etc.) ---
	PopupCols4ThresholdW    = 900
	PopupCols3ThresholdW    = 700
	PopupGridInnerXTotalPad = 48
	PopupCellH              = 160
	PopupCellGap            = 8
	PopupGridXPad           = 24
	PopupGridYOff           = 88
	PopupMaxW               = 1480
	PopupMaxH               = 900

	// --- Widgets (buttons/tabs) ---
	WidgetsBackPanelRightInset = 180
	WidgetsBackTopPad          = 24
	WidgetsBackW               = 160
	WidgetsBackH               = 48
	WidgetsBackLabelX          = 20
	WidgetsBackLabelY          = 32

	WidgetsLevelUpW      = 220
	WidgetsLevelUpH      = 56
	WidgetsLevelUpLabelX = 24
	WidgetsLevelUpLabelY = 36

	WidgetsToBattleGapFromRightBtn = 20
	WidgetsToBattleW               = 220
	WidgetsToBattleLabelX          = 70
	WidgetsToBattleLabelY          = 36

	WidgetsSimBattleW      = 160
	WidgetsSimBattleH      = 48
	WidgetsSimBattleTopPad = 16
	WidgetsSimBattleLabelX = 24
	WidgetsSimBattleLabelY = 32
)
