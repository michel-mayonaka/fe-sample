// Package userfs はユーザセーブ関連のファイル入出力を提供します（JSON バックエンド）。
package userfs

import (
	"encoding/json"
	"fmt"
	"os"
	usr "ui_sample/internal/model/user"
)

// LoadTableJSON はユーザテーブルJSONを読み込みます。
func LoadTableJSON(path string) (*usr.Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open user table: %w", err)
	}
	defer func() { _ = f.Close() }()
	var rows []usr.Character
	if err := json.NewDecoder(f).Decode(&rows); err != nil {
		return nil, fmt.Errorf("decode user table: %w", err)
	}
	return usr.NewTable(rows), nil
}

// SaveTableJSON はテーブル内容を JSON (インデント付き) で保存します。
func SaveTableJSON(path string, t *usr.Table) error {
	if t == nil {
		return fmt.Errorf("nil table")
	}
	b, err := json.MarshalIndent(t.Slice(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}

// LoadUserWeaponsJSON は usr_weapons.json を読み込みます。
func LoadUserWeaponsJSON(path string) ([]usr.OwnWeapon, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open usr_weapons: %w", err)
	}
	var rows []usr.OwnWeapon
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("decode usr_weapons: %w", err)
	}
	return rows, nil
}

// SaveUserWeaponsJSON は usr_weapons.json として保存します。
func SaveUserWeaponsJSON(path string, rows []usr.OwnWeapon) error {
	buf, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, buf, 0644)
}

// LoadUserItemsJSON は usr_items.json を読み込みます。
func LoadUserItemsJSON(path string) ([]usr.OwnItem, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open usr_items: %w", err)
	}
	var rows []usr.OwnItem
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("decode usr_items: %w", err)
	}
	return rows, nil
}

// SaveUserItemsJSON は usr_items.json として保存します。
func SaveUserItemsJSON(path string, rows []usr.OwnItem) error {
	buf, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, buf, 0644)
}
