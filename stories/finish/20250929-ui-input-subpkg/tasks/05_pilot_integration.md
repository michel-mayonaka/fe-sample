# 05_pilot_integration — 代表1画面での適用

## 対象候補
- `internal/game/scenes/character_list` または `status`。

## 手順
- import を `gamesvc` → `ui/input` に置換。
- `Press/Down` 呼び出しはそのまま、`Action` 参照先のみ差し替え。
- 動作確認（Confirm/Cancel/トグル系）。

## 検証
- `make mcp` が通ること。
- 目視での入力反応確認。

