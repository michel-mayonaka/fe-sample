package layout

import (
	uicore "ui_sample/internal/game/service/ui"
)

// BattleStartButtonRect はバトル開始ボタンの矩形を返します。
func BattleStartButtonRect(sw, sh int) (x, y, w, h int) {
	w, h = uicore.SimStartBtnWPx(), uicore.SimStartBtnHPx()
	x = (sw - w) / 2
	y = sh - uicore.ListMarginPx() - h
	return
}

// AutoRunButtonRect は開始ボタンの右隣（同サイズ/間隔S(20)）。
func AutoRunButtonRect(sw, sh int) (int, int, int, int) {
	bx, by, bw, bh := BattleStartButtonRect(sw, sh)
	gap := uicore.SimAutoRunGapPx()
	return bx + bw + gap, by, bw, bh
}
