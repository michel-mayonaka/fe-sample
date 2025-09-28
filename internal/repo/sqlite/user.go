// Package sqlite は将来の SQLite バックエンド実装のスケルトンです。
package sqlite

import (
    "context"
    "database/sql"
    "ui_sample/internal/repo"
    "ui_sample/internal/user"
)

// Ensure implementations
var (
    _ repo.UserRepo = (*UserRepo)(nil)
)

// UserRepo は SQLite 版のユーザリポジトリ（スケルトン）です。
type UserRepo struct {
    DB  *sql.DB
    Ctx context.Context
    // cache は最小限の互換用テーブル（暫定）。将来はSELECTに置換。
    cache *user.Table
}

// NewUserRepo は接続を受け取り、スケルトンを返します（未実装）。
func NewUserRepo(ctx context.Context, db *sql.DB) *UserRepo { return &UserRepo{DB: db, Ctx: ctx} }

func (r *UserRepo) Find(id string) (user.Character, bool) {
    if r.cache == nil { return user.Character{}, false }
    return r.cache.Find(id)
}

func (r *UserRepo) Update(c user.Character) {
    if r.cache == nil { return }
    r.cache.UpdateCharacter(c)
}

func (r *UserRepo) Save() error { return nil }

func (r *UserRepo) Table() *user.Table { return r.cache }

