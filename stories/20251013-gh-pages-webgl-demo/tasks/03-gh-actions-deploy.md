# 03 GitHub Actions（Pages）デプロイ

ステータス: [完了]

## 目的
`main` 更新で自動的に GitHub Pages にデモを公開する CI/CD を構築する。

## 方針（ドラフト）
- Pages ワークフロー: `actions/upload-pages-artifact` → `actions/deploy-pages`。
- 権限: `permissions: pages: write, id-token: write`。
- 成果物: `site/` ディレクトリをアップロード。

## 手順
1) `setup-go` で Go を用意 → 02 の手順で `site/` を生成。
2) `upload-pages-artifact` で `site/` を成果物化。
3) `deploy-pages` で公開。環境名 `github-pages`。
4) `main` への push と手動 `workflow_dispatch` を用意。

## 成功条件（DoD）
- `main` へ push 後、Pages が自動更新される。
- ジョブが `MCP_STRICT=0` の既存 CI と共存し、失敗しない。

## 備考
- 初回のみ、リポジトリの Settings → Pages で "Build and deployment = GitHub Actions" を選択してください。
- 公開URLはデプロイ完了ステップの `page_url` 出力に表示されます（README 更新は別タスク 07 で実施）。
