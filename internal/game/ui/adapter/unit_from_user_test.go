package adapter

import (
    "testing"
    "github.com/hajimehoshi/ebiten/v2"
    gdata "ui_sample/internal/game/data"
    "ui_sample/internal/model"
    usr "ui_sample/internal/model/user"
)

// fakeProvFull は TableProvider を満たすテスト用実装です。
type fakeProvFull struct {
    wt *model.WeaponTable
    it *model.ItemDefTable
    ut *usr.Table
    uw []usr.OwnWeapon
    ui []usr.OwnItem
}

func (f fakeProvFull) WeaponsTable() *model.WeaponTable { return f.wt }
func (f fakeProvFull) ItemsTable() *model.ItemDefTable  { return f.it }
func (f fakeProvFull) UserWeapons() []usr.OwnWeapon     { return f.uw }
func (f fakeProvFull) UserItems() []usr.OwnItem         { return f.ui }
func (f fakeProvFull) UserTable() *usr.Table            { return f.ut }
func (f fakeProvFull) EquipKindAt(unitID string, slot int) (bool, bool) {
    if f.ut == nil { return false, false }
    c, ok := f.ut.Find(unitID); if !ok { return false, false }
    if slot < 0 || slot >= len(c.Equip) { return false, false }
    er := c.Equip[slot]
    return er.UserWeaponsID != "", er.UserItemsID != ""
}

// mock portrait loader（inventory_* テストの簡易版）
type plMock struct{ called []string; withImage bool }
func (m *plMock) Load(name string) (*ebiten.Image, error) {
    m.called = append(m.called, name)
    if m.withImage { return ebiten.NewImage(1,1), nil }
    return nil, nil
}

func TestUnitFromUser_ResolveNamesAndPortrait(t *testing.T) {
    wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    if err != nil { t.Skipf("weapons table not found: %v", err) }
    it, err := model.LoadItemsJSON("db/master/mst_items.json")
    if err != nil { t.Skipf("items table not found: %v", err) }

    uw := []usr.OwnWeapon{{ID:"uw1", MstWeaponsID:"wp_iron_sword", Uses:40, Max:40}}
    ui := []usr.OwnItem{{ID:"ui1", MstItemsID:"it_vulnerary", Uses:3, Max:3}}
    c := usr.Character{
        ID:"alice", Name:"Alice", Class:"Pegasus", Portrait:"alice.png",
        Level:7, HPMax:26, Stats: usr.Stats{Str:9, Spd:14},
        Equip: []usr.EquipRef{{UserWeaponsID:"uw1"}, {UserItemsID:"ui1"}},
    }
    gdata.SetProvider(fakeProvFull{wt: wt, it: it, ut: usr.NewTable([]usr.Character{c}), uw: uw, ui: ui})
    pl := &plMock{withImage:true}

    u := UnitFromUser(c, pl)
    if u.Name != "Alice" || u.Class != "Pegasus" { t.Fatalf("basic fields mismatch: %+v", u) }
    if u.HP != u.HPMax { t.Errorf("HP init: got %d want %d", u.HP, u.HPMax) }
    if u.Portrait == nil { t.Errorf("expected portrait image to be loaded") }
    if len(pl.called) == 0 || pl.called[0] != "alice.png" { t.Errorf("portrait loader not called: %+v", pl.called) }
    if len(u.Equip) != 2 { t.Fatalf("equip len=%d", len(u.Equip)) }
    // 名前解決（IDのままではない）
    if u.Equip[0].Name == "wp_iron_sword" && u.Equip[1].Name == "it_vulnerary" {
        t.Errorf("expected resolved names, got ids: %+v", u.Equip)
    }
    // 耐久値
    if u.Equip[0].Uses == 0 || u.Equip[0].Max == 0 { t.Errorf("weapon uses/max not set: %+v", u.Equip[0]) }
}

func TestBuildUnitsFromProvider_BuildsAll(t *testing.T) {
    wt, _ := model.LoadWeaponsJSON("db/master/mst_weapons.json")
    it, _ := model.LoadItemsJSON("db/master/mst_items.json")
    uw := []usr.OwnWeapon{{ID:"uw1", MstWeaponsID:"wp_iron_sword", Uses:40, Max:40}}
    ui := []usr.OwnItem{{ID:"ui1", MstItemsID:"it_vulnerary", Uses:3, Max:3}}
    rows := []usr.Character{{ID:"alice", Name:"Alice", Portrait:"alice.png", HPMax:20, Equip: []usr.EquipRef{{UserWeaponsID:"uw1"}}}, {ID:"bob", Name:"Bob", HPMax:18}}
    ut := usr.NewTable(rows)
    gdata.SetProvider(fakeProvFull{wt: wt, it: it, ut: ut, uw: uw, ui: ui})
    pl := &plMock{withImage:false}

    units := BuildUnitsFromProvider(pl)
    if len(units) != 2 { t.Fatalf("units len=%d", len(units)) }
    if units[0].ID != "alice" || units[1].ID != "bob" { t.Errorf("ids order mismatch: %s,%s", units[0].ID, units[1].ID) }
    if units[0].Portrait != nil { t.Errorf("no portrait expected without loader image (got non-nil)") }
}

