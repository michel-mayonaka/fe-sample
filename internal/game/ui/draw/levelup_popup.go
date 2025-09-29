package draw

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	lvl "ui_sample/internal/game/service/levelup"
	uicore "ui_sample/internal/game/service/ui"
	uilayout "ui_sample/internal/game/ui/layout"
)

// DrawLevelUpPopup はレベルアップ結果のポップアップを描画します。
func DrawLevelUpPopup(dst *ebiten.Image, u uicore.Unit, gains lvl.Gains) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	vector.DrawFilledRect(dst, 0, 0, float32(sw), float32(sh), color.RGBA{0, 0, 0, 160}, false)
	pw, ph := uilayout.PopupSize(sw, sh)
	pw, ph = uicore.S(520), uicore.S(480) // 既存UIに合わせ固定サイズ
	px := (sw - pw) / 2
	py := (sh - ph) / 2
	uicore.DrawPanel(dst, float32(px), float32(py), float32(pw), float32(ph))
	uicore.TextDraw(dst, "レベルアップ!", uicore.FaceTitle, px+uicore.S(24), py+uicore.S(56), uicore.ColAccent)
	uicore.TextDraw(dst, fmt.Sprintf("Lv %d", u.Level), uicore.FaceMain, px+uicore.S(24), py+uicore.S(96), uicore.ColText)
	y := py + uicore.S(140)
	line := uicore.S(34)
	drawInc := func(label string, v int) {
		if v > 0 {
			uicore.TextDraw(dst, fmt.Sprintf("%s +%d", label, v), uicore.FaceMain, px+uicore.S(40), y, uicore.ColAccent)
			y += line
		}
	}
	if gains.HPGain > 0 {
		drawInc("HP", gains.HPGain)
	}
	drawInc("力", gains.Inc.Str)
	drawInc("魔力", gains.Inc.Mag)
	drawInc("技", gains.Inc.Skl)
	drawInc("速さ", gains.Inc.Spd)
	drawInc("幸運", gains.Inc.Lck)
	drawInc("守備", gains.Inc.Def)
	drawInc("魔防", gains.Inc.Res)
	drawInc("移動", gains.Inc.Mov)
	uicore.TextDraw(dst, "クリックで閉じる", uicore.FaceSmall, px+pw-uicore.S(180), py+ph-uicore.S(24), uicore.ColText)
}
