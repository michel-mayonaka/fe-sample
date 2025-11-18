# 03_移行ステップ案と後続タスクの整理

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 00:22:56 +0900

## 目的
- main→app 移譲の実装方針とステップを整理し、別ストーリー/Backlog として切り出す準備を整える。

## 完了条件（DoD）
- [x] レイヤリング改善を行う場合のフェーズ分け（例: 入力処理移譲→状態遷移移譲→I/O整理）が整理されている。
- [x] 各フェーズで触れる主なファイル/パッケージが列挙されている。
- [x] 必要に応じて、新しい実装用ストーリーまたは Backlog エントリが作成され、関連リンクが記載されている。

## 作業手順（概略）
- `02_レイヤリング方針の整理` の結果をもとに、実装ステップをフェーズごとに分解する。
- 影響範囲やリスクを考慮しつつ、1ストーリー/PR あたりの妥当な作業量を検討する。
- Backlog/Stories/Discovery の状態を更新し、今後の作業が追跡しやすい形に整える。

## 進捗ログ
- 2025-11-19 00:22:56 +0900: タスク作成。
- 2025-11-19 02:06:30 +0900: Task 01/02 の成果をもとに移行ステップ案の整理を開始。
- 2025-11-19 02:10:45 +0900: フェーズ分割と対象ファイルを整理し、Backlog へ実装ストーリー候補を追記。
- 2025-11-19 02:12:30 +0900: DoD を満たしたため完了。

## 依存／ブロッカー
- `02_レイヤリング方針の整理` の完了。

## 成果物リンク
- 新規ストーリー/Backlog エントリ、移行ステップメモ

## 移行ステップメモ

### フェーズ分割（概観）
1. **Phase 1 — App ファサードと Bootstrap の分離**
   - 目的: `cmd/ui_sample` から `internal/game/app` への直接依存を断ち、`internal/app`（仮: `internal/app/runtime`）にブートストラップ用の公開 API を設ける。
   - 主な変更ファイル: `cmd/ui_sample/main.go`, `internal/game/app/core.go`（新パッケージへ移動）, `internal/app/bootstrap/*.go`, 新規 `internal/app/config/config.go`.
   - ゴール: `main.go` が `app.NewRuntime(cfg)` の呼び出しのみで完結し、乱数・レイアウト・Repo 初期化などは app 側に閉じる。
   - リスク/備考: 既存テスト（`internal/game/app/*_test.go`）の import path 更新、`//go:build !headless` タグの移行漏れに注意。

2. **Phase 2 — 入力レイヤとランタイム制御の再配置**
   - 目的: 物理入力（Ebiten）→抽象入力→Scene Intent の各レイヤを `internal/app/input` に集約し、`Game.updateGlobalToggles` で直接 Ebiten キーを参照しない。
   - 主な変更ファイル: `internal/game/app/game.go`, `internal/game/provider/input/ebiten`, 新規 `internal/app/input/{layout,source,global}.go`, `internal/game/ui/input`.
   - ゴール: `Game` は `app.InputController` から提供される抽象イベントのみを扱い、グローバルイベント（ヘルプ/リロード）は app 層で発火を検知して Scene へ通知する。
   - リスク/備考: Backlog にある「入力レイアウト設定の外部化」と統合可能。設定ファイル導入時のパース失敗フォールバックを先に決める。

3. **Phase 3 — メトリクス/ホットリロードの統一**
   - 目的: 初期化と Backspace 長押しの双方で行っている `uicore.Metrics` への手動コピーを `internal/app/metrics` のヘルパにまとめ、副作用を app 層の `ReloadService` 経由にする。
   - 主な変更ファイル: `internal/game/app/core.go`, `internal/game/app/game.go`, 新規 `internal/app/metrics/metrics.go`, `internal/app/reload/service.go`, `internal/assets`.
   - ゴール: `Game` からは `reload.Trigger(env)` のような API を呼ぶだけで、Repo Reload/Asset Cache Clear/Unit 再構築/メトリクス再適用が一貫して行われる。
   - リスク/備考: `assets.Clear()` を他所からも呼んでいないか確認し、副作用順序（Repo→Asset→Session）をドキュメント化する。

4. **Phase 4 — Env/Session と設定境界の整理**
   - 目的: `Env` が抱えている `UserPath` や `RNG` を `internal/app` が保持し、Scene には必要な値だけを渡す。初期 Scene の指定もコンフィグ駆動にする。
   - 主な変更ファイル: `internal/game/scenes/env.go`, `internal/game/app/core.go`, 新規 `internal/app/options.go`, `docs/ARCHITECTURE.md`.
   - ゴール: Scene 側が app 固有情報に直接依存しない構造を作り、将来的な CLI/テストランナーでも `internal/app` を再利用可能にする。
   - リスク/備考: `scenes.Env` を変更する場合は関連シーンのフィールドアクセスを一括アップデートする必要あり。段階移行のために `EnvMeta` などのサブ構造体追加を検討。

5. **Phase 5 — 仕上げ（ドキュメントとテレメトリ）**
   - 目的: `docs/ARCHITECTURE.md`, `docs/CODEX_CLOUD.md` 等に新しいレイヤリングを反映し、`make mcp` を通して回帰を確認。
   - 主な変更ファイル: docs 系、`Makefile`（必要なら `MCP_STRICT` の扱い明記）。
   - リスク/備考: ここで Backlog との突合を行い、実装完了後に該当エントリをクローズ。

### 追加で想定される Backlog/ストーリー
1. **`2025-11-19: internal/app bootstrap ファサード実装`**
   - Phase 1 をカバー。`cmd/ui_sample` を `app.Run(cfg)` に差し替えるストーリー。
   - Backlog エントリとして追記（`stories/BACKLOG.md` 参照）。

2. **`2025-11-19: グローバル入力とホットリロードの app 層統合`**
   - Phase 2〜3 をカバー。入力とリロードの責務整理を段階適用するストーリー。
   - Backlog に追記済み。

3. 既存の Backlog 「入力レイアウト設定の外部化（2025-09-30）」と Phase 2 を統合し、設定ファイル導入を同フェーズで扱う。

4. Phase 4 以降で必要になりそうなサブタスク（`Env` 再設計、ドキュメント更新）は、実装フェーズ開始時に Discovery として切り出す。
