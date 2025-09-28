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
3) README を編集し、関連タスク（`tasks/...md`）やドキュメントをリンク。
4) コミット: Conventional Commits で `[story]` を含むサマリ推奨。
5) 進捗更新: セッション終了時にステータスと日付入りログを追記。

## 完了ストーリーのアーカイブ運用
- 完了したストーリーは `stories/finish/` 配下へ移動します。
- 推奨コマンド: `make finish-story SLUG=ref-architecture` または `make finish-story PATH=stories/20250928-ref-architecture`
- 移動時に `README.md` の先頭ステータスを `[完了]` に自動更新します（スクリプトが置換）。

## テンプレート
- 雛形は `stories/_TEMPLATE/README.md` に配置。
- 直接コピーする場合: `cp -R stories/_TEMPLATE stories/20250928-your-story && vi stories/20250928-your-story/README.md`

## 運用ヒント
- 1 セッション = 1〜数工程。差分は小さく、ロールバック容易に。
- ストーリーは“目的と受け入れ基準”を先に固定。途中の仕様変更は「決定事項」に追記。
- 実装前後で `make mcp`、必要に応じて `go test -race -cover` を実施。
