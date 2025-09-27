package game

// Forecast は簡易命中・ダメージを返す（後で2RN/相性/地形対応）。
func Forecast(att, def Unit) ForecastResult {
	hit := 80 + att.S.Skl*2 + att.S.Lck/2 - (def.S.Spd*2 + def.S.Lck)
	if hit < 0 {
		hit = 0
	}
	if hit > 100 {
		hit = 100
	}
	dmg := att.S.Str + att.W.MT - def.S.Def
	if dmg < 0 {
		dmg = 0
	}
	return ForecastResult{HitDisp: hit, Dmg: dmg, Crit: att.W.Crit}
}
