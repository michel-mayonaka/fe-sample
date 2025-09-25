# Repository Guidelines

本ドキュメントは、本リポジトリに貢献する際の最小限の実務ガイドです。UI サンプル用途を想定しています（現状の構成に合わせて適宜更新してください）。

## Project Structure & Module Organization
- `cmd/ui_sample/main.go`: エントリポイント。
- `internal/ui`: UI レイアウト・入力イベント処理。
- `internal/game`: ゲームループ（`Update/Draw/Layout`）や状態管理。
- `assets/`: 画像・フォント・音源等のアセット。
- `pkg/`: 複数バイナリ/ツールから再利用するライブラリ（必要時）。
- `testdata/`: テスト専用アセット。

## Build, Test, and Development Commands
- `go run ./cmd/ui_sample`: ローカル実行。
- `go build -o bin/ui_sample ./cmd/ui_sample`: バイナリ生成。
- `go test ./...`: ユニットテスト一括実行。
- `go test ./... -race -cover`: 競合検出＋カバレッジ計測。
- `go fmt ./...`: 公式フォーマッタで整形。
- `golangci-lint run`: Lint 実行（導入済みの場合）。

## Coding Style & Naming Conventions
- インデント: タブ（Go 標準）。行長は 120 目安。
- パッケージ名: 小文字単語連結（例: `ui`, `gamecore`）。
- ファイル名: 小文字スネーク（例: `button_test.go`）。
- エクスポート識別子: `UpperCamelCase`、GoDoc コメントを付与。
- エラー: `fmt.Errorf("%w", err)` でラップ、文脈は `errors.Join` 等で付加。

## Testing Guidelines
- フレームワーク: 標準 `testing`。
- 命名: テスト `TestXxx`、ベンチ `BenchmarkXxx`、例: `TestButtonHover`。
- カバレッジ目標: 70% 目安（描画はロジック分離やスナップショットで担保）。
- 実行: `go test ./...` を基本、描画依存は条件付きスキップを使用。

## Commit & Pull Request Guidelines
- 言語: 以後、コミットメッセージと PR タイトル/本文は日本語で記述してください。
- コミット規約: Conventional Commits 準拠（日本語サマリ）。種別例: `feat`, `fix`, `docs`, `refactor`, `test`, `chore`。
- 形式例:
  - `feat(ui): ステータス画面にHPバーを追加`
  - `fix(input): Gamepad 接続時の初期化クラッシュを回避`
  - `docs(readme): トラブルシューティングを追記`
- 本文: なぜ必要か、影響範囲、検証手順を簡潔に（72字程度で改行推奨）。
- PR 要件:
  - 目的/背景と主要変更点（箇条書き可）。
  - UI 変更はスクリーンショット/GIF を添付。
  - 動作確認環境（OS/Go 版/ドライバ等）。
  - 関連 Issue を `Closes #123` でリンク。
  - 1 PR は 1 トピック、差分を小さく保つ。

## Security & Configuration Tips (任意)
- アセットは `embed` 利用を推奨（例: `//go:embed assets/*`）。
- 秘密情報は環境変数や `.env` を使用（リポジトリへコミット不可）。
- 開発時は `-tags=ebitendebug` などのデバッグビルドタグを検討。
