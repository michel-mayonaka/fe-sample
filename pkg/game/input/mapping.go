package input

// Layout はデバイスコードから Action へのデフォルト割当を保持します。
// コード値はアダプタ側で対応する実装依存の値へ変換して用います。
type Layout struct {
    Keyboard map[int]Action
    Mouse    map[int]Action
}

// 代表的なキーコード（抽象）。アダプタは実装依存のコードと対応付けます。
const (
    KeyArrowUp = 1
    KeyArrowDown = 2
    KeyArrowLeft = 3
    KeyArrowRight = 4
    KeyZ = 5
    KeyEnter = 6
    KeyX = 7
    KeyEscape = 8
    KeyA = 9
    KeyS = 10
    KeyTab = 11
    Key1 = 12
    Key2 = 13
    Key3 = 14
    KeyBackspace = 15
)

// マウスボタン（抽象）
const (
    MouseLeft = 100
)

// DefaultLayout は標準割当を返します（UI サンプルに合わせた最小セット）。
func DefaultLayout() Layout {
    return Layout{
        Keyboard: map[int]Action{
            KeyArrowUp:    ActionUp,
            KeyArrowDown:  ActionDown,
            KeyArrowLeft:  ActionLeft,
            KeyArrowRight: ActionRight,
            KeyZ:          ActionConfirm,
            KeyEnter:      ActionConfirm,
            KeyX:          ActionCancel,
            KeyEscape:     ActionCancel,
            KeyA:          ActionPrev,
            KeyS:          ActionNext,
            KeyTab:        ActionMenu,
            Key1:          ActionTerrainAtt1, // Shift 併用はアダプタ側で Def に切替
            Key2:          ActionTerrainAtt2,
            Key3:          ActionTerrainAtt3,
        },
        Mouse: map[int]Action{
            MouseLeft: ActionConfirm,
        },
    }
}
