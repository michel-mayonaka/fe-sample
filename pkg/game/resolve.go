package game

import "math/rand"

// ResolveRound は（互換API）地形無しで1回の命中判定とダメージ適用を行います。
// 仕様: 2RN、最小ダメ1、HP下限0。
func ResolveRound(att, def Unit, rng *rand.Rand) (Unit, Unit, string) {
    return ResolveRoundAt(att, def, Terrain{}, Terrain{}, rng)
}

// ResolveRoundAt は地形補正を考慮して1回の命中判定とダメージ適用を行います。
// 判定: 2RN（r1,r2∈[0..99]の平均と表示値の比較）
// クリティカル: weapon.crit + floor(skl/2) - lck（0..100にクランプ）
// ダメージ: max(1, raw) を基準に、クリティカル時は×2。
func ResolveRoundAt(att, def Unit, attTile, defTile Terrain, rng *rand.Rand) (Unit, Unit, string) {
    fr := ForecastAt(att, def, attTile, defTile)
    hitTrue := (rng.Intn(100) + rng.Intn(100)) / 2
    if hitTrue < fr.HitDisp {
        // クリティカル
        crit := att.W.Crit + att.S.Skl/2 - def.S.Lck
        if crit < 0 {
            crit = 0
        }
        if crit > 100 {
            crit = 100
        }
        dmg := fr.Dmg
        if rng.Intn(100) < crit {
            dmg *= 2
        }
        if dmg < 1 {
            dmg = 1
        }
        def.S.HP -= dmg
        if def.S.HP < 0 {
            def.S.HP = 0
        }
        return att, def, sprintf("命中! %dダメージ (HP %d)", dmg, def.S.HP)
    }
    return att, def, "ミス!"
}

// 依存を避けるため、簡易の文字列化を内部で提供（fmt不使用の軽量）。
func sprintf(format string, args ...int) string {
	// フォーマットは限定的: "命中! %dダメージ (HP %d)"
	if len(args) == 2 && format == "命中! %dダメージ (HP %d)" {
		return "命中! " + itoa(args[0]) + "ダメージ (HP " + itoa(args[1]) + ")"
	}
	return ""
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var b [20]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}
