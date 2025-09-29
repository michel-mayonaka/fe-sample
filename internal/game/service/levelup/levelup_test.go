package levelup

import (
	"testing"
	uicore "ui_sample/internal/game/service/ui"
)

func TestRoll_All100_IncrementsAll(t *testing.T) {
	u := uicore.Unit{Growth: uicore.Growth{HP: 100, Str: 100, Mag: 100, Skl: 100, Spd: 100, Lck: 100, Def: 100, Res: 100, Mov: 100}}
	g := Roll(u, func() float64 { return 0.0 }) // 常に当たる
	if g.HPGain != 1 {
		t.Fatalf("HPGain=%d", g.HPGain)
	}
	inc := g.Inc
	if inc.Str*inc.Mag*inc.Skl*inc.Spd*inc.Lck*inc.Def*inc.Res*inc.Mov == 0 {
		t.Fatalf("some stats not incremented: %+v", inc)
	}
}

func TestApply_ClampByClassCaps(t *testing.T) {
	// 剣士（caps: hp_max 55, Str 24, ...）を前提
	u := uicore.Unit{Class: "剣士", Level: 10, HP: 54, HPMax: 54, Stats: uicore.Stats{Str: 24}}
	gains := Gains{HPGain: 5}
	gains.Inc.Str = 2
	Apply(&u, gains, 99)
	if u.Level != 11 {
		t.Errorf("level inc got %d", u.Level)
	}
	if u.HPMax > 55 {
		t.Errorf("HPMax clamp failed: %d", u.HPMax)
	}
	if u.HP > u.HPMax {
		t.Errorf("HP should be clamped: %d>%d", u.HP, u.HPMax)
	}
	if u.Stats.Str > 24 {
		t.Errorf("Str clamp failed: %d", u.Stats.Str)
	}
}
