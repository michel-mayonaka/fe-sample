package scenes

import (
    "ui_sample/internal/model"
)

// 共有武器テーブル（Repo注入で設定）。未設定時は初回アクセスで読み込みキャッシュ。
var wtShared *model.WeaponTable

// SetWeaponTable は内部で共有する武器テーブルを設定します。
func SetWeaponTable(wt *model.WeaponTable) { wtShared = wt }

// WeaponTable は共有武器テーブルを取得します。未設定時はデフォルトJSONから遅延読み込みします。
func WeaponTable() *model.WeaponTable {
    if wtShared != nil { return wtShared }
    if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
        wtShared = wt
        return wtShared
    }
    return nil
}

// weaponTable はパッケージ内互換のためのエイリアスです（既存コード用）。
func weaponTable() *model.WeaponTable { return WeaponTable() }

