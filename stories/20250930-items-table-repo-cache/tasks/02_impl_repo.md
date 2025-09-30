# 02_impl_repo — ItemsRepo 実装

## 実装項目
- `internal/repo/items.go` を追加し、`ItemsRepo` 実装を提供。
- `usecase.App` に `Items ItemsRepo` を追加、コンストラクタ/初期化で注入。
- `ReloadData()` に `Items.Reload()` を統合。

## 参考/類似
- `internal/repo/weapons.go` のキャッシュ/Reload 実装。

## 検証
- [ ] `go build ./...` が通る。
- [ ] `make mcp` がグリーン。

