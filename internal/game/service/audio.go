package service

// Audio は BGM/SE の簡易キューを表すプレースホルダです。
// 具体実装は将来差し替え可能な形を想定します。
type Audio struct{}
// NewAudio はオーディオ管理のプレースホルダを生成します。
func NewAudio() *Audio { return &Audio{} }

// Flush はキューを適用（MVP: 何もしない）。
func (a *Audio) Flush() {}
