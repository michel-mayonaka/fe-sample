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

// NewWeaponsRepo は武器定義リポジトリのスケルトン実装を返します（未配線）。
func NewWeaponsRepo(db *sql.DB) *WeaponsRepo { return &WeaponsRepo{DB: db} }

// Find は武器名で定義を検索します（スケルトン）。
func (r *WeaponsRepo) Find(name string) (model.Weapon, bool) {
    if r.table == nil { return model.Weapon{}, false }
    return r.table.Find(name)
}

// Table は武器テーブルを返します（スケルトン）。
func (r *WeaponsRepo) Table() *model.WeaponTable { return r.table }

// Reload は定義を再読み込みします（スケルトン）。
func (r *WeaponsRepo) Reload() error { return nil }
