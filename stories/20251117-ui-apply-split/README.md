# 20251117-ui-apply-split — `internal/game/service/ui/apply.go` の分割検討

ステータス: [準備完了]
担当: @tkg-engineer
開始: 2025-11-17 02:11:34 +0900

## 目的・背景
- `internal/game/service/ui/apply.go` に集約されている UI メトリクス適用処理を責務単位で分割し、可読性・変更容易性・テスト容易性を向上させる。
- 現状は List/Status/Sim/Popup/Widgets など複数ドメインへの適用が1ファイルに集中しており、変更差分の把握やコンフリクトが発生しやすい。

## スコープ（成果）
- `internal/game/service/ui/apply.go` の責務を整理し、用途別ファイル（例: `apply_list.go`/`apply_status.go`/`apply_sim.go`/`apply_popup.go`/`apply_widgets.go`）への分割方針を決める。
- 分割後も外部 API（呼び出し元の関数シグネチャや公開 I/F）を極力不変に保つ設計案を用意する。
- 命名規約やアーキテクチャドキュメント（docs/NAMING.md, docs/ARCHITECTURE.md）と整合する構成を検討する。

## 受け入れ基準（Definition of Done）
- [ ] 現状の `apply.go` の責務と依存関係が整理されたメモがある。
- [ ] 分割案（ファイル構成と主な関数配置）が決まり、Pros/Cons が比較できる形でまとめられている。
- [ ] 分割を実施するかどうかの方針と、実施する場合のおおまかな移行ステップが明文化されている（実装自体は別ストーリーで対応してもよい）。

## 工程（サブタスク）
- [ ] 現状調査: `apply.go` の責務と依存関係の棚卸し — `stories/20251117-ui-apply-split/tasks/01_audit-apply-go.md`
- [ ] 分割方針の検討: ファイル構成と責務の再配置案作成 — `stories/20251117-ui-apply-split/tasks/02_design-split-plan.md`
- [ ] 方針決定と後続タスク（実装ストーリー/Backlog）の整理 — `stories/20251117-ui-apply-split/tasks/03_followups-and-backlog.md`

## 計画（目安）
- 見積: 1〜2 セッション
- マイルストン: M1: 現状調査 / M2: 分割案ドラフト / M3: 方針決定

## 進捗・決定事項（ログ）
- 2025-11-17 02:11:34 +0900: ストーリー作成（discovery: 2025-09-29-migrated-05 から昇格）
 - 2025-11-17 02:18:15 +0900: README とサブタスクを整備し、分割方針検討の下準備が整ったためステータスを[準備完了]へ更新。

## リスク・懸念
- 分割粒度を誤ると、かえって関数間の依存が複雑になり、可読性が下がるリスク。
- 既存のテストや呼び出し元との整合を保つための移行コスト。

## 関連
- PR: #
- Issue: #
- Docs: `docs/NAMING.md`, `docs/ARCHITECTURE.md`
