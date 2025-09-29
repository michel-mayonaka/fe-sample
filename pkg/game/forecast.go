// Package game は戦闘ロジック（予測/解決）を提供します。
package game

// 三すくみ補正（暫定）。Sword>Axe>Lance>Sword
var triangle = map[string]map[string]struct{ Hit, Mt int }{
	"Sword": {"Axe": {Hit: +10, Mt: +2}, "Lance": {Hit: -10, Mt: -2}},
	"Axe":   {"Lance": {Hit: +10, Mt: +2}, "Sword": {Hit: -10, Mt: -2}},
	"Lance": {"Sword": {Hit: +10, Mt: +2}, "Axe": {Hit: -10, Mt: -2}},
}

// Forecast は（互換API）地形無しで命中表示値/与ダメ/必殺%を返します。
// 命中判定は Resolve 系で行います。
func Forecast(att, def Unit) ForecastResult { return ForecastAt(att, def, Terrain{}, Terrain{}) }

// ForecastAt は地形補正を考慮して命中表示値/与ダメ/必殺%を返します。
// 仕様（tasks/battlepart.md準拠）:
//
//	atk_hit = weapon.hit + skl*2 + floor(lck/2) + attacker_tile_hit
//	def_avo = spd*2 + lck + defender_tile_avoid
//	hit_disp = clamp(atk_hit - def_avo + triangle_hit, 0, 100)
//	raw = str + mt + triangle_mt - (def + defender_tile_def)
//	dmg = max(1, raw)
func ForecastAt(att, def Unit, attTile, defTile Terrain) ForecastResult {
	tri := triangleMod(att.W.Type, def.W.Type)
	baseHit := att.W.Hit + att.S.Skl*2 + att.S.Lck/2 + attTile.Hit
	defAvo := def.S.Spd*2 + def.S.Lck + defTile.Avoid
	hit := clamp(baseHit-defAvo+tri.Hit, 0, 100)
	raw := att.S.Str + att.W.MT + tri.Mt - (def.S.Def + defTile.Def)
	if raw < 1 {
		raw = 1
	}
	return ForecastResult{HitDisp: hit, Dmg: raw, Crit: att.W.Crit}
}

// ForecastAtExplain は地形・相性を含む予測と、内訳（UI向け）を返します。
func ForecastAtExplain(att, def Unit, attTile, defTile Terrain) (ForecastResult, ForecastBreakdown) {
	tri := triangleMod(att.W.Type, def.W.Type)
	weapHit := att.W.Hit
	skl2 := att.S.Skl * 2
	lckHalf := att.S.Lck / 2
	attTileHit := attTile.Hit
	defSpd2 := def.S.Spd * 2
	defLck := def.S.Lck
	defTileAvoid := defTile.Avoid
	baseHit := weapHit + skl2 + lckHalf + attTileHit
	defAvo := defSpd2 + defLck + defTileAvoid
	hit := clamp(baseHit-defAvo+tri.Hit, 0, 100)

	atkStr := att.S.Str
	wpnMt := att.W.MT
	triMt := tri.Mt
	defTotal := def.S.Def + defTile.Def
	raw := atkStr + wpnMt + triMt - defTotal
	dmg := raw
	if dmg < 1 {
		dmg = 1
	}
	fr := ForecastResult{HitDisp: hit, Dmg: dmg, Crit: att.W.Crit}
	bk := ForecastBreakdown{
		WeapHit: weapHit, Skl2: skl2, LckHalf: lckHalf, AttTileHit: attTileHit,
		DefSpd2: defSpd2, DefLck: defLck, DefTileAvoid: defTileAvoid,
		TriangleHit: tri.Hit, HitDisp: hit,
		AtkStr: atkStr, WpnMt: wpnMt, TriangleMt: triMt, DefTotal: defTotal,
		Raw: raw, Dmg: dmg, Crit: att.W.Crit,
	}
	return fr, bk
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

// TriangleRelation は三すくみの関係を表します。
// 便宜上、0=中立, +1=有利, -1=不利 とします。
type TriangleRelation int

const (
	// TriangleNeutral は三すくみが中立であることを示します。
	TriangleNeutral TriangleRelation = 0
	// TriangleAdvantage は有利関係を示します。
	TriangleAdvantage TriangleRelation = 1
	// TriangleDisadvantage は不利関係を示します。
	TriangleDisadvantage TriangleRelation = -1
)

// TriangleRelationOf は武器種同士の関係を返します。
func TriangleRelationOf(attType, defType string) TriangleRelation {
	m := triangleMod(attType, defType)
	if m.Mt > 0 || m.Hit > 0 {
		return TriangleAdvantage
	}
	if m.Mt < 0 || m.Hit < 0 {
		return TriangleDisadvantage
	}
	return TriangleNeutral
}

// TerrainPresetName は既知の簡易地形名を返します（未知は "地形"）。
func TerrainPresetName(t Terrain) string {
	switch {
	case t.Avoid == 0 && t.Def == 0 && t.Hit == 0:
		return "平地"
	case t.Avoid == 20 && t.Def == 1 && t.Hit == 0:
		return "森"
	case t.Avoid == 15 && t.Def == 2 && t.Hit == 0:
		return "砦"
	default:
		return "地形"
	}
}
