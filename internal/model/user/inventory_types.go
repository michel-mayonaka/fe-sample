package user

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

