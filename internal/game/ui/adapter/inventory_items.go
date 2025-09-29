package adapter

import (
	"github.com/hajimehoshi/ebiten/v2"
	gdata "ui_sample/internal/game/data"
	uiview "ui_sample/internal/game/ui/view"
	"ui_sample/internal/model"
	usr "ui_sample/internal/model/user"
)

// BuildItemRows はユーザ所持アイテムと定義テーブル/ユーザテーブルから
// 表示用の行データを構築します。
func BuildItemRows(owns []usr.OwnItem, it *model.ItemDefTable, pl PortraitLoader) []uiview.ItemRow {
	rows := make([]uiview.ItemRow, 0, len(owns))
	for _, oi := range owns {
		name := oi.MstItemsID
		typ, eff := "", ""
		pow := 0
		if it != nil {
			if d, ok := it.FindByID(oi.MstItemsID); ok {
				name, typ, eff, pow = d.Name, d.Type, d.Effect, d.Power
			}
		}
		rows = append(rows, uiview.ItemRow{ID: oi.ID, Name: name, Type: typ, Effect: eff, Power: pow, Uses: oi.Uses, Max: oi.Max})
	}
	var ut *usr.Table
	if p := gdata.Provider(); p != nil {
		ut = p.UserTable()
	}
	if ut == nil {
		return rows
	}
	own := map[string][]uiview.OwnerBadge{}
	for _, c := range ut.Slice() {
		for _, er := range c.Equip {
			if er.UserItemsID == "" {
				continue
			}
			var img *ebiten.Image
			if pl != nil && c.Portrait != "" {
				if im, err := pl.Load(c.Portrait); err == nil {
					img = im
				}
			}
			own[er.UserItemsID] = append(own[er.UserItemsID], uiview.OwnerBadge{Name: c.Name, Portrait: img})
		}
	}
	for i := range rows {
		rows[i].Owners = own[rows[i].ID]
	}
	return rows
}
