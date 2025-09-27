package uicore

import "github.com/hajimehoshi/ebiten/v2"

// Unit は1ユニット分の表示用データモデルです。
type Unit struct {
    ID, Name, Class string
    Level, Exp      int
    HP, HPMax       int
    Stats           Stats
    Equip           []Item
    Weapon          WeaponRanks
    Magic           MagicRanks
    Growth          Growth
    Portrait        *ebiten.Image
}

type Stats struct{ Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int }
type Item struct {
	Name      string
	Uses, Max int
}
type WeaponRanks struct{ Sword, Lance, Axe, Bow string }
type MagicRanks struct{ Anima, Light, Dark, Staff string }
type Growth struct{ HP, Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int }
