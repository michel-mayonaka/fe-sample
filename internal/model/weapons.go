// Package model はマスターデータのスキーマとローダーを提供します。
package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// Weapon は武器の基本性能を表します。
type Weapon struct {
	Name     string `json:"name"`
	Type     string `json:"type"` // Sword/Lance/Axe/Bow/...
	Rank     string `json:"rank"` // E-D-C-B-A-S
	Might    int    `json:"might"`
	Hit      int    `json:"hit"`
	Crit     int    `json:"crit"`
	Weight   int    `json:"weight"`
	RangeMin int    `json:"rmin"`
	RangeMax int    `json:"rmax"`
}

// WeaponTable は武器名から基本性能を引ける簡易テーブルです。
type WeaponTable struct{ byName map[string]Weapon }

// LoadWeaponsJSON は武器定義の JSON を読み込みます。
func LoadWeaponsJSON(path string) (*WeaponTable, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open weapons: %w", err)
	}
	var rows []Weapon
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("decode weapons: %w", err)
	}
	t := &WeaponTable{byName: make(map[string]Weapon, len(rows))}
	for _, w := range rows {
		t.byName[w.Name] = w
	}
	return t, nil
}

// Find は武器名に一致する定義を返します。
func (t *WeaponTable) Find(name string) (Weapon, bool) {
	if t == nil {
		return Weapon{}, false
	}
	w, ok := t.byName[name]
	return w, ok
}
