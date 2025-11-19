# 01_design — ItemsRepo 設計

## 目的
- `usecase.App.ItemsTable()` の直読みを廃止し、Repo/キャッシュ経由へ統一するためのIF/責務を定義する。

## 方針
- 追加: `internal/repo/items.go`
- IF: `type ItemsRepo interface { Table() *model.ItemDefTable; Reload() error }`
- 実装: JSON 読込→`ItemDefTable` を保持（`WeaponsRepo` と同等のキャッシュ戦略）。
- ライフサイクル: `App` 起動時に初期化、`ReloadData()` で再読込。

## ToDo
- [x] IF 定義案を作成（`internal/repo/items.go`）。
- [x] キャッシュ戦略（常駐/遅延/TTL）の比較と決定（今回は常駐）。
- [x] `docs/KNOWLEDGE/data/db-notes.md` へ参照経路の更新メモ。
