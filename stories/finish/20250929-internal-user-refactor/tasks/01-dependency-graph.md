# [task] internal/user 依存グラフの可視化（read-only）

- ステータス: 完了
- 目的: `internal/user` に対する参照元・依存方向・循環リスクを棚卸しし、以降の安全な分割計画に反映する。

## スコープ
- リポジトリ全体の `internal/user` 参照を列挙。
- インポートグラフ上の位置（親/子の関係、うち循環しそうな箇所）を把握。
- コードの変更は行わない。

## 手順
- 参照列挙: `rg -n "\\binternal/user\\b"`
- パッケージ依存: `go list -deps ./... | rg "/internal/user$" -n`
- 具体識別子の使用状況（必要に応じて）: `rg -n "User[A-Z][A-Za-z0-9_]*" internal/user`
- 結果を本ファイル下部「調査メモ」に記載（呼び出し側、目的、懸念）。

## DoD（完了条件）
- 参照元パッケージの一覧と用途メモがまとまっている。
- 想定される抽象（Provider）と具体（Infra）の境界候補を1案提示。
- 次タスク（02）の入力として不足がない。

## 想定リスク/留意
- `generated` やテストのみ使用のものは区別する（`_test.go`）。

## コマンド例
- `rg -n "\\binternal/user\\b"`
- `go list -deps ./... | rg "/internal/user$" -n`

---

### 調査メモ
- 参照元一覧:
- 用途/懸念:
