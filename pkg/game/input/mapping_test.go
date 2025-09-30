package input

import "testing"

func TestDefaultLayout_ContainsExpectedMappings(t *testing.T) {
    lt := DefaultLayout()
    // 矢印
    if lt.Keyboard[KeyArrowUp] != ActionUp || lt.Keyboard[KeyArrowDown] != ActionDown || lt.Keyboard[KeyArrowLeft] != ActionLeft || lt.Keyboard[KeyArrowRight] != ActionRight {
        t.Fatalf("arrow keys mapping mismatch: %#v", lt.Keyboard)
    }
    // 決定/キャンセル
    if lt.Keyboard[KeyZ] != ActionConfirm || lt.Keyboard[KeyEnter] != ActionConfirm {
        t.Fatalf("confirm mapping mismatch: %#v", lt.Keyboard)
    }
    if lt.Keyboard[KeyX] != ActionCancel || lt.Keyboard[KeyEscape] != ActionCancel {
        t.Fatalf("cancel mapping mismatch: %#v", lt.Keyboard)
    }
    // メニュー/次/前
    if lt.Keyboard[KeyTab] != ActionMenu || lt.Keyboard[KeyS] != ActionNext || lt.Keyboard[KeyA] != ActionPrev {
        t.Fatalf("menu/prev/next mapping mismatch: %#v", lt.Keyboard)
    }
    // 地形（攻撃側: 1..3）- Shift の防御側切替はアダプタ側で扱う
    if lt.Keyboard[Key1] != ActionTerrainAtt1 || lt.Keyboard[Key2] != ActionTerrainAtt2 || lt.Keyboard[Key3] != ActionTerrainAtt3 {
        t.Fatalf("terrain attack mapping mismatch: %#v", lt.Keyboard)
    }
    // マウス左=Confirm
    if lt.Mouse[MouseLeft] != ActionConfirm {
        t.Fatalf("mouse left mapping mismatch: %#v", lt.Mouse)
    }
}

