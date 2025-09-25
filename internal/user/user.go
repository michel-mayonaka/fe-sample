package user

import (
    "encoding/json"
    "fmt"
    "os"
)

// Stats は現在の能力値を表します。
type Stats struct {
    Str, Mag, Skl, Spd, Lck, Def, Res, Mov int
}

// Growth は成長率（%）を表します。
type Growth struct {
    Str, Mag, Skl, Spd, Lck, Def, Res, Mov int
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

// Character はユーザの現在状態（セーブデータ）です。
// マスタの ID に対応するスナップショットを保持します。
type Character struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Class string `json:"class"`
    Portrait string `json:"portrait"`
    Level int    `json:"level"`
    Exp   int    `json:"exp"`
    HP    int    `json:"hp"`
    HPMax int    `json:"hp_max"`
    Stats Stats  `json:"stats"`
    Growth Growth `json:"growth"`
    Weapon WeaponRanks `json:"weapon"`
    Magic  MagicRanks  `json:"magic"`
    Equip []Item `json:"equip"`
}

// Table はユーザのキャラクター状態を ID で引ける索引です。
type Table struct {
    byID map[string]Character
}

// LoadFromJSON はユーザテーブルJSONを読み込みます。
func LoadFromJSON(path string) (*Table, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("open user table: %w", err)
    }
    defer f.Close()
    var rows []Character
    if err := json.NewDecoder(f).Decode(&rows); err != nil {
        return nil, fmt.Errorf("decode user table: %w", err)
    }
    t := &Table{byID: make(map[string]Character, len(rows))}
    for _, c := range rows {
        t.byID[c.ID] = c
    }
    return t, nil
}

// Find はIDでユーザ状態を返します。
func (t *Table) Find(id string) (Character, bool) {
    if t == nil { return Character{}, false }
    c, ok := t.byID[id]
    return c, ok
}
