package sim

import (
    uicore "ui_sample/internal/game/service/ui"
)

// BattleStartButtonRect はバトル開始ボタンの矩形を返します。
func BattleStartButtonRect(sw, sh int) (x, y, w, h int) {
    w, h = uicore.S(240), uicore.S(60)
    x = (sw - w) / 2
    y = sh - uicore.ListMarginPx() - h
    return
}

// AutoRunButtonRect は開始ボタンの右隣（同サイズ/間隔S(20)）の矩形を返します。
func AutoRunButtonRect(sw, sh int) (int, int, int, int) {
    bx, by, bw, bh := BattleStartButtonRect(sw, sh)
    gap := uicore.S(20)
    return bx + bw + gap, by, bw, bh
}

