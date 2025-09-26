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
	text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2 へ移行
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/opentype"
	"ui_sample/internal/game"
	"ui_sample/internal/model"
	"ui_sample/internal/user"
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
	ID    string // 識別子（ユーザテーブル ID）
	Name  string // 名前
	Class string // クラス名
	Level int    // 現在レベル
	Exp   int    // 現在経験値
	HP    int    // 現在HP
	HPMax int    // 最大HP

	Stats Stats  // 能力値
	Equip []Item // 装備（耐久制）

	Portrait *ebiten.Image // ポートレート画像（任意）

	Weapon WeaponRanks // 物理系武器ランク
	Magic  MagicRanks  // 魔法系武器ランク
	Growth Growth      // 成長率（%）
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
	HP, Str, Mag, Skl, Spd, Lck, Def, Res, Mov int
}

// SampleUnit はサンプルとなるユニットデータを生成します。
func SampleUnit() Unit {
	// ユーザテーブル（usr_）のみで構築。失敗時はフォールバック。
	if ut, err := user.LoadFromJSON("db/user/usr_characters.json"); err == nil {
		if uc, ok := ut.Find("iris"); ok {
			return unitFromUser(uc)
		}
	}
	// フォールバック
	u := Unit{
		Name:   "アイリス",
		Class:  "ペガサスナイト",
		Level:  7,
		Exp:    56,
		HP:     22,
		HPMax:  26,
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

// unitFromUser はユーザテーブルのキャラクタを UI 用へ変換します。
func unitFromUser(c user.Character) Unit {
	u := Unit{
		ID:    c.ID,
		Name:  c.Name,
		Class: c.Class,
		Level: c.Level,
		Exp:   c.Exp,
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
	for _, it := range c.Equip {
		u.Equip = append(u.Equip, Item{Name: it.Name, Uses: it.Uses, Max: it.Max})
	}
	if c.Portrait != "" {
		if img, _, err := ebitenutil.NewImageFromFile(c.Portrait); err == nil {
			u.Portrait = img
		}
	}
	// HPが未指定（0）の場合は最大値を初期値として採用
	if u.HP == 0 && u.HPMax > 0 {
		u.HP = u.HPMax
	} else {
		u.HP = c.HP
	}
	return u
}

// LoadUnitsFromUser はユーザテーブル（usr_）から UI 用ユニット配列を生成します。
func LoadUnitsFromUser(path string) ([]Unit, error) {
	ut, err := user.LoadFromJSON(path)
	if err != nil {
		return nil, err
	}
	// 安全な順序で反復（ID昇順ではなく定義順）
	units := make([]Unit, 0, len(ut.ByID()))
	for _, c := range ut.Slice() {
		units = append(units, unitFromUser(c))
	}
	return units, nil
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
		if it.Max > 0 {
			uses = fmt.Sprintf("%d/%d", it.Uses, it.Max)
		}
		text.Draw(dst, uses, faceSmall, int(px)+300, lineY, colAccent)
	}
}

// 一覧レイアウト定数
const (
	listMargin      = 24
	listItemH       = 100
	listItemGap     = 12
	listPortraitSz  = 80
	listTitleOffset = 44
)

// ListItemRect は一覧画面の i 番目の行の矩形を返します。
func ListItemRect(sw, _, i int) (x, y, w, h int) {
	panelX, panelY := listMargin, listMargin
	panelW := sw - listMargin*2
	startY := panelY + listTitleOffset + 32
	y = startY + i*(listItemH+listItemGap)
	return panelX + 16, y, panelW - 32, listItemH
}

// DrawCharacterList はユニット一覧を描画します。
func DrawCharacterList(dst *ebiten.Image, units []Unit, hover int) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	// パネル
	drawPanel(dst, float32(listMargin), float32(listMargin), float32(sw-2*listMargin), float32(sh-2*listMargin))
	text.Draw(dst, "ユニット一覧", faceTitle, listMargin+20, listMargin+listTitleOffset, colAccent)
	for i, u := range units {
		x, y, w, h := ListItemRect(sw, sh, i)
		// カード背景
		bg := color.RGBA{30, 45, 78, 255}
		if i == hover {
			bg = color.RGBA{40, 60, 100, 255}
		}
		vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
		vector.DrawFilledRect(dst, float32(x-2), float32(y-2), float32(w+4), float32(h+4), colBorder, false)

		// ポートレート
		px := float32(x + 12)
		py := float32(y + (h-listPortraitSz)/2)
		drawFramedRect(dst, px-2, py-2, listPortraitSz+4, listPortraitSz+4)
		if u.Portrait != nil {
			drawPortrait(dst, u.Portrait, px, py, listPortraitSz, listPortraitSz)
		} else {
			drawPortraitPlaceholder(dst, px, py, listPortraitSz, listPortraitSz)
		}

		// テキスト
		tx := x + 12 + listPortraitSz + 20
		ty := y + 36
		text.Draw(dst, u.Name, faceMain, tx, ty, colText)
		text.Draw(dst, fmt.Sprintf("%s  Lv %d", u.Class, u.Level), faceSmall, tx, ty+26, colAccent)
	}
}

// BackButtonRect はステータス画面に表示する戻るボタンの矩形を返します。
func BackButtonRect(sw, _ int) (x, y, w, h int) {
	panelX, panelY := listMargin, listMargin
	panelW := sw - listMargin*2
	x = panelX + panelW - 180
	y = panelY + 24
	w = 160
	h = 48
	return
}

// DrawBackButton は戻るボタンを描画します。
func DrawBackButton(dst *ebiten.Image, hovered bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := BackButtonRect(sw, sh)
	bg := color.RGBA{50, 70, 110, 255}
	if hovered {
		bg = color.RGBA{70, 100, 150, 255}
	}
	drawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
	text.Draw(dst, "＜ 一覧へ", faceMain, x+20, y+32, colText)
}

// ToBattleButtonRect はステータス画面の「戦闘へ」ボタンの矩形を返します。
func ToBattleButtonRect(sw, sh int) (x, y, w, h int) {
	// レベルアップボタンの左隣に配置
	rx, ry, _, rh := LevelUpButtonRect(sw, sh)
	w, h = 220, rh
	x = rx - 20 - w
	y = ry
	return
}

// DrawToBattleButton は「戦闘へ」ボタンを描画します。
func DrawToBattleButton(dst *ebiten.Image, hovered, enabled bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := ToBattleButtonRect(sw, sh)
	base := color.RGBA{90, 90, 130, 255}
	if !enabled {
		base = color.RGBA{70, 70, 70, 255}
	}
	if hovered && enabled {
		base = color.RGBA{110, 110, 170, 255}
	}
	drawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
	text.Draw(dst, "戦闘へ", faceMain, x+70, y+36, colText)
}

// LevelUpButtonRect はレベルアップボタンの矩形を返します。
func LevelUpButtonRect(sw, sh int) (x, y, w, h int) {
	// 右下に配置
	w, h = 220, 56
	x = sw - listMargin - w
	y = sh - listMargin - h
	return
}

// DrawLevelUpButton はレベルアップボタンを描画します。
func DrawLevelUpButton(dst *ebiten.Image, hovered bool, enabled bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := LevelUpButtonRect(sw, sh)
	base := color.RGBA{80, 130, 60, 255}
	if !enabled {
		base = color.RGBA{70, 70, 70, 255}
	}
	if hovered && enabled {
		base = color.RGBA{100, 170, 80, 255}
	}
	drawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
	label := "レベルアップ"
	if !enabled {
		label = "最大レベル"
	}
	text.Draw(dst, label, faceMain, x+24, y+36, colText)
}

// LevelUpGains はレベルアップで上昇した値を表します。
type LevelUpGains struct {
	Inc    Stats
	HPGain int
}

// RollLevelUp は成長率に基づき上昇値を抽選します（各+1固定）。
func RollLevelUp(u Unit, rnd func() float64) LevelUpGains {
	g := LevelUpGains{}
	prob := func(p int) bool { return p > 0 && rnd()*100 < float64(p) }
	if prob(u.Growth.HP) {
		g.HPGain++
	}
	if prob(u.Growth.Str) {
		g.Inc.Str++
	}
	if prob(u.Growth.Mag) {
		g.Inc.Mag++
	}
	if prob(u.Growth.Skl) {
		g.Inc.Skl++
	}
	if prob(u.Growth.Spd) {
		g.Inc.Spd++
	}
	if prob(u.Growth.Lck) {
		g.Inc.Lck++
	}
	if prob(u.Growth.Def) {
		g.Inc.Def++
	}
	if prob(u.Growth.Res) {
		g.Inc.Res++
	}
	if prob(u.Growth.Mov) {
		g.Inc.Mov++
	}
	return g
}

// ApplyGains は抽選結果をユニットへ反映します（Level+1, HP系は上限に合わせる）。
func ApplyGains(u *Unit, gains LevelUpGains, levelCap int) {
	if u.Level < levelCap {
		u.Level++
	}
	u.HPMax += gains.HPGain
	u.HP += gains.HPGain
	if u.HP > u.HPMax {
		u.HP = u.HPMax
	}
	u.Stats.Str += gains.Inc.Str
	u.Stats.Mag += gains.Inc.Mag
	u.Stats.Skl += gains.Inc.Skl
	u.Stats.Spd += gains.Inc.Spd
	u.Stats.Lck += gains.Inc.Lck
	u.Stats.Def += gains.Inc.Def
	u.Stats.Res += gains.Inc.Res
	u.Stats.Mov += gains.Inc.Mov
	// クラス上限でクランプ
	if caps, err := model.LoadClassCapsJSON("db/master/mst_class_caps.json"); err == nil {
		if c, ok := caps.Find(u.Class); ok {
			clamp := func(v, m int) int {
				if m > 0 && v > m {
					return m
				}
				if v < 0 {
					return 0
				}
				return v
			}
			u.HPMax = clamp(u.HPMax, c.HPMax)
			u.HP = clamp(u.HP, u.HPMax)
			u.Stats.Str = clamp(u.Stats.Str, c.Str)
			u.Stats.Mag = clamp(u.Stats.Mag, c.Mag)
			u.Stats.Skl = clamp(u.Stats.Skl, c.Skl)
			u.Stats.Spd = clamp(u.Stats.Spd, c.Spd)
			u.Stats.Lck = clamp(u.Stats.Lck, c.Lck)
			u.Stats.Def = clamp(u.Stats.Def, c.Def)
			u.Stats.Res = clamp(u.Stats.Res, c.Res)
			u.Stats.Mov = clamp(u.Stats.Mov, c.Mov)
		}
	}
}

// --- 戦闘画面（簡易） ---

// BattleStartButtonRect は戦闘開始ボタンの矩形を返します。
func BattleStartButtonRect(sw, sh int) (x, y, w, h int) {
	w, h = 240, 60
	x = (sw - w) / 2
	y = sh - listMargin - h
	return
}

// DrawBattle は簡易な戦闘プレビュー画面を描画します。
// attacker/defender はUIのUnit。武器は先頭装備を使用します。
func DrawBattle(dst *ebiten.Image, attacker, defender Unit) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	drawPanel(dst, listMargin, listMargin, float32(sw-2*listMargin), float32(sh-2*listMargin))
	// 左右にユニット
	leftX := listMargin + 40
	rightX := sw - listMargin - 560
	topY := listMargin + 80

	drawBattleSide(dst, attacker, leftX, topY)
	drawBattleSide(dst, defender, rightX, topY)

	// 中央見出し
	text.Draw(dst, "戦闘プレビュー", faceTitle, sw/2-120, listMargin+56, colAccent)

	// ボタン
	bx, by, bw, bh := BattleStartButtonRect(sw, sh)
	drawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
	vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), color.RGBA{110, 90, 40, 255}, false)
	text.Draw(dst, "戦闘開始", faceMain, bx+70, by+38, colText)
}

func drawBattleSide(dst *ebiten.Image, u Unit, x, y int) {
	// 顔 + 基本
	drawFramedRect(dst, float32(x), float32(y), 320, 320)
	if u.Portrait != nil {
		drawPortrait(dst, u.Portrait, float32(x), float32(y), 320, 320)
	}
	text.Draw(dst, u.Name, faceTitle, x, y-16, colText)
	text.Draw(dst, fmt.Sprintf("%s  Lv %d", u.Class, u.Level), faceMain, x, y+350, colAccent)
	text.Draw(dst, fmt.Sprintf("HP %d/%d", u.HP, u.HPMax), faceMain, x, y+384, colText)
	drawHPBar(dst, x, y+390, 320, 14, u.HP, u.HPMax)
	// 武器（先頭）
	wepName := "-"
	if len(u.Equip) > 0 {
		wepName = u.Equip[0].Name
	}
	text.Draw(dst, fmt.Sprintf("武器: %s", wepName), faceMain, x, y+420, colText)
}

// DrawLevelUpPopup はレベルアップ結果をポップアップ表示します。
func DrawLevelUpPopup(dst *ebiten.Image, u Unit, gains LevelUpGains) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	// 半透明オーバーレイ
	overlay := color.RGBA{0, 0, 0, 160}
	vector.DrawFilledRect(dst, float32(0), float32(0), float32(sw), float32(sh), overlay, false)
	// パネル
	pw, ph := 520, 480
	px := (sw - pw) / 2
	py := (sh - ph) / 2
	drawPanel(dst, float32(px), float32(py), float32(pw), float32(ph))
	text.Draw(dst, "レベルアップ!", faceTitle, px+24, py+56, colAccent)
	text.Draw(dst, fmt.Sprintf("Lv %d", u.Level), faceMain, px+24, py+96, colText)

	// 上昇した項目のみを表示
	y := py + 140
	line := 34
	drawInc := func(label string, v int) {
		if v > 0 {
			text.Draw(dst, fmt.Sprintf("%s +%d", label, v), faceMain, px+40, y, colAccent)
			y += line
		}
	}
	if gains.HPGain > 0 {
		drawInc("HP", gains.HPGain)
	}
	drawInc("力", gains.Inc.Str)
	drawInc("魔力", gains.Inc.Mag)
	drawInc("技", gains.Inc.Skl)
	drawInc("速さ", gains.Inc.Spd)
	drawInc("幸運", gains.Inc.Lck)
	drawInc("守備", gains.Inc.Def)
	drawInc("魔防", gains.Inc.Res)
	drawInc("移動", gains.Inc.Mov)

	text.Draw(dst, "クリックで閉じる", faceSmall, px+pw-180, py+ph-24, colText)
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
func drawPortraitPlaceholder(dst *ebiten.Image, x, y, _, h float32) {
	text.Draw(dst, "画像なし", faceSmall, int(x+10), int(y+h/2), colAccent)
}

// drawPortrait はポートレート画像を枠内に等比縮小して描画します。
// 線形補間により縮小時のジャギーを低減します。
func drawPortrait(dst *ebiten.Image, img *ebiten.Image, x, y, w, h float32) {
	b := img.Bounds()
	iw, ih := b.Dx(), b.Dy()
	if iw == 0 || ih == 0 {
		return
	}
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
func drawHPBar(dst *ebiten.Image, x, y, w, h int, hp, maxHP int) {
	if maxHP <= 0 {
		maxHP = 1
	}
	// 背景
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), color.RGBA{50, 50, 50, 255}, false)
	// 値
	ratio := float32(hp) / float32(maxHP)
	bw := float32(w) * ratio
	col := color.RGBA{80, 220, 100, 255}
	if ratio < 0.33 {
		col = color.RGBA{220, 80, 80, 255}
	} else if ratio < 0.66 {
		col = color.RGBA{240, 200, 80, 255}
	}
	vector.DrawFilledRect(dst, float32(x), float32(y), bw, float32(h), col, false)
}

// drawStatLineWithGrowth は能力値と成長率(%)を並べて描画します。
func drawStatLineWithGrowth(dst *ebiten.Image, face font.Face, x, y int, label string, v int, g int) {
	text.Draw(dst, label, face, x, y, colText)
	text.Draw(dst, fmt.Sprintf("%2d", v), face, x+64, y, colAccent)
	text.Draw(dst, fmt.Sprintf("%d%%", g), faceSmall, x+120, y, colAccent)
}

// drawRankLine は武器/魔法ランクのラベルと値を描画します。
func drawRankLine(dst *ebiten.Image, face font.Face, x, y int, label, rank string) {
	if rank == "" {
		rank = "-"
	}
	text.Draw(dst, label, face, x, y, colText)
	text.Draw(dst, rank, face, x+120, y, colAccent)
}
