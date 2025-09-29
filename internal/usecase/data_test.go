package usecase

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "testing"
    uicore "ui_sample/internal/game/service/ui"
    "ui_sample/internal/model"
    usr "ui_sample/internal/model/user"
    "ui_sample/internal/infra/userfs"
)

// ---- fakes -----------------------------------------------------------------

type fakeWeaponsRepo struct{ reloaded int }

func (f *fakeWeaponsRepo) Find(string) (model.Weapon, bool) { return model.Weapon{}, false }
func (f *fakeWeaponsRepo) Table() *model.WeaponTable         { return nil }
func (f *fakeWeaponsRepo) Reload() error                     { f.reloaded++; return nil }

type fakeInvRepo2 struct{ reloaded int }

func (f *fakeInvRepo2) Consume(string, int) error { return nil }
func (f *fakeInvRepo2) Save() error               { return nil }
func (f *fakeInvRepo2) Reload() error             { f.reloaded++; return nil }
func (f *fakeInvRepo2) Weapons() []usr.OwnWeapon { return nil }
func (f *fakeInvRepo2) Items() []usr.OwnItem     { return nil }

type fakeUserRepo2 struct{ t *usr.Table; saved int }

func (r *fakeUserRepo2) Find(id string) (usr.Character, bool) { return r.t.Find(id) }
func (r *fakeUserRepo2) Update(c usr.Character)               { r.t.UpdateCharacter(c) }
func (r *fakeUserRepo2) Save() error                           { r.saved++; return nil }
func (r *fakeUserRepo2) Table() *usr.Table                    { return r.t }

// helper: create temporary user table JSON and load it
func newUserTableForTest2(t *testing.T, rows []usr.Character) *usr.Table {
    t.Helper()
    dir := t.TempDir()
    p := filepath.Join(dir, "usr_characters.json")
    b, err := json.Marshal(rows)
    if err != nil { t.Fatal(err) }
    if err := os.WriteFile(p, b, 0644); err != nil { t.Fatal(err) }
    ut, err := userfs.LoadTableJSON(p)
    if err != nil { t.Fatal(err) }
    return ut
}

// ---- tests -----------------------------------------------------------------

func TestReloadData_InvokesRepos(t *testing.T) {
    wr := &fakeWeaponsRepo{}
    ir := &fakeInvRepo2{}
    a := &App{Weapons: wr, Inv: ir}
    if err := a.ReloadData(); err != nil { t.Fatalf("ReloadData error: %v", err) }
    if wr.reloaded != 1 { t.Fatalf("weapons Reload not called: %d", wr.reloaded) }
    if ir.reloaded != 1 { t.Fatalf("inventory Reload not called: %d", ir.reloaded) }
}

type errWeaponsRepo struct{}
func (e errWeaponsRepo) Find(string) (model.Weapon, bool) { return model.Weapon{}, false }
func (e errWeaponsRepo) Table() *model.WeaponTable { return nil }
func (e errWeaponsRepo) Reload() error { return fmt.Errorf("reload weapons error") }

type okInvRepo struct{}
func (okInvRepo) Consume(string, int) error { return nil }
func (okInvRepo) Save() error { return nil }
func (okInvRepo) Reload() error { return nil }
func (okInvRepo) Weapons() []usr.OwnWeapon { return nil }
func (okInvRepo) Items() []usr.OwnItem { return nil }

func TestReloadData_PropagatesError(t *testing.T) {
    a := &App{Weapons: errWeaponsRepo{}, Inv: okInvRepo{}}
    if err := a.ReloadData(); err == nil {
        t.Fatalf("expected error from ReloadData")
    }
}

func TestPersistUnit_UpdatesUserAndSaves(t *testing.T) {
    ut := newUserTableForTest2(t, []usr.Character{{ID: "u1", Name: "A", Level: 1, HP: 10, HPMax: 10, Stats: usr.Stats{Str: 1, Mag: 2, Skl: 3, Spd: 4, Lck: 5, Def: 6, Res: 7, Mov: 8, Bld: 9}}})
    ur := &fakeUserRepo2{t: ut}
    a := &App{Users: ur}
    // 変更を含む UI ユニット
    u := uicore.Unit{
        ID: "u1", Name: "A", Level: 2,
        HP: 12, HPMax: 14,
        Stats: uicore.Stats{Str: 11, Mag: 12, Skl: 13, Spd: 14, Lck: 15, Def: 16, Res: 17, Mov: 18, Bld: 19},
    }
    if err := a.PersistUnit(u); err != nil { t.Fatalf("PersistUnit error: %v", err) }
    c, _ := ur.Find("u1")
    if c.Level != 2 || c.HP != 12 || c.HPMax != 14 {
        t.Fatalf("level/HP mismatch: %+v", c)
    }
    if s := c.Stats; s.Str != 11 || s.Mag != 12 || s.Skl != 13 || s.Spd != 14 || s.Lck != 15 || s.Def != 16 || s.Res != 17 || s.Mov != 18 || s.Bld != 19 {
        t.Fatalf("stats not applied: %+v", s)
    }
    if ur.saved == 0 { t.Fatalf("expected Save to be called") }
}
