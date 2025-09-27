package service

// UI はテキスト描画やウィジェット薄ラッパのプレースホルダです。
// 現状は internal/ui 側の API を直接利用し、本型は将来の抽象化用に確保します。
type UI struct{}

func NewUI() *UI { return &UI{} }

