# [task] UI 系（描画/入力）の移設

- ステータス: 未着手
- 目的: UI 補助ロジック（描画・入力関連）が `scenes` に混在していれば UI 層へ移す。

## 入力
- タスク01の分類表（UI系に分類された項目）
- 規約: `docs/NAMING.md`（役割+領域で命名、`helper` 禁止）

## スコープ
- 新設候補: `internal/game/ui/draw`, `internal/game/ui/input`
- UI系関数/型の移動、必要に応じて公開/非公開の見直し。
- 参照更新（`scenes/*`）。

## 非スコープ
- UI デザイン変更や挙動変更。

## 手順
- パッケージ骨組み作成（`doc.go` で責務を明文化）。
- 関数の移動（ビルドが通る最小単位でコミット）。
- 呼び出しサイトの import/識別子を置換。
- `make mcp` で逐次確認。

## DoD（完了条件）
- UI系の関数が `internal/game/ui/*` に集約され、`scenes` 側からの依存が解消している。
- `make mcp` グリーン。

## コマンド例
- `mkdir -p internal/game/ui/{draw,input}`
- `rg -n "draw|layout|hover|cursor|input" internal/game/scenes || true`
- `make mcp`

