package game

// 最小スキーマ（MVP段階で拡張）
type Stats struct{ HP, Str, Skl, Spd, Lck, Def, Res, Mov int }

type Weapon struct {
	MT, Hit, Crit, Wt int
	RMin, RMax        int
	Type              string // Edge/Pierce/Blunt etc.
}

type Unit struct {
	ID, Name, Class string
	Lv              int
	S               Stats
	W               Weapon
}

type ForecastResult struct {
	HitDisp int
	Dmg     int
	Crit    int
}
