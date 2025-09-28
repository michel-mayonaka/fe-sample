package scenes

import (
    uicore "ui_sample/internal/game/service/ui"
)

// ListItemRect は一覧の i 行目の矩形を返します。
// 他パッケージ（在庫画面など）からも参照されるため、
// シーン共通のユーティリティとして公開します。
func ListItemRect(sw, _ int, i int) (x, y, w, h int) {
    lm := uicore.ListMarginPx()
    panelX, panelY := lm, lm
    panelW := sw - lm*2
    startY := panelY + uicore.ListTitleOffsetPx() + uicore.S(32)
    y = startY + i*(uicore.ListItemHPx()+uicore.ListItemGapPx())
    return panelX + uicore.S(16), y, panelW - uicore.S(32), uicore.ListItemHPx()
}

