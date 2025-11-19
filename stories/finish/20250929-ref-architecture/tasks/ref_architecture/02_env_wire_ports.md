# 02_env_wire_ports — EnvにPort参照を追加

## 目的・背景
- Envが最小の“環境コンテナ”としてPort参照を保持できるようにし、以降の段階移行先を用意。
- 既存 `UseCases` と並存配線（ブリッジ）で互換を保つ。

## 作業項目（変更点）
- `internal/game/scenes/env.go` に `Data DataPort`, `Battle BattlePort`, `Inv InventoryPort` を追加。
- `internal/game/app/core.go` の `NewUIAppGame()` でポート注入（usecase.App を各Portへアサイン）。
- 既存 `Env.App`（合成UseCases）は維持。呼び出し差し替えは次工程。

## 完了条件
- ビルド成功。起動確認（UI表示可能）。
- `Env` に Port フィールドが存在し、nilでないこと（注入済み）。

## 影響範囲
- `scenes/env.go` と `game/app/core.go` のみ（配線）。

## 手順
1) `env.go` に Portフィールドを追加（コメントで役割を明記）。
2) `core.go` で usecase.App を `Data/Battle/Inv` へ注入。
3) `make mcp` 実行→起動（`go run ./cmd/ui_sample`）してクラッシュしないこと。

## 品質ゲート
- `make mcp`

## 動作確認
- 起動し、ユニット一覧が表示される。

## ロールバック
- 追加フィールド/注入コードを戻す。

## Progress / Notes
- 2025-09-28: 着手
- 2025-09-28: 配線・起動確認（`Env.Data/Battle/Inv` に `usecase.App` を注入、`make mcp` 成功）

## 関連
- `docs/architecture/README.md` 4章/5章/12.4
