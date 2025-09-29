// Package adapter は UI モデルと戦闘ロジック間の変換を提供します。
package adapter

import (
	uicore "ui_sample/internal/game/service/ui"
	"ui_sample/internal/model"
	gcore "ui_sample/pkg/game"
)

// UIToGame は UI ユニットを戦闘ロジックの Unit に変換します（先頭装備のみ考慮）。
func UIToGame(wt *model.WeaponTable, u uicore.Unit) gcore.Unit {
	var w model.Weapon
	if len(u.Equip) > 0 && wt != nil {
		if ww, ok := wt.Find(u.Equip[0].Name); ok {
			w = ww
		}
	}
	return gcore.Unit{
		ID: u.ID, Name: u.Name, Class: u.Class, Lv: u.Level,
		S: gcore.Stats{HP: u.HP, Str: u.Stats.Str, Skl: u.Stats.Skl, Spd: u.Stats.Spd, Lck: u.Stats.Lck, Def: u.Stats.Def, Res: u.Stats.Res, Mov: u.Stats.Mov, Bld: u.Stats.Bld},
		W: gcore.Weapon{MT: w.Might, Hit: w.Hit, Crit: w.Crit, Wt: w.Weight, RMin: w.RangeMin, RMax: w.RangeMax, Type: w.Type},
	}
}

// AttackSpeedOf は UI ユニットの攻撃速度を返します（先頭装備を使用）。
// 武器テーブル未設定時は速さ(Spd)を返します。
func AttackSpeedOf(wt *model.WeaponTable, u uicore.Unit) int {
	if wt == nil {
		return u.Stats.Spd
	}
	gu := UIToGame(wt, u)
	return gcore.AttackSpeed(gu)
}
