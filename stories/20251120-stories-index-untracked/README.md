# 20251120-stories-index-untracked — Storiesインデックスの非管理化と derived ビュー運用

ステータス: [準備完了]
担当: @tkg-engineer
開始: 2025-11-20 02:20:06 +0900

## 目的・背景
- `stories/BACKLOG.md` や `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` を各作業ブランチで更新しており、ストーリーごとにコンフリクトや「一覧だけのコミット」が頻発している。
- これらの一覧ファイルは `stories/**/README.md` や Discovery 本文から自動生成できる派生ビューであり、履歴が欲しいのは一覧そのものではなくストーリー/Discovery 本体である。
- 一人運用＋複数 worktree 前提で、インデックス更新を git 運用から切り離し、「見たくなったときにコマンドで生成するだけ」のスタイルに揃えたい。

## スコープ（成果）
- Stories/Discovery の「真のソース」は `stories/**/README.md` および `stories/discovery/**/*.md` に集約し、`stories/BACKLOG.md` と各 INDEX は完全に派生物として扱う方針を整理する。
- `.gitignore` などを通じて `stories/BACKLOG.md`, `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` を git 管理対象から外し、通常の開発フローで `git status` に出てこないようにする。
- `make backlog-index` / `make discovery-index` / `make story-index` を前提とした閲覧フローを文書化し、今後の自分用の「ストーリー一覧の見方」を AGENTS / docs に残す。

## 受け入れ基準（Definition of Done）
- [ ] `.gitignore`（または同等の設定）により、`stories/BACKLOG.md`, `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` が未追跡変更として表示されない。
- [ ] `stories/BACKLOG.md`, `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` を削除した状態からでも、`make backlog-index` / `make discovery-index` / `make story-index` で問題なく再生成できることを確認している。
- [ ] `docs/workflows/stories.md` や `AGENTS.md` などに、「インデックス系ファイルは派生ビューであり git 管理外／手で編集しない」旨の一文と閲覧手順が追記されている。
- [ ] 複数の作業ブランチ／worktree でストーリー作成・更新を行っても、インデックス系ファイルによるコンフリクトや余計なコミットが発生しない運用フローが確認できている。

## 工程（サブタスク）
- [ ] 01_現状運用とインデックス生成スクリプトの棚卸し — `tasks/01_audit-index-and-scripts.md`
- [ ] 02_.gitignore 設計とインデックス非管理化の適用 — `tasks/02_apply-untracked-index-policy.md`
- [ ] 03_docs/AGENTS への方針反映と閲覧フロー整理 — `tasks/03_update-docs-and-workflow.md`

## 計画（目安）
- 見積: 1 セッション（設計〜反映まで）
- マイルストン: M1 方針整理 → M2 `.gitignore` 適用 → M3 docs/AGENTS 更新

## 進捗・決定事項（ログ）
- 2025-11-20 02:20:06 +0900: ストーリー作成
- 2025-11-20 02:20:16 +0900: README の目的・スコープ・DoD/工程を整理し、インデックス非管理化方針の叩き台ができたためステータスを[準備完了]へ更新。

## リスク・懸念
- インデックス系ファイルが git 管理外になることで、一覧の diff や履歴を直接追えなくなる（ただしストーリー/Discovery 本文の履歴で代替可能）。
- 将来複数人運用になった際に、「誰かがうっかりインデックスをコミットする」ケースが出ないよう、docs/AGENTS の記述を十分にしておく必要がある。

## 関連
- PR: #
- Issue: #
- Docs: `docs/workflows/stories.md`, `AGENTS.md`, `.gitignore`
