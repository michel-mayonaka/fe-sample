# 01_discovery_current_usage — 現状の入力参照の棚卸し

## 目的
Scenes/アプリ層での入力参照箇所を洗い出し、`ui/input` への切替影響を見積もる。

## やること
- `rg` で `ctx.Input` / `gamesvc.<Action>` / `ebiten.IsKeyPressed` 直接参照を抽出。
- 代表画面（`character_list` ほか）での利用パターン（Press/Down、マウス合流）を確認。
- 特殊操作（地形切替 1/2/3 + Shift、マウスLeft→Confirm 合流）の扱いを整理。

## 期待アウトプット
- 参照一覧と簡易分類（Press/Down/連打/長押し）。
- 置換難易度の評価（低/中/高）。

