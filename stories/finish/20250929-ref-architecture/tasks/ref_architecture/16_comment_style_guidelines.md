# 16_comment_style_guidelines — コメント記法の統一方針

## 目的・背景
- GoDoc 準拠の最小コメント規約を明文化し、lint（revive）指摘の方向性と docs/コードの責務分担を揃える。

## 作業項目（変更点）
- `docs/KNOWLEDGE/engineering/comment-style.md` を新規追加（以下を網羅）。
  - パッケージコメント必須（1 行目に役割を簡潔に）。
  - 公開識別子は「X は …」で 1 行要約（日本語）。
  - 詳細は docs 側（ARCHITECTURE/API など）に寄せ、コード内は 1–2 行を上限。
  - 例外: インターフェース契約や注意事項（データ競合・集約境界）は 2–3 行で補足可。
  - JSDoc ライクな @param/@return は原則使用しない（GoDoc の慣習に従う）。
  - 日本語原則。英語が一般名詞のときは併記可（例: Transaction, Provider）。
- 主要パッケージ（例: `internal/game/scenes`, `internal/game/service/ui`, `internal/repo/sqlite`）にパッケージコメントを追記。

## 完了条件
- `docs/KNOWLEDGE/engineering/comment-style.md` が存在し、合意されたルールが明記されている。
- 最低 3 パッケージにパッケージコメントを追加し、revive の package-comments が減少。

## 影響範囲
- docs と一部パッケージの先頭コメントのみ（挙動変更なし）。

## 手順
1) `docs/KNOWLEDGE/engineering/comment-style.md` を追加。
2) 代表パッケージにパッケージコメントを追加（`scenes`/`ui`/`repo/sqlite`）。
3) `make mcp` でビルドを確認。lint の完全解消は別途。

## 品質ゲート
- `make mcp`

## 動作確認
- なし（コメントのみ）。

## ロールバック
- 追加ファイルを削除、既存コメントは残して問題なし。

## Progress / Notes
- 2025-09-28: タスク追加
- 2025-09-28: `docs/KNOWLEDGE/engineering/comment-style.md` を追加、代表パッケージ（scenes/inventory, scenes/sim, scenes, render, util）にパッケージコメントを付与、mcp通過

## 関連
- `docs/architecture/README.md`（コメントとドキュメントの役割分担）
