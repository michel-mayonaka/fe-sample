# [task] 現状棚卸しと分類表作成

- ステータス: 進行中
- 目的: `scenes/helper` 系（現 `rect_helpers.go` を含む）の中身を洗い出し、UI系/サービス系/純ロジック系に分類して移設方針を確定する。

## 入力
- 対象コード: `internal/game/scenes/*.go`（特に `rect_helpers.go`）
- 規約: `docs/NAMING.md`, `docs/ARCHITECTURE.md`

## スコープ
- 関数/型の一覧化（定義/参照箇所、依存パッケージ、純粋性）。
- 分類: UI（描画/入力）, サービス（遷移/演出/状態管理）, 純ロジック（幾何/計算）。
- 移設先案の決定（パッケージ名・ファイル名・公開/非公開）。

## 非スコープ
- 実コードの移動・削除（後続タスク）。

## 手順
- 現状把握:
  - `rg -n "scenes/(helper|rect_helpers)" internal || true`
  - `rg -n "\bPointIn\b" internal`
- 依存確認: `gofmt -r` ではなくコードリーディングで副作用/参照の有無を確認。
- 分類表を本ファイル末尾に追記（関数名/責務/移設先/備考）。
- 命名方針を決定（例: `PointIn` → `geom.RectContains`）。

## DoD（完了条件）
- 分類表が合意済みで、後続タスクの対象/移設先が確定している。
- 影響範囲（参照一覧）が出揃っている。

## コマンド例
- `rg -n "scenes/(helper|rect_helpers)" internal`
- `rg -n "\bPointIn\b"`
- `make mcp`（影響確認用・現状は通ること）

---
### 分類表（追記用）

| 識別子 | 責務 | 現在地 | 提案先 | 備考 |
|---|---|---|---|---|
| `PointIn(px,py,x,y,w,h)` | 矩形内判定 | `internal/game/scenes/rect_helpers.go` | `pkg/game/geom`（`RectContains` 等） | 純ロジック。全シーンから利用。

