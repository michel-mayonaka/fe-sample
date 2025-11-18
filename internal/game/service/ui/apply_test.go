package uicore

import (
	"reflect"
	"testing"
)

func TestApplyMetrics_PartialDoesNotOverride(t *testing.T) {
	def := DefaultMetrics()
	t.Cleanup(func() { ApplyMetrics(def) })

	var m Metrics
	m.List.Margin = def.List.Margin + 5
	ApplyMetrics(m)

	if ListMargin != def.List.Margin+5 {
		t.Fatalf("margin not applied: got=%d", ListMargin)
	}
	// unchanged fields must remain
	if ListItemH != def.List.ItemH || ListItemGap != def.List.ItemGap {
		t.Fatalf("unexpected override: h=%d gap=%d", ListItemH, ListItemGap)
	}
}

func TestApplyListTargets(t *testing.T) {
	var captured Metrics
	targets := captureListTargets(&captured)

	var src Metrics
	src.List.Margin = 10
	src.List.ItemH = 20
	src.List.ItemGap = 5
	src.List.PortraitSize = 64
	src.List.TitleOffset = 12
	src.List.HeaderTopGap = 3
	src.List.ItemsTopGap = 8
	src.List.PanelInnerPaddingX = 6
	src.List.TitleXOffset = 9
	src.List.HeaderBaseX = 7
	src.List.RowTextOffsetX = 11
	src.List.RowTextOffsetY = 13
	src.List.RowBorderPad = 2
	src.List.RowRightIconSize = 14
	src.List.RowRightIconGap = 4
	src.List.HeaderColumnsItems = []int{1, 2, 3}
	src.List.HeaderColumnsWeapons = []int{4, 5, 6, 7}
	src.List.RowColumnsItems = []int{8, 9}
	src.List.RowColumnsWeapons = []int{10, 11, 12}
	src.Line.Main = 24
	src.Line.Small = 16

	applyList(targets, src)

	if !reflect.DeepEqual(captured.List, src.List) {
		t.Fatalf("list mismatch: got=%+v want=%+v", captured.List, src.List)
	}
	if captured.Line != src.Line {
		t.Fatalf("line mismatch: got=%+v want=%+v", captured.Line, src.Line)
	}

	assertSliceCopied(t, "HeaderColumnsItems", captured.List.HeaderColumnsItems, src.List.HeaderColumnsItems)
	assertSliceCopied(t, "HeaderColumnsWeapons", captured.List.HeaderColumnsWeapons, src.List.HeaderColumnsWeapons)
	assertSliceCopied(t, "RowColumnsItems", captured.List.RowColumnsItems, src.List.RowColumnsItems)
	assertSliceCopied(t, "RowColumnsWeapons", captured.List.RowColumnsWeapons, src.List.RowColumnsWeapons)
}

func TestApplyStatusTargets(t *testing.T) {
	var captured Metrics
	targets := captureStatusTargets(&captured)

	var src Metrics
	src.Status.PanelPad = 10
	src.Status.PortraitSize = 40
	src.Status.TextGapX = 20
	src.Status.NameOffsetY = 30
	src.Status.ClassGapFromName = 12
	src.Status.LevelGapFromName = 14
	src.Status.HPGapFromName = 16
	src.Status.HPBarGapFromName = 18
	src.Status.HPBarW = 200
	src.Status.HPBarH = 14
	src.Status.StatsTopGap = 80
	src.Status.StatsLineH = 32
	src.Status.StatsColGap = 64
	src.Status.WeaponRanksXExtra = 42
	src.Status.RankLineH = 20
	src.Status.MagicRanksTopExtra = 70
	src.Status.EquipTitleGapY = 22
	src.Status.EquipLineH = 24
	src.Status.EquipRectYOffset = 8
	src.Status.EquipRectW = 180
	src.Status.EquipRectH = 20
	src.Status.EquipLabelGapX = 6
	src.Status.EquipUsesX = 150

	applyStatus(targets, src)

	if captured.Status != src.Status {
		t.Fatalf("status mismatch: got=%+v want=%+v", captured.Status, src.Status)
	}
}

func TestApplySimTargets(t *testing.T) {
	var captured Metrics
	targets := captureSimTargets(&captured)

	var src Metrics
	src.Sim.StartBtnW = 120
	src.Sim.StartBtnH = 50
	src.Sim.AutoRunGap = 16
	src.Sim.TitleYOffset = 40
	src.Sim.TitleXOffsetFromCenter = 80
	src.Sim.Terrain.ButtonW = 90
	src.Sim.Terrain.ButtonH = 34
	src.Sim.Terrain.BaseYFromBottom = 60
	src.Sim.Terrain.LeftBaseXOffset = 18
	src.Sim.Terrain.RightBaseXInset = 44
	src.Sim.Terrain.ButtonGap = 6
	src.Sim.Terrain.LabelLeftXOffset = 12
	src.Sim.Terrain.LabelYOffsetFromBottom = 22
	src.Sim.Preview.LeftXPad = 14
	src.Sim.Preview.RightXInset = 32
	src.Sim.Preview.TopYFromMargin = 28
	src.Sim.Preview.CardW = 420
	src.Sim.Preview.CardH = 360
	src.Sim.Preview.CardInnerPad = 10
	src.Sim.Preview.PortraitSize = 96
	src.Sim.Preview.ClassOffsetY = 18
	src.Sim.Preview.HPLabelX = 200
	src.Sim.Preview.HPLabelY = 210
	src.Sim.Preview.HPBarX = 220
	src.Sim.Preview.HPBarY = 230
	src.Sim.Preview.HPBarW = 300
	src.Sim.Preview.HPBarH = 12
	src.Sim.Preview.NameOffsetX = 24
	src.Sim.Preview.NameOffsetY = 26
	src.Sim.Preview.LineY = 280
	src.Sim.Preview.BaseMin = 90
	src.Sim.Preview.BaseMax = 140
	src.Sim.Preview.WrapPad = 12
	src.Sim.Preview.LogPadX = 8
	src.Sim.Preview.LogPadY = 4

	applySim(targets, src)

	if captured.Sim != src.Sim {
		t.Fatalf("sim mismatch: got=%+v want=%+v", captured.Sim, src.Sim)
	}
}

func TestApplyPopupTargets(t *testing.T) {
	var captured Metrics
	targets := capturePopupTargets(&captured)

	var src Metrics
	src.Popup.Cols4ThresholdW = 900
	src.Popup.Cols3ThresholdW = 700
	src.Popup.GridInnerXTotalPad = 60
	src.Popup.CellH = 150
	src.Popup.CellGap = 10
	src.Popup.GridXPad = 30
	src.Popup.GridYOff = 80
	src.Popup.MaxW = 1500
	src.Popup.MaxH = 980

	applyPopup(targets, src)

	if captured.Popup != src.Popup {
		t.Fatalf("popup mismatch: got=%+v want=%+v", captured.Popup, src.Popup)
	}
}

func TestApplyWidgetsTargets(t *testing.T) {
	var captured Metrics
	targets := captureWidgetsTargets(&captured)

	var src Metrics
	src.Widgets.Back.PanelRightInset = 100
	src.Widgets.Back.TopPad = 16
	src.Widgets.Back.W = 140
	src.Widgets.Back.H = 42
	src.Widgets.Back.LabelX = 8
	src.Widgets.Back.LabelY = 18
	src.Widgets.LevelUp.W = 200
	src.Widgets.LevelUp.H = 60
	src.Widgets.LevelUp.LabelX = 20
	src.Widgets.LevelUp.LabelY = 30
	src.Widgets.ToBattle.GapFromRightBtn = 14
	src.Widgets.ToBattle.LabelX = 24
	src.Widgets.ToBattle.LabelY = 28
	src.Widgets.ToBattle.W = 220
	src.Widgets.SimBattle.W = 180
	src.Widgets.SimBattle.H = 50
	src.Widgets.SimBattle.TopPad = 12
	src.Widgets.SimBattle.LabelX = 10
	src.Widgets.SimBattle.LabelY = 26

	applyWidgets(targets, src)

	if captured.Widgets != src.Widgets {
		t.Fatalf("widgets mismatch: got=%+v want=%+v", captured.Widgets, src.Widgets)
	}
}

func captureListTargets(buf *Metrics) *metricsTargets {
	return &metricsTargets{
		List: listTargets{
			Margin:               &buf.List.Margin,
			ItemH:                &buf.List.ItemH,
			ItemGap:              &buf.List.ItemGap,
			PortraitSize:         &buf.List.PortraitSize,
			TitleOffset:          &buf.List.TitleOffset,
			HeaderTopGap:         &buf.List.HeaderTopGap,
			ItemsTopGap:          &buf.List.ItemsTopGap,
			PanelInnerPaddingX:   &buf.List.PanelInnerPaddingX,
			TitleXOffset:         &buf.List.TitleXOffset,
			HeaderBaseX:          &buf.List.HeaderBaseX,
			RowTextOffsetX:       &buf.List.RowTextOffsetX,
			RowTextOffsetY:       &buf.List.RowTextOffsetY,
			RowBorderPad:         &buf.List.RowBorderPad,
			RowRightIconSize:     &buf.List.RowRightIconSize,
			RowRightIconGap:      &buf.List.RowRightIconGap,
			HeaderColumnsItems:   &buf.List.HeaderColumnsItems,
			HeaderColumnsWeapons: &buf.List.HeaderColumnsWeapons,
			RowColumnsItems:      &buf.List.RowColumnsItems,
			RowColumnsWeapons:    &buf.List.RowColumnsWeapons,
		},
		Line: lineTargets{
			Main:  &buf.Line.Main,
			Small: &buf.Line.Small,
		},
	}
}

func captureStatusTargets(buf *Metrics) *metricsTargets {
	return &metricsTargets{
		Status: statusTargets{
			PanelPad:           &buf.Status.PanelPad,
			PortraitSize:       &buf.Status.PortraitSize,
			TextGapX:           &buf.Status.TextGapX,
			NameOffsetY:        &buf.Status.NameOffsetY,
			ClassGapFromName:   &buf.Status.ClassGapFromName,
			LevelGapFromName:   &buf.Status.LevelGapFromName,
			HPGapFromName:      &buf.Status.HPGapFromName,
			HPBarGapFromName:   &buf.Status.HPBarGapFromName,
			HPBarW:             &buf.Status.HPBarW,
			HPBarH:             &buf.Status.HPBarH,
			StatsTopGap:        &buf.Status.StatsTopGap,
			StatsLineH:         &buf.Status.StatsLineH,
			StatsColGap:        &buf.Status.StatsColGap,
			WeaponRanksXExtra:  &buf.Status.WeaponRanksXExtra,
			RankLineH:          &buf.Status.RankLineH,
			MagicRanksTopExtra: &buf.Status.MagicRanksTopExtra,
			EquipTitleGapY:     &buf.Status.EquipTitleGapY,
			EquipLineH:         &buf.Status.EquipLineH,
			EquipRectYOffset:   &buf.Status.EquipRectYOffset,
			EquipRectW:         &buf.Status.EquipRectW,
			EquipRectH:         &buf.Status.EquipRectH,
			EquipLabelGapX:     &buf.Status.EquipLabelGapX,
			EquipUsesX:         &buf.Status.EquipUsesX,
		},
	}
}

func captureSimTargets(buf *Metrics) *metricsTargets {
	return &metricsTargets{
		Sim: simTargets{
			StartBtnW:              &buf.Sim.StartBtnW,
			StartBtnH:              &buf.Sim.StartBtnH,
			AutoRunGap:             &buf.Sim.AutoRunGap,
			TitleYOffset:           &buf.Sim.TitleYOffset,
			TitleXOffsetFromCenter: &buf.Sim.TitleXOffsetFromCenter,
			Terrain: terrainTargets{
				ButtonW:                &buf.Sim.Terrain.ButtonW,
				ButtonH:                &buf.Sim.Terrain.ButtonH,
				BaseYFromBottom:        &buf.Sim.Terrain.BaseYFromBottom,
				LeftBaseXOffset:        &buf.Sim.Terrain.LeftBaseXOffset,
				RightBaseXInset:        &buf.Sim.Terrain.RightBaseXInset,
				ButtonGap:              &buf.Sim.Terrain.ButtonGap,
				LabelLeftXOffset:       &buf.Sim.Terrain.LabelLeftXOffset,
				LabelYOffsetFromBottom: &buf.Sim.Terrain.LabelYOffsetFromBottom,
			},
			Preview: previewTargets{
				LeftXPad:       &buf.Sim.Preview.LeftXPad,
				RightXInset:    &buf.Sim.Preview.RightXInset,
				TopYFromMargin: &buf.Sim.Preview.TopYFromMargin,
				CardW:          &buf.Sim.Preview.CardW,
				CardH:          &buf.Sim.Preview.CardH,
				CardInnerPad:   &buf.Sim.Preview.CardInnerPad,
				PortraitSize:   &buf.Sim.Preview.PortraitSize,
				ClassOffsetY:   &buf.Sim.Preview.ClassOffsetY,
				HPLabelX:       &buf.Sim.Preview.HPLabelX,
				HPLabelY:       &buf.Sim.Preview.HPLabelY,
				HPBarX:         &buf.Sim.Preview.HPBarX,
				HPBarY:         &buf.Sim.Preview.HPBarY,
				HPBarW:         &buf.Sim.Preview.HPBarW,
				HPBarH:         &buf.Sim.Preview.HPBarH,
				NameOffsetX:    &buf.Sim.Preview.NameOffsetX,
				NameOffsetY:    &buf.Sim.Preview.NameOffsetY,
				LineY:          &buf.Sim.Preview.LineY,
				BaseMin:        &buf.Sim.Preview.BaseMin,
				BaseMax:        &buf.Sim.Preview.BaseMax,
				WrapPad:        &buf.Sim.Preview.WrapPad,
				LogPadX:        &buf.Sim.Preview.LogPadX,
				LogPadY:        &buf.Sim.Preview.LogPadY,
			},
		},
	}
}

func capturePopupTargets(buf *Metrics) *metricsTargets {
	return &metricsTargets{
		Popup: popupTargets{
			Cols4ThresholdW:    &buf.Popup.Cols4ThresholdW,
			Cols3ThresholdW:    &buf.Popup.Cols3ThresholdW,
			GridInnerXTotalPad: &buf.Popup.GridInnerXTotalPad,
			CellH:              &buf.Popup.CellH,
			CellGap:            &buf.Popup.CellGap,
			GridXPad:           &buf.Popup.GridXPad,
			GridYOff:           &buf.Popup.GridYOff,
			MaxW:               &buf.Popup.MaxW,
			MaxH:               &buf.Popup.MaxH,
		},
	}
}

func captureWidgetsTargets(buf *Metrics) *metricsTargets {
	return &metricsTargets{
		Widgets: widgetsTargets{
			Back: widgetBackTargets{
				PanelRightInset: &buf.Widgets.Back.PanelRightInset,
				TopPad:          &buf.Widgets.Back.TopPad,
				W:               &buf.Widgets.Back.W,
				H:               &buf.Widgets.Back.H,
				LabelX:          &buf.Widgets.Back.LabelX,
				LabelY:          &buf.Widgets.Back.LabelY,
			},
			LevelUp: widgetLevelUpTargets{
				W:      &buf.Widgets.LevelUp.W,
				H:      &buf.Widgets.LevelUp.H,
				LabelX: &buf.Widgets.LevelUp.LabelX,
				LabelY: &buf.Widgets.LevelUp.LabelY,
			},
			ToBattle: widgetToBattleTargets{
				GapFromRightBtn: &buf.Widgets.ToBattle.GapFromRightBtn,
				LabelX:          &buf.Widgets.ToBattle.LabelX,
				LabelY:          &buf.Widgets.ToBattle.LabelY,
				W:               &buf.Widgets.ToBattle.W,
			},
			SimBattle: widgetSimBattleTargets{
				W:      &buf.Widgets.SimBattle.W,
				H:      &buf.Widgets.SimBattle.H,
				TopPad: &buf.Widgets.SimBattle.TopPad,
				LabelX: &buf.Widgets.SimBattle.LabelX,
				LabelY: &buf.Widgets.SimBattle.LabelY,
			},
		},
	}
}

func assertSliceCopied(t *testing.T, name string, got, src []int) {
	t.Helper()
	if !reflect.DeepEqual(got, src) {
		t.Fatalf("%s mismatch: got=%v want=%v", name, got, src)
	}
	if len(got) > 0 && len(src) > 0 && &got[0] == &src[0] {
		t.Fatalf("%s should be copied slice", name)
	}
}
