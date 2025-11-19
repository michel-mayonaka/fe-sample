# Vibe‑kanban を用いた作業フロー

Vibe‑kanban は、ストーリー単位の作業を管理するためのタスクボードとして利用します。

## ローカル起動（例）
- Node.js/npm が入っている環境で、リポジトリとは別の作業用ディレクトリなどから次を実行します:
  - `npx vibe-kanban`
- 起動後に表示される URL（例: `http://127.0.0.1:64909/...`）へブラウザでアクセスします。

## ストーリーの作業開始
- ローカルで立ち上げている Vibe‑kanban にてストーリーの作業を行います（URL は環境により異なります）。
  - 例: `http://127.0.0.1:64909/projects/.../tasks`
- 作業を行いたいストーリーの slug をタイトルに指定して作業を開始します。
  - 例: `20250929-input-decouple-ui-plan-c ストーリーの作業開始`

## Codex との連携
- 作業の中で新たな提案や、分割して実行するタスク/ストーリーが発生した場合は、Vibe‑kanban 内で稼働している Codex（AI エージェント）が `stories/discovery/` に Discovery を自動的に追加します。
- 完了したタスクやストーリーは、Codex によるレビュー/マージフロー（PR 作成〜マージ）と紐付けて運用します。

### Discovery 生成ルール（作業ブランチでの後続タスク）
- 後続タスクや新しいストーリー候補は、常に `stories/discovery/` 直下に `YYYYMMDD-HHMMSS-slug.md` として作成します（`ステータス: [open]` を既定とする）。
- Discovery には、提案元ストーリーを示す `関連ストーリー: 2025xxxx-some-story` を必ず付与し、どの作業から生まれたかを追跡できるようにします。
- `accepted/`, `consumed/`, `declined/` 配下へ直接ファイルを置くのは人間/スクリプトの責務とし、AI エージェントはこれらのディレクトリへ直接書き込みません。

### Discovery の状態遷移（シンプル運用）
- 「やる」と決めた Discovery は、`FROM_DISCOVERY=stories/discovery/....md make new-story SLUG=...` でストーリー化します。
  - `scripts/new_story.sh` から `scripts/consume_discovery.sh` が呼ばれ、対象 Discovery は `stories/discovery/consumed/` へ移動し、`関連ストーリー:` に新ストーリーの slug が記録されます。
- 「やらない」と決めた Discovery は、`make decline-discovery FILE=...` で `stories/discovery/declined/` へ移動します。
- 一人運用では、`accepted/` や Backlog（`stories/BACKLOG.md`）を必須のステップとはせず、必要になったときにだけ使う拡張レイヤーとします（基本は `open → consumed（＋story）` または `open → declined` の二経路）。
