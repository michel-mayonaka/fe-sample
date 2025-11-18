//go:build !headless

package uicore

// applySim はシミュレーション画面（Terrain/Preview含む）のメトリクスを適用します。
func applySim(t *metricsTargets, m Metrics) {
	if t == nil {
		return
	}
	assignPositive(t.Sim.StartBtnW, m.Sim.StartBtnW)
	assignPositive(t.Sim.StartBtnH, m.Sim.StartBtnH)
	assignPositive(t.Sim.AutoRunGap, m.Sim.AutoRunGap)
	assignPositive(t.Sim.TitleYOffset, m.Sim.TitleYOffset)
	assignPositive(t.Sim.TitleXOffsetFromCenter, m.Sim.TitleXOffsetFromCenter)

	assignPositive(t.Sim.Terrain.ButtonW, m.Sim.Terrain.ButtonW)
	assignPositive(t.Sim.Terrain.ButtonH, m.Sim.Terrain.ButtonH)
	assignPositive(t.Sim.Terrain.BaseYFromBottom, m.Sim.Terrain.BaseYFromBottom)
	assignPositive(t.Sim.Terrain.LeftBaseXOffset, m.Sim.Terrain.LeftBaseXOffset)
	assignPositive(t.Sim.Terrain.RightBaseXInset, m.Sim.Terrain.RightBaseXInset)
	assignPositive(t.Sim.Terrain.ButtonGap, m.Sim.Terrain.ButtonGap)
	assignPositive(t.Sim.Terrain.LabelLeftXOffset, m.Sim.Terrain.LabelLeftXOffset)
	assignPositive(t.Sim.Terrain.LabelYOffsetFromBottom, m.Sim.Terrain.LabelYOffsetFromBottom)

	assignPositive(t.Sim.Preview.LeftXPad, m.Sim.Preview.LeftXPad)
	assignPositive(t.Sim.Preview.RightXInset, m.Sim.Preview.RightXInset)
	assignPositive(t.Sim.Preview.TopYFromMargin, m.Sim.Preview.TopYFromMargin)
	assignPositive(t.Sim.Preview.CardW, m.Sim.Preview.CardW)
	assignPositive(t.Sim.Preview.CardH, m.Sim.Preview.CardH)
	assignPositive(t.Sim.Preview.CardInnerPad, m.Sim.Preview.CardInnerPad)
	assignPositive(t.Sim.Preview.PortraitSize, m.Sim.Preview.PortraitSize)
	assignPositive(t.Sim.Preview.ClassOffsetY, m.Sim.Preview.ClassOffsetY)
	assignPositive(t.Sim.Preview.HPLabelX, m.Sim.Preview.HPLabelX)
	assignPositive(t.Sim.Preview.HPLabelY, m.Sim.Preview.HPLabelY)
	assignPositive(t.Sim.Preview.HPBarX, m.Sim.Preview.HPBarX)
	assignPositive(t.Sim.Preview.HPBarY, m.Sim.Preview.HPBarY)
	assignPositive(t.Sim.Preview.HPBarW, m.Sim.Preview.HPBarW)
	assignPositive(t.Sim.Preview.HPBarH, m.Sim.Preview.HPBarH)
	assignPositive(t.Sim.Preview.NameOffsetX, m.Sim.Preview.NameOffsetX)
	assignPositive(t.Sim.Preview.NameOffsetY, m.Sim.Preview.NameOffsetY)
	assignPositive(t.Sim.Preview.LineY, m.Sim.Preview.LineY)
	assignPositive(t.Sim.Preview.BaseMin, m.Sim.Preview.BaseMin)
	assignPositive(t.Sim.Preview.BaseMax, m.Sim.Preview.BaseMax)
	assignPositive(t.Sim.Preview.WrapPad, m.Sim.Preview.WrapPad)
	assignPositive(t.Sim.Preview.LogPadX, m.Sim.Preview.LogPadX)
	assignPositive(t.Sim.Preview.LogPadY, m.Sim.Preview.LogPadY)
}
