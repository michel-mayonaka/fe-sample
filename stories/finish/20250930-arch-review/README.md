# 20250930-arch-review — アーキテクチャ再レビューと現況確認＋BACKLOG整備

ステータス: [完了]
担当: @tkg-engineer

## 目的・背景
- 設計ドキュメント（`docs/architecture/README.md` 等）と現行実装の乖離を可視化し、改善候補を抽出して Backlog を最新化する。
- 今後の変更衝突・リワークを抑制し、優先度付きの改善計画に反映する。

## スコープ（成果）
- 現況レポート（依存関係/境界/責務/データフローの要約）: `stories/20250930-arch-review/report/current_state.md`
- 乖離・論点の一覧（根拠つき）: `report/gaps_and_issues.md`
- 改善提案と対応案のドラフト: `report/proposals.md`
- アーキテクチャ見直し提案を含む（境界/依存方向/パッケージ再配置/Provider↔Portの整理など）
- Backlog 追記（調査で得た改善点/提案を3件以上、優先度順で）
- 最小のドキュメント整合（命名/アーキ/公開APIの差分を最小修正）

- [ ] `stories/20250930-arch-review/report/current_state.md` を作成し、主要コンポーネント/依存/データフローを記述。
- [ ] 乖離と論点を `report/gaps_and_issues.md` に列挙（影響/優先度/根拠を付与）。
- [ ] 改善提案（アーキテクチャ見直しを含む）を `report/proposals.md` に整理し、Backlog に3件以上を新規追加（目的/背景/DoD/関連を含む）。
- [ ] 必要に応じて `docs/architecture/README.md`/`docs/KNOWLEDGE/engineering/naming.md`/`docs/SPECS/reference/api.md` を最小修正。
- [ ] `make mcp` がグリーン（オフライン対応時は `MCP_OFFLINE=1` も確認）。

## 工程（サブタスク）
- [ ] 01: 対象範囲の棚卸し（ディレクトリ/依存の洗い出し） — `tasks/01-inventory.md`
- [ ] 02: ドキュメントと実装の比較 — `tasks/02-compare-docs.md`
- [ ] 03: 乖離・リスクの抽出 — `tasks/03-gaps-risks.md`
- [ ] 04: 改善提案の整理とBacklog追記 — `tasks/04-proposals-backlog.md`
- [ ] 05: 最小のドキュメント修正 — `tasks/05-doc-sync.md`

## 計画（目安）
- 見積: 4–6 時間（1セッション）
- マイルストン: M1 現況レポート / M2 提案+Backlog追記 / M3 ドキュメント最小修正

## 進捗・決定事項（ログ）
- 2025-09-30: ストーリー作成・雛形整備（Backlog昇格）

## リスク・懸念
- 例: 依存の変更、CI制約 など

## 関連
- PR: #
- Issue: #
- Docs: `docs/...`
