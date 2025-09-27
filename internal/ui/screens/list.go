package uiscreens

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    uicore "ui_sample/internal/ui/core"
)

// ListItemRect は一覧の i 行目の矩形を返します。
func ListItemRect(sw, _ int, i int) (x, y, w, h int) {
    lm := uicore.ListMarginPx()
    panelX, panelY := lm, lm
    panelW := sw - lm*2
    startY := panelY + uicore.ListTitleOffsetPx() + uicore.S(32)
    y = startY + i*(uicore.ListItemHPx()+uicore.ListItemGapPx())
    return panelX + uicore.S(16), y, panelW - uicore.S(32), uicore.ListItemHPx()
}

// DrawCharacterList はユニットの一覧を描画します。
func DrawCharacterList(dst *ebiten.Image, units []uicore.Unit, hover int) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := uicore.ListMarginPx()
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    uicore.TextDraw(dst, "ユニット一覧", uicore.FaceTitle, lm+uicore.S(20), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
    for i, u := range units {
        x, y, w, h := ListItemRect(sw, sh, i)
        bg := color.RGBA{30, 45, 78, 255}
        if i == hover {
            bg = color.RGBA{40, 60, 100, 255}
        }
        vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
        vector.DrawFilledRect(dst, float32(x-uicore.S(2)), float32(y-uicore.S(2)), float32(w+uicore.S(4)), float32(h+uicore.S(4)), uicore.ColBorder, false)
        ps := uicore.ListPortraitSzPx()
        px := float32(x + uicore.S(12))
        py := float32(y + (h-ps)/2)
        uicore.DrawFramedRect(dst, px-float32(uicore.S(2)), py-float32(uicore.S(2)), float32(ps+uicore.S(4)), float32(ps+uicore.S(4)))
        if u.Portrait != nil {
            uicore.DrawPortrait(dst, u.Portrait, px, py, float32(ps), float32(ps))
        } else {
            uicore.DrawPortraitPlaceholder(dst, px, py, float32(ps), float32(ps))
        }
        tx := x + uicore.S(12) + ps + uicore.S(20)
        ty := y + uicore.S(36)
        uicore.TextDraw(dst, u.Name, uicore.FaceMain, tx, ty, uicore.ColText)
        uicore.TextDraw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceSmall, tx, ty+uicore.S(26), uicore.ColAccent)
    }
}
