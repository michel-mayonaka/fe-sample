package layout

import (
    uicore "ui_sample/internal/game/service/ui"
)

// ChooseUnitItemRect はポップアップ内の i 番目のキャラ項目の矩形（グリッド）を返します。
func ChooseUnitItemRect(sw, sh, i, _ int) (x, y, w, h int) {
    pw, ph := PopupSize(sw, sh)
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    // グリッド設定
    cols := 5
    if pw < uicore.S(900) { cols = 4 }
    if pw < uicore.S(700) { cols = 3 }
    cellW := (pw - uicore.S(48)) / cols // 左右マージン+隙間
    cellH := uicore.S(160)
    gap := uicore.S(8)
    gridX := px + uicore.S(24)
    gridY := py + uicore.S(88)
    col := i % cols
    row := i / cols
    x = gridX + col*(cellW+gap)
    y = gridY + row*(cellH+gap)
    w = cellW
    h = cellH
    return
}

// PopupSize は選択ポップアップの幅・高さを返します。
func PopupSize(sw, sh int) (w, h int) {
    w = int(float32(sw) * 0.8)
    h = int(float32(sh) * 0.72)
    if w > uicore.S(1480) { w = uicore.S(1480) }
    if h > uicore.S(900) { h = uicore.S(900) }
    return
}

