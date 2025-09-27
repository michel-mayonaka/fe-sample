package service

// Assets は画像/音/フォント等の管理プレースホルダです。
// MVP では UI 層（internal/assets）に委譲し、将来の統合に備えます。
type Assets struct{}

func NewAssets() *Assets { return &Assets{} }

