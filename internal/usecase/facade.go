// Package usecase はユースケース層（アプリケーションサービス）を提供します。
package usecase

import (
	"math/rand"
	"ui_sample/internal/repo"
)

// App はユースケースの最小集合を束ねます。
type App struct {
	Weapons repo.WeaponsRepo
	Users   repo.UserRepo
	Inv     repo.InventoryRepo
	RNG     *rand.Rand
}

// New はリポジトリと乱数源を注入して App を生成します。
func New(users repo.UserRepo, weapons repo.WeaponsRepo, inv repo.InventoryRepo, rng *rand.Rand) *App {
	return &App{Weapons: weapons, Users: users, Inv: inv, RNG: rng}
}
