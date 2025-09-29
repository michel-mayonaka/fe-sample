# [story] `ui/input` サブパッケージ導入の検討

- ステータス: 提案中
- 日付: 2025-09-29
- 参照: docs/ARCHITECTURE.md, docs/NAMING.md, stories/BACKLOG.md

## 背景
現状、抽象入力は `internal/game/service/input.go` にあり、Scenes は `gamesvc.Action` と `(*Input).Press/Down` を直接参照している。UI 補助の責務分割（`ui/layout|draw|view|adapter`）に揃えるなら、入力→意図(Intent)の解釈も UI 側のサブパッケージとして明示し、依存方向と命名を整えたい。

## 目的
- 「生の入力（Ebiten依存）」と「抽象アクション/インテント解釈」を分離し、UI 層で再利用可能にする。
- Scenes が参照する型・関数を `ui/input` に集約し、探索性を向上。
- 将来的なゲームパッド/タッチ等の拡張、リマップ機能の受け皿を用意。

## スコープ
- 設計案の比較と方針決定（維持/移行/分割）。
- `internal/game/ui/input` の最小API定義（`Action` と `Reader` インターフェース想定）。
- 既存 `service.Input` を活かすアダプタ方針の提示（段階移行）。
- 代表1画面での試験導入（Scenes 側の import を `ui/input` に切替）。

## 非スコープ
- 全Sceneの一括移行（別ストーリーで段階適用）。
- フル機能のキーバインドUI/永続化（将来検討）。

## 成果物 / DoD
- 設計決定のドキュメント（Pros/Cons、採用/見送りの根拠）をこのREADMEに記録。
- `internal/game/ui/input` ディレクトリに最小ファイルを追加（型定義/GoDoc）。
- 代表1画面（例: `character_list`）で `ui/input.Reader` を使用してビルド成功。
- `make mcp` グリーン、lint 0 を目標。

## 工程（サブタスク）
- [ ] tasks/01_discovery_current_usage.md — 現状の入力参照箇所の棚卸し
- [ ] tasks/02_design_options.md — 案比較（A: 現状維持、B: ui/input へ移行、C: `pkg/game/input` + アダプタ）
- [ ] tasks/03_define_min_api.md — `Action`/`Reader` の最小API定義と命名確定
- [ ] tasks/04_spike_adapter.md — 既存 `service.Input` → `ui/input.Reader` アダプタのスパイク
- [ ] tasks/05_pilot_integration.md — 代表1画面での切替と検証
- [ ] tasks/06_docs_update.md — ARCHITECTURE/API の更新と決定事項の反映

## 影響範囲
- `internal/game/service/input.go`（参照側のimport変更/アダプタ追加）
- Scenes: `internal/game/scenes/...` の入力参照

## リスクと対策
- リスク: 命名/依存のねじれによる参照断裂
  - 対策: 段階移行（型は互換インターフェースで橋渡し、まずは1画面）。
- リスク: Ebiten API への暗黙依存が残る
  - 対策: `ui/input` からはEbiten型を直接参照しない。

## 計画（目安）
- 見積: 0.5–1日（スパイク+代表1画面）
- マイルストン: M1 設計合意 → M2 スパイク → M3 パイロット適用

## 進捗・決定事項（ログ）
- 2025-09-29: バックログから昇格し、方針検討を開始。
- 2025-09-29: `internal/game/ui/input` を追加し、最小API(Action/Reader)を定義。
- 2025-09-29: `character_list`/`status`/`inventory`/`sim` を `ui/input` 参照に切替。
- 2025-09-29: `game.Ctx.Input` を `ui/input.Reader` に変更（実装は従来の `service.Input` を供給）。
- 2025-09-29: `ServiceAdapter`（`service.Input` を包む）を追加し、単体テストを作成・通過。
- 2025-09-29: docs(API/ARCHITECTURE) 追記を反映。
