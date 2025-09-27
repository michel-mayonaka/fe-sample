package uiscreens

import (
    "fmt"
    "math/rand"
    "ui_sample/internal/adapter"
    "ui_sample/internal/model"
    uicore "ui_sample/internal/ui/core"
    gcore "ui_sample/pkg/game"
)

// SimulateBattleCopy はコピーを用いた簡易模擬戦を行い、結果ログを返します（永続化なし）。
// 互換API: 地形なし。
func SimulateBattleCopy(atk, def uicore.Unit, rng *rand.Rand) (uicore.Unit, uicore.Unit, []string) {
    return SimulateBattleCopyWithTerrain(atk, def, gcore.Terrain{}, gcore.Terrain{}, rng)
}

// SimulateBattleCopyWithTerrain は左右地形を考慮して1手交換+追撃までをシミュレーションします（永続化なし）。
func SimulateBattleCopyWithTerrain(atk, def uicore.Unit, attT, defT gcore.Terrain, rng *rand.Rand) (uicore.Unit, uicore.Unit, []string) {
    a := atk
    d := def
    logs := []string{"模擬戦開始", fmt.Sprintf("%s vs %s", a.Name, d.Name)}
    // 武器テーブル（共有/フォールバック）
    wt := weaponTable()
    if wt == nil {
        if wtf, err := model.LoadWeaponsJSON("db/master/mst_weapons.json"); err == nil { wtShared = wtf; wt = wtShared }
    }
    if wt == nil { return a, d, append(logs, "[エラー] 武器定義を読み込めませんでした") }
    ga := adapter.UIToGame(wt, a)
    gd := adapter.UIToGame(wt, d)
    // 攻→反→追
    logs = append(logs, a.Name+" の攻撃")
    ga2, gd2, line1 := gcore.ResolveRoundAt(ga, gd, attT, defT, rng)
    if line1 != "" { logs = append(logs, line1) }
    d.HP = gd2.S.HP
    if d.HP < 0 { d.HP = 0 }
    // 反撃（射程チェック）
    dist := 1
    canCounter := gd2.W.RMin <= dist && dist <= gd2.W.RMax
    if d.HP > 0 && canCounter {
        logs = append(logs, d.Name+" の反撃")
        gd3, ga3, line2 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, rng)
        if line2 != "" { logs = append(logs, line2) }
        a.HP = ga3.S.HP
        if a.HP < 0 { a.HP = 0 }
        ga2, gd2 = ga3, gd3
    } else {
        a.HP = ga2.S.HP
    }
    // 追撃（AS差>=3）
    if d.HP > 0 && gcore.DoubleAdvantage(ga, gd) {
        logs = append(logs, a.Name+" の追撃")
        ga4, gd4, line3 := gcore.ResolveRoundAt(ga2, gd2, attT, defT, rng)
        if line3 != "" { logs = append(logs, line3) }
        d.HP = gd4.S.HP
        if d.HP < 0 { d.HP = 0 }
        a.HP = ga4.S.HP
    } else if a.HP > 0 && canCounter && gcore.DoubleAdvantage(gd, ga) {
        logs = append(logs, d.Name+" の追撃")
        gd4, ga4, line4 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, rng)
        if line4 != "" { logs = append(logs, line4) }
        a.HP = ga4.S.HP
        if a.HP < 0 { a.HP = 0 }
        d.HP = gd4.S.HP
    }
    if a.HP <= 0 { logs = append(logs, a.Name+" は倒れた") }
    if d.HP <= 0 { logs = append(logs, d.Name+" は倒れた") }
    logs = append(logs, "模擬戦終了")
    return a, d, logs
}

