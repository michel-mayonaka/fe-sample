package game

import "testing"

func TestForecastAt_BasicFormulaAndClamp(t *testing.T) {
    att := Unit{S: Stats{Str: 10, Skl: 5, Spd: 6, Lck: 4, Def: 2}, W: Weapon{MT: 5, Hit: 80, Crit: 0, Type: "Sword"}}
    def := Unit{S: Stats{Str: 0, Skl: 0, Spd: 5, Lck: 3, Def: 4}, W: Weapon{Type: "Axe"}}
    // 三すくみ: Sword>Axe で Hit+10, Mt+2
    fr := ForecastAt(att, def, Terrain{}, Terrain{})
    // atk_hit = 80 + skl*2 + floor(lck/2) = 80 + 10 + 2 = 92
    // def_avo = spd*2 + lck = 10 + 3 = 13
    // hit = 92 - 13 + 10 = 89
    wantHit := 89
    if fr.HitDisp != wantHit {
        t.Fatalf("HitDisp got=%d want=%d", fr.HitDisp, wantHit)
    }
    // raw = str + mt + tri.mt - def = 10 + 5 + 2 - 4 = 13
    // dmg = max(1, raw) = 13
    if fr.Dmg != 13 {
        t.Fatalf("Dmg got=%d want=13", fr.Dmg)
    }
}

func TestForecastAt_TerrainEffectAndFloor(t *testing.T) {
    att := Unit{S: Stats{Str: 6, Skl: 0, Spd: 0, Lck: 0, Def: 0}, W: Weapon{MT: 3, Hit: 50, Type: "Sword"}}
    def := Unit{S: Stats{Spd: 0, Lck: 0, Def: 8}, W: Weapon{Type: "Lance"}}
    // tri: Sword vs Lance は不利( -10, -2 )
    // raw = 6 + 3 - 2 - (8 + defTile.Def)
    // defTile.Def=3 なら raw = -4 → dmg = 1 に丸め
    fr := ForecastAt(att, def, Terrain{}, Terrain{Def: 3})
    if fr.Dmg != 1 {
        t.Fatalf("terrain Def not applied or min damage not enforced: dmg=%d", fr.Dmg)
    }
    // Avoid効果: defTile.Avoid=20 → 命中が下がる
    fr2 := ForecastAt(att, def, Terrain{}, Terrain{Avoid: 20})
    if fr2.HitDisp >= fr.HitDisp {
        t.Fatalf("terrain Avoid not applied: hit2=%d hit1=%d", fr2.HitDisp, fr.HitDisp)
    }
}
