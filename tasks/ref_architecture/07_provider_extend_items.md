# 07_provider_extend_items — ItemsTable の Provider 化

## 目的・背景
- アイテム定義も Provider から取得できるようにし、参照経路を統一。

## 作業項目（変更点）
- `internal/game/data.TableProvider` に `ItemsTable() *model.ItemDefTable` を追加。
- `usecase.App` に実装を追加し `SetProvider(app)` で提供。
- scenes のアイテム参照を Provider 経由へ置換。

## 完了条件
- ビルド成功。アイテム一覧/装備で従来通り表示。

## 影響範囲
- data/provider.go, usecase（実装）, scenes/inventory の参照。

## 手順
1) Provider IF拡張→usecase側に実装→`SetProvider(app)` に変更なし。
2) `popup_item.go` などの直接ロードを Provider 参照へ切替。
3) `make mcp` 実行。

## 品質ゲート
- `make mcp`

## 動作確認
- 在庫→アイテムタブで表示崩れ・クラッシュがないこと。

## ロールバック
- IF拡張差分を戻し、直接ロードに戻す。

## Progress / Notes
- YYYY-MM-DD: 着手
- YYYY-MM-DD: 置換完了

## 関連
- `docs/ARCHITECTURE.md` 4章/7章

