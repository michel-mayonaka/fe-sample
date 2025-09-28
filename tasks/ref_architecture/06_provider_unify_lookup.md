# 06_provider_unify_lookup — Provider参照の統一チェック

## 目的・背景
- UI からのテーブル参照を `gdata.Provider` に統一し、取得経路を一本化。

## 作業項目（変更点）
- `WeaponsTable` などの参照が Provider 経由になっているか全件確認・是正。

## 完了条件
- `rg -n "WeaponsTable\(" internal` の結果が scenes→Provider 経由のみであること。
- ビルド成功。

## 影響範囲
- scenes配下の参照箇所。

## 手順
1) `rg -n "WeaponsTable\(" internal` で現状把握。
2) `Env`/`UseCases` 経由の参照があれば Provider 経由へ修正。
3) `make mcp` 実行。

## 品質ゲート
- `make mcp`

## 動作確認
- 表示・プレビューが従来通り機能。

## ロールバック
- 変更前の参照に戻す（局所差分）。

## Progress / Notes
- YYYY-MM-DD: 着手
- YYYY-MM-DD: 統一完了

## 関連
- `docs/ARCHITECTURE.md` 2章/7章

