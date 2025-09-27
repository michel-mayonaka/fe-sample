package uiscreens

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck
	"golang.org/x/image/font"
	"ui_sample/internal/game"
	"ui_sample/internal/ui/core"
)

func DrawStatus(dst *ebiten.Image, u uicore.Unit) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	margin := float32(24)
	panelX, panelY := margin, margin
	panelW, panelH := float32(sw)-margin*2, float32(sh)-margin*2
	uicore.DrawPanel(dst, panelX, panelY, panelW, panelH)
	px, py := panelX+24, panelY+24
	pw, ph := float32(320), float32(320)
	uicore.DrawFramedRect(dst, px, py, pw, ph)
	if u.Portrait != nil {
		uicore.DrawPortrait(dst, u.Portrait, px, py, pw, ph)
	} else {
		uicore.DrawPortraitPlaceholder(dst, px, py, pw, ph)
	}
	tx := int(px + pw + 32)
	ty := int(py + 44)
	text.Draw(dst, u.Name, uicore.FaceTitle, tx, ty, uicore.ColAccent)
	text.Draw(dst, u.Class, uicore.FaceMain, tx, ty+40, uicore.ColText)
	text.Draw(dst, fmt.Sprintf("Lv %d / %d    経験値 %02d / %d", u.Level, game.LevelCap, u.Exp, game.LevelUpExp), uicore.FaceMain, tx, ty+70, uicore.ColText)
	text.Draw(dst, fmt.Sprintf("HP %d/%d", u.HP, u.HPMax), uicore.FaceMain, tx, ty+110, uicore.ColText)
	uicore.DrawHPBar(dst, tx, ty+116, 300, 14, u.HP, u.HPMax)
	statsTop := ty + 160
	line := 34
	colGap := 180
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+0*line, "力", u.Stats.Str, u.Growth.Str)
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+1*line, "魔力", u.Stats.Mag, u.Growth.Mag)
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+2*line, "技", u.Stats.Skl, u.Growth.Skl)
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+3*line, "速さ", u.Stats.Spd, u.Growth.Spd)
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+0*line, "幸運", u.Stats.Lck, u.Growth.Lck)
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+1*line, "守備", u.Stats.Def, u.Growth.Def)
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+2*line, "魔防", u.Stats.Res, u.Growth.Res)
	drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+3*line, "移動", u.Stats.Mov, u.Growth.Mov)
	wrX := tx + 2*colGap + 64
	wrY := ty
	text.Draw(dst, "武器レベル", uicore.FaceMain, wrX, wrY, uicore.ColAccent)
	rline := 32
	drawRankLine(dst, uicore.FaceMain, wrX, wrY+1*rline, "剣", u.Weapon.Sword)
	drawRankLine(dst, uicore.FaceMain, wrX, wrY+2*rline, "槍", u.Weapon.Lance)
	drawRankLine(dst, uicore.FaceMain, wrX, wrY+3*rline, "斧", u.Weapon.Axe)
	drawRankLine(dst, uicore.FaceMain, wrX, wrY+4*rline, "弓", u.Weapon.Bow)
	mrX := wrX
	mrY := wrY + (4+1)*rline + 16
	text.Draw(dst, "魔法レベル", uicore.FaceMain, mrX, mrY, uicore.ColAccent)
	drawRankLine(dst, uicore.FaceMain, mrX, mrY+1*rline, "理", u.Magic.Anima)
	drawRankLine(dst, uicore.FaceMain, mrX, mrY+2*rline, "光", u.Magic.Light)
	drawRankLine(dst, uicore.FaceMain, mrX, mrY+3*rline, "闇", u.Magic.Dark)
	drawRankLine(dst, uicore.FaceMain, mrX, mrY+4*rline, "杖", u.Magic.Staff)
	equipTitleY := int(py + ph + 56)
	text.Draw(dst, "装備", uicore.FaceMain, int(px), equipTitleY, uicore.ColAccent)
	for i, it := range u.Equip {
		lineY := equipTitleY + 30 + i*30
		text.Draw(dst, "- "+it.Name, uicore.FaceSmall, int(px)+14, lineY, uicore.ColText)
		uses := "-"
		if it.Max > 0 {
			uses = fmt.Sprintf("%d/%d", it.Uses, it.Max)
		}
		text.Draw(dst, uses, uicore.FaceSmall, int(px)+300, lineY, uicore.ColAccent)
	}
}

// ローカル描画補助
func drawStatLineWithGrowth(dst *ebiten.Image, face font.Face, x, y int, label string, v, g int) {
	text.Draw(dst, label, face, x, y, uicore.ColText)
	text.Draw(dst, fmt.Sprintf("%2d", v), face, x+64, y, uicore.ColAccent)
	text.Draw(dst, fmt.Sprintf("%d%%", g), uicore.FaceSmall, x+120, y, uicore.ColAccent)
}

func drawRankLine(dst *ebiten.Image, face font.Face, x, y int, label, rank string) {
	if rank == "" {
		rank = "-"
	}
	text.Draw(dst, label, face, x, y, uicore.ColText)
	text.Draw(dst, rank, face, x+120, y, uicore.ColAccent)
}
