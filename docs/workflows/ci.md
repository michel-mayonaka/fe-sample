# CI 構成（概要）

本プロジェクトの CI は、ローカルでの `make` ベースの検証と整合が取れるように構成されています。

## ジョブ構成
- `smoke-offline`
  - 目的: オフライン再現性の確認（vendor 必須）。
- `build-and-lint`
  - 目的: `make mcp`（非 strict）でロジック検証＋ lint を行う。
- `ui-build-strict`
  - 目的: 依存導入後の UI ビルド厳格チェック（`MCP_STRICT=1 make check-ui` 相当）。

## 補足
- ワークフロー定義ファイル: `.github/workflows/ci.yml`
- ローカルとの対応関係は `docs/workflows/local-dev.md` を参照。

