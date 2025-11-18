# 02_レイヤリング方針の整理 — 望ましい責務分担と依存方向

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 00:22:56 +0900

## 目的
- main/app 間の望ましい責務分担と依存方向を整理し、今後の実装時に参照できるレイヤリング方針をまとめる。

## 完了条件（DoD）
- [x] main が担うべき最小限の責務（例: 起動・DI・エントリポイント）が定義されている。
- [x] `internal/app` が担うべき責務（入力処理・状態遷移など）が一覧化されている。
- [x] 依存方向（main → app → usecase/repo/ui など）がテキストまたは簡易図で表現されている。

## 作業手順（概略）
- `01_現状調査` の結果をもとに、責務と依存を理想的な形に再配置する案を考える。
- 既存の docs/ARCHITECTURE.md と整合するように、用語やレイヤ名を合わせる。
- 実装時の指針として使える程度の粒度で、方針をドキュメント化する。

## 進捗ログ
- 2025-11-19 00:22:56 +0900: タスク作成。
- 2025-11-19 01:58:30 +0900: Task 01 の棚卸し結果と設計ドキュメントを踏まえ、理想的なレイヤリング整理を開始。
- 2025-11-19 02:05:00 +0900: main/app の責務分担と依存方向の方針を確定し、DoD を満たしたため完了。

## 依存／ブロッカー
- `01_現状調査` の完了。

## 成果物リンク
- レイヤリング方針メモ/設計スケッチ

## レイヤリング方針メモ

### 全体方針
- main（`cmd/ui_sample`）は「起動設定＋DIの入口」に徹し、UI 実装・状態機構・入力処理に関するコードは `internal/app` へ委譲する。main から直接 `internal/game` 以下へ触れない。
- `internal/app` は「アプリケーションレイヤ」として 1) ブートストラップ（依存の準備）、2) ランタイム制御（入力/状態/SceneStack）、3) 外部境界（設定・ホットリロード）を分離し、UI（描画）/Usecase（ビジネス）/Repo（永続化）との矢印が一方向になるよう整理する。
- 依存の流れは `cmd/ui_sample` → `internal/app` → `internal/usecase` → `internal/repo`/`db` とし、`internal/game/ui` へは `internal/app` が橋渡し役になる（UI コード自身は repo/usecase を直参照しない）。

### main が担う責務（最小コア）
1. **起動コンテキスト決定**: 画面モード（ウィンドウ/フルスクリーン）、ユーザディレクトリ、実行モード（headless/offline）など CLI 引数や環境変数から決める。
2. **DI 初期化**: `internal/app` のブートストラップ API（例: `app.New(config)`）を呼び、戻り値として `ebiten.Game`（または `app.App`）を取得するだけにする。各 repo/usecase/config の詳細は app 側で吸収。
3. **ランタイム開始/終了フック**: `ebiten.RunGame` や将来の headless ランナーを選択し、終了コード処理を行う。ログ/トレース設定など OS 依存の処理も main に置く。
4. **DI 設定の注入のみ**: 設定ファイルパスやフィーチャートグル等を構造体で app に渡し、main 内で状態を保持しない。

### `internal/app` が担う責務
1. **Bootstrap（初期化レイヤ）**
   - Repo/Usecase の生成と例外処理（JSON 読み込みエラーの扱い、fallback データ供給）。
   - `gdata.SetProvider` や `uiadapter.BuildUnitsFromProvider` 等、game 層への依存注入をまとめる。
   - `Env`/`Session` の生成と初期 Scene 選択を統合し、main からのシーン差し替え要求に応えられる API（例: `app.WithInitialScene(...)`）を用意する。
2. **Runtime 制御**
   - 入力: `ginput.Layout` の決定、`provider/input` ソース生成、`uinput.Reader` ラップまでを `internal/app/input` サブパッケージにまとめ、抽象入力しか UI へ流さない。
   - 状態遷移: `Runner` や SceneStack の管理、`ShouldPop` 判定、グローバル Intent（例: ヘルプトグル）を app 層で検出し、Scene に通知する仕組みを提供する。
   - ホットリロード: メトリクス再適用・`assets.Clear()`・`uiadapter.BuildUnitsFromProvider` 等の副作用を `internal/app/reload` のようなユーティリティに隔離し、`Game` から直接 infra を呼ばない。
3. **外部境界の橋渡し**
   - 設定ファイル／ユーザデータパスの決定、`config.Default*` の上書き受け取り。
   - Ebiten へのプラットフォーム依存API呼び出し（ウィンドウ/TPS/タイトル）を `internal/app/platform` に閉じ、main からは `app.SetupWindow(settings)` を呼ぶだけにする。
   - Env へ渡す `UserPath` や `RNG` を app が保持し、Scene 側からは `Env` 経由でのみ取得できるよう統制する。

### 入力/状態/描画の分離ポリシー
- 入力は「物理入力」→「抽象入力」→「Scene Intent」の 3 層で扱う。物理入力（Ebiten）は `internal/app/input/provider`、抽象入力（`uinput.Reader`）は `internal/app/input/adapter`、Scene Intent 生成は Scene 各自が担当する。
- グローバルキー（ヘルプ、リロードなど）は `internal/app` で抽象アクションとして先に消費し、Scene へ副作用の結果（例: `app.Events.Reloaded`) を送る。こうすることで UI が Ebiten の生キー群を直接参照しない。

### 依存方向（文章図）
```
cmd/ui_sample (main)
    ↓ (設定/DIのみ)
internal/app
    ↓            ↘
internal/game/app (SceneStack, Env, Session 制御)
    ↓                ↘
internal/game/scenes   internal/game/ui/*
    ↓                      ↓
internal/usecase --------→ gdata.Provider()
    ↓
internal/repo + db/json
```
- `internal/app` は `internal/game/app`（もしくはその後継）を内包する形で整備し、`internal/game` から `repo`/`usecase` へ向かう依存が逆流しないようにする。
- UI（描画層）が直接 `internal/repo` や `config` に触れないよう、Env/Session/Provider を窓口にする。

### 実装へ向けた具体的な分担案
1. `internal/app/bootstrap` に `type Config struct { UserPath string; AssetsPath string; LayoutOverride *ginput.Layout; InitialScene scene.Factory }` を用意し、main からは Config を構成して渡すだけにする。
2. `internal/app/metrics` にメトリクス変換ヘルパを配置し、初期化とリロードの両方から利用。`copyMetrics(dst *uicore.Metrics, src cuim.Metrics)` のように一本化して重複を排除。
3. `internal/app/reload` で Backspace 長押し処理を関数化（例: `HandleReload(env *scenes.Env, assets assets.Cache, metricsLoader Loader)`）。`Game.updateGlobalToggles` からは挙動を呼ぶだけにして、副作用の所在を明確にする。
4. `internal/app/input/global.go` を設け、ヘルプ表示/リロードなどのショートカットをアクションマップとして宣言。`H` キーなどプラットフォーム固有の判定はここで完結させる。
5. `cmd/ui_sample/main.go` は将来的に `func main(){ app.Run(app.Config{...}) }` 程度に縮小し、Ebiten 依存を内包した `app.Run` がエラー処理まで面倒を見る。

### 今後の実装ステップの指針（抜粋）
1. `internal/game/app` を `internal/app/runtime` へ移動し、`Game` 構造体を `app.Runtime` として再公開する準備をする。
2. メトリクス適用ヘルパを新設し、初期化・ホットリロード双方で利用することで UI から `config/uimetrics` 依存を外す。
3. 入力周りを層ごとにモジュール化し、Ebiten から直接キーコードを引く箇所を `internal/app/input` に閉じる。
4. main から `internal/game/app` への直接 import を無くし、`internal/app` のファサード関数経由でのみ起動できるようにする。
