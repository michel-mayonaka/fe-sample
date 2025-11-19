# インデックス系ファイルと生成スクリプトの棚卸し

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-20 02:20:06 +0900

## 目的
- `stories/BACKLOG.md`, `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` の役割と、どのスクリプト/Make ターゲットから生成されているかを整理する。
- インデックス系ファイルを git 管理から外しても問題にならない前提（参照元や期待される運用）を確認し、ストーリー本体/Discovery を「真のソース」とみなせることを確かめる。

## 完了条件（DoD）
- [ ] `scripts/gen_backlog.sh`, `scripts/gen_discovery_index.sh`, `scripts/gen_story_index.sh` の挙動と生成先ファイルが簡単なメモとして残っている。
- [ ] インデックス系ファイルが存在しない状態からでも、各スクリプトがエラーなくファイルを再生成できることを確認している。
- [ ] コードや docs からインデックス系ファイルを「永続的なソース」と見なしている箇所がない（またはあっても本ストーリーで一緒に是正する）ことを確認している。

## 作業手順（概略）
- `scripts/gen_backlog.sh`, `scripts/gen_discovery_index.sh`, `scripts/gen_story_index.sh` を読み、生成ロジックと前提を把握する。
- `rg 'BACKLOG.md|discovery/INDEX.md|finish/INDEX.md'` などでリポジトリ内の参照箇所を洗い出す。
- 必要であれば簡単なノートを `stories/20251120-stories-index-untracked/` 配下に残し、後続タスクで参照できるようにする。

## 進捗ログ
- 2025-11-20 02:20:16 +0900: タスク作成。

## 依存／ブロッカー
- 特になし。

## 成果物リンク
- メモ（任意）: `stories/20251120-stories-index-untracked/notes_index-audit.md`

