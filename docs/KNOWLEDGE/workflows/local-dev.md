# ローカル開発・検証フロー

ローカル環境でのビルド/テスト/検証手順をまとめます。codex-cloud の詳細な設定や挙動は `docs/KNOWLEDGE/ops/codex-cloud.md` を参照してください。

## よく使う make ターゲット
- 最短スモーク（論理層のみ・vendor 前提）: `make smoke`
- 包括検証（mcp をオフラインで実行）: `MCP_OFFLINE=1 make offline`
- 通常（オンライン）: `make mcp`
- UI 厳格検証（Linux で依存導入後）: `MCP_STRICT=1 make check-ui`

## 出力ディレクトリ
- バイナリ: `out/bin/`
- ログ: `out/logs/`
- カバレッジ: `out/coverage/`

詳細なオプションや codex-cloud 周りの情報は `docs/KNOWLEDGE/ops/codex-cloud.md` を参照してください。
