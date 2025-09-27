package scenes

import (
    "math/rand"

    "ui_sample/internal/model"
    "ui_sample/internal/repo"
    uicore "ui_sample/internal/game/service/ui"
    scpopup "ui_sample/internal/game/scenes/common/popup"
    "ui_sample/internal/user"
    gcore "ui_sample/pkg/game"
)

// AppPort は Scene 層が必要とする最小限のユースケース境界です。
type AppPort interface {
    ReloadData() error
    WeaponsTable() *model.WeaponTable
    PersistUnit(u uicore.Unit) error
    RunBattleRound(units []uicore.Unit, selIndex int, attT, defT gcore.Terrain) ([]uicore.Unit, []string, bool, error)
    Inventory() repo.InventoryRepo
}

// Env は UI シーン間で共有する状態とユースケースを束ねます。
type Env struct {
    App       AppPort
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
