# 乖離と論点（2025-09-30）

- Provider の UI 型依存: `internal/game/data.TableProvider` が `uicore.Unit` を返す（UI 型に依存）。
  - 影響: Provider を別UIへ差し替える難度が上がる。テストダブル作成時に UI 依存が伝播。
  - 根拠: `UserUnitByID(id string) (uicore.Unit, bool)` を定義。

- ItemsTable の取得経路が一貫していない: `usecase.App.ItemsTable()` が都度 `model.LoadItemsJSON` を直呼び。
  - 影響: キャッシュ不整合/経路の複数化による保守コスト増。オフライン動作の制約が読みづらい。
  - 根拠: `internal/usecase/data.go` 実装。

- `internal/game/util` という汎用ディレクトリの存在（中身空）。
  - 影響: 命名規約（util/helpers禁止）に反する前例になる。将来の責務混在の温床。
  - 根拠: ルート構成の観測。

- UI メトリクス適用の集中度: `uicore.Metrics` の適用と構造が `apply.go` に集中（既にBacklogあり）。
  - 影響: 変更衝突/追跡性の低下。
  - 根拠: `internal/game/service/ui/apply.go` の肥大化。
