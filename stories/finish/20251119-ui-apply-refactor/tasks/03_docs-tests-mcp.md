# 03_テスト/ドキュメント/CI — `apply_test.go` 拡張と `make mcp`

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 01:36:42 +0900

## 目的
- ドメイン別 apply 関数のテスト網羅とドキュメント更新を行い、CI (`make mcp`) で回帰がないことを確認する。

## 完了条件（DoD）
- [x] `apply_test.go` もしくは追加の `_test.go` に、`metricsTargets` を使ったドメイン別単体テストが追加されている。
- [x] `docs/SPECS/reference/api.md` と `docs/architecture/README.md` が分割後の構成を説明している。
- [x] `make mcp` が成功し、結果ログがタスクに記録されている。

## 作業手順（概略）
1. テスト: `DefaultMetrics` スナップショット + `metricsTargets` をモック化し、List/Status/Sim 等の適用を検証するケースを追加。
2. ドキュメント: `ApplyMetrics` の所在ファイルや分割構成について説明し、既存の `apply.go` 参照を更新。
3. CI: `make mcp` を実行し、`cmd/ui_sample` のホットリロード経路チェック結果をログに残す。

## 進捗ログ
- 2025-11-19 01:36:42 +0900: タスク作成。
- 2025-11-19 10:30:00 +0900: `apply_test.go` に `capture*Targets` ヘルパと List/Status/Sim/Popup/Widgets の個別テスト、`assertSliceCopied` を追加し `go test ./internal/game/service/ui` で通過。
- 2025-11-19 10:45:00 +0900: `docs/SPECS/reference/api.md`/`docs/architecture/README.md` を新構成へ更新し、`make mcp` を完走（`check-ui` 含む）したログを取得。

## 依存／ブロッカー
- Task 02（ドメイン分割）完了後にテスト/ドキュメントを更新する。

## 成果物リンク
- テスト: `internal/game/service/ui/apply_test.go`
- ドキュメント: `docs/SPECS/reference/api.md`, `docs/architecture/README.md`
- `make mcp` ログ/証跡: 2025-11-19 10:45 +0900 実行分（ローカル `make mcp` 成功）
