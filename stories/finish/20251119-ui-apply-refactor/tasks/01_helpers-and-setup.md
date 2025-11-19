# 01_ヘルパ導入と準備 — `apply_helpers.go` / `metricsTargets`

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 01:36:42 +0900

## 目的
- ドメイン分割の前提として、非ゼロチェックやスライスコピーを共通化するヘルパと、グローバル変数群へアクセスする `metricsTargets` を整備する。

## 完了条件（DoD）
- [x] `apply_helpers.go`（仮）に `assignPositive`, `assignSlice`, `copyInts`, `metricsTargets` などが実装されている。
- [x] `ApplyMetrics` 内で `metricsTargets` を生成し、既存ロジックはヘルパ経由でアクセスする形に置き換わっている（リファクタ第一段階、機能変更なし）。
- [x] 既存テスト（`apply_test.go`）が通り、ビルドが壊れていない。

## 作業手順（概略）
1. `layout.go` のグローバル変数を `metricsTargets` 構造体へマッピングする（ポインタもしくは setter 関数）。
2. スライス/配列コピーの共通関数を作成し、既存の `append([]int(nil), ...)` を段階的に置き換え。
3. `ApplyMetrics` のトップを `targets := newMetricsTargets()` のように書き換え、以降のリファクタで再利用できる状態にする。

## 進捗ログ
- 2025-11-19 01:36:42 +0900: タスク作成。
- 2025-11-19 09:35:00 +0900: `metricsTargets` 構造体と `assignPositive`/`assignSlice`/`copyInts` を `apply_helpers.go` に実装。
- 2025-11-19 10:05:00 +0900: `ApplyMetrics` が `newMetricsTargets()`→`applyList/...` を呼ぶ形に整理し、`go test ./internal/game/service/ui` でリグレッションなしを確認。

## 依存／ブロッカー
- なし（単独で着手可）。

## 成果物リンク
- コード: `internal/game/service/ui/apply_helpers.go`, `internal/game/service/ui/apply.go`
