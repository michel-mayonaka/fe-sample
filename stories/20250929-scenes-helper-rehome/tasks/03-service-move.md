# [task] サービス系（遷移/演出/状態連携）の移設

- ステータス: 未着手
- 目的: 画面遷移・アニメーション・状態連携などの横断サービスを `internal/game/service/*` に移し、`scenes` から分離する。

## 入力
- タスク01の分類表（サービス系に分類された項目）
- 既存 `internal/game/service/*` の構成

## スコープ
- 新設候補: `internal/game/service/transition`, `internal/game/service/easing` など（実体に合わせて決定）。
- 関数/型の移動と API 整形（必要に応じインタフェース化）。
- 参照更新（`scenes/*`）。

## 非スコープ
- 新規演出の追加。

## 手順
- 既存サービス構成を踏襲し、責務単位でサブパッケージ切り出し。
- 呼び出しの依存方向（`scenes -> service` の一方向）を維持。
- `make mcp` で逐次確認。

## DoD（完了条件）
- サービス系の共通処理が `internal/game/service/*` に移設されている。
- 依存方向の逆転がない（`service` が `scenes` を参照しない）。
- `make mcp` グリーン。

## コマンド例
- `rg -n "transition|fade|anim|ease|state" internal/game/scenes || true`
- `make mcp`

