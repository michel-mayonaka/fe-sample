package view

import "github.com/hajimehoshi/ebiten/v2"

// OwnerBadge は所有者名とポートレートの簡易表示情報（UI表示用）です。
type OwnerBadge struct {
	Name     string
	Portrait *ebiten.Image
}

// ItemRow はアイテム一覧行（プレゼンテーション用の結合データ）です。
type ItemRow struct {
	ID         string
	Name, Type string
	Effect     string
	Power      int
	Uses, Max  int
	Owners     []OwnerBadge
}

// WeaponRow は武器一覧行（プレゼンテーション用の結合データ）です。
type WeaponRow struct {
	ID                 string
	Name, Type, Rank   string
	Might, Hit, Crit   int
	Weight             int
	RangeMin, RangeMax int
	Uses, Max          int
	Owners             []OwnerBadge
}
