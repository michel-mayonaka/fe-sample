package service

// Camera はシンプルな2Dカメラ（平行移動＋スケール）です。
type Camera struct {
    X, Y  float64
    Scale float64
}
// NewCamera はデフォルト倍率のカメラを生成します。
func NewCamera() *Camera { return &Camera{Scale: 1} }
