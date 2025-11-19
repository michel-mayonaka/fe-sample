# 20251119-app-input-reload-integration — 入力層とホットリロードの app 層統合

ステータス: [準備中]
担当: @tkg-engineer
開始: 2025-11-19 02:06:54 +0900

## 目的・背景
- `internal/game/app/game.go` が Ebiten の生キー判定やメトリクス再読み込みを直接扱っており、抽象入力レイヤを経由しない経路が残っている。
- Backspace 長押し時のホットリロードやヘルプ表示などのグローバル操作を `internal/app` 層へ集約し、UI ロジックがプラットフォーム依存コードを持たないようにする（stories/20251119-layering-main-to-app の Phase2〜3）。

## スコープ（成果）
- `internal/app/input` に物理入力→抽象入力の変換、グローバルアクション検知、設定ファイルレイアウト適用を集約する。
- `internal/app/reload`（仮）を追加し、メトリクス適用・Repo Reload・Asset クリア・ユニット再構築の手順を共通化する。
- `Game.updateGlobalToggles` から Ebiten 依存と重複ロジックを排除し、app 層イベント経由の制御に差し替える。

## 受け入れ基準（Definition of Done）
- [ ] `internal/game/app/game.go` で直接 `ebiten.IsKeyPressed` を呼んでいない（すべて `internal/app/input` 経由）。
- [ ] 初期化時とリロード時のメトリクス適用が単一のヘルパを通り、副作用順序がドキュメント化されている。
- [ ] Backspace リロード／ヘルプトグル等の動作確認ログがあり、`make mcp` がグリーン。

## 工程（サブタスク）
- [ ] 設計 — 入力層/グローバル操作方針整理 — `stories/20251119-app-input-reload-integration/tasks/01_design-input-and-reload.md`
- [ ] 実装 — input/reload モジュール追加と `Game` 差し替え — `stories/20251119-app-input-reload-integration/tasks/02_implement-input-modules.md`
- [ ] 統合 — ReloadService/メトリクス適用一本化 — `stories/20251119-app-input-reload-integration/tasks/03_unify-reload-service.md`

## 計画（目安）
- 見積: 2〜3 セッション（設計+実装+統合）
- マイルストン: M1 方針合意 → M2 入力モジュール実装 → M3 ReloadService 完了

## 進捗・決定事項（ログ）
- 2025-11-19 02:06:54 +0900: ストーリー作成（Backlog: グローバル入力とホットリロード統合 から昇格）

## リスク・懸念
- Ebiten 依存を抽象化する際に既存の headless テスト構成が壊れる可能性。
- 入力レイアウト設定の外部化（別 Backlog）と衝突するリスクがあるため、計画段階で統合方針を決める必要がある。

## 関連
- PR: #
- Issue: #
- Docs: `docs/architecture/README.md`, `docs/CONFIG.md` (追加予定)
