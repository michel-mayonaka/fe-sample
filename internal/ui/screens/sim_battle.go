package uiscreens

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "math/rand"
    "ui_sample/internal/adapter"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/ui/core"
    gcore "ui_sample/pkg/game"
)

// ボタンRectは widgets パッケージ側に統一しています。

// SimulateBattleCopy はコピーを用いた簡易模擬戦を行い、結果ログを返します（永続化なし）。
func SimulateBattleCopy(atk, def uicore.Unit, rng *rand.Rand) (uicore.Unit, uicore.Unit, []string) {
    a := atk
    d := def
    logs := []string{"模擬戦開始", fmt.Sprintf("%s vs %s", a.Name, d.Name)}
    // 3ラウンド上限 or どちらか撃破
    for r := 1; r <= 3 && a.HP > 0 && d.HP > 0; r++ {
        // 武器テーブル
        wt := weaponTable()
        if wt == nil {
            if wtf, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil { wtShared = wtf; wt = wtShared }
        }
        if wt == nil { break }
        ga := adapter.UIToGame(wt, a)
        gd := adapter.UIToGame(wt, d)
        // 攻撃
        logs = append(logs, fmt.Sprintf("[R%d] %s の攻撃", r, a.Name))
        ga2, gd2, line1 := gcore.ResolveRoundAt(ga, gd, gcore.Terrain{}, gcore.Terrain{}, rng)
        if line1 != "" { logs = append(logs, line1) }
        d.HP = gd2.S.HP
        if d.HP < 0 { d.HP = 0 }
        // 反撃
        dist := 1
        canCounter := gd2.W.RMin <= dist && dist <= gd2.W.RMax
        if d.HP > 0 && canCounter {
            logs = append(logs, fmt.Sprintf("[R%d] %s の反撃", r, d.Name))
            gd3, ga3, line2 := gcore.ResolveRoundAt(gd2, ga2, gcore.Terrain{}, gcore.Terrain{}, rng)
            if line2 != "" { logs = append(logs, line2) }
            a.HP = ga3.S.HP
            if a.HP < 0 { a.HP = 0 }
            ga2, gd2 = ga3, gd3
        } else {
            a.HP = ga2.S.HP
        }
        // 追撃
        if d.HP > 0 && gcore.DoubleAdvantage(ga, gd) {
            logs = append(logs, fmt.Sprintf("[R%d] %s の追撃", r, a.Name))
            ga4, gd4, line3 := gcore.ResolveRoundAt(ga2, gd2, gcore.Terrain{}, gcore.Terrain{}, rng)
            if line3 != "" { logs = append(logs, line3) }
            d.HP = gd4.S.HP
            if d.HP < 0 { d.HP = 0 }
            a.HP = ga4.S.HP
        } else if a.HP > 0 && canCounter && gcore.DoubleAdvantage(gd, ga) {
            logs = append(logs, fmt.Sprintf("[R%d] %s の追撃", r, d.Name))
            gd4, ga4, line4 := gcore.ResolveRoundAt(gd2, ga2, gcore.Terrain{}, gcore.Terrain{}, rng)
            if line4 != "" { logs = append(logs, line4) }
            a.HP = ga4.S.HP
            if a.HP < 0 { a.HP = 0 }
            d.HP = gd4.S.HP
        }
    }
    if a.HP <= 0 { logs = append(logs, fmt.Sprintf("%s は倒れた", a.Name)) }
    if d.HP <= 0 { logs = append(logs, fmt.Sprintf("%s は倒れた", d.Name)) }
    logs = append(logs, "模擬戦終了")
    return a, d, logs
}

func simulateHit(atk, def uicore.Unit, rng *rand.Rand) (uicore.Unit, uicore.Unit, string) {
    // 武器威力（共有テーブル）
    might := 0
    if len(atk.Equip) > 0 {
        if wt := weaponTable(); wt != nil {
            if w, ok := wt.Find(atk.Equip[0].Name); ok { might = w.Might }
        }
    }
    ga := gcore.Unit{ID: atk.ID, Name: atk.Name, Class: atk.Class, Lv: atk.Level,
        S: gcore.Stats{HP: atk.HP, Str: atk.Stats.Str, Skl: atk.Stats.Skl, Spd: atk.Stats.Spd, Lck: atk.Stats.Lck, Def: atk.Stats.Def, Res: atk.Stats.Res, Mov: atk.Stats.Mov, Bld: atk.Stats.Bld},
        W: gcore.Weapon{MT: might}}
    gd := gcore.Unit{ID: def.ID, Name: def.Name, Class: def.Class, Lv: def.Level,
        S: gcore.Stats{HP: def.HP, Str: def.Stats.Str, Skl: def.Stats.Skl, Spd: def.Stats.Spd, Lck: def.Stats.Lck, Def: def.Stats.Def, Res: def.Stats.Res, Mov: def.Stats.Mov, Bld: def.Stats.Bld}}
	_, gd, line := gcore.ResolveRound(ga, gd, rng)
	def.HP = gd.S.HP
	if def.HP < 0 {
		def.HP = 0
	}
	if line == "" {
		line = fmt.Sprintf("命中! 0ダメージ (HP %d/%d)", def.HP, def.HPMax)
	}
	return atk, def, line
}

// DrawSimulationBattle は模擬戦の結果を画面に描画します。
func DrawSimulationBattle(dst *ebiten.Image, atk, def uicore.Unit, logs []string) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := uicore.ListMarginPx()
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
	// 左右ユニット
    leftX := lm + uicore.S(40)
    rightX := sw - lm - uicore.S(560)
    topY := lm + uicore.S(80)
	drawSide(dst, atk, leftX, topY)
	drawSide(dst, def, rightX, topY)
	// ログパネル
    lw, lh := uicore.S(800), uicore.S(260)
    lx := (sw - lw) / 2
    ly := sh - lm - lh - uicore.S(20)
	uicore.DrawFramedRect(dst, float32(lx), float32(ly), float32(lw), float32(lh))
	vector.DrawFilledRect(dst, float32(lx), float32(ly), float32(lw), float32(lh), color.RGBA{25, 30, 50, 220}, false)
	y := ly + 36
    for i := len(logs) - 1; i >= 0 && y < ly+lh-uicore.S(10); i-- { // 下に新しいログを表示
        text.Draw(dst, logs[i], uicore.FaceSmall, lx+16, y, uicore.ColText)
        y += uicore.S(22)
    }
    text.Draw(dst, "模擬戦", uicore.FaceTitle, sw/2-uicore.S(60), lm+uicore.S(56), uicore.ColAccent)
}

func drawSide(dst *ebiten.Image, u uicore.Unit, x, y int) {
    sz := uicore.S(320)
    uicore.DrawFramedRect(dst, float32(x), float32(y), float32(sz), float32(sz))
    if u.Portrait != nil {
        uicore.DrawPortrait(dst, u.Portrait, float32(x), float32(y), float32(sz), float32(sz))
    }
    text.Draw(dst, u.Name, uicore.FaceTitle, x, y-uicore.S(16), uicore.ColText)
    text.Draw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceMain, x, y+uicore.S(350), uicore.ColAccent)
    text.Draw(dst, uicore.Itoa(u.HP)+"/"+uicore.Itoa(u.HPMax), uicore.FaceMain, x, y+uicore.S(384), uicore.ColText)
    uicore.DrawHPBar(dst, x, y+uicore.S(390), sz, uicore.S(14), u.HP, u.HPMax)
    // 攻撃速度（一元化ロジック） - 小さめフォント + 行間確保
    as := adapter.AttackSpeedOf(weaponTable(), u)
    text.Draw(dst, fmt.Sprintf("攻撃速度 %d", as), uicore.FaceSmall, x, y+uicore.S(410), uicore.ColText)
}
