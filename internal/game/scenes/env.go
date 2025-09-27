package scenes

import (
    "math/rand"

    app "ui_sample/internal/app"
    "ui_sample/internal/ui"
    "ui_sample/internal/user"
)

// Env は UI シーン間で共有する状態とユースケースを束ねます。
type Env struct {
    App       *app.App
    UserTable *user.Table
    UserPath  string
    RNG       *rand.Rand

    Units    []ui.Unit
    SelIndex int

    // ステータス/在庫で共有する状態
    PopupActive     bool
    PopupGains      ui.LevelUpGains
    PopupJustOpened bool
    CurrentSlot     int
    SelectingEquip  bool
    SelectingIsWeapon bool
    InvTab          int // 0=武器,1=アイテム
    HoverInv        int
}

func (e *Env) Selected() ui.Unit {
    if e == nil || len(e.Units) == 0 { return ui.SampleUnit() }
    if e.SelIndex < 0 || e.SelIndex >= len(e.Units) { return e.Units[0] }
    return e.Units[e.SelIndex]
}

func (e *Env) SetSelected(u ui.Unit) {
    if e == nil || len(e.Units) == 0 { return }
    if e.SelIndex < 0 || e.SelIndex >= len(e.Units) { return }
    e.Units[e.SelIndex] = u
}

