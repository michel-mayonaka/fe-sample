# 01 計画とスコープ確定

ステータス: [完了]

## 目的
GitHub Pages での公開方式・配置構成・命名を確定し、以降の作業判断を一本化する。

## 決定事項（最終）
- 公開方式: GitHub Pages（Actions デプロイ）。ブランチ直配置（`/docs`）は採用しない。
- 公開パス: ルート直下（`/`）。相対参照で `index.html`/`ui_sample.wasm`/`wasm_exec.js` を同階層に配置。
- 成果物配置: `site/` 直下（CIで artifact→Pages）。画像は当面 `assets/01_iris.png` のみ配布。
- ストーリー識別子: `20251013-gh-pages-webgl-demo`。

## 手順
1) 代替案比較の結果、Pages（Actions）に確定。
2) 公開 URL は Actions 実行後の `Deploy to GitHub Pages` 出力を参照（環境 `github-pages`）。
3) 成果物ディレクトリ名は `site/` に確定。
4) Pages 有効化はリポジトリ設定で実施（初回のみ）。

## 成功条件（DoD）
- 方式・パス・ディレクトリが README に反映されている。
- 影響範囲（CI・ビルド・アセット）が整理されている。
