// Package model はマスターデータのスキーマとローダーを提供します。
package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// Stats は各種能力値を表します。
type Stats struct {
	Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int
}

// Growth は各能力の成長率（%）を表します。
type Growth struct {
	Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int
}

// Item は耐久制装備（武器・消耗品）を表します。
type Item struct {
	Name string `json:"name"`
	Uses int    `json:"uses"`
	Max  int    `json:"max"`
}

// WeaponRanks は物理系武器のランクを表します。
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

// Character はマスターテーブルの1レコードです。
type Character struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Class    string      `json:"class"`
	Portrait string      `json:"portrait"` // 画像パス（任意）
	Level    int         `json:"level"`
	Exp      int         `json:"exp"`
	HP       int         `json:"hp"`
	HPMax    int         `json:"hp_max"`
	Stats    Stats       `json:"stats"`
	Growth   Growth      `json:"growth"`
	Weapon   WeaponRanks `json:"weapon"`
	Magic    MagicRanks  `json:"magic"`
	Equip    []Item      `json:"equip,omitempty"`
	// 参照方式: ユーザ所持テーブルのIDリストを参照（装備/所持）
	UserWeaponsID []string `json:"user_weapons_id,omitempty"`
	UserItemsID   []string `json:"user_items_id,omitempty"`
}

// Table は読み込んだキャラクターマスタの索引です。
type Table struct {
	byID map[string]Character
}

// LoadFromJSON はJSONファイルからマスタを読み込みます。
func LoadFromJSON(path string) (*Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open master: %w", err)
	}
	defer func() { _ = f.Close() }()
	var rows []Character
	if err := json.NewDecoder(f).Decode(&rows); err != nil {
		return nil, fmt.Errorf("decode master: %w", err)
	}
	t := &Table{byID: make(map[string]Character, len(rows))}
	for _, c := range rows {
		t.byID[c.ID] = c
	}
	return t, nil
}

// Find はIDに一致するキャラクターを返します。
func (t *Table) Find(id string) (Character, bool) {
	if t == nil {
		return Character{}, false
	}
	c, ok := t.byID[id]
	return c, ok
}
