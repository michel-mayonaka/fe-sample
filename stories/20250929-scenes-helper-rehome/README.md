# [story] scenes/helper の機能別分割と配置見直し

- ステータス: 進行中
- 日付: 2025-09-29
- 参照: docs/NAMING.md, docs/ARCHITECTURE.md

## 背景
`scenes/helper` は汎用名かつ多責務化しやすく、規約上も避けるべき。役割単位に分割して再配置する。

## 目的
- 具体的な役割名・層に基づく配置へ再編し、再利用性/テスト容易性を高める。

## スコープ
- UI 描画/入力補助 → `internal/game/ui/draw`, `internal/game/ui/input` 等へ。
- 横断サービス（アニメ/遷移等） → `internal/game/service/<feature>` へ。
- 幾何/数学など純粋ロジック → ドメインなら `pkg/game/geom`、アプリ固有なら `internal/math/geom` へ。
- 参照更新、不要コードの削除。

## 非スコープ
- 新機能追加やデザイン改変。

## 成果物 / DoD
- `scenes/helper` への依存が解消され、具体パッケージへ置換されている。
- `make mcp` 成功、動作回帰確認が完了。
- docs/NAMING.md にパッケージ命名例を追記（`helper` 禁止の明文化）。

## 影響範囲
- `scenes/*`（呼び出し差し替え）
- `internal/game/ui/*`, `internal/game/service/*`, `pkg/game/*`

## リスクと対策
- リスク: 隠れた副作用/状態共有の崩れ。
  - 対策: ミニマムなパイロット（1〜2箇所）から段階適用、ユニットテスト追加。

## 計測/検証
- `rg scenes/helper` が 0 件になること。重複/類似コードの削減。

## 次アクション（タスク化方針）
- 01_中身の棚卸し（分類表作成）
- 02_UI系の移設
- 03_サービス系の移設
- 04_純粋ロジックの移設
- 05_参照掃除＋Lint/Doc 更新
