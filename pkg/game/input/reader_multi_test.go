package input

import "testing"

// 複数アクションの同時押下と順序の検証
func TestEdgeReader_MultiActions(t *testing.T) {
    var r EdgeReader
    // 初期
    r.Step(ControlState{})
    if r.Down(ActionLeft) || r.Down(ActionRight) || r.Press(ActionConfirm) {
        t.Fatalf("unexpected state on init")
    }
    // 左+右 同時押下（ありえないが辺検出の健全性確認）
    s := ControlState{}
    s.Set(ActionLeft, true)
    s.Set(ActionRight, true)
    r.Step(s)
    if !r.Press(ActionLeft) || !r.Press(ActionRight) {
        t.Fatalf("expected press on both edges")
    }
    // 次フレーム: 右だけ離す
    s2 := s
    s2.Set(ActionRight, false)
    r.Step(s2)
    if r.Press(ActionLeft) || !r.Down(ActionLeft) {
        t.Fatalf("left should be held without press")
    }
    if r.Down(ActionRight) || r.Press(ActionRight) {
        t.Fatalf("right should be released")
    }
    // 次フレーム: Confirm 単押し
    s3 := ControlState{}
    s3.Set(ActionConfirm, true)
    r.Step(s3)
    if !r.Press(ActionConfirm) || !r.Down(ActionConfirm) {
        t.Fatalf("confirm press/down expected")
    }
}

