package uiscreens

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "ui_sample/internal/model"
    "ui_sample/internal/ui/core"
    gcore "ui_sample/pkg/game"
)

func BattleStartButtonRect(sw, sh int) (x, y, w, h int) {
	w, h = 240, 60
	x = (sw - w) / 2
	y = sh - uicore.ListMargin - h
	return
}

// DrawBattle は後方互換（平地扱い）。
func DrawBattle(dst *ebiten.Image, attacker, defender uicore.Unit) {
    DrawBattleWithTerrain(dst, attacker, defender, gcore.Terrain{}, gcore.Terrain{})
}

// DrawBattleWithTerrain は左右の地形を指定してプレビューを描画します。
func DrawBattleWithTerrain(dst *ebiten.Image, attacker, defender uicore.Unit, attT, defT gcore.Terrain) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    uicore.DrawPanel(dst, uicore.ListMargin, uicore.ListMargin, float32(sw-2*uicore.ListMargin), float32(sh-2*uicore.ListMargin))
    leftX := uicore.ListMargin + 40
    rightX := sw - uicore.ListMargin - 560
    topY := uicore.ListMargin + 80
    drawBattleSide(dst, attacker, leftX, topY)
    drawBattleSide(dst, defender, rightX, topY)
    text.Draw(dst, "戦闘プレビュー", uicore.FaceTitle, sw/2-120, uicore.ListMargin+56, uicore.ColAccent)
    bx, by, bw, bh := BattleStartButtonRect(sw, sh)
    uicore.DrawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
    vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), color.RGBA{110, 90, 40, 255}, false)
    text.Draw(dst, "戦闘開始", uicore.FaceMain, bx+70, by+38, uicore.ColText)

    // 予測値表示（/pkg/game.ForecastAt）
    if frAtk, frAtkBk, frDef, frDefBk, canCounter, ok := forecastBothWithTerrainExplain(attacker, defender, attT, defT); ok {
        // 左側（攻撃側の予測）
        ax := leftX
        ay := topY + 460
        text.Draw(dst, fmt.Sprintf("命中 %d%%  与ダメ %d  必殺 %d%%", frAtk.HitDisp, frAtk.Dmg, frAtk.Crit), uicore.FaceMain, ax, ay, uicore.ColText)
        // 右側（反撃側の予測）
        dx := rightX
        dy := topY + 460
        if canCounter {
            text.Draw(dst, fmt.Sprintf("(反撃) 命中 %d%%  与ダメ %d  必殺 %d%%", frDef.HitDisp, frDef.Dmg, frDef.Crit), uicore.FaceMain, dx, dy, uicore.ColText)
        } else {
            text.Draw(dst, "(反撃不可) 命中 -  与ダメ -  必殺 -", uicore.FaceMain, dx, dy, color.RGBA{180,180,180,255})
        }
        // 地形表示
        text.Draw(dst, terrainLabel(attT), uicore.FaceSmall, ax, ay+26, color.RGBA{200,220,255,255})
        text.Draw(dst, terrainLabel(defT), uicore.FaceSmall, dx, dy+26, color.RGBA{200,220,255,255})
        // 内訳（簡潔表示）
        text.Draw(dst, fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
            frAtkBk.WeapHit, frAtkBk.Skl2, frAtkBk.LckHalf, frAtkBk.AttTileHit,
            frAtkBk.DefSpd2, frAtkBk.DefLck, frAtkBk.DefTileAvoid, frAtkBk.TriangleHit, frAtkBk.HitDisp,
        ), uicore.FaceSmall, ax, ay+50, color.RGBA{210,230,255,255})
        text.Draw(dst, fmt.Sprintf("[与ダメ内訳] 力:%d + 威力:%d + 相性:%+d - 守備合計:%d = %d",
            frAtkBk.AtkStr, frAtkBk.WpnMt, frAtkBk.TriangleMt, frAtkBk.DefTotal, frAtkBk.Dmg,
        ), uicore.FaceSmall, ax, ay+70, color.RGBA{210,230,255,255})
        if canCounter {
            text.Draw(dst, fmt.Sprintf("[命中内訳] 武器H%d + 技×2:%d + 幸/2:%d + 地命中:%d - (速×2:%d + 幸:%d + 地回避:%d) + 相性:%+d = %d",
                frDefBk.WeapHit, frDefBk.Skl2, frDefBk.LckHalf, frDefBk.AttTileHit,
                frDefBk.DefSpd2, frDefBk.DefLck, frDefBk.DefTileAvoid, frDefBk.TriangleHit, frDefBk.HitDisp,
            ), uicore.FaceSmall, dx, dy+50, color.RGBA{210,230,255,255})
            text.Draw(dst, fmt.Sprintf("[与ダメ内訳] 力:%d + 威力:%d + 相性:%+d - 守備合計:%d = %d",
                frDefBk.AtkStr, frDefBk.WpnMt, frDefBk.TriangleMt, frDefBk.DefTotal, frDefBk.Dmg,
            ), uicore.FaceSmall, dx, dy+70, color.RGBA{210,230,255,255})
        }
    }

    // 相性ラベル
    if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
        aType := weaponTypeOf(wt, attacker)
        dType := weaponTypeOf(wt, defender)
        lbl, col := triangleLabel(aType, dType)
        text.Draw(dst, "相性: "+lbl, uicore.FaceMain, leftX, topY+490, col)
    }
    // ヘルプ: 地形切替
    text.Draw(dst, "[地形切替] 攻: 1=平地 2=森 3=砦 / 防: Shift+1/2/3", uicore.FaceSmall, leftX, topY+516, color.RGBA{190,200,210,255})
}

func drawBattleSide(dst *ebiten.Image, u uicore.Unit, x, y int) {
	uicore.DrawFramedRect(dst, float32(x), float32(y), 320, 320)
	if u.Portrait != nil {
		uicore.DrawPortrait(dst, u.Portrait, float32(x), float32(y), 320, 320)
	}
	text.Draw(dst, u.Name, uicore.FaceTitle, x, y-16, uicore.ColText)
	text.Draw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceMain, x, y+350, uicore.ColAccent)
	text.Draw(dst, uicore.Itoa(u.HP)+"/"+uicore.Itoa(u.HPMax), uicore.FaceMain, x, y+384, uicore.ColText)
	uicore.DrawHPBar(dst, x, y+390, 320, 14, u.HP, u.HPMax)
	wep := "-"
	if len(u.Equip) > 0 {
		wep = u.Equip[0].Name
	}
    text.Draw(dst, "武器: "+wep, uicore.FaceMain, x, y+420, uicore.ColText)
}

// forecastBoth は UI ユニット2体から /pkg/game の予測結果を返します。
func forecastBoth(atk, def uicore.Unit) (gcore.ForecastResult, gcore.ForecastResult, bool, gcore.Terrain, gcore.Terrain, bool) {
    wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    if err != nil {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, gcore.Terrain{}, gcore.Terrain{}, false
    }
    ga, ok := toGameUnit(wt, atk)
    if !ok {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, gcore.Terrain{}, gcore.Terrain{}, false
    }
    gd, ok := toGameUnit(wt, def)
    if !ok {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, gcore.Terrain{}, gcore.Terrain{}, false
    }
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
    wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    if err != nil {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, false
    }
    ga, ok := toGameUnit(wt, atk)
    if !ok {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, false
    }
    gd, ok := toGameUnit(wt, def)
    if !ok {
        return gcore.ForecastResult{}, gcore.ForecastResult{}, false, false
    }
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
    wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    if err != nil {
        return gcore.ForecastResult{}, gcore.ForecastBreakdown{}, gcore.ForecastResult{}, gcore.ForecastBreakdown{}, false, false
    }
    ga, ok := toGameUnit(wt, atk)
    if !ok {
        return gcore.ForecastResult{}, gcore.ForecastBreakdown{}, gcore.ForecastResult{}, gcore.ForecastBreakdown{}, false, false
    }
    gd, ok := toGameUnit(wt, def)
    if !ok {
        return gcore.ForecastResult{}, gcore.ForecastBreakdown{}, gcore.ForecastResult{}, gcore.ForecastBreakdown{}, false, false
    }
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

func toGameUnit(wt *model.WeaponTable, u uicore.Unit) (gcore.Unit, bool) {
    var w model.Weapon
    if len(u.Equip) > 0 {
        if ww, ok := wt.Find(u.Equip[0].Name); ok {
            w = ww
        }
    }
    gu := gcore.Unit{
        ID: u.ID, Name: u.Name, Class: u.Class, Lv: u.Level,
        S: gcore.Stats{HP: u.HP, Str: u.Stats.Str, Skl: u.Stats.Skl, Spd: u.Stats.Spd, Lck: u.Stats.Lck, Def: u.Stats.Def, Res: u.Stats.Res, Mov: u.Stats.Mov},
        W: gcore.Weapon{MT: w.Might, Hit: w.Hit, Crit: w.Crit, Wt: w.Weight, RMin: w.RangeMin, RMax: w.RangeMax, Type: w.Type},
    }
    return gu, true
}

func defaultTerrains() (gcore.Terrain, gcore.Terrain) {
    // MVP: 平地（回避0/防御0/命中0）を両者に適用。後続でUI/マップに接続。
    return gcore.Terrain{}, gcore.Terrain{}
}

func terrainLabel(t gcore.Terrain) string {
    name := "平地"
    if t.Avoid == 20 && t.Def == 1 {
        name = "森"
    } else if t.Avoid == 15 && t.Def == 2 {
        name = "砦"
    }
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

func triangleLabel(attType, defType string) (string, color.Color) {
    // Sword > Axe > Lance > Sword
    switch attType {
    case "Sword":
        if defType == "Axe" {
            return "有利", uicore.ColAccent
        } else if defType == "Lance" {
            return "不利", color.RGBA{255, 170, 70, 255}
        }
    case "Axe":
        if defType == "Lance" {
            return "有利", uicore.ColAccent
        } else if defType == "Sword" {
            return "不利", color.RGBA{255, 170, 70, 255}
        }
    case "Lance":
        if defType == "Sword" {
            return "有利", uicore.ColAccent
        } else if defType == "Axe" {
            return "不利", color.RGBA{255, 170, 70, 255}
        }
    }
    return "中立", color.RGBA{180, 180, 180, 255}
}
