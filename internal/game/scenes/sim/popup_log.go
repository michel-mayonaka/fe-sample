package sim

import (
	"github.com/hajimehoshi/ebiten/v2"
	"ui_sample/internal/game"
	scenes "ui_sample/internal/game/scenes"
	uidraw "ui_sample/internal/game/ui/draw"
	uinput "ui_sample/internal/game/ui/input"
)

// LogView は戦闘ログのポップアップビューです。
// Host の logs を描画し、Confirm でポップアップを閉じます。
type LogView struct {
	Host   *Sim
	sw, sh int
}

// NewLogView は戦闘ログのポップアップビューを生成します。
func NewLogView(host *Sim) *LogView { return &LogView{Host: host} }

// Update はポップアップ自身の入力処理を行います。
func (v *LogView) Update(ctx *game.Ctx) (game.Scene, error) {
	v.sw, v.sh = ctx.ScreenW, ctx.ScreenH
	intents := v.scHandleInput(ctx)
	v.scAdvance(intents)
	v.scFlush(ctx)
	return nil, nil
}

// Draw はログポップアップを描画します。
func (v *LogView) Draw(dst *ebiten.Image) { uidraw.DrawBattleLogOverlay(dst, v.Host.logs) }

// --- 内部: scHandleInput → scAdvance → scFlush --------------------------------------

type lvIntentKind int

const (
	lvNone lvIntentKind = iota
	lvClose
)

type lvIntent struct{ Kind lvIntentKind }

func (lvIntent) IsSceneIntent() {}

func (v *LogView) scHandleInput(ctx *game.Ctx) []scenes.Intent {
	intents := make([]scenes.Intent, 0, 1)
	if ctx != nil && ctx.Input != nil {
		if ctx.Input.Press(uinput.Confirm) {
			intents = append(intents, lvIntent{Kind: lvClose})
		}
	}
	return intents
}

func (v *LogView) scAdvance(intents []scenes.Intent) {
	for _, any := range intents {
		it, ok := any.(lvIntent)
		if !ok {
			continue
		}
		if it.Kind == lvClose {
			if v.Host != nil {
				v.Host.logPopup = false
			}
		}
	}
}

func (v *LogView) scFlush(_ *game.Ctx) { /* 今はなし */ }
