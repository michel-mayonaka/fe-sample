# 01 Port API 再設計

ステータス: [完了]
担当: @codex

## 目的
- `internal/game/data.TableProvider` を UI 非依存の純粋な参照 Port として再定義する。

## ToDo
- 現行 Provider 実装と利用箇所を棚卸しする。
- UI 型へ依存しているエクスポート API を洗い出す。
- 代替となる内部モデル/DTO の定義方針を固める。
- インターフェース変更案と影響範囲を README/設計メモにまとめる。

## 進捗ログ
- 2025-09-30: タスクを起票。
- 2025-09-30: `internal/game/data.TableProvider` から `UserUnitByID` を削除し、UI 依存を解消。
