# Ebiten UI Sample — FE風ステータス画面

ファイアーエムブレムのステータス画面風 UI を Ebiten で描画する最小サンプルです。画像や外部フォントは使わず、`basicfont` とベクタ描画のみで構成しています。

## 要件
- Go 1.22 以上（推奨: 1.25 系）
- 初回は依存取得のためネットワーク接続が必要

## セットアップ & 実行
```sh
# Go バージョン確認（古い場合は更新）
go version

# 依存取得（go.sum 生成）
go mod tidy

# 実行
go run ./cmd/ui_sample
```

## ビルド
```sh
go build -o bin/ui_sample ./cmd/ui_sample
```
※ Apple Silicon では `GOARCH=arm64` が既定です。意図しない `amd64` クロスビルドは避けてください。

## 操作
- `H`: ヘルプ表示 / `Esc`: ヘルプを閉じる
- `Backspace`: サンプル値を再読み込み

## 構成
- `cmd/ui_sample/main.go`: エントリ・ゲームループ
- `internal/ui/ui.go`: パネル/テキスト/HPバー/能力値の描画
- `assets/`: 画像等を追加する場合に利用

## トラブルシューティング
- go.mod の Go 版エラー（例:「go 1.22 だが最大 1.17」）
  - Go を更新（macOS/Homebrew 例）: `brew update && brew upgrade && brew install go`
- 実行時 SIGSEGV（purego/gamepad 周辺）
  - 依存更新: `go get github.com/hajimehoshi/ebiten/v2@latest && go get github.com/ebitengine/purego@latest && go mod tidy`
  - Apple Silicon で `GOARCH` が `arm64` であることを確認: `go env GOARCH`
  - 外部コントローラを一旦外して起動、改善後に接続
  - それでも改善しない場合: `go clean -modcache && go mod tidy`
- Linux でビルド失敗
  - 必要パッケージ例（Debian/Ubuntu）: `sudo apt-get install -y libx11-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libasound2-dev`

デザイン調整は `internal/ui/ui.go` の色・座標を編集してください。必要に応じて画像/TTF を `assets/` に追加し、`embed` で組み込み可能です。
