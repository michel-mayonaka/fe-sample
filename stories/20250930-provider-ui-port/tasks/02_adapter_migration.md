# 02 UI Adapter 変換移行

ステータス: [完了]
担当: @codex

## 目的
- ユーザデータ→`uicore.Unit` の変換責務を `service/ui/adapter` に集約し、Provider から排除する。

## ToDo
- 既存の `UserUnitByID` 呼び出し箇所で必要な属性を整理する。
- Adapter 層に新しい変換関数/構造体を設計・実装する。
- 変換処理のテスト（既存/追加）を整備して回帰を防ぐ。
- Provider から移行したコードのクリーニングとコメント整備を行う。

## 進捗ログ
- 2025-09-30: タスクを起票。
- 2025-09-30: 変換は `uicore.UnitFromUser` を採用済み。`inventory.refreshUnitByID` を新経路へ移行。
