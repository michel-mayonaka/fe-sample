# 20250929-arch-doc-review — アーキテクチャドキュメントの再レビュー

ステータス: [完了]
担当: @yourname

## 目的・背景
- 現行実装（Scene/Service/Provider/Usecase/Model）と `docs/ARCHITECTURE.md` の不整合を解消し、境界・依存方向・責務を最新状態へ明確化する。

## スコープ（成果）
- `docs/ARCHITECTURE.md` の更新（Scene ライフサイクル: input→intent→advance→render の明記、永続は `Env.Data` 経由に統一）。
- 関連ドキュメントの整合（`docs/NAMING.md`、`README.md` の該当箇所）。
- 主要な設計図/関係図のアップデート（テキスト/ASCII 図で可）。

## 受け入れ基準（Definition of Done）
- [ ] 差分の洗い出しメモが残る（不整合一覧 / 採用・却下の判断を記録）
- [ ] `docs/ARCHITECTURE.md` 更新 PR が作成され、リンク/参照整合が取れている
- [ ] `make mcp` がグリーン
- [ ] 後続リファクタ（inventory/status/character_list）に向けた設計チェックリストが共有される

## 工程（サブタスク）
- [ ] [01_audit_current_vs_docs.md](tasks/01_audit_current_vs_docs.md): 現状コードとドキュメントの比較監査
- [ ] [02_update_architecture_docs.md](tasks/02_update_architecture_docs.md): ARCHITECTURE 更新（境界・依存・図）
- [ ] [03_define_scene_lifecycle_and_data_port.md](tasks/03_define_scene_lifecycle_and_data_port.md): Scene ライフサイクルと DataPort の規約明文化
- [ ] [04_prepare_refactor_checklist_for_scenes.md](tasks/04_prepare_refactor_checklist_for_scenes.md): 後続リファクタ用チェックリスト策定

## 計画（目安）
- 見積: X 時間 / セッション
- マイルストン: M1 / M2 / M3

## 進捗・決定事項（ログ）
- 2025-09-29: ストーリー作成。サブタスク草案を追加。
- 2025-09-29: 監査初版を作成（`artifacts/audit_findings.md`）。InventoryPort の Repo 依存を改善候補として提示。
- 2025-09-29: 方針B（Provider 統一）で PoC 実装。`UserWeapons/UserItems` を Provider に追加、在庫参照を Scenes から切替。Docs/API/ARCHITECTURE を更新。
- 2025-09-29: リファクタチェックリストを作成（`artifacts/scene_refactor_checklist.md`）。Task 04 を完了。

## リスク・懸念
- 例: 依存の変更、CI制約 など

## 関連
- PR: #
- Issue: #
- Docs: `docs/ARCHITECTURE.md`, `docs/NAMING.md`, `README.md`
