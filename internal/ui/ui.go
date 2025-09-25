package ui

import (
    "fmt"
    "image/color"
    "math"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/text"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/font/gofont/goregular"
    "golang.org/x/image/font/opentype"
)

// パネル配色（FE風）
var (
    colPanelBG   = color.RGBA{R: 0x20, G: 0x3b, B: 0x73, A: 0xFF}
    colPanelDark = color.RGBA{R: 0x14, G: 0x2a, B: 0x54, A: 0xFF}
    colBorder    = color.RGBA{R: 0xd9, G: 0xb9, B: 0x6e, A: 0xFF}
    colAccent    = color.RGBA{R: 0x7a, G: 0xc0, B: 0xff, A: 0xFF}
    colText      = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xFF}
)

var (
    faceTitle font.Face
    faceMain  font.Face
    faceSmall font.Face
)

func init() {
    // Go Regular フォントを大きめサイズで初期化。失敗時は basicfont にフォールバック。
    if ft, err := opentype.Parse(goregular.TTF); err == nil {
        if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 36, DPI: 96, Hinting: font.HintingNone}); err == nil {
            faceTitle = f
        }
        if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 24, DPI: 96, Hinting: font.HintingNone}); err == nil {
            faceMain = f
        }
        if f, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 18, DPI: 96, Hinting: font.HintingNone}); err == nil {
            faceSmall = f
        }
    }
    if faceTitle == nil {
        faceTitle = basicfont.Face7x13
    }
    if faceMain == nil {
        faceMain = basicfont.Face7x13
    }
    if faceSmall == nil {
        faceSmall = basicfont.Face7x13
    }
}

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

    Ranks WeaponRanks
}

type Stats struct {
    Str, Mag, Skl, Spd, Lck, Def, Res, Mov int
}

// 武器レベル（FE 風）
type WeaponRanks struct {
    Sword string
    Lance string
    Axe   string
    Bow   string
    Tome  string
    Staff string
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
        Ranks: WeaponRanks{Sword: "D", Lance: "B", Axe: "-", Bow: "-", Tome: "-", Staff: "-"},
    }
    if img, _, err := ebitenutil.NewImageFromFile("assets/01_iris.png"); err == nil {
        u.Portrait = img
    }
    return u
}

// DrawStatus はメインのステータス画面を描画。
func DrawStatus(dst *ebiten.Image, u Unit) {
    // 画面サイズに合わせたパネル
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    margin := float32(24)
    panelX, panelY := margin, margin
    panelW, panelH := float32(sw)-margin*2, float32(sh)-margin*2
    drawPanel(dst, panelX, panelY, panelW, panelH)

    // 左: ポートレート
    px, py := panelX+24, panelY+24
    pw, ph := float32(320), float32(320)
    drawFramedRect(dst, px, py, pw, ph)
    if u.Portrait != nil {
        drawPortrait(dst, u.Portrait, px, py, pw, ph)
    } else {
        drawPortraitPlaceholder(dst, px, py, pw, ph)
    }

    // 上: 名前/クラス/レベル
    tx := int(px + pw + 32)
    ty := int(py + 44)
    text.Draw(dst, u.Name, faceTitle, tx, ty, colAccent)
    text.Draw(dst, u.Class, faceMain, tx, ty+40, colText)
    text.Draw(dst, fmt.Sprintf("Lv %d  EXP %02d", u.Level, u.Exp), faceMain, tx, ty+40+30, colText)

    // HP
    text.Draw(dst, fmt.Sprintf("HP %d/%d", u.HP, u.HPMax), faceMain, tx, ty+40+30+40, colText)
    drawHPBar(dst, tx, ty+40+30+46, 600, 14, u.HP, u.HPMax)

    // 能力値（2カラム）
    statsTop := ty + 40 + 30 + 46 + 40
    line := 28
    colGap := 180
    drawStatLine(dst, faceMain, tx+0*colGap, statsTop+0*line, "STR", u.Stats.Str)
    drawStatLine(dst, faceMain, tx+0*colGap, statsTop+1*line, "MAG", u.Stats.Mag)
    drawStatLine(dst, faceMain, tx+0*colGap, statsTop+2*line, "SKL", u.Stats.Skl)
    drawStatLine(dst, faceMain, tx+0*colGap, statsTop+3*line, "SPD", u.Stats.Spd)
    drawStatLine(dst, faceMain, tx+1*colGap, statsTop+0*line, "LCK", u.Stats.Lck)
    drawStatLine(dst, faceMain, tx+1*colGap, statsTop+1*line, "DEF", u.Stats.Def)
    drawStatLine(dst, faceMain, tx+1*colGap, statsTop+2*line, "RES", u.Stats.Res)
    drawStatLine(dst, faceMain, tx+1*colGap, statsTop+3*line, "MOV", u.Stats.Mov)

    // 武器レベル（右側）
    wrX := tx + 2*colGap + 64
    wrY := ty
    text.Draw(dst, "Weapon Rank", faceMain, wrX, wrY, colAccent)
    rline := 28
    drawRankLine(dst, faceMain, wrX, wrY+1*rline, "Sword", u.Ranks.Sword)
    drawRankLine(dst, faceMain, wrX, wrY+2*rline, "Lance", u.Ranks.Lance)
    drawRankLine(dst, faceMain, wrX, wrY+3*rline, "Axe", u.Ranks.Axe)
    drawRankLine(dst, faceMain, wrX, wrY+4*rline, "Bow", u.Ranks.Bow)
    drawRankLine(dst, faceMain, wrX, wrY+5*rline, "Tome", u.Ranks.Tome)
    drawRankLine(dst, faceMain, wrX, wrY+6*rline, "Staff", u.Ranks.Staff)

    // 装備（ポートレートの下段）
    equipTitleY := int(py + ph + 48)
    text.Draw(dst, "Equipment", faceMain, int(px), equipTitleY, colAccent)
    for i, it := range u.Equip {
        text.Draw(dst, fmt.Sprintf("- %s", it), faceSmall, int(px)+14, equipTitleY+30+i*26, colText)
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
    text.Draw(dst, "No Portrait", faceSmall, int(x+10), int(y+h/2), colAccent)
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
    op.Filter = ebiten.FilterLinear
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

func drawStatLine(dst *ebiten.Image, face font.Face, x, y int, label string, v int) {
    text.Draw(dst, label, face, x, y, colText)
    text.Draw(dst, fmt.Sprintf("%2d", v), face, x+64, y, colAccent)
}

func drawRankLine(dst *ebiten.Image, face font.Face, x, y int, label, rank string) {
    if rank == "" { rank = "-" }
    text.Draw(dst, label, face, x, y, colText)
    text.Draw(dst, rank, face, x+120, y, colAccent)
}
