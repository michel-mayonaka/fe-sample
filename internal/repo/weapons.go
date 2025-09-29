package repo

import (
	"ui_sample/internal/model"
)

// WeaponsRepo は武器定義へアクセスするための最小インタフェースです。
type WeaponsRepo interface {
	// Find は武器名で定義を取得します。
	Find(name string) (model.Weapon, bool)
	// Table は内部テーブルへの参照を返します（読み取り専用の想定）。
	Table() *model.WeaponTable
	// Reload は基のJSONを再読み込みし、キャッシュを更新します。
	Reload() error
}

// JSONWeaponsRepo は JSON から読み込む簡易実装です（起動時に1度ロード）。
type JSONWeaponsRepo struct {
	path  string
	table *model.WeaponTable
}

// NewJSONWeaponsRepo は武器定義をロードしてキャッシュします。
func NewJSONWeaponsRepo(path string) (*JSONWeaponsRepo, error) {
	wt, err := model.LoadWeaponsJSON(path)
	if err != nil {
		return nil, err
	}
	return &JSONWeaponsRepo{path: path, table: wt}, nil
}

// Find は武器名から定義を検索します。
func (r *JSONWeaponsRepo) Find(name string) (model.Weapon, bool) {
	if r == nil || r.table == nil {
		return model.Weapon{}, false
	}
	return r.table.Find(name)
}

// Table は内部の武器テーブルへの参照を返します。
func (r *JSONWeaponsRepo) Table() *model.WeaponTable { return r.table }

// Reload は武器定義JSONを再読み込みします。
func (r *JSONWeaponsRepo) Reload() error {
	wt, err := model.LoadWeaponsJSON(r.path)
	if err != nil {
		return err
	}
	r.table = wt
	return nil
}
