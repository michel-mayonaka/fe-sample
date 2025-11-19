下記アーキテクチャ設計に見直しを検討したい。
現行の状況と合わせて調整した方が良い内容がある場合は提案もして欲しい
調整が終わった後doc下にドキュメント化を行う。

0. 全体方針（原則）

Scene / Actor / Service の3層。ECSは必要な島だけ（AIや大量エフェクト）に後付け可。

Update順序を契約化して“どこで何をやるか”を固定化。デバッグしやすい。

データ駆動（CSV/TSV/JSON）＋**ホットリロード（開発時のみ）**で調整コスト削減。

描画はレイヤ固定＋アトラス化でドローコール最小化。

純ロジックは純関数化してgo test可能に。リプレイ（入力＋seed固定）で再現性確保。

**“使う側が定義する小さなinterface”**が基本。実装は暗黙実装で疎結合。

1. ディレクトリ構成（最初から育つ形）
/cmd/game/main.go

/internal/game/
  app/                   # ebiten.Game（心臓）: 旧 app.go をここへ
    app.go
  ctx.go                 # Frame Context（Δt, 入力スナップ, カメラ等）
  scene.go               # Scene, SceneStack

  scenes/
    title/
      title.go
    battle/
      battle.go          # 盤面/ターン/AI入口（port経由でデータ参照）
    result/

  actor/
    actor.go             # IActor: Update/Draw/Alive
    unit.go              # FEユニット（表現用。ドメインのUnitStateはdomain/modelへ）
    cursor.go
    fx.go                # 汎用エフェクト

  service/               # ← UIサービス専用（入出力/表示系のみ）
    assets.go            # 画像/音/フォントのキャッシュ
    input.go             # 抽象入力（Action/Axis）
    audio.go             # BGM/SEキュー
    camera.go
    ui.go                # 文字描画, ウィジェット軽ラッパ
    hotload.go           # データホットリロード（debugビルドのみ, Master差し替え通知）

  domain/                # ← UI非依存の“ゲームのルール”
    model/               # 構造体: Save/UnitState/Stats/Defs…
      defs.go            # ItemDef/WeaponDef/ClassDef/TerrainDef
      state.go           # GameState/ItemRef/UnitState/Progress…
      combat.go          # CombatInput/Outcome
      world.go           # 盤面の軽量表現（座標/地形IDなど）
      types.go           # 列挙/ID型/タグ
    rules/               # 純関数: 予測/A*/致傷/状態機械
      predict.go         # 命中/クリ/ダメ予測
      path.go            # A*（純関数＋任意キャッシュフック）
      turn.go            # フェーズ遷移（状態機械）
      injury.go          # 致傷テーブル
    port/                # “使う側が定義する”データアクセスIF
      master.go          # MasterRepository interface
      save.go            # SaveRepository interface

  repository/            # ← port を実装する“外側”
    master_tsv/          # TSV/CSVローダ（PC開発用）
      master_tsv.go
      index.go
    master_embed/        # go:embed配布用（リリースビルド向け）
      master_embed.go
    save_file/           # JSONセーブ（PC）
      save_file.go
    save_web/            # IndexedDB/LocalStorage（WASM）
      save_web.go
    migrate/             # セーブ互換/移行
      migrate.go
      versions.go

  world/                 # 画面側の盤面制御（描画や選択/ハイライト等）
    world.go             # タイル/地形/高さ/遮蔽（表示・選択）
    overlay.go           # 可視範囲/移動範囲/ハイライト
    adapter.go           # domain.model.World 等との相互変換

  render/
    layers.go            # 背景/世界/影/エフェクト/UI の順序制御

  data/
    tables/              # CSV/TSV（ユニット/武器/地形/状態/AI重み…）
    maps/                # マップ（CSVタイル + JSONオブジェクト）

  assets/
    images/atlas.png
    images/atlas.json    # スプライト座標
    fonts/
    audio/

  util/
    rng.go               # seed固定の乱数
    geom.go              # 2D幾何・格子座標変換
    file.go              # 読み込み
    platform.go          # 環境分岐（デスクトップ/WASM）

/test/
  replay/                # リプレイ記録（入力列）
  golden/                # 予測・経路のゴールデンファイル
Makefile
go.mod


2. コアAPI（最小の型とインタフェース）
// scene.go
package game

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update(ctx *Ctx) (next Scene, err error) // nilで継続 / 次シーンでPush（Pop/Replaceは別メソッド）
	Draw(screen *ebiten.Image)
}

type SceneStack struct{ stack []Scene }
/* Current/Push/Pop/Replace 実装は省略 */

// ctx.go（フレーム共通の読み取り専用情報＋サービス）
type Ctx struct {
	DT      float64
	Frame   uint64
	ScreenW int
	ScreenH int

	Input  *service.Input
	Assets *service.Assets
	Audio  *service.Audio
	Camera *service.Camera
	UI     *service.UI
	Rand   *util.Rand
	Debug  bool
}

// actor/actor.go
type IActor interface {
	Update(ctx *game.Ctx) bool            // falseで破棄
	Draw(dst *ebiten.Image)
	Layer() int                           // 描画順（世界=100, FX=200, UI=300 など）
}

3. Update順序“契約”

固定順序（どのSceneでも同じ思想）：

Input：Snapshot()でフレーム固定

Script/AI：意思決定。重い処理は分割（N体/フレーム）

Physics/Board：座標更新／ZOC／当たり

Resolve：コマンド確定（ダメ適用・状態更新）

Audio：BGM/SEキュー適用

GC/Spawn：死活整理・生成

Draw：レイヤ順で一気に

“どこでやるか”を迷わないのが品質。違反は失敗という自覚で揃える。

4. データ駆動（CSV/JSONスキーマ最小）
4.1 ユニット/クラス表（TSV例）
id	name	class	hp	str	skl	spd	int	res	mov	weapon_types	passives
u_soldier	新兵	兵士	18	6	5	5	2	1	4	lance	shield_wall
u_archer	射手	弓兵	16	5	5	5	2	1	4	bow	cover_fire


文字→IDは**snake_case**で一意に。

passivesはカンマ区切りで複数可。

4.2 武器表
id	name	type	mt	hit	crit	wgt	range	effects
wp_lance1	訓練用ランス	lance	6	80	0	6	1	armor_bonus:10
wp_bow1	短弓	bow	4	75	0	3	2	none

4.3 地形表
id  name  avoid  def  height  cover  move_cost
plain 平地  0     0    0       0      1
forest森    20    1    0       1      2
high丘     10    0    1       0      2
wall壁     0     4    0       1      INF

4.4 マップ（タイルCSV + JSON）

CSV（数字＝地形IDインデックス）

JSON（敵配置、勝利条件、初期天候など）

{
  "width": 20, "height": 12, "tileset": "tileset01",
  "units": [
    {"id":"u_soldier", "side":"player", "x":2, "y":10, "level":1, "weapon":"wp_lance1"},
    {"id":"u_goblin",  "side":"enemy",  "x":12,"y":4,  "ai":"ambush"}
  ],
  "win": {"type":"escape", "x":19, "y":0},
  "weather": "fog",
  "scripts": ["intro_cutscene", "first_blood_tutorial"]
}


ホットリロード（debugタグ）でtables/*.tsv/maps/*.json更新を即反映できるように。

5. 戦術ロジック（純関数の核）
// world/predict.go — 命中/クリ/ダメの最小式（テスト可能）
type CombatInput struct {
	Atk, Def UnitState
	Terrain  Terrain
	Weather  Weather
}
type CombatOutcome struct {
	HitChance int
	CritChance int
	DamageMin int
	DamageMax int
}

func Predict(ci CombatInput) CombatOutcome {
	hit := clamp(100 + ci.Atk.Skl*2 + heightBonus(ci.Atk, ci.Def) - ci.Def.Spd*2 - fogPenalty(ci.Weather), 5, 99)
	crit := clamp(ci.Atk.Skl/2 + backstab(ci.Atk, ci.Def) - ci.Def.Res/3, 0, 80)
	mt := ci.Atk.Wpn.Mt - armor(ci.Def, ci.Atk.Wpn)
	return CombatOutcome{HitChance: hit, CritChance: crit, DamageMin: max(0, mt-2), DamageMax: max(0, mt+2)}
}


I/Oが明確な純関数。ゴールデンテストで回帰を守る。

6. 描画設計（レイヤ・アトラス・整数座標）

レイヤ順：Background(50) → Tile(100) → Units(120) → Shadows(140) → FX(200) → UI(300)

画像は1枚アトラスにまとめ、atlas.jsonで座標管理。

ドット絵は整数配置＋FilterNearest。

テキストはビットマップフォント（フォント描画はドローコールが嵩む）。

// render/layers.go
func DrawLayer(dst *ebiten.Image, actors []actor.IActor, min, max int) {
	for _, a := range actors {
		if l := a.Layer(); l >= min && l < max { a.Draw(dst) }
	}
}

7. 入力（抽象アクション＋スナップショット）
// service/input.go（短縮版）
type Action int
const (Up Down Left Right Confirm Cancel Menu Next Prev ActionCount)

type Input struct{ curr, prev [ActionCount]bool; mapKey map[ebiten.Key]Action }
func (i *Input) Snapshot() { /* キー→Actionに投影し、prev/curr更新 */ }
func (i *Input) Press(a Action) bool { return !i.prev[a] && i.curr[a] }
func (i *Input) Down(a Action) bool  { return i.curr[a] }


抽象アクションに落とすと、パッド/キーボード/リバインドに強い。

押した瞬間（トグルUI等）と押下中（移動）を分ける。

8. セーブ/ロード & リプレイ（再現性）

セーブ：GameState{Scene名, World状態, RNG seed, インベントリ…}をJSONで保存。

リプレイ：seed + 入力列（フレームごとのActionビット列）を記録。

バグ再現はこのペアを渡せば一撃。

type Replay struct {
	Seed int64
	Inputs [][]bool // frame -> [ActionCount]bool
}

9. テスト戦略（最初から回す）

ユニットテスト：Predict, Path, TurnStateMachine は純関数としてgo test。

ゴールデン：予測HUDの数値・A*の経路ノード列をファイルに保存し比較。

プロパティ（rapid / fuzz）：

「命中は0〜100」「高所有利＞＝0」「ZOCでコスト非減少」などを自動検証。

E2E軽テスト：scenes/battleに擬似入力を流し、Worldの状態遷移をアサート。

10. ビルド・ツール・スクリプト

Makefile（抜粋）

run: ## 開発実行
	go run ./cmd/game

test: ## ユニットテスト
	go test ./...

golden: ## ゴールデン更新
	go test ./world -run TestPredict -update

race: ## データ競合検出
	go run -race ./cmd/game

lint: ## ざっくりlint（reviveなど）
	revive -config revive.toml ./...

assets: ## アトラス生成（例: go-bindata/独自スクリプト）
	go run tools/atlasgen/main.go assets/images/ > assets/images/atlas.json


go.modポリシー

外部依存は最小（ebiten, go-cmp, testifyどちらか、必要ならrapid）。

画像パイプラインは自前スクリプトで十分（PNG集約→atlas.json生成）。

11. コーディング規約（運用の型）

命名：idはsnake_case、Goの識別子はUpperCamel（export）/lowerCamel（internal）。

レシーバ：状態変更→ポインタ、純関数→値。

公開/非公開：外へ出したくなった瞬間にインタフェース（最小）を切る。

パニック禁止（ライブラリ層）。致命だけpanic、あとはerrorで明示。

コメント：exportedは名前で始まる一文。reviveのexportedが通るように。

データは不可変視：ロードしたテーブルは読み取り専用にして事故を防ぐ。

12. パフォーマンス・懸念と対策

NewImage/SubImage乱発 → 起動時キャッシュ、毎フレーム作らない。

ドローコール多発 → アトラス＋同スプライトの連続描画、UIはまとめて。

巨大マップのA* → コスト場でダイクストラ事前計算 or 分割更新。

AIフレーム落ち → N体/フレームで回す、非アクティブ敵は粗い思考。

浮動小数のにじみ → 整数配置、拡大は整数倍、FilterはNearest。

データ肥大 → テーブルは列削る勇気、複合効果はタグ＋スクリプト少数に留める。

13. 最初の実装マイルストーン（2週間スプリント想定）

Day1–2：App/SceneStack/Ctx/Input/Assetsの雛形＋title→battle遷移

Day3–5：World（タイル/地形/高さ/遮蔽）＋カーソル移動

Day6–7：ユニット配置→Predict()で予測HUD表示

Day8–9：コマンド確定→ダメ適用→ターン切替

Day10–11：A*実装・移動範囲表示・ZOC

Day12–13：簡易AI（最短接近→攻撃）を分割更新で

Day14：ゴールデン/リプレイ/ホットリロード接続、最低限の演出（ヒットFX）

14. 参考テンプレ（最小Battleシーンの骨）
type Battle struct{
	world   *world.World
	actors  []actor.IActor
	hud     *HUD
}

func NewBattle(w *world.World, a []actor.IActor) *Battle {
	sort.Slice(a, func(i, j int) bool { return a[i].Layer() < a[j].Layer() })
	return &Battle{world:w, actors:a, hud:NewHUD()}
}

func (b *Battle) Update(ctx *game.Ctx) (game.Scene, error) {
	// 1. 入力
	ctx.Input.Snapshot()

	// 2. AI/スクリプト（重いなら分割）
	b.world.UpdateAI(ctx)

	// 3. 盤面・衝突・ZOC
	b.world.UpdateBoard(ctx)

	// 4. 確定（攻撃/回復/状態）
	b.world.Resolve(ctx)

	// 5. 音
	ctx.Audio.Flush()

	// 6. GC/Spawn
	b.pruneActors()

	return nil, nil
}

 func (b *Battle) Draw(dst *ebiten.Image) {
	b.world.DrawTiles(dst)
	for _, a := range b.actors { if a.Layer()<200 { a.Draw(dst) } }
	b.world.DrawShadows(dst)
	for _, a := range b.actors { if a.Layer()>=200 && a.Layer()<300 { a.Draw(dst) } }
	b.hud.Draw(dst)
}

---

進捗ログ（2025-09-28）

- Phase0: スケルトン導入（完了）
  - 追加: `internal/game/{ctx.go, scene.go, actor/actor.go, service/{input.go,assets.go,audio.go,camera.go,ui.go,hotload.go}, render/layers.go, util/rng.go}`
  - 目的: Scene/Actor/Service の最小 API を先行追加し、既存 UI と併存可能に。
- ドキュメント（完了）
  - `docs/architecture/README.md`: 本ファイルの方針を現状に合わせて整理（更新順序契約/移行計画/現状→目標マッピング）。
  - `docs/SPECS/reference/api.md`: `Ctx/Scene/SceneStack/IActor/service.Input` を追記。
- 軽い実配線（着手）
  - `cmd/ui_sample/main.go`: 抽象入力 `service.Input` を導入し、Backspace リロードのみ置換（`BindKey(Backspace→Menu)`→`Down(Menu)`）。従来の挙動を維持。
  - ビルド確認: `go build ./internal/game/...` と `go build ./cmd/ui_sample` 成功。

次アクション（Phase1 部分適用）

- 入力: `updateList/Status/SimBattle` 内の一部キー操作を `service.Input` へ段階移行（Press/Down 切替ルールを明文化）。
- リロード導線: `updateGlobalToggles` の Backspace リロードを `App.ReloadData()` に寄せ、画像キャッシュ `assets.Clear()` と一括化（UI層の直接I/O削減）。
- 画面遷移: 小径で `SceneStack` を導入（Title→List または List→SimBattle）し、契約運用を実地化。
- ドキュメント更新: 本ログ追記の継続、`docs/architecture/README.md` に Press/Down の運用表を追加。

進捗ログ（2025-09-28 午前・Phase1 一部適用）

- 抽象入力の段階導入（Confirm/Cancel）：
  - `cmd/ui_sample/main.go`
    - `modeBattle`: 戦闘開始=Confirm、戻る=Cancel、ログ閉じ=Confirm に置換（マウス操作は維持）。
    - `modeSimBattle`: 戻る=Cancel、ログ閉じ=Confirm、キーボード戦闘開始=Confirm。自動実行スクロール等は従来の矢印/PgUp/PgDn を維持。
    - 模擬戦のユニット選択ポップアップ: キャンセル=Cancel（Esc/X から移行）。
    - ステータス→一覧/在庫→一覧の戻る: Cancel で統一（ボタン/マウスは維持）。
- リロード導線の統一：
  - Backspace 長押し（= `service.Input.Menu`）で `App.ReloadData()` を呼び、`ui.SetWeaponTable()` と `assets.Clear()` を一括適用。
  - `modeSimBattle` 内の毎フレーム再読み込みを撤去（パフォーマンス改善・意図の一元化）。
- 誤爆防止：
  - リロードは「長押し（約0.5秒）」でのみ発火するようフレームカウンタを導入（`reloadHold`）。
  - Status 画面の装備解除ショートカット（Delete/Backspace）は存続し、短押しではリロードが走らない。
- ビルド確認: `go build ./cmd/ui_sample` 成功。

次アクション（Phase1 継続）
- 入力: List/Status での残キー（W/I/E/数字スロット）を `service.Input` へ順次寄せ替え（UI文言も更新）。
- 画面遷移: `SceneStack` の最小導入（List→SimBattle）を試験的に実装し、Update順序契約の検証を開始。
- テスト: `service.Input` の Snapshot/Press/Down の単体テストを追加。

進捗ログ（2025-09-28 昼・Phase1 継続適用）

- 抽象入力の拡張と適用：
  - 追加アクション: `OpenWeapons/OpenItems/EquipToggle/Slot1..5/Unassign`。
  - `cmd/ui_sample/main.go`
    - 一覧: W/I → `OpenWeapons/OpenItems` に置換。
    - ステータス: E → `EquipToggle`、数字1..5 → `Slot1..5`、装備解除 → `Unassign`（Delete）。
  - 地形切替（1/2/3）は暫定で直接キーのまま（今後検討）。
- ドキュメント更新：`docs/architecture/README.md`/`docs/SPECS/reference/api.md` に拡張アクションを追記。
- ビルド確認: `go build ./cmd/ui_sample` 成功。

ドキュメント移管（2025-09-28）

- ARCHITECTURE.md は「あるべき姿」に限定。以下の現状/移行系の内容を本ファイルへ移管：
  - 現状→目標マッピング（参考）
    - Application: `internal/app`（ユースケース）… 継続利用。
    - Domain/Rules: `pkg/game` … 継続利用（テスト済）。
    - Repository: `internal/repo` … 継続利用（JSON/キャッシュ）。
    - UI: `internal/ui/...` … 継続利用（将来 `service.UI` 経由へ薄層化）。
    - Assets: `internal/assets` … 現状使用（将来 `service.Assets` に統合）。
  - 段階的移行計画（Phase 0→4）
    1) Phase0: `internal/game/{ctx,scene,actor,service,render,util}` 追加（非侵入）
    2) Phase1: 入力抽象の導入と既存キー置換（本タスクで進行中）
    3) Phase2: 部分画面で `SceneStack` 導入（List→SimBattle）
    4) Phase3: 資産/画像キャッシュの `service.Assets` への集約＋App.Reload* 統合
    5) Phase4: マップを `world+actor+render` で分割しレイヤ描画に移行
  - 採用/保留（メモ）
    - 採用: Scene/Actor/Service の最小 API、Update 順序契約、抽象入力
    - 保留: ECS 本格導入、TSV/CSV への全面移行、Hotloader 実装、`embed` 化
  - Hotloader 現状: 実装はプレースホルダ（no-op）。今後、`tables/*.tsv`/`maps/*.json` の監視→通知に対応予定。

進捗ログ（2025-09-28 午後・SceneStack 最小導入）

- SceneStack を導入（旧UIと併存）：
  - `cmd/ui_sample/scenes.go`: `listScene` と `simScene` を追加（`game.Scene` 実装）。
  - `Game` に `useScenes` と `stack` を追加。`NewGame` で `listScene` をプッシュ。
  - `Update/Draw` を Scene 優先に分岐。Scene 使用時は旧 `mode` の分岐をスキップ。
  - 遷移: 一覧→模擬戦は `updateListMode` 内の既存遷移フラグ（`modeSimBattle`）を勘案し、Scene をプッシュ。戻るは `modeList` への復帰を検知してポップ。
- 重複ロジックの抽出：
  - `closeSimLogIfRequested`/`handleSimBack`/`updateSimBattleCore` を追加し、Scene更新から利用。
- ビルド確認: `go build ./cmd/ui_sample` 成功。

フォローアップ修正（2025-09-28 夕）

- 回帰対応: List→Status/Inventory へ遷移不能になっていた問題を修正。
  - `cmd/ui_sample/scenes.go`
    - `statusScene` と `invScene` を新規追加。
    - `listScene.Update` で `modeStatus/modeInventory` を検知して各 Scene を Push。
  - `cmd/ui_sample/main.go`
    - Pop 条件を `simScene/statusScene/invScene` に拡張。
  - 動作: List→Status/Inventory/SimBattle の遷移／戻るが復旧。
  - `go build ./cmd/ui_sample` 成功。

次セッションへの引き継ぎ（優先度順）

1) app 層への寄せ（Phase2 準備）
   - `internal/app` に SceneStack 管理を移譲（`App.Update(ctx)/Draw` を追加）。
   - `game.Ctx` を `DT/Frame` 付きで生成し、`NewGame` にタイマを導入。
   - `App.ReloadData()` を Backspace 長押しのハンドラに集約（main から切り離し）。

2) 入力抽象の適用完了
   - 地形切替（1/2/3, Shift+）を `service.Input` の拡張（TerrainAtt1..3 / TerrainDef1..3）で置換。
   - 画面ヘルプ文言（W/I/E/DELETE）を抽象アクション名・説明に同期。

3) ドキュメント
   - `docs/architecture/README.md`: Ctx/Scene/Update順序の簡易図を追加（図示のみ、進捗は記載しない）。

4) テスト
   - `service.Input` の `Snapshot/Press/Down` のユニットテスト（長押し/連打境界）。
   - Scene 遷移の最小 E2E（list→status→back, list→inv→back, list→sim→back）。

進捗ログ（2025-09-28 夜・Phase1 仕上げ + ctx）

- 抽象入力の拡張（地形切替）とUI置換：
  - 追加アクション: `TerrainAtt1..3`, `TerrainDef1..3` を `service.Input` に定義。
  - 実装: `SnapshotWith` を導入し、Shift+1/2/3 を防御側、1/2/3 を攻撃側にマッピング。
  - UI更新: 旧 `inpututil.IsKeyJustPressed(1/2/3 + Shift)` を全面置換（`cmd/ui_sample/main.go`）。
- Ctx の時間情報を実装：
  - `game.Ctx` の `DT/Frame` を `NewGame`/`Update` から供給（Scene駆動時）。
  - 目的: 将来的な AI/スクリプト分割やアニメ時間管理の基盤。
- テスト追加：
  - `internal/game/service/input_test.go`: 
    - `Confirm` の Press/Down 挙動、`Menu` 長押しの Press 一度きり性、`Shift+2`→`TerrainDef2`/`2`→`TerrainAtt2` のマッピングを検証。
- ドキュメント更新：
  - `docs/SPECS/reference/api.md`: 新アクション群と `SnapshotWith` を追記。

残タスク（Phase2 準備継続）
- SceneStack 管理の `internal/game/app` への移譲（ランナー追加）
  - `Runner`（仮）: `stack game.SceneStack`, `Update(ctx)`, `Draw(screen)` を提供。
  - 当面は Pop 条件のコールバック（例: `AfterUpdate(sc) bool`）で `cmd` 側のモード復帰検知を委譲。
- `cmd/ui_sample` の Scene 実体を `internal/...` へ段階移行し、`app` 層から直接呼べる形に整理。

進捗ログ（2025-09-28 夜遅・SceneStack 移譲）

- `internal/game/app/runner.go` を実装し、`cmd/ui_sample` から Runner を利用する形に移譲。
  - `Game` 構造体: `stack` → `runner` へ置換、`AfterUpdate` で Pop 条件（一覧へ戻ったら Pop）を委譲。
  - `Update/Draw`: 直接 Stack を触らず Runner の `Update/Draw` を呼び出すよう変更。
- ビルド/テスト: `go build ./cmd/ui_sample` / `go test ./...` 成功。

追加進捗（2025-09-28 深夜・Scene 内部移設）

- `internal/game/scenes` を追加し、`List/Status/Inventory/Sim` の4 Scene を Host 越しに実装。
- `cmd/ui_sample` 側に Host アダプタ `scenesHost` を追加。
- 旧 `cmd/ui_sample/scenes.go` を削除し、初期 Push を `scenes.List` に変更。

フォローアップ（2025-09-28 深夜・main.go スリム化）

- `cmd/ui_sample/main.go` から責務を分割。起動エントリのみを残し、以下へ移設：
  - `cmd/ui_sample/game_core.go`: 画面定数・`Game` 構造体・共通ヘルパ
  - `cmd/ui_sample/app_bootstrap.go`: `NewGame` と Repo 初期化
  - `cmd/ui_sample/game_loop.go`: `Update/Draw/Layout` とグローバル操作
  - `cmd/ui_sample/mode_list.go`: 一覧モードの更新/描画
  - `cmd/ui_sample/mode_status.go`: ステータス更新/描画 + 装備同期
  - `cmd/ui_sample/mode_inventory.go`: 在庫更新/描画
  - `cmd/ui_sample/mode_sim.go`: 模擬戦更新/描画
  - `cmd/ui_sample/scenes_host.go`: `scenes.Host` アダプタ
- 目的: エントリと実装の分離、可読性/保守性の向上、差分の小分け。

設計変更（2025-09-28 さらに整理・scenesへ集約）

- 入力・更新・描画ロジックを `internal/game/scenes` 下へ移設。
  - 追加: `scenes.Env`（App/ユーザテーブル/ユニット配列など共有状態）
  - 実装: `List/Status/Inventory/Sim` が `game.Scene` を実装し、`Update(ctx *game.Ctx)` で抽象入力（`ctx.Input`）を参照。
  - Pop判定: `ShouldPop() bool` を導入し、`Runner.AfterUpdate` で type assert によりポップ制御。
- `cmd/ui_sample` からはモード別関数を撤去し、Runner 駆動に一本化。

次の移行ステップ（提案）
- Scene 実体を `cmd/` から `internal/game/scenes` 配下へ移設（差分を最小化しつつ、UI依存部は残留）。
- Runner に `Replace`/`PopAll` 等のヘルパーを追加（必要性が出たら）。
