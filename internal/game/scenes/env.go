package scenes

import (
    "math/rand"

    "ui_sample/internal/repo"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/user"
    gcore "ui_sample/pkg/game"
)

// UseCases は Scene 層が必要とする最小限のユースケース境界です。
// 参照テーブルは gdata.Provider に統一するため、WeaponsTable は含めません。
type UseCases interface {
    ReloadData() error
    PersistUnit(u uicore.Unit) error
    RunBattleRound(units []uicore.Unit, selIndex int, attT, defT gcore.Terrain) ([]uicore.Unit, []string, bool, error)
    Inventory() repo.InventoryRepo
    EquipWeapon(unitID string, slot int, userWeaponID string) error
    EquipItem(unitID string, slot int, userItemID string) error
}

// Env は UI シーン間で共有する状態とユースケースを束ねます。
type Env struct {
    // App は合成UseCases（段階移行用の後方互換）。
    App       UseCases
    // Data/Battle/Inv は分割Port（最小依存での参照）。
    Data      DataPort
    Battle    BattlePort
    Inv       InventoryPort
    UserTable *user.Table
    UserPath  string
    RNG       *rand.Rand

    *Session // UI 間で共有する一時状態（埋め込みでプロモート）
}
