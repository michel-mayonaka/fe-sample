# pkg/game インポート実態棚卸しメモ

作成日: 2025-11-17
ストーリー: stories/20250929-pkg-game-visibility/README.md

## 1. 調査方法

- `rg "ui_sample/pkg/game" -n` をレポジトリルートで実行し、`pkg/game` 本体とサブパッケージ（`geom`/`input`）を含む import を列挙した。
- 該当ファイルを確認し、「どの層がどのような用途で `pkg/game` に依存しているか」を分類した。
- 外部モジュール（このリポジトリ外）については Go モジュール依存情報が無いため、ここでは「存在しない」とみなしている（現状は UI サンプル単体モジュール）。

## 2. インポート元パッケージ一覧（同一モジュール内）

### 2-1. 戦闘ロジック本体 (`ui_sample/pkg/game`)

- `internal/usecase/battle.go`
  - import: `gcore "ui_sample/pkg/game"`
  - 用途:
    - `RunBattleRound` 内で `adapter.UIToGame` により UI ユニットを `gcore.Unit` に変換し、`ResolveRoundAt` や `DoubleAdvantage` を用いて 1 ラウンド戦闘を解決している。
    - 地形指定には `gcore.Terrain` を使用し、結果は UI ユニットとユーザセーブに反映される。

- `internal/usecase/battle_test.go`, `internal/usecase/run_battle_more_test.go`
  - import: `gcore "ui_sample/pkg/game"`
  - 用途:
    - テストから直接 `gcore.Unit` や `gcore.Terrain` を組み立て、戦闘ユースケースの挙動を検証している。

- `internal/game/scenes/sim/*.go`
  - 代表: `internal/game/scenes/sim/sim.go`, `engine.go`, `logic.go`
  - import: `gcore "ui_sample/pkg/game"`
  - 用途:
    - シミュレーション画面で、戦闘結果やログを表示するために `gcore` の型と関数を利用。

- `internal/game/ui/draw/sim_battle.go`
  - import: `gcore "ui_sample/pkg/game"`
  - 用途:
    - `DrawBattleWithTerrain` 内で `ForecastAt`/`ForecastAtExplain`/`ForecastBreakdown` を用い、左右のユニットの命中・与ダメ・必殺率のプレビューと内訳を描画している。

- `internal/game/scenes/ports_battle.go`
  - import: `gcore "ui_sample/pkg/game"`
  - 用途:
    - Scene 側の Port 定義で `gcore.Terrain` や戦闘関連の引数/戻り値を使用し、UI 層と usecase 層の境界として戦闘ドメイン型を露出している。

- `internal/adapter/unit.go`
  - import: `gcore "ui_sample/pkg/game"`
  - 用途:
    - `UIToGame` で `uicore.Unit` → `gcore.Unit` への変換に `gcore.Stats`/`gcore.Weapon` を利用。
    - `AttackSpeedOf` で UI ユニットの攻撃速度計算に `gcore.AttackSpeed` を利用。

### 2-2. 幾何ユーティリティ (`ui_sample/pkg/game/geom`)

- `internal/game/scenes/inventory/inventory.go`, `popup_item.go`, `popup_weapon.go`
  - import: `"ui_sample/pkg/game/geom"`
  - 用途:
    - マウス座標とボタン矩形のヒットテストなど、UI 上の矩形判定に `geom.RectContains` を使用。

- `internal/game/scenes/status/status.go`, `internal/game/scenes/character_list/character_list.go`
  - import: `"ui_sample/pkg/game/geom"`
  - 用途:
    - ステータス画面やキャラクター一覧で、行・スロットなどの矩形ヒットテストに `RectContains` を利用。

### 2-3. 入力ドメイン (`ui_sample/pkg/game/input`)

- `internal/game/provider/input/ebiten/source.go`
  - import: `ginput "ui_sample/pkg/game/input"`
  - 用途:
    - Ebiten のキー/マウス状態を `ginput.ControlState`/`ginput.Event` に投影するソース実装。
    - キーコードやマウスボタンを `ginput.Action*` にマッピングし、ドメイン側の抽象入力として扱う。

- `internal/game/app/game.go`, `internal/game/app/core.go`
  - import: `ginput "ui_sample/pkg/game/input"`
  - 用途:
    - アプリケーションゲームの初期化時にドメイン入力レイアウトやラッパーを構成し、`Ctx` 経由で Scene へ渡す。

- `internal/game/ui/input/types.go`
  - import: `ginput "ui_sample/pkg/game/input"`
  - 用途:
    - 旧 UI サービス層との互換のために、ドメイン側の `Action`/`Reader` をラップするアダプタ (`DomainAdapter`) を提供。

## 3. 外部モジュールからの利用有無

- `rg "ui_sample/pkg/game"` の結果から、import 元はすべて同一モジュール（このリポジトリ内）の `internal/...` に限られている。
- go.mod 上も `module ui_sample` のみであり、他モジュールから `ui_sample/pkg/game` を参照している事実は現時点では確認できない。
- よって、**現時点の可視性変更は「将来の外部再利用可能性」にのみ影響し、既存コード（このリポジトリ内）の動作には影響しない** とみなせる。

## 4. 可視性を internal に変更した場合の影響概算

ここでは「`pkg/game` を `internal/gamecore` のような internal 配下に移動し、パスを変更する」ケースを想定する（実際のリネームは非スコープ）。

- ビルド観点
  - Go の `internal` ルール上、同一モジュール内の `internal/...` からの import は許可されるため、**パス修正さえ行えば** `internal/usecase` や `internal/game/...` は問題なくビルド可能。
  - 一方で、モジュール外（将来の別リポジトリやツール）からは import 不可となる。

- 修正が必要になるファイル数（概算）
  - `rg "ui_sample/pkg/game" -n` でヒットした Go ファイルが 10 数件（本ストーリー時点）。
  - これらはすべて import パスを書き換える必要がある（例: `ui_sample/pkg/game` → `ui_sample/internal/gamecore` など）。
  - `pkg/game` 配下のテスト (`pkg/game/*.go` のうち `_test.go`) も、パッケージ名や import を合わせて修正が必要。

- 互換性/テストへの影響
  - モジュール外からの再利用は不可能になるため、将来 CLI やサーバから戦闘ロジックを直接呼びたい場合は「別の公開パッケージ」を新設する必要が出てくる。
  - 逆に、モジュール内のテスト (`make mcp` に含まれる `go test ./pkg/... ./internal/...`) は、import パスさえ揃えれば従来どおり動作する見込み。

## 5. 小結

- `pkg/game` は現状、**このリポジトリ内の UI 層/ユースケース層からのみ参照されており、外部モジュールからの利用実績は無い**。
- 可視性を `internal` に変更しても、同一モジュール内に限れば import は継続可能であり、「既存コードのビルドエラー」は機械的なパス修正で解消できるレベルと見積もられる。
- 一方で、`internal` 化すると将来の外部再利用（CLI・サーバ・別ツール）が制約されるため、「今後どこまでライブラリとして扱いたいか」が判断の主な論点となる。

このメモは 02_インポート実態の棚卸し の成果物とし、後続タスク（比較表・スパイク）での判断材料とする。

