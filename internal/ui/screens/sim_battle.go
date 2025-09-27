package uiscreens

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"ui_sample/internal/model"
	uicore "ui_sample/internal/ui/core"
)

// ボタンRectは widgets パッケージ側に統一しています。

// SimulateBattleCopy はコピーを用いた簡易模擬戦を行い、結果ログを返します（永続化なし）。
func SimulateBattleCopy(atk, def uicore.Unit, rng *rand.Rand) (uicore.Unit, uicore.Unit, []string) {
	a := atk
	d := def
	logs := []string{"模擬戦開始", fmt.Sprintf("%s vs %s", a.Name, d.Name)}
	// 3ラウンド上限 or どちらか撃破
	for r := 1; r <= 3 && a.HP > 0 && d.HP > 0; r++ {
		logs = append(logs, fmt.Sprintf("[R%d] %s の攻撃", r, a.Name))
		a, d, line := simulateHit(a, d, rng)
		logs = append(logs, line)
		if d.HP <= 0 {
			break
		}
		logs = append(logs, fmt.Sprintf("[R%d] %s の反撃", r, d.Name))
		var line2 string
		d, a, line2 = simulateHit(d, a, rng)
		logs = append(logs, line2)
		_ = a // use after assign to satisfy ineffassign
	}
	if a.HP <= 0 {
		logs = append(logs, fmt.Sprintf("%s は倒れた", a.Name))
	}
	if d.HP <= 0 {
		logs = append(logs, fmt.Sprintf("%s は倒れた", d.Name))
	}
	logs = append(logs, "模擬戦終了")
	return a, d, logs
}

func simulateHit(atk, def uicore.Unit, rng *rand.Rand) (uicore.Unit, uicore.Unit, string) {
	// 武器威力
	might := 0
	if len(atk.Equip) > 0 {
		if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
			if w, ok := wt.Find(atk.Equip[0].Name); ok {
				might = w.Might
			}
		}
	}
	hit := 80 + atk.Stats.Skl*2 + atk.Stats.Lck/2 - (def.Stats.Spd*2 + def.Stats.Lck)
	if hit < 0 {
		hit = 0
	}
	if hit > 100 {
		hit = 100
	}
	if rng.Intn(100) < hit {
		dmg := atk.Stats.Str + might - def.Stats.Def
		if dmg < 0 {
			dmg = 0
		}
		def.HP -= dmg
		if def.HP < 0 {
			def.HP = 0
		}
		return atk, def, fmt.Sprintf("命中! %dダメージ (HP %d/%d)", dmg, def.HP, def.HPMax)
	}
	return atk, def, "ミス!"
}

// DrawSimulationBattle は模擬戦の結果を画面に描画します。
func DrawSimulationBattle(dst *ebiten.Image, atk, def uicore.Unit, logs []string) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	uicore.DrawPanel(dst, uicore.ListMargin, uicore.ListMargin, float32(sw-2*uicore.ListMargin), float32(sh-2*uicore.ListMargin))
	// 左右ユニット
	leftX := uicore.ListMargin + 40
	rightX := sw - uicore.ListMargin - 560
	topY := uicore.ListMargin + 80
	drawSide(dst, atk, leftX, topY)
	drawSide(dst, def, rightX, topY)
	// ログパネル
	lw, lh := 800, 260
	lx := (sw - lw) / 2
	ly := sh - uicore.ListMargin - lh - 20
	uicore.DrawFramedRect(dst, float32(lx), float32(ly), float32(lw), float32(lh))
	vector.DrawFilledRect(dst, float32(lx), float32(ly), float32(lw), float32(lh), color.RGBA{25, 30, 50, 220}, false)
	y := ly + 36
	for i := len(logs) - 1; i >= 0 && y < ly+lh-10; i-- { // 下に新しいログを表示
		text.Draw(dst, logs[i], uicore.FaceSmall, lx+16, y, uicore.ColText)
		y += 22
	}
	text.Draw(dst, "模擬戦", uicore.FaceTitle, sw/2-60, uicore.ListMargin+56, uicore.ColAccent)
}

func drawSide(dst *ebiten.Image, u uicore.Unit, x, y int) {
	uicore.DrawFramedRect(dst, float32(x), float32(y), 320, 320)
	if u.Portrait != nil {
		uicore.DrawPortrait(dst, u.Portrait, float32(x), float32(y), 320, 320)
	}
	text.Draw(dst, u.Name, uicore.FaceTitle, x, y-16, uicore.ColText)
	text.Draw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceMain, x, y+350, uicore.ColAccent)
	text.Draw(dst, uicore.Itoa(u.HP)+"/"+uicore.Itoa(u.HPMax), uicore.FaceMain, x, y+384, uicore.ColText)
	uicore.DrawHPBar(dst, x, y+390, 320, 14, u.HP, u.HPMax)
}
