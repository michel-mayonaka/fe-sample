# CI 構成（概要）

本プロジェクトの CI は、ローカルでの `make` ベースの検証と整合が取れるように構成されています。

## ジョブ構成
- `smoke-offline`
  - 目的: オフライン再現性の確認（vendor 必須）。
- `build-and-lint`
  - 目的: `make mcp`（非 strict）でロジック検証＋ lint を行う。
  - 追加: Story 索引生成（`make story-index`）とメタ検証（`./scripts/validate_meta.sh`）を実行し、ストーリー関連メタの整合を確認する。
- `ui-build-strict`
  - 目的: 依存導入後の UI ビルド厳格チェック（`MCP_STRICT=1 make check-ui` 相当）。

## Story 検証/索引生成
- 完了ストーリー索引:
  - ローカル: ストーリーを `finish` へアーカイブしたあとに `make story-index` を実行し、`stories/finish/INDEX.md` を更新してコミットする。
  - CI: `build-and-lint` ジョブ内で `make story-index` を実行し、生成スクリプト自体の正常動作を確認する（CI では生成結果をコミットしない）。
- メタデータ検証:
  - スクリプト: `./scripts/validate_meta.sh [paths...]` は Story/Discovery の `ステータス`/`担当`/`開始` キー有無をチェックし、欠落があれば `[warn]` を出しつつ終了コード `1` を返す。
  - ローカル: `make validate-meta` は `./scripts/validate_meta.sh || true` を呼び、警告のみ表示してローカル開発を止めない運用とする。
  - CI（非strict）: `build-and-lint` ジョブでは `./scripts/validate_meta.sh` の結果を警告としてログに残すのみとし、既存ストーリーのメタ整理が完了するまでは CI fail にはしない。
  - 将来の strict 化案:
    - PR で変更された `stories/**/*.md` のみを引数に `./scripts/validate_meta.sh` へ渡し、終了コードを CI の fail 条件として扱う。

## 補足
- ワークフロー定義ファイル: `.github/workflows/ci.yml`
- ローカルとの対応関係は `docs/workflows/local-dev.md` を参照。
