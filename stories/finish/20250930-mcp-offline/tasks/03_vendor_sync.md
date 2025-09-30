# 03 実装: vendor-sync ターゲット追加と運用

ステータス: [未着手]

## 目的
- 依存更新時に `vendor/` を最新化し、レビュー/CI での再現性を確保する。

## 作業
- `Makefile` に `vendor-sync` を追加（`go mod tidy && go mod vendor`）。
- `README.md` に「依存更新フロー」を追記（更新→`make vendor-sync`→PR）。

## 注意
- 初回ベンダリングはオンライン環境で実行が必要。

## DoD
- `make vendor-sync` が成功し、`vendor/` が生成/更新される。

## 進捗ログ
- 2025-09-30: 作成。

