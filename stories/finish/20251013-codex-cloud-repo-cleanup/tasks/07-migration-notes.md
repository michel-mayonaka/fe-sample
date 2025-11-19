# 07 既存開発者向け移行ノート

## 目的
ローカル開発者や既存ブランチが新しいリポジトリ方針に移行する際の手順を提供する。

## 記載項目（ドラフト）
- Make タスクの変更点と置き換え表。
- ベンダリング更新手順（`go mod vendor` 等）。
- 出力ディレクトリ変更に伴う注意（クリーン/キャッシュ）。

## 成功条件（DoD）
- 既存ブランチでの移行に詰まらないこと。

## 実装メモ（結果）
- 追加: `docs/MIGRATION_20251013_CODEX_CLOUD.md`（置き換え表・注意点）。
- 参照: `docs/KNOWLEDGE/ops/codex-cloud.md`, `docs/KNOWLEDGE/ops/offline.md`, `.github/workflows/ci.yml`。
