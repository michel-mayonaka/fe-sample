package sqlite

import (
    "database/sql"
    "ui_sample/internal/model"
    "ui_sample/internal/repo"
)

// Ensure implementation
var _ repo.WeaponsRepo = (*WeaponsRepo)(nil)

// WeaponsRepo は SQLite 版の武器定義リポジトリ（スケルトン）です。
type WeaponsRepo struct {
    DB *sql.DB
    // 暫定互換: テーブルキャッシュ
    table *model.WeaponTable
}

func NewWeaponsRepo(db *sql.DB) *WeaponsRepo { return &WeaponsRepo{DB: db} }

func (r *WeaponsRepo) Find(name string) (model.Weapon, bool) {
    if r.table == nil { return model.Weapon{}, false }
    return r.table.Find(name)
}

func (r *WeaponsRepo) Table() *model.WeaponTable { return r.table }

func (r *WeaponsRepo) Reload() error { return nil }

