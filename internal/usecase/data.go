package usecase

import (
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/model"
    "ui_sample/internal/user"
    "ui_sample/internal/config"
)

// ReloadData は JSON バックエンドのキャッシュを再読み込みします（UI資産のクリアは呼び出し側で実施）。
func (a *App) ReloadData() error {
    if a == nil { return nil }
    if a.Weapons != nil { if err := a.Weapons.Reload(); err != nil { return err } }
    if a.Inv != nil { if err := a.Inv.Reload(); err != nil { return err } }
    return nil
}

// WeaponsTable は共有用の武器テーブル参照を返します（gdata.Provider 用）。
func (a *App) WeaponsTable() *model.WeaponTable {
    if a == nil || a.Weapons == nil { return nil }
    return a.Weapons.Table()
}

// ItemsTable はアイテム定義テーブルを返します（軽量用途: JSONから都度ロード）。
// 将来的にキャッシュやRepo経由に最適化可能です。
func (a *App) ItemsTable() *model.ItemDefTable {
    it, _ := model.LoadItemsJSON(config.DefaultItemsPath)
    return it
}

// PersistUnit は UI ユニットの現在値をユーザセーブへ反映して保存します。
func (a *App) PersistUnit(u uicore.Unit) error {
    if a == nil || a.Users == nil { return nil }
    c, ok := a.Users.Find(u.ID)
    if !ok { return nil }
    c.Level = u.Level
    c.HP = u.HP
    c.HPMax = u.HPMax
    c.Stats = user.Stats{Str: u.Stats.Str, Mag: u.Stats.Mag, Skl: u.Stats.Skl, Spd: u.Stats.Spd, Lck: u.Stats.Lck, Def: u.Stats.Def, Res: u.Stats.Res, Mov: u.Stats.Mov, Bld: u.Stats.Bld}
    a.Users.Update(c)
    return a.Users.Save()
}

