// Package ui は、FE風ステータス画面のUI描画と
// それに付随するデータモデルを提供します。
package ui

import (
    "fmt"
    "image/color"
    "math"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    resourceFonts "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
    "github.com/hajimehoshi/ebiten/v2/text"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "ui_sample/internal/model"
    "ui_sample/internal/game"
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/font/opentype"
)

// パネル配色（FE風）。コントラスト高めの青系を基調とします。
var (
    colPanelBG   = color.RGBA{R: 0x20, G: 0x3b, B: 0x73, A: 0xFF}
    colPanelDark = color.RGBA{R: 0x14, G: 0x2a, B: 0x54, A: 0xFF}
    colBorder    = color.RGBA{R: 0xd9, G: 0xb9, B: 0x6e, A: 0xFF}
    colAccent    = color.RGBA{R: 0x7a, G: 0xc0, B: 0xff, A: 0xFF}
    colText      = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xFF}
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
    // 日本語を含む M+ 1p Regular を使用（Ebiten examples リソース）。
    if ft, err := opentype.Parse(resourceFonts.MPlus1pRegular_ttf); err == nil {
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

// Unit は1ユニット分の表示用データモデルです。
type Unit struct {
    Name  string         // 名前
    Class string         // クラス名
    Level int            // 現在レベル
    Exp   int            // 現在経験値
    HP    int            // 現在HP
    HPMax int            // 最大HP

    Stats Stats          // 能力値
    Equip []Item         // 装備（耐久制）

    Portrait *ebiten.Image // ポートレート画像（任意）

    Weapon WeaponRanks   // 物理系武器ランク
    Magic  MagicRanks    // 魔法系武器ランク
    Growth Growth        // 成長率（%）
}

// Stats は各種能力値を表します。
type Stats struct {
    Str, Mag, Skl, Spd, Lck, Def, Res, Mov int // 力/魔力/技/速さ/幸運/守備/魔防/移動
}

// Item は耐久制装備（武器・消耗品）を表します。
type Item struct {
    Name string // 名称
    Uses int    // 残り使用回数
    Max  int    // 使用可能数上限
}

// WeaponRanks は物理系武器のランクを表します。
type WeaponRanks struct {
    Sword string // 剣
    Lance string // 槍
    Axe   string // 斧
    Bow   string // 弓
}

// MagicRanks は魔法系のランクを表します。
type MagicRanks struct {
    Anima string // 理
    Light string // 光
    Dark  string // 闇
    Staff string // 杖
}

// Growth は各能力の成長率（%）を表します。
type Growth struct {
    Str, Mag, Skl, Spd, Lck, Def, Res, Mov int
}

// SampleUnit はサンプルとなるユニットデータを生成します。
func SampleUnit() Unit {
    // マスタから読めれば優先。失敗時は内蔵の簡易データを返す。
    if t, err := model.LoadFromJSON("assets/master/characters.json"); err == nil {
        if c, ok := t.Find("iris"); ok {
            return unitFromCharacter(c)
        }
    }
    // フォールバック
    u := Unit{
        Name:  "アイリス",
        Class: "ペガサスナイト",
        Level: 7,
        Exp:   56,
        HP:    22,
        HPMax: 26,
        Stats:  Stats{Str: 9, Mag: 0, Skl: 12, Spd: 14, Lck: 8, Def: 6, Res: 7, Mov: 7},
        Equip:  []Item{{Name: "アイアンランス", Uses: 35, Max: 45}, {Name: "ジャベリン", Uses: 12, Max: 20}, {Name: "傷薬", Uses: 3, Max: 3}},
        Weapon: WeaponRanks{Sword: "D", Lance: "B", Axe: "-", Bow: "-"},
        Magic:  MagicRanks{Anima: "-", Light: "-", Dark: "-", Staff: "-"},
        Growth: Growth{Str: 45, Mag: 10, Skl: 55, Spd: 65, Lck: 50, Def: 20, Res: 35, Mov: 0},
    }
    if img, _, err := ebitenutil.NewImageFromFile("assets/01_iris.png"); err == nil {
        u.Portrait = img
    }
    return u
}

// unitFromCharacter はマスタの Character を UI 用 Unit に変換します。
func unitFromCharacter(c model.Character) Unit {
    u := Unit{
        Name:  c.Name,
        Class: c.Class,
        Level: c.Level,
        Exp:   c.Exp,
        HP:    c.HP,
        HPMax: c.HPMax,
        Stats: Stats(c.Stats),
        Weapon: WeaponRanks{
            Sword: c.Weapon.Sword,
            Lance: c.Weapon.Lance,
            Axe:   c.Weapon.Axe,
            Bow:   c.Weapon.Bow,
        },
        Magic: MagicRanks{
            Anima: c.Magic.Anima,
            Light: c.Magic.Light,
            Dark:  c.Magic.Dark,
            Staff: c.Magic.Staff,
        },
        Growth: Growth(c.Growth),
    }
    // 装備変換
    for _, it := range c.Equip {
        u.Equip = append(u.Equip, Item{Name: it.Name, Uses: it.Uses, Max: it.Max})
    }
    // 画像
    if c.Portrait != "" {
        if img, _, err := ebitenutil.NewImageFromFile(c.Portrait); err == nil {
            u.Portrait = img
        }
    }
    return u
}

// DrawStatus はユニットのステータス画面を描画します。
// 渡された画像サイズ（例: 1920x1080）に合わせてレイアウトされます。
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
    text.Draw(dst, fmt.Sprintf("Lv %d / %d    経験値 %02d / %d", u.Level, game.LevelCap, u.Exp, game.LevelUpExp), faceMain, tx, ty+40+30, colText)

    // HP（行間は少し広め）
    text.Draw(dst, fmt.Sprintf("HP %d/%d", u.HP, u.HPMax), faceMain, tx, ty+40+30+40, colText)
    drawHPBar(dst, tx, ty+40+30+46, 300, 14, u.HP, u.HPMax)

    // 能力値（2カラム）
    statsTop := ty + 40 + 30 + 46 + 48
    line := 34
    colGap := 180
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+0*line, "力", u.Stats.Str, u.Growth.Str)
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+1*line, "魔力", u.Stats.Mag, u.Growth.Mag)
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+2*line, "技", u.Stats.Skl, u.Growth.Skl)
    drawStatLineWithGrowth(dst, faceMain, tx+0*colGap, statsTop+3*line, "速さ", u.Stats.Spd, u.Growth.Spd)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+0*line, "幸運", u.Stats.Lck, u.Growth.Lck)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+1*line, "守備", u.Stats.Def, u.Growth.Def)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+2*line, "魔防", u.Stats.Res, u.Growth.Res)
    drawStatLineWithGrowth(dst, faceMain, tx+1*colGap, statsTop+3*line, "移動", u.Stats.Mov, u.Growth.Mov)

    // 武器レベル（右側・上段）
    wrX := tx + 2*colGap + 64
    wrY := ty
    text.Draw(dst, "武器レベル", faceMain, wrX, wrY, colAccent)
    rline := 32
    drawRankLine(dst, faceMain, wrX, wrY+1*rline, "剣", u.Weapon.Sword)
    drawRankLine(dst, faceMain, wrX, wrY+2*rline, "槍", u.Weapon.Lance)
    drawRankLine(dst, faceMain, wrX, wrY+3*rline, "斧", u.Weapon.Axe)
    drawRankLine(dst, faceMain, wrX, wrY+4*rline, "弓", u.Weapon.Bow)

    // 魔法レベル（右側・下段）
    mrX := wrX
    mrY := wrY + (4+1)*rline + 16 // 見出し1行 + 武器4行 + 余白
    text.Draw(dst, "魔法レベル", faceMain, mrX, mrY, colAccent)
    drawRankLine(dst, faceMain, mrX, mrY+1*rline, "理", u.Magic.Anima)
    drawRankLine(dst, faceMain, mrX, mrY+2*rline, "光", u.Magic.Light)
    drawRankLine(dst, faceMain, mrX, mrY+3*rline, "闇", u.Magic.Dark)
    drawRankLine(dst, faceMain, mrX, mrY+4*rline, "杖", u.Magic.Staff)

    // 装備（ポートレートの下段）
    equipTitleY := int(py + ph + 56)
    text.Draw(dst, "装備", faceMain, int(px), equipTitleY, colAccent)
    for i, it := range u.Equip {
        lineY := equipTitleY + 30 + i*30
        // 名称
        text.Draw(dst, fmt.Sprintf("- %s", it.Name), faceSmall, int(px)+14, lineY, colText)
        // 耐久（右寄せ目安のカラム位置）
        uses := "-"
        if it.Max > 0 { uses = fmt.Sprintf("%d/%d", it.Uses, it.Max) }
        text.Draw(dst, uses, faceSmall, int(px)+300, lineY, colAccent)
    }
}

// drawPanel は立体感のあるパネル（外枠・影付き）を描画します。
func drawPanel(dst *ebiten.Image, x, y, w, h float32) {
    // 外枠
    vector.DrawFilledRect(dst, x-2, y-2, w+4, h+4, colBorder, false)
    // 影
    vector.DrawFilledRect(dst, x+2, y+2, w, h, colPanelDark, false)
    // 内部
    vector.DrawFilledRect(dst, x, y, w, h, colPanelBG, false)
}

// drawFramedRect は金色の縁取りを持つ矩形を描画します。
func drawFramedRect(dst *ebiten.Image, x, y, w, h float32) {
    vector.DrawFilledRect(dst, x-2, y-2, w+4, h+4, colBorder, false)
    vector.DrawFilledRect(dst, x, y, w, h, color.RGBA{30, 45, 78, 255}, false)
}

// drawPortraitPlaceholder は画像未設定時のプレースホルダテキストを描画します。
func drawPortraitPlaceholder(dst *ebiten.Image, x, y, w, h float32) {
    text.Draw(dst, "画像なし", faceSmall, int(x+10), int(y+h/2), colAccent)
}

// drawPortrait はポートレート画像を枠内に等比縮小して描画します。
// 線形補間により縮小時のジャギーを低減します。
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

// drawHPBar はHP割合に応じたカラーで水平バーを描画します。
// x,y: 左上座標, w,h: バーサイズ, hp/max: 現在HP/最大HP。
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

// drawStatLine は単一の能力値ラベルと数値を描画します。
func drawStatLine(dst *ebiten.Image, face font.Face, x, y int, label string, v int) {
    text.Draw(dst, label, face, x, y, colText)
    text.Draw(dst, fmt.Sprintf("%2d", v), face, x+64, y, colAccent)
}

// drawStatLineWithGrowth は能力値と成長率(%)を並べて描画します。
func drawStatLineWithGrowth(dst *ebiten.Image, face font.Face, x, y int, label string, v int, g int) {
    text.Draw(dst, label, face, x, y, colText)
    text.Draw(dst, fmt.Sprintf("%2d", v), face, x+64, y, colAccent)
    text.Draw(dst, fmt.Sprintf("%d%%", g), faceSmall, x+120, y, colAccent)
}

// drawRankLine は武器/魔法ランクのラベルと値を描画します。
func drawRankLine(dst *ebiten.Image, face font.Face, x, y int, label, rank string) {
    if rank == "" { rank = "-" }
    text.Draw(dst, label, face, x, y, colText)
    text.Draw(dst, rank, face, x+120, y, colAccent)
}
