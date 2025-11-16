// Package usecase はユースケース層（アプリケーションサービス）を提供します。
package usecase

import (
	"math/rand"
	"ui_sample/internal/repo"
)

// App はユースケースの最小集合を束ねます。
type App struct {
	Weapons repo.WeaponsRepo
	Items   repo.ItemsRepo
	Users   repo.UserRepo
	Inv     repo.InventoryRepo
	RNG     *rand.Rand
}

// New はリポジトリと乱数源を注入して App を生成します。
func New(users repo.UserRepo, weapons repo.WeaponsRepo, items repo.ItemsRepo, inv repo.InventoryRepo, rng *rand.Rand) *App {
	return &App{Weapons: weapons, Items: items, Users: users, Inv: inv, RNG: rng}
}
