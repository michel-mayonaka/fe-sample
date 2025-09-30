# 03_replace_calls — 呼び出し置換

## 目的
- 既存の `usecase.App.ItemsTable()` 直読みを Repo 経由に統一し、副作用のない参照経路にする。

## 対象候補
- `internal/usecase/data.go` の `ItemsTable()` 実装。
- `internal/game/scenes/inventory/popup_item.go`（Provider 経由を維持しつつ、背後がRepo化されていることを確認）。

## 手順
- [ ] `usecase.App.ItemsTable()` を `a.Items.Table()` に切替。
- [ ] `rg -n 'LoadItemsJSON|ItemsTable\(\)'` で残存確認。
- [ ] 主要画面で表示/動作確認（一覧/装備選択）。

