# 20250930-items-table-repo-cache — ItemsTable 取得の Repo 化・キャッシュ一貫化

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-09-30 00:00:00 +0900

## 目的・背景
- `usecase.App.ItemsTable()` の都度JSON直読みを廃止し、Repo/キャッシュ経由に統一して性能とオフライン挙動を安定化する。
- 参照経路の分散を解消し、`TableProvider.ItemsTable()` からの一貫提供を徹底する。

## スコープ（成果）
- `internal/repo/items.go`（仮）に `ItemsRepo` を追加（`Table()/Reload()` を提供）。
- `usecase.App` に `Items ItemsRepo` を保持し、`ItemsTable()` を Repo 経由へ切替。
- 既存呼び出し箇所の整合（UI/Scene は `gdata.Provider().ItemsTable()` を利用）。
- 最小のドキュメント更新（`docs/DB_NOTES.md`/`docs/API.md`）。

## 受け入れ基準（Definition of Done）
- [x] `internal/usecase/data.go` に JSON 直読みが残っていない（`rg -n 'LoadItemsJSON'` が 0）。
- [x] `internal/repo/items.go` 相当のキャッシュ実装があり、`Reload()` で再読込可能。
- [x] 主要経路の動作確認（在庫一覧/装備ポップアップ）。
- [x] `make mcp` がグリーン。

## 工程（サブタスク）
- [x] 設計: Repo/キャッシュ方針の決定 — `tasks/01_design.md`
- [x] 実装: ItemsRepo を追加し App へ配線 — `tasks/02_impl_repo.md`
- [x] 置換: 呼び出し箇所の更新と確認 — `tasks/03_replace_calls.md`
- [x] 文書: API/DB_NOTES 最小更新 — `tasks/04_docs.md`

## 計画（目安）
- 見積: 2–3 時間（1セッション）
- マイルストン: M1 設計 / M2 実装 / M3 置換+検証

## 進捗・決定事項（ログ）
- 2025-09-30 00:00:00 +0900: ストーリー作成（BACKLOG昇格）
- 2025-10-01 05:11:00 +0900: ステータスを[準備完了]へ更新（レビュー完了）
- 2025-11-16 00:00:00 +0900: ItemsRepo 実装・App 配線・ItemsTable の Repo 化、および docs/API.md・docs/DB_NOTES.md を更新（AIエージェントによる実装）。
- 2025-11-16 00:00:00 +0900: `go test ./internal/usecase` まで確認済。`go test ./...` では既存の sim/levelup テストが失敗しており、本ストーリーでは未対応。

## リスク・懸念
- キャッシュのライフサイクル（Reload/ホットリロード）設計誤り。
- 今後の DB（SQLite）移行時のIF差異。

## 関連
- Discovery: `stories/discovery/consumed/2025-09-30-migrated-04.md`
- 実装候補: `internal/usecase/data.go`, `internal/repo/`, `internal/game/data/provider.go`
- Docs: `docs/DB_NOTES.md`, `docs/API.md`

- 2025-11-17 01:23:03 +0900: アーカイブ（finish へ移動）
