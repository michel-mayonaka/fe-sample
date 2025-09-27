// Package uiwidgets はボタン等の汎用UIウィジェット描画を提供します。
package uiwidgets

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    uicore "ui_sample/internal/game/service/ui"
)

// BackButtonRect は画面右上の戻るボタンの矩形を返します。
func BackButtonRect(sw, _ int) (x, y, w, h int) {
    lm := uicore.ListMarginPx()
    panelX, panelY := lm, lm
    panelW := sw - lm*2
    x = panelX + panelW - uicore.S(180)
    y = panelY + uicore.S(24)
    w = uicore.S(160)
    h = uicore.S(48)
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
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), bg, false)
    uicore.TextDraw(dst, "＜ 一覧へ", uicore.FaceMain, x+uicore.S(20), y+uicore.S(32), uicore.ColText)
}

// LevelUpButtonRect はレベルアップボタンの矩形を返します。
func LevelUpButtonRect(sw, sh int) (x, y, w, h int) {
    lm := uicore.ListMarginPx()
    w, h = uicore.S(220), uicore.S(56)
    x = sw - lm - w
    y = sh - lm - h
    return
}

// DrawLevelUpButton はレベルアップボタンを描画します。
func DrawLevelUpButton(dst *ebiten.Image, hovered, enabled bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := LevelUpButtonRect(sw, sh)
	base := color.RGBA{80, 130, 60, 255}
	if !enabled {
		base = color.RGBA{70, 70, 70, 255}
	}
	if hovered && enabled {
		base = color.RGBA{100, 170, 80, 255}
	}
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
    label := "レベルアップ"
    if !enabled {
        label = "最大レベル"
    }
    uicore.TextDraw(dst, label, uicore.FaceMain, x+uicore.S(24), y+uicore.S(36), uicore.ColText)
}

// ToBattleButtonRect は「戦闘へ」ボタンの矩形を返します。
func ToBattleButtonRect(sw, sh int) (x, y, w, h int) {
    rx, ry, _, rh := LevelUpButtonRect(sw, sh)
    w, h = uicore.S(220), rh
    x = rx - uicore.S(20) - w
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
    uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
    vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
    uicore.TextDraw(dst, "戦闘へ", uicore.FaceMain, x+uicore.S(70), y+uicore.S(36), uicore.ColText)
}

// SimBattleButtonRect は一覧画面の右上に表示する「模擬戦」ボタンの矩形です。
func SimBattleButtonRect(sw, _ int) (x, y, w, h int) {
    lm := uicore.ListMarginPx()
    w, h = uicore.S(160), uicore.S(48)
    x = sw - lm - w
    y = lm + uicore.S(16)
    return
}

// DrawSimBattleButton は一覧画面の「模擬戦」ボタンを描画します。
func DrawSimBattleButton(dst *ebiten.Image, hovered, enabled bool) {
	sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
	x, y, w, h := SimBattleButtonRect(sw, sh)
	base := color.RGBA{40, 60, 100, 255}
	if !enabled {
		base = color.RGBA{70, 70, 70, 255}
	}
	if hovered && enabled {
		base = color.RGBA{60, 90, 150, 255}
	}
	uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
    label := "模擬戦"
    if hovered && enabled {
        label = "> 模擬戦 <"
    }
    uicore.TextDraw(dst, label, uicore.FaceMain, x+uicore.S(24), y+uicore.S(32), uicore.ColText)
}

// DrawAutoRunButton はバトル画面下部の「自動実行/停止」ボタンを描画します。
func DrawAutoRunButton(dst *ebiten.Image, hovered bool, running bool) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    // 位置は API 側の AutoRunButtonRect を使って上位で計算しにくいため、ここでは開始ボタンからの相対は使わず、
    // 呼び出し側で Rect を計算済みと想定して同等スタイルで描画する簡易版とします。
    // ただし、この関数単体では矩形を計算しないため、実プロジェクトでは Draw* と Rect を揃えるのが望ましいです。
    // 互換のため、開始ボタンの右隣（AutoRunButtonRect）を前提に、そこへ描画します。
    // Rect の再計算
    bx, by, bw, bh := uicore.S(0), uicore.S(0), uicore.S(240), uicore.S(60)
    if startX, startY, startW, startH := BattleStartButtonRectCompat(sw, sh); startW > 0 {
        gap := uicore.S(20)
        bx, by, bw, bh = startX+startW+gap, startY, startW, startH
    }
    uicore.DrawFramedRect(dst, float32(bx), float32(by), float32(bw), float32(bh))
    base := color.RGBA{110, 90, 40, 255}
    if running {
        base = color.RGBA{150, 60, 60, 255}
    } else if hovered {
        base = color.RGBA{140, 110, 50, 255}
    }
    vector.DrawFilledRect(dst, float32(bx), float32(by), float32(bw), float32(bh), base, false)
    label := "自動実行"
    if running { label = "停止" }
    uicore.TextDraw(dst, label, uicore.FaceMain, bx+uicore.S(70), by+uicore.S(38), uicore.ColText)
}

// BattleStartButtonRectCompat は widgets から開始ボタン位置へアクセスするための薄い互換関数です。
// 原則として screens に置くべきですが、現状の依存分離を保ちつつ見た目を揃えるための便宜的な実装です。
func BattleStartButtonRectCompat(_ , _ int) (int, int, int, int) {
    // 実体は screens 側。widgets からは参照しない方針のため、
    // ここでは API 経由で取得するのが理想だが、循環参照を避けるため、呼び出し側で Rect を計算済み前提にする。
    // 本関数はダミーとして 0 を返す（呼び出し側で使わない）。
    return 0, 0, 0, 0
}
