# 20250929-docs-naming — 命名規約ドキュメント作成

ステータス: [完了]
担当: @tkg-engineer

## 目的・背景
- リポジトリ横断の命名規約を専用ドキュメント（`docs/KNOWLEDGE/engineering/naming.md`）に集約し、レビュー/実装の判断基準を明確化する。
- 既存の断片（`README.md` のスタイル方針、`docs/architecture/README.md` の「命名・スタイル」補足、コミット規約）を統合し矛盾を解消する。
- 将来の SQLite 移行や UI リファクタ時に命名の一貫性を保ち、差分を最小化する。

## スコープ（成果）
- `docs/KNOWLEDGE/engineering/naming.md` を新規追加（Go/DB/JSON/アセット/コミット規約の観点を網羅）。
- 対象範囲:
  - Go: パッケージ名、ファイル名、エクスポート/非エクスポート識別子、初期ism（`ID/HTTP` など）、テスト命名（`*_test.go`）。
  - DB/データ: テーブル接頭辞（`mst_`/`usr_`）、カラム/JSONキー（`snake_case`）、ID命名、移行時の互換方針。
  - アセット: ディレクトリ/ファイル命名（用途先頭 + `snake_case`、例: `assets/ui/status_panel.png`）。
  - ドキュメント/PR: リンク追加（`README.md`/`ARCHITECTURE.md`）、Conventional Commits との整合例。
  - 網羅対象（チェックリスト）: ファイル/ディレクトリ/パッケージ/型（構造体・インタフェース）/クラス相当/メソッド/関数/フィールド/変数/定数/レシーバ/エラー名/テスト名/DBテーブル/カラム/JSONキー/アセット/スクリプト/Makeターゲット/コミット・PR タイトル。

## 受け入れ基準（Definition of Done）
- [x] `docs/KNOWLEDGE/engineering/naming.md` が存在し、上記範囲の章立てとサンプル（短文）を含む。
- [x] `README.md` と `docs/architecture/README.md` から `docs/KNOWLEDGE/engineering/naming.md` への参照リンクを追加。
- [x] Lint（`golangci-lint`）と `make check-all` が通る（ドキュメント追加のみ想定）。
- [x] 代表箇所の命名を 1〜2 点だけ試験的に是正し、ガイドの運用性を確認（任意）。
 - [x] 「網羅対象チェックリスト」を `docs/KNOWLEDGE/engineering/naming.md` 冒頭に掲載し、未定事項は `TBD` と明記。

## 工程（サブタスク）
- [x] 章立て案の作成（Go/DB/データ/アセット/ドキュメント）。
- [x] `docs/KNOWLEDGE/engineering/naming.md` 下書き（最小サンプルとアンチパターン例）。
- [x] 既存文書の整合チェック（`README.md`/`docs/architecture/README.md`/`docs/KNOWLEDGE/engineering/comment-style.md`）。
- [x] 参照リンクの追記（トップドキュメントから誘導）。
- [x] 軽微な是正 PR（今回は不要のため N/A として扱う）。
- [x] 最終レビューと採択（`[完了]` へ更新）。

## 計画（目安）
- 見積: 2〜3 時間 / 1 セッション
- マイルストン: M1 章立て → M2 下書き → M3 リンク/微修正

## 進捗・決定事項（ログ）
- 2025-09-29: ストーリー起票。ドキュメント化前にストーリーで目的/DoD を固定。
- 2025-09-29: `docs/KNOWLEDGE/engineering/naming.md` に詳細ルールと具体例（Go/DB/アセット/コミット）を追記。
- 2025-09-29: 解像度サフィックスは `@2x/@3x` に決定。列挙定数は型名プレフィックス + UpperCamel + `Unknown` 起点で確定。
- 2025-09-29: 命名是正（試験）: `internal/game/util` パッケージを `internal/game/rng` に改名、`*util.Rand` → `*rng.Rand`。テスト補助関数 `wtForTest` → `weaponsTableForTest` に改名。ビルド/テストOK。
- 2025-09-29: Lint（優先度高）対応: `errcheck`（テストの `Close` を明示処理）、公開識別子に GoDoc 追記（inventory/sim/status/helpers/service）。
- 2025-09-29: NAMING 検証（リネーム）: `internal/game/scenes/helpers.go` → `rect_helpers.go`、`internal/game/ctx.go` → `frame_context.go`。`make mcp` でグリーン。
 - 2025-09-29: NAMING に「汎用名の避け方（helpers/util の代替）」節を追記。`util/helpers` の代替語と判断フローを明文化。
 - 2025-09-29: 最終レビュー完了。本ストーリーを [完了] としてクローズ。

## リスク・懸念
- 既存コードの命名差異が広範囲に及ぶ場合、是正コストが増大する。
- DB/JSON の互換性要件（外部ツール/既存データ）により一部ルールを緩和する可能性。

## 関連
- Docs: `docs/architecture/README.md`、`docs/KNOWLEDGE/engineering/comment-style.md`、`docs/REF_STORIES.md`
- Tasks: `tasks/todo_docs.md`
- PR/Issue: 未作成（本ストーリー採択後に作成）
