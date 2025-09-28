package sim

import (
    "fmt"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    scenes "ui_sample/internal/game/scenes"
    gcore "ui_sample/pkg/game"
)

// Sim は模擬戦画面の Scene 実装です。
//
// 主な責務:
// - 攻撃側/防御側の一時ユニットを保持し、戦闘をシミュレート
// - 地形選択・自動実行トグルなどの入力を意図(Intent)へ変換
// - ログのポップアップ表示と遷移終了（戻る）判定
//
// 更新フローは character_list と同一で、Update → scHandleInput → scAdvance → scFlush の順で処理します。
type Sim struct{
    E *scenes.Env
    simAtk uicore.Unit
    simDef uicore.Unit
    logs   []string
    logPopup bool
    auto bool
    autoCD int
    turn int
    attTerrain gcore.Terrain
    defTerrain gcore.Terrain
    attSel int
    defSel int
    pop bool
    sw, sh int
    // ホバー状態
    backHovered bool
    startHovered bool
    autoHovered bool
    attHover int // -1/0..2
    defHover int // -1/0..2
}

func NewSim(e *scenes.Env, atk, def uicore.Unit) *Sim { return &Sim{E:e, simAtk:atk, simDef:def, turn:1, attHover:-1, defHover:-1} }
func (s *Sim) ShouldPop() bool { return s.pop }

// Intent 種別
type IntentKind int

const (
    intentNone IntentKind = iota
    intentBack
    intentCloseLog
    intentRunOne
    intentToggleAuto
    intentSetTerrainAtt // Index: 0..2
    intentSetTerrainDef // Index: 0..2
)
type Intent struct{ Kind IntentKind; Index int }
func (Intent) IsSceneIntent() {}

// scContract はパッケージ内コンパイル保証のためのインターフェースです。
// Sim が必要な sc* メソッドを実装していることを確認します。
type scContract interface{
    scHandleInput(ctx *game.Ctx) []scenes.Intent
    scAdvance([]scenes.Intent)
    scFlush(ctx *game.Ctx)
}
var _ scContract = (*Sim)(nil)

// Update は状態更新の入口です。
// フロー: scHandleInput → scAdvance → scFlush。次シーンは本シーン内で完結するため nil を返します。
func (s *Sim) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    intents := s.scHandleInput(ctx)
    s.scAdvance(intents)
    s.scFlush(ctx)
    return nil, nil
}

// runOne は 1 ターン分の戦闘を実行し、結果ログを追加します。
func (s *Sim) runOne(leftFirst bool){
    if leftFirst {
        a,d,lines := scenes.SimulateBattleCopyWithTerrain(s.simAtk, s.simDef, s.attTerrain, s.defTerrain, s.E.RNG)
        s.simAtk, s.simDef = a,d; s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simAtk.Name)}, lines...)
    } else {
        a,d,lines := scenes.SimulateBattleCopyWithTerrain(s.simDef, s.simAtk, s.defTerrain, s.attTerrain, s.E.RNG)
        s.simDef, s.simAtk = a,d; s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simDef.Name)}, lines...)
    }
    s.logPopup=true; s.turn++
}

// Draw は模擬戦の盤面・UI とログポップアップを描画します。
func (s *Sim) Draw(dst *ebiten.Image){
    canStart := s.simAtk.HP>0 && s.simDef.HP>0 && !s.logPopup
    scenes.DrawBattleWithTerrain(dst, s.simAtk, s.simDef, s.attTerrain, s.defTerrain, canStart)
    uiwidgets.DrawTerrainButtons(dst, s.attSel, s.defSel)
    mx,my := ebiten.CursorPosition()
    ax,ay,aw,ah := scenes.AutoRunButtonRect(s.sw, s.sh)
    s.autoHovered = scenes.PointIn(mx,my,ax,ay,aw,ah)
    uiwidgets.DrawAutoRunButton(dst, s.autoHovered, s.auto)
    if s.logPopup { scenes.DrawBattleLogOverlay(dst, s.logs) }
    if s.turn<=0 { s.turn=1 }
    leftFirst := (s.turn%2==1); label := "先攻: "; if leftFirst { label+=s.simAtk.Name } else { label+=s.simDef.Name }
    ebitenutil.DebugPrintAt(dst, label, uicore.ListMarginPx()+uicore.S(40), uicore.ListMarginPx()+uicore.S(56))
    bx,by,bw,bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = scenes.PointIn(mx,my,bx,by,bw,bh)
    uiwidgets.DrawBackButton(dst, s.backHovered)
}

// --- 内部: scHandleInput → scAdvance → scFlush --------------------------------------

// scHandleInput は“入力→意図(Intent)”へ変換し、描画用のホバー状態を更新します。
func (s *Sim) scHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 6)
    mx, my := ebiten.CursorPosition()
    // ホバー更新
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = scenes.PointIn(mx,my,bx,by,bw,bh)
    sx, sy, sw2, sh2 := scenes.BattleStartButtonRect(s.sw, s.sh)
    s.startHovered = scenes.PointIn(mx,my,sx,sy,sw2,sh2)
    ax, ay, aw, ah := scenes.AutoRunButtonRect(s.sw, s.sh)
    s.autoHovered = scenes.PointIn(mx,my,ax,ay,aw,ah)
    s.attHover, s.defHover = -1, -1
    for i:=0; i<3; i++ {
        tx,ty,tw,th := uiwidgets.TerrainButtonRect(s.sw, s.sh, true, i)
        if scenes.PointIn(mx,my,tx,ty,tw,th) { s.attHover = i }
        dx,dy,dw,dh := uiwidgets.TerrainButtonRect(s.sw, s.sh, false, i)
        if scenes.PointIn(mx,my,dx,dy,dw,dh) { s.defHover = i }
    }

    if ctx != nil && ctx.Input != nil {
        if s.logPopup {
            if ctx.Input.Press(gamesvc.Confirm) { intents = append(intents, Intent{Kind: intentCloseLog}) }
            return intents
        }
        if ctx.Input.Press(gamesvc.Cancel) { intents = append(intents, Intent{Kind: intentBack}) }
        if s.startHovered && (s.simAtk.HP>0 && s.simDef.HP>0) && ctx.Input.Press(gamesvc.Confirm) {
            intents = append(intents, Intent{Kind: intentRunOne})
        }
        if s.autoHovered && ctx.Input.Press(gamesvc.Confirm) { intents = append(intents, Intent{Kind: intentToggleAuto}) }
        if s.attHover >= 0 && ctx.Input.Press(gamesvc.Confirm) { intents = append(intents, Intent{Kind: intentSetTerrainAtt, Index: s.attHover}) }
        if s.defHover >= 0 && ctx.Input.Press(gamesvc.Confirm) { intents = append(intents, Intent{Kind: intentSetTerrainDef, Index: s.defHover}) }
        // キーショートカット（1/2/3, Shift+1/2/3）
        if ctx.Input.Press(gamesvc.TerrainAtt1) { intents = append(intents, Intent{Kind: intentSetTerrainAtt, Index: 0}) }
        if ctx.Input.Press(gamesvc.TerrainAtt2) { intents = append(intents, Intent{Kind: intentSetTerrainAtt, Index: 1}) }
        if ctx.Input.Press(gamesvc.TerrainAtt3) { intents = append(intents, Intent{Kind: intentSetTerrainAtt, Index: 2}) }
        if ctx.Input.Press(gamesvc.TerrainDef1) { intents = append(intents, Intent{Kind: intentSetTerrainDef, Index: 0}) }
        if ctx.Input.Press(gamesvc.TerrainDef2) { intents = append(intents, Intent{Kind: intentSetTerrainDef, Index: 1}) }
        if ctx.Input.Press(gamesvc.TerrainDef3) { intents = append(intents, Intent{Kind: intentSetTerrainDef, Index: 2}) }
    }
    return intents
}

// scAdvance は意図を解釈して状態機械を前進させ、副作用（戦闘実行/ログ切替/遷移フラグ）を反映します。
func (s *Sim) scAdvance(intents []scenes.Intent) {
    for _, any := range intents {
        it, ok := any.(Intent); if !ok { continue }
        switch it.Kind {
        case intentCloseLog:
            s.logPopup = false
        case intentBack:
            s.pop = true
        case intentRunOne:
            leftFirst := (s.turn%2==1)
            s.runOne(leftFirst)
        case intentToggleAuto:
            s.auto = !s.auto
            if s.auto { s.logPopup = false }
        case intentSetTerrainAtt:
            s.attSel = it.Index
            switch it.Index { case 0: s.attTerrain = gcore.Terrain{}; case 1: s.attTerrain = gcore.Terrain{Avoid:20,Def:1}; case 2: s.attTerrain = gcore.Terrain{Avoid:15,Def:2} }
        case intentSetTerrainDef:
            s.defSel = it.Index
            switch it.Index { case 0: s.defTerrain = gcore.Terrain{}; case 1: s.defTerrain = gcore.Terrain{Avoid:20,Def:1}; case 2: s.defTerrain = gcore.Terrain{Avoid:15,Def:2} }
        }
    }
    // 自動実行
    canStart := s.simAtk.HP>0 && s.simDef.HP>0
    leftFirst := (s.turn%2==1)
    if s.auto && canStart && !s.logPopup {
        if s.autoCD>0 { s.autoCD-- } else { s.runOne(leftFirst); s.autoCD=10 }
    }
}

// scFlush はフレーム末尾の副作用処理用フックです（現状なし）。
func (s *Sim) scFlush(_ *game.Ctx) { /* 今はなし */ }
