// Package characterlist は、ユニット一覧（リスト）画面のシーンを提供します。
//
// 主な責務:
// - 画面内のユニット行ホバー状態の管理
// - クリックでの選択 → ステータス画面への遷移
// - ショートカット操作（武器/アイテム一覧）から在庫画面へ遷移
// - 模擬戦ボタン押下後の「攻撃側/防御側」2段階選択と Sim シーンへの遷移
package characterlist

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "ui_sample/internal/game"
    scenes "ui_sample/internal/game/scenes"
    inventory "ui_sample/internal/game/scenes/inventory"
    sim "ui_sample/internal/game/scenes/sim"
    status "ui_sample/internal/game/scenes/status"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    uidraw "ui_sample/internal/game/ui/draw"
    uinput "ui_sample/internal/game/ui/input"
    uilayout "ui_sample/internal/game/ui/layout"
    "ui_sample/pkg/game/geom"
)

// List はユニット一覧画面の Scene 実装です。
//
// 入力処理の優先順位は次の通りです:
// 1) 模擬戦選択モード中は選択ポップアップの処理を最優先
// 2) 通常時はユニット行クリック→ステータス画面へ遷移
// 3) ショートカット（W:武器一覧 / I:アイテム一覧）で在庫画面へ遷移
type List struct {
	// E: シーン間共有の環境（選択ユニット、ユーザテーブル、ユースケース等）
	E *scenes.Env
	// hoverIndex: マウスが乗っている行のインデックス（-1 は非ホバー）
	hoverIndex int
	// simBtnHovered: 模擬戦ボタンのホバー状態
	simBtnHovered bool
	// simSelecting: 模擬戦の選択モード中か（true: ポップアップ表示中）
	simSelecting bool
	// simSelectStep: 0=攻撃側を選択, 1=防御側を選択（2段階選択の状態）
	simSelectStep int
	// chooseHover: 選択ポップアップ内のホバー行（-1 は非ホバー）
	chooseHover int
	// tmpAtk: 先に選ばれた攻撃側ユニット
	tmpAtk uicore.Unit
	// sw, sh: 直近フレームのスクリーン幅/高さキャッシュ（矩形計算に使用）
	sw, sh int
	// next: このフレームで遷移すべき次シーン（最後に返す）
	next game.Scene
}

// NewList はユニット一覧シーンを生成します。
// 生成時点ではホバー無し（hoverIndex=-1）、模擬戦選択モードは無効です。
// 引数: e — シーン間で共有する環境（ユニット一覧やユーザ情報等）。
// 戻り値: 新しい List インスタンス。
func NewList(e *scenes.Env) *List { return &List{E: e, hoverIndex: -1} }

// Update は状態更新の入口（司令塔）です。
// フロー: handleInput → advance → flush → s.next を返却。
// 引数: ctx — 画面サイズや入力状態を含むフレームコンテキスト。
// 戻り値: 次に積むシーン（なければ nil）、エラー（通常 nil）。
func (s *List) Update(ctx *game.Ctx) (game.Scene, error) {
	s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
	intents := s.scHandleInput(ctx)
	s.scAdvance(intents)
	s.scFlush(ctx)
	nxt := s.next
	s.next = nil
	return nxt, nil
}

// Draw は一覧と UI 部品の描画を行います。
// 模擬戦の選択モード中は半透明のポップアップを重ねて候補を提示します。
// 引数: dst — 描画先の ebiten.Image。
// 戻り値: なし。
func (s *List) Draw(dst *ebiten.Image) {
	s.drawCharacterList(dst)
	uiwidgets.DrawSimBattleButton(dst, s.simBtnHovered, len(s.E.Units) > 1)
	// 画面左上にショートカットのヒントを表示
	ebitenutil.DebugPrintAt(dst, "W: 武器一覧 / I: アイテム一覧", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
	if s.simSelecting {
		// 選択段階に応じてタイトルを切り替え
		title := "模擬戦: 攻撃側を選択"
		if s.simSelectStep == 1 {
			title = "模擬戦: 防御側を選択"
		}
		uidraw.DrawChooseUnitPopup(dst, title, s.E.Units, s.chooseHover)
	}
}

// パッケージ内コンパイル保証: 必須の cl* メソッドを実装しているか
type scContract interface {
	scHandleInput(ctx *game.Ctx) []scenes.Intent
	scAdvance(intents []scenes.Intent)
	scFlush(ctx *game.Ctx)
}

var _ scContract = (*List)(nil)

// IntentKind は入力の意味（意図）を表します。
type IntentKind int

const (
	// 先頭は 0 から開始
	_ IntentKind = iota
	// intentOpenStatus: 一覧のホバー行を選択してステータス画面を開く。
	intentOpenStatus
	// intentOpenWeapons: 武器在庫一覧を開く（SelectingEquip=false）。
	intentOpenWeapons
	// intentOpenItems: アイテム在庫一覧を開く（SelectingEquip=false）。
	intentOpenItems
	// intentOpenSimSelect: 模擬戦の選択モードを開始（攻撃側選択へ）。
	intentOpenSimSelect
	// intentCancelSim: 模擬戦の選択モードをキャンセル。
	intentCancelSim
	// intentSelectAtk: 模擬戦の攻撃側を Index で確定。
	intentSelectAtk
	// intentSelectDef: 模擬戦の防御側を Index で確定し、Sim へ遷移。
	intentSelectDef
)

// Intent は入力の意味を圧縮したものです。
type Intent struct {
	Kind  IntentKind
	Index int // ユニットインデックス等
}

// IsSceneIntent は scenes.Intent のマーカーです。
// 引数: なし。
// 戻り値: なし。
func (Intent) IsSceneIntent() {}

// scHandleInput は“入力→意図(Intent)”へ変換し、描画用キャッシュを更新します。
// 引数: ctx — 入力状態を参照するフレームコンテキスト。
// 戻り値: フレーム内に発生した意図の列（処理優先度順）。
func (s *List) scHandleInput(ctx *game.Ctx) []scenes.Intent {
	intents := make([]scenes.Intent, 0, 3)
	// マウス座標はフレーム先頭で固定し、結果をメンバへ保存
    mx, my := ctx.CursorX, ctx.CursorY
	// ホバー更新（一覧）
	s.hoverIndex = -1
	for i := range s.E.Units {
		x, y, w, h := uilayout.ListItemRect(s.sw, s.sh, i)
		if geom.RectContains(mx, my, x, y, w, h) {
			s.hoverIndex = i
		}
	}
	// 模擬戦ボタンのホバー
	bx, by, bw, bh := uiwidgets.SimBattleButtonRect(s.sw, s.sh)
	s.simBtnHovered = geom.RectContains(mx, my, bx, by, bw, bh)

	// 選択ポップアップ中のホバー更新を先に計算（このフレームの確定に反映）
	if s.simSelecting {
		s.chooseHover = -1
		for i := range s.E.Units {
			x, y, w, h := uilayout.ChooseUnitItemRect(s.sw, s.sh, i, len(s.E.Units))
			if geom.RectContains(mx, my, x, y, w, h) {
				s.chooseHover = i
			}
		}
	} else {
		s.chooseHover = -1
	}

	// キー/アクション意図（chooseHover 計算後に生成）
	if ctx != nil && ctx.Input != nil {
		if ctx.Input.Press(uinput.OpenWeapons) {
			intents = append(intents, Intent{Kind: intentOpenWeapons})
		}
		if ctx.Input.Press(uinput.OpenItems) {
			intents = append(intents, Intent{Kind: intentOpenItems})
		}
		if s.simSelecting {
			if ctx.Input.Press(uinput.Cancel) {
				intents = append(intents, Intent{Kind: intentCancelSim})
			}
			if s.chooseHover >= 0 && ctx.Input.Press(uinput.Confirm) {
				if s.simSelectStep == 0 {
					intents = append(intents, Intent{Kind: intentSelectAtk, Index: s.chooseHover})
				} else {
					intents = append(intents, Intent{Kind: intentSelectDef, Index: s.chooseHover})
				}
			}
		} else {
			if s.hoverIndex >= 0 && ctx.Input.Press(uinput.Confirm) {
				intents = append(intents, Intent{Kind: intentOpenStatus, Index: s.hoverIndex})
			}
			if s.simBtnHovered && len(s.E.Units) > 1 && ctx.Input.Press(uinput.Confirm) {
				intents = append(intents, Intent{Kind: intentOpenSimSelect})
			}
		}
	}
	return intents
}

// scAdvance は状態機械を進め、副作用（遷移生成など）を集約します。
// 引数: intents — このフレームに発生した意図列。
// 戻り値: なし。
func (s *List) scAdvance(intents []scenes.Intent) {
	for _, any := range intents {
		it, ok := any.(Intent)
		if !ok {
			continue
		}
		if s.next != nil {
			break
		} // 最初の決定を優先し上書きを避ける
		switch it.Kind {
		case intentOpenWeapons:
			s.E.InvTab, s.E.HoverInv = 0, -1
			s.E.SelectingEquip, s.E.SelectingIsWeapon = false, true
			s.next = inventory.NewInventory(s.E)
		case intentOpenItems:
			s.E.InvTab, s.E.HoverInv = 1, -1
			s.E.SelectingEquip, s.E.SelectingIsWeapon = false, false
			s.next = inventory.NewInventory(s.E)
		case intentOpenStatus:
			if it.Index >= 0 && it.Index < len(s.E.Units) {
				s.E.SelIndex = it.Index
				s.next = status.NewStatus(s.E)
			}
		case intentOpenSimSelect:
			s.simSelecting = true
			s.simSelectStep = 0
			s.chooseHover = -1
		case intentCancelSim:
			s.simSelecting = false
			s.simSelectStep = 0
			s.chooseHover = -1
		case intentSelectAtk:
			if it.Index >= 0 && it.Index < len(s.E.Units) {
				s.tmpAtk = s.E.Units[it.Index]
				s.simSelectStep = 1
			}
		case intentSelectDef:
			if it.Index >= 0 && it.Index < len(s.E.Units) {
				def := s.E.Units[it.Index]
				s.next = sim.NewSim(s.E, s.tmpAtk, def)
				// モードは呼び出し先で閉じる前提（戻ってきたらリセット）
				s.simSelecting = false
				s.simSelectStep = 0
				s.chooseHover = -1
			}
		}
	}
}

// scFlush はフレーム末尾の副作用処理（Audio/GC 等）をまとめます。
// 引数: ctx — フレームコンテキスト（未使用時は _ で可）。
// 戻り値: なし。
func (s *List) scFlush(_ *game.Ctx) {
	// 今は特になし（将来の Audio 発火や遅延解放に備えたフック）
}

// drawCharacterList はこのシーンのユニット一覧を描画します（非公開メソッド）。
// 引数: dst — 描画先の ebiten.Image。
// 戻り値: なし。
func (s *List) drawCharacterList(dst *ebiten.Image) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	lm := uicore.ListMarginPx()
	uicore.DrawPanel(dst, float32(lm), float32(lm), float32(sw-2*lm), float32(sh-2*lm))
	uicore.TextDraw(dst, "ユニット一覧", uicore.FaceTitle, lm+uicore.ListTitleXOffsetPx(), lm+uicore.ListTitleOffsetPx(), uicore.ColAccent)
	for i, u := range s.E.Units {
		x, y, w, h := uilayout.ListItemRect(sw, sh, i)
		bg := color.RGBA{30, 45, 78, 255}
		if i == s.hoverIndex {
			bg = color.RGBA{40, 60, 100, 255}
		}
		vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
		vector.DrawFilledRect(dst, float32(x-uicore.S(2)), float32(y-uicore.S(2)), float32(w+uicore.S(4)), float32(h+uicore.S(4)), uicore.ColBorder, false)
		ps := uicore.ListPortraitSzPx()
		px := float32(x + uicore.S(12))
		py := float32(y + (h-ps)/2)
		uicore.DrawFramedRect(dst, px-float32(uicore.S(2)), py-float32(uicore.S(2)), float32(ps+uicore.S(4)), float32(ps+uicore.S(4)))
		if u.Portrait != nil {
			uicore.DrawPortrait(dst, u.Portrait, px, py, float32(ps), float32(ps))
		} else {
			uicore.DrawPortraitPlaceholder(dst, px, py, float32(ps), float32(ps))
		}
		tx := x + uicore.S(12) + ps + uicore.S(20)
		ty := y + uicore.S(36)
		uicore.TextDraw(dst, u.Name, uicore.FaceMain, tx, ty, uicore.ColText)
		uicore.TextDraw(dst, u.Class+"  Lv "+uicore.Itoa(u.Level), uicore.FaceSmall, tx, ty+uicore.S(26), uicore.ColAccent)
	}
}
