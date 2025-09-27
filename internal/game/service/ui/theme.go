package uicore

import "image/color"

// テーマ色（UI共通）。
var (
	// ColPanelBG はパネルの基調色です。
	ColPanelBG   = color.RGBA{R: 0x20, G: 0x3b, B: 0x73, A: 0xFF}
	// ColPanelDark はパネルの影色です。
	ColPanelDark = color.RGBA{R: 0x14, G: 0x2a, B: 0x54, A: 0xFF}
	// ColBorder は金枠の色です。
	ColBorder    = color.RGBA{R: 0xd9, G: 0xb9, B: 0x6e, A: 0xFF}
	// ColAccent は強調テキスト色です。
	ColAccent    = color.RGBA{R: 0x7a, G: 0xc0, B: 0xff, A: 0xFF}
	// ColText は標準テキスト色です。
	ColText      = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xFF}
)
