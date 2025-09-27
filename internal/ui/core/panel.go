package uicore

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func DrawPanel(dst *ebiten.Image, x, y, w, h float32) {
	vector.DrawFilledRect(dst, x-2, y-2, w+4, h+4, ColBorder, false)
	vector.DrawFilledRect(dst, x+2, y+2, w, h, ColPanelDark, false)
	vector.DrawFilledRect(dst, x, y, w, h, ColPanelBG, false)
}

func DrawFramedRect(dst *ebiten.Image, x, y, w, h float32) {
	vector.DrawFilledRect(dst, x-2, y-2, w+4, h+4, ColBorder, false)
	vector.DrawFilledRect(dst, x, y, w, h, color.RGBA{30, 45, 78, 255}, false)
}

func DrawHPBar(dst *ebiten.Image, x, y, w, h int, hp, maxHP int) {
	if maxHP <= 0 {
		maxHP = 1
	}
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), color.RGBA{50, 50, 50, 255}, false)
	ratio := float32(hp) / float32(maxHP)
	bw := float32(w) * ratio
	col := color.RGBA{80, 220, 100, 255}
	if ratio < 0.33 {
		col = color.RGBA{220, 80, 80, 255}
	} else if ratio < 0.66 {
		col = color.RGBA{240, 200, 80, 255}
	}
	vector.DrawFilledRect(dst, float32(x), float32(y), bw, float32(h), col, false)
}
