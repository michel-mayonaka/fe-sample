package ebiteninput

import (
    "github.com/hajimehoshi/ebiten/v2"
    ginput "ui_sample/pkg/game/input"
)

// Source は Ebiten 由来の物理入力をドメインの ControlState/Event に変換します。
type Source struct {
    layout ginput.Layout
}

// NewSource は指定レイアウトでソースを生成します。
func NewSource(layout ginput.Layout) *Source { return &Source{layout: layout} }

// Poll は現在のキー/マウス状態をレイアウトに基づき論理状態へ投影します。
func (s *Source) Poll() ginput.ControlState {
    var st ginput.ControlState
    // キーボード
    for key, act := range s.layout.Keyboard {
        if ebiten.IsKeyPressed(toEbitenKey(key)) {
            st.Set(act, true)
        }
    }
    // シフト修飾で地形を防御側へ切替（1/2/3 のみ特別扱い）
    shift := ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight)
    if shift {
        if ebiten.IsKeyPressed(ebiten.Key1) {
            st.Set(ginput.ActionTerrainAtt1, false)
            st.Set(ginput.ActionTerrainDef1, true)
        }
        if ebiten.IsKeyPressed(ebiten.Key2) {
            st.Set(ginput.ActionTerrainAtt2, false)
            st.Set(ginput.ActionTerrainDef2, true)
        }
        if ebiten.IsKeyPressed(ebiten.Key3) {
            st.Set(ginput.ActionTerrainAtt3, false)
            st.Set(ginput.ActionTerrainDef3, true)
        }
    }
    // マウス
    for btn, act := range s.layout.Mouse {
        if mouseButtonDown(btn) {
            st.Set(act, true)
        }
    }
    return st
}

// Events は最小実装として nil を返します（必要時拡張）。
func (s *Source) Events() []ginput.Event { return nil }

// Position は現在のマウス座標を返します。
func (s *Source) Position() (int, int) { return ebiten.CursorPosition() }

// mouseButtonDown は抽象ボタンコードを Ebiten へ解決します。
func mouseButtonDown(btn int) bool {
    switch btn {
    case ginput.MouseLeft:
        return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
    default:
        return false
    }
}

// toEbitenKey は抽象キーコードを Ebiten のキーへ変換します。
func toEbitenKey(code int) ebiten.Key {
    switch code {
    case ginput.KeyArrowUp:
        return ebiten.KeyArrowUp
    case ginput.KeyArrowDown:
        return ebiten.KeyArrowDown
    case ginput.KeyArrowLeft:
        return ebiten.KeyArrowLeft
    case ginput.KeyArrowRight:
        return ebiten.KeyArrowRight
    case ginput.KeyZ:
        return ebiten.KeyZ
    case ginput.KeyEnter:
        return ebiten.KeyEnter
    case ginput.KeyX:
        return ebiten.KeyX
    case ginput.KeyEscape:
        return ebiten.KeyEscape
    case ginput.KeyA:
        return ebiten.KeyA
    case ginput.KeyS:
        return ebiten.KeyS
    case ginput.KeyTab:
        return ebiten.KeyTab
    case ginput.Key1:
        return ebiten.Key1
    case ginput.Key2:
        return ebiten.Key2
    case ginput.Key3:
        return ebiten.Key3
    case ginput.KeyBackspace:
        return ebiten.KeyBackspace
    default:
        return ebiten.Key(-1)
    }
}
