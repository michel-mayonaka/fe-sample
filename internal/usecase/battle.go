package usecase

import (
	"fmt"
	"math/rand"
	"time"
	"ui_sample/internal/adapter"
	uicore "ui_sample/internal/game/service/ui"
	gcore "ui_sample/pkg/game"
)

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
	// RNG フォールバック: 未注入時でもパニックしないように都度生成
	rng := a.RNG
	if rng == nil {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	ga2, gd2, line1 := gcore.ResolveRoundAt(ga, gd, attT, defT, rng)
	atkCount++
	if line1 != "" {
		logs = append(logs, line1)
	}
	// 反撃
	dist := 1
	canCounter := gd2.W.RMin <= dist && dist <= gd2.W.RMax
	if gd2.S.HP > 0 && canCounter {
		logs = append(logs, def.Name+" の反撃")
		gd3, ga3, line2 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, rng)
		defCount++
		if line2 != "" {
			logs = append(logs, line2)
		}
		ga2, gd2 = ga3, gd3
	}
	// 追撃（AS差>=3）
	if gd2.S.HP > 0 && gcore.DoubleAdvantage(ga, gd) {
		logs = append(logs, atk.Name+" の追撃")
		ga4, gd4, line3 := gcore.ResolveRoundAt(ga2, gd2, attT, defT, rng)
		atkCount++
		if line3 != "" {
			logs = append(logs, line3)
		}
		ga2, gd2 = ga4, gd4
	} else if ga2.S.HP > 0 && canCounter && gcore.DoubleAdvantage(gd, ga) {
		logs = append(logs, def.Name+" の追撃")
		gd4, ga4, line4 := gcore.ResolveRoundAt(gd2, ga2, defT, attT, rng)
		defCount++
		if line4 != "" {
			logs = append(logs, line4)
		}
		ga2, gd2 = ga4, gd4
	}
	logs = append(logs, "戦闘終了")

	// UIへHP反映
	atk.HP = ga2.S.HP
	def.HP = gd2.S.HP
	// 使用回数を消費（攻撃1回ごとに1消費）: UI 表示更新
	if len(atk.Equip) > 0 && atk.Equip[0].Uses > 0 {
		use := atkCount
		if use > atk.Equip[0].Uses {
			use = atk.Equip[0].Uses
		}
		atk.Equip[0].Uses -= use
	}
	if len(def.Equip) > 0 && def.Equip[0].Uses > 0 {
		use := defCount
		if use > def.Equip[0].Uses {
			use = def.Equip[0].Uses
		}
		def.Equip[0].Uses -= use
	}
	units[atkIdx] = atk
	units[defIdx] = def

	// ユーザテーブルへ反映・保存（両者）: HP等は usr_characters へ、耐久は usr_weapons/items へ
	if a.Users != nil {
		if c, ok := a.Users.Find(atk.ID); ok {
			c.HP = atk.HP
			c.HPMax = atk.HPMax
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

	// 耐久は usr_weapons.json / usr_items.json に保存
	// 攻撃側
	if len(atk.Equip) > 0 && atkCount > 0 && a.Inv != nil {
		_ = a.Inv.Consume(atk.Equip[0].ID, atkCount)
		_ = a.Inv.Save()
	}
	// 防御側
	if len(def.Equip) > 0 && defCount > 0 && a.Inv != nil {
		_ = a.Inv.Consume(def.Equip[0].ID, defCount)
		_ = a.Inv.Save()
	}
	return units, logs, true, nil
}
