package scenes

import (
    "ui_sample/internal/game/data"
    "ui_sample/internal/model"
)

// 共有武器テーブル（Repo注入で設定）。未設定時は初回アクセスで読み込みキャッシュ。
var wtShared *model.WeaponTable

// SetWeaponTable は内部で共有する武器テーブルを設定します。
func SetWeaponTable(wt *model.WeaponTable) { wtShared = wt }

// WeaponTable は共有武器テーブルを取得します。未設定時はデフォルトJSONから遅延読み込みします。
func WeaponTable() *model.WeaponTable {
    // 1) 推奨: プロバイダ（app）から取得
    if p := data.Provider(); p != nil {
        if wt := p.WeaponsTable(); wt != nil { return wt }
    }
    // 2) 互換: 直接注入済みの共有参照
    if wtShared != nil { return wtShared }
    // 3) フォールバック: 直接JSONを読む（ツール/スタンドアロン用）
    if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
        wtShared = wt
        return wtShared
    }
    return nil
}

// weaponTable はパッケージ内互換のためのエイリアスです（既存コード用）。
func weaponTable() *model.WeaponTable { return WeaponTable() }
