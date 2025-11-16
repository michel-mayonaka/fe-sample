package status

import (
    "testing"
    "ui_sample/internal/game"
    "ui_sample/internal/game/scenes"
    uicore "ui_sample/internal/game/service/ui"
    uinput "ui_sample/internal/game/ui/input"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
)

// fakeReader は uinput.Reader の最小モックです。
type fakeReader struct{ press map[uinput.Action]bool }
func (f fakeReader) Press(a uinput.Action) bool { return f.press[a] }
func (f fakeReader) Down(a uinput.Action) bool  { return f.press[a] }

func TestBackButtonClick_ReturnsToList(t *testing.T) {
    // Env/Session 準備（選択ユニット1件）
    env := &scenes.Env{Session: &scenes.Session{Units: []uicore.Unit{{ID:"u1", Name:"Alice"}}, SelIndex: 0}}
    s := NewStatus(env)
    s.sw, s.sh = 1920, 1080
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    cx, cy := bx + bw/2, by + bh/2 // ボタン内
    // Confirm を一度だけ発火
    r := fakeReader{press: map[uinput.Action]bool{uinput.Confirm: true}}
    ctx := &game.Ctx{ScreenW: s.sw, ScreenH: s.sh, CursorX: cx, CursorY: cy, Input: r}

    intents := s.scHandleInput(ctx)
    s.scAdvance(intents)
    if !s.ShouldPop() {
        t.Fatalf("back button click did not trigger pop")
    }
}

