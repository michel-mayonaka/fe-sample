# 12_cleanup_legacy_docs — 旧ドキュメント整理

## 目的・背景
- 重複/旧版ドキュメント（例: `docs/architecture.md` 小文字）を整理し、参照先を一本化。

## 作業項目（変更点）
- 旧版の削除または残置時は先頭に非推奨バナーを追加し新アーキテクチャへリンク。
- READMEなどのリンクを新 `ARCHITECTURE.md` に統一。

## 完了条件
- `docs/architecture.md` の扱いが決定（削除 or 非推奨化）。
- 既存リンクが404にならない。

## 影響範囲
- docs配下・README/その他リンク元。

## 手順
1) 参照箇所検索 `rg -n "docs/architecture.md|ARCHITECTURE.md"`。
2) 方針に従い削除/非推奨バナー追記/リンク更新。
3) `make mcp` 実行。

## 品質ゲート
- `make mcp`

## 動作確認
- 目視でリンク遷移が期待通り。

## ロールバック
- 削除したファイルを復元、またはリンクを元に戻す。

## Progress / Notes
- 2025-09-28: 着手
- 2025-09-28: 完了（小文字版なしを確認・リンクは `ARCHITECTURE.md` に統一済み）

## 関連
- `docs/ARCHITECTURE.md`
