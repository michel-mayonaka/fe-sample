package ui

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "ui_sample/internal/game"
)

func DrawStatus(dst *ebiten.Image, u Unit) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    margin := float32(24)
    panelX, panelY := margin, margin
    panelW, panelH := float32(sw)-margin*2, float32(sh)-margin*2
    drawPanel(dst, panelX, panelY, panelW, panelH)

    px, py := panelX+24, panelY+24
    pw, ph := float32(320), float32(320)
    drawFramedRect(dst, px, py, pw, ph)
    if u.Portrait != nil { drawPortrait(dst, u.Portrait, px, py, pw, ph) } else { drawPortraitPlaceholder(dst, px, py, pw, ph) }

    tx := int(px + pw + 32)
    ty := int(py + 44)
    text.Draw(dst, u.Name, faceTitle, tx, ty, colAccent)
    text.Draw(dst, u.Class, faceMain, tx, ty+40, colText)
    text.Draw(dst, fmt.Sprintf("Lv %d / %d    経験値 %02d / %d", u.Level, game.LevelCap, u.Exp, game.LevelUpExp), faceMain, tx, ty+70, colText)

    text.Draw(dst, fmt.Sprintf("HP %d/%d", u.HP, u.HPMax), faceMain, tx, ty+110, colText)
    drawHPBar(dst, tx, ty+116, 300, 14, u.HP, u.HPMax)

    statsTop := ty + 160
    line := 34
    colGap := 180
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+0*line, "力", u.Stats.Str, u.Growth.Str)
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+1*line, "魔力", u.Stats.Mag, u.Growth.Mag)
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+2*line, "技", u.Stats.Skl, u.Growth.Skl)
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+3*line, "速さ", u.Stats.Spd, u.Growth.Spd)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+0*line, "幸運", u.Stats.Lck, u.Growth.Lck)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+1*line, "守備", u.Stats.Def, u.Growth.Def)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+2*line, "魔防", u.Stats.Res, u.Growth.Res)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+3*line, "移動", u.Stats.Mov, u.Growth.Mov)

    // 武器レベル・魔法レベル
    wrX := tx + 2*colGap + 64
    wrY := ty
    text.Draw(dst, "武器レベル", faceMain, wrX, wrY, colAccent)
    rline := 32
    drawRankLine(dst, faceMain, wrX, wrY+1*rline, "剣", u.Weapon.Sword)
    drawRankLine(dst, faceMain, wrX, wrY+2*rline, "槍", u.Weapon.Lance)
    drawRankLine(dst, faceMain, wrX, wrY+3*rline, "斧", u.Weapon.Axe)
    drawRankLine(dst, faceMain, wrX, wrY+4*rline, "弓", u.Weapon.Bow)
    mrX := wrX
    mrY := wrY + (4+1)*rline + 16
    text.Draw(dst, "魔法レベル", faceMain, mrX, mrY, colAccent)
    drawRankLine(dst, faceMain, mrX, mrY+1*rline, "理", u.Magic.Anima)
    drawRankLine(dst, faceMain, mrX, mrY+2*rline, "光", u.Magic.Light)
    drawRankLine(dst, faceMain, mrX, mrY+3*rline, "闇", u.Magic.Dark)
    drawRankLine(dst, faceMain, mrX, mrY+4*rline, "杖", u.Magic.Staff)

    // 装備
    equipTitleY := int(py + ph + 56)
    text.Draw(dst, "装備", faceMain, int(px), equipTitleY, colAccent)
    for i, it := range u.Equip {
        lineY := equipTitleY + 30 + i*30
        text.Draw(dst, "- "+it.Name, faceSmall, int(px)+14, lineY, colText)
        uses := "-"; if it.Max > 0 { uses = fmt.Sprintf("%d/%d", it.Uses, it.Max) }
        text.Draw(dst, uses, faceSmall, int(px)+300, lineY, colAccent)
    }
}

