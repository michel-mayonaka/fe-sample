package adapter

import (
    uicore "ui_sample/internal/game/service/ui"
    usr "ui_sample/internal/model/user"
)

// init で uicore 側のブリッジ関数に adapter 実装を登録します。
func init() {
    uicore.UnitFromUserFunc = func(c usr.Character) uicore.Unit {
        return UnitFromUser(c, AssetsPortraitLoader{})
    }
    uicore.BuildUnitsFromProviderFunc = func() []uicore.Unit {
        return BuildUnitsFromProvider(AssetsPortraitLoader{})
    }
}

