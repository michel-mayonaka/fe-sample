# 13_repo_sqlite_skeleton — SQLite リポジトリ骨組み

## 目的・背景
- 将来のDB移行（JSON→SQLite）に備え、`repo/sqlite` の骨組みを追加（未配線）。

## 作業項目（変更点）
- ディレクトリ `internal/repo/sqlite` 追加、`user.go`/`weapons.go`/`inventory.go` のIFスタブ。
- `go:build sqlite` などのビルドタグコメントを付与（未使用時に無害）。
- README に「将来の移行先」注記を追加。

## 完了条件
- ビルド成功（スタブは未使用だがエラーにならない）。

## 影響範囲
- internal/repo/sqlite/* のみ（参照なし）。

## 手順
1) スタブファイルを追加しIF満たすダミー実装を記述（`panic("not implemented")` 禁止）。
2) Lint で不要警告が出ないよう最小関数体にする。
3) `make mcp` 実行。

## 品質ゲート
- `make mcp`

## 動作確認
- なし（未配線）。

## ロールバック
- 追加ディレクトリを削除。

## Progress / Notes
- 2025-09-28: 着手
- 2025-09-28: 追加・mcp通過（sqlite/{user,weapons,inventory}.go を追加）

## 関連
- `docs/architecture/README.md` 9章
