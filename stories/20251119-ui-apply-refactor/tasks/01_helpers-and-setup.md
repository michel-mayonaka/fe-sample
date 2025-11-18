# 01_ヘルパ導入と準備 — `apply_helpers.go` / `metricsTargets`

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-19 01:36:42 +0900

## 目的
- ドメイン分割の前提として、非ゼロチェックやスライスコピーを共通化するヘルパと、グローバル変数群へアクセスする `metricsTargets` を整備する。

## 完了条件（DoD）
- [ ] `apply_helpers.go`（仮）に `assignPositive`, `assignSlice`, `copyInts`, `metricsTargets` などが実装されている。
- [ ] `ApplyMetrics` 内で `metricsTargets` を生成し、既存ロジックはヘルパ経由でアクセスする形に置き換わっている（リファクタ第一段階、機能変更なし）。
- [ ] 既存テスト（`apply_test.go`）が通り、ビルドが壊れていない。

## 作業手順（概略）
1. `layout.go` のグローバル変数を `metricsTargets` 構造体へマッピングする（ポインタもしくは setter 関数）。
2. スライス/配列コピーの共通関数を作成し、既存の `append([]int(nil), ...)` を段階的に置き換え。
3. `ApplyMetrics` のトップを `targets := newMetricsTargets()` のように書き換え、以降のリファクタで再利用できる状態にする。

## 進捗ログ
- 2025-11-19 01:36:42 +0900: タスク作成。

## 依存／ブロッカー
- なし（単独で着手可）。

## 成果物リンク
- PR/コミット: （後続で記載）
