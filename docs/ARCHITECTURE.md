# アーキテクチャ指針 v1（Scene/Actor/Service）

最終更新: 2025-09-28

## 原則
- Scene / Actor / Service の3層。ECS は必要箇所に限定して後付け可能。
- Update 順序を契約として固定し、デバッグ容易性を最優先。
- データ駆動（TSV/JSON）＋開発時ホットリロードで調整コストを削減。
- 描画は固定レイヤ＋アトラス化。純ロジックは純関数化し `go test` 可能に。
- 「使う側が定義する小さな interface」を基本とし疎結合化。

## ディレクトリ構成（あるべき姿）

`/internal/game/`
- `app/`            ebiten.Game の心臓。フレーム駆動、SceneStack 管理。
- `ctx.go`          フレームコンテキスト（DT/Frame/Screen/Input/Services）。
- `scene.go`        `Scene` インタフェースと `SceneStack` 実装。
- `scenes/`         画面群（title/battle/result など）。
- `actor/`          `IActor` と具体 Actor（unit/cursor/fx）。
- `service/`        入出力・資産・音・カメラ・UI・ホットリロード。
- `domain/`         UI 非依存のゲームルール（model/rules/port）。
- `repository/`     port 実装（tsv/embed/save_file/save_web/migrate）。
- `world/`          表示側の盤面制御（タイル/ハイライト/アダプタ）。
- `render/`         レイヤ描画（背景/世界/影/FX/UI）。
- `data/`           `tables/` `maps/`（CSV/TSV/JSON）。
- `assets/`         アトラス/フォント/音源。
- `util/`           RNG/幾何/IO/プラットフォーム分岐。

## コア API（最小）
```go
// game.Scene
Update(ctx *game.Ctx) (next game.Scene, err error)
Draw(screen *ebiten.Image)

// game.SceneStack
Current() Scene; Push(Scene); Pop() Scene; Replace(Scene); Size() int

// game.Ctx
DT/Frame/ScreenW/ScreenH と Input/Assets/Audio/Camera/UI/Rand/Debug

// actor.IActor
Update(*game.Ctx) bool; Draw(*ebiten.Image); Layer() int
```

## Update 順序契約
1) Input（Snapshot 固定）
2) Script/AI（重い処理は分割）
3) Physics/Board（座標・ZOC・当たり）
4) Resolve（コマンド確定・状態更新）
5) Audio（キュー適用）
6) GC/Spawn（死活整理）
7) Draw（レイヤ順）

すべての Scene はこの順序を遵守します。

## データ駆動とホットリロード
- マスタは `mst_*`、ユーザは `usr_*` の接頭辞で管理。
- テーブルは CSV/TSV/JSON を許容。開発時は `tables/*.tsv` と `maps/*.json` をホットリロード。
- マスタ→ユーザの上書きモデルを維持し、一貫した読取 API を提供。

## 入力（抽象アクション）
- `service.Input` を追加（Up/Down/Left/Right/Confirm/Cancel/Menu/Next/Prev）
  に加え、実運用便宜のため `OpenWeapons/OpenItems/EquipToggle/Slot1..5/Unassign` を定義。
- Snapshot/Press/Down を提供。キーボード→アクションへの投影で UI/パッド差分を吸収。

Press/Down の運用ルール（暫定）
- Press: UIトグル/決定/戻るなど「瞬間」操作（例: Confirm で戦闘開始、Cancel で戻る）。
- Down: 押下継続の意味がある操作（例: Menu=Backspace 長押しでデータリロード）。

## テスト戦略
- `pkg/game` の純関数をユニットテストで担保。
- Adapter/Repo/レイアウト計算など UI 非依存の純関数をテスト化。
- 予測数値/A* 経路等はゴールデンファイルで検証。`replay` による再現性確保。
