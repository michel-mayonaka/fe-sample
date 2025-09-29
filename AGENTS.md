# AGENTS — 開発者/エージェント向けガイド（最新）

このリポジトリで作業する人間/エージェント向けの最小実務ガイドです。詳細は各ドキュメントへリンクします（クリックで開けます）。

## クイックリンク
- 命名規約: docs/NAMING.md
- コメント記法: docs/COMMENT_STYLE.md
- アーキテクチャ: docs/ARCHITECTURE.md
- ストーリー運用: docs/REF_STORIES.md
- DB メモ: docs/DB_NOTES.md
- ルート概要: README.md

## よく使うコマンド
- 検証一式: `make mcp`（vet/build/lint/test を実行）
- コンパイル確認: `make check-all`（UI まで）
- テスト: `make test-all`（`-race -cover` 既定）
- 実行: `go run ./cmd/ui_sample`
- ビルド: `go build -o bin/ui_sample ./cmd/ui_sample`

## ストーリー駆動の運用
- 新規作成: `make new-story SLUG=<slug>` → `stories/YYYYMMDD-slug/README.md`
- 作成フロー（要点）:
  - README を先に作成し、目的/スコープ/DoD を確定（制作者レビュー必須）。
  - サブタスクは 1 タスク=1 MD をストーリー配下 `tasks/` に作成（連番推奨）。詳細: docs/REF_STORIES.md
  - 完了/アーカイブ: `make finish-story SLUG=<slug>`（ステータスを自動で [完了] に更新）。

## 命名・コーディング（要約）
- ルール全文: docs/NAMING.md
- コード: UpperCamel（export）/lowerCamel（internal）、初期ismは `ID/HTTP/URL` などを全大文字維持。
- データ: `snake_case`、マスタ `mst_*`／ユーザ `usr_*`、JSON キーはタグで対応。
- アセット: `用途_部位_状態_バリアント@倍率.ext`（例: `status_panel_dark@2x.png`）。
- 列挙: 型付き `iota`、`XxxUnknown` を 0 に据える。
- 汎用名は避ける: `util`/`helpers` 等は禁止。例: 乱数→`internal/game/rng`、矩形判定→`rect*_`。

## コメント・ドキュメント
- エクスポート識別子は 1 行 GoDoc 必須（revive の exported を満たす）。
- 詳述は Markdown に集約（コード内は最小限）。
- 新規/変更 API は docs/API.md に追記。設計背景は docs/ARCHITECTURE.md、命名/表記は docs/NAMING.md に整合。

## Lint/テスト基準
- Lint は `.golangci.yml` に準拠。ローカルは非必須扱いだが、可能な限り 0 issue を維持。
- 変更時は `make mcp` を実施し、`pkg/...` と `internal/usecase` のユニットテストが通ること。

## コミット規約（日本語・Conventional Commits）
- 例: `feat(ui): ステータス画面にHPバーを追加`、`docs(naming): NAMING に詳細規則を追記`。
- 本文に「目的／影響範囲／検証手順」を簡潔に。ストーリー作業は `[story]` をサマリに含めてもよい。

## 変更の流儀（エージェント向け）
- 小さく進める: ストーリー→サブタスク→実装→`make mcp`→進捗ログ更新。
- 命名は docs/NAMING.md を最優先。`util/helpers/common` 等の曖昧名は採らない。
- 既存規約と衝突する場合はストーリーで先に方針決定→ドキュメントを更新してから変更する。
- 影響の大きいリネームは試験的に 1〜2 箇所から始め、DoD で運用確認して段階適用。

---

以下は参考の詳細（上記リンク先を正とします）。

### 構成（サマリ）
- `cmd/ui_sample/`: エントリ・ゲームループ
- `internal/game/`: Scene/Service/UI/Provider/Usecase 依存の境界
- `internal/model` / `internal/user`: マスタ/ユーザデータのスキーマとローダ
- `db/master` / `db/user`: JSON 置き場（将来 SQLite に移行）
- `pkg/game`: ドメインロジック（テスト対象）
- `assets/`: 画像・フォント・音源

### セキュリティ/設定（任意）
- アセットは `embed` 検討、秘密情報は `.env`/環境変数で管理。
