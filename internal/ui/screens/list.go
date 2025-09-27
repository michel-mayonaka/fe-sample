package uiscreens

import (
	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"ui_sample/internal/ui/core"
)

func ListItemRect(sw, _ int, i int) (x, y, w, h int) {
	panelX, panelY := uicore.ListMargin, uicore.ListMargin
	panelW := sw - uicore.ListMargin*2
	startY := panelY + uicore.ListTitleOffset + 32
	y = startY + i*(uicore.ListItemH+uicore.ListItemGap)
	return panelX + 16, y, panelW - 32, uicore.ListItemH
}

func DrawCharacterList(dst *ebiten.Image, units []uicore.Unit, hover int) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	uicore.DrawPanel(dst, float32(uicore.ListMargin), float32(uicore.ListMargin), float32(sw-2*uicore.ListMargin), float32(sh-2*uicore.ListMargin))
	text.Draw(dst, "ユニット一覧", uicore.FaceTitle, uicore.ListMargin+20, uicore.ListMargin+uicore.ListTitleOffset, uicore.ColAccent)
	for i, u := range units {
		x, y, w, h := ListItemRect(sw, sh, i)
		bg := color.RGBA{30, 45, 78, 255}
		if i == hover {
			bg = color.RGBA{40, 60, 100, 255}
		}
		vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
		vector.DrawFilledRect(dst, float32(x-2), float32(y-2), float32(w+4), float32(h+4), uicore.ColBorder, false)
		px := float32(x + 12)
		py := float32(y + (h-uicore.ListPortraitSz)/2)
		uicore.DrawFramedRect(dst, px-2, py-2, uicore.ListPortraitSz+4, uicore.ListPortraitSz+4)
		if u.Portrait != nil {
			uicore.DrawPortrait(dst, u.Portrait, px, py, uicore.ListPortraitSz, uicore.ListPortraitSz)
		} else {
			uicore.DrawPortraitPlaceholder(dst, px, py, uicore.ListPortraitSz, uicore.ListPortraitSz)
		}
		tx := x + 12 + uicore.ListPortraitSz + 20
		ty := y + 36
		text.Draw(dst, u.Name, uicore.FaceMain, tx, ty, uicore.ColText)
		text.Draw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceSmall, tx, ty+26, uicore.ColAccent)
	}
}
