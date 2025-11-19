# 02_分割方針の検討 — ファイル構成と責務再配置案

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-17 02:13:34 +0900

## 目的
- 現状調査の結果を踏まえ、`apply.go` をどのような単位（List/Status/Sim/Popup/Widgets 等）でファイル分割するか、および各ファイルの責務と API を定義する。

## 完了条件（DoD）
- [x] 候補となるファイル構成（例: `apply_list.go`/`apply_status.go`/`apply_sim.go`/`apply_popup.go`/`apply_widgets.go`）が提示されている。
- [x] 各ファイルが担当する責務と、公開/内部関数の境界がラフに決められている。
- [x] 外部 API を極力変更しないための方針（ラッパ関数/型の配置など）が記述されている。

## 作業手順（概略）
- `01_現状調査` の結果を元に、自然な責務の切れ目（画面単位/メトリクス種別など）を検討する。
- 命名規約（docs/KNOWLEDGE/engineering/naming.md）とアーキテクチャ方針（docs/architecture/README.md）を確認し、それに沿ったファイル名・レイヤ分割を考える。
- 複数案がある場合は Pros/Cons を列挙し、推奨案を1つ選ぶ。

## 進捗ログ
- 2025-11-17 02:13:34 +0900: タスク作成。
- 2025-11-19 00:40:33 +0900: 現状調査の結果を反映しつつ分割案ドラフトを作成。
- 2025-11-19 00:48:12 +0900: ドメイン別 `apply_*.go` 案と API 方針を確定し、DoD を満たしたため完了。

## 依存／ブロッカー
- `01_現状調査` の完了。

## 成果物リンク
- 分割方針メモ: stories/20251117-ui-apply-split/tasks/02_design-split-plan.md

## 分割方針メモ

### ゴール
- `ApplyMetrics`/`Metrics` の公開 API は維持し、呼び出し側（`internal/game/app` 等）の変更を不要にする。
- ドメイン単位（List/Status/Sim/Popup/Widgets）で責務を切り出し、将来の追加/改修を局所化する。
- 条件ロジック（非ゼロチェック、スライスコピー）を共通化し、重複による漏れを防ぐ。

### 候補構成
| 案 | 概要 | Pros | Cons |
| --- | --- | --- | --- |
| A: ドメイン別 `apply_*.go` | `apply.go` は API ラッパーのみ残し、`apply_list.go` などに `func applyList(dst *metricsTargets, src Metrics)` を配置。 | 既存ファイルを段階的に縮小でき、責務単位でレビューしやすい。`go test` への影響も軽微。 | 小さなファイルが5+個に増え、`ApplyMetrics` 内からの呼び順管理が必要。 |
| B: サブパッケージ化 | `internal/game/service/ui/apply` 配下にサブパッケージを作り、`uicore.ApplyMetrics` がそこへ委譲。 | メトリクス適用コードをパッケージとして独立させられる。 | `uicore` 直下の公開 API を動かす必要があり、`docs/SPECS/reference/api.md` などの参照修正が大きくなる。 |

→ **推奨: 案A**（現行 API を保ったままファイル分割のみ実施）。

### 推奨構成（案A詳細）
- `apply.go`: 公開 struct (`Metrics`) と API (`DefaultMetrics`, `ApplyMetrics`) を保持。`ApplyMetrics` 内ではセクション別関数へ委譲し、グローバルの import 影響を最小化。
- `apply_helpers.go`: 共通ヘルパ `func copyInts(src []int) []int`、`func assignPositive(dst *int, val int)` などをまとめる。将来的にゼロ許容を再検討する際もここを変更すればよい。
- `apply_list.go`: List + Line 領域を担当。`applyList(m Metrics)` の内部でリスト系と行間をまとめて処理。
- `apply_status.go`: Status 領域の 21 フィールドを担当。
- `apply_sim.go`: Sim 基本ボタン、Terrain、Preview を 1 ファイルに集約。
- `apply_popup.go`: Popup 関連9フィールド。
- `apply_widgets.go`: Widgets ボタン群。`widgets` ディレクトリ配下の利用が明確なので分離価値が高い。

### 公開/内部境界
- 外部向けは現状どおり:
  - `type Metrics struct { ... }`
  - `func DefaultMetrics() Metrics`
  - `func ApplyMetrics(m Metrics)`
- 内部向け:
  - `type metricsTargets struct`（`layout.go` の変数セットをまとめたハンドル。`*metricsTargets` を各 apply 関数へ渡し、テストで差し替え可能にする案）
  - `func applyList(t *metricsTargets, m Metrics)`
  - `func applyStatus(t *metricsTargets, m Metrics)` 等
  - `func applyPopup(...)`, `func applyWidgets(...)`
  - ヘルパ `func assignPositive(dst *int, val int)`、`func assignSlice(dst *[]int, src []int)`

※ `metricsTargets` は `layout.go` のグローバル変数に直接アクセスする単純構造体（フィールドはポインタまたは setter 関数）。これによりユニットテストでモックを渡しやすくなる。

### Pros/Cons（推奨案に対する詳細）
- Pros
  - メトリクス追加時は該当ドメインファイルのみ修正すればよく、コンフリクトを減らせる。
  - `metricsTargets` を挟むことで後続の設定バックエンド（例: JSON→構造体→別ランタイム）にも対応しやすい。
  - ヘルパ群を共有でき、非ゼロ判定の仕様変更も一箇所で済む。
- Cons
  - ファイル間で `metricsTargets` を共有するため、構造体の更新忘れがあるとビルドエラーになりにくい（テストでの検証が必要）。
  - 小さい変更でも複数ファイルに分散するため、`go test` の観点では差分追跡が煩雑になる可能性。

### 実装ステップ（案）
1. `apply.go` に `metricsTargets` とヘルパ関数を追加し、既存ロジックを段階的に移行。初期コミットでは機能変更なしを担保。
2. ドメイン別に `apply_*.go` を新規作成。`ApplyMetrics` 内の該当部分を `applyList(&targets, m)` などに置換。
3. `apply_test.go` を拡張し、`metricsTargets` を使った単体テストを追加（例: List 部分のみ差し替えて動作確認）。
4. `docs/SPECS/reference/api.md`/`docs/architecture/README.md` のパスを `apply.go` → `apply_*.go` へ更新（必要に応じて「内部的には分割」と注記）。

### 保留・追加検討事項
- 非ゼロ判定仕様を見直すか（例: `*int` を使って「0でも適用」の意思表示を可能にする）。現ストーリーでは方針のみ記載し、実装ストーリーで扱う。
- `metricsTargets` を導入するか、または単純に `applyList(m Metrics)` が直接グローバルへ代入する方式にするか（テスト容易性 vs シンプルさのトレードオフ）。今回はテスト容易性優先のため、`metricsTargets` 案を推奨。
