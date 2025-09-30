# 開発者向けワークフロー（最小ガイド）

目的: AI駆動開発で「何を・誰と・どこを確認するか」だけを最短で把握する。

## 誰とやり取りするか
- Codex（AIコーディングエージェント）: 設計相談→開始指示→実装/レビュー→フィニッシュまでを対話で進行。
- Vibe‑kanban: ストーリー/タスクの状態管理と合意事項の記録（DoD/決定事項）。

## 典型フロー（開発者の行動だけ）
1) ディスカッション/合意: Vibe‑kanban のカードに目的/スコープ/DoD を記録し、Codex と方針を固める。
2) 開始指示: 「YYYYMMDD-slug を開始」の一言で実装着手を許可（指示が出るまでは実装禁止）。
3) 実装/確認: Codex が変更を提案・実装。コード変更時は `make mcp` 成功が前提。疑問は対話で都度解消。
4) レビュー/終了: DoD満たす→フィニッシュ（索引/Backlogは自動再生成）。

## どのドキュメントを確認するか（チェックリスト）
- 命名規約: `docs/NAMING.md`（識別子/ファイル/パッケージ）
- ストーリー運用: `docs/REF_STORIES.md`（Discovery/Backlog/Story の扱い）
- アーキテクチャ: `docs/ARCHITECTURE.md`（層・依存・境界）
- コメント記法: `docs/COMMENT_STYLE.md`（GoDoc 最小ルール）
- ワークフロー本書: `docs/WORKFLOW.md`（本ドキュメント）
- 参考: AI内部の動作を知りたい場合は `docs/AI_OPERATIONS.md` を参照。

## 作業開始ルール（重要）
- ストーリー作成直後の実装着手は禁止。開始指示があるまで「ディスカッション/レビュー/Story配下MDの更新」のみ可。
- コード変更を含む作業は `make mcp` 成功が前提。マージは基本 Squash。

## 最小コマンド（覚えるのはこれだけ）
- ストーリー関連: `make new-story SLUG=...` / `make finish-story SLUG=...`
- Discovery関連: `make new-discovery` / `make promote-discovery` / `make consume-discovery`
- 索引・生成: `make story-index` / `make backlog-index`
