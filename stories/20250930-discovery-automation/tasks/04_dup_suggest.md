# 04_dup_suggest — 重複サジェスト（new_discovery）

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-09-30 18:42:45 +0900

## 目的
- 類似タイトル/キーワードの既存 Discovery/Backlog を表示し、重複起票を抑止する。

## 完了条件（DoD）
- [ ] `new_discovery.sh` 実行時にタイトル/slug で `rg -i` を行い、上位ヒットを提示。
- [ ] `--no-suggest` オプションで抑制可能。

## 進捗ログ
- 2025-09-30 18:42:45 +0900: タスク起票

- 2025-09-30 18:53:35 +0900: new_discovery.sh に重複サジェストを追加（NO_SUGGEST=1で抑制可）
