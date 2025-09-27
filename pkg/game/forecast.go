package game

// 三すくみ補正（暫定）。Sword>Axe>Lance>Sword
var triangle = map[string]map[string]struct{ Hit, Mt int }{
	"Sword": {"Axe": {Hit: +10, Mt: +2}, "Lance": {Hit: -10, Mt: -2}},
	"Axe":   {"Lance": {Hit: +10, Mt: +2}, "Sword": {Hit: -10, Mt: -2}},
	"Lance": {"Sword": {Hit: +10, Mt: +2}, "Axe": {Hit: -10, Mt: -2}},
}

// Forecast は命中表示値/与ダメ/必殺%を返す（2RNの表示値はmeanで、判定はResolve側）。
func Forecast(att, def Unit) ForecastResult {
	tri := triangleMod(att.W.Type, def.W.Type)
	baseHit := att.W.Hit + att.S.Skl*2 + att.S.Lck/2 - (def.S.Spd*2 + def.S.Lck)
	hit := clamp(baseHit+tri.Hit, 0, 100)
	dmg := att.S.Str + att.W.MT + tri.Mt - def.S.Def
	if dmg < 0 {
		dmg = 0
	}
	return ForecastResult{HitDisp: hit, Dmg: dmg, Crit: att.W.Crit}
}

func triangleMod(attType, defType string) (m struct{ Hit, Mt int }) {
	if m2, ok := triangle[attType]; ok {
		if r, ok := m2[defType]; ok {
			return r
		}
	}
	return
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
