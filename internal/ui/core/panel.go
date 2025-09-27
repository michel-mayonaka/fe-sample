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
    // 背景と枠
    bx, by, bw, bh := float32(x), float32(y), float32(w), float32(h)
    border := float32(S(2))
    if border < 1 { border = 1 }
    // 枠（外側）
    vector.DrawFilledRect(dst, bx-border, by-border, bw+border*2, bh+border*2, ColBorder, false)
    // 背景
    vector.DrawFilledRect(dst, bx, by, bw, bh, color.RGBA{40, 48, 64, 255}, false)
    ratio := float32(hp) / float32(maxHP)
    bw := float32(w) * ratio
    col := color.RGBA{80, 220, 100, 255}
    if ratio < 0.33 {
        col = color.RGBA{220, 80, 80, 255}
    } else if ratio < 0.66 {
        col = color.RGBA{240, 200, 80, 255}
    }
    vector.DrawFilledRect(dst, bx, by, bw, bh, col, false)
}
