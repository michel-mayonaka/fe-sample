# 開発者向けワークフロー（最小ガイド）

本プロジェクトではAI駆動開発を採用しており、主に下記サービスを使用して進行します。

## Codex（AIコーディングエージェント）
### ストーリーの作成
- ストーリーの作成は主にstories/BACKLOG.mdにあるものから作成するか、突発ストーリーの作成をお願いするかの2択に
なります。
- 作成の際にはmasterブランチにて下記のようなプロンプトを実行して作成します。
```
Codex へ:
開発者向けワークフローガイド策定 のストーリーを作成したい
```
- 具体的なストーリー運用に関してはリンクをチェックしてください: `docs/REF_STORIES.md`（Discovery/Backlog/Story の扱い）
- Codexではストーリーの作成・レビュー・承認・コミットまでを行い実際の作業はVibe‑kanbanを使用して行います。
 - コード変更を伴う作業は `feat/<slug>` ブランチで実施（ドキュメント/ストーリーのみは master 直コミット可）。

### stories/discovery の昇格について
- Vibe‑kanbanからの課題提案は `stories/discovery/`（open）へ起票されます。
  - 未整理の課題になるのでstories/BACKLOG.mdに載せるかの判断を定期的にCodexに伝えるようお願いします。
    - 見送った場合は `stories/discovery/declined/`、採択は `stories/discovery/accepted/` に格納されます（Backlog は accepted から自動生成）。
- 昇格してストーリー化したら FROM_DISCOVERY を使って `stories/discovery/consumed/` へ退避されます（Backlog からも自動で消えます）。
- ※Codex に直接依頼して生成したストーリーはdiscoveryが存在しない形になります。

## Vibe‑kanban
### ストーリーの作業開始（Vibe‑kanban）
- local で立ち上げている Vibe‑kanban にてストーリーの作業を行います（URL は環境により異なる例）。
  - 例: http://127.0.0.1:64909/projects/a6dcabe1-fa65-4110-8fab-810427163ed8/tasks
- 作業を行いたいストーリーの slug をタイトルに指定して作業を開始します。
例
```
20250929-input-decouple-ui-plan-c ストーリーの作業開始
```
- 作業の中で新たな提案や分割して実行するタスクやストーリーが発生した場合はVibe‑kanban内で稼働しているCodexが `stories/discovery/` に自動的に内容を追加してくれます。
- 作業が完了したらレビューを行い問題なければマージを行なってください。
