package ui

import (
    "fmt"
    "image/color"
    "math"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/text"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "golang.org/x/image/font/basicfont"
)

// パネル配色（FE風）
var (
    colPanelBG   = color.RGBA{R: 0x20, G: 0x3b, B: 0x73, A: 0xFF}
    colPanelDark = color.RGBA{R: 0x14, G: 0x2a, B: 0x54, A: 0xFF}
    colBorder    = color.RGBA{R: 0xd9, G: 0xb9, B: 0x6e, A: 0xFF}
    colAccent    = color.RGBA{R: 0x7a, G: 0xc0, B: 0xff, A: 0xFF}
    colText      = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xFF}
)

// Unit はステータス表示用の単純なデータ。
type Unit struct {
    Name  string
    Class string
    Level int
    Exp   int
    HP    int
    HPMax int

    Stats Stats
    Equip []string

    Portrait *ebiten.Image
}

type Stats struct {
    Str, Mag, Skl, Spd, Lck, Def, Res, Mov int
}

// SampleUnit は画面用ダミーデータ。
func SampleUnit() Unit {
    u := Unit{
        Name:  "Iris",
        Class: "Pegasus Knight",
        Level: 7,
        Exp:   56,
        HP:    22,
        HPMax: 26,
        Stats: Stats{Str: 9, Mag: 0, Skl: 12, Spd: 14, Lck: 8, Def: 6, Res: 7, Mov: 7},
        Equip: []string{"Iron Lance", "Javelin", "Vulnerary"},
    }
    if img, _, err := ebitenutil.NewImageFromFile("assets/01_iris.png"); err == nil {
        u.Portrait = img
    }
    return u
}

// DrawStatus はメインのステータス画面を描画。
func DrawStatus(dst *ebiten.Image, u Unit) {
    // メインパネル
    drawPanel(dst, 8, 8, 304, 224)

    // 左: ポートレート枠（ダミー）
    drawFramedRect(dst, 16, 20, 96, 96)
    if u.Portrait != nil {
        drawPortrait(dst, u.Portrait, 16, 20, 96, 96)
    } else {
        drawPortraitPlaceholder(dst, 16, 20, 96, 96)
    }

    // 上: 名前/クラス/レベル
    face := basicfont.Face7x13
    text.Draw(dst, u.Name, face, 128, 36, colAccent)
    text.Draw(dst, u.Class, face, 128, 52, colText)
    text.Draw(dst, fmt.Sprintf("Lv %d  EXP %02d", u.Level, u.Exp), face, 128, 68, colText)

    // HP バー
    text.Draw(dst, fmt.Sprintf("HP %d/%d", u.HP, u.HPMax), face, 128, 88, colText)
    drawHPBar(dst, 128, 94, 168, 8, u.HP, u.HPMax)

    // 右: 能力値
    left := 128
    top := 120
    line := 14
    drawStatLine(dst, left, top+0*line, "STR", u.Stats.Str)
    drawStatLine(dst, left, top+1*line, "MAG", u.Stats.Mag)
    drawStatLine(dst, left, top+2*line, "SKL", u.Stats.Skl)
    drawStatLine(dst, left, top+3*line, "SPD", u.Stats.Spd)
    drawStatLine(dst, left, top+4*line, "LCK", u.Stats.Lck)
    drawStatLine(dst, left, top+5*line, "DEF", u.Stats.Def)
    drawStatLine(dst, left, top+6*line, "RES", u.Stats.Res)
    drawStatLine(dst, left, top+7*line, "MOV", u.Stats.Mov)

    // 下: 装備
    text.Draw(dst, "Equipment", face, 16, 132, colAccent)
    for i, it := range u.Equip {
        text.Draw(dst, fmt.Sprintf("- %s", it), face, 24, 150+i*14, colText)
    }
}

func drawPanel(dst *ebiten.Image, x, y, w, h float32) {
    // 外枠
    vector.DrawFilledRect(dst, x-2, y-2, w+4, h+4, colBorder, false)
    // 影
    vector.DrawFilledRect(dst, x+2, y+2, w, h, colPanelDark, false)
    // 内部
    vector.DrawFilledRect(dst, x, y, w, h, colPanelBG, false)
}

func drawFramedRect(dst *ebiten.Image, x, y, w, h float32) {
    vector.DrawFilledRect(dst, x-2, y-2, w+4, h+4, colBorder, false)
    vector.DrawFilledRect(dst, x, y, w, h, color.RGBA{30, 45, 78, 255}, false)
}

func drawPortraitPlaceholder(dst *ebiten.Image, x, y, w, h float32) {
    face := basicfont.Face7x13
    text.Draw(dst, "No Portrait", face, int(x+10), int(y+h/2), colAccent)
}

func drawPortrait(dst *ebiten.Image, img *ebiten.Image, x, y, w, h float32) {
    iw, ih := img.Size()
    if iw == 0 || ih == 0 { return }
    sx := float64(w) / float64(iw)
    sy := float64(h) / float64(ih)
    s := math.Min(sx, sy)
    sw := float64(iw) * s
    sh := float64(ih) * s
    tx := float64(x) + (float64(w)-sw)/2
    ty := float64(y) + (float64(h)-sh)/2
    var op ebiten.DrawImageOptions
    op.GeoM.Scale(s, s)
    op.GeoM.Translate(tx, ty)
    dst.DrawImage(img, &op)
}

func drawHPBar(dst *ebiten.Image, x, y, w, h int, hp, max int) {
    if max <= 0 { max = 1 }
    // 背景
    vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), color.RGBA{50, 50, 50, 255}, false)
    // 値
    ratio := float32(hp) / float32(max)
    bw := float32(w) * ratio
    col := color.RGBA{80, 220, 100, 255}
    if ratio < 0.33 {
        col = color.RGBA{220, 80, 80, 255}
    } else if ratio < 0.66 {
        col = color.RGBA{240, 200, 80, 255}
    }
    vector.DrawFilledRect(dst, float32(x), float32(y), bw, float32(h), col, false)
}

func drawStatLine(dst *ebiten.Image, x, y int, label string, v int) {
    face := basicfont.Face7x13
    text.Draw(dst, label, face, x, y, colText)
    text.Draw(dst, fmt.Sprintf("%2d", v), face, x+64, y, colAccent)
}
