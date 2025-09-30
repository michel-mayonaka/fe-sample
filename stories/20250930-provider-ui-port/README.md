# 20250930-provider-ui-port — Provider から UI 型依存を分離（Port/Adapter 再整理）

ステータス: [進行中]
担当: @codex

## 目的・背景
- `internal/game/data.TableProvider` を UI 非依存の参照 Port に整備し、UI 実装との結合を解消して再利用性とテスト容易性を高める。
- 現行の `UserUnitByID(id) (uicore.Unit, bool)` が UI 層の型を返すため、Provider が UI 実装へ結合している。

## スコープ（成果）
- TableProvider インターフェースから UI 型を返す API を排除し、データ変換を UI Adapter 層に移管する。
- `service/ui/adapter` におけるユーザデータ→UI モデル変換の再編と必要な補完構造の導入。
- 関連する呼び出し箇所の差し替えと影響範囲のビルド/テスト整合を確認する。

## 受け入れ基準（Definition of Done）
- [x] `internal/game/data.TableProvider` から `UserUnitByID` を撤去し、代替 Port/API が導入されている。
- [x] UI 側でユーザモデルから `uicore.Unit` への変換が完結し、呼び出し元が新 API に移行している（`inventory.refreshUnitByID`）。
- [x] `make mcp` が成功し、影響範囲のテストがグリーン（ヘッドレス環境では `TEST_TAGS=headless` を付与）。
- [x] 必要に応じて `docs/API.md` / `docs/ARCHITECTURE.md` へ設計更新が反映されている。

## 工程（サブタスク）
- [x] `stories/20250930-provider-ui-port/tasks/01_port_api.md`
- [x] `stories/20250930-provider-ui-port/tasks/02_adapter_migration.md`
- [x] `stories/20250930-provider-ui-port/tasks/03_callsite_verification.md`

## 計画（目安）
- 見積: 1.5 セッション
- マイルストン: M1 Port設計固め → M2 Adapter移行 → M3 呼び出し置換と検証

## 進捗・決定事項（ログ）
- 2025-09-30: ストーリー作成、Backlog から昇格。
- 2025-09-30: Provider から `UserUnitByID` を削除。`inventory.refreshUnitByID` を `UserTable().Find` → `uicore.UnitFromUser` へ置換。関連テスト/ドキュメントを更新。
- 2025-09-30: オフライン/ヘッドレス実行での `make mcp` 成功（`TEST_TAGS=headless` 追加、`uicore` に `headless` ビルドタグを導入）。

## リスク・懸念
- Provider API 変更に伴う予期しない呼び出し漏れ。
- UI Adapter での変換ロジックが肥大化する可能性。

## 関連
- Docs: `docs/ARCHITECTURE.md`
- 参考: `internal/game/data/provider.go`, `internal/game/service/ui/adapter`
