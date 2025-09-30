# Ebiten UI Sample — FE風ステータス画面

ファイアーエムブレムのステータス画面風 UI を Ebiten で描画する最小サンプルです。画像や外部フォントは使わず、`basicfont` とベクタ描画のみで構成しています。

## 要件
- Go 1.25 以上（推奨: 最新 1.25 系）
- 初回は依存取得のためネットワーク接続が必要

補足（toolchain/代替手順）
- Go 1.21 以上をお使いの場合は、`go env -w GOTOOLCHAIN=auto` を設定すると `go.mod` の `go 1.25.0` に合わせてツールチェーンを自動取得できます。
- それ未満の環境では Go を 1.25 系へ更新してください（例: macOS/Homebrew は `brew update && brew upgrade go`）。

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

## コンパイルチェック
- 依存含めて静的検査＋ビルド検証（バイナリは残さない）
```sh
make check
```

## Lint（golangci-lint）
- インストール（macOS/Homebrew 例）: `brew install golangci-lint`
- 実行: `make lint` または `golangci-lint run`
- 整形: `make fmt`

## テスト
- 既定（ロジック/ユースケースのみ）: `make test-all`
- UI 関連含む（adapter/levelup 等）: `make test-all-ui`

CI（GitHub Actions）
- 本リポジトリは `make mcp`（vet/build/lint）を CI で実行します。
- Go 1.25.x を固定し、Go のビルドキャッシュ/モジュールキャッシュを保存します。
- GUI 依存のビルドは CI 環境で失敗しうるため、既定でスキップ（`MCP_STRICT=0`）します。

## 表示仕様（解像度）
- 論理解像度: 1920×1080（`Layout`）
- ウィンドウ: 1920×1080（可変リサイズ）
- フォント: Go Regular（タイトル36px、本文24px、注釈18px）
- ポートレート枠: 320×320（画像は等比縮小・線形補間）

## 操作
- `H`: ヘルプ表示 / `Esc`: ヘルプを閉じる
- `Backspace`: サンプル値を再読み込み

## 構成
- `cmd/ui_sample/`: エントリ・ゲームループ
- `internal/game/scenes/`: 画面（一覧/ステータス/在庫/模擬戦）とポート定義
- `internal/game/ui/`
  - `layout/`: レイアウト計算（矩形/サイズ）
  - `draw/`: 描画関数（見た目）
  - `view/`: 表示用データ型（行モデルなど）
  - `adapter/`: view-model 生成（テーブル/ユーザ→view、`PortraitLoader` 抽象）
- `internal/game/service/ui/`: テキスト/パネル等のUIユーティリティ、`SampleUnit()` など
- `internal/game/service/levelup/`: レベルアップの抽選/反映ロジック
- `pkg/game/geom/`: 幾何の純ロジック（`RectContains` など、テスト付き）
- `assets/`: 画像等を追加する場合に利用（例: `assets/01_iris.png`）
- `internal/model/`: マスタ定義とJSONローダー
- `internal/model/user/`: ユーザ（セーブ）データの純粋モデル
- `internal/infra/userfs/`: ユーザデータのJSON入出力（バックエンド）
- `db/master/*.json`: 各種マスタ（`mst_`）
- `db/user/*.json`: ユーザ状態（`usr_`）

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
- 別名で使う場合は `internal/game/service/ui/load.go` の `SampleUnit()` の読み込みパスを変更してください。

## データ構成（mst_/usr_）
- マスタ: `db/master/mst_characters.json`
  - 役割: 初期値のみを保持（名前/クラス/成長率/初期装備の上限 など）
  - 注意: レベルごとの能力は保持しない（成長率に依存し可変のため）
- ユーザ: `db/user/usr_characters.json`
  - 役割: 現在値を保持（Lv/Exp/HP/能力値/装備残耐久 など）
- 表示: UI はユーザテーブルのみで構築（usr_）。マスタは初期投入用。

将来: SQLite へ移行予定（`docs/DB_NOTES.md` 参照）

## ドキュメント
- 命名規約: `docs/NAMING.md`
- コメント記法: `docs/COMMENT_STYLE.md`
- アーキテクチャ: `docs/ARCHITECTURE.md`
- ストーリー運用: `docs/REF_STORIES.md`
- ワークフロー: `docs/WORKFLOW.md`
- DB メモ: `docs/DB_NOTES.md`

## 戦闘画面（簡易）
- ステータス画面右下の「戦闘へ」→ 戦闘プレビュー
- 「戦闘開始」で1ラウンド（攻撃→反撃）を解決し、HP/耐久を `db/user/usr_characters.json` に保存
- ダメージ計算: `攻撃 = 力 + 武器威力 - 相手守備`（命中は簡易式）

デザイン調整は `internal/game/service/ui` と `internal/game/ui/{layout,draw}` の色・座標・フォントサイズを編集してください。必要に応じて画像/TTF を `assets/` に追加し、`embed` で組み込み可能です。
## オフライン実行
- 一度オンラインで `make vendor-sync` を実行して `vendor/` を更新し、コミットします。
- 以降はオフラインで以下を実行できます。
```sh
GOPROXY=off MCP_OFFLINE=1 make mcp
```
