# [story] internal/user 再編（モデル/I-O 分離）

- ステータス: [完了]
- 日付: 2025-09-29
- 参照: docs/ARCHITECTURE.md, docs/NAMING.md, docs/REF_STORIES.md

## 背景
`internal/user` にモデル・入出力・境界が混在し、依存関係の明確化と再利用性が阻害されている可能性がある。

## 目的
- モデル、ポート（抽象）、具体実装（アダプタ）を明確に分離し、保守容易性とテスト容易性を向上させる。

## スコープ
- 型/スキーマを `internal/model/user` へ集約。
- 具体実装（JSON/FS 等）を `internal/infra/userfs` に抽出。
- 既存参照の置換とインポート循環の解消。

## 非スコープ
- ドメインロジックの仕様変更。
- ストレージ方式変更（例: SQLite 化）。

## 成果物 / DoD
- `internal/user` の混在を解消し、`internal/model/user` と `internal/infra/userfs` に分離済み。
- 既存 `repo` インタフェース群（User/Inventory/Weapons）を Port として継続採用し循環なし。
- `make mcp` 成功（vet/build/lint/test すべてOK）。
- docs（README/API/ARCHITECTURE/AGENTS）を現構成に同期済み。

## 影響範囲
- `internal/game/*`（usecase/provider）
- `internal/infra/*`（データアクセス）
- 単体テスト（ユーザ関連）

## リスクと対策
- リスク: 広範な import 置換による影響。
  - 対策: 「探して置換」を最小単位で段階適用、`git grep`/`rg` で差分範囲を限定。
- リスク: 循環依存の発生。
  - 対策: provider を最上位の抽象として固定し、infra から usecase への参照を禁止。

## 計測/検証
- `make mcp` 結果、テストカバレッジの変化、インポートグラフの簡素化（ノード/エッジ数）。

## 次アクション（タスク化方針）
- 01_依存グラフ可視化（read-only）
- 02_provider インタフェース確定
- 03_model/user へ型移設
- 04_infra 実装新設＋差し替え
- 05_ドキュメント更新
