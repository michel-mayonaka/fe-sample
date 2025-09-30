# ストーリー運用ガイド（ref_stories）

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
1) スラグを決める（例: `ref-architecture`）。
2) 生成コマンド（推奨）を使う: `make new-story SLUG=ref-architecture`
   - 生成物: `stories/$(date +%Y%m%d)-$(SLUG)/README.md`（テンプレート展開）
3) README を編集（目的/スコープ/DoD を確定）。
4) 制作者レビューを実施して README の内容を確定。
5) サブタスクMDを作成（必須）: ストーリーディレクトリ直下に 1 サブタスク = 1 MD で作成し、進捗管理を行う。
   - 置き場所例: `stories/YYYYMMDD-slug/tasks/01_outline.md`、`.../tasks/02_impl.md`
   - README の「工程」欄から各サブタスクMDへリンクする。
   - 横断的な改善メモは `stories/BACKLOG.md` に記載し、採択時に当該ストーリー配下へ移す（ルート直下 `tasks/` は廃止）。
6) コミット: Conventional Commits で `[story]` を含むサマリ推奨。
7) 進捗更新: セッション終了時にステータスと日付入りログを追記。

## 完了ストーリーのアーカイブ運用
- 完了したストーリーは `stories/finish/` 配下へ移動します。
- 推奨コマンド: `make finish-story SLUG=ref-architecture` または `make finish-story PATH=stories/20250928-ref-architecture`
- 移動時に `README.md` の先頭ステータスを `[完了]` に自動更新します（スクリプトが置換）。

## Backlog（将来のストーリー候補）運用
- 置き場所: `stories/BACKLOG.md`（自動生成）。
- 生成: `make backlog-index`（ソースは `stories/discovery/accepted/*.md`）。
- 記法: 1エントリ=1セクションで「目的/背景/DoD/関連」を簡潔に（テンプレ自動出力）。
- 昇格フロー: 採択が決まったら `make promote-discovery FILE=<path>` で Backlog に反映。詳細設計はストーリー新規作成後に移行。
- 見直し: セッション/週次の終わりに discovery をトリアージ（open/promoted/declined）。

### Discovery（並行作業向けの衝突回避）
- 目的: 複数ストーリーの並行作業で `stories/BACKLOG.md` への同時追記によるコンフリクトを避ける。
- 置き場所: `stories/discovery/`（1課題=1 MD）。
- 新規作成: `make new-discovery SLUG=<slug> [TITLE=..] [PRIORITY=P2] [STORY=YYYYMMDD-slug]`
  - 生成物: `stories/discovery/YYYYMMDD-HHMMSS-<slug>.md`（秒精度の作成時刻を含む）
- 採択（Backlogへ昇格）: `make promote-discovery FILE=<path> [PRIORITY=Px]`
  - 挙動: BACKLOG にエントリを追加、`状態: [promoted]` へ更新し `stories/discovery/accepted/` へ移動。
- 見送り（やらない判断）: `make decline-discovery FILE=<path> [REASON=..]`
  - 挙動: `状態: [declined]` へ更新し、`stories/discovery/declined/` へ移動。理由をログに追記。
 - 索引: `make discovery-index` で一覧を生成（`stories/discovery/INDEX.md`）。
※ Backlog 追記は「採択時のみ」とし、発見段階の共有は discovery ファイルで行う運用を推奨。

## テンプレート
- 雛形は `stories/_TEMPLATE/README.md` に配置。
- 直接コピーする場合: `cp -R stories/_TEMPLATE stories/20250928-your-story && vi stories/20250928-your-story/README.md`

## 運用ヒント
- 1 セッション = 1〜数工程。差分は小さく、ロールバック容易に。
- ストーリーは“目的と受け入れ基準”を先に固定。途中の仕様変更は「決定事項」に追記。
- 実装前後で `make mcp`、必要に応じて `go test -race -cover` を実施。

### 事例: UI 責務分割パターン（layout/draw/view/adapter）
- 動機: `helper` 的な曖昧名を排し、探索性と再利用性を高める。
- 構成:
  - `ui/layout`: 副作用のない矩形計算（座標/サイズ）。
  - `ui/draw`: 見た目の描画（テキスト/パネル/ボタン等）。
  - `ui/view`: 表示用データ（行モデルなど）。
  - `ui/adapter`: テーブル/ユーザ→view の変換（`PortraitLoader` で画像読込を抽象化）。
- 適用例: `scenes/list_layout` → `ui/layout/list.go`、`scenes/view_status_draw` → `ui/draw/status.go`、
  在庫行の描画/生成 → `ui/draw/inventory_list.go` / `ui/adapter/inventory_*.go`。

### サブタスク運用（必須）
- サブタスクは一つ一つを別 MD として、該当ストーリーディレクトリ内に作成し、進捗管理を行うこと。
- サブタスクMDの作成は README 作成後、制作者レビューが終了し内容が確定してから行うこと。
- 命名は実行順が分かるよう連番を推奨（例: `01_outline.md`、`02_impl.md`、`03_review.md`）。

#### ステータスと進捗更新（追記ルール）
- 各サブタスクMDの先頭に「ステータス」欄を設けること。
  - 例: `ステータス: [未着手]/[進行中]/[レビュー待ち]/[完了]`
- 作業する度に該当サブタスクのステータスと進捗ログを更新すること。
- 推奨: `## 進捗ログ` に日時（秒精度）＋一行で記録（例: `2025-09-29 14:23:05 +0900: 入力列挙を定義しテスト追加`）。
- ストーリーREADMEの「工程（サブタスク）」チェックボックスと整合させる（完了時は [x] に変更）。
### メタキー（共通）
- 先頭に以下のキーを置く（Story/Discovery 共通）:
  - `ステータス: [proposed|open|in_progress|promoted|declined|done]`
  - `担当: @handle`
  - `開始: YYYY-MM-DD HH:MM:SS +0900`
  - `優先度: P0|P1|P2|P3`（Discovery で主に使用）
  - 必要に応じて `関連ストーリー:` を追記。
