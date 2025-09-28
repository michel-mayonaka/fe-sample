# Scene単位の入力処理一元化（設計メモ）

## 背景
- 現状、入力のプリミティブは `internal/game/service/input.go`（抽象アクション/スナップショット）に実装。
- `internal/game/app/game.go` の `Update()` で毎フレーム `Input.Snapshot()` を呼び、`updateGlobalToggles()` で一部グローバル操作（ヘルプ表示、データ再読込）を処理。
- 各 Scene は `ctx.Input.Press(...)` を用いつつ、マウス入力は `ebiten.CursorPosition`/`inpututil.IsMouseButtonJustPressed` を直接参照。

本タスクは「プリミティブは据え置きつつ、Scene単位でキーマップを組み込み・切替できる」構造への移行方針をまとめる。

## 目標
- Sceneごとの追加/上書きキーマップを、Push/Pop に同期して適用する。
- グローバル操作（ヘルプ/再読込）も Action 化し、可能なら共通の入力面に寄せる。
- 既存の `Press/Down` 互換性を維持しつつ、導入は段階的に行えるようにする。

## 設計方針
- 入力オーバーレイ（レイヤ）方式：
  - `Input` に「ベースのキーマップ」と「オーバーレイキーマップのスタック」を持たせる。
  - 解決順序は「最上位オーバーレイ → … → ベース」。最初にヒットした Action を採用。
- Scene キーマップは Scene Push で `PushMap()`、Pop で `PopMap()` する。
- 長押し（再読込トグル用）などは `Input` がフレーム継続数を保持し、`HeldFor(a Action) int` を提供（任意）。
- マウスは当面、各 Scene が直接参照（プリミティブ据え置き）。将来的に抽象化する余地は残す。

## 仕様案

### Input API 拡張（`internal/game/service/input.go`）
```go
type Input struct {
    curr, prev [ActionCount]bool
    mapKey     map[ebiten.Key]Action // ベース
    layers     []map[ebiten.Key]Action // 追加レイヤ（Scene毎）
    heldFrames [ActionCount]int // 任意: 長押しカウント
}

func (i *Input) PushMap(m map[ebiten.Key]Action) { i.layers = append(i.layers, m) }
func (i *Input) PopMap() { if n := len(i.layers); n>0 { i.layers = i.layers[:n-1] } }

// SnapshotWith 内のキー→Action解決手順（擬似コード）
// for key in allKeys:
//   if any layer maps key -> a { curr[a] = true; continue }
//   else if base map maps key -> a { curr[a] = true }

func (i *Input) HeldFor(a Action) int { return i.heldFrames[a] } // 任意
```

### Runner 拡張（`internal/game/app/runner.go`）
```go
// Scene が実装可能なオプションIF
type SceneKeymap interface {
    Keymap() map[ebiten.Key]gamesvc.Action
}

// Push時: 最前面Sceneが SceneKeymap を実装していれば Input.PushMap()
// Pop時: 対応するレイヤを PopMap()
```

### Game 側（`internal/game/app/game.go`）
- `updateGlobalToggles()` は、可能なら `Action` ベースに寄せる。
  - 例: `HelpToggle`（押下でトグル）、`ReloadHold`（`HeldFor()`>=閾値）。
  - 互換のため、当面は現行実装のままでも可。段階移行。

### Scene 実装パターン（例）
```go
// character_list.go
func (s *List) Keymap() map[ebiten.Key]gamesvc.Action {
    return map[ebiten.Key]gamesvc.Action{
        ebiten.KeyW: gamesvc.OpenWeapons,
        ebiten.KeyI: gamesvc.OpenItems,
    }
}
// Update 内は従来通り ctx.Input.Press(...) を使うだけ（差分最小）。
```

## 実装手順（段階的）
1) Input にレイヤAPI（PushMap/PopMap）と解決順序の変更、必要なら `HeldFor` を追加。単体テストを `input_test.go` に拡充。
2) Runner に SceneKeymap 検出/適用を追加。Stack の Push/Pop タイミングに連動。
3) `character_list` と `status` から導入し、キーマップを Scene 提供に切替（挙動に変化がないことを確認）。
4) `sim`/`inventory` も順次対応。不要になった Game 側のハードコードは段階的に削減。
5) `updateGlobalToggles` の Action 化（任意）。長押しのしきい値は 30f（約0.5秒/60FPS）を既定。

## テスト方針
- 入力:
  - レイヤが上書き優先で解決されること（同一キー競合時）。
  - レイヤの Push/Pop 順に応じて期待動作になること。
  - `Press/Down` の遷移が既存と互換であること。
  - `HeldFor`（導入時）でフレーム数が正しく増減すること。
- ランナー:
  - ダミーSceneを用意し、Pushで Keymap 適用、Popで解除されることを確認。

## 影響範囲 / 互換性
- 既存 Scene の `ctx.Input.Press/Down` 呼び出しはそのまま利用可能。
- import 循環は発生しない（Scene → gamesvc の一方向参照）。
- マウス入力は当面据え置き（将来、抽象レイヤ追加の余地あり）。

## リスク / 留意点
- レイヤ解除漏れ（Pop忘れ）に注意。Runner の Stack 操作と1:1で管理する実装にする。
- 同一キーを複数Sceneが使う場合、最前面Sceneが優先される設計で合意する。
- 長押しやリピート（キーリピート）を導入する場合は、UI意図に応じたデフォルト値/可変設定を検討。

## 工数目安
- 実装＋テスト: 2〜3時間（既存テスト流用/拡張前提）。

---
担当メモ（2025-09-28）
- 現在の入力処理位置: `service/input.go`、Scene内のマウス直参照、Gameの一部グローバル処理。
- 本メモの方針であれば、導入は小さなPRに分割可能（Input→Runner→Scene順）。

