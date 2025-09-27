// Package popup は各種ポップアップUIの描画を提供します。
package popup

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "golang.org/x/image/font"
    "image/color"
    uicore "ui_sample/internal/game/service/ui"
)

// ChooseUnitItemRect はポップアップ内の i 番目のキャラ項目の矩形を返します。
// 固定で最大5列のグリッドに配置します。
func ChooseUnitItemRect(sw, sh, i, _ int) (x, y, w, h int) {
    pw, ph := popupSize(sw, sh)
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    // グリッド設定
    cols := 5
    if pw < uicore.S(900) { cols = 4 }
    if pw < uicore.S(700) { cols = 3 }
    cellW := (pw - uicore.S(48)) / cols // 左右マージン+隙間
    cellH := uicore.S(160)
    gap := uicore.S(8)
    gridX := px + uicore.S(24)
    gridY := py + uicore.S(88)
    col := i % cols
    row := i / cols
    x = gridX + col*(cellW+gap)
    y = gridY + row*(cellH+gap)
    w = cellW
    h = cellH
    return
}

// DrawChooseUnitPopup はキャラ一覧（アイコン+名前）をポップアップで描画します。
// hover はハイライトするインデックス（-1 で無し）。
func DrawChooseUnitPopup(dst *ebiten.Image, title string, units []uicore.Unit, hover int) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    // 背景ディマー
    vector.DrawFilledRect(dst, 0, 0, float32(sw), float32(sh), color.RGBA{0, 0, 0, 140}, false)
    // パネル
    pw, ph := popupSize(sw, sh)
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    uicore.DrawPanel(dst, float32(px), float32(py), float32(pw), float32(ph))
    uicore.TextDraw(dst, title, uicore.FaceTitle, px+uicore.S(24), py+uicore.S(56), uicore.ColAccent)
    // グリッド描画
    for i, u := range units {
        x, y, w, h := ChooseUnitItemRect(sw, sh, i, len(units))
        // アイテム背景
        base := color.RGBA{30, 45, 78, 255}
        if i == hover { base = color.RGBA{40, 60, 110, 255} }
        vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), base, false)
        uicore.DrawFramedRect(dst, float32(x), float32(y), float32(w), float32(h))
        // ポートレート
        ps := uicore.S(96)
        px2 := float32(x + (w-ps)/2)
        py2 := float32(y + uicore.S(12))
        uicore.DrawFramedRect(dst, px2-float32(uicore.S(2)), py2-float32(uicore.S(2)), float32(ps+uicore.S(4)), float32(ps+uicore.S(4)))
        if u.Portrait != nil {
            uicore.DrawPortrait(dst, u.Portrait, px2, py2, float32(ps), float32(ps))
        } else {
            uicore.DrawPortraitPlaceholder(dst, px2, py2, float32(ps), float32(ps))
        }
        // 名前
        tw := int(font.MeasureString(uicore.FaceSmall, u.Name) >> 6)
        uicore.TextDraw(dst, u.Name, uicore.FaceSmall, x+(w-tw)/2, y+uicore.S(128), uicore.ColText)
    }
    // ヒント
    hint := "クリックで選択 / Escでキャンセル"
    tw := int(font.MeasureString(uicore.FaceSmall, hint) >> 6)
    uicore.TextDraw(dst, hint, uicore.FaceSmall, px+(pw-tw)/2, py+ph-uicore.S(20), color.RGBA{210, 220, 240, 255})
}

func popupSize(sw, sh int) (w, h int) {
    w = int(float32(sw) * 0.8)
    h = int(float32(sh) * 0.72)
    if w > uicore.S(1480) { w = uicore.S(1480) }
    if h > uicore.S(900) { h = uicore.S(900) }
    return
}

