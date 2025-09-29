# [task] Infra 実装の新設と参照差し替え

- ステータス: 完了
- 目的: Provider を満たす具体実装（例: FS/JSON ローダ等）を `internal/infra/userfs` に実装し、呼び出し元を差し替える。

## 入力
- タスク02の Provider 定義、タスク03のモデル配置。

## スコープ
- `internal/infra/userfs`（仮名）に実装を追加。
- コンポジションルート（`cmd/ui_sample` 等）で Provider を注入。
- 旧 `internal/user` 実装からの呼び出しを置換/削除。

## 非スコープ
- ストレージ方式の刷新（形式変更やSQLite移行は対象外）。

## 手順
- 新パッケージ作成、`provider` インタフェースを満たす実装を追加。
- 依存注入（生成/渡し先のコード）を更新。
- 旧コードの参照を削除し、`rg scenes|usecase` で漏れがないか確認。
- `make mcp` 実行、ユニットテスト補強（必要に応じ）。

## DoD（完了条件）
- 旧 `internal/user` 実装への依存がなくなっている。
- `usecase` は `provider` のみ参照し、`infra` へは依存していない。
- `make mcp` グリーン。

## コマンド例
- `rg -n "\binternal/user\b"`（0件であること）
- `go test ./... -race -cover`
