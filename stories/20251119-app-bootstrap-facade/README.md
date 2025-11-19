# 20251119-app-bootstrap-facade — internal/app ブートストラップファサード実装

ステータス: [準備中]
担当: @tkg-engineer
開始: 2025-11-19 02:06:44 +0900

## 目的・背景
- `cmd/ui_sample/main.go` が `internal/game/app` へ直接依存し、乱数/Repo/Env 初期化をすべて抱えている現状を是正する。
- `internal/app` にアプリ起動の責務を集約し、main 側は設定注入＋ `app.Run(cfg)` を呼ぶだけの構成にすることでレイヤリング整合を保つ（stories/20251119-layering-main-to-app の Phase1）。

## スコープ（成果）
- `internal/app` 配下に公開ファサード（仮: `bootstrap.Config`, `runtime.Runtime`）を追加し、repo/usecase/metrics/scene 初期化を一箇所で完了させる。
- `cmd/ui_sample` から `internal/game/app` への直接 import を排除し、`make mcp` が通る状態で新 API を適用する。
- docs/architecture/README.md などの構成記述を更新し、開発者が新しい起動フローを参照できるようにする。

## 受け入れ基準（Definition of Done）
- [ ] `cmd/ui_sample/main.go` の責務が「設定収集→app.Run」で完結し、アプリ固有の初期化コードが残っていない。
- [ ] 新しい `internal/app` ファサードが repo/usecase/Env/Session/metrics 初期化とウィンドウ設定を担い、`make mcp` がグリーンである。
- [ ] docs/architecture/README.md（および関連ドキュメント）が最新のレイヤリング/起動フローを説明している。

## 工程（サブタスク）
- [ ] 設計: `internal/app` ファサードの API と責務整理 — `stories/20251119-app-bootstrap-facade/tasks/01_design-bootstrap-facade.md`
- [ ] 実装: runtime/bootstrap 追加と main 以外からの利用準備 — `stories/20251119-app-bootstrap-facade/tasks/02_implement-runtime.md`
- [ ] 移行: main 差し替えと docs 更新 — `stories/20251119-app-bootstrap-facade/tasks/03_migrate-main-and-docs.md`

## 計画（目安）
- 見積: 2 セッション（設計1 / 実装1）
- マイルストン: M1 設計合意 → M2 runtime 実装 → M3 main/doc 更新

## 進捗・決定事項（ログ）
- 2025-11-19 02:06:44 +0900: ストーリー作成（Backlog: internal/app bootstrap ファサード実装 から昇格）

## リスク・懸念
- `internal/game/app` からファイル移動を行うため、大量差分になりコンフリクトが発生しやすい。
- Ebiten 依存をラップする際に headless ビルドとの両立を崩す可能性がある。

## 関連
- PR: #
- Issue: #
- Docs: `docs/architecture/README.md`, `docs/KNOWLEDGE/ops/codex-cloud.md`
