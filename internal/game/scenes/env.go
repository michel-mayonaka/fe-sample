package scenes

import (
    "math/rand"

    app "ui_sample/internal/app"
    uicore "ui_sample/internal/game/service/ui"
    scpopup "ui_sample/internal/game/scenes/common/popup"
    "ui_sample/internal/user"
)

// Env は UI シーン間で共有する状態とユースケースを束ねます。
type Env struct {
    App       *app.App
    UserTable *user.Table
    UserPath  string
    RNG       *rand.Rand

    Units    []uicore.Unit
    SelIndex int

    // ステータス/在庫で共有する状態
    PopupActive     bool
    PopupGains      scpopup.LevelUpGains
    PopupJustOpened bool
    CurrentSlot     int
    SelectingEquip  bool
    SelectingIsWeapon bool
    InvTab          int // 0=武器,1=アイテム
    HoverInv        int
}

func (e *Env) Selected() uicore.Unit {
    if e == nil || len(e.Units) == 0 { return uicore.SampleUnit() }
    if e.SelIndex < 0 || e.SelIndex >= len(e.Units) { return e.Units[0] }
    return e.Units[e.SelIndex]
}

func (e *Env) SetSelected(u uicore.Unit) {
    if e == nil || len(e.Units) == 0 { return }
    if e.SelIndex < 0 || e.SelIndex >= len(e.Units) { return }
    e.Units[e.SelIndex] = u
}
