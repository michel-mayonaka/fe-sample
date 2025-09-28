package sim

import (
    "fmt"
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "golang.org/x/image/font"
    "ui_sample/internal/adapter"
    uicore "ui_sample/internal/game/service/ui"
    gdata "ui_sample/internal/game/data"
    "ui_sample/internal/model"
    gcore "ui_sample/pkg/game"
)

// DrawBattleWithTerrain は左右の地形を指定してプレビューを描画します。
// startEnabled が false の場合、開始ボタンはグレーアウト表示になります。
func DrawBattleWithTerrain(dst *ebiten.Image, attacker, defender uicore.Unit, attT, defT gcore.Terrain, startEnabled bool) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := uicore.ListMarginPx()
    // 枠とヘッダ
    drawBattleHeader(dst, sw, sh, lm)
    // 左右サイド
    leftX := lm + uicore.S(40)
    rightX := sw - lm - uicore.S(560)
    topY := lm + uicore.S(80)
    drawBattleSide(dst, attacker, leftX, topY)
    drawBattleSide(dst, defender, rightX, topY)
    // 開始ボタン
    drawStartButton(dst, sw, sh, startEnabled)

    // 予測値表示（/pkg/game.ForecastAt）
    if frAtk, frAtkBk, frDef, frDefBk, canCounter, ok := forecastBothWithTerrainExplain(attacker, defender, attT, defT); ok {
        ax, ay := leftX, topY+uicore.S(460)
        dx, dy := rightX, topY+uicore.S(460)
        // 左右の基本行
        drawForecastLeft(dst, ax, ay, attacker, defender, frAtk, frAtkBk)
        drawForecastRight(dst, dx, dy, sw, lm, canCounter, frDef, frDefBk)
        // 地形ラベル
        drawTerrainLabels(dst, sw, lm, ax, ay, dx, dy, attT, defT)
    }

    // ヘルプ: 地形切替（下部寄せ）
    uicore.TextDraw(dst, "[地形切替] 攻: 1=平地 2=森 3=砦 / 防: Shift+1/2/3", uicore.FaceSmall, leftX, sh-uicore.ListMarginPx()-uicore.S(190), color.RGBA{190,200,210,255})
}

func drawBattleHeader(dst *ebiten.Image, sw, sh, lm int) {
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    uicore.TextDraw(dst, "戦闘プレビュー", uicore.FaceTitle, sw/2-uicore.S(120), uicore.ListMarginPx()+uicore.S(56), uicore.ColAccent)
}

func drawStartButton(dst *ebiten.Image, sw, sh int, enabled bool) {
    bx, by, bw, bh := BattleStartButtonRect(sw, sh)
    uicore.DrawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
    fill := color.RGBA{110, 90, 40, 255}
    labelCol := uicore.ColText
    if !enabled {
        fill = color.RGBA{80, 80, 80, 200}
        labelCol = color.RGBA{200, 200, 200, 180}
    }
    vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), fill, false)
    uicore.TextDraw(dst, "戦闘開始", uicore.FaceMain, bx+uicore.S(70), by+uicore.S(38), labelCol)
}

// 左側の予測（攻撃側）: 基本行 + 相性 + 内訳
func drawForecastLeft(dst *ebiten.Image, x, y int, atk, def uicore.Unit, fr gcore.ForecastResult, bk gcore.ForecastBreakdown) {
    baseLine := fmt.Sprintf("命中 %d%%  与ダメ %d  必殺 %d%%", fr.HitDisp, fr.Dmg, fr.Crit)
    uicore.TextDraw(dst, baseLine, uicore.FaceMain, x, y, uicore.ColText)
    if p := gdata.Provider(); p != nil {
        if wt := p.WeaponsTable(); wt != nil {
        aType := weaponTypeOf(wt, atk)
        dType := weaponTypeOf(wt, def)
        lbl, col := triangleRelationLabelColor(gcore.TriangleRelationOf(aType, dType))
        w := int(font.MeasureString(uicore.FaceMain, baseLine) >> 6)
        uicore.TextDraw(dst, "  相性: ", uicore.FaceMain, x+w+8, y, uicore.ColText)
        w2 := int(font.MeasureString(uicore.FaceMain, "  相性: ") >> 6)
        uicore.TextDraw(dst, lbl, uicore.FaceMain, x+w+8+w2, y, col)
        }
    }
    // 内訳
    wrap := dynamicWrap(dst.Bounds().Dx(), uicore.ListMarginPx(), true)
    leftHitLine := fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
        bk.WeapHit, bk.Skl2, bk.LckHalf, bk.AttTileHit, bk.DefSpd2, bk.DefLck, bk.DefTileAvoid, bk.TriangleHit, bk.HitDisp)
    y2 := uicore.DrawWrapped(dst, uicore.FaceSmall, leftHitLine, x, y+uicore.S(26), uicore.ColText, wrap, uicore.LineHSmallPx())
    _ = y2
    dmgLine := fmt.Sprintf("[与ダメ内訳] 力:%d + 武器D%d + 相性:%+d - 守備合計:%d = %d",
        bk.AtkStr, bk.WpnMt, bk.TriangleMt, bk.DefTotal, bk.Dmg)
    _ = uicore.DrawWrapped(dst, uicore.FaceSmall, dmgLine, x, y+uicore.S(26)+uicore.LineHSmallPx(), uicore.ColText, wrap, uicore.LineHSmallPx())
}

// 右側の予測（防御側）
func drawForecastRight(dst *ebiten.Image, x, y, sw, lm int, canCounter bool, fr gcore.ForecastResult, bk gcore.ForecastBreakdown) {
    label := "反撃不可"
    if canCounter {
        label = fmt.Sprintf("命中 %d%%  与ダメ %d  必殺 %d%%", fr.HitDisp, fr.Dmg, fr.Crit)
    }
    wrap := dynamicWrap(sw, lm, false)
    uicore.TextDraw(dst, label, uicore.FaceMain, x, y, uicore.ColText)
    if canCounter {
        rightHitLine := fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
            bk.WeapHit, bk.Skl2, bk.LckHalf, bk.AttTileHit, bk.DefSpd2, bk.DefLck, bk.DefTileAvoid, bk.TriangleHit, bk.HitDisp)
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, rightHitLine, x, y+uicore.S(26), uicore.ColText, wrap, uicore.LineHSmallPx())
        dmgLine := fmt.Sprintf("[与ダメ内訳] 力:%d + 武器D%d + 相性:%+d - 守備合計:%d = %d",
            bk.AtkStr, bk.WpnMt, bk.TriangleMt, bk.DefTotal, bk.Dmg)
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, dmgLine, x, y+uicore.S(26)+uicore.LineHSmallPx(), uicore.ColText, wrap, uicore.LineHSmallPx())
    }
}

// 地形ラベル（左右）
func drawTerrainLabels(dst *ebiten.Image, sw, lm, ax, ay, dx, dy int, attT, defT gcore.Terrain) {
    wrapLeft := dynamicWrap(sw, lm, true)
    wrapRight := dynamicWrap(sw, lm, false)
    _ = uicore.DrawWrapped(dst, uicore.FaceSmall, terrainLine(attT), ax, ay+uicore.S(26), color.RGBA{200,220,255,255}, wrapLeft, uicore.LineHSmallPx())
    _ = uicore.DrawWrapped(dst, uicore.FaceSmall, terrainLine(defT), dx, dy+uicore.S(26), color.RGBA{200,220,255,255}, wrapRight, uicore.LineHSmallPx())
}

func drawBattleSide(dst *ebiten.Image, u uicore.Unit, x, y int) {
    sz := uicore.S(320)
    uicore.DrawFramedRect(dst, float32(x), float32(y), float32(sz), float32(sz))
    if u.Portrait != nil {
        uicore.DrawPortrait(dst, u.Portrait, float32(x), float32(y), float32(sz), float32(sz))
    }
    uicore.TextDraw(dst, u.Name, uicore.FaceTitle, x, y-uicore.S(16), uicore.ColText)
    uicore.TextDraw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceMain, x, y+uicore.S(350), uicore.ColAccent)
    uicore.TextDraw(dst, uicore.Itoa(u.HP)+"/"+uicore.Itoa(u.HPMax), uicore.FaceMain, x, y+uicore.S(384), uicore.ColText)
    uicore.DrawHPBar(dst, x, y+uicore.S(390), sz, uicore.S(14), u.HP, u.HPMax)
    wep := "-"
    if len(u.Equip) > 0 { wep = u.Equip[0].Name }
    // 攻撃速度
    var wt *model.WeaponTable
    if p := gdata.Provider(); p != nil { wt = p.WeaponsTable() }
    as := adapter.AttackSpeedOf(wt, u)
    lineY := y + uicore.S(410)
    uicore.TextDraw(dst, fmt.Sprintf("攻撃速度 %d", as), uicore.FaceSmall, x, lineY, uicore.ColText)
    lineY += uicore.LineHSmallPx()
    uicore.TextDraw(dst, "武器: "+wep, uicore.FaceSmall, x, lineY, uicore.ColText)
}

// dynamicWrap は左右列の折返し幅をウィンドウ幅とマージンから決めます。
// 左右で若干の差をつけ、最小/最大幅を設けます。
func dynamicWrap(sw, lm int, left bool) int {
    panelW := sw - 2*lm
    base := panelW/2 - uicore.S(100)
    if !left { base = panelW/2 - uicore.S(140) }
    if base < uicore.S(420) { base = uicore.S(420) }
    if base > uicore.S(900) { base = uicore.S(900) }
    return base
}

// explain 版
func forecastBothWithTerrainExplain(atk, def uicore.Unit, attT, defT gcore.Terrain) (gcore.ForecastResult, gcore.ForecastBreakdown, gcore.ForecastResult, gcore.ForecastBreakdown, bool, bool) {
    var wt *model.WeaponTable
    if p := gdata.Provider(); p != nil { wt = p.WeaponsTable() }
    if wt == nil {
        return gcore.ForecastResult{}, gcore.ForecastBreakdown{}, gcore.ForecastResult{}, gcore.ForecastBreakdown{}, false, false
    }
    ga := adapter.UIToGame(wt, atk)
    gd := adapter.UIToGame(wt, def)
    dist := 1
    canCounter := gd.W.RMin <= dist && dist <= gd.W.RMax
    frAtk, bkAtk := gcore.ForecastAtExplain(ga, gd, attT, defT)
    frDef := gcore.ForecastResult{}
    bkDef := gcore.ForecastBreakdown{}
    if canCounter {
        frDef, bkDef = gcore.ForecastAtExplain(gd, ga, defT, attT)
    }
    return frAtk, bkAtk, frDef, bkDef, canCounter, true
}

func terrainLine(t gcore.Terrain) string {
    name := gcore.TerrainPresetName(t)
    return fmt.Sprintf("地形: %s (回避+%d 防御+%d 命中+%d)", name, t.Avoid, t.Def, t.Hit)
}

func weaponTypeOf(wt *model.WeaponTable, u uicore.Unit) string {
    if len(u.Equip) == 0 { return "" }
    if w, ok := wt.Find(u.Equip[0].Name); ok { return w.Type }
    return ""
}

func triangleRelationLabelColor(rel gcore.TriangleRelation) (string, color.Color) {
    switch rel {
    case gcore.TriangleAdvantage:
        return "有利", uicore.ColAccent
    case gcore.TriangleDisadvantage:
        return "不利", color.RGBA{255, 170, 70, 255}
    default:
        return "中立", color.RGBA{180, 180, 180, 255}
    }
}

// DrawBattleLogOverlay は全画面を半透明で覆い、中央にログを表示します。
func DrawBattleLogOverlay(dst *ebiten.Image, logs []string) {
    if len(logs) == 0 { return }
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    vector.DrawFilledRect(dst, 0, 0, float32(sw), float32(sh), color.RGBA{0, 0, 0, 140}, false)
    pw, ph := int(float32(sw)*0.7), 300
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    uicore.DrawFramedRect(dst, float32(px), float32(py), float32(pw), float32(ph))
    vector.DrawFilledRect(dst, float32(px), float32(py), float32(pw), float32(ph), color.RGBA{25, 30, 50, 230}, false)
    uicore.TextDraw(dst, "戦闘ログ", uicore.FaceMain, px+16, py+32, uicore.ColAccent)
    maxLines := (ph - uicore.S(80)) / uicore.LineHSmallPx()
    start := 0
    if len(logs) > maxLines { start = len(logs) - maxLines }
    y := py + 58
    for i := start; i < len(logs); i++ {
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, logs[i], px+uicore.S(16), y, uicore.ColText, pw-uicore.S(32), uicore.LineHSmallPx())
        y += uicore.LineHSmallPx()
    }
    hint := "クリック または Z/Enter で閉じる"
    tw := int(font.MeasureString(uicore.FaceSmall, hint) >> 6)
    uicore.TextDraw(dst, hint, uicore.FaceSmall, px+(pw-tw)/2, py+ph-16, color.RGBA{210, 220, 240, 255})
}
