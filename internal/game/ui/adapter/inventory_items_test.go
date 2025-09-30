package adapter

import (
    "testing"
    "github.com/hajimehoshi/ebiten/v2"
    gdata "ui_sample/internal/game/data"
    usr "ui_sample/internal/model/user"
    "ui_sample/internal/model"
)

type mockPL struct{ called []string; withImage bool }

func (m *mockPL) Load(name string) (*ebiten.Image, error) {
    m.called = append(m.called, name)
    if m.withImage {
        return ebiten.NewImage(1,1), nil
    }
    return nil, nil
}

// fakeProv はテスト用の簡易 TableProvider 実装です（必要メソッドのみ）。
type fakeProv struct{ ut *usr.Table }
func (f fakeProv) WeaponsTable() *model.WeaponTable                 { return nil }
func (f fakeProv) ItemsTable() *model.ItemDefTable                  { return nil }
func (f fakeProv) UserWeapons() []usr.OwnWeapon                     { return nil }
func (f fakeProv) UserItems() []usr.OwnItem                         { return nil }
func (f fakeProv) UserTable() *usr.Table                            { return f.ut }
func (f fakeProv) EquipKindAt(string, int) (bool, bool)             { return false, false }

// ui 型はここでは不要

func TestBuildItemRows_OwnersAndPortraits(t *testing.T) {
    owns := []usr.OwnItem{{ID:"u_item_1", MstItemsID:"mst_potion", Uses:3, Max:10}}
    ut := usr.NewTable([]usr.Character{{
        ID:"char_1", Name:"Alice", Portrait:"alice.png",
        Equip: []usr.EquipRef{{UserItemsID: "u_item_1"}},
    }})
    gdata.SetProvider(fakeProv{ut: ut})
    pl := &mockPL{withImage:true}
    rows := BuildItemRows(owns, nil, pl)
    if len(rows) != 1 { t.Fatalf("rows len=%d", len(rows)) }
    r := rows[0]
    if r.Name != "mst_potion" { t.Errorf("want name passthrough, got %s", r.Name) }
    if r.Uses != 3 || r.Max != 10 { t.Errorf("uses/max not kept: %d/%d", r.Uses, r.Max) }
    if len(r.Owners) != 1 { t.Fatalf("owners len=%d", len(r.Owners)) }
    if r.Owners[0].Name != "Alice" { t.Errorf("owner name=%s", r.Owners[0].Name) }
    if r.Owners[0].Portrait == nil { t.Errorf("expected portrait image loaded") }
    if len(pl.called) == 0 || pl.called[0] != "alice.png" { t.Errorf("portrait loader not called correctly: %+v", pl.called) }
}

func TestBuildItemRows_WithDefinitions(t *testing.T) {
    owns := []usr.OwnItem{{ID:"u_item_1", MstItemsID:"it_vulnerary", Uses:2, Max:10}}
    it, err := model.LoadItemsJSON("db/master/mst_items.json")
    if err != nil { t.Skipf("items table not available: %v", err) }
    gdata.SetProvider(fakeProv{})
    rows := BuildItemRows(owns, it, nil)
    if len(rows) != 1 { t.Fatalf("rows len=%d", len(rows)) }
    r := rows[0]
    if r.Name == "it_vulnerary" { t.Errorf("expected name resolved, got id: %s", r.Name) }
    if r.Type == "" || r.Effect == "" || r.Power <= 0 {
        t.Errorf("expected fields from definitions, got type=%s effect=%s power=%d", r.Type, r.Effect, r.Power)
    }
}
