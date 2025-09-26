package ui

import (
    resourceFonts "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/font/opentype"
)

// 画面内で用いるフォントフェイス群。
// Title: 見出し、Main: 本文、Small: 注釈やサブ情報に使用します。
var (
    faceTitle font.Face
    faceMain  font.Face
    faceSmall font.Face
)

// init は日本語フォント（M+ 1p）を初期化します。
// 失敗時は basicfont にフォールバックします。
func init() {
    if ft, err := opentype.Parse(resourceFonts.MPlus1pRegular_ttf); err == nil {
        if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 36, DPI: 96, Hinting: font.HintingNone}); err == nil { faceTitle = f }
        if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 24, DPI: 96, Hinting: font.HintingNone}); err == nil { faceMain = f }
        if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 18, DPI: 96, Hinting: font.HintingNone}); err == nil { faceSmall = f }
    }
    if faceTitle == nil { faceTitle = basicfont.Face7x13 }
    if faceMain == nil { faceMain = basicfont.Face7x13 }
    if faceSmall == nil { faceSmall = basicfont.Face7x13 }
}

