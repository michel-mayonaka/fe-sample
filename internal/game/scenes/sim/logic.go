package sim

import (
	"fmt"
	scenes "ui_sample/internal/game/scenes"
	gcore "ui_sample/pkg/game"
)

// scAdvance は意図を解釈して状態機械を前進させ、副作用（戦闘実行/ログ切替/遷移フラグ）を反映します。
func (s *Sim) scAdvance(intents []scenes.Intent) {
	for _, any := range intents {
		it, ok := any.(Intent)
		if !ok {
			continue
		}
		switch it.Kind {
		case intentBack:
			s.pop = true
		case intentRunOne:
			leftFirst := (s.turn%2 == 1)
			s.runOne(leftFirst)
		case intentToggleAuto:
			s.auto = !s.auto
			if s.auto {
				s.logPopup = false
			}
		case intentSetTerrainAtt:
			s.attSel = it.Index
			switch it.Index {
			case 0:
				s.attTerrain = gcore.Terrain{}
			case 1:
				s.attTerrain = gcore.Terrain{Avoid: 20, Def: 1}
			case 2:
				s.attTerrain = gcore.Terrain{Avoid: 15, Def: 2}
			}
		case intentSetTerrainDef:
			s.defSel = it.Index
			switch it.Index {
			case 0:
				s.defTerrain = gcore.Terrain{}
			case 1:
				s.defTerrain = gcore.Terrain{Avoid: 20, Def: 1}
			case 2:
				s.defTerrain = gcore.Terrain{Avoid: 15, Def: 2}
			}
		}
	}
	// 自動実行
	canStart := s.simAtk.HP > 0 && s.simDef.HP > 0
	leftFirst := (s.turn%2 == 1)
	if s.auto && canStart && !s.logPopup {
		if s.autoCD > 0 {
			s.autoCD--
		} else {
			s.runOne(leftFirst)
			s.autoCD = 10
		}
	}
}

// runOne は 1 ターン分の戦闘を実行し、結果ログを追加します。
func (s *Sim) runOne(leftFirst bool) {
	if leftFirst {
		a, d, lines := SimulateBattleCopyWithTerrain(s.simAtk, s.simDef, s.attTerrain, s.defTerrain, s.E.RNG)
		s.simAtk, s.simDef = a, d
		s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simAtk.Name)}, lines...)
	} else {
		a, d, lines := SimulateBattleCopyWithTerrain(s.simDef, s.simAtk, s.defTerrain, s.attTerrain, s.E.RNG)
		s.simDef, s.simAtk = a, d
		s.logs = append([]string{fmt.Sprintf("ターン %d 先攻: %s", s.turn, s.simDef.Name)}, lines...)
	}
	s.logPopup = true
	s.turn++
}
