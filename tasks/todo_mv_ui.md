移設ポリシー（短冊）

service/ui/ … フォント・画像・レイアウト・テーマ・汎用ウィジェット（文脈レス）

scenes/ … 画面（battle, item_list…）・ゲーム文脈のポップアップ

domain/model/ … user.go, inventory.go のようなデータ構造

domain/rules/ … ユーザーや在庫に関する更新ロジック（あれば）

service/ui/widgets/ … モーダル枠・リスト・パネル等の再利用UI

具体的な移設マップ
1) ui/core → internal/game/service/ui/

fonts.go → service/ui/fonts.go

image.go → service/ui/image.go

layout.go → service/ui/layout.go

load.go → service/ui/load.go（アセット読込のUI側ヘルパ）

metrics.go → service/ui/metrics.go（文字幅・行高など）

panel.go → service/ui/panel.go（汎用パネルなら widgets でもOK）

textutil.go → service/ui/text.go

theme.go → service/ui/theme.go

types.go → service/ui/types.go

util.go → service/ui/util.go

ここは**Ebiten寄りの“見た目/配置ツール”**だけにする。ゲームの状態は持たない。

2) ui/widgets → internal/game/service/ui/widgets/

そのまま移動。ゲーム非依存を守る（引数はデータではなく文字列や色、コールバックなど）。

3) ui/screens/* → internal/game/scenes/...

battle.go → scenes/battle/view.go（HUDや盤面の“画面ロジック”）

battle_sim.go → scenes/battle_sim/scene.go

item_list.go → scenes/inventory/item_list.go

weapon_list.go → scenes/inventory/weapon_list.go

list.go / status.go → 各シーンに分割 or scenes/common/ へ再配置

画面（Screen）はSceneの責務。UIツールに入れると境界が曖昧になる。

4) ui/popup/*

choose_unit.go / levelup.go

ゲーム文脈あり → scenes/common/popup/choose_unit.go, .../levelup.go

枠だけの汎用モーダルなら service/ui/widgets/modal.go にベースを置き、
その上にシーン側でコンテンツを載せる二段構えがベスト。

5) ui/api.go

service/ui/api.go に移動して、シーンから使う最小インタフェースを定義：

// service/ui/api.go
package ui
type Painter interface {
    Text(x, y int, s string, style TextStyle)
    Panel(r Rect, style PanelStyle)
    Image(img *ebiten.Image, op *ebiten.DrawImageOptions)
}
type Layout interface { Row(...Widget) Widget; Col(...Widget) Widget /* ... */ }


シーン層は この抽象だけ見る（具体はこのパッケージ内に隠す）。

6) ui/user/* → internal/game/domain/model/

inventory.go → domain/model/inventory.go

user.go → domain/model/user.go

もし操作関数が入っているなら、副作用のない更新は domain/rules/ へ、
I/Oや保存は repository/ へ分離。

最終ディレクトリ（関連部分のみ）
/internal/game/
  service/
    ui/
      api.go
      fonts.go
      image.go
      layout.go
      load.go
      metrics.go
      panel.go
      text.go
      theme.go
      types.go
      util.go
      widgets/
        button.go
        modal.go
        listview.go
        scrollbar.go
  scenes/
    battle/
      scene.go
      view.go
      hud.go
      popup/
        levelup.go
        choose_unit.go
    battle_sim/
      scene.go
    inventory/
      item_list.go
      weapon_list.go
    common/
      widgets/   # シーン間再利用の“文脈薄め”UI（必要なら）
  domain/
    model/
      user.go
      inventory.go
      defs.go
      state.go
    rules/
      inventory_ops.go   # 在庫の純関数更新（必要な場合）
  repository/
    save_file/...
    master_tsv/...

依存の向き（再確認）
scenes/*  →  service/ui (Painter/Layout/Widgets)
scenes/*  →  domain/model (+ rules) → repository/*（port経由）
service/ui は domain を知らない（ゲーム文脈を持たない）

実務Tips（移行を安全に）

先にパッケージ名を直す：package ui → package battle などへ変更してから git mv すると差分が読みやすい。

UIツールのAPIを“細い”抽象に：Painter/Layout を経由させると循環importを防げる。

widgetsはProps駆動に：データ構造を渡さず、表示用Props（テキスト/選択肢/ハンドラ）で渡す。

grep -R "package main" で cmd/ に迷い込んだロジックがないか最後に一掃。

迷った時の判定基準

そのファイルはゲームの意味（ユニット、在庫、レベル）を知っているか？ → Scene/Domain

見た目だけ/汎用コンポーネントか？ → service/ui

保存・読み込みか？ → repository

---

進捗メモ（2025-09-28）
- 完了: ui/core → internal/game/service/ui へ移設（パッケージ名は暫定で `uicore` 維持）。
- 完了: ui/widgets → internal/game/service/ui/widgets へ移設。参照を全体置換。
- 完了: screens → scenes へ移設（一覧/ステータス/バトル/在庫・行構築/模擬戦）。`internal/ui/screens` と `internal/ui/api.go` を撤去。
- 完了: scenes 側で `BattleStartButtonRect`/`AutoRunButtonRect`/`Draw*` 群を提供し、呼び出し元から `ui` 依存を解消。
- 完了: `SetWeaponTable` は `scenes.SetWeaponTable` に移動。再読み込み処理も追随。
- 完了: `LoadUnitsFromUser` は `uicore.LoadUnitsFromUser` を直接利用。
- 保留: ui/popup の scenes/common/popup への移設。
- 保留: service/ui に抽象（Painter/Layout）導入。

次アクション
- screens/* の各描画関数を scenes 下へ整理し、`ui` 経由呼び出しの段階的解消。
- `service/ui` に抽象 API を追加し、scenes は抽象にのみ依存する形へ。

