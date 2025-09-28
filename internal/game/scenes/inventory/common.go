package inventory

import (
    "github.com/hajimehoshi/ebiten/v2"
)

// OwnerBadge は所有者名とポートレートの簡易表示情報です（popup共通）。
type OwnerBadge struct {
    Name     string
    Portrait *ebiten.Image
}

