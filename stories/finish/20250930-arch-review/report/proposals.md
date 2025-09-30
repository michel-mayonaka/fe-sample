# 改善提案ドラフト（アーキ見直し含む）

以下は優先度（高→低）、効果、概算コストを併記した提案です。

1) Provider の UI 型依存を解消（Port/Adapter 再整理）
- 目的: `internal/game/data.TableProvider` を純粋な参照Portにし、UI 型への依存を削除。
- 変更案:
  - `UserUnitByID` を Provider から外し、`service/ui` の `adapter` に配置（`UnitFromUser` を利用）。
  - Scenes は `Provider.UserTable().Find(id)` → `uicore.UnitFromUser(c)` の経路に統一。
- 効果: 依存方向の明確化、テストダブル容易化、他UI移植性の向上。
- 概算コスト: 小〜中（呼び出し箇所の置換 + テスト調整）。

2) ItemsTable 取得を Repo 化し、キャッシュ一貫化
- 目的: `usecase.App.ItemsTable()` の直読みを廃止し、Repo/キャッシュに一本化。
- 変更案:
  - `repo.ItemsRepo` の導入、または `WeaponsRepo` に隣接する `ItemDefTable` キャッシュを追加。
  - `usecase.App.ItemsTable()` は Repo 経由に切替。
- 効果: I/O 経路の統一、オフライン運用の明確化、性能の安定。
- 概算コスト: 中（Repo 追加 + 注入配線 + テスト）。

3) `internal/game/util` の撤去または責務特化へ改称
- 目的: 命名規約順守と将来の責務混在回避。
- 変更案:
  - 空なら削除。必要なら `internal/game/{rng,rect,debug}` 等に明確化して再配置。
- 効果: 規約一貫性、探索性向上。
- 概算コスト: 極小。

4) 入力レイアウトの外部化（設定駆動）
- 目的: キー/マウス割当を `config` で切替可能にし、`provider/input/ebiten` のレイアウト注入を標準化。
- 変更案:
  - `config/input_layout.json` を追加し、`app.NewUIAppGame` で読み込み→`ginput.Layout` を生成。
- 効果: カスタマイズ性向上、将来的なデバイス拡張の土台。
- 概算コスト: 小。

5) UI メトリクス適用の分割（apply.go 分解）
- 目的: 既存Backlogの具体化。責務単位で分割し変更衝突を低減。
- 変更案: `apply_list.go`/`apply_status.go`/`apply_sim.go`/`apply_popup.go`/`apply_widgets.go` に分割。
- 効果: 可読性/保守性/衝突低減。
- 概算コスト: 中。
