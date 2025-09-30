# [story] ストーリー運用の強化（テンプレ/秒精度ログ/finish索引/スクリプト）

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-09-30 17:53:37 +0900

## 目的・背景
- ストーリー/タスク運用の表記揺れと可視性不足（完了一覧の俯瞰）を解消する。
- 秒精度のタイムスタンプを標準化し、履歴の機械可読性を高める。

## スコープ（成果）
- テンプレの秒精度タイムスタンプ対応（Story/Task）。
- 生成/アーカイブスクリプト強化（表記揺れ吸収、完了ログ追記）。
- 完了ストーリー索引 `stories/finish/INDEX.md` の自動生成と Make ターゲット追加。
- ガイド更新（REF_STORIES のログ記法を秒精度へ）。
- レポジトリ直下 `tasks/` の廃止（内容を該当ストーリー配下 `tasks/` または `stories/BACKLOG.md` へ移管）。

## 非スコープ
- CI での厳格検証や PR テンプレの導入は別ストーリーで検討（参考として README に残す）。

## 受け入れ基準（Definition of Done）
- [x] `stories/_TEMPLATE/README.md` に `開始:` と秒精度ログ例を追加。
- [x] `stories/_TEMPLATE/tasks/01_sample.md` を追加（秒精度ログ例）。
- [x] `scripts/new_story.sh` が作成直後の README に秒精度 `開始:` と初期ログを出力。
- [x] `scripts/finish_story.sh` が `ステータス:` の表記揺れを吸収して `[完了]` 化し、秒精度のアーカイブログを追記。
- [x] `scripts/gen_story_index.sh` により `stories/finish/INDEX.md` を生成できる。
- [x] `make story-index` で上記スクリプトを実行可能。
- [x] `docs/REF_STORIES.md` のログ記法に秒精度を明記。
- [x] 直下 `tasks/` の廃止（全ファイル移管済み、残ファイル 0、CI/README 参照箇所の更新済み）。

## 工程（サブタスク）
- [x] テンプレ修正（Story/Task）
- [x] new_story.sh 改修（秒精度/ディレクトリ名をタイトルへ）
- [x] finish_story.sh 改修（表記揺れ対応＋完了ログ追記）
- [x] 完了索引スクリプト `scripts/gen_story_index.sh` 追加
- [x] Makefile に `story-index` 追加
- [x] 直下 `tasks/` の棚卸し（存否/移管先の決定）
- [x] 各ファイルの移管（ストーリー配下/Backlog へ）
- [x] `tasks/` ディレクトリの廃止（空を確認のうえ削除）
- [ ] CI 統合（任意。別PRでも可）
- [ ] 既存 Story の先頭表記統一（必要箇所のみ）

## 計画（目安）
- 見積: 0.5〜1.0h（実装） + 0.5h（既存 Story の軽微是正） + 0.5〜1.0h（直下 `tasks/` の移管・廃止）

## 進捗・決定事項（ログ）
- 2025-09-30 17:53:37 +0900: ストーリー作成
- 2025-09-30 17:54:10 +0900: テンプレ/スクリプト実装＆Make 追加

## リスク・懸念
- sed/awk の差異（BSD/GNU）により置換が一部環境で失敗する可能性 → BSD/GNU 両対応を実装済み。

## 関連
- Docs: `docs/REF_STORIES.md`

- 2025-09-30 17:58:38 +0900: DoD に直下 tasks 廃止を追加し、計画を調整

- 2025-09-30 18:05:29 +0900: 直下 tasks の棚卸し開始（移管先の決定/分類）

- 2025-09-30 18:05:36 +0900: 直下 tasks の棚卸し開始（移管先の決定/分類）
- 2025-09-30 18:05:36 +0900: Backlog へ内容を統合（battle_map/data_repo/docs/domain/sqlite/testing/ui_scaling）
- 2025-09-30 18:05:36 +0900: finish 配下への歴史タスク移管（arch-doc-review/arch-review/scenes-helper-rehome/mcp-offline/battle-notes）

- 2025-09-30 18:47:31 +0900: アーカイブ（finish へ移動）
