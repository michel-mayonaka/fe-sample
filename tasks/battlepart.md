要件定義：FE風シミュレーションRPG・バトルパート（MVP）
0. 用語

ユニット：プレイヤー/敵のキャラ。

タイル：グリッド1マス。

相性：三すくみ（Edge＞Pierce＞Blunt＞Edge）。

ZOC（Zone of Control）：敵隣接マスの通過制限。

1. 目的・非目的
1.1 目的（このMVPで必ず成立させること）

1マップ完結の戦闘：プレイヤー側/敵側のターン交代で全滅 or 勝利条件達成まで進行。

プレイアブルUI：キーボード操作で選択→移動→攻撃→結果表示ができる。

情報の見える化：命中/被ダメ予測UI、相性表示、地形効果を事前提示。

セーブ/ロード：戦闘中スナップショットをJSONで保存・読み込み。

最小AI：敵が“到達→攻撃”の合理的な行動を1手先で選べる。

検証可能性：ヘッドレス検証ツールで同じ結果が再現できる（seed固定）。

1.2 非目的（今回やらない）

クラスチェンジ/支援会話/装備錬成/マップ間進行/ストーリー。

オンライン要素/難易度選択/複雑なスキルツリー。

高度な演出（簡易アニメ/エフェクトの範囲に留める）。

2. 開発スタック・構成

言語：Go 1.22+

描画：Ebitengine（Ebiten）

ルート構成：

/cmd/game          # 実行本体
/pkg/game          # ゲーム状態・ロジック（UIに依存しない）★テスト対象
/pkg/ui            # 入力・レンダ・UI
/assets            # 画像/フォント（仮素材） 
/tools/simcli      # ヘッドレスUpdate Nフレーム
/tools/render      # 1フレーム描画→PNG
/mcp/make_runner   # make verify を叩くMCP（任意）
/artifacts         # 検証成果物（ログ/ssim差分/verify.json）
AGENTS.md
Makefile

3. ゲームループ仕様
3.1 ステートマシン
Title → MapIntro → PlayerPhase
→ (SelectUnit → ShowMove → Move → ActionSelect → CombatForecast → CombatResolve → EndAction)
→ EnemyPhase（AI行動）
→ Victory | Defeat | PlayerPhase

3.2 入力（デフォルト）

←↑→↓：カーソル移動／範囲選択

Z/Enter：決定、X/Esc：キャンセル

Tab：行動未済ユニットへジャンプ

S：セーブ、L：ロード、Space：アニメ高速化

4. データモデル（JSON想定・例）
{
  "tileset": {
    "plain":  {"move_cost":1, "avoid":0,  "def":0 },
    "forest": {"move_cost":2, "avoid":20, "def":1 },
    "fort":   {"move_cost":2, "avoid":15, "def":2, "heal":10 }
  },
  "weapons": {
    "iron_edge":   {"type":"Edge",   "mt":6, "hit":85, "crit":0, "rng":[1,1], "wt":5},
    "iron_pierce": {"type":"Pierce", "mt":5, "hit":90, "crit":5, "rng":[1,2], "wt":5},
    "iron_blunt":  {"type":"Blunt",  "mt":8, "hit":75, "crit":0, "rng":[1,1], "wt":7}
  },
  "map": {
    "w": 16, "h": 12,
    "tiles": [ ["plain","plain","forest", "..."] ],
    "player_units": [
      {"id":"p1","name":"Sable","class":"Vanguard","pos":[2,9],
       "stats":{"hp":22,"str":7,"skl":6,"spd":7,"lck":3,"def":6,"res":1,"mov":5},
       "weapon":"iron_edge", "tags":["infantry"], "lv":1}
    ],
    "enemy_units": [
      {"id":"e1","name":"Bandit","class":"Raider","pos":[11,3],
       "stats":{"hp":20,"str":6,"skl":4,"spd":5,"lck":1,"def":4,"res":0,"mov":5},
       "weapon":"iron_blunt", "tags":["infantry"], "lv":1}
    ],
    "victory":"DefeatAll", "defeat":"LeaderDown"
  }
}

5. ルール（MVP）
5.1 移動

A*またはDijkstraで移動力×地形コスト。

ZOC：敵隣接マスは通過不可（ただしtags: ["flying"]は無視）。

5.2 射程 & 反撃

武器ごとにrng[min,max]。射程外は反撃不可。

5.3 相性（三すくみ）

Edge＞Pierce＞Blunt＞Edge

補正：命中±10、威力±2（MVP基準。定数化）

5.4 命中・ダメージ式（2RN“真ヒット”）
atk_hit = weapon.hit + skl*2 + floor(lck/2) + terrain_hit_bonus(attacker_tile)
def_avo = spd*2 + lck + terrain_avoid(defender_tile)
hit_disp = clamp(atk_hit - def_avo + triangle_hit, 0, 100)

# 2RN
r1, r2 in [0..99]; hit_true = floor((r1+r2)/2)
is_hit = (hit_true < hit_disp)

crit = weapon.crit + floor(attacker.skl/2) - defender.lck
is_crit = (rand < crit%)

raw = attacker.str + weapon.mt + triangle_mt - defender.def
dmg = max(1, is_crit ? floor(raw*2) : raw)

5.5 地形

avoidは回避上昇、defは被ダメ軽減（上式のdefに加算）。

healは自軍ターン開始時に％回復。

6. UI/演出

移動候補の塗りつぶし（到達可能タイル）

行動済みマーカー

戦闘前プレビュー：

表示：命中% / 与ダメ / 反撃ダメ / 必殺% / 相性

色：有利=青、劣勢=橙、中立=灰

簡易戦闘演出：攻撃時のヒット/ミス演出（点滅/SE）

7. 敵AI（最小）

候補：到達マス×攻撃対象を列挙。

評価関数（線形）：
score = E[dmg_to_enemy] - E[dmg_to_self] + kill_bonus(10) + terrain_bonus - death_risk

最高スコア行動を選択。攻撃不可なら高回避地形へ前進。

8. セーブ/ロード

スナップショット内容：マップID、全ユニット状態（HP/pos/weapon/acted）、乱数seed/消費数、フェーズ、ターン数。

形式：artifacts/save_slot1.json 等。

ロード後に同一行動で結果一致（seed再現）。

9. テスト・検証（Make＋MCPで固定）

make lint：golangci-lint run

make test：go test ./pkg/...（ロジック層のみ）

make sim：tools/simcli -scene battle1 -steps 300 -seed 42 > artifacts/sim.json

make render：tools/render artifacts/frame.png && python tools/ssim_check.py golden.png artifacts/frame.png

make verify：上記すべて＋artifacts/verify.json生成（HMAC付き）

AGENTS.md（抜粋）

- 回答前に `make_runner.run_make("verify")` を実行し、payload.exit==0 かつ Proof 行を提示すること。
- 失敗時は結論を出さず、修正案のみ列挙すること。
- UI/UXや式の説明は、必ずゲーム内表示と内部値の一致を確認してから記述すること。

10. 受け入れ基準（Acceptance Criteria）

 タイトル→戦闘→勝利/敗北の一連が完全に操作できる。

 命中/与被ダメ予測が実際の結果と±1ダメ以内で一致（乱数の当たり外れは除く）。

 相性UIが三すくみ通りに反映（命中±10/威力±2）。

 ZOCで迂回が必要な経路が視覚化される。

 セーブ/ロードで再現（同じ操作で同じ結果）。

 make verify が0 exitで完走、verify.json（HMAC付き）生成。

 敵AIが最低1体は合理的ターゲット選択を行う（自滅行動を取らない）。

11. 実装ガイド（抜粋）
11.1 ロジック層（/pkg/game）はUI非依存

Game{Phase, Turn, Units, Map, RNG}

ComputeReachable(unit) / ComputeAttackables(unit, pos)

Forecast(attacker, defender)（上記式の結果を返す）

ApplyCombat(attacker, defender)（結果を適用）

11.2 UI層（/pkg/ui）

状態遷移は明示enumで管理（Select/Move/ActionSelect/Preview/Resolve）。

描画は座標系ユーティリティで一元化（グリッド→スクリーン）。

12. 禁止事項（脱・“シミュレーターだけ”）

CLIだけで完結する実装（UIなし）は禁止。

予測UIと内部計算の不一致を放置しない。

戦闘を自動スキップして結果だけ返すモードをデフォルトにしない。

公式FEの固有名/配色/UI意匠を模倣しない（抽象語彙で置換）。

13. マイルストーン

M1（Day1-2）：/pkg/game のコアロジック・テスト通過、simcliで再現OK。

M2（Day3-4）：/pkg/uiで移動/攻撃UI、プレビュー/解決、相性表示。

M3（Day5）：AI/セーブ/make verify 締め、verify.jsonとSSIM閾値通過。

M4（Day6）：演出・効果音の最低限追加、入力ヘルプ、エッジケース潰し。