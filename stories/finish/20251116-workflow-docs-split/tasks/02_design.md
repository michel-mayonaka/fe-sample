# 02_design — docs/workflows/ 構成案

## 目的
- `docs/workflows/` ディレクトリのファイル構成（ファイル名と役割）を設計し、レビュー可能な案としてまとめる。

## 作業項目
- 01_audit の棚卸し結果をもとに、主要な論点ごとのファイル候補を列挙する。
- 命名規約（`docs/NAMING.md`）との整合を確認しつつ、ファイル名を決める。
- 将来の拡張を見据えて、無理のない粒度で分割単位を提案する。

## アウトプット
- `docs/workflows/` 配下の構成案（例: `overview.md`, `stories.md`, `ci.md`, `local-dev.md` など）。
- README 受け入れ基準と矛盾しないことを確認したメモ。

## 構成案（初期）
- `docs/workflows/overview.md`
  - 開発者向けワークフローの全体像。
  - Codex・stories/discovery・Vibe‑kanban・CI・ローカル検証の関係を俯瞰して説明。
  - 詳細は各専用ファイル（stories/local-dev/ci/vibe-kanban）へのリンクで案内。

- `docs/workflows/stories.md`
  - 現行 `docs/REF_STORIES.md` の内容をベースに、「ストーリー/Discovery/Backlog 運用」の中心ドキュメントとする。
  - Codex との連携（stories/BACKLOG, discovery 昇格, FROM_DISCOVERY など）もこちらに集約。
  - 将来的に Vibe‑kanban との連携部分も必要に応じてここから参照。

- `docs/workflows/local-dev.md`
  - ローカル開発・検証の流れを整理（`make smoke` / `offline` / `mcp` / `check-ui` など）。
  - codex-cloud や `docs/CODEX_CLOUD.md` への詳細リンクもここから誘導。
  - 出力ディレクトリ構成（`out/bin`, `out/logs`, `out/coverage`）も必要であれば要約。

- `docs/workflows/ci.md`
  - CI 構成の概要（ジョブ列、目的、依存）を記述。
  - 実際のワークフロー定義ファイル（`.github/workflows/ci.yml`）への参照。
  - 将来 CI が増えた場合の追記先としても利用。

- `docs/workflows/vibe-kanban.md`（名称は要調整候補）
  - Vibe‑kanban を用いた日々の作業フロー（ストーリーの開始/完了、Codex との連携）を整理。
  - URL 例や、ストーリー slug をタイトルに入れる運用などを記述。
  - ストーリー運用（stories.md）からもリンクする。

## メモ（README/REF_STORIES との関係）
- `docs/WORKFLOW.md` は最終的に `docs/workflows/overview.md` 相当の内容に集約し、入口としての役割に絞る想定。
- `docs/REF_STORIES.md` は `docs/workflows/stories.md` へ実体を移し、旧ファイルは「ストーリー運用ガイドは `docs/workflows/stories.md` を参照」の stub とする方針が妥当そう。
- README からの「ワークフロー/ストーリー関連」リンクは、原則 `docs/workflows/overview.md` または `docs/workflows/stories.md` のどちらかに寄せる。
