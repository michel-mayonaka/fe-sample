# 20250930-provider-ui-decouple — Provider から UI 型依存を分離（Port/Adapter 再整理）

ステータス: [準備完了]
担当: @tkg-engineer
開始: 2025-09-30 00:00:00 +0900

## 目的・背景
- `internal/game/data.TableProvider` を純参照Portに保ち、UI 型（`uicore.Unit` 等）への依存を排除する。
- `UI ↔ Adapter` 境界へ変換責務を集約し、移植性とテスト容易性を向上する。

## スコープ（成果）
- Provider が返す型を `model/*`/`user/*` に限定（UI型ゼロ）。
- UI 変換は `internal/game/ui/adapter` に集約（ユニット/装備表示など）。
- 呼び出し箇所を見直し、必要に応じて Adapter 経由へ差し替え。
- 最小のドキュメント整合（`docs/ARCHITECTURE.md`/`docs/API.md`）。

## 受け入れ基準（Definition of Done）
- [ ] `internal/game/data/` 直下で `uicore.` を参照していない（`rg -n 'uicore\.' internal/game/data` が 0）。
- [ ] Provider IF（`provider.go`）に UI 型が含まれないことを確認。
- [ ] 主要画面（一覧/ステータス/シミュ）で動作確認。
- [ ] `make mcp` がグリーン。

## 工程（サブタスク）
- [ ] 現状調査: 依存と使用箇所の棚卸し — `tasks/01_inventory.md`
- [ ] 方針確定: Port/Adapter の役割定義 — `tasks/02_design.md`
- [ ] 置換: 呼び出しの再配線と最小修正 — `tasks/03_replace_calls.md`
- [ ] テスト/文書: 回帰確認とドキュメント更新 — `tasks/04_tests_docs.md`

## 計画（目安）
- 見積: 2 時間（1セッション）
- マイルストン: M1 調査 / M2 置換 / M3 検証

## 進捗・決定事項（ログ）
- 2025-09-30 00:00:00 +0900: ストーリー作成（BACKLOG昇格）
- 2025-10-01 05:11:00 +0900: ステータスを[準備完了]へ更新（レビュー完了）

## リスク・懸念
- UI/Adapter への集約で一時的に呼び出しが増える可能性。
- 既存テストのモック差し替えコスト。

## 関連
- Discovery: `stories/discovery/consumed/2025-09-30-migrated-03.md`
- 実装箇所: `internal/game/data/provider.go`, `internal/game/ui/adapter/`, `docs/ARCHITECTURE.md`
