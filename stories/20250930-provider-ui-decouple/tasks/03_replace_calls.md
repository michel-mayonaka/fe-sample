# 03_replace_calls — 呼び出しの再配線

## 目的
- UI 型依存の呼び出しを Adapter 経由へ切替し、Provider の純化を完了する。

## 対象候補
- `internal/game/scenes/*` でのユニット/装備表示生成箇所。
- 旧来の `UserUnitByID` 相当を期待していたコード（存在すれば）。

## 手順
- [x] Provider 参照箇所を確認し、UI 型生成は Adapter へ移動。
- [x] `rg -n "uicore\.|UserUnitByID"` で再チェック。
- [ ] 表示/シミュ/ステータスで動作確認。
