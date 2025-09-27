package game

import (
    "math/rand"
    "testing"
)

func TestResolveRound_2RN_Bounds(t *testing.T) {
    rng := rand.New(rand.NewSource(1))
    att := Unit{S: Stats{HP: 20, Str: 5, Skl: 5, Spd: 5, Lck: 0}, W: Weapon{MT: 5, Hit: 999, Type: "Sword"}}
    def := Unit{S: Stats{HP: 20, Def: 0, Lck: 0}}
    // HitDisp は 100 にクランプ → 常に命中
    a1, d1, _ := ResolveRound(att, def, rng)
    if d1.S.HP == def.S.HP { // ダメージが入っていない
        t.Fatal("expected hit but got miss")
    }
    // 逆に、命中0なら絶対に当たらない
    att.W.Hit = -999
    def.S.HP = 20
    _, d2, _ := ResolveRound(att, def, rng)
    if d2.S.HP != 20 {
        t.Fatal("expected miss but got hit")
    }
    _ = a1
}

func TestResolveRound_CritAndMinDamage(t *testing.T) {
    rng := rand.New(rand.NewSource(7))
    att := Unit{S: Stats{HP: 20, Str: 5, Skl: 0, Spd: 0, Lck: 0}, W: Weapon{MT: 1, Hit: 100, Crit: 100, Type: "Sword"}}
    def := Unit{S: Stats{HP: 20, Def: 0, Lck: 0}}
    // 100% クリティカル
    _, d, _ := ResolveRound(att, def, rng)
    fr := Forecast(att, def)
    want := fr.Dmg * 2
    if def.S.HP-d.S.HP != want {
        t.Fatalf("crit damage mismatch: got=%d want=%d", def.S.HP-d.S.HP, want)
    }
}

func TestResolveRoundAt_TerrainApplied(t *testing.T) {
    rng := rand.New(rand.NewSource(42))
    att := Unit{S: Stats{HP: 20, Str: 6, Skl: 0, Spd: 0, Lck: 0}, W: Weapon{MT: 3, Hit: 100, Type: "Sword"}}
    def := Unit{S: Stats{HP: 20, Def: 5, Lck: 0}}
    _, d1, _ := ResolveRoundAt(att, def, Terrain{}, Terrain{Def: 0}, rng)
    // Def+3なら被ダメが3減る（ただし最小1）
    def.S.HP = 20
    _, d2, _ := ResolveRoundAt(att, def, Terrain{}, Terrain{Def: 3}, rng)
    if (20-d2.S.HP) >= (20-d1.S.HP) {
        t.Fatalf("expected terrain Def to reduce damage: dmg(def=3)=%d dmg(def=0)=%d", 20-d2.S.HP, 20-d1.S.HP)
    }
}

