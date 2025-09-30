package usecase

import (
    "encoding/json"
    "math/rand"
    "os"
    "testing"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/model"
    usr "ui_sample/internal/model/user"
    gcore "ui_sample/pkg/game"
)

// ---- fakes -----------------------------------------------------------------

type fakeWeapons struct{ t *model.WeaponTable }

func (f *fakeWeapons) Find(string) (model.Weapon, bool) { return model.Weapon{}, false }
func (f *fakeWeapons) Table() *model.WeaponTable        { return f.t }
func (f *fakeWeapons) Reload() error                    { return nil }

type fakeUsers struct {
	t     *usr.Table
	saved int
}

func (f *fakeUsers) Find(id string) (usr.Character, bool) { return f.t.Find(id) }
func (f *fakeUsers) Update(c usr.Character)               { f.t.UpdateCharacter(c) }
func (f *fakeUsers) Save() error                          { f.saved++; return nil }
func (f *fakeUsers) Table() *usr.Table                    { return f.t }

type fakeInv struct {
	consumed map[string]int
	saved    int
}

func (f *fakeInv) Consume(id string, n int) error {
	if f.consumed == nil {
		f.consumed = map[string]int{}
	}
	f.consumed[id] += n
	return nil
}
func (f *fakeInv) Save() error              { f.saved++; return nil }
func (f *fakeInv) Reload() error            { return nil }
func (f *fakeInv) Weapons() []usr.OwnWeapon { return nil }
func (f *fakeInv) Items() []usr.OwnItem     { return nil }

func weaponsTableForTest() *model.WeaponTable {
	rows := []model.Weapon{
		{ID: "w_sword", Name: "Iron Sword", Type: "Sword", Rank: "E", Might: 5, Hit: 95, Crit: 0, Weight: 5, RangeMin: 1, RangeMax: 1},
		{ID: "w_axe", Name: "Iron Axe", Type: "Axe", Rank: "E", Might: 6, Hit: 80, Crit: 0, Weight: 8, RangeMin: 1, RangeMax: 1},
	}
	// JSON経由でロードして本番経路と同等に構築
	return mustLoadWeapons(rows)
}

func mustLoadWeapons(rows []model.Weapon) *model.WeaponTable {
	// serialize to temp and load via model API
	// This keeps tests aligned with production loader behavior
	b, _ := jsonMarshal(rows)
	p := writeTempFile(b)
	wt, err := model.LoadWeaponsJSON(p)
	if err != nil {
		panic(err)
	}
	return wt
}

func jsonMarshal(v any) ([]byte, error) { return json.Marshal(v) }

func writeTempFile(b []byte) string {
	f, _ := os.CreateTemp("", "wt-*.json")
	defer func() { _ = f.Close() }()
	_, _ = f.Write(b)
	return f.Name()
}

// ---- tests -----------------------------------------------------------------

func TestRunBattleRound_UpdatesHP_Uses_Saves(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	wt := weaponsTableForTest()
	// user table: two chars with equipment refs
	ut := newUserTableForTest2(t, []usr.Character{{ID: "u1", Name: "A", HP: 20, HPMax: 20}, {ID: "u2", Name: "B", HP: 20, HPMax: 20}})
	ar := &fakeUsers{t: ut}
	ir := &fakeInv{}
	a := &App{Weapons: &fakeWeapons{t: wt}, Users: ar, Inv: ir, RNG: rng}

	units := []uicore.Unit{
		{ID: "u1", Name: "A", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 8, Skl: 6, Spd: 7, Lck: 0, Def: 2, Bld: 6}, Equip: []uicore.Item{{ID: "uw_1", Name: "Iron Sword", Uses: 10, Max: 40}}},
		{ID: "u2", Name: "B", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 7, Skl: 5, Spd: 6, Lck: 0, Def: 2, Bld: 8}, Equip: []uicore.Item{{ID: "uw_2", Name: "Iron Axe", Uses: 10, Max: 45}}},
	}
	gotUnits, logs, ok, err := a.RunBattleRound(units, 0, gflat(), gflat())
	if err != nil || !ok {
		t.Fatalf("RunBattleRound error=%v ok=%v", err, ok)
	}
	if len(logs) == 0 {
		t.Fatalf("expected logs")
	}
	// 双方1回ずつ消費想定（速度差は1）
	if gotUnits[0].Equip[0].Uses != 9 || gotUnits[1].Equip[0].Uses != 9 {
		t.Fatalf("uses not decremented: %+v vs %+v", gotUnits[0].Equip[0], gotUnits[1].Equip[0])
	}
	if ar.saved == 0 {
		t.Fatalf("user Save not called")
	}
	if ir.saved == 0 || ir.consumed["uw_1"] == 0 || ir.consumed["uw_2"] == 0 {
		t.Fatalf("inventory not saved/consumed: saved=%d cons=%v", ir.saved, ir.consumed)
	}
}

func TestRunBattleRound_AttackerDouble_ConsumesTwo(t *testing.T) {
    // 乱数事象（ミス/クリティカル）により追撃や死亡が揺れるのを避けるため、
    // シードを固定して再現性を確保する。
    rng := rand.New(rand.NewSource(1))
	wt := weaponsTableForTest()
	ut := newUserTableForTest2(t, []usr.Character{{ID: "u1"}, {ID: "u2"}})
	a := &App{Weapons: &fakeWeapons{t: wt}, Users: &fakeUsers{t: ut}, Inv: &fakeInv{}, RNG: rng}
	// 追撃条件: AS差>=3 を満たすように a(Spd=12,Bld=8, Wt5)=AS11, d(Spd=8,Bld=5, Wt8)=AS5
	units := []uicore.Unit{
		{ID: "u1", Name: "A", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 8, Skl: 6, Spd: 12, Lck: 0, Def: 2, Bld: 8}, Equip: []uicore.Item{{ID: "uw_1", Name: "Iron Sword", Uses: 5, Max: 40}}},
		{ID: "u2", Name: "B", HP: 20, HPMax: 20, Stats: uicore.Stats{Str: 7, Skl: 5, Spd: 8, Lck: 0, Def: 2, Bld: 5}, Equip: []uicore.Item{{ID: "uw_2", Name: "Iron Axe", Uses: 5, Max: 45}}},
	}
	got, _, ok, err := a.RunBattleRound(units, 0, gflat(), gflat())
	if err != nil || !ok {
		t.Fatalf("RunBattleRound: %v", err)
	}
	// 攻撃側は2回、守備側は1回（反撃）
	if got[0].Equip[0].Uses != 3 {
		t.Fatalf("attacker uses want 3 got %d", got[0].Equip[0].Uses)
	}
	if got[1].Equip[0].Uses != 4 {
		t.Fatalf("defender uses want 4 got %d", got[1].Equip[0].Uses)
	}
}

func gflat() gcore.Terrain { return gcore.Terrain{} }
