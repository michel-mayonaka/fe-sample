package repo

import (
	"ui_sample/internal/model"
)

// ItemsRepo はアイテム定義テーブルへの最小アクセスインタフェースです。
//
// 現状はテーブル全体の参照と再読み込みのみを提供し、
// 名前やIDでの検索は model.ItemDefTable 側に委譲します。
type ItemsRepo interface {
	// Table は内部キャッシュされたアイテム定義テーブルへの参照を返します。
	Table() *model.ItemDefTable
	// Reload は基のJSONを再読み込みし、キャッシュを更新します。
	Reload() error
}

// JSONItemsRepo は JSON から読み込む簡易実装です（起動時に1度ロード）。
type JSONItemsRepo struct {
	path  string
	table *model.ItemDefTable
}

// NewJSONItemsRepo はアイテム定義をロードしてキャッシュします。
func NewJSONItemsRepo(path string) (*JSONItemsRepo, error) {
	it, err := model.LoadItemsJSON(path)
	if err != nil {
		return nil, err
	}
	return &JSONItemsRepo{path: path, table: it}, nil
}

// Table は内部のアイテム定義テーブルへの参照を返します。
func (r *JSONItemsRepo) Table() *model.ItemDefTable { return r.table }

// Reload はアイテム定義JSONを再読み込みします。
func (r *JSONItemsRepo) Reload() error {
	it, err := model.LoadItemsJSON(r.path)
	if err != nil {
		return err
	}
	r.table = it
	return nil
}
