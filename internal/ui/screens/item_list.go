package uiscreens

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    uicore "ui_sample/internal/ui/core"
)

// ItemRow はユーザ所持アイテム（耐久）+マスタ情報の結合行です。
type ItemRow struct {
    ID               string
    Name, Type       string
    Effect           string
    Power            int
    Uses, Max        int
    Owners           []OwnerBadge
}

// DrawItemList はアイテム一覧を描画します（usr_items.json 由来）。
func DrawItemList(dst *ebiten.Image, rows []ItemRow, hover int) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := uicore.ListMarginPx()
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    // タイトル
    uicore.TextDraw(dst, "アイテム一覧", uicore.FaceTitle, lm+uicore.S(20), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
    // ヘッダ
    startY := lm + uicore.ListTitleOffsetPx() + uicore.S(8)
    hx := lm + uicore.S(36)
    hy := startY
    uicore.TextDraw(dst, "名称", uicore.FaceSmall, hx+uicore.S(0), hy, uicore.ColText)
    uicore.TextDraw(dst, "種別", uicore.FaceSmall, hx+uicore.S(560), hy, uicore.ColText)
    uicore.TextDraw(dst, "効果", uicore.FaceSmall, hx+uicore.S(720), hy, uicore.ColText)
    uicore.TextDraw(dst, "数値", uicore.FaceSmall, hx+uicore.S(900), hy, uicore.ColText)
    uicore.TextDraw(dst, "耐久", uicore.FaceSmall, hx+uicore.S(1000), hy, uicore.ColText)

    for i, it := range rows {
        x, y, width, h := ListItemRect(sw, sh, i)
        bg := color.RGBA{30, 45, 78, 255}
        if i == hover { bg = color.RGBA{40, 60, 100, 255} }
        vector.DrawFilledRect(dst, float32(x), float32(y), float32(width), float32(h), bg, false)
        vector.DrawFilledRect(dst, float32(x-uicore.S(2)), float32(y-uicore.S(2)), float32(width+uicore.S(4)), float32(h+uicore.S(4)), uicore.ColBorder, false)

        tx := x + uicore.S(20)
        ty := y + uicore.S(36)
        uicore.TextDraw(dst, it.Name, uicore.FaceMain, tx, ty, uicore.ColText)
        uicore.TextDraw(dst, it.Type, uicore.FaceSmall, tx+uicore.S(540), ty, uicore.ColAccent)
        uicore.TextDraw(dst, it.Effect, uicore.FaceSmall, tx+uicore.S(700), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", it.Power), uicore.FaceSmall, tx+uicore.S(880), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d/%d", it.Uses, it.Max), uicore.FaceSmall, tx+uicore.S(980), ty, uicore.ColAccent)
        // 所有者アイコン（最大1人）
        if len(it.Owners) > 0 {
            ob := it.Owners[len(it.Owners)-1]
            icon := uicore.S(24)
            ox := x + width - uicore.S(12) - icon
            oy := y + (h-icon)/2
            uicore.DrawFramedRect(dst, float32(ox), float32(oy), float32(icon), float32(icon))
            if ob.Portrait != nil {
                uicore.DrawPortrait(dst, ob.Portrait, float32(ox), float32(oy), float32(icon), float32(icon))
            } else {
                uicore.DrawPortraitPlaceholder(dst, float32(ox), float32(oy), float32(icon), float32(icon))
            }
        }
    }
}
