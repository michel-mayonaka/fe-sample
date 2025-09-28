package scenes

import (
    uicore "ui_sample/internal/game/service/ui"
    gcore "ui_sample/pkg/game"
)

// BattlePort は戦闘解決（本番）を実行するユースケース境界です。
type BattlePort interface {
    // RunBattleRound は選択中ユニットを基点に1ラウンド解決します。
    RunBattleRound(units []uicore.Unit, selIndex int, attT, defT gcore.Terrain) ([]uicore.Unit, []string, bool, error)
}
