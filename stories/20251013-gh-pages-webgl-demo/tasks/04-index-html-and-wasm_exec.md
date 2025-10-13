# 04 ローダ HTML と `wasm_exec.js`

ステータス: [完了]

## 目的
ブラウザで WASM を起動する最小の `index.html` と `wasm_exec.js` の取り扱いを確定する。

## 方針（ドラフト）
- 直下に `index.html` を置き、`wasm_exec.js` を相対参照。
- `WebAssembly.instantiateStreaming` 不可時は `fetch` + `instantiate` にフォールバック。
- 既知の制約（初回操作で音声解放など）を画面に明記。

## 手順
1) HTML 雛形を用意（タイトル、キャンバス領域、起動スクリプト）。
2) エラー時のユーザー向けメッセージ（対応ブラウザ/手順）を追加。
3) ベースパス配下でも動くよう相対パス検証。

## 成功条件（DoD）
- ローカル `site/index.html` で WASM 起動を確認できる。
- GitHub Pages で同一構成がそのまま動作する。

## 実装
- `web/index.html` を追加し、`make site` 時に `site/index.html` として配置。
- `instantiateStreaming` 不可時のフォールバック実装済み。
