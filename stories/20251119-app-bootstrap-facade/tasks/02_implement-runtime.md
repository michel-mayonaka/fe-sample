# 02_実装 — app.Runtime/Run の追加

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-19 02:06:44 +0900

## 目的
- 設計済みファサードに沿って `internal/app` 配下へ `bootstrap`/`runtime` モジュールを追加し、Repo/Usecase/Env 初期化を集中させる。

## 完了条件（DoD）
- [ ] `internal/app/bootstrap`（仮）で repo/usecase/metrics 初期化が行われ、`Game` 生成まで完結する。
- [ ] `cmd/ui_sample/main.go` が新 API を呼ぶだけの構成になっている（旧 `internal/game/app` への直接依存が無い）。
- [ ] `make mcp` が通り、UI サンプルが従来同等に起動できることを確認。

## 成果物
- 実装済みコードと差分リンク
- 動作確認ログ

## 進捗ログ
- 2025-11-19 02:06:44 +0900: タスク登録。
