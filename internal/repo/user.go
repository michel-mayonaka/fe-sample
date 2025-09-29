package repo

import (
    usr "ui_sample/internal/model/user"
    "ui_sample/internal/infra/userfs"
)

// UserRepo はユーザセーブデータへの最小アクセスを抽象化します。
type UserRepo interface {
    Find(id string) (usr.Character, bool)
    Update(c usr.Character)
    Save() error
    Table() *usr.Table
}

// JSONUserRepo は JSON バックエンドの簡易実装です（ロード/セーブ）。
type JSONUserRepo struct {
    path string
    t    *usr.Table
}

// NewJSONUserRepo は JSON を読み込み、テーブルをキャッシュします。
func NewJSONUserRepo(path string) (*JSONUserRepo, error) {
    ut, err := userfs.LoadTableJSON(path)
    if err != nil {
        return nil, err
    }
    return &JSONUserRepo{path: path, t: ut}, nil
}

// Find はIDでユーザキャラクターを検索します。
func (r *JSONUserRepo) Find(id string) (usr.Character, bool) {
    if r == nil || r.t == nil { return usr.Character{}, false }
    return r.t.Find(id)
}

// Update はキャラクターを更新します（ID一致時）。
func (r *JSONUserRepo) Update(c usr.Character) {
    if r == nil || r.t == nil { return }
    r.t.UpdateCharacter(c)
}

// Save はテーブルを元のJSONへ保存します。
func (r *JSONUserRepo) Save() error {
    if r == nil || r.t == nil { return nil }
    return userfs.SaveTableJSON(r.path, r.t)
}

// Table は内部テーブルへの参照を返します。
func (r *JSONUserRepo) Table() *usr.Table { return r.t }
