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

# 実行（フルHD 1920x1080）
go run ./cmd/ui_sample
```

## ビルド
```sh
go build -o bin/ui_sample ./cmd/ui_sample
```
※ Apple Silicon では `GOARCH=arm64` が既定です。意図しない `amd64` クロスビルドは避けてください。

## 表示仕様（解像度）
- 論理解像度: 1920×1080（`Layout`）
- ウィンドウ: 1920×1080（可変リサイズ）
- フォント: Go Regular（タイトル36px、本文24px、注釈18px）
- ポートレート枠: 320×320（画像は等比縮小・線形補間）

## 操作
- `H`: ヘルプ表示 / `Esc`: ヘルプを閉じる
- `Backspace`: サンプル値を再読み込み

## 構成
- `cmd/ui_sample/main.go`: エントリ・ゲームループ
- `internal/ui/ui.go`: パネル/テキスト/HPバー/能力値の描画、ポートレート画像の表示
- `assets/`: 画像等を追加する場合に利用（例: `assets/01_iris.png`）
- `internal/model/`: キャラクターマスタのモデルとJSONローダー
- `assets/master/characters.json`: キャラクターのマスタデータ（ID索引）
- `internal/user/`: ユーザ（セーブ）データのモデルとJSONローダー
- `assets/user/party.json`: 現在のユーザ状態（上書き用）

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

## 画像の追加（ポートレート表示）
- `assets/01_iris.png` を配置すると、左上のポートレート枠に表示されます。
- 別名で使う場合は `internal/ui/ui.go` の `SampleUnit()` で読み込みパスを変更してください。

## データ構成（マスタ/ユーザ）
- マスタ: `assets/master/characters.json`
  - 役割: 初期値のみを保持（名前/クラス/成長率/初期装備の上限 など）
  - 注意: レベルごとの能力は保持しない（成長率に依存し可変のため）
- ユーザ: `assets/user/party.json`
  - 役割: 現在値を保持（Lv/Exp/HP/能力値/装備残耐久 など）
- 合成: UI はマスタを読み、ユーザ値があれば上書きして表示

将来: SQLite へ移行予定（`docs/DB_NOTES.md` 参照）

デザイン調整は `internal/ui/ui.go` の色・座標・フォントサイズを編集してください。必要に応じて画像/TTF を `assets/` に追加し、`embed` で組み込み可能です。
