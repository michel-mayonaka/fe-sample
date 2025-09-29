package usecase

import (
    "testing"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/model"
    usr "ui_sample/internal/model/user"
    gcore "ui_sample/pkg/game"
)

// Test: 反撃射程外なら防御側は耐久を消費しない
func TestRunBattleRound_NoCounter_OutOfRange(t *testing.T) {
    rows := []model.Weapon{
        {ID: "w_sword", Name: "Iron Sword", Type: "Sword", Rank: "E", Might: 5, Hit: 95, Weight: 5, RangeMin: 1, RangeMax: 1},
        {ID: "w_bow",   Name: "Iron Bow",   Type: "Bow",   Rank: "E", Might: 4, Hit: 90, Weight: 5, RangeMin: 2, RangeMax: 2},
    }
    wt := mustLoadWeapons(rows)
    ut := newUserTableForTest2(t, []usr.Character{{ID: "u1", Name: "A", HP: 20, HPMax: 20}, {ID: "u2", Name: "B", HP: 20, HPMax: 20}})
    ar := &fakeUsers{t: ut}
    ir := &fakeInv{}
    a := &App{Weapons: &fakeWeapons{t: wt}, Users: ar, Inv: ir}

    units := []uicore.Unit{
        {ID: "u1", Name: "A", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 8, Skl: 6, Spd: 7, Def: 2, Bld: 6}, Equip: []uicore.Item{{ID: "uw_1", Name: "Iron Sword", Uses: 5, Max: 40}}},
        {ID: "u2", Name: "B", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 7, Skl: 5, Spd: 6, Def: 2, Bld: 6}, Equip: []uicore.Item{{ID: "uw_2", Name: "Iron Bow", Uses: 5, Max: 30}}},
    }
    got, _, ok, err := a.RunBattleRound(units, 0, gcore.Terrain{}, gcore.Terrain{})
    if err != nil || !ok { t.Fatalf("RunBattleRound: %v", err) }
    if got[0].Equip[0].Uses != 4 { t.Fatalf("attacker should consume 1, got %d", got[0].Equip[0].Uses) }
    if got[1].Equip[0].Uses != 5 { t.Fatalf("defender should not consume, got %d", got[1].Equip[0].Uses) }
}

// Test: レポジトリ未注入でもパニックせずUI更新は行う
func TestRunBattleRound_NoRepos_NoPanic(t *testing.T) {
    wt := mustLoadWeapons([]model.Weapon{{ID: "w_sword", Name: "Iron Sword", Type: "Sword", Rank: "E", Might: 5, Hit: 95, Weight: 5, RangeMin: 1, RangeMax: 1}})
    a := &App{Weapons: &fakeWeapons{t: wt}} // Users/Inv は nil
    units := []uicore.Unit{
        {ID: "u1", Name: "A", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 8, Skl: 6, Spd: 7, Def: 2, Bld: 6}, Equip: []uicore.Item{{ID: "uw_1", Name: "Iron Sword", Uses: 5, Max: 40}}},
        {ID: "u2", Name: "B", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 7, Skl: 5, Spd: 6, Def: 2, Bld: 6}, Equip: []uicore.Item{{Name: "Iron Sword", Uses: 5, Max: 40}}},
    }
    got, _, ok, err := a.RunBattleRound(units, 0, gcore.Terrain{}, gcore.Terrain{})
    if err != nil || !ok { t.Fatalf("RunBattleRound: %v", err) }
    if got[0].HP == 20 && got[1].HP == 20 { t.Fatalf("expected some HP change") }
}

// Test: 不正な選択インデックスは無視
func TestRunBattleRound_InvalidSelIndex(t *testing.T) {
    wt := mustLoadWeapons([]model.Weapon{{ID: "w_sword", Name: "Iron Sword", Type: "Sword", Rank: "E", Might: 5, Hit: 95, Weight: 5, RangeMin: 1, RangeMax: 1}})
    a := &App{Weapons: &fakeWeapons{t: wt}}
    units := []uicore.Unit{{ID: "u1", Name: "A"}}
    got, logs, ok, err := a.RunBattleRound(units, 5, gcore.Terrain{}, gcore.Terrain{})
    if err != nil { t.Fatalf("err: %v", err) }
    if ok || len(logs) != 0 || len(got) != 1 { t.Fatalf("should be no-op for invalid index") }
}
