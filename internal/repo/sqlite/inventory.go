package sqlite

import (
    "database/sql"
    "ui_sample/internal/repo"
    "ui_sample/internal/user"
)

// Ensure implementation
var _ repo.InventoryRepo = (*InventoryRepo)(nil)

// InventoryRepo は SQLite 版の在庫リポジトリ（スケルトン）です。
type InventoryRepo struct {
    DB *sql.DB
    // 暫定互換: スナップショット
    weapons []user.OwnWeapon
    items   []user.OwnItem
}

func NewInventoryRepo(db *sql.DB) *InventoryRepo { return &InventoryRepo{DB: db} }

func (r *InventoryRepo) Consume(id string, count int) error { return nil }
func (r *InventoryRepo) Save() error                       { return nil }
func (r *InventoryRepo) Reload() error                     { return nil }
func (r *InventoryRepo) Weapons() []user.OwnWeapon         { return append([]user.OwnWeapon(nil), r.weapons...) }
func (r *InventoryRepo) Items() []user.OwnItem             { return append([]user.OwnItem(nil), r.items...) }

