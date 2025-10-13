# 02 WASM ビルド手順と Make 整備

## 目的
`GOOS=js GOARCH=wasm` で UI サンプルをビルドし、再現性の高い出力（`site/` など）を得る。

## 方針（ドラフト）
- 出力: `site/ui_sample.wasm`
- 付帯: `site/wasm_exec.js`（`$(go env GOROOT)/misc/wasm/wasm_exec.js` をコピー）
- Make 目標: `make wasm` / `make site`（クリーン含む）

## 手順
1) ビルドコマンド確定: `GOOS=js GOARCH=wasm go build -o site/ui_sample.wasm ./cmd/ui_sample`
2) `wasm_exec.js` を `site/` に複製。
3) 依存アセットのコピー規則（`assets/` など）を定義。
4) `MCP_OFFLINE=1` でも動くよう `-mod=vendor` を考慮。

## 成功条件（DoD）
- ローカルで `site/` が生成され、`.wasm` と `wasm_exec.js` が揃う。
- `make mcp` と競合しない（CI で問題なし）。

