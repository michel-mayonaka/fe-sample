# 03_policy — 分割ポリシーとルール定義

## 目的
- WORKFLOW 系ドキュメントを今後も継続的に保守しやすくするため、分割ポリシーと更新ルールを明文化する。

## 作業項目
- どの種類の情報をどのファイルに置くかのルールを言語化する。
- 他のドキュメント（`docs/ARCHITECTURE.md`, `docs/REF_STORIES.md` など）との役割分担を整理する。
- 今後 WORKFLOW 関連の情報を追加・変更する際のガイドライン（どのファイルを編集するか）を決める。

## アウトプット
- 分割ポリシーと運用ルール（短いセクション単位で可）。
- 必要であれば `docs/WORKFLOW.md` または新ファイルへの記載案。

## 分割ポリシー（案）
- 役割ベースで配置する:
  - ワークフロー全体像・入口 → `docs/workflow/overview.md`
  - ストーリー/Discovery/Backlog 運用 → `docs/workflow/stories.md`
  - ローカル開発・検証フロー → `docs/workflow/local-dev.md`
  - CI 構成・ジョブ概要 → `docs/workflow/ci.md`
  - Vibe‑kanban 等のツール連携 → `docs/workflow/vibe-kanban.md`
- 層の考え方:
  - `README.md`: 「どう始めるか」の超入口。詳細は必ず `docs/workflow/*.md` へリンク。
  - `docs/WORKFLOW.md`: 将来的には `overview.md` へのブリッジ（入口 stub）に寄せ、詳細は `docs/workflow/` に集約。
  - その他のドキュメント（`ARCHITECTURE.md`, `CODEX_CLOUD.md` 等）は、それぞれの詳細を保持しつつ、必要に応じて workflow 側から参照する。

## 他ドキュメントとの役割分担
- `docs/ARCHITECTURE.md`
  - システム構成・レイヤリング・依存関係など「静的な構造」の説明を担う。
  - ワークフロー（どの順番で何を実行するか）は極力書かず、必要に応じて `docs/workflow/overview.md` へのリンクを貼る程度に留める。
- `docs/REF_STORIES.md`
  - 実体としては `docs/workflow/stories.md` に統合する。
  - 旧ファイルは「ストーリー運用ガイドは `docs/workflow/stories.md` を参照」の stub とし、内容は残さない方針。
- `docs/CODEX_CLOUD.md`
  - codex-cloud や `make mcp` 周辺の詳細な説明を担当。
  - `docs/workflow/local-dev.md` は「日々どう使うか」のガイドとし、詳細は CODEX_CLOUD へリンクする。

## 更新ガイドライン
- ワークフロー関連の新情報を追加する際は、まず「どの観点か」を決める:
  - ストーリー運用の手順/ルール → `docs/workflow/stories.md`
  - ローカルでの検証/ビルドの流れ → `docs/workflow/local-dev.md`
  - CI のジョブや検証内容 → `docs/workflow/ci.md`
  - Vibe‑kanban やツール連携 → `docs/workflow/vibe-kanban.md`
  - それらを横断して俯瞰した説明 → `docs/workflow/overview.md`
- `README.md` にワークフロー情報を足したくなった場合:
  - 原則として `docs/workflow/*.md` へのリンクを追加し、詳細説明はそちらに追記する。
- 既存ドキュメントとの重複を見つけた場合:
  - どちらを「正」とするかを決め、もう片方は stub 化またはリンク案内のみにする。
  - 今回はストーリー運用に関して `docs/workflow/stories.md` を「正」とし、`docs/REF_STORIES.md` は stub にする。
