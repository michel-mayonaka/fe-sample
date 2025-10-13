# 05: golangci-lint v2 移行と lint 対象の限定

ステータス: [完了]
担当: @tkg-engineer（想定）
開始: 2025-10-13 21:18:00 +0900

## 目的
- `.golangci.yml (version: 2)` と整合するよう、CI で golangci-lint v2 を用い、UI 依存で壊れないよう lint 対象をロジック層（`./pkg/... ./internal/usecase`）へ限定する。

## 完了条件（DoD）
- [x] GHA で v2 のバイナリを明示インストール。
- [x] Makefile の `lint-ci` をロジック層に限定。
- [x] `build-and-lint` で UI 依存が原因の typecheck 失敗を回避。

## 変更内容
- `.github/workflows/ci.yml`: `github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest` をインストール。
- `Makefile`: `LINT_PKGS ?= ./pkg/... ./internal/usecase` を導入し、`lint-ci` に適用。
- `.golangci.yml`: `run.build-tags: headless` を追加し、ヘッドレス型定義（`uicore`）で型チェック可能に。

## 作業ログ
- 2025-10-13 21:18:30 +0900: v2 を明示インストールし、バージョン表示で確認（IOP=++）
- 2025-10-13 21:22:00 +0900: lint 対象をロジック層に限定、push 済み（IOP=++）

## 成果物リンク
- `.github/workflows/ci.yml`
- `Makefile`
