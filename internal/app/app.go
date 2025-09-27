package app

import (
    "fmt"
    "math/rand"
    "ui_sample/internal/adapter"
    "ui_sample/internal/repo"
    uicore "ui_sample/internal/ui/core"
    "ui_sample/internal/user"
    gcore "ui_sample/pkg/game"
)

// App はユースケースの最小集合を束ねます。
type App struct {
    Weapons repo.WeaponsRepo
    Users   repo.UserRepo
    RNG     *rand.Rand
}

// New はリポジトリと乱数源を注入して App を生成します。
func New(users repo.UserRepo, weapons repo.WeaponsRepo, rng *rand.Rand) *App {
    return &App{Weapons: weapons, Users: users, RNG: rng}
}

// RunBattleRound は選択中ユニットと次ユニットの1ラウンド戦闘を解決し、
// UIユニット配列を更新、ユーザセーブへ反映・保存します。
func (a *App) RunBattleRound(units []uicore.Unit, selIndex int, attT, defT gcore.Terrain) ([]uicore.Unit, []string, bool, error) {
    if a == nil || len(units) < 2 {
        return units, nil, false, nil
    }
    if selIndex < 0 || selIndex >= len(units) {
        return units, nil, false, nil
    }
    atkIdx := selIndex
    defIdx := (selIndex + 1) % len(units)
    atk := units[atkIdx]
    def := units[defIdx]

    ga := adapter.UIToGame(a.Weapons.Table(), atk)
    gd := adapter.UIToGame(a.Weapons.Table(), def)

    logs := []string{"戦闘開始", atk.Name + " の攻撃"}
    ga2, gd2, line1 := gcore.ResolveRoundAt(ga, gd, attT, defT, a.RNG)
    if line1 != "" { logs = append(logs, line1) }
    // 反撃
    if gd2.S.HP > 0 {
        dist := 1
        canCounter := gd2.W.RMin <= dist && dist <= gd2.W.RMax
        if canCounter {
            logs = append(logs, def.Name+" の反撃")
            gd3, ga3, line2 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, a.RNG)
            if line2 != "" { logs = append(logs, line2) }
            ga2, gd2 = ga3, gd3
        }
    }
    logs = append(logs, "戦闘終了")

    // UIへHP反映
    atk.HP = ga2.S.HP
    def.HP = gd2.S.HP
    // 使用回数を1消費（攻撃側先頭装備）
    if len(atk.Equip) > 0 && atk.Equip[0].Uses > 0 { atk.Equip[0].Uses-- }
    units[atkIdx] = atk
    units[defIdx] = def

    // ユーザテーブルへ反映・保存（両者）
    if a.Users != nil {
        if c, ok := a.Users.Find(atk.ID); ok {
            c.HP = atk.HP
            c.HPMax = atk.HPMax
            if len(c.Equip) > 0 && len(atk.Equip) > 0 {
                c.Equip[0].Uses = atk.Equip[0].Uses
            }
            a.Users.Update(c)
        }
        if c2, ok := a.Users.Find(def.ID); ok {
            c2.HP = def.HP
            c2.HPMax = def.HPMax
            a.Users.Update(c2)
        }
        if err := a.Users.Save(); err != nil {
            return units, logs, true, fmt.Errorf("save user: %w", err)
        }
    }
    return units, logs, true, nil
}

// PersistUnit は UI ユニットの現在値をユーザセーブへ反映して保存します。
func (a *App) PersistUnit(u uicore.Unit) error {
    if a == nil || a.Users == nil { return nil }
    c, ok := a.Users.Find(u.ID)
    if !ok { return nil }
    c.Level = u.Level
    c.HP = u.HP
    c.HPMax = u.HPMax
    c.Stats = user.Stats{Str: u.Stats.Str, Mag: u.Stats.Mag, Skl: u.Stats.Skl, Spd: u.Stats.Spd, Lck: u.Stats.Lck, Def: u.Stats.Def, Res: u.Stats.Res, Mov: u.Stats.Mov}
    a.Users.Update(c)
    return a.Users.Save()
}
