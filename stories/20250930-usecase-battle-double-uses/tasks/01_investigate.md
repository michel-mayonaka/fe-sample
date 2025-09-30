# 01 調査: 再現と原因切り分け

ステータス: [未着手]

## 目的
- 追撃時に攻撃側の Uses が2消費されない原因を特定する。

## 観測
- 失敗例: `attacker uses want 3 got 4`
- テスト: `internal/usecase/battle_test.go:114`（`TestRunBattleRound_AttackerDouble_ConsumesTwo`）

## 切り分けポイント
- `pkg/game.DoubleAdvantage(a,d)` が期待通り true になっているか。
- `internal/adapter.UIToGame` の武器重量 `Wt` が正しいか（Name/ID 突合）。
- `atkCount` の加算タイミングと追撃分岐の到達可否（HP条件/反撃条件）
- RNG により初撃/反撃の命中で分岐が変わる影響の有無。

## DoD
- 原因仮説を 1 つに絞り、修正方針を本ファイル末尾に記載。

## 進捗ログ
- 2025-09-30: 作成。

