package adapter

import (
    "testing"
    usr "ui_sample/internal/model/user"
)

type mockPL struct{ called []string; withImage bool }

func (m *mockPL) Load(name string) (*ebiten.Image, error) {
    m.called = append(m.called, name)
    if m.withImage {
        return ebiten.NewImage(1,1), nil
    }
    return nil, nil
}

// ebiten import needs to be after func to avoid unused during formatting
import "github.com/hajimehoshi/ebiten/v2"

func TestBuildItemRows_OwnersAndPortraits(t *testing.T) {
    owns := []usr.OwnItem{{ID:"u_item_1", MstItemsID:"mst_potion", Uses:3, Max:10}}
    ut := usr.NewTable([]usr.Character{{
        ID:"char_1", Name:"Alice", Portrait:"alice.png",
        Equip: []usr.EquipRef{{UserItemsID: "u_item_1"}},
    }})
    pl := &mockPL{withImage:true}
    rows := BuildItemRows(owns, nil, ut, pl)
    if len(rows) != 1 { t.Fatalf("rows len=%d", len(rows)) }
    r := rows[0]
    if r.Name != "mst_potion" { t.Errorf("want name passthrough, got %s", r.Name) }
    if r.Uses != 3 || r.Max != 10 { t.Errorf("uses/max not kept: %d/%d", r.Uses, r.Max) }
    if len(r.Owners) != 1 { t.Fatalf("owners len=%d", len(r.Owners)) }
    if r.Owners[0].Name != "Alice" { t.Errorf("owner name=%s", r.Owners[0].Name) }
    if r.Owners[0].Portrait == nil { t.Errorf("expected portrait image loaded") }
    if len(pl.called) == 0 || pl.called[0] != "alice.png" { t.Errorf("portrait loader not called correctly: %+v", pl.called) }
}

