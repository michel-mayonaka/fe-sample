//go:build !headless

package uicore

// applyStatus はステータス画面メトリクスを適用します。
func applyStatus(t *metricsTargets, m Metrics) {
	if t == nil {
		return
	}
	assignPositive(t.Status.PanelPad, m.Status.PanelPad)
	assignPositive(t.Status.PortraitSize, m.Status.PortraitSize)
	assignPositive(t.Status.TextGapX, m.Status.TextGapX)
	assignPositive(t.Status.NameOffsetY, m.Status.NameOffsetY)
	assignPositive(t.Status.ClassGapFromName, m.Status.ClassGapFromName)
	assignPositive(t.Status.LevelGapFromName, m.Status.LevelGapFromName)
	assignPositive(t.Status.HPGapFromName, m.Status.HPGapFromName)
	assignPositive(t.Status.HPBarGapFromName, m.Status.HPBarGapFromName)
	assignPositive(t.Status.HPBarW, m.Status.HPBarW)
	assignPositive(t.Status.HPBarH, m.Status.HPBarH)
	assignPositive(t.Status.StatsTopGap, m.Status.StatsTopGap)
	assignPositive(t.Status.StatsLineH, m.Status.StatsLineH)
	assignPositive(t.Status.StatsColGap, m.Status.StatsColGap)
	assignPositive(t.Status.WeaponRanksXExtra, m.Status.WeaponRanksXExtra)
	assignPositive(t.Status.RankLineH, m.Status.RankLineH)
	assignPositive(t.Status.MagicRanksTopExtra, m.Status.MagicRanksTopExtra)
	assignPositive(t.Status.EquipTitleGapY, m.Status.EquipTitleGapY)
	assignPositive(t.Status.EquipLineH, m.Status.EquipLineH)
	assignPositive(t.Status.EquipRectYOffset, m.Status.EquipRectYOffset)
	assignPositive(t.Status.EquipRectW, m.Status.EquipRectW)
	assignPositive(t.Status.EquipRectH, m.Status.EquipRectH)
	assignPositive(t.Status.EquipLabelGapX, m.Status.EquipLabelGapX)
	assignPositive(t.Status.EquipUsesX, m.Status.EquipUsesX)
}
