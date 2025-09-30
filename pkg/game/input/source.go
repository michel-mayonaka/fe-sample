package input

// Source は物理入力から論理状態/イベントを供給する抽象です。
// Ebiten 等の具体 API はアダプタ側で扱い、本インタフェースは純粋な取得手段のみを表します。
type Source interface {
    // Poll は現在の論理状態のスナップショットを返します。
    Poll() ControlState
    // Events は直近フレームの素イベント列を返します（任意実装）。
    Events() []Event
}

