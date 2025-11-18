//go:build !headless

package uicore

// applyPopup はポップアップ画面メトリクスを適用します。
func applyPopup(t *metricsTargets, m Metrics) {
	if t == nil {
		return
	}
	assignPositive(t.Popup.Cols4ThresholdW, m.Popup.Cols4ThresholdW)
	assignPositive(t.Popup.Cols3ThresholdW, m.Popup.Cols3ThresholdW)
	assignPositive(t.Popup.GridInnerXTotalPad, m.Popup.GridInnerXTotalPad)
	assignPositive(t.Popup.CellH, m.Popup.CellH)
	assignPositive(t.Popup.CellGap, m.Popup.CellGap)
	assignPositive(t.Popup.GridXPad, m.Popup.GridXPad)
	assignPositive(t.Popup.GridYOff, m.Popup.GridYOff)
	assignPositive(t.Popup.MaxW, m.Popup.MaxW)
	assignPositive(t.Popup.MaxH, m.Popup.MaxH)
}
