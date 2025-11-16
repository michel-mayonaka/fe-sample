# 代表 specs へのメタデータ適用

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-17 00:35:10 +0900

## 目的
- 設計したメタキーと状態表現を、実際の spec ファイルに適用し、運用イメージと現実的な更新コストを確認する。
- 今後の Story で参照しやすい「見本 spec」を用意する。

## 完了条件（DoD）
- [x] 少なくとも 1 つの system spec（例: `docs/specs/system/battle_basic.md`）にメタ情報が追加されている。
- [x] 少なくとも 1 つの ui spec（例: `docs/specs/ui/status_screen.md`）にメタ情報が追加されている。
- [x] 付与したメタ情報が、設計したメタキー・状態のルールに沿っている。

## 作業手順（概略）
- タスク01で設計したメタキーと状態の一覧を確認する。
- 対象となる spec ファイルを選定し、ファイル先頭にメタ情報ブロック（例: 状態・主な実装・最新ストーリー）を追加する。
- 追加後の見た目と編集しやすさを確認し、必要があればメタキー案を微修正する。

## 進捗ログ
- 2025-11-17 00:35:10 +0900: タスク作成。
- 2025-11-17 00:38:00 +0900: `docs/specs/ui/status_screen.md` と `docs/specs/system/battle_basic.md` の先頭に、状態/主な実装/最新ストーリーのメタ情報ブロックを追加。

## 依存／ブロッカー
- `stories/20251117-spec-status-metadata/tasks/01_spec_metadata_design.md` の内容（メタキー設計）。

## 成果物リンク
- `docs/specs/system/battle_basic.md`
- `docs/specs/ui/status_screen.md`
