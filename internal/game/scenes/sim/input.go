package sim

import (
    "github.com/hajimehoshi/ebiten/v2"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    scenes "ui_sample/internal/game/scenes"
    uilayout "ui_sample/internal/game/ui/layout"
    "ui_sample/pkg/game/geom"
)

// scHandleInput は“入力→意図(Intent)”へ変換し、描画用のホバー状態を更新します。
func (s *Sim) scHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 6)
    if s.logPopup { return intents }
    mx, my := ebiten.CursorPosition()
    // ホバー更新
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = geom.RectContains(mx,my,bx,by,bw,bh)
    sx, sy, sw2, sh2 := uilayout.BattleStartButtonRect(s.sw, s.sh)
    s.startHovered = geom.RectContains(mx,my,sx,sy,sw2,sh2)
    ax, ay, aw, ah := uilayout.AutoRunButtonRect(s.sw, s.sh)
    s.autoHovered = geom.RectContains(mx,my,ax,ay,aw,ah)
    s.attHover, s.defHover = -1, -1
    for i:=0; i<3; i++ {
        tx,ty,tw,th := uiwidgets.TerrainButtonRect(s.sw, s.sh, true, i)
        if geom.RectContains(mx,my,tx,ty,tw,th) { s.attHover = i }
        dx,dy,dw,dh := uiwidgets.TerrainButtonRect(s.sw, s.sh, false, i)
        if geom.RectContains(mx,my,dx,dy,dw,dh) { s.defHover = i }
    }

    if ctx != nil && ctx.Input != nil {
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
