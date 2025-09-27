package uicore

import (
	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
	"math"
)

func DrawPortraitPlaceholder(dst *ebiten.Image, x, y, _, h float32) {
	text.Draw(dst, "画像なし", FaceSmall, int(x+10), int(y+h/2), ColAccent)
}

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
