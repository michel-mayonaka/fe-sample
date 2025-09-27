package user

import (
    "encoding/json"
    "fmt"
    "os"
)

// OwnWeapon はユーザが所持する武器インスタンス（耐久）です。
type OwnWeapon struct {
    ID           string `json:"id"`
    MstWeaponsID string `json:"mst_weapons_id"`
    Uses         int    `json:"uses"`
    Max          int    `json:"max"`
}

// OwnItem はユーザが所持する消耗品インスタンス（耐久/使用回数）です。
type OwnItem struct {
    ID          string `json:"id"`
    MstItemsID  string `json:"mst_items_id"`
    Uses        int    `json:"uses"`
    Max         int    `json:"max"`
}

// LoadUserWeaponsJSON は usr_weapons.json を読み込みます。
func LoadUserWeaponsJSON(path string) ([]OwnWeapon, error) {
    b, err := os.ReadFile(path)
    if err != nil { return nil, fmt.Errorf("open usr_weapons: %w", err) }
    var rows []OwnWeapon
    if err := json.Unmarshal(b, &rows); err != nil { return nil, fmt.Errorf("decode usr_weapons: %w", err) }
    return rows, nil
}

// SaveUserWeaponsJSON は usr_weapons.json として保存します。
func SaveUserWeaponsJSON(path string, rows []OwnWeapon) error {
    buf, err := json.MarshalIndent(rows, "", "  ")
    if err != nil { return err }
    return os.WriteFile(path, buf, 0644)
}

// LoadUserItemsJSON は usr_items.json を読み込みます。
func LoadUserItemsJSON(path string) ([]OwnItem, error) {
    b, err := os.ReadFile(path)
    if err != nil { return nil, fmt.Errorf("open usr_items: %w", err) }
    var rows []OwnItem
    if err := json.Unmarshal(b, &rows); err != nil { return nil, fmt.Errorf("decode usr_items: %w", err) }
    return rows, nil
}

// SaveUserItemsJSON は usr_items.json として保存します。
func SaveUserItemsJSON(path string, rows []OwnItem) error {
    buf, err := json.MarshalIndent(rows, "", "  ")
    if err != nil { return err }
    return os.WriteFile(path, buf, 0644)
}
