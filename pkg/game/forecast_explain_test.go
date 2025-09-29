package game

import "testing"

func TestForecastAtExplain_BreakdownMatches(t *testing.T) {
    att := Unit{S: Stats{Str: 8, Skl: 7, Spd: 5, Lck: 4, Def: 0}, W: Weapon{MT: 5, Hit: 70, Crit: 5, Type: "Sword"}}
    def := Unit{S: Stats{Str: 0, Skl: 0, Spd: 6, Lck: 3, Def: 4}, W: Weapon{Type: "Axe"}}
    attT := Terrain{Hit: 5}
    defT := Terrain{Avoid: 15, Def: 2}
    fr := ForecastAt(att, def, attT, defT)
    fr2, bk := ForecastAtExplain(att, def, attT, defT)
    if fr != fr2 { t.Fatalf("Forecast mismatch: %+v vs %+v", fr, fr2) }
    // 命中: weapHit + skl*2 + lck/2 + attTileHit - (defSpd*2 + defLck + defTileAvoid) + triHit
    hit := bk.WeapHit + bk.Skl2 + bk.LckHalf + bk.AttTileHit - (bk.DefSpd2 + bk.DefLck + bk.DefTileAvoid) + bk.TriangleHit
    if clamp(hit, 0, 100) != bk.HitDisp { t.Fatalf("hit breakdown mismatch: got=%d want=%d", hit, bk.HitDisp) }
    // ダメージ: atkStr + wpnMt + triMt - (def + defTileDef)
    raw := bk.AtkStr + bk.WpnMt + bk.TriangleMt - bk.DefTotal
    wantDmg := raw
    if wantDmg < 1 { wantDmg = 1 }
    if bk.Dmg != wantDmg { t.Fatalf("dmg breakdown mismatch: got=%d want=%d", bk.Dmg, wantDmg) }
}

func TestTriangleRelationOf(t *testing.T) {
    if got := TriangleRelationOf("Sword", "Axe"); got != TriangleAdvantage { t.Fatalf("Sword>Axe should be advantage") }
    if got := TriangleRelationOf("Sword", "Lance"); got != TriangleDisadvantage { t.Fatalf("Sword<Lance should be disadvantage") }
    if got := TriangleRelationOf("Sword", "Sword"); got != TriangleNeutral { t.Fatalf("same type should be neutral") }
}

func TestTerrainPresetName(t *testing.T) {
    if n := TerrainPresetName(Terrain{}); n != "平地" { t.Fatalf("flat: %s", n) }
    if n := TerrainPresetName(Terrain{Avoid: 20, Def: 1}); n != "森" { t.Fatalf("forest: %s", n) }
    if n := TerrainPresetName(Terrain{Avoid: 15, Def: 2}); n != "砦" { t.Fatalf("fort: %s", n) }
    if n := TerrainPresetName(Terrain{Avoid: 1}); n != "地形" { t.Fatalf("other: %s", n) }
}

func TestDoubleAdvantage_Boundary(t *testing.T) {
    a := Unit{S: Stats{Spd: 10, Bld: 5}, W: Weapon{Wt: 8}}
    d := Unit{S: Stats{Spd: 10, Bld: 5}, W: Weapon{Wt: 10}}
    // AS(a)=10-(8-5)=7, AS(d)=10-(10-5)=5 → 差2 → false
    if DoubleAdvantage(a, d) { t.Fatal("expected false at diff=2") }
    // aの体格を+1 → AS=8, 差3 → true
    a.S.Bld = 6
    if !DoubleAdvantage(a, d) { t.Fatal("expected true at diff=3") }
}

