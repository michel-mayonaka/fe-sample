# 01_audit_current_vs_docs — 現状とドキュメントの比較監査

ステータス: [完了]

## 目的
- 現行コードと `docs/architecture/README.md` の不整合を体系的に洗い出す。

## 作業概要
- レイヤ/境界: Scene / Service / Provider / Usecase / Model / DataPort の役割と依存方向を確認。
- ライフサイクル: `input→intent→advance→render` の実装有無とズレを抽出。
- 永続処理: 直接書込がないか、`Env.Data` 経由に統一されているかを確認。
- 命名規約: `docs/KNOWLEDGE/engineering/naming.md` との不一致（例: util/helpers/common 名）を確認。

## 手順（チェックリスト）
- [ ] 対象範囲の決定（internal/game/*, pkg/game, docs/*）
- [ ] 依存方向マップの作成（簡易リスト/ASCII 図）
- [ ] 不整合リストの作成（根拠ファイル/行参照付き）
- [ ] 修正方針（採用/却下/保留）を暫定記載

## 成果物
- `artifacts/audit_findings.md`（不整合一覧と提案）

## DoD
- [x] 不整合一覧がレビュアブルな形で揃う
- [x] 次タスクで必要な決定事項の草案がある
