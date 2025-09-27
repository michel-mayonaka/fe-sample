package uicore

import (
    "image/color"
    "sync"

    "github.com/hajimehoshi/ebiten/v2"
    textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
    "golang.org/x/image/font"
)

var (
    faceCacheMu sync.RWMutex
    faceCache   = map[font.Face]*textv2.GoXFace{}
)

// v2Face は x/image/font.Face を text/v2 の Face にラップします。
// GoXFace はキャッシュを内部に持つため、インスタンスを使い回します。
func v2Face(face font.Face) *textv2.GoXFace {
    if face == nil { return nil }
    faceCacheMu.RLock()
    if f, ok := faceCache[face]; ok {
        faceCacheMu.RUnlock()
        return f
    }
    faceCacheMu.RUnlock()
    f := textv2.NewGoXFace(face)
    faceCacheMu.Lock()
    faceCache[face] = f
    faceCacheMu.Unlock()
    return f
}

// TextDraw は text/v1 互換の座標系（x,y=ベースライン）で文字列を描画します。
// 旧APIの `text.Draw(dst, s, face, x, y, col)` と同等の見た目を目指します。
func TextDraw(dst *ebiten.Image, s string, face font.Face, x, y int, col color.Color) {
    vf := v2Face(face)
    if vf == nil || dst == nil || s == "" { return }
    m := vf.Metrics()
    var op textv2.DrawOptions
    // text/v1 互換: (x,y)=ベースライン → v2は上端原点なので Ascent 分だけ上にずらす
    op.GeoM.Translate(float64(x), float64(y)-m.HAscent)
    op.ColorScale.ScaleWithColor(col)
    // 行送りは v1 と同等（Height = Ascent+Descent+LineGap）
    op.LineSpacing = m.HAscent + m.HDescent + m.HLineGap
    op.PrimaryAlign = textv2.AlignStart
    op.SecondaryAlign = textv2.AlignStart
    textv2.Draw(dst, s, vf, &op)
}
