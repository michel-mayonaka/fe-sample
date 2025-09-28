# 11_docs_sync — ドキュメント同期

## 目的・背景
- 実装変更（Port/Provider/依存図）をドキュメントへ反映し、学習コストと齟齬を減らす。

## 作業項目（変更点）
- `docs/ARCHITECTURE.md` の更新（依存図・ディレクトリ構成の最終化）。
- `docs/API.md` に主要Portの簡単な利用例を追記。
- `README.md` にアーキテクチャへのリンクと開発コマンド（`make mcp`）を追記。

## 完了条件
- リンク切れなし。`rg -n "TODO|WIP" docs` が 0 件（残置は明示）。

## 影響範囲
- docs配下・README。

## 手順
1) 変更に合わせて各mdを更新。
2) `make mcp` 実行。

## 品質ゲート
- `make mcp`

## 動作確認
- 目視でリンク・章立てが正しく参照できる。

## ロールバック
- md差分を戻す。

## Progress / Notes
- YYYY-MM-DD: 着手
- YYYY-MM-DD: 完了

## 関連
- `docs/ARCHITECTURE.md`

