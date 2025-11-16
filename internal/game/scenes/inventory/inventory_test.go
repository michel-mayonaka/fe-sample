package inventory

import (
    "testing"
    "ui_sample/internal/game"
    "ui_sample/internal/game/scenes"
    uicore "ui_sample/internal/game/service/ui"
    uinput "ui_sample/internal/game/ui/input"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
)

type fakeReader struct{ press map[uinput.Action]bool }
func (f fakeReader) Press(a uinput.Action) bool { return f.press[a] }
func (f fakeReader) Down(a uinput.Action) bool  { return f.press[a] }

func TestBackButtonClick_PopsInventory(t *testing.T) {
    env := &scenes.Env{Session: &scenes.Session{Units: []uicore.Unit{{ID:"u1"}}, SelIndex: 0}}
    s := NewInventory(env)
    s.sw, s.sh = 1920, 1080
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    cx, cy := bx + bw/2, by + bh/2
    r := fakeReader{press: map[uinput.Action]bool{uinput.Confirm: true}}
    ctx := &game.Ctx{ScreenW: s.sw, ScreenH: s.sh, CursorX: cx, CursorY: cy, Input: r}
    intents := s.scHandleInput(ctx)
    s.scAdvance(intents)
    if !s.ShouldPop() {
        t.Fatalf("inventory back button did not trigger pop")
    }
}

