//go:build !headless

package uicore

// applyWidgets は共通ウィジェットボタン群のメトリクスを適用します。
func applyWidgets(t *metricsTargets, m Metrics) {
	if t == nil {
		return
	}
	assignPositive(t.Widgets.Back.PanelRightInset, m.Widgets.Back.PanelRightInset)
	assignPositive(t.Widgets.Back.TopPad, m.Widgets.Back.TopPad)
	assignPositive(t.Widgets.Back.W, m.Widgets.Back.W)
	assignPositive(t.Widgets.Back.H, m.Widgets.Back.H)
	assignPositive(t.Widgets.Back.LabelX, m.Widgets.Back.LabelX)
	assignPositive(t.Widgets.Back.LabelY, m.Widgets.Back.LabelY)

	assignPositive(t.Widgets.LevelUp.W, m.Widgets.LevelUp.W)
	assignPositive(t.Widgets.LevelUp.H, m.Widgets.LevelUp.H)
	assignPositive(t.Widgets.LevelUp.LabelX, m.Widgets.LevelUp.LabelX)
	assignPositive(t.Widgets.LevelUp.LabelY, m.Widgets.LevelUp.LabelY)

	assignPositive(t.Widgets.ToBattle.GapFromRightBtn, m.Widgets.ToBattle.GapFromRightBtn)
	assignPositive(t.Widgets.ToBattle.LabelX, m.Widgets.ToBattle.LabelX)
	assignPositive(t.Widgets.ToBattle.LabelY, m.Widgets.ToBattle.LabelY)
	assignPositive(t.Widgets.ToBattle.W, m.Widgets.ToBattle.W)

	assignPositive(t.Widgets.SimBattle.W, m.Widgets.SimBattle.W)
	assignPositive(t.Widgets.SimBattle.H, m.Widgets.SimBattle.H)
	assignPositive(t.Widgets.SimBattle.TopPad, m.Widgets.SimBattle.TopPad)
	assignPositive(t.Widgets.SimBattle.LabelX, m.Widgets.SimBattle.LabelX)
	assignPositive(t.Widgets.SimBattle.LabelY, m.Widgets.SimBattle.LabelY)
}
