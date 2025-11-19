package uimetrics

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// Metrics は外部ファイルで定義する UI メトリクスです。
// JSON キーは lowerCamel とし、docs/KNOWLEDGE/engineering/naming.md に整合します。
type Metrics struct {
	Base struct {
		W int `json:"w"`
		H int `json:"h"`
	} `json:"base"`
	List struct {
		Margin               int   `json:"margin"`
		ItemH                int   `json:"itemH"`
		ItemGap              int   `json:"itemGap"`
		PortraitSize         int   `json:"portraitSize"`
		TitleOffset          int   `json:"titleOffset"`
		HeaderTopGap         int   `json:"headerTopGap"`
		ItemsTopGap          int   `json:"itemsTopGap"`
		PanelInnerPaddingX   int   `json:"panelInnerPaddingX"`
		TitleXOffset         int   `json:"titleXOffset"`
		HeaderBaseX          int   `json:"headerBaseX"`
		RowTextOffsetX       int   `json:"rowTextOffsetX"`
		RowTextOffsetY       int   `json:"rowTextOffsetY"`
		RowBorderPad         int   `json:"rowBorderPad"`
		RowRightIconSize     int   `json:"rowRightIconSize"`
		RowRightIconGap      int   `json:"rowRightIconGap"`
		HeaderColumnsWeapons []int `json:"headerColumnsWeapons"`
		HeaderColumnsItems   []int `json:"headerColumnsItems"`
		RowColumnsWeapons    []int `json:"rowColumnsWeapons"`
		RowColumnsItems      []int `json:"rowColumnsItems"`
	} `json:"list"`
	Line struct {
		Main  int `json:"main"`
		Small int `json:"small"`
	} `json:"line"`
	Status struct {
		PanelPad           int `json:"panelPad"`
		PortraitSize       int `json:"portraitSize"`
		TextGapX           int `json:"textGapX"`
		NameOffsetY        int `json:"nameOffsetY"`
		ClassGapFromName   int `json:"classGapFromName"`
		LevelGapFromName   int `json:"levelGapFromName"`
		HPGapFromName      int `json:"hpGapFromName"`
		HPBarGapFromName   int `json:"hpBarGapFromName"`
		HPBarW             int `json:"hpBarW"`
		HPBarH             int `json:"hpBarH"`
		StatsTopGap        int `json:"statsTopGap"`
		StatsLineH         int `json:"statsLineH"`
		StatsColGap        int `json:"statsColGap"`
		WeaponRanksXExtra  int `json:"weaponRanksXExtra"`
		RankLineH          int `json:"rankLineH"`
		MagicRanksTopExtra int `json:"magicRanksTopExtra"`
		EquipTitleGapY     int `json:"equipTitleGapY"`
		EquipLineH         int `json:"equipLineH"`
		EquipRectYOffset   int `json:"equipRectYOffset"`
		EquipRectW         int `json:"equipRectW"`
		EquipRectH         int `json:"equipRectH"`
		EquipLabelGapX     int `json:"equipLabelGapX"`
		EquipUsesX         int `json:"equipUsesX"`
	} `json:"status"`
	Sim struct {
		StartBtnW              int `json:"startBtnW"`
		StartBtnH              int `json:"startBtnH"`
		AutoRunGap             int `json:"autoRunGap"`
		TitleYOffset           int `json:"titleYOffset"`
		TitleXOffsetFromCenter int `json:"titleXOffsetFromCenter"`
		Terrain                struct {
			ButtonW                int `json:"buttonW"`
			ButtonH                int `json:"buttonH"`
			BaseYFromBottom        int `json:"baseYFromBottom"`
			LeftBaseXOffset        int `json:"leftBaseXOffset"`
			RightBaseXInset        int `json:"rightBaseXInset"`
			ButtonGap              int `json:"buttonGap"`
			LabelLeftXOffset       int `json:"labelLeftXOffset"`
			LabelYOffsetFromBottom int `json:"labelYOffsetFromBottom"`
		} `json:"terrain"`
		Preview struct {
			LeftXPad       int `json:"leftXPad"`
			RightXInset    int `json:"rightXInset"`
			TopYFromMargin int `json:"topYFromMargin"`
			CardW          int `json:"cardW"`
			CardH          int `json:"cardH"`
			CardInnerPad   int `json:"cardInnerPad"`
			PortraitSize   int `json:"portraitSize"`
			ClassOffsetY   int `json:"classOffsetY"`
			HPLabelX       int `json:"hpLabelX"`
			HPLabelY       int `json:"hpLabelY"`
			HPBarX         int `json:"hpBarX"`
			HPBarY         int `json:"hpBarY"`
			HPBarW         int `json:"hpBarW"`
			HPBarH         int `json:"hpBarH"`
			NameOffsetX    int `json:"nameOffsetX"`
			NameOffsetY    int `json:"nameOffsetY"`
			LineY          int `json:"lineY"`
			BaseMin        int `json:"baseMin"`
			BaseMax        int `json:"baseMax"`
			WrapPad        int `json:"wrapPad"`
			LogPadX        int `json:"logPadX"`
			LogPadY        int `json:"logPadY"`
		} `json:"preview"`
	} `json:"sim"`
	Popup struct {
		Cols4ThresholdW    int `json:"cols4ThresholdW"`
		Cols3ThresholdW    int `json:"cols3ThresholdW"`
		GridInnerXTotalPad int `json:"gridInnerXTotalPad"`
		CellH              int `json:"cellH"`
		CellGap            int `json:"cellGap"`
		GridXPad           int `json:"gridXPad"`
		GridYOff           int `json:"gridYOff"`
		MaxW               int `json:"maxW"`
		MaxH               int `json:"maxH"`
	} `json:"popup"`
	Widgets struct {
		Back struct {
			PanelRightInset int `json:"panelRightInset"`
			TopPad          int `json:"topPad"`
			W               int `json:"w"`
			H               int `json:"h"`
			LabelX          int `json:"labelX"`
			LabelY          int `json:"labelY"`
		} `json:"back"`
		LevelUp struct {
			W      int `json:"w"`
			H      int `json:"h"`
			LabelX int `json:"labelX"`
			LabelY int `json:"labelY"`
		} `json:"levelUp"`
		ToBattle struct {
			GapFromRightBtn int `json:"gapFromRightBtn"`
			LabelX          int `json:"labelX"`
			LabelY          int `json:"labelY"`
			W               int `json:"w"`
		} `json:"toBattle"`
		SimBattle struct {
			W      int `json:"w"`
			H      int `json:"h"`
			TopPad int `json:"topPad"`
			LabelX int `json:"labelX"`
			LabelY int `json:"labelY"`
		} `json:"simBattle"`
	} `json:"widgets"`
}

// Default は現行のビルトイン値を返します（ファイル不在時のフォールバック）。
func Default() Metrics {
	var m Metrics
	m.Base.W, m.Base.H = 1920, 1080
	m.List.Margin = 24
	m.List.ItemH = 100
	m.List.ItemGap = 12
	m.List.PortraitSize = 80
	m.List.TitleOffset = 44
	m.List.HeaderTopGap = 8
	m.List.ItemsTopGap = 32
	m.List.PanelInnerPaddingX = 16
	m.List.TitleXOffset = 20
	m.List.HeaderBaseX = 36
	m.List.RowTextOffsetX = 20
	m.List.RowTextOffsetY = 36
	m.List.RowBorderPad = 2
	m.List.RowRightIconSize = 24
	m.List.RowRightIconGap = 12
	m.List.HeaderColumnsItems = []int{0, 560, 720, 900, 1000}
	m.List.HeaderColumnsWeapons = []int{0, 560, 680, 760, 840, 920, 1000, 1080, 1160}
	m.List.RowColumnsItems = []int{0, 540, 700, 880, 980}
	m.List.RowColumnsWeapons = []int{0, 540, 660, 750, 830, 910, 990, 1070, 1150}
	m.Line.Main = 26
	m.Line.Small = 22
	// Status defaults
	m.Status.PanelPad = 24
	m.Status.PortraitSize = 320
	m.Status.TextGapX = 32
	m.Status.NameOffsetY = 44
	m.Status.ClassGapFromName = 40
	m.Status.LevelGapFromName = 70
	m.Status.HPGapFromName = 110
	m.Status.HPBarGapFromName = 116
	m.Status.HPBarW = 300
	m.Status.HPBarH = 14
	m.Status.StatsTopGap = 160
	m.Status.StatsLineH = 34
	m.Status.StatsColGap = 180
	m.Status.WeaponRanksXExtra = 64
	m.Status.RankLineH = 32
	m.Status.MagicRanksTopExtra = 5*32 + 16
	m.Status.EquipTitleGapY = 56
	m.Status.EquipLineH = 30
	m.Status.EquipRectYOffset = 20
	m.Status.EquipRectW = 360
	m.Status.EquipRectH = 26
	m.Status.EquipLabelGapX = 14
	m.Status.EquipUsesX = 300
	// Sim defaults
	m.Sim.StartBtnW = 240
	m.Sim.StartBtnH = 60
	m.Sim.AutoRunGap = 20
	m.Sim.TitleYOffset = 56
	m.Sim.TitleXOffsetFromCenter = 120
	m.Sim.Terrain.ButtonW = 120
	m.Sim.Terrain.ButtonH = 40
	m.Sim.Terrain.BaseYFromBottom = 200
	m.Sim.Terrain.LeftBaseXOffset = 40
	m.Sim.Terrain.RightBaseXInset = 560
	m.Sim.Terrain.ButtonGap = 16
	m.Sim.Terrain.LabelLeftXOffset = 40
	m.Sim.Terrain.LabelYOffsetFromBottom = 230
	m.Sim.Preview.LeftXPad = 40
	m.Sim.Preview.RightXInset = 560
	m.Sim.Preview.TopYFromMargin = 80
	m.Sim.Preview.CardW = 520
	m.Sim.Preview.CardH = 420
	m.Sim.Preview.CardInnerPad = 16
	m.Sim.Preview.PortraitSize = 120
	m.Sim.Preview.ClassOffsetY = 48
	m.Sim.Preview.HPLabelX = 160
	m.Sim.Preview.HPLabelY = 90
	m.Sim.Preview.HPBarX = 200
	m.Sim.Preview.HPBarY = 86
	m.Sim.Preview.HPBarW = 280
	m.Sim.Preview.HPBarH = 14
	m.Sim.Preview.NameOffsetX = 16
	m.Sim.Preview.NameOffsetY = 16
	m.Sim.Preview.LineY = 410
	m.Sim.Preview.BaseMin = 420
	m.Sim.Preview.BaseMax = 900
	m.Sim.Preview.WrapPad = 80
	m.Sim.Preview.LogPadX = 16
	m.Sim.Preview.LogPadY = 0
	// Popup defaults
	m.Popup.Cols4ThresholdW = 900
	m.Popup.Cols3ThresholdW = 700
	m.Popup.GridInnerXTotalPad = 48
	m.Popup.CellH = 160
	m.Popup.CellGap = 8
	m.Popup.GridXPad = 24
	m.Popup.GridYOff = 88
	m.Popup.MaxW = 1480
	m.Popup.MaxH = 900
	// Widgets defaults
	m.Widgets.Back.PanelRightInset = 180
	m.Widgets.Back.TopPad = 24
	m.Widgets.Back.W = 160
	m.Widgets.Back.H = 48
	m.Widgets.Back.LabelX = 20
	m.Widgets.Back.LabelY = 32
	m.Widgets.LevelUp.W = 220
	m.Widgets.LevelUp.H = 56
	m.Widgets.LevelUp.LabelX = 24
	m.Widgets.LevelUp.LabelY = 36
	m.Widgets.ToBattle.GapFromRightBtn = 20
	m.Widgets.ToBattle.W = 220
	m.Widgets.ToBattle.LabelX = 70
	m.Widgets.ToBattle.LabelY = 36
	m.Widgets.SimBattle.W = 160
	m.Widgets.SimBattle.H = 48
	m.Widgets.SimBattle.TopPad = 16
	m.Widgets.SimBattle.LabelX = 24
	m.Widgets.SimBattle.LabelY = 32
	return m
}

// Load は単一ファイルから JSON を読み込みます。
func Load(path string) (Metrics, error) {
	if path == "" {
		return Metrics{}, errors.New("empty path")
	}
	f, err := os.Open(path)
	if err != nil {
		return Metrics{}, err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return Metrics{}, err
	}
	var m Metrics
	if err := json.Unmarshal(b, &m); err != nil {
		return Metrics{}, err
	}
	return m, nil
}

// LoadOrDefault は user→master の順で読み込み、無ければデフォルトを返します。
func LoadOrDefault(userPath, masterPath string) Metrics {
	if m, err := Load(userPath); err == nil {
		return m
	}
	if m, err := Load(masterPath); err == nil {
		return m
	}
	return Default()
}
