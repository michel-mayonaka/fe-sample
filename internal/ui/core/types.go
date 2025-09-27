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

// Stats は表示用の基礎能力値です。
type Stats struct{ Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int }

// Item は装備/アイテムの表示用スナップショットです。
type Item struct {
    ID        string
    Name      string
    Uses, Max int
}

// WeaponRanks は武器熟練度の表示用ランクです。
type WeaponRanks struct{ Sword, Lance, Axe, Bow string }

// MagicRanks は魔法熟練度の表示用ランクです。
type MagicRanks struct{ Anima, Light, Dark, Staff string }

// Growth は成長率（%）の表示用値です。
type Growth struct{ HP, Str, Mag, Skl, Spd, Lck, Def, Res, Mov, Bld int }
