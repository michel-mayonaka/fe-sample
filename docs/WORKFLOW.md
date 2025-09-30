# 開発者向けワークフロー（WORKFLOW）

目的: 5分で全体像を掴み、迷いなく手を動かせる最小ガイド。

## ライフサイクル概要
- Discovery → Backlog(accepted) → Story → Consumed(退避) → Finish（索引反映）
- 索引/自動生成:
  - Backlog: `make backlog-index`（accepted/*.md がソース。目的/背景/DoDを各3行で展開＋リンク）
  - 完了索引: `make story-index`

## 日々の流れ（サマリ）
- 朝: Discovery のトリアージ（48h以内に open/promote/decline を確定）
- 作業開始: 明示の開始指示が出た Story のみ着手（詳細は「作業開始ルール」）
- 終了時: DoDチェック → `make finish-story` → 索引/Backlog 再生成

## プロジェクト管理
- 新規課題の起票: `make new-discovery SLUG=... [TITLE=..] [PRIORITY=Px] [STORY=YYYYMMDD-slug]`
  - 実行時に類似候補をサジェスト（重複抑止）
- 採択（Backlogへ）: `make promote-discovery FILE=... [PRIORITY=Px]`
  - accepted/ へ移動。Backlog は accepted のみから自動生成
- ストーリー化と退避:
  - `FROM_DISCOVERY=<accepted.md> make new-story SLUG=...`
  - または `make consume-discovery FILE=<accepted.md> STORY_DIR=stories/YYYYMMDD-slug`
  - consumed/ へ退避（Backlogから自動で消える）

## 機能開発
- ブランチ: `feat/<slug>` 推奨。コミットは Conventional Commits（`[story:YYYYMMDD-slug]` 任意）
- 最低ゲート: 本体コードに対する変更は `make mcp` 成功が前提（vet/build/lint/test）
- PRの本文: 目的/影響範囲/検証手順（手動確認やコマンド）を簡潔に記載
- マージ: Squash を既定（履歴の簡潔性）。必要に応じ Rebase 可

## 並行作業/衝突回避
- 起票時に「競合しない計画単位」を明記（対象ディレクトリ/主要型/想定変更点）
- Discovery は 1課題=1ファイル。Backlog は accepted のみ参照→ストーリー化で自動的に痩せる

## レビュー/完了
- レビュー観点: 目的とDoD一致/副作用範囲/命名規約/テスト/ロールバック
- 完了手順:
  - DoD 完了 → `make finish-story SLUG=...` → `make story-index` → `make backlog-index`
  - 関連ファイル/リンクの整合（README/Docs/ストーリー）

## 作業開始ルール（重要）
- ストーリー作成直後の実装着手は禁止（ディスカッション/レビュー/Story配下MDの更新のみ可）
- 実装開始には「開始」指示が必要（例: 「20250930-xxx を開始」）
- ルール違反の変更は Revert または Story に移管してやり直す

## コマンド早見表
- Discovery: `make new-discovery` / `make promote-discovery` / `make decline-discovery`
- ストーリー: `make new-story SLUG=...` / `make finish-story SLUG=...`
- 退避: `make consume-discovery FILE=... STORY_DIR=stories/YYYYMMDD-slug`
- 索引/生成: `make story-index` / `make backlog-index` / `make discovery-index`
- 検証: `make validate-meta`（必須メタの警告）/ `make mcp`

