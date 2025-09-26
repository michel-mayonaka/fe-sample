package ui

import (
    "fmt"
    "image/color"
    text "github.com/hajimehoshi/ebiten/v2/text" //nolint:staticcheck // TODO: text/v2
    "github.com/hajimehoshi/ebiten/v2/vector"
    "ui_sample/internal/model"
)

type LevelUpGains struct { Inc Stats; HPGain int }

func RollLevelUp(u Unit, rnd func() float64) LevelUpGains {
    g := LevelUpGains{}
    prob := func(p int) bool { return p > 0 && rnd()*100 < float64(p) }
    if prob(u.Growth.HP)  { g.HPGain++ }
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

func ApplyGains(u *Unit, gains LevelUpGains, levelCap int) {
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
            clamp := func(v, m int) int { if m>0 && v>m { return m }; if v<0 { return 0 }; return v }
            u.HPMax = clamp(u.HPMax, c.HPMax); u.HP = clamp(u.HP, u.HPMax)
            u.Stats.Str = clamp(u.Stats.Str, c.Str); u.Stats.Mag = clamp(u.Stats.Mag, c.Mag)
            u.Stats.Skl = clamp(u.Stats.Skl, c.Skl); u.Stats.Spd = clamp(u.Stats.Spd, c.Spd)
            u.Stats.Lck = clamp(u.Stats.Lck, c.Lck); u.Stats.Def = clamp(u.Stats.Def, c.Def)
            u.Stats.Res = clamp(u.Stats.Res, c.Res); u.Stats.Mov = clamp(u.Stats.Mov, c.Mov)
        }
    }
}

func DrawLevelUpPopup(dst *ebiten.Image, u Unit, gains LevelUpGains) {
    sw, sh := dst.Bounds().Dx(), dst.Bounds().Dy()
    overlay := color.RGBA{0, 0, 0, 160}
    vector.DrawFilledRect(dst, float32(0), float32(0), float32(sw), float32(sh), overlay, false)
    pw, ph := 520, 480
    px := (sw - pw) / 2
    py := (sh - ph) / 2
    drawPanel(dst, float32(px), float32(py), float32(pw), float32(ph))
    textDraw(dst, "レベルアップ!", faceTitle, px+24, py+56, colAccent)
    textDraw(dst, fmt.Sprintf("Lv %d", u.Level), faceMain, px+24, py+96, colText)
    y := py + 140; line := 34
    drawInc := func(label string, v int){ if v>0 { textDraw(dst, fmt.Sprintf("%s +%d", label, v), faceMain, px+40, y, colAccent); y+=line } }
    if gains.HPGain>0 { drawInc("HP", gains.HPGain) }
    drawInc("力", gains.Inc.Str); drawInc("魔力", gains.Inc.Mag); drawInc("技", gains.Inc.Skl)
    drawInc("速さ", gains.Inc.Spd); drawInc("幸運", gains.Inc.Lck); drawInc("守備", gains.Inc.Def)
    drawInc("魔防", gains.Inc.Res); drawInc("移動", gains.Inc.Mov)
    textDraw(dst, "クリックで閉じる", faceSmall, px+pw-180, py+ph-24, colText)
}

