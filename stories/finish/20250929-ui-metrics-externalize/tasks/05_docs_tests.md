# 05 — ドキュメント更新と簡易テスト

ステータス: [完了]

## 目的
- 仕様をドキュメントに反映し、最低限のユニットテストでリグレッションを抑止する。

## 対象
- `docs/architecture/README.md`（UIメトリクスの所在と読み込みタイミング）
- `docs/SPECS/reference/api.md`（必要なら Ctx/サービスの項に言及）
- `stories/BACKLOG.md`（当該エントリを整理）
- テスト: `internal/config/uimetrics`（ロード/フォールバック/優先順位）
- テスト: `internal/game/service/ui`（`ApplyMetrics` の部分適用）

## 成果物
- ドキュメント差分（反映済）
- 追加テスト:
  - `internal/config/uimetrics/metrics_test.go`
  - `internal/game/service/ui/apply_test.go`
- `make mcp` グリーン（-race -cover 既定）

## 進捗ログ
- 2025-09-29: 雛形作成
- 2025-09-29: 代表画面への適用完了。ドキュメント反映作業に着手。
- 2025-09-29: ローダ/適用テストを追加し単体で緑を確認。
