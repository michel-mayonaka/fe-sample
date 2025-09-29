package game

// AttackSpeed は攻撃速度(AS)を計算します。
// 仕様: AS = 速さ - max(0, 武器重さ - 体格)
func AttackSpeed(u Unit) int {
	penalty := u.W.Wt - u.S.Bld
	if penalty < 0 {
		penalty = 0
	}
	return u.S.Spd - penalty
}

// DoubleAdvantage は a が d に対して追撃可能か（AS差>=3）を返します。
func DoubleAdvantage(a, d Unit) bool {
	return AttackSpeed(a) >= AttackSpeed(d)+3
}
