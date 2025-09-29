package levelup

import (
    "ui_sample/internal/model"
    uicore "ui_sample/internal/game/service/ui"
)

// Gains はレベルアップで上昇した値を表します。
type Gains struct { Inc uicore.Stats; HPGain int }

// Roll は成長率に基づいて上昇ステータスを抽選します。
func Roll(u uicore.Unit, rnd func() float64) Gains {
    g := Gains{}
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

// Apply は抽選結果をユニットへ反映し、上限とHP整合を保ちます。
func Apply(u *uicore.Unit, gains Gains, levelCap int) {
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

