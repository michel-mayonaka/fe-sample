package adapter

import (
    "testing"
    usr "ui_sample/internal/model/user"
)

func TestBuildWeaponRows_OwnersAndPortraits(t *testing.T) {
    owns := []usr.OwnWeapon{{ID:"u_weap_1", MstWeaponsID:"mst_sword", Uses:12, Max:40}}
    ut := usr.NewTable([]usr.Character{{
        ID:"char_2", Name:"Bob", Portrait:"bob.png",
        Equip: []usr.EquipRef{{UserWeaponsID: "u_weap_1"}},
    }})
    pl := &mockPL{withImage:false}
    rows := BuildWeaponRows(owns, nil, ut, pl)
    if len(rows) != 1 { t.Fatalf("rows len=%d", len(rows)) }
    r := rows[0]
    if r.Name != "mst_sword" { t.Errorf("want name passthrough, got %s", r.Name) }
    if r.Uses != 12 || r.Max != 40 { t.Errorf("uses/max not kept: %d/%d", r.Uses, r.Max) }
    if len(r.Owners) != 1 { t.Fatalf("owners len=%d", len(r.Owners)) }
    if r.Owners[0].Name != "Bob" { t.Errorf("owner name=%s", r.Owners[0].Name) }
    if r.Owners[0].Portrait != nil { t.Errorf("expected nil portrait when loader returns nil") }
}

