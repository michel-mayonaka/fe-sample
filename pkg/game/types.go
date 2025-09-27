package game

// Stats は戦闘計算で用いる最小の能力値セットです。
type Stats struct{ HP, Str, Skl, Spd, Lck, Def, Res, Mov, Bld int }

// Weapon は武器の基本性能を表します。
type Weapon struct {
    MT, Hit, Crit, Wt int
    RMin, RMax        int
    Type              string // Sword/Lance/Axe etc.
}

// Unit は戦闘ロジックにおけるユニットの最小構成です。
type Unit struct {
	ID, Name, Class string
	Lv              int
	S               Stats
	W               Weapon
}

// ForecastResult は命中表示値/与ダメ/必殺% の結果です。
type ForecastResult struct {
    HitDisp int
    Dmg     int
    Crit    int
}

// Terrain はMVP用の簡易地形定義です。
//
// - Avoid: 回避上昇（防御側の回避に加算）
// - Def: 防御上昇（被ダメ計算時の防御に加算）
// - Hit: 命中上昇（攻撃側の命中に加算。多くのタイルでは0想定）
// - Heal: ターン開始時の％回復（本MVPのロジックでは未使用）
type Terrain struct{ Avoid, Def, Hit, Heal int }

// Distance はタイル距離（マンハッタン距離想定）を表します。
type Distance int

// ForecastBreakdown は命中/ダメージ計算の内訳を保持します（UI可視化用）。
type ForecastBreakdown struct {
    // 命中
    WeapHit   int
    Skl2      int
    LckHalf   int
    AttTileHit int
    DefSpd2   int
    DefLck    int
    DefTileAvoid int
    TriangleHit  int
    HitDisp      int
    // ダメージ
    AtkStr    int
    WpnMt     int
    TriangleMt int
    DefTotal  int // defender.Def + defTile.Def
    Raw       int
    Dmg       int
    Crit      int
}
