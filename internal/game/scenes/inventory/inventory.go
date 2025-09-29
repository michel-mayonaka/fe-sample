package inventory

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "ui_sample/internal/game"
    gamesvc "ui_sample/internal/game/service"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    scenes "ui_sample/internal/game/scenes"
    "ui_sample/pkg/game/geom"
)

// Inventory は在庫画面の Scene 実装です。
//
// 主な責務:
// - 武器/アイテムのタブ切替と一覧表示
// - スロット選択時の装備割り当て（オーナー移動を含む）
// - 戻る操作でのシーン終了
//
// 更新フローは character_list と同一で、Update → scHandleInput → scAdvance → scFlush の順で処理します。
type Inventory struct{
    E *scenes.Env
    pop bool
    sw, sh int
    // ホバー状態（描画に使用）
    tabHover int // -1=なし, 0=武器, 1=アイテム
    backHovered bool
    wv *WeaponView
    iv *ItemView
}
// NewInventory は在庫画面の Scene を生成します。
func NewInventory(e *scenes.Env) *Inventory {
    s := &Inventory{E:e, tabHover: -1}
    s.wv = NewWeaponView(e, s)
    s.iv = NewItemView(e, s)
    return s
}
// ShouldPop は本シーンが終了要求（pop）状態かを返します。
func (s *Inventory) ShouldPop() bool { return s.pop }

// IntentKind は在庫画面における入力意図の種別です。
type IntentKind int

const (
    intentNone IntentKind = iota
    intentBack
    intentTabWeapons
    intentTabItems
    intentChooseRow // Index に行番号
)

// Intent は入力を意味表現に変換したものです。
type Intent struct{
    Kind IntentKind
    Index int
}
// IsSceneIntent は scenes.Intent 実装のマーカーです。
func (Intent) IsSceneIntent() {}

// scContract はパッケージ内コンパイル保証のためのインターフェースです。
// Inventory が必要な sc* メソッドを実装していることを確認します。
type scContract interface {
    scHandleInput(ctx *game.Ctx) []scenes.Intent
    scAdvance(intents []scenes.Intent)
    scFlush(ctx *game.Ctx)
}
var _ scContract = (*Inventory)(nil)

// Update は状態更新の入口です。
// フロー: scHandleInput → scAdvance → scFlush。次シーンはこの関数では返しません（呼び出し元で積み替え）。
func (s *Inventory) Update(ctx *game.Ctx) (game.Scene, error) {
    s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
    // サブビュー更新（タブに応じて）
    if s.E.InvTab == 0 { _, _ = s.wv.Update(ctx) } else { _, _ = s.iv.Update(ctx) }
    intents := s.scHandleInput(ctx)
    s.scAdvance(intents)
    s.scFlush(ctx)
    return nil, nil
}

// Draw はタブとリスト、戻るボタン等の UI を描画します。
func (s *Inventory) Draw(dst *ebiten.Image) {
    // 本体（タブに応じて）
    if s.E.InvTab == 0 { s.wv.Draw(dst) } else { s.iv.Draw(dst) }
    // タブ描画
    tabW, tabH := uicore.S(160), uicore.S(44)
    lm := uicore.ListMarginPx()
    tx := lm + uicore.S(20)
    ty := lm + uicore.S(12)
    // 武器タブ
    uicore.DrawFramedRect(dst, float32(tx), float32(ty), float32(tabW), float32(tabH))
    baseW := color.RGBA{40, 60, 110, 255}
    if s.E.InvTab == 0 { baseW = color.RGBA{70, 100, 160, 255} }
    vector.DrawFilledRect(dst, float32(tx), float32(ty), float32(tabW), float32(tabH), baseW, false)
    uicore.TextDraw(dst, "武器", uicore.FaceMain, tx+uicore.S(56), ty+uicore.S(30), uicore.ColText)
    // アイテムタブ
    tx2 := tx + tabW + uicore.S(10)
    uicore.DrawFramedRect(dst, float32(tx2), float32(ty), float32(tabW), float32(tabH))
    baseI := color.RGBA{40, 60, 110, 255}
    if s.E.InvTab == 1 { baseI = color.RGBA{70, 100, 160, 255} }
    vector.DrawFilledRect(dst, float32(tx2), float32(ty), float32(tabW), float32(tabH), baseI, false)
    uicore.TextDraw(dst, "アイテム", uicore.FaceMain, tx2+uicore.S(34), ty+uicore.S(30), uicore.ColText)
    // 戻る
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    mx, my := ebiten.CursorPosition()
    s.backHovered = geom.RectContains(mx, my, bx, by, bw, bh)
    uiwidgets.DrawBackButton(dst, s.backHovered)
    if s.E.SelectingEquip { ebitenutil.DebugPrintAt(dst, "クリックでスロットに装備", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10)) }
}

// 装備確定の実処理は各サブビュー（WeaponView/ItemView）へ移動しました。

// refreshUnitByID はユーザテーブルの変更を UI 表示用ユニット配列へ反映します。
func (s *Inventory) refreshUnitByID(id string) {
    if s.E == nil || s.E.UserTable == nil { return }
    c, ok := s.E.UserTable.Find(id); if !ok { return }
    u := uicore.UnitFromUser(c)
    for i := range s.E.Units {
        if s.E.Units[i].ID == id {
            s.E.Units[i] = u
        }
    }
}

// --- 内部: scHandleInput → scAdvance → scFlush --------------------------------------

// scHandleInput は“入力→意図(Intent)”へ変換し、ホバー状態を更新します。
func (s *Inventory) scHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 3)
    mx, my := ebiten.CursorPosition()
    // ホバー更新
    s.tabHover = -1
    tabW, tabH := uicore.S(160), uicore.S(44)
    lm := uicore.ListMarginPx()
    tx := lm + uicore.S(20)
    ty := lm + uicore.S(12)
    if geom.RectContains(mx, my, tx, ty, tabW, tabH) { s.tabHover = 0 }
    tx2 := tx + tabW + uicore.S(10)
    if geom.RectContains(mx, my, tx2, ty, tabW, tabH) { s.tabHover = 1 }

    // ボタンホバー
    bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
    s.backHovered = geom.RectContains(mx, my, bx, by, bw, bh)

    if ctx != nil && ctx.Input != nil {
        if ctx.Input.Press(gamesvc.Cancel) { intents = append(intents, Intent{Kind: intentBack}) }
        if ctx.Input.Press(gamesvc.Confirm) {
            switch s.tabHover {
            case 0: intents = append(intents, Intent{Kind: intentTabWeapons})
            case 1: intents = append(intents, Intent{Kind: intentTabItems})
            }
            if s.backHovered { intents = append(intents, Intent{Kind: intentBack}) }
        }
    }
    return intents
}

// scAdvance は意図に基づき、タブ切替・装備確定・戻る等の状態変更を行います。
func (s *Inventory) scAdvance(intents []scenes.Intent) {
    for _, any := range intents {
        it, ok := any.(Intent); if !ok { continue }
        switch it.Kind {
        case intentBack:
            s.pop = true
        case intentTabWeapons:
            if s.E.Inv != nil && s.E.Inv.Inventory() != nil { s.E.InvTab, s.E.HoverInv = 0, -1 }
        case intentTabItems:
            if s.E.Inv != nil && s.E.Inv.Inventory() != nil { s.E.InvTab, s.E.HoverInv = 1, -1 }
        }
    }
}

// scFlush はフレーム末尾の副作用処理用フックです（現状なし）。
func (s *Inventory) scFlush(_ *game.Ctx) { /* 今はなし */ }
