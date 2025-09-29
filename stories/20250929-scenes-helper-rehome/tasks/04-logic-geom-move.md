# [task] 純ロジック（幾何/計算）の移設（`pkg/game/geom`）

- ステータス: 完了
- 目的: シーンに依存しない純粋関数をドメイン層へ移し、テスト容易性と再利用性を高める。

## 入力
- タスク01の分類表（純ロジック項目）

## スコープ
- 新設: `pkg/game/geom`
- 既存: `internal/game/scenes/rect_helpers.go` の `PointIn` を移設・改名（候補: `RectContains(px,py,x,y,w,h int) bool`）。
- 呼び出し側の参照更新（`scenes.PointIn` → `geom.RectContains` 等）。

## 非スコープ
- 仕様変更（境界の扱いなどのロジック変更）。現仕様を維持。

## 手順
- パッケージ作成: `pkg/game/geom/doc.go`（責務を記述）。
- 実装追加: `rect.go` に `RectContains` を実装。
- 単体テスト: `rect_test.go`（境界含有/外側/負座標の最小ケース）。
- 置換: `rg -n "\bPointIn\b"` の呼び出しを差し替え。
- 旧定義削除: `internal/game/scenes/rect_helpers.go` を削除。
- `make mcp` 実行。

## DoD（完了条件）
- `pkg/game/geom` にテスト付きで実装が存在する。
- `rg PointIn` が 0 件。
- `make mcp` グリーン。

## コマンド例
- `mkdir -p pkg/game/geom`
- `rg -n "\bPointIn\b"`
- `go test ./... -race -cover`
