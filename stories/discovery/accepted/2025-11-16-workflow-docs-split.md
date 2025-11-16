# discovery: docs/WORKFLOW.md をディレクトリ化して細分化を行う

ステータス: [promoted]
担当: @tkg-engineer
作成: 2025-11-16 00:00:00 +0900
優先度: P2
提案元: backlog-request
関連ストーリー: N/A

## 目的
- `docs/WORKFLOW.md` を章単位に分割し、`docs/workflow/` 配下へ整理して参照性・更新性を高める。

## 背景
- 単一ファイルに多くの内容が集約され肥大化しており、章内リンクや関連ドキュメントとの整合が取りづらい。

## DoD候補
- `docs/workflow/` ディレクトリを新設し、主要章（概要/ローカル検証/CI方針/変数運用/MCP_STRICT・MCP_OFFLINE/ストーリー運用フロー など）に分割。
- 目次（索引）ファイル `docs/workflow/README.md` を用意し、分割先へのリンクを集約。
- 既存参照（`AGENTS.md`/`README.md`/`docs/REF_STORIES.md`）のリンクを新構成へ更新。
- `make mcp` がグリーン、`make check-all` も成功（UI ビルドは既定運用に従う）。
- 変更履歴をコミットメッセージに記録（目的/影響範囲/検証手順）。

## 関連
- docs/WORKFLOW.md, AGENTS.md, docs/REF_STORIES.md

## 進捗ログ
- 2025-11-16 00:00:00 +0900: Backlog へ追加（docs/WORKFLOW.md の分割・ディレクトリ化の提案）。

