package status

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "ui_sample/internal/game"
    gdata "ui_sample/internal/game/data"
    scenes "ui_sample/internal/game/scenes"
    inventory "ui_sample/internal/game/scenes/inventory"
    lvl "ui_sample/internal/game/service/levelup"
    uicore "ui_sample/internal/game/service/ui"
    uiwidgets "ui_sample/internal/game/service/ui/widgets"
    uidraw "ui_sample/internal/game/ui/draw"
    uinput "ui_sample/internal/game/ui/input"
    uilayout "ui_sample/internal/game/ui/layout"
    "ui_sample/pkg/game/geom"
)

// Status はステータス画面の Scene 実装です。
//
// 主な責務:
// - 選択ユニットの基礎情報・装備・能力を表示
// - レベルアップのポップアップ表示と確定
// - 装備スロット選択から在庫画面への遷移
//
// 更新フローは character_list と同一で、Update → scHandleInput → scAdvance → scFlush の順で処理します。
type Status struct {
	E           *scenes.Env
	pop         bool
	sw, sh      int
	next        game.Scene
	backHovered bool
	lvHovered   bool
}

// NewStatus はステータス画面の Scene を生成します。
func NewStatus(e *scenes.Env) *Status { return &Status{E: e} }

// ShouldPop は本シーンが終了要求（pop）状態かを返します。
func (s *Status) ShouldPop() bool { return s.pop }

// IntentKind はステータス画面における入力意図の種別です。
type IntentKind int

const (
	intentNone IntentKind = iota
	intentBack
	intentClosePopup
	intentLevelUp
	intentOpenInvSlot // Index にスロット番号
	intentOpenInvWeapons
	intentOpenInvItems
)

// Intent は入力を意味表現に変換したものです。
type Intent struct {
	Kind  IntentKind
	Index int
}

// IsSceneIntent は scenes.Intent 実装のマーカーです。
func (Intent) IsSceneIntent() {}

// scContract はパッケージ内コンパイル保証のためのインターフェースです。
// Status が必要な sc* メソッドを実装していることを確認します。
type scContract interface {
	scHandleInput(ctx *game.Ctx) []scenes.Intent
	scAdvance([]scenes.Intent)
	scFlush(ctx *game.Ctx)
}

var _ scContract = (*Status)(nil)

// Update は状態更新の入口です。
// フロー: scHandleInput → scAdvance → scFlush。次シーンを返す場合は在庫画面などの遷移先を返却します。
func (s *Status) Update(ctx *game.Ctx) (game.Scene, error) {
	s.sw, s.sh = ctx.ScreenW, ctx.ScreenH
	intents := s.scHandleInput(ctx)
	s.scAdvance(intents)
	s.scFlush(ctx)
	nxt := s.next
	s.next = nil
	return nxt, nil
}

// Draw はステータス UI、戻る/レベルアップボタン、必要に応じてポップアップを描画します。
func (s *Status) Draw(dst *ebiten.Image) {
    // 本体（ステータス）
    unit := s.E.Selected()
    uidraw.DrawStatus(dst, unit)
    // 戻る/レベルアップは scHandleInput で更新済みのホバー状態を使って描画
    uiwidgets.DrawBackButton(dst, s.backHovered)
    uiwidgets.DrawLevelUpButton(dst, s.lvHovered, unit.Level < game.LevelCap && !s.E.PopupActive)
	if s.E.PopupActive {
		uidraw.DrawLevelUpPopup(dst, unit, s.E.PopupGains)
	}
	ebitenutil.DebugPrintAt(dst, "装備: E/数字/DELETE で操作", uicore.ListMarginPx()+uicore.S(20), uicore.ListMarginPx()+uicore.S(10))
}

// --- 内部: scHandleInput → scAdvance → scFlush --------------------------------------

// scHandleInput は“入力→意図(Intent)”へ変換し、ホバー状態やポップアップの直後状態を更新します。
func (s *Status) scHandleInput(ctx *game.Ctx) []scenes.Intent {
    intents := make([]scenes.Intent, 0, 4)
    mx, my := ctx.CursorX, ctx.CursorY
	wasPopup := s.E != nil && s.E.PopupActive
	// 戻る/レベルアップ ホバー
	bx, by, bw, bh := uiwidgets.BackButtonRect(s.sw, s.sh)
	s.backHovered = geom.RectContains(mx, my, bx, by, bw, bh)
	lvx, lvy, lvw, lvh := uiwidgets.LevelUpButtonRect(s.sw, s.sh)
	s.lvHovered = geom.RectContains(mx, my, lvx, lvy, lvw, lvh)

	if ctx != nil && ctx.Input != nil {
		if ctx.Input.Press(uinput.Cancel) {
			intents = append(intents, Intent{Kind: intentBack})
		}

		// ポップアップ中の操作
		if s.E.PopupActive {
			if s.E.PopupJustOpened {
				s.E.PopupJustOpened = false
			} else if ctx.Input.Press(uinput.Confirm) {
				intents = append(intents, Intent{Kind: intentClosePopup})
			}
			// ポップアップを閉じる確定と同フレームでは以降を処理しない
			if wasPopup {
				return intents
			}
		}

		// レベルアップ
		unit := s.E.Selected()
		if unit.Level < game.LevelCap && !s.E.PopupActive && s.lvHovered && ctx.Input.Press(uinput.Confirm) {
			intents = append(intents, Intent{Kind: intentLevelUp})
		}

		// スロット選択→在庫
		for i := 0; i < 5; i++ {
			x, y, w, h := uilayout.EquipSlotRect(s.sw, s.sh, i)
			if geom.RectContains(mx, my, x, y, w, h) && ctx.Input.Press(uinput.Confirm) {
				intents = append(intents, Intent{Kind: intentOpenInvSlot, Index: i})
				break
			}
		}
		// ショートカット（E/I）
		if ctx.Input.Press(uinput.EquipToggle) {
			intents = append(intents, Intent{Kind: intentOpenInvWeapons})
		}
		if ctx.Input.Press(uinput.OpenItems) {
			intents = append(intents, Intent{Kind: intentOpenInvItems})
		}
	}
	return intents
}

// scAdvance は意図に基づき、戻る・ポップアップ閉鎖・レベルアップ・在庫遷移などの状態変更を行います。
func (s *Status) scAdvance(intents []scenes.Intent) {
	for _, any := range intents {
		it, ok := any.(Intent)
		if !ok {
			continue
		}
		switch it.Kind {
		case intentBack:
			s.pop = true
		case intentClosePopup:
			s.E.PopupActive = false
		case intentLevelUp:
			unit := s.E.Selected()
			gains := lvl.Roll(unit, s.E.RNG.Float64)
			lvl.Apply(&unit, gains, game.LevelCap)
			s.E.SetSelected(unit)
			s.E.PopupGains, s.E.PopupActive, s.E.PopupJustOpened = gains, true, true
			// 永続化は Usecase(DataPort) に委譲し、UI から Repo へ直接書き込みしない
			if s.E.Data != nil {
				_ = s.E.Data.PersistUnit(unit)
			}
		case intentOpenInvSlot:
			i := it.Index
			s.E.CurrentSlot = i
			s.E.SelectingEquip = true
			s.E.HoverInv = -1
			// 既装備の種別からタブを初期選択（なければ武器）
			s.E.SelectingIsWeapon = true
			s.E.InvTab = 0
			if p := gdata.Provider(); p != nil {
				unit := s.E.Selected()
				hasW, hasI := p.EquipKindAt(unit.ID, i)
				if hasI {
					s.E.SelectingIsWeapon = false
					s.E.InvTab = 1
				}
				if hasW {
					s.E.SelectingIsWeapon = true
					s.E.InvTab = 0
				}
			}
			s.next = inventory.NewInventory(s.E)
		case intentOpenInvWeapons:
			s.E.InvTab, s.E.HoverInv = 0, -1
			s.E.SelectingEquip, s.E.SelectingIsWeapon = true, true
			s.next = inventory.NewInventory(s.E)
		case intentOpenInvItems:
			s.E.InvTab, s.E.HoverInv = 1, -1
			s.E.SelectingEquip, s.E.SelectingIsWeapon = true, false
			s.next = inventory.NewInventory(s.E)
		}
	}
}

// scFlush はフレーム末尾の副作用処理用フックです（現状なし）。
func (s *Status) scFlush(_ *game.Ctx) { /* 今はなし */ }
