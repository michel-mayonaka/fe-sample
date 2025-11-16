package sim

import (
    "testing"
    "ui_sample/internal/game"
    "ui_sample/internal/game/scenes"
    uicore "ui_sample/internal/game/service/ui"
    uinput "ui_sample/internal/game/ui/input"
)

type fakeReader struct{ press map[uinput.Action]bool }
func (f fakeReader) Press(a uinput.Action) bool { return f.press[a] }
func (f fakeReader) Down(a uinput.Action) bool  { return f.press[a] }

func TestCancel_PopsSimScene(t *testing.T) {
    env := &scenes.Env{Session: &scenes.Session{Units: []uicore.Unit{{ID:"att"},{ID:"def"}}, SelIndex: 0}}
    s := NewSim(env, env.Session.Units[0], env.Session.Units[1])
    s.sw, s.sh = 1920, 1080
    r := fakeReader{press: map[uinput.Action]bool{uinput.Cancel: true}}
    ctx := &game.Ctx{ScreenW: s.sw, ScreenH: s.sh, CursorX: 0, CursorY: 0, Input: r}
    intents := s.scHandleInput(ctx)
    s.scAdvance(intents)
    if !s.ShouldPop() {
        t.Fatalf("sim scene cancel did not trigger pop")
    }
}

func TestBackButtonClick_PopsSimScene(t *testing.T) {
    env := &scenes.Env{Session: &scenes.Session{Units: []uicore.Unit{{ID:"att"},{ID:"def"}}, SelIndex: 0}}
    s := NewSim(env, env.Session.Units[0], env.Session.Units[1])
    s.sw, s.sh = 1920, 1080
    // カーソルを戻るボタン内へ置く
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    cx, cy := bx + bw/2, by + bh/2
    r := fakeReader{press: map[uinput.Action]bool{uinput.Confirm: true}}
    ctx := &game.Ctx{ScreenW: s.sw, ScreenH: s.sh, CursorX: cx, CursorY: cy, Input: r}
    intents := s.scHandleInput(ctx)
    s.scAdvance(intents)
    if !s.ShouldPop() {
        t.Fatalf("sim back button click did not trigger pop")
    }
}

func TestAutoRunButton_Toggles(t *testing.T) {
    env := &scenes.Env{Session: &scenes.Session{Units: []uicore.Unit{{ID:"att", HP:10, HPMax:10},{ID:"def", HP:10, HPMax:10}}, SelIndex: 0}}
    s := NewSim(env, env.Session.Units[0], env.Session.Units[1])
    s.sw, s.sh = 1920, 1080
    // カーソルを自動実行ボタン上へ置いて Confirm
    ax, ay, aw, ah := uilayout.AutoRunButtonRect(s.sw, s.sh)
    cx, cy := ax + aw/2, ay + ah/2
    r := fakeReader{press: map[uinput.Action]bool{uinput.Confirm: true}}
    ctx := &game.Ctx{ScreenW: s.sw, ScreenH: s.sh, CursorX: cx, CursorY: cy, Input: r}
    intents := s.scHandleInput(ctx)
    // 前提: まだポップアップは出ていない
    if len(intents) == 0 {
        t.Fatalf("no intents captured for auto-run toggle")
    }
    s.scAdvance(intents)
    if !s.auto {
        t.Fatalf("auto-run was not toggled on")
    }
}
