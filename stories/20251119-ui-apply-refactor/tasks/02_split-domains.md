# 02_ドメイン別ファイル分割 — List/Status/Sim/Popup/Widgets

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-19 01:36:42 +0900

## 目的
- `apply.go` に残る大部分の代入処理を、責務単位の `apply_list.go` / `apply_status.go` / `apply_sim.go` / `apply_popup.go` / `apply_widgets.go` へ切り出す。

## 完了条件（DoD）
- [ ] 各新ファイルに `func apply<List|Status|Sim|Popup|Widgets>(t *metricsTargets, m Metrics)` が実装され、`ApplyMetrics` から呼び出されている。
- [ ] ファイルごとに GoDoc コメント（日本語1行）を付与し、ビルドタグ `//go:build !headless` を付けている。
- [ ] `go test ./internal/game/service/ui` が成功し、`apply.go` には公開 API と `ApplyMetrics` の呼び出しのみが残っている。

## 作業手順（概略）
1. List/Line セクションを `apply_list.go` に移動し、ヘルパで代入。
2. Status 系・Sim 系（Terrain/Preview 含む）・Popup・Widgets を順番に移設。
3. `ApplyMetrics` を読みやすい順序（List→Status→Sim→Popup→Widgets）で内部関数呼び出しに置換。

## 進捗ログ
- 2025-11-19 01:36:42 +0900: タスク作成。

## 依存／ブロッカー
- Task 01（ヘルパ導入）完了後に着手するとスムーズ。

## 成果物リンク
- PR/コミット: （後続で記載）
