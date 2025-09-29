# Known Issues — 最終棚卸し (2025-09-29)

本ファイルは 14_final_audit の結果、既知の課題と今後の対応方針を集約します。

## 優先度: 高
- テスト不足: ユースケース層 `usecase.App`
  - 現状: `./internal/usecase` カバレッジ 42.9%。
  - 対応: `RunBattleRound` のHP/耐久/保存呼び出しを検証するテストを追加し 60% 目標へ。
  - 進捗: `ReloadData`/`PersistUnit` のテストを追加済み。

## 完了（記録）
- `/pkg/game` のカバレッジ強化は完了
  - 以前: 48.9% → 現在: 87.2%（`ForecastAtExplain`/三すくみ/地形/AS境界を追加）。

## 優先度: 中
- Lint 実行の環境依存エラー
  - 事象: `golangci-lint` が `context loading failed: no go files to analyze` を出力。
  - 補足: ローカル実行は非必須扱い（Makefile でメッセージ化）。CI 導入時に再検証。
  - 対応案: CI では `-mod=readonly` と `GOWORK=off` を明示、`go mod tidy` を前段に挿入。
- UI スナップショットテスト未整備
  - 対応: 主要ビュー（ステータス/在庫/戦闘プレビュー）の最小スナップショット仕組みを `testdata/` と併用して導入（描画依存は条件付きスキップ）。

## 優先度: 低（情報）
- 実行環境に JRE がない場合のコンソール出力
  - 事象: 一部環境で `Unable to locate a Java Runtime` が出力される。
  - 補足: ビルド/テストへの影響は無し（Makefile で許容）。

---
更新ルール: issue 解消時に本ファイルから削除し、PR/コミットに根拠リンクを残すこと。
