package uicore

import (
	resourceFonts "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/opentype"
)

var (
	FaceTitle font.Face
	FaceMain  font.Face
	FaceSmall font.Face
)

func init() {
	if ft, err := opentype.Parse(resourceFonts.MPlus1pRegular_ttf); err == nil {
		if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 36, DPI: 96, Hinting: font.HintingNone}); err == nil {
			FaceTitle = f
		}
		if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 24, DPI: 96, Hinting: font.HintingNone}); err == nil {
			FaceMain = f
		}
		if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 18, DPI: 96, Hinting: font.HintingNone}); err == nil {
			FaceSmall = f
		}
	}
	if FaceTitle == nil {
		FaceTitle = basicfont.Face7x13
	}
	if FaceMain == nil {
		FaceMain = basicfont.Face7x13
	}
	if FaceSmall == nil {
		FaceSmall = basicfont.Face7x13
	}
}
