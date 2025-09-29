# 10_remove_aggregated_usecases — 合成UseCasesの段階的撤去

## 目的・背景
- Env/Scene の依存を個別Portへ完全移行し、合成 `UseCases` を撤去して最小依存を徹底。

## 作業項目（変更点）
- `Env.App`（合成）を参照している箇所を `Env.Data/Battle/Inv` に全置換。
- 置換完了後に `Env.App` フィールドと関連インターフェースを削除。

## 完了条件
- `rg -n "\.App\." internal/game/scenes` が 0 件。
- ビルド成功、主要画面の動作確認OK。

## 影響範囲
- scenes配下全体、env.go、app/core.go（配線）。

## 手順
1) 検索で参照箇所洗出し→Sceneごとに置換。
2) 置換後に `env.go` の `UseCases`/`App` を削除。
3) `core.go` から `App` 注入を削除し、Port注入のみ残す。
4) `make mcp`、起動確認。

## 品質ゲート
- `make mcp`

## 動作確認
- 一覧/ステータス/在庫/（戦闘導線があれば）戦闘が従来通り動く。

## ロールバック
- フィールド削除前のコミットへ戻す（段階的に進める）。

## Progress / Notes
- 2025-09-28: 着手
- 2025-09-28: 置換/削除完了（Env.App/UseCases撤去、core配線更新、`go build` OK）

## 関連
- `docs/ARCHITECTURE.md` 12.4
