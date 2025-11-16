package usecase

import (
	"testing"
	usr "ui_sample/internal/model/user"
)

// スロット拡張とオーナーなしの装備
func TestEquipWeapon_SlotExpand_NoOwner(t *testing.T) {
	ut := newUserTableForTest2(t, []usr.Character{{ID: "u1", Name: "A"}})
	ur := &fakeUserRepo{t: ut}
	a := New(ur, nil, nil, fakeInvRepo{}, nil)
	if err := a.EquipWeapon("u1", 3, "uw_x"); err != nil {
		t.Fatalf("EquipWeapon: %v", err)
	}
	c, _ := ur.Find("u1")
	if len(c.Equip) <= 3 || c.Equip[3].UserWeaponsID != "uw_x" {
		t.Fatalf("slot expand failed: %+v", c.Equip)
	}
	if ur.saved == 0 {
		t.Fatalf("expected Save to be called")
	}
}

func TestEquipItem_SlotExpand_NoOwner(t *testing.T) {
	ut := newUserTableForTest2(t, []usr.Character{{ID: "u1", Name: "A"}})
	ur := &fakeUserRepo{t: ut}
	a := New(ur, nil, nil, fakeInvRepo{}, nil)
	if err := a.EquipItem("u1", 2, "ui_x"); err != nil {
		t.Fatalf("EquipItem: %v", err)
	}
	c, _ := ur.Find("u1")
	if len(c.Equip) <= 2 || c.Equip[2].UserItemsID != "ui_x" {
		t.Fatalf("slot expand failed: %+v", c.Equip)
	}
	if ur.saved == 0 {
		t.Fatalf("expected Save to be called")
	}
}
