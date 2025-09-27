package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "ui_sample/internal/game"
)

// listScene は一覧画面を Scene 化した薄いアダプタです。
type listScene struct{ g *Game }

func (s *listScene) Update(_ *game.Ctx) (game.Scene, error) {
    mx, my := ebiten.CursorPosition()
    s.g.updateListMode(mx, my)
    // 遷移検出: 旧ロジックが modeSimBattle をセットしたら Scene をプッシュ
    if s.g.mode == modeSimBattle {
        // 旧モードは維持しない
        s.g.mode = modeList
        return &simScene{g: s.g}, nil
    }
    return nil, nil
}

func (s *listScene) Draw(screen *ebiten.Image) { s.g.drawList(screen) }

// simScene は模擬戦画面を Scene 化した薄いアダプタです。
type simScene struct{ g *Game }

func (s *simScene) Update(_ *game.Ctx) (game.Scene, error) {
    // 既存の modeSimBattle 部分と同等の更新を実行
    // ログポップアップ中は閉じる操作のみ受け付け
    if s.g.simLogPopup {
        if s.g.closeSimLogIfRequested() {
            return nil, nil
        }
        return nil, nil
    }
    // 戻る（ボタン/Cancel）
    s.g.handleSimBack()
    if s.g.mode == modeList { // 旧ロジックは modeList をセット
        return nil, nil // Pop は呼び出し側で検出
    }
    // 自動実行/ボタン/開始など通常更新
    s.g.updateSimBattleCore()
    return nil, nil
}

func (s *simScene) Draw(screen *ebiten.Image) { s.g.drawSimBattle(screen) }

