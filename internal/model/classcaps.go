package model

import (
    "encoding/json"
    "fmt"
    "os"
)

// ClassCaps はクラスごとの能力上限を表します。
type ClassCaps struct {
    Class string `json:"class"`
    HPMax int    `json:"hp_max"`
    Str   int    `json:"Str"`
    Mag   int    `json:"Mag"`
    Skl   int    `json:"Skl"`
    Spd   int    `json:"Spd"`
    Lck   int    `json:"Lck"`
    Def   int    `json:"Def"`
    Res   int    `json:"Res"`
    Mov   int    `json:"Mov"`
}

type ClassCapsTable struct { byClass map[string]ClassCaps }

func LoadClassCapsJSON(path string) (*ClassCapsTable, error) {
    b, err := os.ReadFile(path)
    if err != nil { return nil, fmt.Errorf("open class caps: %w", err) }
    var rows []ClassCaps
    if err := json.Unmarshal(b, &rows); err != nil { return nil, fmt.Errorf("decode class caps: %w", err) }
    t := &ClassCapsTable{byClass: make(map[string]ClassCaps, len(rows))}
    for _, r := range rows { t.byClass[r.Class] = r }
    return t, nil
}

func (t *ClassCapsTable) Find(class string) (ClassCaps, bool) {
    if t == nil { return ClassCaps{}, false }
    c, ok := t.byClass[class]
    return c, ok
}

