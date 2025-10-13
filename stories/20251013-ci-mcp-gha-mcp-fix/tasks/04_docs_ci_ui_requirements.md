# 04: CI 運用ドキュメント更新（strict/non-strict 運用）

ステータス: [未着手]
担当: @tkg-engineer（想定）
開始: 2025-10-13 16:39:05 +0900

## 目的
- CI における UI ビルドの扱い（非 strict 既定＋別ジョブでの strict 検証）を AGENTS/README に明記し、開発者が判断/再現できるようにする。

## 完了条件（DoD）
- [ ] `AGENTS.md` の「よく使うコマンド」「作業開始ルール」に CI での `MCP_STRICT` 取扱を追記。
- [ ] `README.md` に Linux での依存導入手順（apt パッケージ例）と `MCP_STRICT=1 make check-ui` の確認手順を追記。
- [ ] 変更内容が `make mcp`（docs を含め）で整合している。

## 作業手順（概略）
- 依存一覧と目的（ヘッダ欠如の回避）を簡潔に説明。
- 非 strict/strict の 2 段構えの意図（安定した mcp と品質担保）を明記。

## 成果物リンク
- 変更ファイル: `AGENTS.md`, `README.md`

