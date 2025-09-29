package scenes

import (
	lvl "ui_sample/internal/game/service/levelup"
	uicore "ui_sample/internal/game/service/ui"
)

// Session は UI シーン間で共有される“表示用の一時状態”をまとめた構造体です。
// - 論理データの保存やリポジトリアクセスは UseCases に委譲します。
// - ここは UI の選択状態やポップアップ表示状態など、描画に近い値のみを保持します。
type Session struct {
	// 一覧/選択
	Units    []uicore.Unit
	SelIndex int

	// ステータス/在庫で共有する状態
	PopupActive       bool
	PopupGains        lvl.Gains
	PopupJustOpened   bool
	CurrentSlot       int
	SelectingEquip    bool
	SelectingIsWeapon bool
	InvTab            int // 0=武器,1=アイテム
	HoverInv          int
}

// Selected は現在選択中のユニットを返します（範囲外は先頭、未設定はサンプル）。
func (s *Session) Selected() uicore.Unit {
	if s == nil || len(s.Units) == 0 {
		return uicore.SampleUnit()
	}
	i := s.SelIndex
	if i < 0 || i >= len(s.Units) {
		i = 0
	}
	return s.Units[i]
}

// SetSelected は現在選択中のユニットを書き換えます（範囲外は無視）。
func (s *Session) SetSelected(u uicore.Unit) {
	if s == nil || len(s.Units) == 0 {
		return
	}
	if s.SelIndex < 0 || s.SelIndex >= len(s.Units) {
		return
	}
	s.Units[s.SelIndex] = u
}
