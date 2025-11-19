# 02_ドメイン別ファイル分割 — List/Status/Sim/Popup/Widgets

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 01:36:42 +0900

## 目的
- `apply.go` に残る大部分の代入処理を、責務単位の `apply_list.go` / `apply_status.go` / `apply_sim.go` / `apply_popup.go` / `apply_widgets.go` へ切り出す。

## 完了条件（DoD）
- [x] 各新ファイルに `func apply<List|Status|Sim|Popup|Widgets>(t *metricsTargets, m Metrics)` が実装され、`ApplyMetrics` から呼び出されている。
- [x] ファイルごとに GoDoc コメント（日本語1行）を付与し、ビルドタグ `//go:build !headless` を付けている。
- [x] `go test ./internal/game/service/ui` が成功し、`apply.go` には公開 API と `ApplyMetrics` の呼び出しのみが残っている。

## 作業手順（概略）
1. List/Line セクションを `apply_list.go` に移動し、ヘルパで代入。
2. Status 系・Sim 系（Terrain/Preview 含む）・Popup・Widgets を順番に移設。
3. `ApplyMetrics` を読みやすい順序（List→Status→Sim→Popup→Widgets）で内部関数呼び出しに置換。

## 進捗ログ
- 2025-11-19 01:36:42 +0900: タスク作成。
- 2025-11-19 09:50:00 +0900: `apply_list.go`/`apply_status.go`/`apply_sim.go`/`apply_popup.go`/`apply_widgets.go` を追加、各ファイルに GoDoc＋ビルドタグを付与。
- 2025-11-19 10:05:00 +0900: `ApplyMetrics` をセクション呼び出しのみに縮約し、`go test ./internal/game/service/ui` で確認。

## 依存／ブロッカー
- Task 01（ヘルパ導入）完了後に着手するとスムーズ。

## 成果物リンク
- コード: `internal/game/service/ui/apply_list.go`, `internal/game/service/ui/apply_status.go`, `internal/game/service/ui/apply_sim.go`, `internal/game/service/ui/apply_popup.go`, `internal/game/service/ui/apply_widgets.go`
