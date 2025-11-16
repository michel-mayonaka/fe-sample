package usecase

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"ui_sample/internal/infra/userfs"
	usr "ui_sample/internal/model/user"
)

// --- fakes ---------------------------------------------------------------

type fakeUserRepo struct {
	t     *usr.Table
	saved int
}

func (r *fakeUserRepo) Find(id string) (usr.Character, bool) { return r.t.Find(id) }
func (r *fakeUserRepo) Update(c usr.Character)               { r.t.UpdateCharacter(c) }
func (r *fakeUserRepo) Save() error                          { r.saved++; return nil }
func (r *fakeUserRepo) Table() *usr.Table                    { return r.t }

type fakeInvRepo struct{}

func (f fakeInvRepo) Consume(string, int) error { return nil }
func (f fakeInvRepo) Save() error               { return nil }
func (f fakeInvRepo) Reload() error             { return nil }
func (f fakeInvRepo) Weapons() []usr.OwnWeapon  { return nil }
func (f fakeInvRepo) Items() []usr.OwnItem      { return nil }

// helper: create temporary user table JSON and load it
func newUserTableForTest(t *testing.T, rows []usr.Character) *usr.Table {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "usr_characters.json")
	b, err := json.Marshal(rows)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, b, 0644); err != nil {
		t.Fatal(err)
	}
	ut, err := userfs.LoadTableJSON(p)
	if err != nil {
		t.Fatal(err)
	}
	return ut
}

// TestEquipWeapon_TransfersOwnership は武器の所有者移譲と巻き戻しを検証します。
func TestEquipWeapon_TransfersOwnership(t *testing.T) {
	ut := newUserTableForTest(t, []usr.Character{
		{ID: "u1", Name: "A", Equip: []usr.EquipRef{{UserWeaponsID: "uw_2"}}},
		{ID: "u2", Name: "B", Equip: []usr.EquipRef{{UserWeaponsID: "uw_1"}}},
	})
	ur := &fakeUserRepo{t: ut}
	a := New(ur, nil, nil, fakeInvRepo{}, nil)

	if err := a.EquipWeapon("u1", 0, "uw_1"); err != nil {
		t.Fatalf("equip weapon: %v", err)
	}
	c1, _ := ur.Find("u1")
	c2, _ := ur.Find("u2")
	if got := c1.Equip[0].UserWeaponsID; got != "uw_1" {
		t.Fatalf("u1 slot0 = %s, want uw_1", got)
	}
	if got := c2.Equip[0].UserWeaponsID; got != "uw_2" {
		t.Fatalf("u2 slot0 = %s, want uw_2 (rollback)", got)
	}
	if ur.saved == 0 {
		t.Fatalf("expected Save to be called")
	}
}

// TestEquipItem_TransfersOwnership はアイテムの所有者移譲と巻き戻しを検証します。
func TestEquipItem_TransfersOwnership(t *testing.T) {
	ut := newUserTableForTest(t, []usr.Character{
		{ID: "u1", Name: "A", Equip: []usr.EquipRef{{}, {UserItemsID: "ui_2"}}},
		{ID: "u2", Name: "B", Equip: []usr.EquipRef{{}, {UserItemsID: "ui_1"}}},
	})
	ur := &fakeUserRepo{t: ut}
	a := New(ur, nil, nil, fakeInvRepo{}, nil)

	if err := a.EquipItem("u1", 1, "ui_1"); err != nil {
		t.Fatalf("equip item: %v", err)
	}
	c1, _ := ur.Find("u1")
	c2, _ := ur.Find("u2")
	if got := c1.Equip[1].UserItemsID; got != "ui_1" {
		t.Fatalf("u1 slot1 = %s, want ui_1", got)
	}
	if got := c2.Equip[1].UserItemsID; got != "ui_2" {
		t.Fatalf("u2 slot1 = %s, want ui_2 (rollback)", got)
	}
	if ur.saved == 0 {
		t.Fatalf("expected Save to be called")
	}
}
