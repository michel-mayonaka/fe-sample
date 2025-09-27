package scenes

import (
    "fmt"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    gcore "ui_sample/pkg/game"
)

// Sim シーン（模擬戦）
type Sim struct{
    E *Env
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
}

func NewSim(e *Env, atk, def uicore.Unit) *Sim { return &Sim{E:e, simAtk:atk, simDef:def, turn:1} }
func (s *Sim) ShouldPop() bool { return s.pop }

func (s *Sim) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    // ログポップアップ中は閉じるのみ
    if s.logPopup {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.logPopup=false; return nil,nil }
        if ctx!=nil && ctx.Input!=nil && ctx.Input.Press(gamesvc.Confirm) { s.logPopup=false; return nil,nil }
        return nil,nil
    }
    // 戻る
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    mx, my := ebiten.CursorPosition()
    if pointIn(mx,my,bx,by,bw,bh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.pop=true; return nil,nil }
    if ctx!=nil && ctx.Input!=nil && ctx.Input.Press(gamesvc.Cancel) { s.pop=true; return nil,nil }
    // 実行
    canStart := s.simAtk.HP > 0 && s.simDef.HP > 0
    bx2, by2, bw2, bh2 := BattleStartButtonRect(s.sw, s.sh)
    leftFirst := (s.turn%2==1)
    if canStart && pointIn(mx,my,bx2,by2,bw2,bh2) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        s.runOne(leftFirst)
    }
    if canStart && ctx!=nil && ctx.Input!=nil && ctx.Input.Press(gamesvc.Confirm) { s.runOne(leftFirst) }
    // 自動実行
    if s.auto && canStart && !s.logPopup {
        if s.autoCD>0 { s.autoCD-- } else { s.runOne(leftFirst); s.autoCD=10 }
    }
    // 自動実行トグル
    ax, ay, aw, ah := AutoRunButtonRect(s.sw, s.sh)
    if pointIn(mx,my,ax,ay,aw,ah) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.auto=!s.auto; if s.auto { s.logPopup=false } }
    // 地形ボタン
    for i:=0;i<3;i++{
        ax,ay,aw,ah := uiwidgets.TerrainButtonRect(s.sw, s.sh, true, i)
        if pointIn(mx,my,ax,ay,aw,ah) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.attSel=i; switch i{case 0:s.attTerrain=gcore.Terrain{}; case 1:s.attTerrain=gcore.Terrain{Avoid:20,Def:1}; case 2:s.attTerrain=gcore.Terrain{Avoid:15,Def:2}} }
        dx,dy,dw,dh := uiwidgets.TerrainButtonRect(s.sw, s.sh, false, i)
        if pointIn(mx,my,dx,dy,dw,dh) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) { s.defSel=i; switch i{case 0:s.defTerrain=gcore.Terrain{}; case 1:s.defTerrain=gcore.Terrain{Avoid:20,Def:1}; case 2:s.defTerrain=gcore.Terrain{Avoid:15,Def:2}} }
    }
    return nil,nil
}

func (s *Sim) runOne(leftFirst bool){
    if leftFirst {
        a,d,lines := SimulateBattleCopyWithTerrain(s.simAtk, s.simDef, s.attTerrain, s.defTerrain, s.E.RNG)
        s.simAtk, s.simDef = a,d; s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simAtk.Name)}, lines...)
    } else {
        a,d,lines := SimulateBattleCopyWithTerrain(s.simDef, s.simAtk, s.defTerrain, s.attTerrain, s.E.RNG)
        s.simDef, s.simAtk = a,d; s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simDef.Name)}, lines...)
    }
    s.logPopup=true; s.turn++
}

func (s *Sim) Draw(dst *ebiten.Image){
    canStart := s.simAtk.HP>0 && s.simDef.HP>0 && !s.logPopup
    DrawBattleWithTerrain(dst, s.simAtk, s.simDef, s.attTerrain, s.defTerrain, canStart)
    uiwidgets.DrawTerrainButtons(dst, s.attSel, s.defSel)
    mx,my := ebiten.CursorPosition(); ax,ay,aw,ah := AutoRunButtonRect(s.sw, s.sh)
    uiwidgets.DrawAutoRunButton(dst, pointIn(mx,my,ax,ay,aw,ah), s.auto)
    if s.logPopup { DrawBattleLogOverlay(dst, s.logs) }
    if s.turn<=0 { s.turn=1 }
    leftFirst := (s.turn%2==1); label := "先攻: "; if leftFirst { label+=s.simAtk.Name } else { label+=s.simDef.Name }
    ebitenutil.DebugPrintAt(dst, label, uicore.ListMarginPx()+uicore.S(40), uicore.ListMarginPx()+uicore.S(56))
    bx,by,bw,bh := uiwidgets.BackButtonRect(s.sw, s.sh); uiwidgets.DrawBackButton(dst, pointIn(mx,my,bx,by,bw,bh))
}
