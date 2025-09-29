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
- [ ] 依存方向が UI → アダプタ → ドメイン の一方向であることを図で説明。
- [ ] 主要型/関数の命名とパッケージ配置が `docs/NAMING.md` に整合。
- [ ] 拡張（新入力デバイス追加）の手順が 3 ステップで説明可能。

## 作業メモ
- ドメイン: `pkg/game/input`
  - `type InputEvent struct { Kind, Code, Value, Mods }`
  - `type ControlState struct { Up,Down,Left,Right,Confirm,Cancel,... }`
  - `type InputSource interface { Poll() ControlState; Events() []InputEvent }`
- アダプタ: `internal/game/provider/input/ebiten`
  - ebiten API から状態を取得し、`ControlState` を構築。
