# [task] 参照掃除・Lint・ドキュメント更新

- ステータス: 未着手
- 目的: 参照の取りこぼしを解消し、Lint/テストを全通過。ドキュメントへ反映する。

## スコープ
- 参照掃除: `rg scenes/helper|PointIn` が 0 件であることの確認。
- Lint/ビルド/テストの通過（`make mcp`）。
- ドキュメント更新: `docs/NAMING.md` に `helper` 禁止と代替命名例の追記、`docs/ARCHITECTURE.md` に層の責務例追記。

## 非スコープ
- 新機能追加。

## 手順
- `rg` による残存参照の撲滅（import/識別子/コメント）。
- Lint/テスト実行、警告是正。
- ドキュメントの差分作成・リンク確認。

## DoD（完了条件）
- `rg -n "scenes/helper|\bPointIn\b"` が 0 件。
- `make mcp` グリーン。
- 上記ドキュメントの該当節が更新済み。

## コマンド例
- `rg -n "scenes/helper|\bPointIn\b" || true`
- `make mcp`

