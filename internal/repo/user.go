package repo

import (
    "ui_sample/internal/user"
)

// UserRepo はユーザセーブデータへの最小アクセスを抽象化します。
type UserRepo interface {
    Find(id string) (user.Character, bool)
    Update(c user.Character)
    Save() error
    Table() *user.Table
}

// JSONUserRepo は JSON バックエンドの簡易実装です（ロード/セーブ）。
type JSONUserRepo struct {
    path string
    t    *user.Table
}

// NewJSONUserRepo は JSON を読み込み、テーブルをキャッシュします。
func NewJSONUserRepo(path string) (*JSONUserRepo, error) {
    ut, err := user.LoadFromJSON(path)
    if err != nil {
        return nil, err
    }
    return &JSONUserRepo{path: path, t: ut}, nil
}

func (r *JSONUserRepo) Find(id string) (user.Character, bool) {
    if r == nil || r.t == nil { return user.Character{}, false }
    return r.t.Find(id)
}

func (r *JSONUserRepo) Update(c user.Character) {
    if r == nil || r.t == nil { return }
    r.t.UpdateCharacter(c)
}

func (r *JSONUserRepo) Save() error {
    if r == nil || r.t == nil { return nil }
    return r.t.Save(r.path)
}

func (r *JSONUserRepo) Table() *user.Table { return r.t }

