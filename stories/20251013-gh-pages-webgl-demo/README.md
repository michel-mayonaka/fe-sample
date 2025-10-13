# 2025-10-13 — GitHub Pages で WebGL デモを公開（`20251013-gh-pages-webgl-demo`）

- 状態: 準備完了
- 作成: 2025-10-13
- オーナー: TBD
- 関連: README.md, docs/ARCHITECTURE.md, docs/NAMING.md, docs/REF_STORIES.md, docs/WORKFLOW.md

## 目的
GitHub Pages 上で本リポジトリの WebGL（WASM）デモを一般公開し、ブラウザで手早く試せる導線を提供する。

## 背景
- 体験のハードルを下げ、レビューや共有を容易にする。
- CI 経由で常に最新の `main` を自動デプロイする。

## スコープ
- WASM ビルドの整備（`GOOS=js GOARCH=wasm`）。
- `index.html`（ローダ）と `wasm_exec.js` の配置方針決定。
- GitHub Actions による `gh-pages` ブランチ(or GitHub Pages) への自動デプロイ。
- 動作確認（Chrome/Firefox/Safari/iOS）の最小マトリクス。

## 非スコープ
- 大規模最適化や新機能追加。
- サイト全体のデザイン刷新。

## 成果物（Deliverables）
- 公開 URL（GitHub Pages）とスクリーンショット1枚。
- CI 設定ファイル（デプロイ用ワークフロー）。
- `README.md` に公開リンクとローカル実行手順を追記。

## 完了の定義（DoD）
- GitHub Pages 上の URL でデモが起動し、主要 3 ブラウザ（Chrome/Firefox/Safari）で表示・入力が動作する。
- `main` への push で自動ビルド・自動反映される（手動作業不要）。
- CI は `MCP_STRICT=0 make mcp` を通過し、別ジョブで WASM ビルドが成功。
- `README.md` に利用方法・既知の制約・サポートブラウザを明記。

## リスクと対策
- Safari/iOS の WebGL/オーディオ制限: 互換設定・初回ユーザー操作誘導を明記。
- WASM サイズ肥大: `-ldflags "-s -w"` などの基本最適化、不要アセット除外。
- MIME/パス問題: GH Pages での配置パス/相対パスを事前に統一。

## 依存関係
- Go/Ebitengine の対応バージョン確認。
- GitHub Pages の権限設定（Pages 有効化、`gh-pages` へのデプロイ許可）。

## タスク
1. tasks/01-plan-and-scope.md — 公開方式/構成の最終決定
2. tasks/02-wasm-build-setup.md — WASM ビルド手順と Make タスク整備
3. tasks/03-gh-actions-deploy.md — CI/CD（Pages）ワークフロー追加
4. tasks/04-index-html-and-wasm_exec.md — ローダ HTML と `wasm_exec.js` 方針
5. tasks/05-browser-qa-matrix.md — 動作確認マトリクスと検証観点
6. tasks/06-perf-and-bundle-size.md — サイズ/初回起動時間の予算と対策
7. tasks/07-docs-and-announcement.md — README 追記と告知

## 開始指示（テンプレ）
実装開始時は、以下のように明示してください。

> 20251013-gh-pages-webgl-demo を開始
