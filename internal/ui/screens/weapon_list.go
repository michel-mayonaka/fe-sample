package uiscreens

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    uicore "ui_sample/internal/ui/core"
)

// OwnerBadge は所有者名とポートレートの簡易表示情報です。
type OwnerBadge struct {
    Name     string
    Portrait *ebiten.Image
}

// WeaponRow はユーザ所持武器（耐久含む）+マスタ情報の結合行です。
type WeaponRow struct {
    ID                string
    Name, Type, Rank  string
    Might, Hit, Crit  int
    Weight            int
    RangeMin, RangeMax int
    Uses, Max         int
    Owners            []OwnerBadge
}

// DrawWeaponList は武器一覧を描画します（usr_weapons.json 由来）。
func DrawWeaponList(dst *ebiten.Image, rows []WeaponRow, hover int) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    lm := uicore.ListMarginPx()
    uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
    // タイトル
    uicore.TextDraw(dst, "武器一覧", uicore.FaceTitle, lm+uicore.S(20), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
    // ヘッダ行
    startY := lm + uicore.ListTitleOffsetPx() + uicore.S(8)
    hx := lm + uicore.S(36)
    hy := startY
    uicore.TextDraw(dst, "名称", uicore.FaceSmall, hx+uicore.S(0), hy, uicore.ColText)
    uicore.TextDraw(dst, "種別", uicore.FaceSmall, hx+uicore.S(560), hy, uicore.ColText)
    uicore.TextDraw(dst, "ﾗﾝｸ", uicore.FaceSmall, hx+uicore.S(680), hy, uicore.ColText)
    uicore.TextDraw(dst, "威力", uicore.FaceSmall, hx+uicore.S(760), hy, uicore.ColText)
    uicore.TextDraw(dst, "命中", uicore.FaceSmall, hx+uicore.S(840), hy, uicore.ColText)
    uicore.TextDraw(dst, "必殺", uicore.FaceSmall, hx+uicore.S(920), hy, uicore.ColText)
    uicore.TextDraw(dst, "重さ", uicore.FaceSmall, hx+uicore.S(1000), hy, uicore.ColText)
    uicore.TextDraw(dst, "射程", uicore.FaceSmall, hx+uicore.S(1080), hy, uicore.ColText)
    uicore.TextDraw(dst, "耐久", uicore.FaceSmall, hx+uicore.S(1160), hy, uicore.ColText)

    // 各行
    for i, w := range rows {
        x, y, width, h := ListItemRect(sw, sh, i)
        // 背景
        bg := color.RGBA{30, 45, 78, 255}
        if i == hover { bg = color.RGBA{40, 60, 100, 255} }
        vector.DrawFilledRect(dst, float32(x), float32(y), float32(width), float32(h), bg, false)
        vector.DrawFilledRect(dst, float32(x-uicore.S(2)), float32(y-uicore.S(2)), float32(width+uicore.S(4)), float32(h+uicore.S(4)), uicore.ColBorder, false)

        // テキスト
        tx := x + uicore.S(20)
        ty := y + uicore.S(36)
        // 名称
        uicore.TextDraw(dst, w.Name, uicore.FaceMain, tx, ty, uicore.ColText)
        // 列
        uicore.TextDraw(dst, w.Type, uicore.FaceSmall, tx+uicore.S(540), ty, uicore.ColAccent)
        uicore.TextDraw(dst, w.Rank, uicore.FaceSmall, tx+uicore.S(660), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Might), uicore.FaceSmall, tx+uicore.S(750), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Hit), uicore.FaceSmall, tx+uicore.S(830), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Crit), uicore.FaceSmall, tx+uicore.S(910), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d", w.Weight), uicore.FaceSmall, tx+uicore.S(990), ty, uicore.ColAccent)
        rng := fmt.Sprintf("%d", w.RangeMin)
        if w.RangeMax != w.RangeMin { rng = fmt.Sprintf("%d-%d", w.RangeMin, w.RangeMax) }
        uicore.TextDraw(dst, rng, uicore.FaceSmall, tx+uicore.S(1070), ty, uicore.ColAccent)
        uicore.TextDraw(dst, fmt.Sprintf("%d/%d", w.Uses, w.Max), uicore.FaceSmall, tx+uicore.S(1150), ty, uicore.ColAccent)
        // 所有者アイコン（最大1人）
        if len(w.Owners) > 0 {
            ob := w.Owners[len(w.Owners)-1]
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
