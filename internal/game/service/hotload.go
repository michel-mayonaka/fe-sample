package service

// Hotloader は開発時のテーブル/マップ等のホットリロード検出用プレースホルダです。
// 実装は別ビルドタグ（例: ebitendebug）で差し替える想定です。
type Hotloader struct{}

// NewHotloader はホットリロード検出のプレースホルダを生成します。
func NewHotloader() *Hotloader { return &Hotloader{} }

// Poll は変更有無を返します（MVP: 常に false）。
func (h *Hotloader) Poll() bool { return false }
