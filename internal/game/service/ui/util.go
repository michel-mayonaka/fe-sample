package uicore

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
	"ui_sample/internal/assets"
	"ui_sample/internal/infra/userfs"
	"ui_sample/internal/model"
	usr "ui_sample/internal/model/user"
)

// Itoa は整数を10進文字列に変換します。
func Itoa(n int) string { return strconv.Itoa(n) }

var (
	cachedMaster     *model.Table
	cachedWeapons    *model.WeaponTable
	cachedItems      *model.ItemDefTable
	cachedUsrWeapons []usr.OwnWeapon
	cachedUsrItems   []usr.OwnItem
	cacheLoaded      bool
)

func ensureCaches() {
	if cacheLoaded {
		return
	}
	if t, err := model.LoadFromJSON("db/master/mst_characters.json"); err == nil {
		cachedMaster = t
	}
	if wt, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil {
		cachedWeapons = wt
	}
	if it, err := model.LoadItemsJSON("db/master/mst_items.json"); err == nil {
		cachedItems = it
	}
	if uw, err := userfs.LoadUserWeaponsJSON("db/user/usr_weapons.json"); err == nil {
		cachedUsrWeapons = uw
	}
	if ui, err := userfs.LoadUserItemsJSON("db/user/usr_items.json"); err == nil {
		cachedUsrItems = ui
	}
	cacheLoaded = true
}

// UnitFromUser はユーザのキャラクターレコードから表示用の Unit を構築します。
func UnitFromUser(c usr.Character) Unit {
	ensureCaches()
	u := Unit{
		ID: c.ID, Name: c.Name, Class: c.Class, Level: c.Level, Exp: c.Exp,
		HPMax: c.HPMax, Stats: Stats(c.Stats),
		Weapon: WeaponRanks{Sword: c.Weapon.Sword, Lance: c.Weapon.Lance, Axe: c.Weapon.Axe, Bow: c.Weapon.Bow},
		Magic:  MagicRanks{Anima: c.Magic.Anima, Light: c.Magic.Light, Dark: c.Magic.Dark, Staff: c.Magic.Staff},
		Growth: Growth(c.Growth),
	}
	// 新参照方式: マスタで user_*_id を引く → usr_* からUses等を取得 → マスタ定義名へ解決
	// 現行仕様: usr_* 参照を優先（旧仕様の直接埋め込みは非対応）
	// 優先: ユーザ側の equip 参照（usr_* のID）
	if len(c.Equip) > 0 {
		for _, er := range c.Equip {
			if er.UserWeaponsID != "" {
				for _, ow := range cachedUsrWeapons {
					if ow.ID == er.UserWeaponsID {
						name := ow.MstWeaponsID
						if cachedWeapons != nil {
							if w, ok := cachedWeapons.FindByID(ow.MstWeaponsID); ok {
								name = w.Name
							}
						}
						u.Equip = append(u.Equip, Item{ID: ow.ID, Name: name, Uses: ow.Uses, Max: ow.Max})
						break
					}
				}
			} else if er.UserItemsID != "" {
				for _, oi := range cachedUsrItems {
					if oi.ID == er.UserItemsID {
						name := oi.MstItemsID
						if cachedItems != nil {
							if it, ok := cachedItems.FindByID(oi.MstItemsID); ok {
								name = it.Name
							}
						}
						u.Equip = append(u.Equip, Item{ID: oi.ID, Name: name, Uses: oi.Uses, Max: oi.Max})
						break
					}
				}
			}
		}
	}
	// 次点: マスタに user_*_id があればそれを参照
	if cachedMaster != nil {
		key := c.MstCharactersID
		if key == "" {
			key = c.ID
		}
		if mc, ok := cachedMaster.Find(key); ok {
			for _, uwid := range mc.UserWeaponsID {
				for _, ow := range cachedUsrWeapons {
					if ow.ID == uwid {
						name := ow.MstWeaponsID
						if cachedWeapons != nil {
							if w, ok := cachedWeapons.FindByID(ow.MstWeaponsID); ok {
								name = w.Name
							}
						}
						u.Equip = append(u.Equip, Item{ID: ow.ID, Name: name, Uses: ow.Uses, Max: ow.Max})
						break
					}
				}
			}
			for _, uiid := range mc.UserItemsID {
				for _, oi := range cachedUsrItems {
					if oi.ID == uiid {
						name := oi.MstItemsID
						if cachedItems != nil {
							if it, ok := cachedItems.FindByID(oi.MstItemsID); ok {
								name = it.Name
							}
						}
						u.Equip = append(u.Equip, Item{ID: oi.ID, Name: name, Uses: oi.Uses, Max: oi.Max})
						break
					}
				}
			}
		}
	}
	// 互換: 旧仕様（ユーザ側 Equip 埋め込み）は未対応のためスキップ
	if c.Portrait != "" {
		if img, err := assets.LoadImage(c.Portrait); err == nil {
			u.Portrait = img
		}
	}
	if u.HP == 0 && u.HPMax > 0 {
		u.HP = u.HPMax
	} else {
		u.HP = c.HP
	}
	return u
}

// LoadUnitsFromUser はユーザテーブルJSONから Unit の配列を構築します。
func LoadUnitsFromUser(path string) ([]Unit, error) {
	ut, err := userfs.LoadTableJSON(path)
	if err != nil {
		return nil, err
	}
	units := make([]Unit, 0, len(ut.ByID()))
	for _, c := range ut.Slice() {
		units = append(units, UnitFromUser(c))
	}
	return units, nil
}

// DrawWrapped は文字列を最大幅 maxW で折り返して描画します。
// 戻り値は最後に描画した行のY座標（次の行を書く起点）です。
func DrawWrapped(dst *ebiten.Image, face font.Face, s string, x, y int, col color.Color, maxW, lineH int) int {
	if maxW <= 0 {
		TextDraw(dst, s, face, x, y, col)
		return y + lineH
	}
	// ルーン単位で貪欲折り返し
	line := ""
	for _, r := range s {
		trial := line + string(r)
		if int(font.MeasureString(face, trial)>>6) > maxW {
			TextDraw(dst, line, face, x, y, col)
			y += lineH
			line = string(r)
			continue
		}
		line = trial
	}
	if line != "" {
		TextDraw(dst, line, face, x, y, col)
		y += lineH
	}
	return y
}
