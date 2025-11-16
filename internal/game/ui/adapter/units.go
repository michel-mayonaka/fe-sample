package adapter

import (
    gdata "ui_sample/internal/game/data"
    uicore "ui_sample/internal/game/service/ui"
    usr "ui_sample/internal/model/user"
)

// BuildUnitsFromUserTable はユーザテーブルから UI 用ユニット配列を構築します。
func BuildUnitsFromUserTable(ut *usr.Table, pl PortraitLoader) []uicore.Unit {
    if ut == nil {
        return nil
    }
    rows := ut.Slice()
    units := make([]uicore.Unit, 0, len(rows))
    for _, c := range rows {
        units = append(units, UnitFromUser(c, pl))
    }
    return units
}

// BuildUnitsFromProvider は Provider 経由でユーザテーブルを取得し、UI 用ユニット配列を構築します。
func BuildUnitsFromProvider(pl PortraitLoader) []uicore.Unit {
    p := gdata.Provider()
    if p == nil {
        return nil
    }
    return BuildUnitsFromUserTable(p.UserTable(), pl)
}

