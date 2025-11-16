# 20251117-spec-status-metadata — specs 状態メタデータ整備

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-17 00:29:58 +0900

## 目的・背景
- `docs/specs/` 配下の仕様ごとに「実装・テストの反映状況」が分かるようにメタデータを整理し、仕様と実装のズレを早期に検知しやすくする。
- 「Backlog/Story = 未完の仕事」「spec = 振る舞いの単一ソース＋状態」という役割分担を明確にし、運用コストを下げる。

## スコープ（成果）
- system/ui specs に共通のメタキー（例: `状態`, `主な実装`, `最新ストーリー`）を定義したガイドを用意する。
- 代表的な既存 spec（例: `docs/specs/ui/status_screen.md`）にメタ情報を実際に付与し、運用イメージを示す。
- Story/Discovery/Backlog から spec のメタ情報をどう更新するかの軽いワークフローを整理する（docs/specs/AGENTS.md 等への追記を含む）。

## 受け入れ基準（Definition of Done）
- [x] specs 用メタキー（`状態: [spec-only|impl-partial|impl-done|impl-tested]` など）がドキュメントとして明文化されている。
- [x] 少なくとも 1 つ以上の system spec と 1 つ以上の ui spec に、メタ情報が実際に付与されている。
- [x] Story/Discovery/Backlog から spec の `状態` を更新する標準フローが文章で示されている。

## 工程（サブタスク）
- [x] メタキー案と状態遷移の整理（リンク: `stories/20251117-spec-status-metadata/tasks/01_spec_metadata_design.md`）
- [x] 代表 spec へのメタ情報付与（リンク: `stories/20251117-spec-status-metadata/tasks/02_apply_metadata_to_specs.md`）
- [x] AGENTS/README への運用ルール追記（リンク: `stories/20251117-spec-status-metadata/tasks/03_update_agents_and_docs.md`）

## 計画（目安）
- 見積: 1〜2 セッション
- マイルストン: M1 メタキー案整理 / M2 代表 spec への適用 / M3 ドキュメント反映

## 進捗・決定事項（ログ）
- 2025-11-17 00:29:58 +0900: ストーリー作成（仕様と実装の状態を spec メタデータで管理する方針を検討するため）。
- 2025-11-17 00:38:45 +0900: メタキー設計と代表 spec への適用、および AGENTS/README へのルール追記を実施。

## リスク・懸念
- 既存 Story/Backlog の運用とメタデータ更新のルールが噛み合わないと、どちらか一方が形骸化する可能性がある。
- 状態の粒度を細かくしすぎると更新コストが高くなるため、最小限の段階に抑える必要がある。

## 関連
- Docs: `docs/specs/AGENTS.md`, `docs/specs/README.md`, `docs/workflows/stories.md`

- 2025-11-17 00:38:01 +0900: アーカイブ（finish へ移動）

- 2025-11-17 00:46:15 +0900: アーカイブ（finish へ移動）
