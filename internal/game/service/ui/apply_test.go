package uicore

import "testing"

func TestApplyMetrics_PartialDoesNotOverride(t *testing.T) {
	// snapshot defaults
	def := DefaultMetrics()
	// apply partial: only Margin changes
	var m Metrics
	m.List.Margin = def.List.Margin + 5
	ApplyMetrics(m)
	if ListMargin != def.List.Margin+5 {
		t.Fatalf("margin not applied: got=%d", ListMargin)
	}
	// other values must remain unchanged
	if ListItemH != def.List.ItemH || ListItemGap != def.List.ItemGap {
		t.Fatalf("unexpected override: h=%d gap=%d", ListItemH, ListItemGap)
	}
}
