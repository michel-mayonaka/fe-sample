// Package main は Ebiten を用いた FE 風ステータスUIサンプルの
// エントリポイントを提供します。
package main

import (
    "fmt"
    "image/color"
    "math/rand"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "ui_sample/internal/app"
    "ui_sample/internal/config"
    "ui_sample/internal/model"
    "ui_sample/internal/repo"
    "ui_sample/internal/assets"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/ui/core"
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
    simLogPopup bool
    // 模擬戦・選択フロー
    simSelecting bool
    simSelectStep int // 0=攻撃側選択,1=防御側選択
    chooseHover int
    simTurn int // 1始まり。奇数=左先攻, 偶数=右先攻
    simLogScroll int
    simAutoEnded bool
    // 地形選択
    attTerrainSel int
    defTerrainSel int
    // 自動実行
    simAuto bool
    simAutoCooldown int

    // 戦闘プレビュー用地形（暫定: 手動切替）
    attTerrain gcore.Terrain
    defTerrain gcore.Terrain

    // 戦闘ログ（攻撃→反撃の結果など）
    battleLogs []string
    battleLogPopup bool

    // App（ユースケース）
    app *app.App

    // 抽象入力（移行中）
    ginput *gamesvc.Input

    // Scene 運用（最小導入）
    useScenes bool
    stack     game.SceneStack

    // 一覧（武器/アイテム）
    weapons   []ui.WeaponRow
    items     []ui.ItemRow
    invTab    int // 0=武器, 1=アイテム
    hoverInv  int

    // 装備変更
    selectingEquip bool
    selectingIsWeapon bool
    currentSlot int

    // Backspace長押し判定用（データリロードの誤爆防止）
    reloadHold int
}

type screenMode int

const (
    modeList screenMode = iota
    modeStatus
    modeBattle
    modeSimBattle
    modeInventory
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
    // ローカルの userTable が存在する場合は HP/Max のみ同期
    if g.userTable != nil {
        atkIdx := g.selIndex
        defIdx := (g.selIndex + 1) % len(g.units)
        atk := g.units[atkIdx]
        def := g.units[defIdx]
        if c, ok := g.userTable.Find(atk.ID); ok {
            c.HP = atk.HP
            c.HPMax = atk.HPMax
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
    // 抽象入力初期化（Backspace→Menu に割当し、従来のリロード操作を維持）
    g.ginput = gamesvc.NewInput()
    g.ginput.BindKey(ebiten.KeyBackspace, gamesvc.Menu)
    g.attTerrainSel, g.defTerrainSel = 0, 0
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
            inv, _ := appInitInventoryRepo(config.DefaultUserWeaponsPath, config.DefaultUserItemsPath, config.DefaultWeaponsPath)
            g.app = app.New(ur, wr, inv, g.rng)
            // UIへ武器テーブルを共有
            ui.SetWeaponTable(wr.Table())
        }
    }
    // メトリクス初期化（論理解像度）
    uicore.SetBaseResolution(screenW, screenH)
    // SceneStack 最小導入（一覧→模擬戦の遷移のみ）
    g.useScenes = true
    g.stack.Push(&listScene{g: g})
    return g
}

func appInitUserRepo(path string) (*repo.JSONUserRepo, error) {
    return repo.NewJSONUserRepo(path)
}

func appInitWeaponsRepo(path string) (*repo.JSONWeaponsRepo, error) {
    return repo.NewJSONWeaponsRepo(path)
}

func appInitInventoryRepo(usrW, usrI, mstW string) (*repo.JSONInventoryRepo, error) {
    return repo.NewJSONInventoryRepo(usrW, usrI, mstW, "db/master/mst_items.json")
}

// Update は毎フレームの更新処理を行います。
func (g *Game) Update() error {
    // 抽象入力スナップショット（移行中のため併用）
    if g.ginput != nil { g.ginput.Snapshot() }
    g.updateGlobalToggles()

    if g.useScenes {
        // Scene 駆動
        ctx := &game.Ctx{ScreenW: screenW, ScreenH: screenH, Input: g.ginput}
        if sc := g.stack.Current(); sc != nil {
            if next, _ := sc.Update(ctx); next != nil { g.stack.Push(next) }
            // Pop 条件: simScene で一覧へ戻ったらポップ
            if _, ok := sc.(*simScene); ok && g.mode == modeList { g.stack.Pop() }
        }
        return nil
    }

    // 入力（マウス）
    mx, my := ebiten.CursorPosition()
    switch g.mode {
    case modeList:
        g.updateListMode(mx, my)
    case modeStatus:
        g.updateStatusMode(mx, my)
    case modeBattle:
        // 戻る
        bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
        // ログポップアップ表示中はポップアップを優先（クリック/Z/Enterで閉じる）
        if g.battleLogPopup {
            if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || (g.ginput != nil && g.ginput.Press(gamesvc.Confirm)) {
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
        // キー操作: Confirmで戦闘、Cancelで戻る
        if canStart && g.ginput != nil && g.ginput.Press(gamesvc.Confirm) {
            g.runBattleRound()
        }
        if g.ginput != nil && g.ginput.Press(gamesvc.Cancel) {
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
        // ログポップアップ中は閉じる操作のみ受け付け
        if g.simLogPopup {
            if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || (g.ginput != nil && g.ginput.Press(gamesvc.Confirm)) {
                g.simLogPopup = false
                g.simAutoEnded = false
                g.simLogScroll = 0
            }
            return nil
        }
        // 戻る
        bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
        if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.mode = modeList
            g.simActive = false
        }
        if g.ginput != nil && g.ginput.Press(gamesvc.Cancel) {
            g.mode = modeList
            g.simActive = false
        }
        // データ再読み込みは Backspace（Menu）に集約
        // 戦闘開始（コピーでシミュレーション）
        bx2, by2, bw2, bh2 := ui.BattleStartButtonRect(screenW, screenH)
        // 自動実行ボタン
        ax, ay, aw, ah := ui.AutoRunButtonRect(screenW, screenH)
        aHovered := pointIn(mx, my, ax, ay, aw, ah)
        if aHovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.simAuto = !g.simAuto
            if g.simAuto { g.simLogPopup = false }
        }
        // 実行可能条件: 両者HP>0
        canStart := g.simAtk.HP > 0 && g.simDef.HP > 0
        leftFirst := (g.simTurn%2 == 1)
        if canStart && pointIn(mx, my, bx2, by2, bw2, bh2) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            if leftFirst {
                a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simAtk, g.simDef, g.attTerrain, g.defTerrain, g.rng)
                g.simAtk, g.simDef = a, d
                g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simAtk.Name)}, logs...)
            } else {
                a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simDef, g.simAtk, g.defTerrain, g.attTerrain, g.rng)
                // a=右, d=左
                g.simDef, g.simAtk = a, d
                g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simDef.Name)}, logs...)
            }
            g.simLogPopup = true
            g.simTurn++
        }
        if canStart && g.ginput != nil && g.ginput.Press(gamesvc.Confirm) {
            if leftFirst {
                a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simAtk, g.simDef, g.attTerrain, g.defTerrain, g.rng)
                g.simAtk, g.simDef = a, d
                g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simAtk.Name)}, logs...)
            } else {
                a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simDef, g.simAtk, g.defTerrain, g.attTerrain, g.rng)
                g.simDef, g.simAtk = a, d
                g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simDef.Name)}, logs...)
            }
            g.simLogPopup = true
            g.simTurn++
        }
        // 自動実行: 一定クールダウンで連続ターンを再生（決着で停止）
        if g.simAuto && canStart && !g.simLogPopup {
            if g.simAutoCooldown > 0 {
                g.simAutoCooldown--
            } else {
                if leftFirst {
                    a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simAtk, g.simDef, g.attTerrain, g.defTerrain, g.rng)
                    g.simAtk, g.simDef = a, d
                    g.simLogs = append(g.simLogs, append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simAtk.Name)}, logs...)...)
                } else {
                    a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simDef, g.simAtk, g.defTerrain, g.attTerrain, g.rng)
                    g.simDef, g.simAtk = a, d
                    g.simLogs = append(g.simLogs, append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simDef.Name)}, logs...)...)
                }
                g.simTurn++
                g.simAutoCooldown = 10
                if g.simAtk.HP <= 0 || g.simDef.HP <= 0 {
                    g.simAuto = false
                    g.simLogPopup = true
                    g.simAutoEnded = true
                }
            }
        }
        // 自動実行ポップアップのスクロール（ホイール/矢印/PgUp/PgDn）
        if g.simAuto {
            _, wy := ebiten.Wheel()
            if wy != 0 { g.simLogScroll -= int(wy) }
            if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) { g.simLogScroll++ }
            if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) { if g.simLogScroll > 0 { g.simLogScroll-- } }
            if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) { g.simLogScroll += 5 }
            if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) { g.simLogScroll -= 5; if g.simLogScroll < 0 { g.simLogScroll = 0 } }
            // 最新追従（自動実行中は新ログでスクロールが負方向にならないよう下限0を保つ）
            if g.simLogScroll < 0 { g.simLogScroll = 0 }
        }
        // 地形ボタン（クリック選択）
        mx, my := ebiten.CursorPosition()
        for i := 0; i < 3; i++ {
            ax, ay, aw, ah := ui.TerrainButtonRect(screenW, screenH, true, i)
            if pointIn(mx, my, ax, ay, aw, ah) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
                g.attTerrainSel = i
                switch i { case 0: g.attTerrain = terrainPlain(); case 1: g.attTerrain = terrainForest(); case 2: g.attTerrain = terrainFort() }
            }
            dx, dy, dw, dh := ui.TerrainButtonRect(screenW, screenH, false, i)
            if pointIn(mx, my, dx, dy, dw, dh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
                g.defTerrainSel = i
                switch i { case 0: g.defTerrain = terrainPlain(); case 1: g.defTerrain = terrainForest(); case 2: g.defTerrain = terrainFort() }
            }
        }
        // キー（互換操作）: 1/2/3 = 攻、Shift+1/2/3 = 防
        if inpututil.IsKeyJustPressed(ebiten.Key1) {
            if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) { g.defTerrainSel = 0; g.defTerrain = terrainPlain() } else { g.attTerrainSel = 0; g.attTerrain = terrainPlain() }
        }
        if inpututil.IsKeyJustPressed(ebiten.Key2) {
            if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) { g.defTerrainSel = 1; g.defTerrain = terrainForest() } else { g.attTerrainSel = 1; g.attTerrain = terrainForest() }
        }
        if inpututil.IsKeyJustPressed(ebiten.Key3) {
            if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) { g.defTerrainSel = 2; g.defTerrain = terrainFort() } else { g.attTerrainSel = 2; g.attTerrain = terrainFort() }
        }
    case modeInventory:
        g.updateInventory()
        }
        return nil
}

// updateGlobalToggles はヘルプ表示やデータ再読み込みなどのグローバル操作を処理します。
func (g *Game) updateGlobalToggles() {
    // Backspace 長押しで再読み込み（App+画像キャッシュを一括）
    if g.ginput != nil && g.ginput.Down(gamesvc.Menu) {
        g.reloadHold++
        if g.reloadHold == 30 { // 約0.5秒（60FPS時）
            if g.app != nil {
                _ = g.app.ReloadData()
                ui.SetWeaponTable(g.app.WeaponsTable())
            }
            assets.Clear()
            // UIユニットを再構築
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
    } else {
        g.reloadHold = 0
    }
    if ebiten.IsKeyPressed(ebiten.KeyH) {
        g.showHelp = true
    } else if g.ginput != nil && g.ginput.Down(gamesvc.Cancel) {
        g.showHelp = false
    }
}

// updateListMode は一覧画面での入力処理を行います。
func (g *Game) updateListMode(mx, my int) {
    // 行ホバー検出
    g.hoverIndex = -1
    for i := range g.units {
        x, y, w, h := ui.ListItemRect(screenW, screenH, i)
        if pointIn(mx, my, x, y, w, h) {
            g.hoverIndex = i
        }
    }
    // 行クリックでステータスへ遷移
    if g.hoverIndex >= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        g.selIndex = g.hoverIndex
        if g.selIndex >= 0 && g.selIndex < len(g.units) {
            g.unit = g.units[g.selIndex]
        }
        g.mode = modeStatus
        return
    }

    // ショートカット: 武器/アイテム一覧を開く
    if g.ginput != nil && g.ginput.Press(gamesvc.OpenWeapons) {
        if g.app != nil && g.app.Inv != nil {
            g.weapons = ui.BuildWeaponRowsWithOwners(g.app.Inv.Weapons(), g.app.WeaponsTable(), g.userTable)
            g.invTab, g.hoverInv = 0, -1
            g.selectingEquip, g.selectingIsWeapon = false, true
            g.mode = modeInventory
            return
        }
    }
    if g.ginput != nil && g.ginput.Press(gamesvc.OpenItems) {
        if g.app != nil && g.app.Inv != nil {
            if it, err := model.LoadItemsJSON("db/master/mst_items.json"); err == nil {
                g.items = ui.BuildItemRowsWithOwners(g.app.Inv.Items(), it, g.userTable)
            }
            g.invTab, g.hoverInv = 1, -1
            g.selectingEquip, g.selectingIsWeapon = false, false
            g.mode = modeInventory
            return
        }
    }

    // 模擬戦選択フロー開始（ボタン）
    bx, by, bw, bh := ui.SimBattleButtonRect(screenW, screenH)
    if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && len(g.units) > 1 {
        g.simSelecting = true
        g.simSelectStep = 0
        g.chooseHover = -1
        return
    }
    // 選択ポップアップ操作中
    if g.simSelecting {
        // キャンセル
        if g.ginput != nil && g.ginput.Press(gamesvc.Cancel) {
            g.simSelecting = false
            g.simSelectStep = 0
            g.chooseHover = -1
            return
        }
        // ホバー更新
        g.chooseHover = -1
        for i := range g.units {
            x, y, w, h := ui.ChooseUnitItemRect(screenW, screenH, i, len(g.units))
            if pointIn(mx, my, x, y, w, h) {
                g.chooseHover = i
            }
        }
        // クリックで選択
        if g.chooseHover >= 0 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            if g.simSelectStep == 0 {
                g.simAtk = g.units[g.chooseHover]
                g.simSelectStep = 1
                return
            }
            // 防御側選択で確定
            g.simDef = g.units[g.chooseHover]
            g.simLogs = nil
            g.simLogPopup = false
            g.simAuto = false
            g.simAutoCooldown = 0
            g.simAutoEnded = false
            g.simLogScroll = 0
            g.simTurn = 1
            // 地形は既定（平地）から開始
            g.attTerrainSel, g.defTerrainSel = 0, 0
            g.attTerrain, g.defTerrain = terrainPlain(), terrainPlain()
            g.simSelecting = false
            g.mode = modeSimBattle
            return
        }
    }
}

// updateStatusMode はステータス画面での入力処理を行います。
func (g *Game) updateStatusMode(mx, my int) {
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
        // 保存（App経由） + ローカルテーブル同期
        if g.userTable != nil {
            if c, ok := g.userTable.Find(g.unit.ID); ok {
                c.Level = g.unit.Level
                c.HPMax = g.unit.HPMax
                c.Stats = user.Stats{Str: g.unit.Stats.Str, Mag: g.unit.Stats.Mag, Skl: g.unit.Stats.Skl, Spd: g.unit.Stats.Spd, Lck: g.unit.Stats.Lck, Def: g.unit.Stats.Def, Res: g.unit.Stats.Res, Mov: g.unit.Stats.Mov}
                g.userTable.UpdateCharacter(c)
            }
        }
        if g.app != nil { _ = g.app.PersistUnit(g.unit) }
    }
    // ポップアップ閉じる
    if g.popupActive {
        if g.popupJustOpened {
            g.popupJustOpened = false
        } else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.popupActive = false
        }
    }
    // スロット操作: クリックで選択 + 一覧を開く
    for i := 0; i < 5; i++ {
        sx, sy, swd, shd := ui.EquipSlotRect(screenW, screenH, i)
        if pointIn(mx, my, sx, sy, swd, shd) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.currentSlot = i
            if g.app != nil && g.app.Inv != nil {
                if i == 0 {
                    g.weapons = ui.BuildWeaponRowsWithOwners(g.app.Inv.Weapons(), g.app.WeaponsTable(), g.userTable)
                    g.invTab = 0
                    g.hoverInv = -1
                    g.selectingEquip = true
                    g.selectingIsWeapon = true
                    g.mode = modeInventory
                } else {
                    if it, err := model.LoadItemsJSON("db/master/mst_items.json"); err == nil {
                        g.items = ui.BuildItemRowsWithOwners(g.app.Inv.Items(), it, g.userTable)
                    }
                    g.invTab = 1
                    g.hoverInv = -1
                    g.selectingEquip = true
                    g.selectingIsWeapon = false
                    g.mode = modeInventory
                }
            }
        }
    }
    // スロット操作: 数字キー 1..5
    if g.ginput != nil && g.ginput.Press(gamesvc.Slot1) { g.currentSlot = 0 }
    if g.ginput != nil && g.ginput.Press(gamesvc.Slot2) { g.currentSlot = 1 }
    if g.ginput != nil && g.ginput.Press(gamesvc.Slot3) { g.currentSlot = 2 }
    if g.ginput != nil && g.ginput.Press(gamesvc.Slot4) { g.currentSlot = 3 }
    if g.ginput != nil && g.ginput.Press(gamesvc.Slot5) { g.currentSlot = 4 }
    // スロット解除
    if (g.ginput != nil && g.ginput.Press(gamesvc.Unassign)) {
        if g.userTable != nil {
            if c, ok := g.userTable.Find(g.unit.ID); ok {
                for len(c.Equip) <= g.currentSlot { c.Equip = append(c.Equip, user.EquipRef{}) }
                c.Equip[g.currentSlot] = user.EquipRef{}
                // 末尾の空要素を圧縮
                j := len(c.Equip)
                for j > 0 {
                    if c.Equip[j-1].UserItemsID == "" && c.Equip[j-1].UserWeaponsID == "" { j-- } else { break }
                }
                c.Equip = c.Equip[:j]
                g.userTable.UpdateCharacter(c)
                _ = g.userTable.Save(g.userPath)
                g.unit = uicore.UnitFromUser(c)
                g.units[g.selIndex] = g.unit
            }
        }
    }
    // 装備付け替え開始（ショートカット）
    if g.ginput != nil && g.ginput.Press(gamesvc.EquipToggle) {
        if g.app != nil && g.app.Inv != nil {
            g.weapons = ui.BuildWeaponRowsWithOwners(g.app.Inv.Weapons(), g.app.WeaponsTable(), g.userTable)
            g.invTab = 0
            g.hoverInv = -1
            g.selectingEquip = true
            g.selectingIsWeapon = true
            g.mode = modeInventory
        }
    }
    if g.ginput != nil && g.ginput.Press(gamesvc.OpenItems) {
        if g.app != nil && g.app.Inv != nil {
            if it, err := model.LoadItemsJSON("db/master/mst_items.json"); err == nil {
                g.items = ui.BuildItemRowsWithOwners(g.app.Inv.Items(), it, g.userTable)
            }
            g.invTab = 1
            g.hoverInv = -1
            g.selectingEquip = true
            g.selectingIsWeapon = false
            g.mode = modeInventory
        }
    }
    // 戻るボタン/キー
    bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
    if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { g.mode = modeList }
    if g.ginput != nil && g.ginput.Press(gamesvc.Cancel) { g.mode = modeList }
}

// closeSimLogIfRequested は模擬戦ログポップアップの閉じ操作を処理します。
func (g *Game) closeSimLogIfRequested() bool {
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || (g.ginput != nil && g.ginput.Press(gamesvc.Confirm)) {
        g.simLogPopup = false
        g.simAutoEnded = false
        g.simLogScroll = 0
        return true
    }
    return false
}

// handleSimBack は模擬戦画面の戻る操作を処理します。
func (g *Game) handleSimBack() {
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
    if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        g.mode = modeList
        g.simActive = false
    }
    if g.ginput != nil && g.ginput.Press(gamesvc.Cancel) {
        g.mode = modeList
        g.simActive = false
    }
}

// updateSimBattleCore は模擬戦画面の主要更新（開始/自動実行/地形選択）を処理します。
func (g *Game) updateSimBattleCore() {
    // 戦闘開始（コピーでシミュレーション）
    mx, my := ebiten.CursorPosition()
    bx2, by2, bw2, bh2 := ui.BattleStartButtonRect(screenW, screenH)
    // 自動実行ボタン
    ax, ay, aw, ah := ui.AutoRunButtonRect(screenW, screenH)
    aHovered := pointIn(mx, my, ax, ay, aw, ah)
    if aHovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        g.simAuto = !g.simAuto
        if g.simAuto { g.simLogPopup = false }
    }
    // 実行可能条件: 両者HP>0
    canStart := g.simAtk.HP > 0 && g.simDef.HP > 0
    leftFirst := (g.simTurn%2 == 1)
    if canStart && pointIn(mx, my, bx2, by2, bw2, bh2) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        if leftFirst {
            a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simAtk, g.simDef, g.attTerrain, g.defTerrain, g.rng)
            g.simAtk, g.simDef = a, d
            g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simAtk.Name)}, logs...)
        } else {
            a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simDef, g.simAtk, g.defTerrain, g.attTerrain, g.rng)
            g.simDef, g.simAtk = a, d
            g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simDef.Name)}, logs...)
        }
        g.simLogPopup = true
        g.simTurn++
    }
    if canStart && g.ginput != nil && g.ginput.Press(gamesvc.Confirm) {
        if leftFirst {
            a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simAtk, g.simDef, g.attTerrain, g.defTerrain, g.rng)
            g.simAtk, g.simDef = a, d
            g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simAtk.Name)}, logs...)
        } else {
            a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simDef, g.simAtk, g.defTerrain, g.attTerrain, g.rng)
            g.simDef, g.simAtk = a, d
            g.simLogs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simDef.Name)}, logs...)
        }
        g.simLogPopup = true
        g.simTurn++
    }
    // 自動実行: 一定クールダウンで連続ターンを再生（決着で停止）
    if g.simAuto && canStart && !g.simLogPopup {
        if g.simAutoCooldown > 0 {
            g.simAutoCooldown--
        } else {
            if leftFirst {
                a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simAtk, g.simDef, g.attTerrain, g.defTerrain, g.rng)
                g.simAtk, g.simDef = a, d
                g.simLogs = append(g.simLogs, append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simAtk.Name)}, logs...)...)
            } else {
                a, d, logs := ui.SimulateBattleCopyWithTerrain(g.simDef, g.simAtk, g.defTerrain, g.attTerrain, g.rng)
                g.simDef, g.simAtk = a, d
                g.simLogs = append(g.simLogs, append([]string{fmt.Sprintf("ターン %d 先攻: %s", g.simTurn, g.simDef.Name)}, logs...)...)
            }
            g.simTurn++
            g.simAutoCooldown = 10
            if g.simAtk.HP <= 0 || g.simDef.HP <= 0 {
                g.simAuto = false
                g.simLogPopup = true
                g.simAutoEnded = true
            }
        }
    }
    // 自動実行ポップアップのスクロール（ホイール/矢印/PgUp/PgDn）
    if g.simAuto {
        _, wy := ebiten.Wheel()
        if wy != 0 { g.simLogScroll -= int(wy) }
        if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) { g.simLogScroll++ }
        if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) { if g.simLogScroll > 0 { g.simLogScroll-- } }
        if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) { g.simLogScroll += 5 }
        if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) { g.simLogScroll -= 5; if g.simLogScroll < 0 { g.simLogScroll = 0 } }
        if g.simLogScroll < 0 { g.simLogScroll = 0 }
    }
    // 地形ボタン（クリック選択）
    mx, my = ebiten.CursorPosition()
    for i := 0; i < 3; i++ {
        ax, ay, aw, ah := ui.TerrainButtonRect(screenW, screenH, true, i)
        if pointIn(mx, my, ax, ay, aw, ah) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.attTerrainSel = i
            switch i { case 0: g.attTerrain = terrainPlain(); case 1: g.attTerrain = terrainForest(); case 2: g.attTerrain = terrainFort() }
        }
        dx, dy, dw, dh := ui.TerrainButtonRect(screenW, screenH, false, i)
        if pointIn(mx, my, dx, dy, dw, dh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            g.defTerrainSel = i
            switch i { case 0: g.defTerrain = terrainPlain(); case 1: g.defTerrain = terrainForest(); case 2: g.defTerrain = terrainFort() }
        }
    }
    // キー（互換操作）: 1/2/3 = 攻、Shift+1/2/3 = 防（暫定: 直接キー）
    if inpututil.IsKeyJustPressed(ebiten.Key1) {
        if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) { g.defTerrainSel = 0; g.defTerrain = terrainPlain() } else { g.attTerrainSel = 0; g.attTerrain = terrainPlain() }
    }
    if inpututil.IsKeyJustPressed(ebiten.Key2) {
        if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) { g.defTerrainSel = 1; g.defTerrain = terrainForest() } else { g.attTerrainSel = 1; g.attTerrain = terrainForest() }
    }
    if inpututil.IsKeyJustPressed(ebiten.Key3) {
        if ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) { g.defTerrainSel = 2; g.defTerrain = terrainFort() } else { g.attTerrainSel = 2; g.attTerrain = terrainFort() }
    }
}

// updateInventory は在庫タブ画面の入力処理をまとめたものです。
func (g *Game) updateInventory() {
    screenW, screenH := screenW, screenH
    // タブ切替
    tabW, tabH := uicore.S(160), uicore.S(44)
    lm := uicore.ListMarginPx()
    tx := lm + uicore.S(20)
    ty := lm + uicore.S(12)
    rxW, rxI := tx, tx+tabW+uicore.S(10)
    mx, my := ebiten.CursorPosition()
    if pointIn(mx, my, rxW, ty, tabW, tabH) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { g.invTab = 0 }
    if pointIn(mx, my, rxI, ty, tabW, tabH) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { g.invTab = 1 }
    // リスト選択
    g.hoverInv = -1
    if g.invTab == 0 {
        for i := range g.weapons {
            x, y, w, h := ui.ListItemRect(screenW, screenH, i)
            if pointIn(mx, my, x, y, w, h) { g.hoverInv = i }
            if g.selectingEquip && g.selectingIsWeapon && g.hoverInv == i && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
                chosen := g.weapons[i]
                if g.userTable != nil {
                    if c, ok := g.userTable.Find(g.unit.ID); ok {
                        var prev user.EquipRef
                        if g.currentSlot < len(c.Equip) { prev = c.Equip[g.currentSlot] }
                        ownerID := ""
                        ownerSlot := -1
                        for _, oc := range g.userTable.Slice() {
                            for idx, er := range oc.Equip {
                                if er.UserWeaponsID == chosen.ID { ownerID = oc.ID; ownerSlot = idx; break }
                            }
                            if ownerID != "" { break }
                        }
                        if ownerID != "" {
                            if oc, ok2 := g.userTable.Find(ownerID); ok2 {
                                for len(oc.Equip) <= ownerSlot { oc.Equip = append(oc.Equip, user.EquipRef{}) }
                                oc.Equip[ownerSlot] = prev
                                j := len(oc.Equip)
                                for j > 0 { if oc.Equip[j-1].UserItemsID == "" && oc.Equip[j-1].UserWeaponsID == "" { j-- } else { break } }
                                oc.Equip = oc.Equip[:j]
                                g.userTable.UpdateCharacter(oc)
                                g.refreshUnitByID(ownerID)
                            }
                        }
                        for len(c.Equip) <= g.currentSlot { c.Equip = append(c.Equip, user.EquipRef{}) }
                        c.Equip[g.currentSlot] = user.EquipRef{UserWeaponsID: chosen.ID}
                        g.userTable.UpdateCharacter(c)
                        _ = g.userTable.Save(g.userPath)
                        g.unit = uicore.UnitFromUser(c)
                        g.units[g.selIndex] = g.unit
                        g.selectingEquip = false
                        g.mode = modeStatus
                        break
                    }
                }
            }
        }
    } else {
        for i := range g.items {
            x, y, w, h := ui.ListItemRect(screenW, screenH, i)
            if pointIn(mx, my, x, y, w, h) { g.hoverInv = i }
            if g.selectingEquip && !g.selectingIsWeapon && g.hoverInv == i && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
                chosen := g.items[i]
                if g.userTable != nil {
                    if c, ok := g.userTable.Find(g.unit.ID); ok {
                        var prev user.EquipRef
                        if g.currentSlot < len(c.Equip) { prev = c.Equip[g.currentSlot] }
                        ownerID := ""
                        ownerSlot := -1
                        for _, oc := range g.userTable.Slice() {
                            for idx, er := range oc.Equip {
                                if er.UserItemsID == chosen.ID { ownerID = oc.ID; ownerSlot = idx; break }
                            }
                            if ownerID != "" { break }
                        }
                        if ownerID != "" {
                            if oc, ok2 := g.userTable.Find(ownerID); ok2 {
                                for len(oc.Equip) <= ownerSlot { oc.Equip = append(oc.Equip, user.EquipRef{}) }
                                oc.Equip[ownerSlot] = prev
                                j := len(oc.Equip)
                                for j > 0 { if oc.Equip[j-1].UserItemsID == "" && oc.Equip[j-1].UserWeaponsID == "" { j-- } else { break } }
                                oc.Equip = oc.Equip[:j]
                                g.userTable.UpdateCharacter(oc)
                                g.refreshUnitByID(ownerID)
                            }
                        }
                        for len(c.Equip) <= g.currentSlot { c.Equip = append(c.Equip, user.EquipRef{}) }
                        c.Equip[g.currentSlot] = user.EquipRef{UserItemsID: chosen.ID}
                        g.userTable.UpdateCharacter(c)
                        _ = g.userTable.Save(g.userPath)
                        g.unit = uicore.UnitFromUser(c)
                        g.units[g.selIndex] = g.unit
                        g.selectingEquip = false
                        g.mode = modeStatus
                        break
                    }
                }
            }
        }
    }
    // 戻る
    bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
    if pointIn(mx, my, bx, by, bw, bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { g.mode = modeList }
    if g.ginput != nil && g.ginput.Press(gamesvc.Cancel) { g.mode = modeList }
}

// refreshUnitByID は g.userTable の内容から該当IDの UIユニットを再構築して差し替えます。
func (g *Game) refreshUnitByID(id string) {
    if g == nil || g.userTable == nil { return }
    c, ok := g.userTable.Find(id)
    if !ok { return }
    u := uicore.UnitFromUser(c)
    for i := range g.units {
        if g.units[i].ID == id {
            g.units[i] = u
            if g.selIndex == i {
                g.unit = u
            }
            break
        }
    }
}

// Draw は画面描画を行います。
func (g *Game) Draw(screen *ebiten.Image) {
    // ウィンドウサイズからスケール更新
    uicore.UpdateMetricsFromWindow()
    uicore.MaybeUpdateFontFaces()
    screen.Fill(color.RGBA{12, 18, 30, 255})
    if g.useScenes {
        if sc := g.stack.Current(); sc != nil {
            sc.Draw(screen)
        }
        if g.showHelp {
            ebitenutil.DebugPrintAt(screen, "H: ヘルプ表示切替 / ESC: 閉じる\nBackspace: サンプル値を再読み込み", 16, screenH-64)
        }
        return
    }
    switch g.mode {
    case modeList:
        g.drawList(screen)
    case modeStatus:
        g.drawStatus(screen)
    case modeBattle:
        // Phase3: 実戦画面は非推奨（将来撤去）。現状はステータスから遷移しないため通常到達しない。
        fallthrough
    case modeSimBattle:
        g.drawSimBattle(screen)
    case modeInventory:
        g.drawInventory(screen)
    }
	if g.showHelp {
		ebitenutil.DebugPrintAt(screen, "H: ヘルプ表示切替 / ESC: 閉じる\nBackspace: サンプル値を再読み込み", 16, screenH-64)
	}
}

// drawList は一覧画面を描画し、UIボタン類を表示します。
// （後続PR）描画処理は draw* 系へ段階的に分割予定
// Layout は論理解像度（内部解像度）を返します。
func (g *Game) Layout(_, _ int) (int, int) {
    return screenW, screenH
}

// drawList は一覧画面の描画を行います。
func (g *Game) drawList(screen *ebiten.Image) {
    // 本体（一覧）
    ui.DrawCharacterList(screen, g.units, g.hoverIndex)
    // 模擬戦ボタン（統一スタイル）
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.SimBattleButtonRect(screenW, screenH)
    hovered := pointIn(mx, my, bx, by, bw, bh)
    ui.DrawSimBattleButton(screen, hovered, len(g.units) > 1)
    // ショートカットガイド
    ebitenutil.DebugPrintAt(screen, "W: 武器一覧 / I: アイテム一覧", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
    // 選択フローのガイド
    if g.simSelecting {
        title := "模擬戦: 攻撃側を選択"
        if g.simSelectStep == 1 { title = "模擬戦: 防御側を選択" }
        ui.DrawChooseUnitPopup(screen, title, g.units, g.chooseHover)
    }
}

// drawStatus はステータス画面の描画を行います。
func (g *Game) drawStatus(screen *ebiten.Image) {
    // 本体（ステータス）
    ui.DrawStatus(screen, g.unit)
    // 戻るボタン
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
    hovered := pointIn(mx, my, bx, by, bw, bh)
    ui.DrawBackButton(screen, hovered)
    ebitenutil.DebugPrintAt(screen, fmt.Sprintf("選択中スロット: %d", g.currentSlot+1), uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
    // レベルアップボタン
    lvx, lvy, lvw, lvh := ui.LevelUpButtonRect(screenW, screenH)
    lvHovered := pointIn(mx, my, lvx, lvy, lvw, lvh)
    ui.DrawLevelUpButton(screen, lvHovered, g.unit.Level < game.LevelCap && !g.popupActive)
    if g.popupActive {
        ui.DrawLevelUpPopup(screen, g.unit, g.popupGains)
    }
    // Phase3: ステータス画面の「戦闘へ」ボタンは削除
}

// drawSimBattle は新バトルシミュレータの描画を行います。
func (g *Game) drawSimBattle(screen *ebiten.Image) {
    // 新バトルシミュレータ（battleレイアウトを使用）
    canStart := g.simAtk.HP > 0 && g.simDef.HP > 0 && !g.simLogPopup
    ui.DrawBattleWithTerrain(screen, g.simAtk, g.simDef, g.attTerrain, g.defTerrain, canStart)
    // 地形ボタン
    attIdx := g.attTerrainSel
    defIdx := g.defTerrainSel
    ui.DrawTerrainButtons(screen, attIdx, defIdx)
    // 自動実行ボタン
    mx, my := ebiten.CursorPosition()
    ax, ay, aw, ah := ui.AutoRunButtonRect(screenW, screenH)
    aHovered := pointIn(mx, my, ax, ay, aw, ah)
    ui.DrawAutoRunButton(screen, aHovered, g.simAuto)
    // ログ表示（自動実行中はポップアップでスクロール可）
    if g.simAuto {
        ui.DrawBattleLogOverlayScroll(screen, g.simLogs, g.simLogScroll)
    } else if g.simLogPopup {
        if g.simAutoEnded {
            ui.DrawBattleLogOverlayScroll(screen, g.simLogs, g.simLogScroll)
        } else {
            ui.DrawBattleLogOverlay(screen, g.simLogs)
        }
    }
    // 先攻表示（ヘッダ下）
    if g.simTurn <= 0 { g.simTurn = 1 }
    leftFirst := (g.simTurn%2 == 1)
    label := "先攻: "
    if leftFirst { label += g.simAtk.Name } else { label += g.simDef.Name }
    ebitenutil.DebugPrintAt(screen, label, uicore.ListMarginPx()+uicore.S(40), uicore.ListMarginPx()+uicore.S(56))
    // 既定のポップアップ（上の分岐で描画済み）
    mx, my = ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
    ui.DrawBackButton(screen, pointIn(mx, my, bx, by, bw, bh))
}

// drawInventory は在庫タブ画面の描画を行います。
func (g *Game) drawInventory(screen *ebiten.Image) {
    // 本体（タブに応じて）
    if g.invTab == 0 { ui.DrawWeaponList(screen, g.weapons, g.hoverInv) } else { ui.DrawItemList(screen, g.items, g.hoverInv) }
    // タブ描画
    tabW, tabH := uicore.S(160), uicore.S(44)
    lm := uicore.ListMarginPx()
    tx := lm + uicore.S(20)
    ty := lm + uicore.S(12)
    // 武器タブ
    uicore.DrawFramedRect(screen, float32(tx), float32(ty), float32(tabW), float32(tabH))
    baseW := color.RGBA{40, 60, 110, 255}
    if g.invTab == 0 { baseW = color.RGBA{70, 100, 160, 255} }
    vector.DrawFilledRect(screen, float32(tx), float32(ty), float32(tabW), float32(tabH), baseW, false)
    uicore.TextDraw(screen, "武器", uicore.FaceMain, tx+uicore.S(56), ty+uicore.S(30), uicore.ColText)
    // アイテムタブ
    tx2 := tx + tabW + uicore.S(10)
    uicore.DrawFramedRect(screen, float32(tx2), float32(ty), float32(tabW), float32(tabH))
    baseI := color.RGBA{40, 60, 110, 255}
    if g.invTab == 1 { baseI = color.RGBA{70, 100, 160, 255} }
    vector.DrawFilledRect(screen, float32(tx2), float32(ty), float32(tabW), float32(tabH), baseI, false)
    uicore.TextDraw(screen, "アイテム", uicore.FaceMain, tx2+uicore.S(34), ty+uicore.S(30), uicore.ColText)
    // 戻るボタン
    mx, my := ebiten.CursorPosition()
    bx, by, bw, bh := ui.BackButtonRect(screenW, screenH)
    hovered := pointIn(mx, my, bx, by, bw, bh)
    ui.DrawBackButton(screen, hovered)
    if g.selectingEquip {
        ebitenutil.DebugPrintAt(screen, "クリックでスロットに装備", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
    }
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
