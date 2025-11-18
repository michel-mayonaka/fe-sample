# 01_現状調査 — main と internal/app の責務/依存棚卸し

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 00:22:56 +0900

## 目的
- `cmd/ui_sample/main.go` と `internal/app` 周辺の責務と依存関係を棚卸しし、どの処理を app 層へ集約すべきか明らかにする。

## 完了条件（DoD）
- [x] main に残っている入力処理・状態遷移・初期化ロジックが一覧化されている。
- [x] `internal/app` 側の責務と、現在の依存関係（Repo/Usecase/UI など）が整理されている。
- [x] レイヤリング改善の観点で問題になっている箇所（UI からの直接I/O 等）が洗い出されている。

## 作業手順（概略）
- main と `internal/app` のコードを読み、処理フローと依存関係を簡略図や箇条書きでまとめる。
- 入力/状態遷移/初期化/終了処理などの観点で、責務を分類する。
- レイヤリング上の違和感がある箇所をメモし、次タスクのインプットとする。

## 進捗ログ
- 2025-11-19 00:22:56 +0900: タスク作成。
- 2025-11-19 01:47:16 +0900: main と `internal/app` のコード読みを開始し、現状棚卸しメモの作成を進行。
- 2025-11-19 02:05:27 +0900: 棚卸しメモが DoD を満たしたため完了。

## 依存／ブロッカー
- `cmd/ui_sample` と `internal/app` 一帯へのアクセス。

## 成果物リンク
- 現状調査メモ

## 現状調査メモ

### main（`cmd/ui_sample/main.go`）
- `main()` は `ebiten.RunGame(gameapp.NewUIAppGame())` を呼ぶだけで、入力/状態遷移/初期化ロジックはすべて `internal/game/app` 側に移譲済み。
- エントリポイント自体は極小だが、`internal/game/app` が UI・I/O・依存初期化をすべて抱えているため、別アプリケーション（CLI/テスト用 headless 等）からの再利用は難しい構造になっている。

### `internal/game/app` の構造と責務

#### `NewUIAppGame`（`core.go`）
- 乱数と抽象入力レイアウトを生成し、`Backspace` を `Menu` にも割当てるなど UI 専用のキーマップをここで決定。`internal/game/provider/input/ebiten.Source` を直接生成しており、物理入力の差し替えポイントが `game` 配下に固定化されている。
- `config.Default*` で決まる JSON ファイル群から `repo.NewJSON*Repo` を都度生成。失敗時は `nil` 許容としてサンプルユニット fallback するが、エラーの扱いはここで握っており、アプリ層としてのリトライや設定変更の余地がない。
- `usecase.New(...)` → `gdata.SetProvider(a)` でゲームロジック側に依存性を注入しつつ、そのまま UI 用ユニット（`uiadapter.BuildUnitsFromProvider`）まで構築。リポジトリ/ユースケース/アセット読み込みの境界が混在。
- 画面メトリクスは `uicore.SetBaseResolution` の後に `config/uimetrics.LoadOrDefault` を呼び、`uicore.Metrics` フィールドへ手動コピーして `uicore.ApplyMetrics`。同じコピー処理が後述のホットリロードでも重複している。
- `scenes.Env` と `scenes.Session` をここで初期化し、`Runner.Stack` に `character_list.NewList(env)` を push。初期シーン決定や Session 構築まで app パッケージに閉じており、main からの差し替え余地はない。
- ウィンドウサイズ/TPS/タイトル設定（`SetupWindow`/`ebiten.SetWindowTitle`）もここで一括実行され、UI 実装が Ebiten 実行環境に直接ロックされている。

#### `Game`（`game.go`）
- `Runner`・`InputSrc`・`EdgeReader`・`uinput.Reader`・`Env`・`Session` など UI 実行に必要な状態をすべて保持し、`ebiten.Game` を実装。main 側は `Game` の参照を得ても個別コンポーネントを差し換えられない。
- `Update()` では `ginput.Source` → `EdgeReader` → `uinput.Reader` まで束ね、`game.Ctx` を構築して `Runner.Update` へ渡す。マウス座標取得もここで `interface{ Position() }` に型アサートして処理している。
- `updateGlobalToggles()` に Backspace 長押し時のデータ/メトリクス再読み込み、`assets.Clear()`、`uiadapter.BuildUnitsFromProvider` の再実行、ヘルプ表示トグル（`ebiten.IsKeyPressed(KeyH)` と `uinput.Cancel`) などグローバルトリガが集約。抽象入力と Ebiten 生キー判定が混在しており、今後キー割当をアプリ層で変更することが難しい。
- 画面描画前に `uicore.UpdateMetricsFromWindow` / `uicore.MaybeUpdateFontFaces` を呼ぶなど、UI の低レベル初期化も `Game` に集中。

#### `Runner`（`runner.go`）
- `game.SceneStack` への Push/Pop/Draw を吸収する薄いラッパーで、`AfterUpdate` に `ShouldPop()` 実装有無を委譲。シーン遷移のフローも `internal/game/app` に内包されており、app 層を跨いだ制御ができない。

#### Env / Session（`internal/game/scenes/env.go`, `session.go`）
- `Env` は `Data/Battle/Inv` の Usecase Port、`UserPath`、`RNG`、`Session` を保持。UI から直接ユーザディレクトリやランダムソースに触れられるため、アプリ固有のコンテキストが game パッケージへ漏れている。
- `Session` は `Units` リストやポップアップ状態など描画寄りの一時状態をまとめており、app 層としての永続ストレージ/設定との分離はできているが、生成箇所が `NewUIAppGame` 固定のためストレージ戦略を変更しづらい。

### 依存先まとめ（`internal/game/app` からの主な import）
- インフラ: `internal/config`, `internal/config/uimetrics`, `internal/repo`, `internal/usecase`, `internal/assets`.
- ゲームロジック/UI: `internal/game/data`, `internal/game/scenes`, `internal/game/service/ui`, `internal/game/ui/adapter`, `internal/game/ui/input`, `pkg/game/input`.
- プラットフォーム: `github.com/hajimehoshi/ebiten/v2`（ウィンドウ・TPS・key polling）、`math/rand`/`time`.
- 結果として「アプリ起動の全責務」が `internal/game/app` に集中し、main 層から細粒度に依存を差し替えることができない。

### レイヤリング上の気になる点
1. **アプリ層とゲーム層の境界不明瞭**: JSON Repo や `usecase.New`、`gdata.SetProvider`、UI ユニット組み立てまで `internal/game/app` が行い、`internal/game` 配下に I/O とインフラ依存が入り込んでいる。
2. **入力管理の一元化不足**: `updateGlobalToggles` で `ebiten.IsKeyPressed(KeyH)` を直接参照する一方、他の入力は `uinput.Reader` 経由。アプリレベルでショートカットを差し替える余地がなく、抽象入力レイヤを通らない経路が存在する。
3. **メトリクス適用処理の重複**: 初期化時と Backspace 長押し時の両方で `uicore.Metrics` へのフィールドコピーが手書きで重複。`internal/app` 層で共有ヘルパを設けない限り、仕様変更時にずれが生じるリスクが高い。
4. **グローバル状態管理の集中**: `Game` が `assets.Clear()`・`Env.Data.ReloadData()`・`uiadapter.BuildUnitsFromProvider()` 等の副作用を直接呼び出しており、UI 以外のコンポーネント（例: CLI, tests, dedicated app）から共通ロジックを再利用しづらい。
5. **初期シーン/Session 固定**: `Runner.Stack.Push(characterlist.NewList(env))` が `internal/game/app` 内にハードコードされ、main 側で別シーンを指定したり、`Env.Session` を差し替える拡張の余地がない。
