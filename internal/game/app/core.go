package app

import (
    "math/rand"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    appuse "ui_sample/internal/app"
    "ui_sample/internal/config"
    "ui_sample/internal/game"
    "ui_sample/internal/game/scenes"
    gamesvc "ui_sample/internal/game/service"
    "ui_sample/internal/repo"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/user"
)

// NewUIAppGame は UI サンプル用にポートを注入し SceneStack を構築した ebiten.Game を返します。
func NewUIAppGame() *Game {
    // 乱数と入力
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    in := gamesvc.NewInput()
    in.BindKey(ebiten.KeyBackspace, gamesvc.Menu)

    // ユーザパス/テーブル
    userPath := config.DefaultUserPath
    var ut *user.Table
    if t, err := user.LoadFromJSON(userPath); err == nil { ut = t }
    // 一覧
    units, _ := uicore.LoadUnitsFromUser(userPath)
    if len(units) == 0 { units = []uicore.Unit{uicore.SampleUnit()} }

    // Ports（JSON）を注入して App を生成
    urepo, _ := repo.NewJSONUserRepo(userPath)
    wrepo, _ := repo.NewJSONWeaponsRepo(config.DefaultWeaponsPath)
    inv, _ := repo.NewJSONInventoryRepo(config.DefaultUserWeaponsPath, config.DefaultUserItemsPath, config.DefaultWeaponsPath, "db/master/mst_items.json")
    a := appuse.New(urepo, wrepo, inv, rng)
    scenes.SetWeaponTable(wrepo.Table())

    // 画面メトリクス初期化
    uicore.SetBaseResolution(screenW, screenH)

    // Env（共有状態）
    env := &scenes.Env{App: a, UserTable: ut, UserPath: userPath, RNG: rng, Units: units, SelIndex: 0}

    // Game（Runner + AfterUpdate）
    g := &Game{Runner: Runner{}, Input: in, Env: env, prevTime: time.Now()}
    g.Runner.AfterUpdate = func(sc game.Scene) bool {
        if p, ok := sc.(interface{ ShouldPop() bool }); ok { return p.ShouldPop() }
        return false
    }
    g.Runner.Stack.Push(scenes.NewList(env))

    // ウィンドウ・TPS
    SetupWindow()
    ebiten.SetWindowTitle("Ebiten UI サンプル - ステータス画面")
    return g
}
