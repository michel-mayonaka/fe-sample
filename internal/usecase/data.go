package usecase

import (
    "ui_sample/internal/config"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/model"
    usr "ui_sample/internal/model/user"
)

// ReloadData は JSON バックエンドのキャッシュを再読み込みします（UI資産のクリアは呼び出し側で実施）。
func (a *App) ReloadData() error {
	if a == nil {
		return nil
	}
	if a.Weapons != nil {
		if err := a.Weapons.Reload(); err != nil {
			return err
		}
	}
	if a.Inv != nil {
		if err := a.Inv.Reload(); err != nil {
			return err
		}
	}
	return nil
}

// WeaponsTable は共有用の武器テーブル参照を返します（gdata.Provider 用）。
func (a *App) WeaponsTable() *model.WeaponTable {
	if a == nil || a.Weapons == nil {
		return nil
	}
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
	if a == nil || a.Users == nil {
		return nil
	}
	c, ok := a.Users.Find(u.ID)
	if !ok {
		return nil
	}
	c.Level = u.Level
	c.HP = u.HP
	c.HPMax = u.HPMax
	c.Stats = usr.Stats{Str: u.Stats.Str, Mag: u.Stats.Mag, Skl: u.Stats.Skl, Spd: u.Stats.Spd, Lck: u.Stats.Lck, Def: u.Stats.Def, Res: u.Stats.Res, Mov: u.Stats.Mov, Bld: u.Stats.Bld}
	a.Users.Update(c)
	return a.Users.Save()
}

// UserWeapons はユーザ所持武器のスナップショットを返します。
func (a *App) UserWeapons() []usr.OwnWeapon {
	if a == nil || a.Inv == nil {
		return nil
	}
	return a.Inv.Weapons()
}

// UserItems はユーザ所持アイテムのスナップショットを返します。
func (a *App) UserItems() []usr.OwnItem {
	if a == nil || a.Inv == nil {
		return nil
	}
	return a.Inv.Items()
}

// UserTable はユーザテーブル（読み取り用）を返します。
func (a *App) UserTable() *usr.Table {
	if a == nil || a.Users == nil {
		return nil
	}
	return a.Users.Table()
}

// UserUnitByID はユーザキャラクタIDから UI 用ユニットを生成して返します。
// Note: 以前は TableProvider 経由で UI 型を返す `UserUnitByID` を提供していたが、
// Port の UI 依存排除に伴い UI 変換は呼び出し側（UI/adapter）へ移管した。

// EquipKindAt は指定スロットに武器/アイテムのどちらが入っているかを返します。
func (a *App) EquipKindAt(unitID string, slot int) (bool, bool) {
	if a == nil || a.Users == nil {
		return false, false
	}
	c, ok := a.Users.Find(unitID)
	if !ok {
		return false, false
	}
	if slot < 0 || slot >= len(c.Equip) {
		return false, false
	}
	er := c.Equip[slot]
	return er.UserWeaponsID != "", er.UserItemsID != ""
}
