# 20250929-ref-architecture — アーキテクチャ再整備（完了）

ステータス: [完了]
担当: @tkg-engineer

## 目的・背景
- Port/Provider/CQRS 方針に沿って UI 依存を整理し、Usecase へコマンド集約・Provider へクエリ統一。
- 最終棚卸し（14_final_audit）で参照経路/コメント/テストを横断確認し、残課題を `tasks/known_issues.md` に集約。

## スコープ（成果）
- Port 定義と配線: `scenes` に `DataPort/BattlePort/InventoryPort`、`usecase.App` 実装。
- UI 依存の是正: 直接 Repo 書込の排除（例: Status の保存は `DataPort.PersistUnit` に委譲）。
- Provider 統一: `gdata.SetProvider(app)` 経由で武器/アイテム定義参照。
- テスト拡充:
  - `/pkg/game`: 予測内訳・三すくみ・地形・AS 境界のユニットテスト追加。
  - `/internal/usecase`: `ReloadData/PersistUnit/RunBattleRound/Equip(Weapon|Item)` のハッピーパス/境界ケース。
- ビルドタスク: `make mcp` に `test-all` を追加（`-race -cover` 付き、対象は `./pkg/... ./internal/usecase`）。

## 受け入れ基準（DoD）
- `make mcp` が成功（`check-all` + Lint 非strict + `test-all`）。
- 主要フロー（一覧/ステータス/在庫/戦闘プレビュー）が従来通り操作可能。
- 既知の課題は `tasks/known_issues.md` に集約（クリティカル 0）。

## 動作確認手順（抜粋）
- 速い確認: `GOFLAGS='-mod=readonly' GOWORK=off go test -count=1 ./pkg/game ./internal/usecase`
- 追加検証: `make mcp`（重い場合は `TEST_FLAGS=""` で軽量化）

## 進捗ログ（要約）
- 2025-09-29: `14_final_audit` 実施。UI 直書きを排除、テスト拡充、`mcp` にテスト統合。

## 関連
- タスク: `tasks/ref_architecture/index.md`, `tasks/ref_architecture/14_final_audit.md`
- 既知課題: `tasks/known_issues.md`
- 設計: `docs/ARCHITECTURE.md`, `docs/API.md`
