package scenes

import (
	uicore "ui_sample/internal/game/service/ui"
)

// DataPort はデータ再読み込み・保存など横断的操作の境界です。
type DataPort interface {
	// ReloadData は参照テーブル等を再読み込みします。
	ReloadData() error
	// PersistUnit はUIユニットの最新状態を保存します。
	PersistUnit(u uicore.Unit) error
}
