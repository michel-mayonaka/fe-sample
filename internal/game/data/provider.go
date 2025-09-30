// Package data はUI層（scenes）が参照する読み取り専用テーブル群のプロバイダを提供します。
package data

import (
	"ui_sample/internal/model"
	usr "ui_sample/internal/model/user"
)

// TableProvider はシーンから参照するテーブル群の最小ポートです。
// 将来的に Items/Classes 等を追加可能ですが、まずは武器のみ。
type TableProvider interface {
    WeaponsTable() *model.WeaponTable
    ItemsTable() *model.ItemDefTable
    // UserInventory: ユーザ在庫の参照（スナップショット）。
    UserWeapons() []usr.OwnWeapon
    UserItems() []usr.OwnItem
    // UserTable: ユーザテーブル参照（読み取り用途）。
    UserTable() *usr.Table
    // EquipKindAt: 指定スロットの種別（武器/アイテムのどちらが入っているか）を返す。
    EquipKindAt(unitID string, slot int) (hasWeapon bool, hasItem bool)
}

var provider TableProvider

// SetProvider はアプリケーション側の実装（例: *app.App）を注入します。
func SetProvider(p TableProvider) { provider = p }

// Provider は現在のテーブルプロバイダを返します（未設定時は nil）。
func Provider() TableProvider { return provider }
