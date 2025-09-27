// Package model はマスターデータのスキーマとローダーを提供します。
package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// ItemDef は消耗品等の基本性能を表します（名前重複回避のため Def を付与）。
type ItemDef struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Type   string `json:"type"`   // consumable, key 等
    Effect string `json:"effect"` // heal 等
    Power  int    `json:"power"`  // 効果量（例: heal=10）
}

// ItemDefTable はアイテム名から定義を引ける簡易テーブルです。
type ItemDefTable struct{ byName map[string]ItemDef; byID map[string]ItemDef }

// LoadItemsJSON はアイテム定義の JSON を読み込みます。
func LoadItemsJSON(path string) (*ItemDefTable, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open items: %w", err)
	}
	var rows []ItemDef
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("decode items: %w", err)
	}
    t := &ItemDefTable{byName: make(map[string]ItemDef, len(rows)), byID: make(map[string]ItemDef, len(rows))}
    for _, it := range rows {
        t.byName[it.Name] = it
        if it.ID != "" { t.byID[it.ID] = it }
    }
    return t, nil
}

// Find は名称に一致するアイテム定義を返します。
func (t *ItemDefTable) Find(name string) (ItemDef, bool) {
    if t == nil {
        return ItemDef{}, false
    }
    it, ok := t.byName[name]
    return it, ok
}

// FindByID はIDでアイテム定義を引きます。
func (t *ItemDefTable) FindByID(id string) (ItemDef, bool) {
    if t == nil { return ItemDef{}, false }
    it, ok := t.byID[id]
    return it, ok
}
