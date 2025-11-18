# 03_テスト/ドキュメント/CI — `apply_test.go` 拡張と `make mcp`

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-19 01:36:42 +0900

## 目的
- ドメイン別 apply 関数のテスト網羅とドキュメント更新を行い、CI (`make mcp`) で回帰がないことを確認する。

## 完了条件（DoD）
- [ ] `apply_test.go` もしくは追加の `_test.go` に、`metricsTargets` を使ったドメイン別単体テストが追加されている。
- [ ] `docs/API.md` と `docs/ARCHITECTURE.md` が分割後の構成を説明している。
- [ ] `make mcp` が成功し、結果ログがタスクに記録されている。

## 作業手順（概略）
1. テスト: `DefaultMetrics` スナップショット + `metricsTargets` をモック化し、List/Status/Sim 等の適用を検証するケースを追加。
2. ドキュメント: `ApplyMetrics` の所在ファイルや分割構成について説明し、既存の `apply.go` 参照を更新。
3. CI: `make mcp` を実行し、`cmd/ui_sample` のホットリロード経路チェック結果をログに残す。

## 進捗ログ
- 2025-11-19 01:36:42 +0900: タスク作成。

## 依存／ブロッカー
- Task 02（ドメイン分割）完了後にテスト/ドキュメントを更新する。

## 成果物リンク
- `make mcp` ログ/証跡: （後続で記載）
