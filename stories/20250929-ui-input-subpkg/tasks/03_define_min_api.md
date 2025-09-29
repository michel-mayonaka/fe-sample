# 03_define_min_api — 最小APIの定義

## 方針
- `internal/game/ui/input` に以下のみ定義：
  - `type Action int` と列挙（既存 `service.Action` と同等の並び）
  - `type Reader interface { Press(Action) bool; Down(Action) bool }`
- Ebiten 依存を持ち込まない。`Snapshot()` はアプリ側（`app.Game.Update`）が継続実行。

## 作業
- `internal/game/ui/input/types.go` を追加（GoDoc 1行必須）。
- 既存 `service.Input` が `Reader` を満たすことを確認（メソッド互換）。
- 命名は docs/NAMING.md に従う（公開: UpperCamel）。

