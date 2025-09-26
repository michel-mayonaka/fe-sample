package ui

import "github.com/hajimehoshi/ebiten/v2"

// Unit は1ユニット分の表示用データモデルです。
type Unit struct {
    ID    string
    Name  string
    Class string
    Level int
    Exp   int
    HP    int
    HPMax int

    Stats  Stats
    Equip  []Item
    Weapon WeaponRanks
    Magic  MagicRanks
    Growth Growth

    Portrait *ebiten.Image
}

// Stats は各種能力値を表します。
type Stats struct { Str, Mag, Skl, Spd, Lck, Def, Res, Mov int }

// Item は耐久制装備（武器・消耗品）を表します。
type Item struct { Name string; Uses, Max int }

// WeaponRanks は物理系武器のランクを表します。
type WeaponRanks struct { Sword, Lance, Axe, Bow string }

// MagicRanks は魔法系のランクを表します。
type MagicRanks struct { Anima, Light, Dark, Staff string }

// Growth は各能力の成長率（%）を表します。
type Growth struct { HP, Str, Mag, Skl, Spd, Lck, Def, Res, Mov int }

