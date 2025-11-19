# [story] scenes/list_layout のUI層への抽出

ステータス: [完了]
- 日付: 2025-09-29
- 参照: docs/architecture/README.md, docs/KNOWLEDGE/engineering/naming.md

## 背景
`scenes/list_layout` には再利用可能な UI レイアウト処理が含まれ、シーン層（オーケストレーション）から分離した方が責務が明確になる。

## 目的
- レイアウトを UI コンポーネント/レイアウト層へ移し、シーンは状態遷移と演出の調停に専念できるようにする。

## スコープ
- 新パッケージの新設: `internal/game/ui/layout/list`（名称は既存APIに合わせ最終決定）。
- 既存 `list_layout` の型/関数/リソースの移動・命名整備。
- 呼び出し側の import/参照更新。

## 非スコープ
- 新規UIデザイン/仕様変更。
- 入出力デバイス処理の拡張。

## 成果物 / DoD
- `scenes` 直下からレイアウト処理が排出され、`internal/game/ui/layout/list` に集約されている。
- 既存シーンは UI レイアウトAPIを経由して利用し、重複コードが削減されている。
- `make mcp` 成功、主要シーンの動作回帰テストがグリーン。
- docs/architecture/README.md に層の説明と参照例を追記。

## 影響範囲
- `scenes/*`（参照更新）
- `internal/game/ui/*`（新設）

## リスクと対策
- リスク: レイアウト寸法/座標の意図せぬ変更。
  - 対策: 既存シーンとのピクセル差分確認、Golden スクリーンショットがあれば比較。

## 計測/検証
- 重複行数の削減、`rg` による `list_layout` 参照の一本化。

## 次アクション（タスク化方針）
- 01_API 抽出と命名確定
- 02_実装移動＋ビルド
- 03_呼び出し側更新
- 04_回帰確認＋ドキュメント更新
