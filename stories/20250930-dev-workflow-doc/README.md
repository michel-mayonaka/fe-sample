# [story] 開発者向けワークフローガイド策定（WORKFLOW.md/AGENTS更新）

ステータス: [提案中]
担当: @tkg-engineer
開始: 2025-09-30 19:43:16 +0900

## 目的・背景
- 開発者が5分で全体像を掴めるワークフロードキュメントを整備し、Discovery→Backlog→Story→Finish の運用と実装着手ルールを明文化する。

## スコープ（成果）
- 新規: `docs/WORKFLOW.md` の追加（実装フロー/プロジェクト管理/並行作業/レビュー/マージ/コマンド早見表）。
- 既存更新: `AGENTS.md` に「作業開始ルール」を追記（ストーリー作成直後の実装禁止、開始指示必須）。
- 既存更新: `README.md` から WORKFLOW への導線追加（リンク）。

## 受け入れ基準（Definition of Done）
- [ ] `docs/WORKFLOW.md` が合意構成で作成されている。
- [ ] `AGENTS.md` に「作業開始ルール（開始指示があるまで実装禁止）」が追記されている。
- [ ] `README.md` に WORKFLOW へのリンクがある。
- [ ] `make story-index`/`make backlog-index` 実行で索引が整合。

## 工程（サブタスク）
- [ ] 01_outline.md（章立て確定）
- [ ] 02_draft_workflow.md（本文初稿）
- [ ] 03_update_agents.md（開始ルール追記）
- [ ] 04_review.md（レビュー反映）
- [ ] 05_publish.md（リンク整備・索引確認）

## 計画（目安）
- 見積: 1.0〜1.5h（初稿/合意/反映）
- マイルストン: M1 章立て → M2 初稿 → M3 レビュー → M4 公開

## 進捗・決定事項（ログ）
- 2025-09-30 19:43:16 +0900: ストーリー作成

## リスク・懸念
- 用語や既存ドキュメントとの不整合 → 章立て段階で整合レビューを実施。

## 関連
- Docs: `docs/REF_STORIES.md`, `AGENTS.md`
