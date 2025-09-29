// Package game は Scene/Actor/Service で構成される
// 軽量なゲームループ基盤を提供します。
package game

import (
	g_rng "ui_sample/internal/game/rng"
	service "ui_sample/internal/game/service"
	uinput "ui_sample/internal/game/ui/input"
)

// Ctx はフレーム共通の読み取り専用情報＋サービス群です。
// Update の呼び出し中は不変として扱います（入力は Snapshot 済み）。
type Ctx struct {
	DT      float64 // 経過秒（前フレーム→今フレーム）
	Frame   uint64  // 累積フレーム
	ScreenW int
	ScreenH int

	Input  uinput.Reader
	Assets *service.Assets
	Audio  *service.Audio
	Camera *service.Camera
	UI     *service.UI
	Rand   *g_rng.Rand
	Debug  bool
}
