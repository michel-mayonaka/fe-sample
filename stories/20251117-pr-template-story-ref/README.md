# 20251117-pr-template-story-ref — PR テンプレ/Story 参照の必須化（軽量）

ステータス: [準備完了]
担当: @tkg-engineer
開始: 2025-11-17 01:58:08 +0900

## 目的・背景
- PR やコミットに対象ストーリー（`YYYYMMDD-slug`）と DoD/検証手順を明示的に含めることで、変更の追跡性とレビュー効率を高める。
- 現状は PR/コミットメッセージから関連ストーリーを特定するのに時間がかかる場合があり、レビュアー/将来の読者にとって負荷になっている。

## スコープ（成果）
- `.github/PULL_REQUEST_TEMPLATE.md` に Story/DoD/検証手順の入力欄が追加されている。
- ストーリー識別子の書式（`YYYYMMDD-slug`）や記入ルールが docs 側に明文化されている。
- 将来 CI で Story 記述を軽く検証できるようにするためのフック（例: 特定キーワードの有無）方針が決まっている。

## 受け入れ基準（Definition of Done）
- [ ] PR テンプレートに Story/DoD/検証手順の項目が追加され、実際の PR で利用できる状態になっている。
- [ ] AGENTS.md または docs/workflows/stories.md に、PR/コミットとストーリーの紐づけルールが追記されている。
- [ ] Backlog 上の discovery `PR テンプレ/Story 参照の必須化（軽量）` が consumed になり、本ストーリーと相互リンクされている。

## 工程（サブタスク）
- [ ] 現行のコミット/PR 運用と Story 参照の実態を確認 — `stories/20251117-pr-template-story-ref/tasks/01_current-pr-workflow.md`
- [ ] PR テンプレ案の作成と運用ルールのドラフト — `stories/20251117-pr-template-story-ref/tasks/02_pr-template-design.md`
- [ ] AGENTS.md / docs/workflows/stories.md への反映と Backlog の更新 — `stories/20251117-pr-template-story-ref/tasks/03_docs-and-backlog-update.md`

## 計画（目安）
- 見積: 1 セッション程度
- マイルストン: M1: 現状調査 / M2: テンプレ案 / M3: docs 反映

## 進捗・決定事項（ログ）
- 2025-11-17 01:58:08 +0900: ストーリー作成（discovery: 2025-09-30-migrated-02 から昇格）
 - 2025-11-17 02:04:12 +0900: README とサブタスクの内容をレビューし、方針に問題がないことを確認したためステータスを[準備完了]へ更新。

## リスク・懸念
- Story 記述が必須になることで、軽微な変更のフローが重くなりすぎる可能性。

## 関連
- PR: #
- Issue: #
- Docs: `AGENTS.md`, `docs/workflows/stories.md`, `.github/PULL_REQUEST_TEMPLATE.md`
