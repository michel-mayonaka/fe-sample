package input

// Reader はフレーム間の辺（立ち上がり）検出と押下状態の問い合わせを提供します。
type Reader interface {
    // Press はこのフレームで新たに押されたか（立ち上がり）を返します。
    Press(Action) bool
    // Down は押下中か（継続）を返します。
    Down(Action) bool
}

// EdgeReader は ControlState の遷移から Press/Down を判定する実装です。
type EdgeReader struct {
    prev, curr ControlState
}

// Step は次のフレームの状態を供給し、内部状態を更新します。
func (r *EdgeReader) Step(next ControlState) {
    r.prev = r.curr
    r.curr = next
}

// Press はこのフレームで新たに押されたか（prev=false, curr=true）を返します。
func (r *EdgeReader) Press(a Action) bool { return !r.prev.Get(a) && r.curr.Get(a) }

// Down は押下中か（curr=true）を返します。
func (r *EdgeReader) Down(a Action) bool { return r.curr.Get(a) }

