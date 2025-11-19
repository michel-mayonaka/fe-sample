# 20251119-internal-game-util-specialize — `internal/game/util` の撤去または責務特化サブパッケージ化

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 00:28:16 +0900

## 目的・背景
- 命名規約（util/helpers 禁止）に沿って、`internal/game/util` のような汎用名ディレクトリを整理し、責務が明確なサブパッケージへ分割または撤去する。
- 空ディレクトリや用途不明のヘルパー置き場をなくし、将来的な拡張やリファクタリング時の混乱を防ぐ。

## スコープ（成果）
- `internal/game/util` 配下の現状（ファイル有無/使われ方）を確認し、不要なものは削除する方針をまとめる。
- 必要な機能がある場合は、役割に応じて `rng`/`rect`/`debug` など責務特化サブパッケージへの再配置案を検討する。
- docs/KNOWLEDGE/engineering/naming.md や docs/architecture/README.md と整合する形で、util 系ディレクトリの扱い方針を整理する。

## 受け入れ基準（Definition of Done）
- [x] `internal/game/util` の現状と利用状況が整理されたメモがある。
- [x] 削除または責務特化サブパッケージ化の方針（どの機能をどこへ移すか）が決まっている。
- [x] README や関連ドキュメントに、util 名を避ける方針と今回の整理の位置付けが明文化されている。

## 工程（サブタスク）
- [x] 現状調査: `internal/game/util` 配下と利用箇所の棚卸し — `stories/20251119-internal-game-util-specialize/tasks/01_audit-util-usage.md`
- [x] 再配置/削除方針の検討: 責務特化サブパッケージ案の作成 — `stories/20251119-internal-game-util-specialize/tasks/02_design-specialized-packages.md`
- [x] ドキュメント更新と後続タスク（実装ストーリー/Backlog）の整理 — `stories/20251119-internal-game-util-specialize/tasks/03_docs-and-followups.md`

## 計画（目安）
- 見積: 1 セッション程度
- マイルストン: M1: 現状調査 / M2: 方針案 / M3: docs/Backlog 整理

## 進捗・決定事項（ログ）
- 2025-11-19 00:28:16 +0900: ストーリー作成（discovery: 2025-09-30-migrated-07 から昇格）
 - 2025-11-19 00:33:53 +0900: README とサブタスクを整備し、`internal/game/util` の整理方針検討の下準備が整ったためステータスを[準備完了]へ更新。
- 2025-11-18 22:55 JST: 現状調査を実施し、`internal/game/util` が削除済であること、`internal/game/rng` へ機能移設済であることを確認。
- 2025-11-18 23:05 JST: 再配置方針・命名ガイド追記・Backlog 整理を完了。util 系ディレクトリの再導入禁止方針を docs/KNOWLEDGE/engineering/naming.md へ明文化し、本ストーリーを完了へ移行。

## リスク・懸念
- 既存コードが util 配下のヘルパーに依存している場合、移動に伴う import 調整やテスト修正が必要になる可能性。

## 関連
- PR: #
- Issue: #
- Docs: `docs/KNOWLEDGE/engineering/naming.md`, `docs/architecture/README.md`
