package sqlite

import (
    "database/sql"
    "ui_sample/internal/repo"
    usr "ui_sample/internal/model/user"
)

// Ensure implementation
var _ repo.InventoryRepo = (*InventoryRepo)(nil)

// InventoryRepo は SQLite 版の在庫リポジトリ（スケルトン）です。
// InventoryRepo は SQLite 版の在庫リポジトリ（スケルトン）です。
// 在庫集約として usr_weapons と usr_items の両方を横断管理します。
type InventoryRepo struct {
    DB *sql.DB
    // 暫定互換: スナップショット
    weapons []usr.OwnWeapon
    items   []usr.OwnItem
}

// NewInventoryRepo は在庫リポジトリのスケルトン実装を返します（未配線）。
func NewInventoryRepo(db *sql.DB) *InventoryRepo { return &InventoryRepo{DB: db} }

// Consume は ID に応じて耐久を減らします（スケルトン）。
func (r *InventoryRepo) Consume(_ string, _ int) error { return nil }
// Save は在庫を保存します（スケルトン）。
func (r *InventoryRepo) Save() error { return nil }
// Reload は在庫を再読み込みします（スケルトン）。
func (r *InventoryRepo) Reload() error { return nil }
// Weapons は所持武器のスナップショットを返します（コピー）。
func (r *InventoryRepo) Weapons() []usr.OwnWeapon { return append([]usr.OwnWeapon(nil), r.weapons...) }
// Items は所持アイテムのスナップショットを返します（コピー）。
func (r *InventoryRepo) Items() []usr.OwnItem { return append([]usr.OwnItem(nil), r.items...) }
