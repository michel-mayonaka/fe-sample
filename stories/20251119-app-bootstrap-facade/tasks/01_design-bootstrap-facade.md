# 01_設計 — app/bootstrap ファサード方針固め

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-19 02:06:44 +0900

## 目的
- `internal/app` で管理すべき初期化責務（Repo, Usecase, Env, Session, Metrics, Window設定）を洗い出し、公開 API の構成を決める。

## 完了条件（DoD）
- [ ] Config 構造体と公開関数（例: `NewRuntime`, `Run`）のインタフェース案が記述されている。
- [ ] 依存挿入の流れ（config→usecase→game/app）とエラーハンドリング方針が決まっている。
- [ ] docs/architecture/README.md との整合点や更新対象がリストアップされている。

## 成果物
- 設計メモ（このファイル末尾 or README 参照）
- 追随が必要なドキュメント一覧

## 進捗ログ
- 2025-11-19 02:06:44 +0900: タスク登録。
