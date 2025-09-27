package popup

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "image/color"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/game/service/ui"
)

// LevelUpGains はレベルアップで上昇した値を表します。
type LevelUpGains struct { Inc uicore.Stats; HPGain int }

// RollLevelUp はユニットの成長率に基づき上昇ステータスを抽選します。
func RollLevelUp(u uicore.Unit, rnd func() float64) LevelUpGains {
    g := LevelUpGains{}
    prob := func(p int) bool { return p > 0 && rnd()*100 < float64(p) }
    if prob(u.Growth.HP) { g.HPGain++ }
    if prob(u.Growth.Str) { g.Inc.Str++ }
    if prob(u.Growth.Mag) { g.Inc.Mag++ }
    if prob(u.Growth.Skl) { g.Inc.Skl++ }
    if prob(u.Growth.Spd) { g.Inc.Spd++ }
    if prob(u.Growth.Lck) { g.Inc.Lck++ }
    if prob(u.Growth.Def) { g.Inc.Def++ }
    if prob(u.Growth.Res) { g.Inc.Res++ }
    if prob(u.Growth.Mov) { g.Inc.Mov++ }
    return g
}

// ApplyGains は抽選結果をユニットへ反映し、上限とHP整合を保ちます。
func ApplyGains(u *uicore.Unit, gains LevelUpGains, levelCap int) {
    if u.Level < levelCap { u.Level++ }
    u.HPMax += gains.HPGain
    u.HP += gains.HPGain
    if u.HP > u.HPMax { u.HP = u.HPMax }
    u.Stats.Str += gains.Inc.Str
    u.Stats.Mag += gains.Inc.Mag
    u.Stats.Skl += gains.Inc.Skl
    u.Stats.Spd += gains.Inc.Spd
    u.Stats.Lck += gains.Inc.Lck
    u.Stats.Def += gains.Inc.Def
    u.Stats.Res += gains.Inc.Res
    u.Stats.Mov += gains.Inc.Mov
    if caps, err := model.LoadClassCapsJSON("db/master/mst_class_caps.json"); err == nil {
        if c, ok := caps.Find(u.Class); ok {
            clamp := func(v, m int) int { if m > 0 && v > m { return m }; if v < 0 { return 0 }; return v }
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

// DrawLevelUpPopup はレベルアップ結果のポップアップを描画します。
func DrawLevelUpPopup(dst *ebiten.Image, u uicore.Unit, gains LevelUpGains) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    overlay := color.RGBA{0, 0, 0, 160}
    vector.DrawFilledRect(dst, 0, 0, float32(sw), float32(sh), overlay, false)
    pw, ph := uicore.S(520), uicore.S(480)
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    uicore.DrawPanel(dst, float32(px), float32(py), float32(pw), float32(ph))
    uicore.TextDraw(dst, "レベルアップ!", uicore.FaceTitle, px+uicore.S(24), py+uicore.S(56), uicore.ColAccent)
    uicore.TextDraw(dst, fmt.Sprintf("Lv %d", u.Level), uicore.FaceMain, px+uicore.S(24), py+uicore.S(96), uicore.ColText)
    y := py + uicore.S(140)
    line := uicore.S(34)
    drawInc := func(label string, v int) {
        if v > 0 {
            uicore.TextDraw(dst, fmt.Sprintf("%s +%d", label, v), uicore.FaceMain, px+uicore.S(40), y, uicore.ColAccent)
            y += line
        }
    }
    if gains.HPGain > 0 { drawInc("HP", gains.HPGain) }
    drawInc("力", gains.Inc.Str)
    drawInc("魔力", gains.Inc.Mag)
    drawInc("技", gains.Inc.Skl)
    drawInc("速さ", gains.Inc.Spd)
    drawInc("幸運", gains.Inc.Lck)
    drawInc("守備", gains.Inc.Def)
    drawInc("魔防", gains.Inc.Res)
    drawInc("移動", gains.Inc.Mov)
    uicore.TextDraw(dst, "クリックで閉じる", uicore.FaceSmall, px+pw-uicore.S(180), py+ph-uicore.S(24), uicore.ColText)
}

