// Package main は Ebiten を用いた FE 風ステータスUIサンプルの
// エントリポイントを提供します。
package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"ui_sample/internal/game"
	"ui_sample/internal/model"
	"ui_sample/internal/ui"
	"ui_sample/internal/user"
)

const (
	// screenW は論理解像度の横幅（ピクセル）です。
	screenW = 1920
	// screenH は論理解像度の縦幅（ピクセル）です。
	screenH = 1080
)

// Game はゲーム状態（UI表示用）を保持します。
type Game struct {
	showHelp bool    // ヘルプ表示フラグ
	unit     ui.Unit // 表示対象ユニット

	// 一覧/詳細の画面モード
	mode       screenMode
	units      []ui.Unit
	selIndex   int
	hoverIndex int

	userTable *user.Table
	userPath  string
	rng       *rand.Rand

	popupActive     bool
	popupGains      ui.LevelUpGains
	popupJustOpened bool
}

type screenMode int

const (
	modeList screenMode = iota
	modeStatus
	modeBattle
)

func pointIn(px, py, x, y, w, h int) bool {
	return px >= x && py >= y && px < x+w && py < y+h
}

// 簡易戦闘の1ラウンドを実行し、結果をUIとユーザJSONへ反映します。
func (g *Game) runBattleRound() {
	if len(g.units) < 2 {
		return
	}
	atkIdx := g.selIndex
	defIdx := (g.selIndex + 1) % len(g.units)
	atk := g.units[atkIdx]
	def := g.units[defIdx]

	// 武器威力（先頭装備）
	wepMight := 0
	if len(atk.Equip) > 0 {
		if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
			if w, ok := wt.Find(atk.Equip[0].Name); ok {
				wepMight = w.Might
			}
		}
		// 使用回数を1つ消費
		if atk.Equip[0].Uses > 0 {
			atk.Equip[0].Uses--
		}
	}
	// 命中（単純化）
	hit := 80 + atk.Stats.Skl*2 + atk.Stats.Lck/2 - (def.Stats.Spd*2 + def.Stats.Lck)
	if hit < 0 {
		hit = 0
	}
	if hit > 100 {
		hit = 100
	}

	// 攻撃
	if g.rng.Intn(100) < hit {
		dmg := atk.Stats.Str + wepMight - def.Stats.Def
		if dmg < 0 {
			dmg = 0
		}
		def.HP -= dmg
		if def.HP < 0 {
			def.HP = 0
		}
	}
	// 反撃（隣接想定）
	if def.HP > 0 {
		// 反撃も簡易。敵の先頭装備威力
		w2 := 0
		if len(def.Equip) > 0 {
			if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
				if w, ok := wt.Find(def.Equip[0].Name); ok {
					w2 = w.Might
				}
			}
		}
		hit2 := 80 + def.Stats.Skl*2 + def.Stats.Lck/2 - (atk.Stats.Spd*2 + atk.Stats.Lck)
		if hit2 < 0 {
			hit2 = 0
		}
		if hit2 > 100 {
			hit2 = 100
		}
		if g.rng.Intn(100) < hit2 {
			dmg := def.Stats.Str + w2 - atk.Stats.Def
			if dmg < 0 {
				dmg = 0
			}
			atk.HP -= dmg
			if atk.HP < 0 {
				atk.HP = 0
			}
		}
	}

	// 反映
	g.units[atkIdx] = atk
	g.units[defIdx] = def
	g.unit = atk

	// 保存（両者）
	if g.userTable != nil {
		if c, ok := g.userTable.Find(atk.ID); ok {
			c.HP = atk.HP
			c.HPMax = atk.HPMax
			// 装備の使用回数反映（先頭のみ簡易）
			if len(c.Equip) > 0 && len(atk.Equip) > 0 {
				c.Equip[0].Uses = atk.Equip[0].Uses
			}
			g.userTable.UpdateCharacter(c)
		}
		if c2, ok := g.userTable.Find(def.ID); ok {
			c2.HP = def.HP
			c2.HPMax = def.HPMax
			g.userTable.UpdateCharacter(c2)
		}
		_ = g.userTable.Save(g.userPath)
	}
}

// NewGame は Game を初期化して返します。
func NewGame() *Game {
	g := &Game{}
	g.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	g.userPath = "db/user/usr_characters.json"
	if ut, err := user.LoadFromJSON(g.userPath); err == nil {
		g.userTable = ut
	}
	// ユーザテーブルから一覧を読み込む
	if us, err := ui.LoadUnitsFromUser(g.userPath); err == nil && len(us) > 0 {
		g.units = us
		g.selIndex = 0
		g.unit = us[0]
	} else {
		// フォールバック
		g.unit = ui.SampleUnit()
		g.units = []ui.Unit{g.unit}
		g.selIndex = 0
	}
	g.mode = modeList
	g.hoverIndex = -1
	return g
}

// Update は毎フレームの更新処理を行います。
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		// データ再読み込み
		if us, err := ui.LoadUnitsFromUser("db/user/usr_characters.json"); err == nil && len(us) > 0 {
			g.units = us
			if g.selIndex >= len(us) {
				g.selIndex = 0
			}
			g.unit = us[g.selIndex]
		} else {
			g.unit = ui.SampleUnit()
			g.units = []ui.Unit{g.unit}
			g.selIndex = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		g.showHelp = true
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.showHelp = false
	}

	// 入力（マウス）
	mx, my := ebiten.CursorPosition()
	switch g.mode {
	case modeList:
		g.hoverIndex = -1
		for i := range g.units {
			x, y, w, h := ui.ListItemRect(screenW, screenH, i)
			if pointIn(mx, my, x, y, w, h) {
				g.hoverIndex = i
			}
		}
		if g.hoverIndex >= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.selIndex = g.hoverIndex
			g.unit = g.units[g.selIndex]
			g.mode = modeStatus
		}
	case modeStatus:
		// レベルアップボタン
		lbx, lby, lbw, lbh := ui.LevelUpButtonRect(screenW, screenH)
		lvEnabled := g.unit.Level < game.LevelCap && !g.popupActive
		if lvEnabled && pointIn(mx, my, lbx, lby, lbw, lbh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			// 抽選 → 反映 → 保存 → ポップアップ表示
			gains := ui.RollLevelUp(g.unit, g.rng.Float64)
			ui.ApplyGains(&g.unit, gains, game.LevelCap)
			g.units[g.selIndex] = g.unit
			g.popupGains = gains
			g.popupActive = true
			g.popupJustOpened = true
			// 保存
			if g.userTable != nil {
				if c, ok := g.userTable.Find(g.unit.ID); ok {
					c.Level = g.unit.Level
					c.HPMax = g.unit.HPMax
					c.Stats = user.Stats{Str: g.unit.Stats.Str, Mag: g.unit.Stats.Mag, Skl: g.unit.Stats.Skl, Spd: g.unit.Stats.Spd, Lck: g.unit.Stats.Lck, Def: g.unit.Stats.Def, Res: g.unit.Stats.Res, Mov: g.unit.Stats.Mov}
					g.userTable.UpdateCharacter(c)
					_ = g.userTable.Save(g.userPath)
				}
			}
		}
		// ポップアップ閉じる
		if g.popupActive {
			if g.popupJustOpened {
				g.popupJustOpened = false
			} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				g.popupActive = false
			}
		}
		bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
		if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.mode = modeList
		}
		// 戦闘へ
		sbx, sby, sbw, sbh := ui.ToBattleButtonRect(screenW, screenH)
		if pointIn(mx, my, sbx, sby, sbw, sbh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if len(g.units) > 1 {
				g.mode = modeBattle
			}
		}
	case modeBattle:
		// 戻る
		bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
		if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.mode = modeStatus
		}
		// 戦闘開始
		bx2, by2, bw2, bh2 := ui.BattleStartButtonRect(screenW, screenH)
		if pointIn(mx, my, bx2, by2, bw2, bh2) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.runBattleRound()
		}
	}
	return nil
}

// Draw は画面描画を行います。
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{12, 18, 30, 255})
	switch g.mode {
	case modeList:
		ui.DrawCharacterList(screen, g.units, g.hoverIndex)
	case modeStatus:
		ui.DrawStatus(screen, g.unit)
		// 戻るボタン
		mx, my := ebiten.CursorPosition()
		bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
		hovered := pointIn(mx, my, bx, by, bw, bh)
		ui.DrawBackButton(screen, hovered)
		// レベルアップボタン
		lvx, lvy, lvw, lvh := ui.LevelUpButtonRect(screenW, screenH)
		lvHovered := pointIn(mx, my, lvx, lvy, lvw, lvh)
		ui.DrawLevelUpButton(screen, lvHovered, g.unit.Level < game.LevelCap && !g.popupActive)
		if g.popupActive {
			ui.DrawLevelUpPopup(screen, g.unit, g.popupGains)
		}
		// 戦闘へ
		sbx, sby, sbw, sbh := ui.ToBattleButtonRect(screenW, screenH)
		sbHovered := pointIn(mx, my, sbx, sby, sbw, sbh)
		ui.DrawToBattleButton(screen, sbHovered, len(g.units) > 1)
	case modeBattle:
		// 対戦相手は次のユニット
		defIdx := (g.selIndex + 1) % len(g.units)
		atk := g.units[g.selIndex]
		def := g.units[defIdx]
		ui.DrawBattle(screen, atk, def)
		mx, my := ebiten.CursorPosition()
		bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
		ui.DrawBackButton(screen, pointIn(mx, my, bx, by, bw, bh))
	}
	if g.showHelp {
		ebitenutil.DebugPrintAt(screen, "H: ヘルプ表示切替 / ESC: 閉じる\nBackspace: サンプル値を再読み込み", 16, screenH-64)
	}
}

// Layout は論理解像度（内部解像度）を返します。
func (g *Game) Layout(_, _ int) (int, int) {
	return screenW, screenH
}

// main はウィンドウを作成しゲームループを開始します。
func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Ebiten UI サンプル - ステータス画面")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
