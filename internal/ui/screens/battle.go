// Package uiscreens は各種画面（バトル/一覧/ステータス）の描画を提供します。
package uiscreens

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "golang.org/x/image/font"
    "image/color"
    "ui_sample/internal/adapter"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/ui/core"
    gcore "ui_sample/pkg/game"
)

// 共有武器テーブル（Repo注入で設定）。未設定時は初回アクセスで読み込みキャッシュ。
var wtShared *model.WeaponTable

// SetWeaponTable は内部で共有する武器テーブルを設定します。
func SetWeaponTable(wt *model.WeaponTable) { wtShared = wt }

// BattleStartButtonRect はバトル開始ボタンの矩形を返します。
func BattleStartButtonRect(sw, sh int) (x, y, w, h int) {
    w, h = uicore.S(240), uicore.S(60)
    x = (sw - w) / 2
    y = sh - uicore.ListMarginPx() - h
    return
}

// DrawBattle は後方互換（平地扱い）。
func DrawBattle(dst *ebiten.Image, attacker, defender uicore.Unit) {
    DrawBattleWithTerrain(dst, attacker, defender, gcore.Terrain{}, gcore.Terrain{}, true)
}

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

// ヘッダとパネル
func drawBattleHeader(dst *ebiten.Image, sw, sh, lm int) {
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    uicore.TextDraw(dst, "戦闘プレビュー", uicore.FaceTitle, sw/2-uicore.S(120), uicore.ListMarginPx()+uicore.S(56), uicore.ColAccent)
}

// 開始ボタン
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
    if wt := weaponTable(); wt != nil {
        aType := weaponTypeOf(wt, atk)
        dType := weaponTypeOf(wt, def)
        lbl, col := triangleRelationLabelColor(gcore.TriangleRelationOf(aType, dType))
        w := int(font.MeasureString(uicore.FaceMain, baseLine) >> 6)
        uicore.TextDraw(dst, "  相性: ", uicore.FaceMain, x+w+8, y, uicore.ColText)
        w2 := int(font.MeasureString(uicore.FaceMain, "  相性: ") >> 6)
        uicore.TextDraw(dst, lbl, uicore.FaceMain, x+w+8+w2, y, col)
    }
    // 内訳
    wrap := dynamicWrap(dst.Bounds().Dx(), uicore.ListMarginPx(), true)
    leftHitLine := fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
        bk.WeapHit, bk.Skl2, bk.LckHalf, bk.AttTileHit, bk.DefSpd2, bk.DefLck, bk.DefTileAvoid, bk.TriangleHit, bk.HitDisp)
    y2 := uicore.DrawWrapped(dst, uicore.FaceSmall, leftHitLine, x, y+uicore.S(50), color.RGBA{210,230,255,255}, wrap, uicore.LineHSmallPx())
    leftDmgLine := fmt.Sprintf("[与ダメ内訳] 力:%d + 威力:%d + 相性:%+d - 守備合計:%d = %d", bk.AtkStr, bk.WpnMt, bk.TriangleMt, bk.DefTotal, bk.Dmg)
    _ = uicore.DrawWrapped(dst, uicore.FaceSmall, leftDmgLine, x, y2, color.RGBA{210,230,255,255}, wrap, uicore.LineHSmallPx())
}

// 右側の予測（反撃側）: 反撃可否に応じた表示 + 内訳
func drawForecastRight(dst *ebiten.Image, x, y, sw, lm int, canCounter bool, fr gcore.ForecastResult, bk gcore.ForecastBreakdown) {
    wrapRight := dynamicWrap(sw, lm, false)
    if canCounter {
        base := fmt.Sprintf("(反撃) 命中 %d%%  与ダメ %d  必殺 %d%%", fr.HitDisp, fr.Dmg, fr.Crit)
        _ = uicore.DrawWrapped(dst, uicore.FaceMain, base, x, y, uicore.ColText, wrapRight, uicore.LineHMainPx())
        rightHit := fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
            bk.WeapHit, bk.Skl2, bk.LckHalf, bk.AttTileHit, bk.DefSpd2, bk.DefLck, bk.DefTileAvoid, bk.TriangleHit, bk.HitDisp)
        y2 := uicore.DrawWrapped(dst, uicore.FaceSmall, rightHit, x, y+uicore.S(50), color.RGBA{210,230,255,255}, wrapRight, uicore.LineHSmallPx())
        rightDmg := fmt.Sprintf("[与ダメ内訳] 力:%d + 威力:%d + 相性:%+d - 守備合計:%d = %d", bk.AtkStr, bk.WpnMt, bk.TriangleMt, bk.DefTotal, bk.Dmg)
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, rightDmg, x, y2, color.RGBA{210,230,255,255}, wrapRight, uicore.LineHSmallPx())
    } else {
        _ = uicore.DrawWrapped(dst, uicore.FaceMain, "(反撃不可) 命中 -  与ダメ -  必殺 -", x, y, color.RGBA{180,180,180,255}, wrapRight, uicore.LineHMainPx())
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
	if len(u.Equip) > 0 {
		wep = u.Equip[0].Name
	}
    // 攻撃速度（一元化ロジック） - 小さめフォント + 行間確保
    as := adapter.AttackSpeedOf(weaponTable(), u)
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
    if !left {
        base = panelW/2 - uicore.S(140)
    }
    if base < uicore.S(420) { base = uicore.S(420) }
    if base > uicore.S(900) { base = uicore.S(900) }
    return base
}


// forecastBothWithTerrain は未使用のため削除予定（互換用）。

// explain 版
func forecastBothWithTerrainExplain(atk, def uicore.Unit, attT, defT gcore.Terrain) (gcore.ForecastResult, gcore.ForecastBreakdown, gcore.ForecastResult, gcore.ForecastBreakdown, bool, bool) {
    wt := weaponTable()
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

// 変換は adapter.UIToGame に集約。


func terrainLine(t gcore.Terrain) string {
    name := gcore.TerrainPresetName(t)
    return fmt.Sprintf("地形: %s (回避+%d 防御+%d 命中+%d)", name, t.Avoid, t.Def, t.Hit)
}

func weaponTypeOf(wt *model.WeaponTable, u uicore.Unit) string {
    if len(u.Equip) == 0 {
        return ""
    }
    if w, ok := wt.Find(u.Equip[0].Name); ok {
        return w.Type
    }
    return ""
}

func weaponTable() *model.WeaponTable {
    if wtShared != nil { return wtShared }
    // フォールバック: 直接JSON読込（初回のみ）。
    if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
        wtShared = wt
        return wtShared
    }
    return nil
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
// 末尾に「クリックまたはZ/Enterで閉じる」ガイドを表示します。
func DrawBattleLogOverlay(dst *ebiten.Image, logs []string) {
    if len(logs) == 0 {
        return
    }
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    // 背景ディマー
    vector.DrawFilledRect(dst, 0, 0, float32(sw), float32(sh), color.RGBA{0, 0, 0, 140}, false)
    // 中央パネル
    pw, ph := int(float32(sw)*0.7), 300
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    uicore.DrawFramedRect(dst, float32(px), float32(py), float32(pw), float32(ph))
    vector.DrawFilledRect(dst, float32(px), float32(py), float32(pw), float32(ph), color.RGBA{25, 30, 50, 230}, false)
    uicore.TextDraw(dst, "戦闘ログ", uicore.FaceMain, px+16, py+32, uicore.ColAccent)
    // ログ本文（下に新しいもの）
    maxLines := (ph - uicore.S(80)) / uicore.LineHSmallPx()
    start := 0
    if len(logs) > maxLines {
        start = len(logs) - maxLines
    }
    y := py + 58
    for i := start; i < len(logs); i++ {
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, logs[i], px+uicore.S(16), y, uicore.ColText, pw-uicore.S(32), uicore.LineHSmallPx())
        y += uicore.LineHSmallPx()
    }
    hint := "クリック または Z/Enter で閉じる"
    tw := int(font.MeasureString(uicore.FaceSmall, hint) >> 6)
    uicore.TextDraw(dst, hint, uicore.FaceSmall, px+(pw-tw)/2, py+ph-16, color.RGBA{210, 220, 240, 255})
}

// DrawBattleLogOverlayScroll は DrawBattleLogOverlay と同レイアウトで、
// 末尾からのオフセット（offset行）を指定してスクロール表示します。
// offset=0 で最新、正の値で過去方向へスクロールします。
func DrawBattleLogOverlayScroll(dst *ebiten.Image, logs []string, offset int) {
    if len(logs) == 0 { return }
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    // 背景ディマー
    vector.DrawFilledRect(dst, 0, 0, float32(sw), float32(sh), color.RGBA{0, 0, 0, 140}, false)
    // 中央パネル（既存と同寸）
    pw, ph := int(float32(sw)*0.7), 300
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    uicore.DrawFramedRect(dst, float32(px), float32(py), float32(pw), float32(ph))
    vector.DrawFilledRect(dst, float32(px), float32(py), float32(pw), float32(ph), color.RGBA{25, 30, 50, 230}, false)
    uicore.TextDraw(dst, "戦闘ログ", uicore.FaceMain, px+16, py+24, uicore.ColAccent)
    // スクロール
    maxLines := (ph - uicore.S(64)) / uicore.LineHSmallPx()
    if maxLines < 1 { maxLines = 1 }
    if offset < 0 { offset = 0 }
    if offset > len(logs)-1 { offset = len(logs)-1 }
    start := len(logs) - maxLines - offset
    if start < 0 { start = 0 }
    end := start + maxLines
    if end > len(logs) { end = len(logs) }
    y := py + 48
    for i := start; i < end; i++ {
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, logs[i], px+uicore.S(16), y, uicore.ColText, pw-uicore.S(32), uicore.LineHSmallPx())
        y += uicore.LineHSmallPx()
    }
    // ヒント
    hint := "上下/ホイールでスクロール・停止で閉じる"
    tw := int(font.MeasureString(uicore.FaceSmall, hint) >> 6)
    uicore.TextDraw(dst, hint, uicore.FaceSmall, px+(pw-tw)/2, py+ph-12, color.RGBA{210, 220, 240, 255})
}

// DrawBattleLogs は画面下部に戦闘ログを表示します。
func DrawBattleLogs(dst *ebiten.Image, logs []string) {
    if len(logs) == 0 {
        return
    }
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    w, h := sw-2*uicore.ListMargin-80, 160
    x := uicore.ListMargin + 40
    y := sh - uicore.ListMargin - h - 10
    uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
    vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), color.RGBA{25, 30, 50, 220}, false)
    uicore.TextDraw(dst, "戦闘ログ", uicore.FaceMain, x+12, y+24, uicore.ColAccent)
    maxLines := (h - uicore.S(40)) / uicore.LineHSmallPx()
    lineY := y + 40
    start := 0
    if len(logs) > maxLines {
        start = len(logs) - maxLines
    }
    for i := start; i < len(logs); i++ {
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, logs[i], x+uicore.S(12), lineY, uicore.ColText, w-uicore.S(24), uicore.LineHSmallPx())
        lineY += uicore.LineHSmallPx()
    }
}
