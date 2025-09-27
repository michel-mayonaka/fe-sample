# バトルパート（MVP）実装計画（rev.1)

本計画は tasks/battlepart.md の要件に基づき、現状実装（一覧/ステータス/簡易戦闘・模擬戦、UI分割済）からMVPに到達するための段階案です。実装順と受け入れ基準を明示します。

## フェーズ0：構成整備（短期）
- ディレクトリ: `/pkg/game` を新設（UI非依存のドメインロジック）。
- データ: `db/master/` に三すくみ表・地形表（後続）を追加予定。
- 既存UIは維持（模擬戦/簡易戦闘は後で /pkg/game へ置換）。

## フェーズ1：コアロジック（/pkg/game）
- 型: `Unit/Stats/Weapon/Terrain/Map/RNG`（最小スキーマ）。
- 関数: `Forecast(att,def)`（2RN/相性/地形/追撃/射程は段階導入）、`Resolve(att,def)`。
- テスト: 2RN、三すくみ補正、追撃閾値、地形補正、反撃有無（seed固定）。

## フェーズ2：UI統合（screens/widgets）
- 模擬戦を `/pkg/game.Forecast/Resolve` 呼び出しへ置換。
- 予測UI（命中/与被ダメ/必殺/相性/地形）を表示。
- キーボード操作の導入（←↑→↓/Z/X/Tab/Space）。

## フェーズ3：AI・セーブ/ロード
- AI（1手先評価関数）: `score = E[dmg_to_enemy] - E[dmg_to_self] + kill_bonus + terrain - death_risk`。
- セーブ: 戦闘スナップショット（seed含む）を `artifacts/save_*.json`。ロードで再現。

## フェーズ4：検証と演出
- ヘッドレス `tools/simcli`、1フレーム `tools/render`、SSIMチェック、`make verify`。
- 演出: HPバーTween、被弾フラッシュ、ログ強調、色分け（有利/劣勢）。

## 受け入れ基準（要約）
- タイトル→戦闘→勝利/敗北の操作完結。
- 予測と結果が±1ダメ以内で一致（乱数除く）。
- 三すくみ命中±10/威力±2が反映。ZOCや回避地形の効果が可視。
- セーブ/ロードで再現。`make verify` が 0 exit。

## 当面の着手範囲（この後実装開始）
- `/pkg/game` のスケルトン（型・`Forecast` の最小実装）とユニットテスト雛形。
- UIからの呼び出し置換は次のコミットで段階導入。

---

## 実装状況（2025-09-27 現在）

フェーズ0：構成整備
- `/pkg/game` 新設済み。UI 非依存のロジック層として運用中。
- 追加予定データ（三すくみ表・地形表の JSON）は未導入（コード内の定数で暫定対応）。

フェーズ1：コアロジック（/pkg/game）
- 型: `Unit/Stats/Weapon/Terrain` 実装済。`Terrain{Avoid,Def,Hit}` を使用。
- 三すくみ: `Sword > Axe > Lance > Sword`（命中±10/威力±2）を実装。
- 予測: `ForecastAt(att,def,attTile,defTile)` 実装（命中表示値・与ダメ・必殺%）。
- 解決: `ResolveRoundAt(att,def,attTile,defTile,rng)` 実装（2RN、最小ダメ=1、クリティカル=武器Crit+床(skl/2)-lck、0..100にクランプ）。
- 内訳: `ForecastAtExplain` と `ForecastBreakdown` を追加（UIに命中/ダメの内訳を表示可能）。
- 射程: `Weapon.RMin/RMax` を参照。反撃可否はUI層で距離1固定チェック（ロジック側は未強制）。
- テスト: `pkg/game/forecast_test.go`, `resolve_test.go` 追加。`make test`（pkgのみ）通過。

フェーズ2：UI統合（screens/widgets）
- バトルプレビュー: `/pkg/game.ForecastAt` を用いて「命中/与ダメ/必殺」を両者表示。
- 反撃可否: 射程に基づく「(反撃不可)」表示対応。
- 相性表示: 有利=青/不利=橙/中立=灰のラベルを表示（内訳は±値で可視化）。
- 地形: 左右の地形（平地/森/砦）をキーで切替（攻: 1/2/3、防: Shift+1/2/3）。プレビューに地形名と補正（回避/防御/命中）を表示。
- 解決: 「戦闘開始」で `ResolveRoundAt` を呼び出し（攻撃→射程可反撃）。
- 操作: 一覧/詳細/戦闘で基本キーバインド導入（←↑→↓/Z/Enter/X/Esc/Tab）。

フェーズ3：AI・セーブ/ロード
- AI: 未実装。
- セーブ: 戦闘スナップショットの保存/復元は未実装。従来のユーザJSONへのHP/耐久反映は維持。

フェーズ4：検証と演出
- 検証: `make mcp`（lint＋vet＋ビルド）整備。サンドボックス/ネットワーク制限下では pkg ビルドへフォールバックさせ完走。
- 演出: 既存HPバーなどの最低限のみ。Tween/被弾フラッシュは未着手。

受け入れ基準の進捗
- 予測と結果の整合: 同一ロジック（ForecastAt/ResolveRoundAt）に統一済。±1以内の一致（乱数除く）を満たす設計。
- 相性UI: 命中±10/威力±2を適用・表示。
- ZOC/経路: 未実装。
- セーブ/ロード再現: 未実装。
- make verify: 未整備（今は `make test`/`make mcp` のみ）。
- 敵AI: 未実装。

既知の課題/ToDo（次フェーズ）
- 地形/三すくみの外部化（`db/master` の定義ファイル追加）。
- マップ/移動/射程2以上の距離計算（`ComputeReachable/ComputeAttackables`）。
- 反撃不可時のボタン無効化/ガイド表示強化。
- セーブ/ロード（seed含む）の JSON 設計とヘッドレス `tools/simcli` 整備。
- `make verify` パイプラインの導入（lint/test/sim/render/SSIM/hmac）。
