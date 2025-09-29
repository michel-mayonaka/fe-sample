package layout

import (
    uicore "ui_sample/internal/game/service/ui"
)

const slotCap = 5

// EquipSlotRect はステータス画面における装備スロット行の矩形を返します（index: 0..4）。
func EquipSlotRect(_ , _ int, index int) (x, y, w, h int) {
    if index < 0 || index >= slotCap { return 0,0,0,0 }
    lm := float32(uicore.ListMarginPx())
    panelX, panelY := lm, lm
    px, py := panelX+float32(uicore.S(24)), panelY+float32(uicore.S(24))
    ph := float32(uicore.S(320))
    equipTitleY := int(py + ph + float32(uicore.S(56)))
    lineY := equipTitleY + uicore.S(30) + index*uicore.S(30)
    x = int(px)
    y = lineY - uicore.S(20)
    w = uicore.S(360)
    h = uicore.S(26)
    return
}

