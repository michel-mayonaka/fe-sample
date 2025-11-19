# 01_現状調査 — apply.go の責務と依存関係の棚卸し

ステータス: [進行中]
担当: @tkg-engineer
開始: 2025-11-17 02:13:34 +0900

## 目的
- `internal/game/service/ui/apply.go` が現在担っている責務と依存関係（呼び出し元/呼び出し先/使用している型など）を整理し、分割時に壊してはいけない前提条件を明らかにする。

## 完了条件（DoD）
- [x] `apply.go` 内の主要な関数/ロジックが、どの画面（List/Status/Sim/Popup/Widgets 等）に対応しているかが一覧化されている。
- [x] 依存しているパッケージや型（例: メトリクス定義、view-model、layout など）が整理されている。
- [x] 分割時に考慮すべき制約（循環依存を避けるべき箇所など）がメモとして残っている。

## 作業手順（概略）
- `internal/game/service/ui/apply.go` の関数一覧とコメントを確認し、役割ごとにグルーピングする。
- 呼び出し元（scenes/ui など）と、呼び出し先/依存パッケージを追って、依存関係の概要を図や箇条書きでまとめる。
- 分割にあたって注意が必要なポイント（共通処理/ヘルパ関数など）を洗い出す。

## 進捗ログ
- 2025-11-17 02:13:34 +0900: タスク作成。
- 2025-11-19 00:25:01 +0900: `apply.go` の棚卸しを実施し、現状メモを追記。

## 依存／ブロッカー
- `internal/game/service/ui` 配下のファイル全体へのアクセス。

## 成果物リンク
- 現状調査メモ: stories/20251117-ui-apply-split/tasks/01_audit-apply-go.md

## 現状調査メモ

### ファイル構成と責務
- `Metrics` 構造体に List/Line/Status/Sim/Popup/Widgets の各グループを内包し、`config/uimetrics.Metrics` からのデータ受け皿になっている（タグなし・Go 内部専用）。
- `DefaultMetrics` は `layout.go` のビルトイン値を `Metrics` へコピーし、スライスは `append([]int(nil), ...)` でディープコピーして初期状態のスナップショットを提供。外部利用は `apply_test.go` のみ。
- `ApplyMetrics` が唯一の公開エントリーポイントで、非ゼロ（スライスは非空）のフィールドだけを `layout.go` のグローバル変数へ書き戻す。0 や負値を上書きする手段が無く、部分適用前提の挙動になっている。
- ビルドタグ `//go:build !headless` により、UI を含まない headless ビルドでは本ファイル全体がコンパイル対象外。

### 呼び出し元／入力元
- ランタイム適用: `internal/game/app/core.go` と `internal/game/app/game.go` が `config/uimetrics.LoadOrDefault` → `uicore.Metrics` へコピー → `ApplyMetrics` を呼び出す経路（初期化時とホットリロード）。
- テスト: `internal/game/service/ui/apply_test.go` が既定値スナップショット＋部分適用のリグレッションをカバー。
- ドキュメント参照: `docs/SPECS/reference/api.md` / `docs/architecture/README.md` で「ApplyMetrics がビルトイン変数を上書きする」と明記されており、公開 API として扱われている。

### 出力先／依存先
- `layout.go` のグローバル変数群（約 70 フィールド）を直接更新。これらは `metrics.go` でスケール済みアクセサ `ListMarginPx()` などにラップされ、UI 実装全体から参照される。
- スライス系（ヘッダ列・行列）は `append` でコピーしており、呼び出し側と共有しない前提。

### ドメイン別の主な利用箇所
- **List + Line**: `internal/game/ui/draw/inventory_list.go`、`ui/layout/list.go`、`scenes/character_list` / `inventory` / `sim` などメインパネルの余白・行 gap を参照。`LineHMain/Small` は `ui/draw/sim_battle.go` ほか複数のテキスト描画で折り返し計算に使用。
- **Status**: `internal/game/ui/draw/status.go` と `ui/layout/status.go` がほぼ全フィールドを利用し、各ポートレート/ステータス列を配置。`scenes/status` がウィジェット状態と連動。
- **Sim（Terrain/Preview 含む）**: `ui/draw/sim_battle.go`, `ui/layout/sim.go`, `service/ui/widgets/terrain.go` がスタートボタン、地形ボタン、プレビューカードの配置を参照。`scenes/sim` の入力制御も `SimTitleYOffset` などを前提とする。
- **Popup**: `ui/layout/popup.go`, `ui/draw/popup_choose_unit.go`, `ui/draw/levelup_popup.go` がしきい値・セルサイズを利用し、`scenes/status` / `scenes/character_list` / `scenes/sim` がポップアップ状態を描画する。
- **Widgets**: `service/ui/widgets/buttons.go` が Back/LevelUp/ToBattle/SimBattle のサイズ・余白を用い、`scenes/status` や `scenes/inventory` が hover 判定に転用。

### 分割時に考慮すべき制約
1. API 互換性: `ApplyMetrics`/`Metrics` の公開シグネチャを残す必要があり、ファイル分割後も単一点から呼べる構成（例: `ApplyMetrics` → 内部の `applyList(m)` など）にする。
2. 非ゼロ条件: 現挙動では「0 を適用できない」ため、責務を分けても条件判定を統一しないと意図しないゼロクリアが発生しうる。必要なら別タスクで仕様見直し。
3. スライスコピー: `append([]int(nil), ...)` を各セクションで忘れると、ロード済み設定の可変参照を保持してしまう。共有ヘルパを定義する案あり。
4. ビルドタグ: 分割ファイルも `//go:build !headless` を付与し、headless ビルドでの欠落を防ぐ。
5. テスト: 現状テストは単一ファイルを前提（`DefaultMetrics` と `ApplyMetrics` 同居）。分割時はテストの import 循環を避けるため、テスト用の共通 `_test.go` から各セクションを exercise する形が望ましい。
6. ドキュメント更新: `docs/SPECS/reference/api.md` / `docs/architecture/README.md` に記載されている参照先パス（`ui/apply.go`）が変わる場合は追随が必要。

### 依存関係サマリ表
| ドメイン | 更新対象フィールド数 | 主利用ファイル例 | 補足 |
| --- | --- | --- | --- |
| List/Line | 22 | `ui/draw/inventory_list.go`, `ui/layout/list.go`, `scenes/character_list` | Inventory/Character 画面の共通レイアウト |
| Status | 21 | `ui/draw/status.go`, `ui/layout/status.go` | ステータス画面のみ、Widgets と位置共有 |
| Sim（含 Terrain/Preview） | 29 | `ui/draw/sim_battle.go`, `ui/layout/sim.go`, `service/ui/widgets/terrain.go` | プレビューやログのレイアウトで LineHSmall とも連動 |
| Popup | 9 | `ui/layout/popup.go`, `ui/draw/popup_choose_unit.go` | 列数しきい値で UI 分岐 |
| Widgets | 13 | `service/ui/widgets/buttons.go`, `scenes/status`, `scenes/inventory` | Back/LevelUp/戦闘開始ボタン配置 |
