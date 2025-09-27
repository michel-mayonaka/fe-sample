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
