# 03 — ローダ実装とDI配線

ステータス: [完了]

## 目的
- JSON メトリクスを読み込むローダを実装し、UI レイアウトで利用可能にする。

## 実装方針
- パッケージ: `internal/config/uimetrics`（I/O と構造体定義）。
- 公開 API 例:
  - `func Load(path string) (Metrics, error)`
  - `func Default() Metrics`（ビルトイン既定値）
- 依存注入:
  - `internal/game/service/ui` に `func ApplyMetrics(Metrics)` を用意し、既存の定数を上書き。
  - 起動時（`cmd/ui_sample/main.go` または Game 初期化）で `Load(DefaultUIMetricsPath)` → `ApplyMetrics`。

## 変更予定ファイル
- `internal/config/config.go`（`DefaultUIMetricsPath`/`DefaultUserUIMetricsPath` 追記 済）
- `internal/config/uimetrics/*.go`（新規 済）
- `internal/game/service/ui/layout.go`（const→var に変更 済）
- `internal/game/service/ui/apply.go`（`ApplyMetrics` 追加 済）
- `internal/game/app/core.go`（ロード/適用を組込み 済）

## 成果物
- ローダ実装済（`internal/config/uimetrics/metrics.go`）
- 起動時ロード/適用済（`internal/game/app/core.go`）
- 既定 JSON 追加（`db/master/mst_ui_metrics.json`）

## 進捗ログ
- 2025-09-29: 雛形作成
- 2025-09-29: 実装/配線完了、`make mcp` グリーン（lintは任意）
