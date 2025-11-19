# AGENTS — 開発者/エージェント向けガイド（最新）

このリポジトリで作業する人間/エージェント向けの最小実務ガイドです。詳細は各ドキュメントへリンクします（クリックで開けます）。

## クイックリンク
- 仕様ハブ: docs/SPECS/README.md
- エージェント向け仕様ガイド: docs/SPECS/AGENTS.md
- ops 1枚絵: docs/ops-overview.md
- ナレッジ目次: docs/KNOWLEDGE/README.md
- 命名規約: docs/KNOWLEDGE/engineering/naming.md
- コメント記法: docs/KNOWLEDGE/engineering/comment-style.md
- アーキテクチャ: docs/architecture/README.md
- ストーリー運用: docs/KNOWLEDGE/workflows/stories.md
- ワークフロー概要: docs/KNOWLEDGE/workflows/overview.md
- Vibe‑kanban 運用: docs/KNOWLEDGE/workflows/vibe-kanban.md
- DB メモ: docs/KNOWLEDGE/data/db-notes.md
- ルート概要: README.md
- codex-cloud 実行: docs/KNOWLEDGE/ops/codex-cloud.md

## よく使うコマンド
- 検証一式: `make mcp`（vet/build/lint/test を実行）
- コンパイル確認: `make check-all`（UI まで）
- テスト: `make test-all`（`-race -cover` 既定）
- 実行: `go run ./cmd/ui_sample`
- ビルド: `go build -o bin/ui_sample ./cmd/ui_sample`
 - Backlog生成: `make backlog-index`（discovery/accepted から自動生成）
 - Discovery索引: `make discovery-index`
 - スモーク（最短）: `make smoke`
 - オフライン包括: `MCP_OFFLINE=1 make offline`
 - 生成物クリーン: `make clean`

### CI/環境変数（UIビルドの扱い）
- `MCP_STRICT`:
  - `0`（既定/CIのmcpジョブ）: 環境依存の UI ビルド失敗（例: `X11/Xlib.h` 欠如）を検出したらスキップし、ロジック層の検証を継続。
  - `1`（厳格）: 上記も失敗として扱う。Linux では X11/GL の開発パッケージ導入が必要。
- `MCP_OFFLINE`: `1` で `-mod=vendor` を強制し、オフライン開発に最適化。

ローカル検証の例:
```sh
# 非strict（既定）: mcp を通す
make mcp

# 厳格: UIも必ず検証（Linuxは依存導入後に）
MCP_STRICT=1 make check-ui
```

## ストーリー駆動の運用
- 新規作成: `make new-story SLUG=<slug>` → `stories/YYYYMMDD-slug/README.md`
- 作成フロー（要点）:
  - README を先に作成し、目的/スコープ/DoD を確定（制作者レビュー必須）。
  - サブタスクは 1 タスク=1 MD をストーリー配下 `tasks/` に作成（連番推奨）。詳細: docs/KNOWLEDGE/workflows/stories.md
  - 完了/アーカイブ: `make finish-story SLUG=<slug>`（ステータスを自動で [完了] に更新）。

### Discovery（Backlogコンフリクト回避）
- 並列作業での課題発見は `stories/discovery/` に 1課題=1MD で起票。
- コマンド: `make new-discovery SLUG=<slug> [TITLE=..] [PRIORITY=Px] [STORY=YYYYMMDD-slug]`
- 採択時は `make promote-discovery FILE=...` で Backlog に昇格（discovery は accepted/ へ移動）。
- 見送り時は `make decline-discovery FILE=...` で declined/ へ移動し、理由を記録。
- 索引/検証: `make discovery-index`, `make validate-meta`（警告のみ）

## 作業開始ルール（重要）
- ストーリー作成直後の実装着手は禁止。
- 実装開始には、明示的な開始指示が必要（例: 「20250930-xxx を開始」）。
- 開始指示前に許可される作業: ディスカッション/レビュー/ストーリー配下の Markdown 更新のみ。
- コード/ドキュメント本体（`docs/*` や `*.go` 等）の変更は開始指示後に行う。
- コード変更を含むPRは `make mcp` 成功が前提。マージは Squash を既定とする。

## 命名・コーディング（要約）
- ルール全文: docs/KNOWLEDGE/engineering/naming.md
- コード: UpperCamel（export）/lowerCamel（internal）、初期ismは `ID/HTTP/URL` などを全大文字維持。
- データ: `snake_case`、マスタ `mst_*`／ユーザ `usr_*`、JSON キーはタグで対応。
- アセット: `用途_部位_状態_バリアント@倍率.ext`（例: `status_panel_dark@2x.png`）。
- 列挙: 型付き `iota`、`XxxUnknown` を 0 に据える。
- 汎用名は避ける: `util`/`helpers` 等は禁止。例: 乱数→`internal/game/rng`、矩形判定→`rect*_`。

## コメント・ドキュメント
- 仕様全体の構成方針は docs/KNOWLEDGE/meta/docs-structure.md を参照。
- エクスポート識別子は 1 行 GoDoc 必須（revive の exported を満たす）。
- 詳述は Markdown に集約（コード内は最小限）。
- 新規/変更 API は docs/SPECS/reference/api.md に追記。設計背景は docs/architecture/README.md、命名/表記は docs/KNOWLEDGE/engineering/naming.md に整合。

## Lint/テスト基準
- Lint は `.golangci.yml` に準拠。ローカルは非必須扱いだが、可能な限り 0 issue を維持。
- 変更時は `make mcp` を実施し、`pkg/...` と `internal/usecase` のユニットテストが通ること。

### CI 方針（UI 依存）
- `make mcp` は非strict（`MCP_STRICT=0`）で実行し、UI ビルドが環境依存で失敗する場合はスキップ。
- 別ジョブ `ui-build-strict` で Linux に X11/GL 依存（例: `xorg-dev libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev`）を導入し、`MCP_STRICT=1 make check-ui` を実行して UI を厳格検証。

## コミット規約（日本語・Conventional Commits）
- 例: `feat(ui): ステータス画面にHPバーを追加`、`docs(naming): NAMING に詳細規則を追記`。
- 本文に「目的／影響範囲／検証手順」を簡潔に。ストーリー作業は `[story]` をサマリに含めてもよい。

## 変更の流儀（エージェント向け）
- 小さく進める: ストーリー→サブタスク→実装→`make mcp`→進捗ログ更新。
- 命名は docs/KNOWLEDGE/engineering/naming.md を最優先。`util/helpers/common` 等の曖昧名は採らない。
- 既存規約と衝突する場合はストーリーで先に方針決定→ドキュメントを更新してから変更する。
- 影響の大きいリネームは試験的に 1〜2 箇所から始め、DoD で運用確認して段階適用。

---

以下は参考の詳細（上記リンク先を正とします）。

### 構成（サマリ）
- `cmd/ui_sample/`: エントリ・ゲームループ
- `internal/game/`: Scene/Service/UI/Provider/Usecase 依存の境界
- `internal/model` / `internal/model/user`: マスタ/ユーザデータの純粋モデル（I/Oなし）
- `internal/infra/userfs`: ユーザデータのJSON入出力
- `db/master` / `db/user`: JSON 置き場（将来 SQLite に移行）
- `pkg/game`: ドメインロジック（テスト対象）
- `assets/`: 画像・フォント・音源

### セキュリティ/設定（任意）
- アセットは `embed` 検討、秘密情報は `.env`/環境変数で管理。
