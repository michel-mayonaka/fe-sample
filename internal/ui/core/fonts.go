// Package uicore は UI の共通描画基盤（フォント/色/レイアウト等）を提供します。
package uicore

import (
    "math"
    resourceFonts "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/font/opentype"
)

// FaceTitle/FaceMain/FaceSmall は現在のスケールに応じた描画用フォントです。
var (
    FaceTitle font.Face
    FaceMain  font.Face
    FaceSmall font.Face

    baseFont   *opentype.Font
    lastScaled float32
)

const (
    baseTitlePt = 36
    baseMainPt  = 24
    baseSmallPt = 18
)

func init() {
    // 初期化時にベースフォントをパース
    if ft, err := opentype.Parse(resourceFonts.MPlus1pRegular_ttf); err == nil {
        baseFont = ft
    }
    // デフォルトの固定サイズで Faces を作成（初期ウィンドウ想定）
    if baseFont != nil {
        setFacesForScale(1.0)
    }
    // フォールバック
    if FaceTitle == nil { FaceTitle = basicfont.Face7x13 }
    if FaceMain == nil { FaceMain = basicfont.Face7x13 }
    if FaceSmall == nil { FaceSmall = basicfont.Face7x13 }
}

// MaybeUpdateFontFaces はスケール変化が十分大きい場合にフォントサイズを更新します。
func MaybeUpdateFontFaces() {
    s := CurrentScale()
    if baseFont == nil {
        // basicfont のまま（スケール非対応）
        return
    }
    // しきい値（±0.05）以上の変化で更新
    if math.Abs(float64(s-lastScaled)) < 0.05 {
        return
    }
    setFacesForScale(s)
}

func setFacesForScale(s float32) {
    if s <= 0 { s = 1 }
    // 過度な縮小/拡大を緩和
    szTitle := clampPt(float64(baseTitlePt)*float64(s), 12, 96)
    szMain  := clampPt(float64(baseMainPt)*float64(s), 10, 64)
    szSmall := clampPt(float64(baseSmallPt)*float64(s), 8,  48)
    if f, err := opentype.NewFace(baseFont, &opentype.FaceOptions{Size: szTitle, DPI: 96, Hinting: font.HintingNone}); err == nil {
        FaceTitle = f
    }
    if f, err := opentype.NewFace(baseFont, &opentype.FaceOptions{Size: szMain, DPI: 96, Hinting: font.HintingNone}); err == nil {
        FaceMain = f
    }
    if f, err := opentype.NewFace(baseFont, &opentype.FaceOptions{Size: szSmall, DPI: 96, Hinting: font.HintingNone}); err == nil {
        FaceSmall = f
    }
    lastScaled = s
}

func clampPt(v, lo, hi float64) float64 {
    if v < lo { return lo }
    if v > hi { return hi }
    return v
}
