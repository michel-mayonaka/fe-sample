package ui

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "github.com/hajimehoshi/ebiten/v2/vector"
)

func ListItemRect(sw, _ int, i int) (x, y, w, h int) {
    panelX, panelY := listMargin, listMargin
    panelW := sw - listMargin*2
    startY := panelY + listTitleOffset + 32
    y = startY + i*(listItemH+listItemGap)
    return panelX + 16, y, panelW - 32, listItemH
}

func DrawCharacterList(dst *ebiten.Image, units []Unit, hover int) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    drawPanel(dst, float32(listMargin), float32(listMargin), float32(sw-2*listMargin), float32(sh-2*listMargin))
    textDraw(dst, "ユニット一覧", faceTitle, listMargin+20, listMargin+listTitleOffset, colAccent)
    for i, u := range units {
        x, y, w, h := ListItemRect(sw, sh, i)
        bg := color.RGBA{30, 45, 78, 255}
        if i == hover { bg = color.RGBA{40, 60, 100, 255} }
        vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
        vector.DrawFilledRect(dst, float32(x-2), float32(y-2), float32(w+4), float32(h+4), colBorder, false)

        px := float32(x + 12)
        py := float32(y + (h-listPortraitSz)/2)
        drawFramedRect(dst, px-2, py-2, listPortraitSz+4, listPortraitSz+4)
        if u.Portrait != nil { drawPortrait(dst, u.Portrait, px, py, listPortraitSz, listPortraitSz) } else { drawPortraitPlaceholder(dst, px, py, listPortraitSz, listPortraitSz) }

        tx := x + 12 + listPortraitSz + 20
        ty := y + 36
        textDraw(dst, u.Name, faceMain, tx, ty, colText)
        textDraw(dst, u.Class+"  Lv "+itoa(u.Level), faceSmall, tx, ty+26, colAccent)
    }
}

// itoa は簡易的な整数→文字列（少数回呼び出しのため fmt 回避）。
func itoa(n int) string { return fmtInt(n) }

