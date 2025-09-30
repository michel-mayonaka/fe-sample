# 01_inventory — 現状調査

## 目的
- Provider と UI 型の結合点を棚卸しし、置換範囲を明確化する。

## 調査観点
- `internal/game/data/` での `uicore.` 参照有無。
- Scenes/UI 側で Provider が返す値に UI 変換を期待していないか。
- Adapter 経由のユニット生成/装備表示の責務分担。

## コマンド例
- `rg -n "uicore\." internal/game/data`
- `rg -n "Provider\(\)" internal/game | sed -n '1,200p'`

## 完了条件
- [ ] 現状の結合点リスト化。
- [ ] 置換の影響範囲見積り。

