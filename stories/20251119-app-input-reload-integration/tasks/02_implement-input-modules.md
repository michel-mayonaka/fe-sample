# 02_実装 — input/reload モジュールの追加

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-19 02:06:50 +0900

## 目的
- Ebiten 入力ソース／抽象入力アダプタ／グローバルアクション検出を `internal/app/input` にまとめ、`Game` から直接キーを読む処理を排除する。

## 完了条件（DoD）
- [ ] `internal/app/input` 下に layout/source/global 等のサブパッケージ（またはファイル）が追加されている。
- [ ] `internal/game/app/game.go` の `updateGlobalToggles` から Ebiten 依存が取り除かれ、app 層のイベント経由でリロード/ヘルプを制御できる。
- [ ] `make mcp` が通り、グローバル操作（ヘルプ表示、Backspace リロード）が従来通り動作する。

## 進捗ログ
- 2025-11-19 02:06:50 +0900: タスク登録。
