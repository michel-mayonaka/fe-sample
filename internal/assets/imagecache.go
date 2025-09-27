package assets

import (
    "sync"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// ImageCache はパス→画像の単純キャッシュです。
type ImageCache struct {
    mu sync.RWMutex
    m  map[string]*ebiten.Image
}

func NewImageCache() *ImageCache { return &ImageCache{m: make(map[string]*ebiten.Image)} }

var DefaultImageCache = NewImageCache()

// LoadImage はキャッシュ経由で画像を取得します。
func LoadImage(path string) (*ebiten.Image, error) {
    if path == "" { return nil, nil }
    DefaultImageCache.mu.RLock()
    if img, ok := DefaultImageCache.m[path]; ok {
        DefaultImageCache.mu.RUnlock()
        return img, nil
    }
    DefaultImageCache.mu.RUnlock()
    img, _, err := ebitenutil.NewImageFromFile(path)
    if err != nil { return nil, err }
    DefaultImageCache.mu.Lock()
    DefaultImageCache.m[path] = img
    DefaultImageCache.mu.Unlock()
    return img, nil
}

