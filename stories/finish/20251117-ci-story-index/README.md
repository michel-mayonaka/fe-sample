# 20251117-ci-story-index — CI にストーリー検証/索引生成を統合

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-17 01:58:08 +0900

## 目的・背景
- PR/Push 時に `make story-index` と最小限の Story 検証（テンプレ必須項目の存在チェックなど）を自動実行し、ストーリー索引とメタ情報の整合性を常に保てるようにする。
- 手動運用だと Story の README/タスク/finish 索引とのズレや更新漏れが発生しやすいため、CI による軽量な検証ステップを追加したい。

## スコープ（成果）
- GitHub Actions（または将来の CI）に `make story-index` と軽量メタ検証スクリプト（`scripts/validate_meta.sh` 等）を組み込むワークフロー案が決まっている。
- `stories/finish/INDEX.md` を自動更新するフローと、レビュー/マージ時の扱い（自動コミットするか否か）が整理されている。
- エラー時の挙動（必須メタ欠落などで CI を fail にする条件）が明文化されている。

## 受け入れ基準（Definition of Done）
- [x] CI ワークフローに Story 検証/索引生成ステップを追加する案が docs/workflows/ci.md 等に記載されている。
- [x] `make story-index`/`make validate-meta` をローカルで実行したときの期待する使い方が簡潔に説明されている。
- [x] Backlog 上の discovery `CI にストーリー検証/索引生成を統合` が consumed になり、本ストーリーと相互リンクされている。

## 工程（サブタスク）
- [ ] 現状の `scripts/gen_story_index.sh` / `scripts/validate_meta.sh` の挙動確認と要件整理 — `stories/20251117-ci-story-index/tasks/01_tools-audit.md`
- [ ] CI ワークフロー案のドラフト作成（ローカル実行含む） — `stories/20251117-ci-story-index/tasks/02_ci-workflow-design.md`
- [ ] docs/workflows/ci.md への反映と Backlog の更新 — `stories/20251117-ci-story-index/tasks/03_docs-and-backlog-update.md`

## 計画（目安）
- 見積: 1〜2 セッション
- マイルストン: M1: 要件整理 / M2: CI 案ドラフト / M3: docs 反映

## 進捗・決定事項（ログ）
- 2025-11-17 01:58:08 +0900: ストーリー作成（discovery: 2025-09-30-migrated-01 から昇格）
 - 2025-11-17 02:04:12 +0900: README とサブタスクの内容をレビューし、CI への統合方針に問題がないことを確認したためステータスを[準備完了]へ更新。
 - 2025-11-19 01:46:49 +0900: 作業開始につきステータスを[進行中]へ更新。
 - 2025-11-19 01:52:18 +0900: Story 索引生成/メタ検証ツールの棚卸し・CI ワークフロー更新・docs/backlog 整理を完了し、受け入れ基準を満たしたためステータスを[完了]へ更新。

## リスク・懸念
- CI 実行時間の増加や、ストーリー未整備時に PR が詰まりすぎるリスク。

## 関連
- PR: #
- Issue: #
- Docs: `docs/workflows/stories.md`, `docs/workflows/ci.md`, `scripts/gen_story_index.sh`, `scripts/validate_meta.sh`
- Discovery: `stories/discovery/consumed/2025-09-30-migrated-01.md`

- 2025-11-19 01:58:43 +0900: アーカイブ（finish へ移動）
