package ui

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "github.com/hajimehoshi/ebiten/v2/vector"
)

func BackButtonRect(sw, _ int) (x, y, w, h int) {
    panelX, panelY := listMargin, listMargin
    panelW := sw - listMargin*2
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
    if hovered { bg = color.RGBA{70, 100, 150, 255} }
    drawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
    vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
    textDraw(dst, "＜ 一覧へ", faceMain, x+20, y+32, colText)
}

func LevelUpButtonRect(sw, sh int) (x, y, w, h int) {
    w, h = 220, 56
    x = sw - listMargin - w
    y = sh - listMargin - h
    return
}

func DrawLevelUpButton(dst *ebiten.Image, hovered, enabled bool) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    x, y, w, h := LevelUpButtonRect(sw, sh)
    base := color.RGBA{80, 130, 60, 255}
    if !enabled { base = color.RGBA{70, 70, 70, 255} }
    if hovered && enabled { base = color.RGBA{100, 170, 80, 255} }
    drawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
    vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
    label := "レベルアップ"
    if !enabled { label = "最大レベル" }
    textDraw(dst, label, faceMain, x+24, y+36, colText)
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
    if !enabled { base = color.RGBA{70, 70, 70, 255} }
    if hovered && enabled { base = color.RGBA{110, 110, 170, 255} }
    drawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
    vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
    textDraw(dst, "戦闘へ", faceMain, x+70, y+36, colText)
}

// textDraw は text.Draw の薄いラッパ（将来的に text/v2 へ移行するための集約点）。
func textDraw(dst *ebiten.Image, s string, face fontFace, x, y int, clr color.Color) {
    // 型エイリアスを避けるために別関数化
    text.Draw(dst, s, face, x, y, clr)
}

// fontFace は text.Draw の引数に合わせた別名インタフェース。
type fontFace interface{}

