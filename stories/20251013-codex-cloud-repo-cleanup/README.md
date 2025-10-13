# 2025-10-13 — codex-cloud 環境対応のリポジトリ整理（`20251013-codex-cloud-repo-cleanup`）

- 状態: 進行中（2025-10-13 開始）
- 作成: 2025-10-13
- オーナー: TBD
- 関連: README.md, docs/WORKFLOW.md, docs/NAMING.md, docs/REF_STORIES.md

## 目的
codex-cloud（ワークスペース書込のみ・ネットワーク制限・承認制実行）環境で、開発/検証が滞りなく行えるよう、リポジトリ構成・Make タスク・ドキュメントを整備する。

## 背景
- 開発環境の標準化とクラウド実行の増加に伴い、ネットワーク制限や一時ディレクトリ制約への適合が必要。
- 既存の `MCP_OFFLINE` 方針を前提に、オフライン/サンドボックスで再現性の高い動作を目指す。

## スコープ
- Make タスクの整理（オフライン前提、重いUI検証の分離）。
- 生成物ディレクトリの統一（例: `out/` or `site/`）と `.gitignore` 方針。
- ベンダリング/キャッシュ戦略（`-mod=vendor` と運用ルール）。
- 最小スモークチェック（論理層のみ）ターゲットの追加設計。
- 開発手順の文書化（codex-cloud 専用の実行手順）。

## 非スコープ
- 依存パッケージの大規模更新や機能追加。
- UI の設計刷新。

## 成果物（Deliverables）
- 整理後のタスク設計案（Make/スクリプトの変更提案）。
- codex-cloud 実行ガイド（README 追記、必要なら `docs/CODEX_CLOUD.md`）。
- スモークチェック手順（コマンド例）と期待結果。

## 完了の定義（DoD）
- codex-cloud 環境で、オフライン前提の検証がワンコマンドで完了する（UI 依存は非strictで迂回）。
- 生成物の出力先とクリーン手順が統一され、再現性が担保される。
- 開発者が README を見ればセットアップ〜検証が実行できる。

## 進捗メモ（2025-10-13）
- Make: `smoke`/`offline`/`clean`/`build-out` を追加し `out/` 集約。
- Docs: `docs/CODEX_CLOUD.md` 追加、`README`/`AGENTS` 更新。
- CI: `smoke-offline` ジョブ追加、README にフロー追記。

## リスクと対策
- ローカル/CI/クラウドでの挙動差: フラグ既定値整理（例: `MCP_STRICT=0`）と明記。
- ベンダリングのメンテコスト: 更新手順を明記し、定期的に検証。

## 依存関係
- Go/Ebitengine の推奨バージョン。
- 既存 CI 方針（`make mcp` 非strict）との整合。

## タスク
1. tasks/01-plan-and-scope.md — 目的/境界/方針の確定
2. tasks/02-env-constraints-checklist.md — 制約整理と適合チェックリスト
3. tasks/03-makefile-offline-and-targets.md — Make タスク再設計（オフライン/スモーク）
4. tasks/04-artifacts-and-gitignore.md — 生成物ディレクトリ統一と無視設定
5. tasks/05-ci-alignment.md — CI 方針との整合（ジョブ分離/依存）
6. tasks/06-docs-codex-cloud-guide.md — 実行ガイド整備（README/Docs）
7. tasks/07-migration-notes.md — 既存開発者向け移行ノート

## 開始指示（テンプレ）
実装開始時は、以下のように明示してください。

> 20251013-codex-cloud-repo-cleanup を開始
