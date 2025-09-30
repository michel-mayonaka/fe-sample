# 001 設計固め（案Cの具体化と命名）

## 目的
- 入力ロジックの UI 非依存化（案C）を設計として確定し、命名・配置を合意する。

## スコープ
- 層の責務と依存関係の定義: `pkg/game/input`（ドメイン）と `internal/game/provider/input/ebiten`（アダプタ）。
- 公開 I/F（例: `InputSource`, `InputEvent`, `ControlState`）の最小集合と拡張方針。
- キー/パッド/マウスのマッピング方針（テーブル駆動 vs コード）。

## 成果物
- 設計メモ（本ファイルの更新）。
- ドキュメント更新案（`docs/ARCHITECTURE.md` 差分案）。

## 受け入れ基準
- [x] 依存方向が UI → アダプタ → ドメイン の一方向であることを図で説明。
- [x] 主要型/関数の命名とパッケージ配置が `docs/NAMING.md` に整合。
- [x] 拡張（新入力デバイス追加）の手順が 3 ステップで説明可能。

## 作業メモ
### レイヤ構成（決定）
```
UI（Scenes, App）
  │  依存（意図のみ参照: Press/Down 等）
  ▼
Adapter（internal/game/provider/input/ebiten）
  │  依存（取得のみ・変換実装）
  ▼
Domain（pkg/game/input）
```

依存原則:
- UI は `pkg/game/input` の公開 I/F（`Reader`/`Action` 等）のみ参照。
- 物理入力（Ebiten 由来）は Adapter が変換して Domain に投影。
- Domain は UI/API 実装（Ebiten 等）を import しない。

### 公開 I/F（Domain: `pkg/game/input`）
- `type Action int`
  - 列挙は `ActionUnknown` を 0（命名規約: 型接頭辞 + Unknown を 0）
  - 既存 UI の操作に対応: `ActionUp/Down/Left/Right/Confirm/Cancel/Menu/.../ActionTerrainDef3` まで。
- `type EventKind int` と `type Event struct { Kind EventKind; Code int; Value float64; Mods Modifier }`
  - `EventKind`: `EventUnknown`, `EventKey`, `EventMouseButton`, `EventMouseWheel`, `EventGamepadButton`, `EventGamepadAxis`。
  - `Modifier` はビットフラグ（`ModShift/ModCtrl/ModAlt/ModMeta`）。
- `type ControlState struct { Up, Down, Left, Right, Confirm, Cancel, ... bool }`
  - UI が即時参照する“意味”の集合（フレームスナップショット）。
- `type Source interface { Poll() ControlState; Events() []Event }`
  - アダプタ実装が満たす。`Poll` は現在状態の投影、`Events` は直近フレームの素イベント列（任意）。
- `type Reader interface { Press(Action) bool; Down(Action) bool }`
  - 辺検出（prev/curr 比較）。実装例: `EdgeReader`（`Step(ControlState)` で更新）。

備考:
- 現行 `internal/game/ui/input.Reader` は段階移行のため `pkg/game/input.Reader` へ別名で委譲するシムを用意し、最終的に置換する。

### アダプタ（`internal/game/provider/input/ebiten`）
- 役割: Ebiten の入力 API を読み取り、ドメイン `ControlState` と `Event` に変換。
- 既定マッピング（最小）:
  - `↑/↓/←/→` → `ActionUp/Down/Left/Right`
  - `Z/Enter` → `ActionConfirm`、`X/Escape` → `ActionCancel`
  - `Tab` → `ActionMenu`、`A/S` → `ActionPrev/Next`
  - 左クリック → `ActionConfirm`
  - `1/2/3` と `Shift+1/2/3` → `ActionTerrainAtt*/ActionTerrainDef*`
- 取得はテーブル駆動（`KeyCode → Action`）。修飾キーは `Modifier` を併用。

### 拡張方針
- マッピングは設定主導にする（JSON/Go テーブル）。UI テストではテーブル差し替えで検証。
- デバイス追加（Gamepad 等）は、`Source` 実装を追加して `Event` 種別で拡張可能。
- `EdgeReader` により `Press/Down` の判定は UI から独立してテスト可能。

### 新デバイス追加の手順（3 ステップ）
1) Domain に `EventKind` の値を追加（例: `EventTouch`）。必要なら `Modifier` も拡張。
2) Adapter に当該デバイス用のスキャン/マッピング処理を追加し、`Poll/Events` で投影。
3) テスト: マッピングテーブルをモックし、`EdgeReader` の辺検出と併せてユニットテストを追加。

### 影響範囲（現状把握）
- `internal/game/service/input.go` が Ebiten へ直接依存（優先移行対象）。
- `internal/game/ui/input/*` は `service.Input` へのアダプタ（段階的に Domain 直参照へ）。
- `internal/game/app/game.go` の一部（グローバルトグル）で Ebiten のキー直接参照あり（UI 層のため許容）。
- `internal/game/ui/*`（draw/view/adapter）は描画系で `*ebiten.Image` を参照（本ストーリーの対象外）。

### Docs 差分案
- `docs/ARCHITECTURE.md`: 「UI 補助サブパッケージ」の入力節を `pkg/game/input`（Domain）＋ Adapter に再記述。
- `docs/API.md`: 公開 I/F（`Action/Event/ControlState/Source/Reader`）を追記。列挙は `ActionUnknown` を 0 に据える方針を明記。
- `docs/NAMING.md`: 列挙・パッケージ命名の整合を確認（型接頭辞・Unknown=0）。

### マイグレーション計画（段階適用）
M1: Domain 型と Adapter の土台を追加（互換レイヤを維持）。
M2: `internal/game/app` を新 `Source+EdgeReader` に切替、`internal/game/ui/input` をシム化。
M3: 旧 `service.Input` を撤去し、テストを `pkg/game/input` へ移設（回帰確認）。
