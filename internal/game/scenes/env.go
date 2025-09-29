package scenes

import (
    "math/rand"
)

// Env は UI シーン間で共有する状態とユースケースを束ねます。
type Env struct {
    // Data/Battle/Inv は分割Port（最小依存での参照）。
    Data      DataPort
    Battle    BattlePort
    Inv       InventoryPort
    UserPath  string
    RNG       *rand.Rand

    *Session // UI 間で共有する一時状態（埋め込みでプロモート）
}
