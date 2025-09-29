# [task] サービス系（遷移/演出/状態連携）の移設

- ステータス: 完了
- 目的: 画面遷移・アニメーション・状態連携などの横断サービスを `internal/game/service/*` に移し、`scenes` から分離する。

## 入力
- タスク01の分類表（サービス系に分類された項目）
- 既存 `internal/game/service/*` の構成

## スコープ
- 新設: `internal/game/service/levelup`（成長抽選/反映）。
- 既存: `scenes/common/popup/levelup.go` のロジック部を移設。
- 参照更新（`status` シーンから `levelup.Roll/Apply` を利用）。

## 非スコープ
- 新規演出の追加。

## 手順
- 既存サービス構成を踏襲し、責務単位でサブパッケージ切り出し。
- 呼び出しの依存方向（`scenes -> service` の一方向）を維持。
- `make mcp` で逐次確認。

## DoD（完了条件）
- 成長抽選/反映ロジックが `service/levelup` に移り、描画は `ui/draw` に分離されている。
- 依存方向の逆転がない（`service` が `scenes` を参照しない）。
- `make mcp` グリーン。

## コマンド例
- `rg -n "transition|fade|anim|ease|state" internal/game/scenes || true`
- `make mcp`
