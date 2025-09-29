package adapter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"ui_sample/internal/assets"
)

// PortraitLoader は名前からポートレート画像を読み込む抽象です。
// テストではモックに差し替え可能です。
type PortraitLoader interface {
	Load(name string) (*ebiten.Image, error)
}

// AssetsPortraitLoader は assets.LoadImage を用いた既定実装です。
type AssetsPortraitLoader struct{}

func (AssetsPortraitLoader) Load(name string) (*ebiten.Image, error) {
	return assets.LoadImage(name)
}
