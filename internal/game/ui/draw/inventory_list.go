package draw

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	uicore "ui_sample/internal/game/service/ui"
	uilayout "ui_sample/internal/game/ui/layout"
	uiview "ui_sample/internal/game/ui/view"
)

// DrawItemListView はアイテム一覧（アイテムビュー）を描画します。
func DrawItemListView(dst *ebiten.Image, rows []uiview.ItemRow, hover int) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	lm := uicore.ListMarginPx()
	uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
	// タイトル
	uicore.TextDraw(dst, "アイテム一覧", uicore.FaceTitle, lm+uicore.ListTitleXOffsetPx(), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
	// ヘッダ
	startY := lm + uicore.ListTitleOffsetPx() + uicore.ListHeaderTopGapPx()
	hx := lm + uicore.ListHeaderBaseXPx()
	hy := startY
	colsH := uicore.ListHeaderColumnsItemsPx()
	if len(colsH) >= 5 {
		uicore.TextDraw(dst, "名称", uicore.FaceSmall, hx+colsH[0], hy, uicore.ColText)
		uicore.TextDraw(dst, "種別", uicore.FaceSmall, hx+colsH[1], hy, uicore.ColText)
		uicore.TextDraw(dst, "効果", uicore.FaceSmall, hx+colsH[2], hy, uicore.ColText)
		uicore.TextDraw(dst, "数値", uicore.FaceSmall, hx+colsH[3], hy, uicore.ColText)
		uicore.TextDraw(dst, "耐久", uicore.FaceSmall, hx+colsH[4], hy, uicore.ColText)
	}

	for i, it := range rows {
		x, y, width, h := uilayout.ListItemRect(sw, sh, i)
		bg := color.RGBA{30, 45, 78, 255}
		if i == hover {
			bg = color.RGBA{40, 60, 100, 255}
		}
		vector.DrawFilledRect(dst, float32(x), float32(y), float32(width), float32(h), bg, false)
		bp := uicore.ListRowBorderPadPx()
		vector.DrawFilledRect(dst, float32(x-bp), float32(y-bp), float32(width+2*bp), float32(h+2*bp), uicore.ColBorder, false)

		tx := x + uicore.ListRowTextOffsetXPx()
		ty := y + uicore.ListRowTextOffsetYPx()
		uicore.TextDraw(dst, it.Name, uicore.FaceMain, tx, ty, uicore.ColText)
		colsR := uicore.ListRowColumnsItemsPx()
		if len(colsR) >= 5 {
			// 0: 名称は tx で描画済み
			uicore.TextDraw(dst, it.Type, uicore.FaceSmall, tx+colsR[1], ty, uicore.ColAccent)
			uicore.TextDraw(dst, it.Effect, uicore.FaceSmall, tx+colsR[2], ty, uicore.ColAccent)
			uicore.TextDraw(dst, fmt.Sprintf("%d", it.Power), uicore.FaceSmall, tx+colsR[3], ty, uicore.ColAccent)
			uicore.TextDraw(dst, fmt.Sprintf("%d/%d", it.Uses, it.Max), uicore.FaceSmall, tx+colsR[4], ty, uicore.ColAccent)
		}

		if n := len(it.Owners); n > 0 {
			ob := it.Owners[n-1]
			icon := uicore.ListRowRightIconSizePx()
			ox := x + width - uicore.ListRowRightIconGapPx() - icon
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

// DrawWeaponListView は武器一覧（武器ビュー）を描画します。
func DrawWeaponListView(dst *ebiten.Image, rows []uiview.WeaponRow, hover int) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	lm := uicore.ListMarginPx()
	uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
	// タイトル
	uicore.TextDraw(dst, "武器一覧", uicore.FaceTitle, lm+uicore.ListTitleXOffsetPx(), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
	// ヘッダ
	startY := lm + uicore.ListTitleOffsetPx() + uicore.ListHeaderTopGapPx()
	hx := lm + uicore.ListHeaderBaseXPx()
	hy := startY
	colsH := uicore.ListHeaderColumnsWeaponsPx()
	if len(colsH) >= 9 {
		uicore.TextDraw(dst, "名称", uicore.FaceSmall, hx+colsH[0], hy, uicore.ColText)
		uicore.TextDraw(dst, "種別", uicore.FaceSmall, hx+colsH[1], hy, uicore.ColText)
		uicore.TextDraw(dst, "ﾗﾝｸ", uicore.FaceSmall, hx+colsH[2], hy, uicore.ColText)
		uicore.TextDraw(dst, "威力", uicore.FaceSmall, hx+colsH[3], hy, uicore.ColText)
		uicore.TextDraw(dst, "命中", uicore.FaceSmall, hx+colsH[4], hy, uicore.ColText)
		uicore.TextDraw(dst, "必殺", uicore.FaceSmall, hx+colsH[5], hy, uicore.ColText)
		uicore.TextDraw(dst, "重さ", uicore.FaceSmall, hx+colsH[6], hy, uicore.ColText)
		uicore.TextDraw(dst, "射程", uicore.FaceSmall, hx+colsH[7], hy, uicore.ColText)
		uicore.TextDraw(dst, "耐久", uicore.FaceSmall, hx+colsH[8], hy, uicore.ColText)
	}

	for i, w := range rows {
		x, y, width, h := uilayout.ListItemRect(sw, sh, i)
		bg := color.RGBA{30, 45, 78, 255}
		if i == hover {
			bg = color.RGBA{40, 60, 100, 255}
		}
		vector.DrawFilledRect(dst, float32(x), float32(y), float32(width), float32(h), bg, false)
		bp := uicore.ListRowBorderPadPx()
		vector.DrawFilledRect(dst, float32(x-bp), float32(y-bp), float32(width+2*bp), float32(h+2*bp), uicore.ColBorder, false)

		tx := x + uicore.ListRowTextOffsetXPx()
		ty := y + uicore.ListRowTextOffsetYPx()
		uicore.TextDraw(dst, w.Name, uicore.FaceMain, tx, ty, uicore.ColText)
		colsR := uicore.ListRowColumnsWeaponsPx()
		if len(colsR) >= 9 {
			// 0: 名称は tx で描画済み
			uicore.TextDraw(dst, w.Type, uicore.FaceSmall, tx+colsR[1], ty, uicore.ColAccent)
			uicore.TextDraw(dst, w.Rank, uicore.FaceSmall, tx+colsR[2], ty, uicore.ColAccent)
			uicore.TextDraw(dst, fmt.Sprintf("%d", w.Might), uicore.FaceSmall, tx+colsR[3], ty, uicore.ColAccent)
			uicore.TextDraw(dst, fmt.Sprintf("%d", w.Hit), uicore.FaceSmall, tx+colsR[4], ty, uicore.ColAccent)
			uicore.TextDraw(dst, fmt.Sprintf("%d", w.Crit), uicore.FaceSmall, tx+colsR[5], ty, uicore.ColAccent)
			uicore.TextDraw(dst, fmt.Sprintf("%d", w.Weight), uicore.FaceSmall, tx+colsR[6], ty, uicore.ColAccent)
		}
		rng := fmt.Sprintf("%d", w.RangeMin)
		if w.RangeMax != w.RangeMin {
			rng = fmt.Sprintf("%d-%d", w.RangeMin, w.RangeMax)
		}
		if len(colsR) >= 9 {
			uicore.TextDraw(dst, rng, uicore.FaceSmall, tx+colsR[7], ty, uicore.ColAccent)
			uicore.TextDraw(dst, fmt.Sprintf("%d/%d", w.Uses, w.Max), uicore.FaceSmall, tx+colsR[8], ty, uicore.ColAccent)
		}

		if n := len(w.Owners); n > 0 {
			ob := w.Owners[n-1]
			icon := uicore.ListRowRightIconSizePx()
			ox := x + width - uicore.ListRowRightIconGapPx() - icon
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
