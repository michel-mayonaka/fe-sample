# 20250930-mcp-offline — make mcp オフライン実行対応

ステータス: [完了]
担当: @tkg-engineer

## 目的・背景
- ネットワーク遮断環境（サンドボックス/機内/社内制限）でも `make mcp`（vet/build/lint/test）を完走させ、ローカル検証を安定化する。

## スコープ（成果）
- `MCP_OFFLINE=1 make mcp` が `.gomodcache` 空かつネットワーク未接続でも成功する。
- モジュール解決は `vendor/` に固定し、`go` コマンドは `-mod=vendor` を使用。
- `vendor-sync` ターゲットを追加し、依存更新フローを明文化。
- `lint` は未導入でもスキップ可（既存挙動を維持）。

- [x] `vendor/` をリポジトリに追加し、`MCP_OFFLINE=1 make mcp` がネットワーク不要で実行可能（テスト結果は現状のロジックに依存）。
- [x] `Makefile` に `MCP_OFFLINE` 分岐と `vendor-sync` を追加（`go mod vendor`）。
- [x] 検証手順: `.gomodcache` を空にせずとも `GOPROXY=off MCP_OFFLINE=1 make mcp` が `vet/build` まで完走。
- [x] ドキュメント: `README.md` と `docs/KNOWLEDGE/ops/offline.md` にオフライン手順を追記。

## 工程（サブタスク）
- [ ] 設計: オフライン方針と適用範囲（`pkg/...`/UI） — `tasks/01_outline.md`
- [ ] 実装: `Makefile` に `MCP_OFFLINE` 分岐追加 — `tasks/02_impl_makefile.md`
- [ ] 実装: `vendor-sync` 追加と運用メモ — `tasks/03_vendor_sync.md`
- [ ] ドキュメント: 手順追記/Backlog 整合 — `tasks/04_docs.md`
- [ ] 検証: オフライン再現テスト — `tasks/05_offline_verify.md`

## 計画（目安）
- 見積: 1.5〜2.0 時間 / 1 セッション
- マイルストン: M1 設計合意 → M2 Makefile 反映 → M3 検証/ドキュメント

## 進捗・決定事項（ログ）
- 2025-09-30: ストーリー作成、方針は `vendor/` 採用で合意見込み。

## リスク・懸念
- 初回 `vendor/` 生成には一度オンライン環境が必要。
- UI 側ビルドは環境依存でスキップ分岐あり（既存挙動を尊重）。

## 関連
- Docs: `docs/REF_STORIES.md`, `README.md`, `.golangci.yml`, `Makefile`
