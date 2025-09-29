package scenes

import (
	"ui_sample/internal/game"
)

// Intent は各シーンの入力を意味レベルに圧縮した共通インタフェースです。
// 具体型は各サブパッケージ（character_list, inventory など）が定義します。
type Intent interface{ IsSceneIntent() }

// Lifecycle は各シーンが備えるべき更新フック群です。
// 実行順: HandleInput → Advance → Flush。
type Lifecycle interface {
	HandleInput(ctx *game.Ctx) []Intent
	Advance(intents []Intent)
	Flush(ctx *game.Ctx)
}
