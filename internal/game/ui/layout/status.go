package layout

import (
	uicore "ui_sample/internal/game/service/ui"
)

const slotCap = 5

// EquipSlotRect はステータス画面における装備スロット行の矩形を返します（index: 0..4）。
func EquipSlotRect(_, _ int, index int) (x, y, w, h int) {
	if index < 0 || index >= slotCap {
		return 0, 0, 0, 0
	}
	lm := float32(uicore.ListMarginPx())
	panelX, panelY := lm, lm
	px, py := panelX+float32(uicore.StatusPanelPadPx()), panelY+float32(uicore.StatusPanelPadPx())
	ph := float32(uicore.StatusPortraitSizePx())
	equipTitleY := int(py + ph + float32(uicore.StatusEquipTitleGapYPx()))
	lineY := equipTitleY + uicore.StatusEquipLineHPx() + index*uicore.StatusEquipLineHPx()
	x = int(px)
	y = lineY - uicore.StatusEquipRectYOffsetPx()
	w = uicore.StatusEquipRectWPx()
	h = uicore.StatusEquipRectHPx()
	return
}
