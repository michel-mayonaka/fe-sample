# [story] Discovery 自動化とBacklog衝突回避（索引/自動生成/重複検知/メタ統一）

ステータス: [進行中]
担当: @tkg-engineer
開始: 2025-09-30 18:42:03 +0900

## 目的・背景
- 並列作業での `BACKLOG.md` 競合を極小化し、Discovery→採択（Backlog）までの流れを自動化・可視化する。
- Discovery と Story のメタを統一し、将来の機械可読な運用に備える。

## スコープ（成果）
- `stories/discovery/accepted/*.md` から `stories/BACKLOG.md` を自動生成（手編集を原則廃止）。
- `stories/discovery/INDEX.md` を生成（open/promoted/declined 一覧）。
- Discovery/Story のメタキーを最小集合に統一（ステータス/担当/開始/優先度）。
- `new_discovery.sh` に既存類似検知（`rg -i`）サジェストを追加。
- 運用ルール（48hトリアージ、昇格は1日1回など）をドキュメント化。

## 受け入れ基準（Definition of Done）
- [ ] `make backlog-index` で `stories/BACKLOG.md` が生成される（accepted/*.md が唯一のソース）。
- [ ] `make discovery-index` で `stories/discovery/INDEX.md` が生成される。
- [ ] Discovery/Story の先頭メタが統一キーで出力/検証される。
- [ ] `new_discovery.sh` 実行時に既存近似タイトルのサジェストが表示される。
- [ ] `docs/REF_STORIES.md`/`AGENTS.md` に運用ルールとコマンドが追記される。

## 工程（サブタスク）
- [ ] 01_backlog_index: accepted から BACKLOG 生成スクリプト
- [ ] 02_discovery_index: discovery の一覧生成
- [ ] 03_meta_unify: メタキー統一とバリデータ（軽量）
- [ ] 04_dup_suggest: new_discovery に重複サジェスト追加
- [ ] 05_docs: ドキュメント更新（REF_STORIES/AGENTS）
- [ ] 06_ops_rules: トリアージ/昇格頻度の明文化

## 計画（目安）
- 見積: 1.5〜2.0h（実装/ドキュメント/検証）

## 進捗・決定事項（ログ）
- 2025-09-30 18:42:03 +0900: ストーリー作成

## リスク・懸念
- 例: 依存の変更、CI制約 など

## 関連
- PR: #
- Issue: #
- Docs: `docs/...`
