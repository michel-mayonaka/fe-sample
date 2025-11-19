# AI Operations（Codex/Vibe 連携の参考資料）

この文書は、AI駆動開発の「内部動作」や自動化スクリプトの詳細を知りたい場合の参考用です。普段の開発では読む必要はありません。

## 概要
- Codex（このリポで使うAIコーディングエージェント）の役割、計画/パッチ適用/検証の流れ。
- ストーリー/Discovery 自動化（scripts/*, Make ターゲット）の連携図。

## 参考リンク
- 1 枚絵: `docs/ops-overview.md`
- 実務ガイド: `docs/KNOWLEDGE/workflows/overview.md`（開発者向けワークフロー概要）
- ストーリー運用: `docs/KNOWLEDGE/workflows/stories.md`
- エージェントガイド: `AGENTS.md`
- コマンド一覧: `Makefile` / `scripts/`

（詳細は必要になった時点で追記）

## Vibe‑kanban 連携用 Codex 設定例

Vibe‑kanban から Codex（AI エージェント）にタスクを実行させる場合の、推奨設定イメージです。

### ワークスペースパス
- `workspace_dir`: リポジトリルート（このファイルと `AGENTS.md` が存在するディレクトリ）
  - 例: ローカルでは `.../ebiten/ui_sample` をワークスペースとして渡す。
  - Codex はワークスペース直下の `AGENTS.md` を自動で読み、ここに記載されたルールを前提として動く。

### System プロンプト例
Vibe‑kanban 側で Codex を起動する際の、固定の system メッセージ例:

> あなたは ebiten/ui_sample リポジトリ専用の AI コーディングエージェントです。  
> - 作業ディレクトリはリポジトリルート（AGENTS.md が存在する場所）とします。  
> - まず `AGENTS.md` を読み、記載されたルール・優先順位に必ず従ってください。  
> - ストーリー/Discovery/タスク運用は `docs/KNOWLEDGE/workflows/stories.md` と `docs/KNOWLEDGE/workflows/vibe-kanban.md` を前提とします。  
> - ユーザーから `stories/YYYYMMDD-slug` のような相対パスが渡された場合、それを「現在のストーリー」とみなし、最初にその `README.md` と `tasks/` 配下を読んでください。  
> - ユーザーから「後続タスクを Discovery として起票して」などと指示された場合は、`stories/discovery/` 直下に `YYYYMMDD-HHMMSS-slug.md` を追加し、`ステータス: [open]` と `関連ストーリー: <元ストーリーの slug>` を必ず記載してください。  
> - 明示的な開始指示が無いストーリーでは、コード（`*.go` や docs 本体）の実装変更は行わず、Markdown での設計・メモ・タスク分割に留めてください。  
> - 変更はできるだけ小さく進め、`make mcp` などの重いコマンドは実行前に必ずユーザーへ確認してください。

この system プロンプトは固定とし、Vibe‑kanban 側からは user メッセージで
- 例: 「`stories/20251117-pr-template-story-ref` の作業開始して」「このストーリーで後続タスク候補を Discovery として起票して」
のように「ストーリーの相対パス」と具体的な依頼内容を渡す想定です。
