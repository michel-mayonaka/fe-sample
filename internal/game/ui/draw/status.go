package draw

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "golang.org/x/image/font"
    "ui_sample/internal/adapter"
    "ui_sample/internal/game"
    uicore "ui_sample/internal/game/service/ui"
    gdata "ui_sample/internal/game/data"
    "ui_sample/internal/model"
)

// DrawStatus はステータス画面を描画します。
func DrawStatus(dst *ebiten.Image, u uicore.Unit) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := float32(uicore.ListMarginPx())
    panelX, panelY := lm, lm
    panelW, panelH := float32(sw)-lm*2, float32(sh)-lm*2
    uicore.DrawPanel(dst, panelX, panelY, panelW, panelH)
    px, py := panelX+float32(uicore.S(24)), panelY+float32(uicore.S(24))
    pw, ph := float32(uicore.S(320)), float32(uicore.S(320))
    tx := int(px + pw + float32(uicore.S(32)))
    ty := int(py + float32(uicore.S(44)))
    drawStatusHeader(dst, u, int(px), int(py), int(pw), int(ph), tx, ty)
    statsTop := ty + uicore.S(160)
    line := uicore.S(34)
    colGap := uicore.S(180)
    drawCoreStats(dst, u, tx, statsTop, line, colGap)
    wrX := tx + 2*colGap + uicore.S(64)
    wrY := ty
    drawWeaponRanks(dst, u, wrX, wrY)
    mrX := wrX
    mrY := wrY + (4+1)*uicore.S(32) + uicore.S(16)
    drawMagicRanks(dst, u, mrX, mrY)
    drawEquipList(dst, u, int(px), int(py), int(ph))
}

// ヘッダ（ポートレート/名前/クラス/レベル/HP）
func drawStatusHeader(dst *ebiten.Image, u uicore.Unit, px, py, pw, ph, tx, ty int) {
    uicore.DrawFramedRect(dst, float32(px), float32(py), float32(pw), float32(ph))
    if u.Portrait != nil {
        uicore.DrawPortrait(dst, u.Portrait, float32(px), float32(py), float32(pw), float32(ph))
    } else {
        uicore.DrawPortraitPlaceholder(dst, float32(px), float32(py), float32(pw), float32(ph))
    }
    uicore.TextDraw(dst, u.Name, uicore.FaceTitle, tx, ty, uicore.ColAccent)
    uicore.TextDraw(dst, u.Class, uicore.FaceMain, tx, ty+uicore.S(40), uicore.ColText)
    uicore.TextDraw(dst, fmt.Sprintf("Lv %d / %d    経験値 %02d / %d", u.Level, game.LevelCap, u.Exp, game.LevelUpExp), uicore.FaceMain, tx, ty+uicore.S(70), uicore.ColText)
    uicore.TextDraw(dst, fmt.Sprintf("HP %d/%d", u.HP, u.HPMax), uicore.FaceMain, tx, ty+uicore.S(110), uicore.ColText)
    uicore.DrawHPBar(dst, tx, ty+uicore.S(116), uicore.S(300), uicore.S(14), u.HP, u.HPMax)
}

// 基本ステータス（成長率付き）+ 派生（攻撃速度）
func drawCoreStats(dst *ebiten.Image, u uicore.Unit, tx, statsTop, line, colGap int) {
    var wt *model.WeaponTable
    if p := gdata.Provider(); p != nil { wt = p.WeaponsTable() }
    atkSpeed := adapter.AttackSpeedOf(wt, u)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+0*line, "力", u.Stats.Str, u.Growth.Str)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+1*line, "魔力", u.Stats.Mag, u.Growth.Mag)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+2*line, "技", u.Stats.Skl, u.Growth.Skl)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+0*colGap, statsTop+3*line, "速さ", u.Stats.Spd, u.Growth.Spd)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+0*line, "幸運", u.Stats.Lck, u.Growth.Lck)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+1*line, "守備", u.Stats.Def, u.Growth.Def)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+2*line, "魔防", u.Stats.Res, u.Growth.Res)
    drawStatLineWithGrowth(dst, uicore.FaceMain, tx+1*colGap, statsTop+3*line, "体格", u.Stats.Bld, u.Growth.Bld)
    drawStatLine(dst, uicore.FaceMain, tx+1*colGap, statsTop+5*line, "攻撃速度", atkSpeed)
}

// 武器ランク
func drawWeaponRanks(dst *ebiten.Image, u uicore.Unit, x, y int) {
    uicore.TextDraw(dst, "武器レベル", uicore.FaceMain, x, y, uicore.ColAccent)
    rline := uicore.S(32)
    drawRankLine(dst, uicore.FaceMain, x, y+1*rline, "剣", u.Weapon.Sword)
    drawRankLine(dst, uicore.FaceMain, x, y+2*rline, "槍", u.Weapon.Lance)
    drawRankLine(dst, uicore.FaceMain, x, y+3*rline, "斧", u.Weapon.Axe)
    drawRankLine(dst, uicore.FaceMain, x, y+4*rline, "弓", u.Weapon.Bow)
}

// 魔法ランク
func drawMagicRanks(dst *ebiten.Image, u uicore.Unit, x, y int) {
    uicore.TextDraw(dst, "魔法レベル", uicore.FaceMain, x, y, uicore.ColAccent)
    rline := uicore.S(32)
    drawRankLine(dst, uicore.FaceMain, x, y+1*rline, "理", u.Magic.Anima)
    drawRankLine(dst, uicore.FaceMain, x, y+2*rline, "光", u.Magic.Light)
    drawRankLine(dst, uicore.FaceMain, x, y+3*rline, "闇", u.Magic.Dark)
    drawRankLine(dst, uicore.FaceMain, x, y+4*rline, "杖", u.Magic.Staff)
}

// 装備リスト
func drawEquipList(dst *ebiten.Image, u uicore.Unit, px, py, ph int) {
    equipTitleY := py + ph + uicore.S(56)
    uicore.TextDraw(dst, "装備", uicore.FaceMain, px, equipTitleY, uicore.ColAccent)
    for i := 0; i < 5; i++ {
        lineY := equipTitleY + uicore.S(30) + i*uicore.S(30)
        label := "- 空 -"
        uses := "-"
        if i < len(u.Equip) {
            it := u.Equip[i]
            if it.Name != "" {
                label = "- " + it.Name
                if it.Max > 0 { uses = fmt.Sprintf("%d/%d", it.Uses, it.Max) }
            }
        }
        uicore.TextDraw(dst, label, uicore.FaceSmall, px+uicore.S(14), lineY, uicore.ColText)
        uicore.TextDraw(dst, uses, uicore.FaceSmall, px+uicore.S(300), lineY, uicore.ColAccent)
    }
}

// ローカル描画補助
func drawStatLineWithGrowth(dst *ebiten.Image, face font.Face, x, y int, label string, v, g int) {
    uicore.TextDraw(dst, label, face, x, y, uicore.ColText)
    uicore.TextDraw(dst, fmt.Sprintf("%2d", v), face, x+uicore.S(64), y, uicore.ColAccent)
    uicore.TextDraw(dst, fmt.Sprintf("%d%%", g), uicore.FaceSmall, x+uicore.S(120), y, uicore.ColAccent)
}

func drawRankLine(dst *ebiten.Image, face font.Face, x, y int, label, rank string) {
    if rank == "" { rank = "-" }
    uicore.TextDraw(dst, label, face, x, y, uicore.ColText)
    uicore.TextDraw(dst, rank, face, x+uicore.S(120), y, uicore.ColAccent)
}

// 成長率のない派生値用の簡易行描画。
func drawStatLine(dst *ebiten.Image, face font.Face, x, y int, label string, v int) {
    uicore.TextDraw(dst, label, face, x, y, uicore.ColText)
    lw := int(font.MeasureString(face, label) >> 6)
    gap := uicore.S(20)
    uicore.TextDraw(dst, fmt.Sprintf("%2d", v), face, x+lw+gap, y, uicore.ColAccent)
}

