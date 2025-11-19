# ストーリー運用ガイド（workflow/stories）

本プロジェクトでは、実装の“まとまり”をストーリー（Story）として管理します。タスクはストーリー配下の工程（小粒度の作業）として扱います。

## 全体運用の概要
- 課題やアイデアはまず Discovery（`stories/discovery/`）として起票し、必要に応じて Backlog（`stories/BACKLOG.md`）へ整理します。
- ストーリーは `master` ブランチ上でのみ作成し、`make new-story` によって起票します（作成経路の詳細は後述の「作成フロー」を参照）。
- 各ストーリーごとに「1ストーリー=1作業ブランチ」を切り、そのブランチ上でタスク MD を進めつつ実装・コミット・PR を行います。
- レビューが通ったら、その作業ブランチ上で README/タスクのステータスと DoD を [完了] に揃え、必要に応じて `make finish-story` で `stories/finish/` へアーカイブしてから PR をマージします。
- 索引系ファイル（`stories/BACKLOG.md`, `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` など）は必要なときに `make backlog-index` / `make discovery-index` / `make story-index` で生成する派生ビューであり、原則コミットせず、ストーリー/Discovery 本文を「真のソース」として扱います。

## 基本ルールと用語
- 置き場所: `stories/`
- ディレクトリ命名: `YYYYMMDD-slug/`（例: `20250928-ref-architecture`）
- ステータス記号: `[完了] / [進行中] / [次セッション] / [保留/N/A]`
- 参照関係: ストーリーは複数のタスク（`tasks/`）やコード変更、PR を束ねます。
  - README の見出し付近にある `[完了] / [進行中] ...` は人間向け表示用、メタキー `ステータス: [...]` はスクリプトや検証用の機械可読ステータスとして使います。

## ストーリー運用

### ストーリーの構成
- `stories/YYYYMMDD-slug/README.md`: 1ファイルで完結（テンプレート使用）
  - 目的・背景 / スコープ / 受け入れ基準
  - 工程（サブタスクとリンク）/ ステータス
  - 計画（見積/マイルストン）/ 進捗・決定 / リスク
  - 関連リンク（PR/Issue/docs）

### 作成フロー
ストーリーの作成経路は次の 2 つに限定します（いずれも `master` ブランチ上で行う前提です）。
- `master` 上で手動プランニングしつつ `make new-story SLUG=...` で新規ストーリーを起票する。
- Discovery/Backlog のエントリを昇格させる形で `make new-story`（必要に応じて `FROM_DISCOVERY=...` 併用）を実行する。
  - 作業ブランチから直接 `make new-story` してストーリーを増やさず、「新しく必要になったストーリー」は一度 Discovery として起票し、後から `master` 上で昇格させます。

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

### 作業ブランチ運用とクローズタイミング
- ストーリーに紐づく実装作業は、原則として「1ストーリー=1作業ブランチ」で行います（例: `feat/20250928-ref-architecture`）。
- ブランチ名・PR タイトル・コミットサマリには、対象ストーリーの slug（`YYYYMMDD-slug`）を含めて紐付けが分かるようにします。
- 作業ブランチでレビューが通ったタイミングで、そのブランチ上でストーリー README / タスクのステータスとチェックボックスを [完了] に揃えます。
- 原則として、このタイミングで `make finish-story` を実行して `stories/finish/` へアーカイブし、ストーリーのクローズ（[完了] + finish-story 済み）までを 1 作業ブランチ内で完結させてから PR をマージします（設計メモなどで finish を遅延したい特殊なケースは、README の進捗ログ等に方針をメモしておきます）。

### サブタスク運用
- サブタスクは一つ一つを別 MD として、該当ストーリーディレクトリ内に作成し、進捗管理を行うこと。
- サブタスク MD の作成は README 作成後、制作者レビューが終了し内容が確定してから行うこと。
- 命名は実行順が分かるよう連番を推奨（例: `01_outline.md`、`02_impl.md`、`03_review.md`）。

ステータスと進捗更新（推奨）:
- 各サブタスク MD の先頭に「ステータス」欄を設ける（例: `ステータス: [未着手]/[進行中]/[レビュー待ち]/[完了]`）。
- 作業する度に該当サブタスクのステータスと進捗ログを更新する。
- `## 進捗ログ` に日時（秒精度）＋一行で記録（例: `2025-09-29 14:23:05 +0900: 入力列挙を定義しテスト追加`）。
- ストーリー README の「工程（サブタスク）」チェックボックスと整合させる（完了時は [x] に変更）。

## Discovery 運用
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

## Backlog 運用
- Backlog（将来のストーリー候補）は Discovery から昇格したストーリー候補の一覧です。
- 置き場所: `stories/BACKLOG.md`（自動生成）。
- 生成: `make backlog-index`（ソースは `stories/discovery/accepted/*.md`）。
- 記法: 1エントリ=1セクションで「目的/背景/DoD/関連」を簡潔に（テンプレ自動出力）。
- 昇格フロー: 採択が決まったら `make promote-discovery FILE=<path>` で Backlog に反映。詳細設計はストーリー新規作成後に移行。
- 見直し: セッション/週次の終わりに discovery をトリアージ（open/promoted/declined）。

## メタキー（共通）
- 先頭に以下のキーを置く（Story/Discovery 共通）:
  - `ステータス: [proposed|open|in_progress|promoted|declined|done]`
  - `担当: @handle`
  - `開始: YYYY-MM-DD HH:MM:SS +0900`
  - `優先度: P0|P1|P2|P3`（Discovery で主に使用）
  - 必要に応じて `関連ストーリー:` を追記。
