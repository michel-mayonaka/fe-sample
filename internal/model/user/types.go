// Package user はユーザ（セーブ）データの純粋なモデル型を提供します。
package user

// Stats は現在の能力値を表します。
type Stats struct {
	Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int
}

// Growth は成長率（%）を表します。
type Growth struct {
	HP, Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int
}

// WeaponRanks は物理武器のランクを表します。
type WeaponRanks struct {
	Sword string `json:"sword"`
	Lance string `json:"lance"`
	Axe   string `json:"axe"`
	Bow   string `json:"bow"`
}

// MagicRanks は魔法系のランクを表します。
type MagicRanks struct {
	Anima string `json:"anima"`
	Light string `json:"light"`
	Dark  string `json:"dark"`
	Staff string `json:"staff"`
}

// Item はユーザセーブにおける装備の残耐久を表します。
type Item struct {
	Name string `json:"name"`
	Uses int    `json:"uses"`
	Max  int    `json:"max"`
}

// EquipRef はユーザ装備の参照ID（usr_*）を保持します。
type EquipRef struct {
	UserWeaponsID string `json:"usr_weapons_id,omitempty"`
	UserItemsID   string `json:"usr_items_id,omitempty"`
}

// Character はユーザの現在状態（セーブデータ）です。
// マスタの ID に対応するスナップショットを保持します。
type Character struct {
	ID string `json:"id"`
	// サロゲートキー（ユーザレコード自身のPK）。任意。
	UID string `json:"uid,omitempty"`
	// マスタ参照: mst_characters の ID
	MstCharactersID string      `json:"mst_characters_id,omitempty"`
	Name            string      `json:"name"`
	Class           string      `json:"class"`
	Portrait        string      `json:"portrait"`
	Level           int         `json:"level"`
	Exp             int         `json:"exp,omitempty"`
	HP              int         `json:"hp,omitempty"`
	HPMax           int         `json:"hp_max"`
	Stats           Stats       `json:"stats"`
	Growth          Growth      `json:"growth"`
	Weapon          WeaponRanks `json:"weapon"`
	Magic           MagicRanks  `json:"magic"`
	Equip           []EquipRef  `json:"equip,omitempty"`
}
