package draw

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"image/color"
	"ui_sample/internal/adapter"
	gdata "ui_sample/internal/game/data"
	uicore "ui_sample/internal/game/service/ui"
	uilayout "ui_sample/internal/game/ui/layout"
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
	leftX := lm + uicore.SimPreviewLeftXPadPx()
	rightX := sw - lm - uicore.SimPreviewRightXInsetPx()
	topY := lm + uicore.SimPreviewTopYFromMarginPx()
	drawBattleSide(dst, attacker, leftX, topY)
	drawBattleSide(dst, defender, rightX, topY)
	// 開始ボタン
	drawStartButton(dst, sw, sh, startEnabled)

	// 予測値表示
	if frAtk, frAtkBk, frDef, frDefBk, canCounter, ok := forecastBothWithTerrainExplain(attacker, defender, attT, defT); ok {
		ax, ay := leftX, topY+uicore.S(460)
		dx, dy := rightX, topY+uicore.S(460)
		drawForecastLeft(dst, ax, ay, attacker, defender, frAtk, frAtkBk)
		drawForecastRight(dst, dx, dy, sw, lm, canCounter, frDef, frDefBk)
		// 地形ラベル
		drawTerrainLabels(dst, sw, lm, ax, ay, dx, dy, attT, defT)
	}

	// ヘルプ: 地形切替（下部寄せ）
	uicore.TextDraw(dst, "[地形切替] 攻: 1=平地 2=森 3=砦 / 防: Shift+1/2/3", uicore.FaceSmall, leftX, sh-uicore.ListMarginPx()-uicore.S(190), color.RGBA{190, 200, 210, 255})
}

func drawBattleHeader(dst *ebiten.Image, sw, sh, lm int) {
	uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
	uicore.TextDraw(dst, "戦闘プレビュー", uicore.FaceTitle, sw/2-uicore.SimTitleXOffsetFromCenterPx(), uicore.ListMarginPx()+uicore.SimTitleYOffsetPx(), uicore.ColAccent)
}

func drawStartButton(dst *ebiten.Image, sw, sh int, enabled bool) {
	bx, by, bw, bh := uilayout.BattleStartButtonRect(sw, sh)
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

// 左側の予測（攻撃側）
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
	_ = uicore.DrawWrapped(dst, uicore.FaceSmall, leftHitLine, x, y+uicore.S(26), uicore.ColText, wrap, uicore.LineHSmallPx())
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

func drawTerrainLabels(dst *ebiten.Image, sw, lm, ax, ay, dx, dy int, attT, defT gcore.Terrain) {
	// 地形ラベル
	as := terrainLine(attT)
	ds := terrainLine(defT)
	uicore.TextDraw(dst, as, uicore.FaceSmall, ax, ay+uicore.S(90), uicore.ColText)
	tw := int(font.MeasureString(uicore.FaceSmall, ds) >> 6)
	uicore.TextDraw(dst, ds, uicore.FaceSmall, sw-lm-uicore.S(560)+uicore.S(540)-tw, dy+uicore.S(90), uicore.ColText)
}

func drawBattleSide(dst *ebiten.Image, u uicore.Unit, x, y int) {
	w := uicore.SimPreviewCardWPx()
	h := uicore.SimPreviewCardHPx()
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), color.RGBA{25, 30, 50, 230}, false)
	// 基本情報
	uicore.TextDraw(dst, u.Name, uicore.FaceMain, x+uicore.SimPreviewNameOffsetXPx(), y+uicore.SimPreviewNameOffsetYPx(), uicore.ColAccent)
	uicore.TextDraw(dst, u.Class, uicore.FaceSmall, x+uicore.SimPreviewNameOffsetXPx(), y+uicore.SimPreviewClassOffsetYPx(), uicore.ColText)
	// ポトレ
	ps := uicore.SimPreviewPortraitSizePx()
	px := x + uicore.SimPreviewCardInnerPadPx()
	py := y + (uicore.SimPreviewCardHPx()-ps)/2
	uicore.DrawFramedRect(dst, float32(px), float32(py), float32(ps), float32(ps))
	if u.Portrait != nil {
		uicore.DrawPortrait(dst, u.Portrait, float32(px), float32(py), float32(ps), float32(ps))
	} else {
		uicore.DrawPortraitPlaceholder(dst, float32(px), float32(py), float32(ps), float32(ps))
	}
	// 主武器など
	wep := "-"
	if len(u.Equip) > 0 && u.Equip[0].Name != "" {
		wep = u.Equip[0].Name
	}
	uicore.TextDraw(dst, "HP: ", uicore.FaceSmall, x+uicore.SimPreviewHPLabelXPx(), y+uicore.SimPreviewHPLabelYPx(), uicore.ColText)
	uicore.DrawHPBar(dst, x+uicore.SimPreviewHPBarXPx(), y+uicore.SimPreviewHPBarYPx(), uicore.SimPreviewHPBarWPx(), uicore.SimPreviewHPBarHPx(), u.HP, u.HPMax)
	// 攻撃速度
	var wt *model.WeaponTable
	if p := gdata.Provider(); p != nil {
		wt = p.WeaponsTable()
	}
	as := adapter.AttackSpeedOf(wt, u)
	lineY := y + uicore.SimPreviewLineYPx()
	uicore.TextDraw(dst, fmt.Sprintf("攻撃速度 %d", as), uicore.FaceSmall, x, lineY, uicore.ColText)
	lineY += uicore.LineHSmallPx()
	uicore.TextDraw(dst, "武器: "+wep, uicore.FaceSmall, x, lineY, uicore.ColText)
}

// dynamicWrap は左右列の折返し幅をウィンドウ幅とマージンから決めます。
func dynamicWrap(sw, lm int, left bool) int {
	panelW := sw - 2*lm
	base := panelW/2 - uicore.S(100)
	if !left {
		base = panelW/2 - uicore.S(140)
	}
	if base < uicore.SimPreviewBaseMinPx() {
		base = uicore.SimPreviewBaseMinPx()
	}
	if base > uicore.SimPreviewBaseMaxPx() {
		base = uicore.SimPreviewBaseMaxPx()
	}
	return base
}

// explain 版
func forecastBothWithTerrainExplain(atk, def uicore.Unit, attT, defT gcore.Terrain) (gcore.ForecastResult, gcore.ForecastBreakdown, gcore.ForecastResult, gcore.ForecastBreakdown, bool, bool) {
	var wt *model.WeaponTable
	if p := gdata.Provider(); p != nil {
		wt = p.WeaponsTable()
	}
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
	if len(u.Equip) == 0 {
		return ""
	}
	if w, ok := wt.Find(u.Equip[0].Name); ok {
		return w.Type
	}
	return ""
}

// DrawBattleLogOverlay は全画面を半透明で覆い、中央にログを表示します。
func DrawBattleLogOverlay(dst *ebiten.Image, logs []string) {
	if len(logs) == 0 {
		return
	}
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	vector.DrawFilledRect(dst, 0, 0, float32(sw), float32(sh), color.RGBA{0, 0, 0, 140}, false)
	pw, ph := int(float32(sw)*0.7), 300
	px := (sw - pw) / 2
	py := (sh - ph) / 2
	uicore.DrawFramedRect(dst, float32(px), float32(py), float32(pw), float32(ph))
	vector.DrawFilledRect(dst, float32(px), float32(py), float32(pw), float32(ph), color.RGBA{25, 30, 50, 230}, false)
	uicore.TextDraw(dst, "戦闘ログ", uicore.FaceMain, px+16, py+32, uicore.ColAccent)
	maxLines := (ph - uicore.SimPreviewWrapPadPx()) / uicore.LineHSmallPx()
	start := 0
	if len(logs) > maxLines {
		start = len(logs) - maxLines
	}
	y := py + 58
	for i := start; i < len(logs); i++ {
		_ = uicore.DrawWrapped(dst, uicore.FaceSmall, logs[i], px+uicore.SimPreviewLogPadXPx(), y+uicore.SimPreviewLogPadYPx(), uicore.ColText, pw-2*uicore.SimPreviewLogPadXPx(), uicore.LineHSmallPx())
		y += uicore.LineHSmallPx()
	}
	hint := "クリック または Z/Enter で閉じる"
	tw := int(font.MeasureString(uicore.FaceSmall, hint) >> 6)
	uicore.TextDraw(dst, hint, uicore.FaceSmall, px+(pw-tw)/2, py+ph-16, color.RGBA{210, 220, 240, 255})
}
