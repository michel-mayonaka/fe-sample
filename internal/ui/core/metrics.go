package uicore

import "github.com/hajimehoshi/ebiten/v2"

var baseW, baseH int = 1920, 1080
var scale float32 = 1.0

// SetBaseResolution は論理解像度（Layoutの返すサイズ）を設定します。
func SetBaseResolution(w, h int) { if w > 0 && h > 0 { baseW, baseH = w, h } }

// UpdateMetricsFromWindow は現在のウィンドウサイズからスケールを算出します。
func UpdateMetricsFromWindow() {
    w, h := ebiten.WindowSize()
    if w <= 0 || h <= 0 || baseW <= 0 || baseH <= 0 { scale = 1; return }
    sx := float32(w) / float32(baseW)
    sy := float32(h) / float32(baseH)
    if sx < sy { scale = sx } else { scale = sy }
    if scale <= 0 { scale = 1 }
}

// S は整数ピクセル値をスケールに応じて拡縮します。
func S(n int) int { return int(float32(n) * scale) }

// 代表的な行高/マージンのスケール済み値。
func ListMarginPx() int      { return S(ListMargin) }
func ListItemHPx() int       { return S(ListItemH) }
func ListItemGapPx() int     { return S(ListItemGap) }
func ListPortraitSzPx() int  { return S(ListPortraitSz) }
func ListTitleOffsetPx() int { return S(ListTitleOffset) }
func LineHMainPx() int       { return S(LineHMain) }
func LineHSmallPx() int      { return S(LineHSmall) }

// CurrentScale は現在のスケール値を返します。
func CurrentScale() float32 { return scale }
