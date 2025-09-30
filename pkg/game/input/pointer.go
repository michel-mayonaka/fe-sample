package input

// Pointer はポイントデバイス（マウス等）の座標を提供します。
// 実装は実デバイス依存の取得を隠蔽します。
type Pointer interface {
    Position() (x, y int)
}

