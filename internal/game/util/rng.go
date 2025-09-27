package util

import "math/rand"

// Rand はテスト容易性のための薄いラッパです。
type Rand struct{ r *rand.Rand }

// New は seed から新しい Rand を作成します。
func New(seed int64) *Rand { return &Rand{r: rand.New(rand.NewSource(seed))} }

// Intn は [0,n) を返します。
func (x *Rand) Intn(n int) int { return x.r.Intn(n) }

// Float64 は [0,1) を返します。
func (x *Rand) Float64() float64 { return x.r.Float64() }

