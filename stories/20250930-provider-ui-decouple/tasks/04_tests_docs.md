# 04_tests_docs — テスト/文書

## 目的
- 回帰を防ぎ、設計ルールをドキュメントに定着させる。

## テスト
- [x] 主要ユースケースの単体/結合テストが通る（`make mcp`）。
- [x] Adapter のユニット変換に最小テストを追加（`unit_from_user_test.go`）。
- [x] ステータス画面の「＜ 一覧へ」回帰テストを追加（`status_test.go`）。
- [x] 在庫画面の戻るボタン回帰テストを追加（`inventory_test.go`）。
- [x] 模擬戦画面の Cancel で戻る回帰テストを追加（`sim_test.go`）。
- [x] 模擬戦画面の「自動実行」/「＜ 一覧へ」ボタンクリック回帰テストを追加。

## 文書
- [x] `docs/ARCHITECTURE.md`/`docs/API.md` の差分更新。
