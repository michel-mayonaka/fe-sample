# 02 修正: 追撃時の耐久消費の是正

ステータス: [未着手]

## 目的
- 追撃が成立した場合に攻撃側の Uses を正しく2消費させる。

## 変更候補
- A) `RunBattleRound` の追撃分岐到達条件の見直し（`gd2.S.HP > 0` のみ/`ga2` 死亡時の扱い）。
- B) `atkCount`/`defCount` 加算箇所の一元化（`ResolveRoundAt` 呼び出し直後に加算）。
- C) RNG 依存を避けるためテストで seed 固定。

## 影響
- UI 表示とインベントリ保存の整合が取れること。

## DoD
- `TestRunBattleRound_AttackerDouble_ConsumesTwo` が安定成功。

## 進捗ログ
- 2025-09-30: 作成。

