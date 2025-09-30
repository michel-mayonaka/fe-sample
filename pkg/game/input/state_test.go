package input

import "testing"

func TestControlState_EqualAndGetSet(t *testing.T) {
    var a, b ControlState
    if !a.Equal(b) {
        t.Fatalf("zero states should be equal")
    }
    a.Set(ActionConfirm, true)
    if a.Equal(b) {
        t.Fatalf("states must differ after Set")
    }
    if !a.Get(ActionConfirm) || b.Get(ActionConfirm) {
        t.Fatalf("Get should reflect Set")
    }
    b.Set(ActionConfirm, true)
    if !a.Equal(b) {
        t.Fatalf("states should be equal after matching Set")
    }
}

func TestEdgeReader_PressDown(t *testing.T) {
    var r EdgeReader
    r.Step(ControlState{}) // frame 1: idle
    if r.Press(ActionConfirm) || r.Down(ActionConfirm) {
        t.Fatalf("unexpected press/down on idle")
    }
    // rising edge
    var s ControlState
    s.Set(ActionConfirm, true)
    r.Step(s)
    if !r.Press(ActionConfirm) || !r.Down(ActionConfirm) {
        t.Fatalf("expected press and down on rising edge")
    }
    // hold
    r.Step(s)
    if r.Press(ActionConfirm) || !r.Down(ActionConfirm) {
        t.Fatalf("no press while held; down should remain true")
    }
    // release
    var z ControlState
    r.Step(z)
    if r.Press(ActionConfirm) || r.Down(ActionConfirm) {
        t.Fatalf("no press/down on release")
    }
}

