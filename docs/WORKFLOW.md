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

## ストーリーの作成（Codex へのプロンプト例／ブランチ）
例1（新規ストーリーをゼロから作る）
```
Codex へ:
新しいストーリーを起票して。タイトルは「dev-workflow-doc」。
目的/スコープ/DoD/タスク雛形を用意して、実装は開始しないで。
```

例2（Backlog/accepted からストーリー化する）
```
Codex へ:
accepted の 2025-10-02-xxx をベースにストーリー化して。
FROM_DISCOVERY を使って退避し、DoD とタスクを整えて。実装は開始しないで。
```

ブランチ方針（コード変更を伴う場合のみ）
- 作業ブランチ: `feat/<slug>` を推奨（例: `feat/dev-workflow-doc`）。
- ドキュメント/ストーリーのみの編集は master 直コミット可。

## ストーリーの作業開始（Vibe‑kanban へのプロンプト例）
例1（単一ストーリーを開始）
```
Vibe‑kanban へ:
カード「20250930-dev-workflow-doc」を開始（In Progress）。
競合しない計画単位: docs/WORKFLOW.md と AGENTS.md の一部追記のみ。
```

例2（並行ストーリー開始時の注意付き依頼）
```
Vibe‑kanban へ:
カード「20251002-provider-ui-decouple」を開始（In Progress）。
競合しない計画単位: internal/game/data/* の Read系I/F、UI層はノータッチ。
他ストーリーと衝突しそうなら指摘して保留にしてください。
```
