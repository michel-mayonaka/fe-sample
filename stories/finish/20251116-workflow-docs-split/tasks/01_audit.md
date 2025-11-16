# 01_audit — WORKFLOW 現状棚卸し

## 目的
- 現行 `docs/WORKFLOW.md` の構成と内容を洗い出し、どの論点をどのファイルへ分割すべきかを把握する。

## 作業項目
- セクション構成と見出しを一覧化する。
- 情報の種類（ストーリー運用/CI/ローカル開発 など）ごとにグルーピングする。
- 分割時に重複・欠落が起きそうな箇所をメモする。

## アウトプット
- 棚卸しメモ（簡単な箇条書きで可）。
- 次タスク（02_design）の入力となる論点リスト。

## 棚卸しメモ
- セクション構成（現状）
  - タイトル: 「開発者向けワークフロー（最小ガイド）」
  - Codex（AIコーディングエージェント）
    - ストーリーの作成
      - stories/BACKLOG.md からのストーリー作成方法
      - 突発ストーリー作成時のプロンプト例（Codex への依頼文）
      - ストーリー運用の詳細は `docs/REF_STORIES.md` を参照する旨
      - Codex と Vibe‑kanban の役割分担（Codex: ストーリー〜コミット、Vibe‑kanban: 実際の作業）
      - ブランチ運用の基本方針（コード変更は `feat/<slug>`、ドキュメント/ストーリーのみは master 可）
    - stories/discovery の昇格について
      - Vibe‑kanban からの課題提案が `stories/discovery/` に起票されること
      - accepted/declined への移動と `stories/BACKLOG.md` への反映
      - 昇格したストーリーの FROM_DISCOVERY 運用（consumed への退避と Backlog からの削除）
      - Codex 直依頼ストーリーは discovery を持たないケースがあること
  - ローカル検証（codex-cloud/オフライン対応）
    - `make smoke` / `MCP_OFFLINE=1 make offline` / `make mcp` / `MCP_STRICT=1 make check-ui` の役割
    - 出力先ディレクトリ（`out/bin`, `out/logs`, `out/coverage`）
    - 詳細は `docs/CODEX_CLOUD.md` 参照
  - CI 構成（概要）
    - ジョブ構成: `smoke-offline` → `build-and-lint` → `ui-build-strict`
    - 各ジョブの目的（smoke/offline, build+lint, UI 厳格）
    - ワークフロー定義ファイル: `.github/workflows/ci.yml`
  - Vibe‑kanban
    - ストーリーの作業開始方法（ローカルの Vibe‑kanban URL 例付き）
    - ストーリー slug をタイトルに含めて作業開始する例
    - 作業中に発生した提案や分割タスクが Vibe‑kanban 内の Codex により `stories/discovery/` に追加されること
    - 作業完了後のレビューとマージの流れ（高レベル）

## 論点のグルーピング（種別）
- ストーリー/Discovery/Backlog 運用（Codex + `docs/REF_STORIES.md` 連携）
- ローカル開発/検証フロー（make コマンド群と codex-cloud）
- CI フロー（GitHub Actions の全体像）
- Vibe‑kanban を用いたタスク運用（ツール連携）

## 分割時の注意点メモ
- `docs/REF_STORIES.md` と「Codex/ストーリー作成」の説明が重複気味なので、ストーリー運用の詳細は workflow/stories 側に集約し、WORKFLOW の「Codex」節は入口程度にとどめる方向が良さそう。
- ローカル検証と CI の説明は、将来的に `docs/workflows/local-dev.md` と `docs/workflows/ci.md` 等に分ける候補。
- Vibe‑kanban の説明はツール/運用寄りで、ストーリー運用やローカル開発と横断するため、どのファイルからリンクするかを 03_policy で整理する必要がある。
