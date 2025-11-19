# .gitignore 設計とインデックス非管理化の適用

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-20 02:20:06 +0900

## 目的
- `stories/BACKLOG.md`, `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` を git 管理から外し、通常の開発フローでインデックス系ファイルを意識しなくてよい状態にする。
- インデックス系ファイルが存在しない/改変された状態でも、`make backlog-index` / `make discovery-index` / `make story-index` によって即座に再生成できることを確認する。

## 完了条件（DoD）
- [ ] `.gitignore`（または同等の仕組み）にインデックス系ファイルを除外するエントリが追加されている。
- [ ] インデックス生成後の状態で `git status` を実行しても、`stories/BACKLOG.md`, `stories/discovery/INDEX.md`, `stories/finish/INDEX.md` が変更として表示されない。
- [ ] インデックス系ファイルを削除したあとに `make backlog-index` / `make discovery-index` / `make story-index` を実行し、エラーなく再生成されることを確認している。

## 作業手順（概略）
- `.gitignore` を編集し、対象ファイル（Backlog/Discovery Index/Finish Index）を除外するパターンを追加する。
- 既存のインデックスを削除または退避し、各 Make ターゲットで再生成できることを確認する。
- 生成後の状態で `git status` を確認し、インデックス系ファイルが開発フローのノイズになっていないことを確認する。

## 進捗ログ
- 2025-11-20 02:20:16 +0900: タスク作成。

## 依存／ブロッカー
- `01_audit-index-and-scripts.md` で前提が整理されているとスムーズ。

## 成果物リンク
- `.gitignore`: `.gitignore`

