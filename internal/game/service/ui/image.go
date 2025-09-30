//go:build !headless

package uicore

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

// DrawPortraitPlaceholder はポートレート未設定時のプレースホルダーを描画します。
func DrawPortraitPlaceholder(dst *ebiten.Image, x, y, _, h float32) {
	TextDraw(dst, "画像なし", FaceSmall, int(x+10), int(y+h/2), ColAccent)
}

// DrawPortrait は画像を枠内に等比縮小して描画します。
func DrawPortrait(dst *ebiten.Image, img *ebiten.Image, x, y, w, h float32) {
	b := img.Bounds()
	iw, ih := b.Dx(), b.Dy()
	if iw == 0 || ih == 0 {
		return
	}
	sx := float64(w) / float64(iw)
	sy := float64(h) / float64(ih)
	s := math.Min(sx, sy)
	sw := float64(iw) * s
	sh := float64(ih) * s
	tx := float64(x) + (float64(w)-sw)/2
	ty := float64(y) + (float64(h)-sh)/2
	var op ebiten.DrawImageOptions
	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(s, s)
	op.GeoM.Translate(tx, ty)
	dst.DrawImage(img, &op)
}
