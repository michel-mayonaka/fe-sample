# docs 現状棚卸し

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-16 20:26:57 +0900

## 目的
- `docs/` 配下のファイル構成と内容を棚卸しし、重複・陳腐化・不足箇所を洗い出す。
- 特にシステム仕様・画面仕様など、将来的に `docs/specs/` に集約したい情報を把握する。

## 完了条件（DoD）
- [ ] `docs/` 配下のファイル一覧と簡易説明が整理されている。
- [x] 仕様相当ドキュメントの現在地（どのファイルに何が書かれているか）が把握できるメモがある。
- [x] 問題点（重複・陳腐化・不足）のメモが Story README か本タスクに記録されている。
- [x] 後続タスクで検討すべき論点が箇条書きでまとまっている。

## 作業手順（概略）
- `docs/` 以下のファイル/ディレクトリを一覧化する。
- 各ファイルの役割・最終更新日・参照想定をざっくり把握する。
- システム仕様・画面仕様など仕様寄りの内容がどこに散らばっているかを確認する。
- 重複・陳腐化・不足が疑われる箇所をメモに起こす。

## 進捗ログ
- 2025-11-16: docs/ 配下のファイル一覧と役割メモを本タスクに追記。
- 2025-11-16 20:26:57 +0900: タスク作成

## 依存／ブロッカー
- なし（単独で実施可能）。

## 成果物リンク
- Docs: `docs/`
- Story: `stories/20251116-docs-cleanup/README.md`

## docs/ 現状メモ（2025-11-16 時点）

### ルート直下
- `docs/AI_OPERATIONS.md`: AI コーディングエージェント/自動化まわりの内部メモ。仕様というより運用/参考資料。
- `docs/API.md`: 外部 API 視点の仕様置き場。現時点では簡易的で、今後拡張余地あり。
- `docs/ARCHITECTURE.md`: 現行の理想アーキテクチャ説明。設計思想/レイヤ構造が中心で、仕様の一部を内包。
- `docs/CODEX_CLOUD.md`: codex-cloud 環境での実行方法。開発環境/運用向け。
- `docs/COMMENT_STYLE.md`: コメント/GoDoc 記法のルール。仕様ではなくスタイルガイド。
- `docs/DB_NOTES.md`: DB/永続化に関するメモと将来の移行方針。
- `docs/DOCS_STRUCTURE.md`: docs 配下全体の構成ポリシー。今回のストーリーで追加したガイド。
- `docs/MIGRATION_20251013_CODEX_CLOUD.md`: codex-cloud 関連の移行メモ。時系列の履歴的な位置づけ。
- `docs/NAMING.md`: 命名規約。コード/データ/アセット名などの準拠先。
- `docs/OFFLINE.md`: オフライン開発/実行まわりの手順メモ。

### specs/
- `docs/specs/README.md`: 仕様ハブの説明。system/ui 等カテゴリと役割を定義。
- `docs/specs/AGENTS.md`: エージェント向けの仕様読み方ガイドと参照優先順位。
- `docs/specs/system/_TEMPLATE.md`: システム仕様のテンプレート。まだ具体仕様は未作成。
- `docs/specs/ui/_TEMPLATE.md`: 画面仕様のテンプレート。まだ具体仕様は未作成。

### workflow/
- `docs/workflow/overview.md`: 開発ワークフロー全体の概要。ストーリー/CI/ローカル開発との関係。
- `docs/workflow/stories.md`: ストーリー/タスク運用の詳細ルール。
- `docs/workflow/local-dev.md`: ローカル開発環境構築と日常的なコマンド運用。
- `docs/workflow/ci.md`: CI ジョブ構成と方針のメモ。
- `docs/workflow/vibe-kanban.md`: Vibe カンバン運用のメモ。開発プロセス寄り。

## 問題点・不足の仮メモ
- system/ui 向けの実際の仕様（例: ステータス画面の仕様、戦闘ロジックの仕様）がまだ `docs/specs/` に書かれていない。
- 一部の仕様的な説明（例: 戦闘ルール、画面の振る舞いなど）が `docs/ARCHITECTURE.md` や README に散在しており、将来的に specs へ切り出した方が良さそう。
- API.md は現状ラフで、system spec や domain spec が増えた際に役割の整理が必要になりそう。

## 後続タスクへの論点メモ
- どの単位で system spec / ui spec を書き始めるか（ユースケース単位かドメイン単位か画面単位か）。
- 既存ドキュメント（README/ARCHITECTURE 等）から、どこまで specs へ移すか vs リンクで参照に留めるかの方針。
- `docs/specs/system/` / `docs/specs/ui/` の細分化（サブディレクトリ構成）は、別バックログとして切り出して段階的に進める（既に Backlog に追加済み）。
- 2025-11-16: docs/ 配下のファイル一覧と役割メモを本タスクに追記。
