# 開発者向けワークフロー概要

本プロジェクトの開発フロー全体像と、関連するドキュメント/ツールの位置づけをまとめます。詳細な手順や運用ルールは `docs/workflow/*.md` を参照してください。

## コンポーネント一覧
- Codex（AI コーディングエージェント）
- ストーリー/Discovery/Backlog（`stories/`）
- Vibe‑kanban（タスクボード）
- ローカル開発・検証（`make` コマンド群）
- CI（GitHub Actions）

## 高レベルな流れ
1. 課題/アイデアを Discovery として起票し、Backlog に整理する。
2. Backlog からストーリーを選び、`make new-story` でストーリーを作成する。
3. Codex と Vibe‑kanban を使いながら、ストーリー配下のタスクを進める。
4. ローカルで `make smoke` / `make mcp` などを実行して検証する。
5. CI で `smoke-offline` → `build-and-lint` → `ui-build-strict` の順に検証する。
6. DoD を満たしたらストーリーを完了し、必要に応じて `make finish-story` でアーカイブする。

各フェーズの詳細は以下を参照してください:
- ストーリー/Discovery/Backlog 運用: `docs/workflow/stories.md`
- ローカル開発・検証フロー: `docs/workflow/local-dev.md`
- CI 構成と検証内容: `docs/workflow/ci.md`
- Vibe‑kanban を用いた日々の運用: `docs/workflow/vibe-kanban.md`

