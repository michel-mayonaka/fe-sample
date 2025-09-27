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
    "ui_sample/internal/app"
    "ui_sample/internal/config"
    "ui_sample/internal/repo"
    "ui_sample/internal/game"
    "ui_sample/internal/ui"
    "ui_sample/internal/user"
    gcore "ui_sample/pkg/game"
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

	// 模擬戦
    simActive bool
    simAtk    ui.Unit
    simDef    ui.Unit
    simLogs   []string

    // 戦闘プレビュー用地形（暫定: 手動切替）
    attTerrain gcore.Terrain
    defTerrain gcore.Terrain

    // 戦闘ログ（攻撃→反撃の結果など）
    battleLogs []string
    battleLogPopup bool

    // App（ユースケース）
    app *app.App
}

type screenMode int

const (
	modeList screenMode = iota
	modeStatus
	modeBattle
	modeSimBattle
)

func pointIn(px, py, x, y, w, h int) bool {
	return px >= x && py >= y && px < x+w && py < y+h
}

// 簡易戦闘の1ラウンドを実行し、結果をUIとユーザJSONへ反映します。
func (g *Game) runBattleRound() {
    if g.app == nil { return }
    updated, logs, popup, _ := g.app.RunBattleRound(g.units, g.selIndex, g.attTerrain, g.defTerrain)
    // UI状態に反映
    g.units = updated
    if g.selIndex >= 0 && g.selIndex < len(g.units) {
        g.unit = g.units[g.selIndex]
    }
    g.battleLogs = logs
    g.battleLogPopup = popup
    // ローカルの userTable が存在する場合は同期（簡易）
    if g.userTable != nil {
        atkIdx := g.selIndex
        defIdx := (g.selIndex + 1) % len(g.units)
        atk := g.units[atkIdx]
        def := g.units[defIdx]
        if c, ok := g.userTable.Find(atk.ID); ok {
            c.HP = atk.HP
            c.HPMax = atk.HPMax
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
    }
}

// toGameUnit は UIユニットを /pkg/game.Unit に変換します。
// toGameUnit は adapter に移行済み（使用箇所は削除）。

func terrainPlain() gcore.Terrain  { return gcore.Terrain{Avoid: 0, Def: 0, Hit: 0} }
func terrainForest() gcore.Terrain { return gcore.Terrain{Avoid: 20, Def: 1, Hit: 0} }
func terrainFort() gcore.Terrain   { return gcore.Terrain{Avoid: 15, Def: 2, Hit: 0} }

// NewGame は Game を初期化して返します。
func NewGame() *Game {
    g := &Game{}
	g.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	g.userPath = config.DefaultUserPath
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
    // 地形の初期値（平地）
    g.attTerrain = gcore.Terrain{}
    g.defTerrain = gcore.Terrain{}

    // App 初期化
    if ur, err := appInitUserRepo(g.userPath); err == nil {
        if wr, err2 := appInitWeaponsRepo(config.DefaultWeaponsPath); err2 == nil {
            g.app = app.New(ur, wr, g.rng)
        }
    }
    return g
}

func appInitUserRepo(path string) (*repo.JSONUserRepo, error) {
    return repo.NewJSONUserRepo(path)
}

func appInitWeaponsRepo(path string) (*repo.JSONWeaponsRepo, error) {
    return repo.NewJSONWeaponsRepo(path)
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
        // キー操作: ↑↓で選択、Z/Enterで詳細、Tabで次へ
        if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
            if g.selIndex > 0 {
                g.selIndex--
                g.unit = g.units[g.selIndex]
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
            if g.selIndex < len(g.units)-1 {
                g.selIndex++
                g.unit = g.units[g.selIndex]
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
            if len(g.units) > 0 {
                g.selIndex = (g.selIndex + 1) % len(g.units)
                g.unit = g.units[g.selIndex]
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyZ) {
            g.mode = modeStatus
        }
        // 模擬戦ボタン
        sbx, sby, sbw, sbh := ui.SimBattleButtonRect(screenW, screenH)
        if pointIn(mx, my, sbx, sby, sbw, sbh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            if len(g.units) > 1 {
                aidx := g.selIndex
                if g.hoverIndex >= 0 {
                    aidx = g.hoverIndex
                }
                didx := (aidx + 1) % len(g.units)
                g.simAtk = g.units[aidx]
                g.simDef = g.units[didx]
                a, d, logs := ui.SimulateBattleCopy(g.simAtk, g.simDef, g.rng)
                g.simAtk, g.simDef, g.simLogs = a, d, logs
                g.simActive = true
                g.mode = modeSimBattle
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
                g.battleLogs = nil
            }
        }
        // キー操作: X/Escで戻る、Z/Enterで戦闘へ
        if inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
            g.mode = modeList
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyZ) {
            if len(g.units) > 1 {
                g.mode = modeBattle
                g.battleLogs = nil
            }
        }
    case modeBattle:
        // 戻る
        bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
        // ログポップアップ表示中はポップアップを優先（クリック/Z/Enterで閉じる）
        if g.battleLogPopup {
            if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inpututil.IsKeyJustPressed(ebiten.KeyZ) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                g.battleLogPopup = false
            }
            return nil
        }
        if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.mode = modeStatus
        }
        // 戦闘開始
        bx2, by2, bw2, bh2 := ui.BattleStartButtonRect(screenW, screenH)
        // 実行可能条件: ログポップアップ非表示 かつ 両者HP>0
        defIdx := (g.selIndex + 1) % len(g.units)
        canStart := !g.battleLogPopup && g.units[g.selIndex].HP > 0 && g.units[defIdx].HP > 0
        if canStart && pointIn(mx, my, bx2, by2, bw2, bh2) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.runBattleRound()
        }
        // キー操作: Z/Enterで戦闘、X/Escで戻る
        if canStart && (inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyZ)) {
            g.runBattleRound()
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
            g.mode = modeStatus
        }
        // 地形切替（1/2/3: 攻撃側、Shift+1/2/3: 防御側）
        if inpututil.IsKeyJustPressed(ebiten.Key1) {
            if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) {
                g.defTerrain = terrainPlain()
            } else {
                g.attTerrain = terrainPlain()
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.Key2) {
            if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) {
                g.defTerrain = terrainForest()
            } else {
                g.attTerrain = terrainForest()
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.Key3) {
            if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) {
                g.defTerrain = terrainFort()
            } else {
                g.attTerrain = terrainFort()
            }
        }
    case modeSimBattle:
        bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
        if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.mode = modeList
            g.simActive = false
        }
        // キー操作: X/Escで戻る
        if inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
            g.mode = modeList
            g.simActive = false
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
		// 模擬戦ボタン（統一スタイル）
		mx, my := ebiten.CursorPosition()
		bx, by, bw, bh := ui.SimBattleButtonRect(screenW, screenH)
		hovered := pointIn(mx, my, bx, by, bw, bh)
		ui.DrawSimBattleButton(screen, hovered, len(g.units) > 1)
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
        canStart := !g.battleLogPopup && atk.HP > 0 && def.HP > 0
        ui.DrawBattleWithTerrain(screen, atk, def, g.attTerrain, g.defTerrain, canStart)
        if g.battleLogPopup {
            ui.DrawBattleLogOverlay(screen, g.battleLogs)
        }
        mx, my := ebiten.CursorPosition()
        bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
        ui.DrawBackButton(screen, pointIn(mx, my, bx, by, bw, bh))
	case modeSimBattle:
		ui.DrawSimulationBattle(screen, g.simAtk, g.simDef, g.simLogs)
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
