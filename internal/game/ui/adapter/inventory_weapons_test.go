package adapter

import (
    "testing"
    gdata "ui_sample/internal/game/data"
    usr "ui_sample/internal/model/user"
    "ui_sample/internal/model"
)

func TestBuildWeaponRows_OwnersAndPortraits(t *testing.T) {
    owns := []usr.OwnWeapon{{ID:"u_weap_1", MstWeaponsID:"mst_sword", Uses:12, Max:40}}
    ut := usr.NewTable([]usr.Character{{
        ID:"char_2", Name:"Bob", Portrait:"bob.png",
        Equip: []usr.EquipRef{{UserWeaponsID: "u_weap_1"}},
    }})
    // Provider を差し替えて Owner バッジ解決を可能に
    gdata.SetProvider(fakeProv{ut: ut})
    pl := &mockPL{withImage:false}
    rows := BuildWeaponRows(owns, nil, pl)
    if len(rows) != 1 { t.Fatalf("rows len=%d", len(rows)) }
    r := rows[0]
    if r.Name != "mst_sword" { t.Errorf("want name passthrough, got %s", r.Name) }
    if r.Uses != 12 || r.Max != 40 { t.Errorf("uses/max not kept: %d/%d", r.Uses, r.Max) }
    if len(r.Owners) != 1 { t.Fatalf("owners len=%d", len(r.Owners)) }
    if r.Owners[0].Name != "Bob" { t.Errorf("owner name=%s", r.Owners[0].Name) }
    if r.Owners[0].Portrait != nil { t.Errorf("expected nil portrait when loader returns nil") }
}

func TestBuildWeaponRows_WithDefinitions(t *testing.T) {
    owns := []usr.OwnWeapon{{ID:"u_w1", MstWeaponsID:"wp_iron_sword", Uses:40, Max:40}}
    wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    if err != nil { t.Skipf("weapons table not available: %v", err) }
    gdata.SetProvider(fakeProv{})
    rows := BuildWeaponRows(owns, wt, nil)
    if len(rows) != 1 { t.Fatalf("rows len=%d", len(rows)) }
    r := rows[0]
    if r.Name == "wp_iron_sword" { t.Errorf("expected name resolved, got id: %s", r.Name) }
    if r.Type == "" || r.Rank == "" || r.Might <= 0 {
        t.Errorf("expected fields from definitions, got type=%s rank=%s might=%d", r.Type, r.Rank, r.Might)
    }
}

// ------- テスト補助 -------

type mockPL struct{ called []string; withImage bool }

func (m *mockPL) Load(name string) (*ebiten.Image, error) {
    m.called = append(m.called, name)
    if m.withImage {
        return ebiten.NewImage(1,1), nil
    }
    return nil, nil
}

import "github.com/hajimehoshi/ebiten/v2"

type fakeProv struct{ ut *usr.Table }
func (f fakeProv) WeaponsTable() *model.WeaponTable                 { return nil }
func (f fakeProv) ItemsTable() *model.ItemDefTable                  { return nil }
func (f fakeProv) UserWeapons() []usr.OwnWeapon                     { return nil }
func (f fakeProv) UserItems() []usr.OwnItem                         { return nil }
func (f fakeProv) UserTable() *usr.Table                            { return f.ut }
func (f fakeProv) UserUnitByID(string) (ui.Unit, bool)              { return ui.Unit{}, false }
func (f fakeProv) EquipKindAt(string, int) (bool, bool)             { return false, false }

import ui "ui_sample/internal/game/service/ui"
