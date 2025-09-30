# 01_backlog_index — accepted から BACKLOG 自動生成

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-09-30 18:42:30 +0900

## 目的
- `stories/BACKLOG.md` を人手編集から脱却し、`stories/discovery/accepted/*.md` を唯一のソースにする。

## 完了条件（DoD）
- [ ] `scripts/gen_backlog.sh` を追加し、accepted/*.md から Backlog セクションを生成。
- [ ] 優先度/日付/タイトル/関連（Story/ファイル）を出力。
- [ ] `make backlog-index` ターゲット追加。
- [ ] 既存 BACKLOG の手編集運用を README/Docs で非推奨化。

## 作業手順（概略）
- メタ抽出（rg/sed/awk）→ 生成テンプレへ整形。
- 既存 BACKLOG のヘッダ/雛形を保持し、エントリ部のみ上書き。

## 進捗ログ
- 2025-09-30 18:42:30 +0900: タスク起票

- 2025-09-30 18:53:35 +0900: gen_backlog.sh と Make 追加、出力を確認
