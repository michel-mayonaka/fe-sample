# [task] 型スキーマの `internal/model/user` への移設

- ステータス: 完了
- 目的: モデル（エンティティ/値オブジェクト）を入出力実装から分離し、`internal/model/user` に集約する。

## 入力
- タスク02で確定した Provider から参照される型一覧。

## スコープ
- `internal/model/user` パッケージ新設（なければ）。
- ユーザ関連の純粋な型・バリデーション（副作用なし）を移設。
- 移設に伴う import 更新（ビルドが通るまで）。

## 非スコープ
- ストレージ/I/O に依存する実装（infraへ）。

## 手順
- 移設候補ファイルの列挙。
- `model` へ移動＋パッケージ名修正。
- 参照側の import 文置換。
- `make mcp` でビルド/テストを通す。

## DoD（完了条件）
- `internal/user` に純粋型が残っていない。
- `internal/model/user` に型が集約され、`usecase`/`provider` から参照可能。
- `make mcp` グリーン。

## コマンド例
- `rg -n "type \w+ struct" internal/user`
- `rg -n "package model" internal/model || true`
