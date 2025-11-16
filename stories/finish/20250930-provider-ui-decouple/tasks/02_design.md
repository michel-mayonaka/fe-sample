# 02_design — 方針確定（Port/Adapter）

## 目的
- Provider を純参照Portとし、UI 型は Adapter に集約する設計ルールを明文化する。

## 方針
- Provider IF は `model/*` と `user/*` のみ返す。
- UI 型 `uicore.*` は `internal/game/ui/adapter` が生成する。
- 変換関数はユニット/装備ごとに最小APIへ分解（テスト可能な純関数）。

## 成果物
- ルール追記: `docs/ARCHITECTURE.md` / `docs/API.md`。（済）
- 置換計画の簡易表（対象/対応/影響）。
