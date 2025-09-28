package sim

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/game"
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
    // サブビュー（ポップアップ）
    lv *LogView
    // ホバー状態
    backHovered bool
    startHovered bool
    autoHovered bool
    attHover int // -1/0..2
    defHover int // -1/0..2
}

func NewSim(e *scenes.Env, atk, def uicore.Unit) *Sim {
    s := &Sim{E:e, simAtk:atk, simDef:def, turn:1, attHover:-1, defHover:-1}
    s.lv = NewLogView(s)
    return s
}
func (s *Sim) ShouldPop() bool { return s.pop }
var _ scContract = (*Sim)(nil)

// Update は状態更新の入口です。
// フロー: scHandleInput → scAdvance → scFlush。次シーンは本シーン内で完結するため nil を返します。
func (s *Sim) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    // ポップアップ優先更新
    if s.logPopup && s.lv != nil { _, _ = s.lv.Update(ctx) }
    intents := s.scHandleInput(ctx)
    s.scAdvance(intents)
    s.scFlush(ctx)
    return nil, nil
}

// Draw は模擬戦の盤面・UI とログポップアップを描画します。
func (s *Sim) Draw(dst *ebiten.Image){
    canStart := s.simAtk.HP>0 && s.simDef.HP>0 && !s.logPopup
    DrawBattleWithTerrain(dst, s.simAtk, s.simDef, s.attTerrain, s.defTerrain, canStart)
    uiwidgets.DrawTerrainButtons(dst, s.attSel, s.defSel)
    mx,my := ebiten.CursorPosition()
    ax,ay,aw,ah := AutoRunButtonRect(s.sw, s.sh)
    s.autoHovered = scenes.PointIn(mx,my,ax,ay,aw,ah)
    uiwidgets.DrawAutoRunButton(dst, s.autoHovered, s.auto)
    if s.logPopup && s.lv != nil { s.lv.Draw(dst) }
    if s.turn<=0 { s.turn=1 }
    leftFirst := (s.turn%2==1); label := "先攻: "; if leftFirst { label+=s.simAtk.Name } else { label+=s.simDef.Name }
    ebitenutil.DebugPrintAt(dst, label, uicore.ListMarginPx()+uicore.S(40), uicore.ListMarginPx()+uicore.S(56))
    bx,by,bw,bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = scenes.PointIn(mx,my,bx,by,bw,bh)
    uiwidgets.DrawBackButton(dst, s.backHovered)
}

// --- 内部: scHandleInput → scAdvance → scFlush --------------------------------------

// scFlush はフレーム末尾の副作用処理用フックです（現状なし）。

func (s *Sim) scFlush(_ *game.Ctx) { /* 今はなし */ }
