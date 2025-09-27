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

// ListMarginPx は一覧パネルのマージン（スケール適用後）を返します。
func ListMarginPx() int      { return S(ListMargin) }
// ListItemHPx は一覧行の高さ（スケール適用後）を返します。
func ListItemHPx() int       { return S(ListItemH) }
// ListItemGapPx は一覧行の間隔（スケール適用後）を返します。
func ListItemGapPx() int     { return S(ListItemGap) }
// ListPortraitSzPx はポートレート枠サイズ（スケール適用後）を返します。
func ListPortraitSzPx() int  { return S(ListPortraitSz) }
// ListTitleOffsetPx はタイトルのYオフセット（スケール適用後）を返します。
func ListTitleOffsetPx() int { return S(ListTitleOffset) }
// LineHMainPx は本文行の高さ（スケール適用後）を返します。
func LineHMainPx() int       { return S(LineHMain) }
// LineHSmallPx は小サイズ行の高さ（スケール適用後）を返します。
func LineHSmallPx() int      { return S(LineHSmall) }

// CurrentScale は現在のスケール値を返します。
func CurrentScale() float32 { return scale }
