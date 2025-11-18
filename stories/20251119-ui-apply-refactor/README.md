# 20251119-ui-apply-refactor — `internal/game/service/ui/apply.go` の分割実装

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 01:36:42 +0900
元ストーリー: 20251117-ui-apply-split（調査/設計）

## 目的・背景
- 調査ストーリーで決定した案A（`apply.go` を API ラッパに残し、ドメイン別 `apply_*.go` へロジックを切り出す）を実装する。
- `apply.go` に集約された 70+ フィールドの代入ロジックを整理し、責務ごとのファイルに分割してコンフリクトとレビュー負荷を軽減する。
- メトリクス適用処理をヘルパ関数/構造体で共通化し、将来のゼロ許容など仕様変更に備える。

## スコープ（成果）
- `apply.go` に公開 API（`Metrics`/`DefaultMetrics`/`ApplyMetrics`）だけを残し、実処理は `apply_helpers.go`、`apply_list.go`、`apply_status.go`、`apply_sim.go`、`apply_popup.go`、`apply_widgets.go` へ分割されている。
- 新設ヘルパ `metricsTargets` や `assignPositive` 等で非ゼロ判定・スライスコピーを共通化し、テストで差し替え可能な構造になっている。
- `apply_test.go` にドメイン別テストが追加され、部分適用の回帰が担保されている。
- ドキュメント（`docs/API.md`, `docs/ARCHITECTURE.md`）が新構成に追随し、`make mcp` が通っている。

## 受け入れ基準（Definition of Done）
- [x] `ApplyMetrics` 本体がセクション別の内部関数呼び出しのみになっている。
- [x] `metricsTargets` を介した単体テストが追加され、List/Status/Sim/Popup/Widgets ごとに適用を検証できる。
- [x] ドキュメントの参照が `apply.go` 単体ではなく分割構成に更新されている。
- [x] `make mcp` が成功し、UI ホットリロード経路（`internal/game/app/game.go`, `core.go`）で動作確認済み。

## 工程（サブタスク）
- [x] 01_ヘルパ導入と準備 — `apply_helpers.go` と `metricsTargets` の追加
- [x] 02_ドメイン別ファイル分割 — List/Status/Sim/Popup/Widgets の移行
- [x] 03_テスト＆ドキュメント／CI 検証 — `apply_test.go` 拡張 + docs 更新 + `make mcp`

## 進捗・決定事項（ログ）
- 2025-11-19 01:36:42 +0900: ストーリー起票。Backlog `[P1] 2025-11-19: ApplyMetrics 分割実装` から昇格。
- 2025-11-19 10:05:00 +0900: Task01/02 を連続実施し、`metricsTargets` + `apply_*.go` へ分割完了。`go test ./internal/game/service/ui` で回帰なしを確認。
- 2025-11-19 10:45:00 +0900: Task03 で `apply_test.go` 拡張・`docs/API.md`/`docs/ARCHITECTURE.md` 更新・`make mcp` まで完了し、ストーリーを DoD 充足としてクローズ。

## リスク・懸念
- `metricsTargets` へのフィールド追加漏れでビルドエラーにならずランタイム不具合となる可能性 → テストで各ドメインのフィールドを網羅。
- 並行作業とのコンフリクト。分割後の大規模差分になるため、段階的コミットと早めの PR を推奨。

## 関連
- Backlog: stories/BACKLOG.md
- Docs: `docs/API.md`, `docs/ARCHITECTURE.md`
- 参照ストーリー: `stories/20251117-ui-apply-split`
