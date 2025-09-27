package uiscreens

import (
	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"ui_sample/internal/ui/core"
)

func BattleStartButtonRect(sw, sh int) (x, y, w, h int) {
	w, h = 240, 60
	x = (sw - w) / 2
	y = sh - uicore.ListMargin - h
	return
}

func DrawBattle(dst *ebiten.Image, attacker, defender uicore.Unit) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	uicore.DrawPanel(dst, uicore.ListMargin, uicore.ListMargin, float32(sw-2*uicore.ListMargin), float32(sh-2*uicore.ListMargin))
	leftX := uicore.ListMargin + 40
	rightX := sw - uicore.ListMargin - 560
	topY := uicore.ListMargin + 80
	drawBattleSide(dst, attacker, leftX, topY)
	drawBattleSide(dst, defender, rightX, topY)
	text.Draw(dst, "戦闘プレビュー", uicore.FaceTitle, sw/2-120, uicore.ListMargin+56, uicore.ColAccent)
	bx, by, bw, bh := BattleStartButtonRect(sw, sh)
	uicore.DrawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
	vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), color.RGBA{110, 90, 40, 255}, false)
	text.Draw(dst, "戦闘開始", uicore.FaceMain, bx+70, by+38, uicore.ColText)
}

func drawBattleSide(dst *ebiten.Image, u uicore.Unit, x, y int) {
	uicore.DrawFramedRect(dst, float32(x), float32(y), 320, 320)
	if u.Portrait != nil {
		uicore.DrawPortrait(dst, u.Portrait, float32(x), float32(y), 320, 320)
	}
	text.Draw(dst, u.Name, uicore.FaceTitle, x, y-16, uicore.ColText)
	text.Draw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceMain, x, y+350, uicore.ColAccent)
	text.Draw(dst, uicore.Itoa(u.HP)+"/"+uicore.Itoa(u.HPMax), uicore.FaceMain, x, y+384, uicore.ColText)
	uicore.DrawHPBar(dst, x, y+390, 320, 14, u.HP, u.HPMax)
	wep := "-"
	if len(u.Equip) > 0 {
		wep = u.Equip[0].Name
	}
	text.Draw(dst, "武器: "+wep, uicore.FaceMain, x, y+420, uicore.ColText)
}
