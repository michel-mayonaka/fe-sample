package ui

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "github.com/hajimehoshi/ebiten/v2/vector"
)

func BattleStartButtonRect(sw, sh int) (x, y, w, h int) {
    w, h = 240, 60
    x = (sw - w) / 2
    y = sh - listMargin - h
    return
}

func DrawBattle(dst *ebiten.Image, attacker, defender Unit) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    drawPanel(dst, listMargin, listMargin, float32(sw-2*listMargin), float32(sh-2*listMargin))
    leftX := listMargin + 40
    rightX := sw - listMargin - 560
    topY := listMargin + 80
    drawBattleSide(dst, attacker, leftX, topY)
    drawBattleSide(dst, defender, rightX, topY)
    text.Draw(dst, "戦闘プレビュー", faceTitle, sw/2-120, listMargin+56, colAccent)
    bx, by, bw, bh := BattleStartButtonRect(sw, sh)
    drawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
    vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), color.RGBA{110, 90, 40, 255}, false)
    textDraw(dst, "戦闘開始", faceMain, bx+70, by+38, colText)
}

func drawBattleSide(dst *ebiten.Image, u Unit, x, y int) {
    drawFramedRect(dst, float32(x), float32(y), 320, 320)
    if u.Portrait != nil { drawPortrait(dst, u.Portrait, float32(x), float32(y), 320, 320) }
    textDraw(dst, u.Name, faceTitle, x, y-16, colText)
    textDraw(dst, u.Class+"  Lv "+itoa(u.Level), faceMain, x, y+350, colAccent)
    textDraw(dst, itoa(u.HP)+"/"+itoa(u.HPMax), faceMain, x, y+384, colText)
    drawHPBar(dst, x, y+390, 320, 14, u.HP, u.HPMax)
    wep := "-"; if len(u.Equip)>0 { wep = u.Equip[0].Name }
    textDraw(dst, "武器: "+wep, faceMain, x, y+420, colText)
}

