# ストーリー運用ガイド（workflow/stories）

本プロジェクトでは、実装の“まとまり”をストーリー（Story）として管理します。タスクはストーリー配下の工程（小粒度の作業）として扱います。

- 置き場所: `stories/`
- ディレクトリ命名: `YYYYMMDD-slug/`（例: `20250928-ref-architecture`）
- ステータス記号: `[完了] / [進行中] / [次セッション] / [保留/N/A]`
- 参照関係: ストーリーは複数のタスク（`tasks/`）やコード変更、PR を束ねます。

## ストーリーの構成
- `stories/YYYYMMDD-slug/README.md`: 1ファイルで完結（テンプレート使用）
  - 目的・背景 / スコープ / 受け入れ基準
  - 工程（サブタスクとリンク）/ ステータス
  - 計画（見積/マイルストン）/ 進捗・決定 / リスク
  - 関連リンク（PR/Issue/docs）

## 作成フロー
1. スラグを決める（例: `ref-architecture`）。
2. 生成コマンド（推奨）を使う: `make new-story SLUG=ref-architecture`
   - 生成物: `stories/$(date +%Y%m%d)-$(SLUG)/README.md`（テンプレート展開）
3. README を編集（目的/スコープ/DoD を確定）。
4. 制作者レビューを実施して README の内容を確定。
5. サブタスク MD を作成（必須）: ストーリーディレクトリ直下に 1 サブタスク = 1 MD で作成し、進捗管理を行う。
   - 置き場所例: `stories/YYYYMMDD-slug/tasks/01_outline.md`、`.../tasks/02_impl.md`
   - README の「工程」欄から各サブタスク MD へリンクする。
   - 横断的な改善メモは `stories/BACKLOG.md` に記載し、採択時に当該ストーリー配下へ移す（ルート直下 `tasks/` は廃止）。
6. コミット: Conventional Commits で `[story]` を含むサマリ推奨。
7. 進捗更新: セッション終了時にステータスと日付入りログを追記。
   - Story 完了時は、README の「受け入れ基準」と各サブタスク MD の「完了条件（DoD）」のチェックボックスを、実際の作業結果と必ず揃えてからステータスを [完了] にする。

## Discovery / Backlog 運用
- Discovery の目的: 複数ストーリーの並行作業で `stories/BACKLOG.md` への同時追記によるコンフリクトを避ける。
- 置き場所: `stories/discovery/`（1課題=1 MD）。
- 新規作成: `make new-discovery SLUG=<slug> [TITLE=..] [PRIORITY=P2] [STORY=YYYYMMDD-slug]`
  - 生成物: `stories/discovery/YYYYMMDD-HHMMSS-<slug>.md`（秒精度の作成時刻を含む）。
- 採択（Backlogへ昇格）: `make promote-discovery FILE=<path> [PRIORITY=Px]`
  - 挙動: BACKLOG にエントリを追加、`状態: [promoted]` へ更新し `stories/discovery/accepted/` へ移動。
- 見送り（やらない判断）: `make decline-discovery FILE=<path> [REASON=..]`
  - 挙動: `状態: [declined]` へ更新し、`stories/discovery/declined/` へ移動。理由をログに追記。
- ストーリー化済みの消化: `make consume-discovery FILE=<accepted.md> STORY_DIR=stories/YYYYMMDD-slug`
  - 挙動: `状態: [consumed]` へ更新し `stories/discovery/consumed/` へ移動、関連ストーリーを追記（次回 `make backlog-index` 時に BACKLOG から自動で消える）。
- 索引: `make discovery-index` で一覧を生成（`stories/discovery/INDEX.md`）。

Backlog（将来のストーリー候補）:
- 置き場所: `stories/BACKLOG.md`（自動生成）。
- 生成: `make backlog-index`（ソースは `stories/discovery/accepted/*.md`）。
- 記法: 1エントリ=1セクションで「目的/背景/DoD/関連」を簡潔に（テンプレ自動出力）。
- 昇格フロー: 採択が決まったら `make promote-discovery FILE=<path>` で Backlog に反映。詳細設計はストーリー新規作成後に移行。
- 見直し: セッション/週次の終わりに discovery をトリアージ（open/promoted/declined）。

## サブタスク運用
- サブタスクは一つ一つを別 MD として、該当ストーリーディレクトリ内に作成し、進捗管理を行うこと。
- サブタスク MD の作成は README 作成後、制作者レビューが終了し内容が確定してから行うこと。
- 命名は実行順が分かるよう連番を推奨（例: `01_outline.md`、`02_impl.md`、`03_review.md`）。

ステータスと進捗更新（推奨）:
- 各サブタスク MD の先頭に「ステータス」欄を設ける（例: `ステータス: [未着手]/[進行中]/[レビュー待ち]/[完了]`）。
- 作業する度に該当サブタスクのステータスと進捗ログを更新する。
- `## 進捗ログ` に日時（秒精度）＋一行で記録（例: `2025-09-29 14:23:05 +0900: 入力列挙を定義しテスト追加`）。
- ストーリー README の「工程（サブタスク）」チェックボックスと整合させる（完了時は [x] に変更）。

## メタキー（共通）
- 先頭に以下のキーを置く（Story/Discovery 共通）:
  - `ステータス: [proposed|open|in_progress|promoted|declined|done]`
  - `担当: @handle`
  - `開始: YYYY-MM-DD HH:MM:SS +0900`
  - `優先度: P0|P1|P2|P3`（Discovery で主に使用）
  - 必要に応じて `関連ストーリー:` を追記。
