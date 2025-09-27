package uiscreens

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "ui_sample/internal/adapter"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/ui/core"
    gcore "ui_sample/pkg/game"
)

// 共有武器テーブル（Repo注入で設定）。未設定時は初回アクセスで読み込みキャッシュ。
var wtShared *model.WeaponTable

func SetWeaponTable(wt *model.WeaponTable) { wtShared = wt }

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
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    leftX := lm + uicore.S(40)
    rightX := sw - lm - uicore.S(560)
    topY := lm + uicore.S(80)
    drawBattleSide(dst, attacker, leftX, topY)
    drawBattleSide(dst, defender, rightX, topY)
    text.Draw(dst, "戦闘プレビュー", uicore.FaceTitle, sw/2-uicore.S(120), uicore.ListMarginPx()+uicore.S(56), uicore.ColAccent)
    bx, by, bw, bh := BattleStartButtonRect(sw, sh)
    uicore.DrawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
    fill := color.RGBA{110, 90, 40, 255}
    labelCol := uicore.ColText
    if !startEnabled {
        fill = color.RGBA{80, 80, 80, 200}
        labelCol = color.RGBA{200, 200, 200, 180}
    }
    vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), fill, false)
    text.Draw(dst, "戦闘開始", uicore.FaceMain, bx+uicore.S(70), by+uicore.S(38), labelCol)

    // 予測値表示（/pkg/game.ForecastAt）
    if frAtk, frAtkBk, frDef, frDefBk, canCounter, ok := forecastBothWithTerrainExplain(attacker, defender, attT, defT); ok {
        // 左側（攻撃側の予測）
        ax := leftX
        ay := topY + uicore.S(460)
        baseLine := fmt.Sprintf("命中 %d%%  与ダメ %d  必殺 %d%%", frAtk.HitDisp, frAtk.Dmg, frAtk.Crit)
        text.Draw(dst, baseLine, uicore.FaceMain, ax, ay, uicore.ColText)
        // 相性を同一行の末尾に配置
        if wt := weaponTable(); wt != nil {
            aType := weaponTypeOf(wt, attacker)
            dType := weaponTypeOf(wt, defender)
            lbl, col := triangleRelationLabelColor(gcore.TriangleRelationOf(aType, dType))
            w := text.BoundString(uicore.FaceMain, baseLine).Dx()
            text.Draw(dst, "  相性: ", uicore.FaceMain, ax+w+8, ay, uicore.ColText)
            w2 := text.BoundString(uicore.FaceMain, "  相性: ").Dx()
            text.Draw(dst, lbl, uicore.FaceMain, ax+w+8+w2, ay, col)
        }
        // 右側（反撃側の予測）
        dx := rightX
        dy := topY + uicore.S(460)
        // 折返し幅（右列）
        wrapRight := dynamicWrap(sw, lm, false)
        if canCounter {
            baseLineR := fmt.Sprintf("(反撃) 命中 %d%%  与ダメ %d  必殺 %d%%", frDef.HitDisp, frDef.Dmg, frDef.Crit)
            _ = uicore.DrawWrapped(dst, uicore.FaceMain, baseLineR, dx, dy, uicore.ColText, wrapRight, uicore.LineHMainPx())
        } else {
            _ = uicore.DrawWrapped(dst, uicore.FaceMain, "(反撃不可) 命中 -  与ダメ -  必殺 -", dx, dy, color.RGBA{180,180,180,255}, wrapRight, uicore.LineHMainPx())
        }
        // 地形表示
        wrapLeft := dynamicWrap(sw, lm, true)
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, terrainLine(attT), ax, ay+uicore.S(26), color.RGBA{200,220,255,255}, wrapLeft, uicore.LineHSmallPx())
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, terrainLine(defT), dx, dy+uicore.S(26), color.RGBA{200,220,255,255}, wrapRight, uicore.LineHSmallPx())
        // 内訳（簡潔表示）: 折り返し描画
        leftHitLine := fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
            frAtkBk.WeapHit, frAtkBk.Skl2, frAtkBk.LckHalf, frAtkBk.AttTileHit,
            frAtkBk.DefSpd2, frAtkBk.DefLck, frAtkBk.DefTileAvoid, frAtkBk.TriangleHit, frAtkBk.HitDisp)
        y2 := uicore.DrawWrapped(dst, uicore.FaceSmall, leftHitLine, ax, ay+uicore.S(50), color.RGBA{210,230,255,255}, wrapLeft, uicore.LineHSmallPx())
        leftDmgLine := fmt.Sprintf("[与ダメ内訳] 力:%d + 威力:%d + 相性:%+d - 守備合計:%d = %d",
            frAtkBk.AtkStr, frAtkBk.WpnMt, frAtkBk.TriangleMt, frAtkBk.DefTotal, frAtkBk.Dmg)
        _ = uicore.DrawWrapped(dst, uicore.FaceSmall, leftDmgLine, ax, y2, color.RGBA{210,230,255,255}, wrapLeft, uicore.LineHSmallPx())
        if canCounter {
            rightHit := fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
                frDefBk.WeapHit, frDefBk.Skl2, frDefBk.LckHalf, frDefBk.AttTileHit,
                frDefBk.DefSpd2, frDefBk.DefLck, frDefBk.DefTileAvoid, frDefBk.TriangleHit, frDefBk.HitDisp)
            yR := uicore.DrawWrapped(dst, uicore.FaceSmall, rightHit, dx, dy+uicore.S(50), color.RGBA{210,230,255,255}, wrapRight, uicore.LineHSmallPx())
            rightDmg := fmt.Sprintf("[与ダメ内訳] 力:%d + 威力:%d + 相性:%+d - 守備合計:%d = %d",
                frDefBk.AtkStr, frDefBk.WpnMt, frDefBk.TriangleMt, frDefBk.DefTotal, frDefBk.Dmg)
            _ = uicore.DrawWrapped(dst, uicore.FaceSmall, rightDmg, dx, yR, color.RGBA{210,230,255,255}, wrapRight, uicore.LineHSmallPx())
        }
    }

    // ヘルプ: 地形切替（下部寄せ）
    text.Draw(dst, "[地形切替] 攻: 1=平地 2=森 3=砦 / 防: Shift+1/2/3", uicore.FaceSmall, leftX, sh-uicore.ListMarginPx()-uicore.S(190), color.RGBA{190,200,210,255})
}

func drawBattleSide(dst *ebiten.Image, u uicore.Unit, x, y int) {
    sz := uicore.S(320)
    uicore.DrawFramedRect(dst, float32(x), float32(y), float32(sz), float32(sz))
    if u.Portrait != nil {
        uicore.DrawPortrait(dst, u.Portrait, float32(x), float32(y), float32(sz), float32(sz))
    }
    text.Draw(dst, u.Name, uicore.FaceTitle, x, y-uicore.S(16), uicore.ColText)
    text.Draw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceMain, x, y+uicore.S(350), uicore.ColAccent)
    text.Draw(dst, uicore.Itoa(u.HP)+"/"+uicore.Itoa(u.HPMax), uicore.FaceMain, x, y+uicore.S(384), uicore.ColText)
    uicore.DrawHPBar(dst, x, y+uicore.S(390), sz, uicore.S(14), u.HP, u.HPMax)
	wep := "-"
	if len(u.Equip) > 0 {
		wep = u.Equip[0].Name
	}
    // 攻撃速度（一元化ロジック） - 小さめフォント + 行間確保
    as := adapter.AttackSpeedOf(weaponTable(), u)
    lineY := y + uicore.S(410)
    text.Draw(dst, fmt.Sprintf("攻撃速度 %d", as), uicore.FaceSmall, x, lineY, uicore.ColText)
    lineY += uicore.LineHSmallPx()
    text.Draw(dst, "武器: "+wep, uicore.FaceSmall, x, lineY, uicore.ColText)
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

// forecastBoth は UI ユニット2体から /pkg/game の予測結果を返します。
func forecastBoth(atk, def uicore.Unit) (gcore.ForecastResult, gcore.ForecastResult, bool, gcore.Terrain, gcore.Terrain, bool) {
    wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    if err != nil {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, gcore.Terrain{}, gcore.Terrain{}, false
    }
    ga := adapter.UIToGame(wt, atk)
    gd := adapter.UIToGame(wt, def)
    // 射程（暫定：距離1固定）。
    dist := 1
    canCounter := gd.W.RMin <= dist && dist <= gd.W.RMax
    // 地形（暫定: 平地固定）
    attT, defT := defaultTerrains()
    frDef := gcore.ForecastResult{}
    if canCounter {
        frDef = gcore.ForecastAt(gd, ga, defT, attT)
    }
    return gcore.ForecastAt(ga, gd, attT, defT), frDef, canCounter, attT, defT, true
}

// forecastBothWithTerrain は指定地形で予測します。
func forecastBothWithTerrain(atk, def uicore.Unit, attT, defT gcore.Terrain) (gcore.ForecastResult, gcore.ForecastResult, bool, bool) {
    wt := weaponTable()
    if wt == nil {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, false
    }
    ga := adapter.UIToGame(wt, atk)
    gd := adapter.UIToGame(wt, def)
    dist := 1
    canCounter := gd.W.RMin <= dist && dist <= gd.W.RMax
    frDef := gcore.ForecastResult{}
    if canCounter {
        frDef = gcore.ForecastAt(gd, ga, defT, attT)
    }
    return gcore.ForecastAt(ga, gd, attT, defT), frDef, canCounter, true
}

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

func defaultTerrains() (gcore.Terrain, gcore.Terrain) {
    // MVP: 平地（回避0/防御0/命中0）を両者に適用。後続でUI/マップに接続。
    return gcore.Terrain{}, gcore.Terrain{}
}

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
    text.Draw(dst, "戦闘ログ", uicore.FaceMain, px+16, py+32, uicore.ColAccent)
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
    tw := text.BoundString(uicore.FaceSmall, hint).Dx()
    text.Draw(dst, hint, uicore.FaceSmall, px+(pw-tw)/2, py+ph-16, color.RGBA{210, 220, 240, 255})
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
    text.Draw(dst, "戦闘ログ", uicore.FaceMain, px+16, py+24, uicore.ColAccent)
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
    tw := text.BoundString(uicore.FaceSmall, hint).Dx()
    text.Draw(dst, hint, uicore.FaceSmall, px+(pw-tw)/2, py+ph-12, color.RGBA{210, 220, 240, 255})
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
    text.Draw(dst, "戦闘ログ", uicore.FaceMain, x+12, y+24, uicore.ColAccent)
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
