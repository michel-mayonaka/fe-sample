package app

import (
    "fmt"
    "math/rand"
    "ui_sample/internal/adapter"
    "ui_sample/internal/model"
    "ui_sample/internal/repo"
    uicore "ui_sample/internal/ui/core"
    "ui_sample/internal/user"
    "ui_sample/internal/assets"
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
    // 攻撃回数（耐久消費用）
    atkCount, defCount := 0, 0
    ga2, gd2, line1 := gcore.ResolveRoundAt(ga, gd, attT, defT, a.RNG)
    atkCount++
    if line1 != "" { logs = append(logs, line1) }
    // 反撃
    dist := 1
    canCounter := gd2.W.RMin <= dist && dist <= gd2.W.RMax
    if gd2.S.HP > 0 && canCounter {
        logs = append(logs, def.Name+" の反撃")
        gd3, ga3, line2 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, a.RNG)
        defCount++
        if line2 != "" { logs = append(logs, line2) }
        ga2, gd2 = ga3, gd3
    }
    // 追撃（AS差>=3）
    asAtk := gcore.AttackSpeed(ga)
    asDef := gcore.AttackSpeed(gd)
    if gd2.S.HP > 0 && gcore.DoubleAdvantage(ga, gd) {
        logs = append(logs, atk.Name+" の追撃")
        ga4, gd4, line3 := gcore.ResolveRoundAt(ga2, gd2, attT, defT, a.RNG)
        atkCount++
        if line3 != "" { logs = append(logs, line3) }
        ga2, gd2 = ga4, gd4
    } else if ga2.S.HP > 0 && canCounter && gcore.DoubleAdvantage(gd, ga) {
        _ = asAtk; _ = asDef // for symmetry / readability
        logs = append(logs, def.Name+" の追撃")
        gd4, ga4, line4 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, a.RNG)
        defCount++
        if line4 != "" { logs = append(logs, line4) }
        ga2, gd2 = ga4, gd4
    }
    logs = append(logs, "戦闘終了")

    // UIへHP反映
    atk.HP = ga2.S.HP
    def.HP = gd2.S.HP
    // 使用回数を消費（攻撃1回ごとに1消費）
    if len(atk.Equip) > 0 && atk.Equip[0].Uses > 0 {
        use := atkCount
        if use > atk.Equip[0].Uses { use = atk.Equip[0].Uses }
        atk.Equip[0].Uses -= use
    }
    if len(def.Equip) > 0 && def.Equip[0].Uses > 0 {
        use := defCount
        if use > def.Equip[0].Uses { use = def.Equip[0].Uses }
        def.Equip[0].Uses -= use
    }
    units[atkIdx] = atk
    units[defIdx] = def

    // ユーザテーブルへ反映・保存（両者）
    if a.Users != nil {
        if c, ok := a.Users.Find(atk.ID); ok {
            c.HP = atk.HP
            c.HPMax = atk.HPMax
            if len(c.Equip) > 0 && len(atk.Equip) > 0 { c.Equip[0].Uses = atk.Equip[0].Uses }
            a.Users.Update(c)
        }
        if c2, ok := a.Users.Find(def.ID); ok {
            c2.HP = def.HP
            c2.HPMax = def.HPMax
            if len(c2.Equip) > 0 && len(def.Equip) > 0 { c2.Equip[0].Uses = def.Equip[0].Uses }
            a.Users.Update(c2)
        }
        if err := a.Users.Save(); err != nil {
            return units, logs, true, fmt.Errorf("save user: %w", err)
        }
    }
    return units, logs, true, nil
}

// ReloadData は JSON バックエンドのキャッシュを再読み込みします。
func (a *App) ReloadData() error {
    if a == nil { return nil }
    if a.Weapons != nil { if err := a.Weapons.Reload(); err != nil { return err } }
    // 画像キャッシュもここでクリア（UI層の依存を薄く保ちつつ、利便性を優先）
    assets.Clear()
    return nil
}

// WeaponsTable は共有用の武器テーブル参照を返します。
// UIはこれを `ui.SetWeaponTable` に渡して利用します。
func (a *App) WeaponsTable() *model.WeaponTable {
    if a == nil || a.Weapons == nil { return nil }
    return a.Weapons.Table()
}

// PersistUnit は UI ユニットの現在値をユーザセーブへ反映して保存します。
func (a *App) PersistUnit(u uicore.Unit) error {
    if a == nil || a.Users == nil { return nil }
    c, ok := a.Users.Find(u.ID)
    if !ok { return nil }
    c.Level = u.Level
    c.HP = u.HP
    c.HPMax = u.HPMax
    c.Stats = user.Stats{Str: u.Stats.Str, Mag: u.Stats.Mag, Skl: u.Stats.Skl, Spd: u.Stats.Spd, Lck: u.Stats.Lck, Def: u.Stats.Def, Res: u.Stats.Res, Mov: u.Stats.Mov, Bld: u.Stats.Bld}
    a.Users.Update(c)
    return a.Users.Save()
}
