# docs/AGENTS への方針反映と閲覧フロー整理

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-20 02:20:06 +0900

## 目的
- インデックス系ファイルを git 管理から外す方針を AGENTS / docs に明記し、将来の自分や他の開発者が迷わないようにする。
- Backlog/Discovery/完了ストーリー一覧を見たいときのコマンド (`make backlog-index`, `make discovery-index`, `make story-index`) と運用パターンをドキュメント化する。

## 完了条件（DoD）
- [ ] `AGENTS.md` または `docs/workflows/stories.md` 等に、「インデックス系ファイルは派生ビューであり git 管理外・直接編集しない」ことが記載されている。
- [ ] Backlog/Discovery/完了ストーリーの一覧を確認する手順（使う Make ターゲットと生成先ファイル）が、1〜2 箇所のドキュメントから辿れる。
- [ ] 本ストーリー `20251120-stories-index-untracked` への簡単な言及（またはリンク）がどこかに残っており、経緯が追える。

## 作業手順（概略）
- `docs/workflows/stories.md` と `AGENTS.md` を読み、インデックス運用に触れるべきセクションを確認する。
- インデックス非管理化の方針と閲覧コマンドを追記し、既存の説明と矛盾しないように調整する。
- 必要に応じて、このストーリーへのリンクや補足メモを追加する。

## 進捗ログ
- 2025-11-20 02:20:16 +0900: タスク作成。

## 依存／ブロッカー
- `.gitignore` 適用後の最終状態を前提とするため、`02_apply-untracked-index-policy.md` の内容を踏まえて更新する。

## 成果物リンク
- `AGENTS.md`: `AGENTS.md`
- ストーリー運用ガイド: `docs/workflows/stories.md`

