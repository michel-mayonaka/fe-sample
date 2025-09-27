package uiwidgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"ui_sample/internal/ui/core"
)

func BackButtonRect(sw, _ int) (x, y, w, h int) {
	panelX, panelY := uicore.ListMargin, uicore.ListMargin
	panelW := sw - uicore.ListMargin*2
	x = panelX + panelW - 180
	y = panelY + 24
	w = 160
	h = 48
	return
}

func DrawBackButton(dst *ebiten.Image, hovered bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := BackButtonRect(sw, sh)
	bg := color.RGBA{50, 70, 110, 255}
	if hovered {
		bg = color.RGBA{70, 100, 150, 255}
	}
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
	text.Draw(dst, "＜ 一覧へ", uicore.FaceMain, x+20, y+32, uicore.ColText)
}

func LevelUpButtonRect(sw, sh int) (x, y, w, h int) {
	w, h = 220, 56
	x = sw - uicore.ListMargin - w
	y = sh - uicore.ListMargin - h
	return
}

func DrawLevelUpButton(dst *ebiten.Image, hovered, enabled bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := LevelUpButtonRect(sw, sh)
	base := color.RGBA{80, 130, 60, 255}
	if !enabled {
		base = color.RGBA{70, 70, 70, 255}
	}
	if hovered && enabled {
		base = color.RGBA{100, 170, 80, 255}
	}
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
	label := "レベルアップ"
	if !enabled {
		label = "最大レベル"
	}
	text.Draw(dst, label, uicore.FaceMain, x+24, y+36, uicore.ColText)
}

func ToBattleButtonRect(sw, sh int) (x, y, w, h int) {
	rx, ry, _, rh := LevelUpButtonRect(sw, sh)
	w, h = 220, rh
	x = rx - 20 - w
	y = ry
	return
}

func DrawToBattleButton(dst *ebiten.Image, hovered, enabled bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := ToBattleButtonRect(sw, sh)
	base := color.RGBA{90, 90, 130, 255}
	if !enabled {
		base = color.RGBA{70, 70, 70, 255}
	}
	if hovered && enabled {
		base = color.RGBA{110, 110, 170, 255}
	}
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
	text.Draw(dst, "戦闘へ", uicore.FaceMain, x+70, y+36, uicore.ColText)
}

// SimBattleButtonRect は一覧画面の右上に表示する「模擬戦」ボタンの矩形です。
func SimBattleButtonRect(sw, _ int) (x, y, w, h int) {
	w, h = 160, 48
	x = sw - uicore.ListMargin - w
	y = uicore.ListMargin + 16
	return
}

// DrawSimBattleButton は一覧画面の「模擬戦」ボタンを描画します。
func DrawSimBattleButton(dst *ebiten.Image, hovered, enabled bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := SimBattleButtonRect(sw, sh)
	base := color.RGBA{40, 60, 100, 255}
	if !enabled {
		base = color.RGBA{70, 70, 70, 255}
	}
	if hovered && enabled {
		base = color.RGBA{60, 90, 150, 255}
	}
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
	label := "模擬戦"
	if hovered && enabled {
		label = "> 模擬戦 <"
	}
	text.Draw(dst, label, uicore.FaceMain, x+24, y+32, uicore.ColText)
}
