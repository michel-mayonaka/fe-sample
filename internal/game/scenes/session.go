package scenes

import (
    uicore "ui_sample/internal/game/service/ui"
    scpopup "ui_sample/internal/game/scenes/common/popup"
)

// Session は UI シーン間で共有される“表示用の一時状態”をまとめた構造体です。
// - 論理データの保存やリポジトリアクセスは UseCases に委譲します。
// - ここは UI の選択状態やポップアップ表示状態など、描画に近い値のみを保持します。
type Session struct {
    // 一覧/選択
    Units    []uicore.Unit
    SelIndex int

    // ステータス/在庫で共有する状態
    PopupActive     bool
    PopupGains      scpopup.LevelUpGains
    PopupJustOpened bool
    CurrentSlot     int
    SelectingEquip  bool
    SelectingIsWeapon bool
    InvTab          int // 0=武器,1=アイテム
    HoverInv        int
}

