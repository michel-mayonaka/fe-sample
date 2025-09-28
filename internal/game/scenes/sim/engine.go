package sim

import (
    "fmt"
    "math/rand"
    "ui_sample/internal/adapter"
    uicore "ui_sample/internal/game/service/ui"
    gdata "ui_sample/internal/game/data"
    "ui_sample/internal/model"
    gcore "ui_sample/pkg/game"
)

// SimulateBattleCopy はコピーを用いた簡易模擬戦（平地）
func SimulateBattleCopy(atk, def uicore.Unit, rng *rand.Rand) (uicore.Unit, uicore.Unit, []string) {
    return SimulateBattleCopyWithTerrain(atk, def, gcore.Terrain{}, gcore.Terrain{}, rng)
}

// SimulateBattleCopyWithTerrain は左右地形を考慮して1手交換+追撃までをシミュレーションします（永続化なし）。
func SimulateBattleCopyWithTerrain(atk, def uicore.Unit, attT, defT gcore.Terrain, rng *rand.Rand) (uicore.Unit, uicore.Unit, []string) {
    a := atk
    d := def
    logs := []string{"模擬戦開始", fmt.Sprintf("%s vs %s", a.Name, d.Name)}
    // 武器テーブル（共有/フォールバック）
    var wt *model.WeaponTable
    if p := gdata.Provider(); p != nil { wt = p.WeaponsTable() }
    // 最終的に provider 前提にするが、安全のため nil チェック
    if wt == nil { return a, d, append(logs, "[エラー] 武器定義を読み込めませんでした") }
    ga := adapter.UIToGame(wt, a)
    gd := adapter.UIToGame(wt, d)
    // 攻→反→追（結果を同一行に出力）
    ga2, gd2, line1 := gcore.ResolveRoundAt(ga, gd, attT, defT, rng)
    if line1 != "" { logs = append(logs, a.Name+" の攻撃 → "+line1) } else { logs = append(logs, a.Name+" の攻撃") }
    d.HP = gd2.S.HP
    if d.HP < 0 { d.HP = 0 }
    // 反撃（射程チェック）
    dist := 1
    canCounter := gd2.W.RMin <= dist && dist <= gd2.W.RMax
    if d.HP > 0 && canCounter {
        gd3, ga3, line2 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, rng)
        if line2 != "" { logs = append(logs, d.Name+" の反撃 → "+line2) } else { logs = append(logs, d.Name+" の反撃") }
        a.HP = ga3.S.HP
        if a.HP < 0 { a.HP = 0 }
        ga2, gd2 = ga3, gd3
    } else {
        a.HP = ga2.S.HP
    }
    // 追撃（AS差>=3）
    if d.HP > 0 && gcore.DoubleAdvantage(ga, gd) {
        ga4, gd4, line3 := gcore.ResolveRoundAt(ga2, gd2, attT, defT, rng)
        if line3 != "" { logs = append(logs, a.Name+" の追撃 → "+line3) } else { logs = append(logs, a.Name+" の追撃") }
        d.HP = gd4.S.HP
        if d.HP < 0 { d.HP = 0 }
        a.HP = ga4.S.HP
    } else if a.HP > 0 && canCounter && gcore.DoubleAdvantage(gd, ga) {
        gd4, ga4, line4 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, rng)
        if line4 != "" { logs = append(logs, d.Name+" の追撃 → "+line4) } else { logs = append(logs, d.Name+" の追撃") }
        a.HP = ga4.S.HP
        if a.HP < 0 { a.HP = 0 }
        d.HP = gd4.S.HP
    }
    // 勝利ログ
    switch {
    case a.HP <= 0 && d.HP <= 0:
        logs = append(logs, "相打ち")
    case a.HP <= 0:
        logs = append(logs, "勝利: "+d.Name)
    case d.HP <= 0:
        logs = append(logs, "勝利: "+a.Name)
    }
    logs = append(logs, "模擬戦終了")
    return a, d, logs
}
