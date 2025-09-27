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
  app.go            // ebiten.Game（心臓）
  ctx.go            // Frame Context（Δt, 入力スナップ, カメラ等）
  scene.go          // Scene, SceneStack
  scenes/
    title/
      title.go
    battle/
      battle.go     // 盤面/ターン/AI入口
    result/
  actor/
    actor.go        // IActor: Update/Draw/Alive
    unit.go         // FEユニット（最低限）
    cursor.go
    fx.go           // 汎用エフェクト
  service/
    assets.go       // 画像/音/フォントのキャッシュ
    input.go        // 抽象入力（Action/Axis）
    audio.go        // BGM/SEキュー
    camera.go
    ui.go           // 文字描画, ウィジェット軽ラッパ
    hotload.go      // データホットリロード（debugビルドのみ）
  world/
    world.go        // タイル/地形/高さ/遮蔽
    turn.go         // フェーズ/状態機械
    predict.go      // 命中/クリ/ダメ予測（純関数）
    path.go         // A*（純関数＋キャッシュ）
  render/
    layers.go       // 背景/世界/影/エフェクト/UI の順序制御
  data/
    tables/         // CSV/TSV（ユニット/武器/地形/状態/AI重み…）
    maps/           // マップ（CSVタイル + JSONオブジェクト）
  assets/
    images/atlas.png
    images/atlas.json   // スプライト座標
    fonts/
    audio/
  util/
    rng.go          // seed固定の乱数
    geom.go         // 2D幾何・格子座標変換
    file.go         // 読み込み
/test/
  replay/           // リプレイ記録（入力列）
  golden/           // 予測・経路のゴールデンファイル
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
