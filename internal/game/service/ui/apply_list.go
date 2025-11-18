//go:build !headless

package uicore

// applyList は一覧およびライン高さメトリクスを適用します。
func applyList(t *metricsTargets, m Metrics) {
	if t == nil {
		return
	}
	assignPositive(t.List.Margin, m.List.Margin)
	assignPositive(t.List.ItemH, m.List.ItemH)
	assignPositive(t.List.ItemGap, m.List.ItemGap)
	assignPositive(t.List.PortraitSize, m.List.PortraitSize)
	assignPositive(t.List.TitleOffset, m.List.TitleOffset)
	assignPositive(t.List.HeaderTopGap, m.List.HeaderTopGap)
	assignPositive(t.List.ItemsTopGap, m.List.ItemsTopGap)
	assignPositive(t.List.PanelInnerPaddingX, m.List.PanelInnerPaddingX)
	assignPositive(t.List.TitleXOffset, m.List.TitleXOffset)
	assignPositive(t.List.HeaderBaseX, m.List.HeaderBaseX)
	assignPositive(t.List.RowTextOffsetX, m.List.RowTextOffsetX)
	assignPositive(t.List.RowTextOffsetY, m.List.RowTextOffsetY)
	assignPositive(t.List.RowBorderPad, m.List.RowBorderPad)
	assignPositive(t.List.RowRightIconSize, m.List.RowRightIconSize)
	assignPositive(t.List.RowRightIconGap, m.List.RowRightIconGap)
	assignSlice(t.List.HeaderColumnsItems, m.List.HeaderColumnsItems)
	assignSlice(t.List.HeaderColumnsWeapons, m.List.HeaderColumnsWeapons)
	assignSlice(t.List.RowColumnsItems, m.List.RowColumnsItems)
	assignSlice(t.List.RowColumnsWeapons, m.List.RowColumnsWeapons)
	assignPositive(t.Line.Main, m.Line.Main)
	assignPositive(t.Line.Small, m.Line.Small)
}
