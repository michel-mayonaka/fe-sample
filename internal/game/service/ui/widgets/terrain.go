package uiwidgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	uicore "ui_sample/internal/game/service/ui"
)

// TerrainButtonRect は左右それぞれの地形選択ボタン矩形を返します。
// left=true が攻撃側（左列）、false が防御側（右列）。idx=0:平地,1:森,2:砦。
func TerrainButtonRect(sw, sh int, left bool, idx int) (x, y, w, h int) {
	lm := uicore.ListMarginPx()
	w, h = uicore.TerrainBtnWPx(), uicore.TerrainBtnHPx()
	baseY := sh - lm - uicore.TerrainBaseYFromBottomPx()
	// 左列/右列の基準X
	baseX := lm + uicore.TerrainLeftBaseXOffsetPx()
	if !left {
		baseX = sw - lm - uicore.TerrainRightBaseXInsetPx()
	}
	x = baseX + idx*(w+uicore.TerrainBtnGapPx())
	y = baseY
	return
}

// DrawTerrainButtons は左右に 3 つずつの地形ボタンを描画します。
// attSel/defSel は 0..2 の選択インデックスです。
func DrawTerrainButtons(dst *ebiten.Image, attSel, defSel int) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	// ラベル
	uicore.TextDraw(dst, "攻地形", uicore.FaceSmall, uicore.ListMarginPx()+uicore.TerrainLabelLeftXOffsetPx(), sh-uicore.ListMarginPx()-uicore.TerrainLabelYOffsetFromBottomPx(), uicore.ColAccent)
	uicore.TextDraw(dst, "防地形", uicore.FaceSmall, sw-uicore.ListMarginPx()-uicore.TerrainRightBaseXInsetPx(), sh-uicore.ListMarginPx()-uicore.TerrainLabelYOffsetFromBottomPx(), uicore.ColAccent)
	// 左右3つずつ
	names := []string{"平地", "森", "砦"}
	for i := 0; i < 3; i++ {
		lx, ly, lw, lh := TerrainButtonRect(sw, sh, true, i)
		drawTerrainButton(dst, lx, ly, lw, lh, names[i], i == attSel)
		rx, ry, rw, rh := TerrainButtonRect(sw, sh, false, i)
		drawTerrainButton(dst, rx, ry, rw, rh, names[i], i == defSel)
	}
}

func drawTerrainButton(dst *ebiten.Image, x, y, w, h int, label string, selected bool) {
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	base := color.RGBA{40, 60, 100, 255}
	if selected {
		base = color.RGBA{70, 100, 160, 255}
	}
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
	uicore.TextDraw(dst, label, uicore.FaceSmall, x+uicore.S(20), y+uicore.S(26), uicore.ColText)
}
