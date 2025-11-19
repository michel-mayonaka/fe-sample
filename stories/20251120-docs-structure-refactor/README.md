# 20251120-docs-structure-refactor — docs 構成の再設計と整理

ステータス: [進行中]
担当: @tkg-engineer
開始: 2025-11-20 01:54:51 +0900

## 目的・背景
- docs/ 配下のアーキテクチャ、ナレッジ、仕様（spec）を整理し、日々の開発・レビュー・Codex 連携で参照しやすい構成にする。
- 現状はファイルの役割と分類が分かりづらく、ストーリー駆動開発や UI/ゲームロジック実装時に必要な情報へ素早く辿りづらい。
- ストーリー/Discovery で蓄積されたナレッジを docs/ に再配置し、運用しやすい「1枚絵＋分割構成」を整える。

## スコープ（成果）
- `docs/ops-overview.md` に「タスク/ナレッジ/アーキ/テスト戦略」を俯瞰できる 1 枚絵を用意する。
- `docs/architecture/`、`docs/KNOWLEDGE/`、`docs/SPECS/` のディレクトリ構成を定義し、既存 docs を移動・リネームして整理する。
- 代表的な ADR（例: `architecture/adr/ADR-0001-scene-architecture.md`）を配置し、アーキ設計の決定履歴への導線を用意する。
- 既存の仕様書を `SPECS/world|gameplay|ui` などに振り分け、画面仕様やシステム仕様から関連コードへ辿りやすくする。

## 受け入れ基準（Definition of Done）
- [ ] `docs/ops-overview.md` に、開発フロー・ストーリー運用・テスト戦略・Codex 連携の概要がまとまっている。
- [ ] `docs/architecture/` / `docs/KNOWLEDGE/` / `docs/SPECS/` の構成と命名方針が README または既存ガイドに明記されている。
- [ ] 既存 docs が新構成に沿って移動されており、壊れたリンクや参照がない（`rg` で旧パスを確認）。
- [ ] `make mcp` が成功し、必要に応じて `docs/SPECS/README.md` や AGENTS の記述も更新されている。

## 工程（サブタスク）
- [ ] docs 現状構成の棚卸しとターゲット構成の詳細設計（`tasks/01_design-docs-structure.md`）
- [ ] 既存 docs の移動・リネームとリンク修正（`tasks/02_migrate-existing-docs.md`）
- [ ] AGENTS / specs README / ops-overview の更新と最終チェック（`tasks/03_update-guides-and-index.md`）

## 計画（目安）
- 見積: 2〜3 セッション（設計 1、移行 1、仕上げ 1）
- マイルストン: 構成設計 → 移行実施 → ガイド/インデックス更新

## 進捗・決定事項（ログ）
- 2025-11-19 17:05:00 +0900: Codex 作業開始指示に伴いステータスを[進行中]へ更新し、ターゲット構成と移行計画の実装に着手。
- 2025-11-20 01:54:51 +0900: ストーリー作成
- 2025-11-20 01:56:59 +0900: README とサブタスク構成を整備し、docs 構成再設計ストーリーの下準備が整ったためステータスを[準備完了]へ更新。

## リスク・懸念
- ファイル移動に伴うリンク切れや、他ドキュメントとの整合性崩れ。
- Codex からの参照パス変更による一時的な混乱 → AGENTS/ガイド類での周知と、パス変更のサポート。

## 関連
- PR: #
- Issue: #
- Docs: `docs/KNOWLEDGE/meta/docs-structure.md`, `docs/SPECS/README.md`, `AGENTS.md`
