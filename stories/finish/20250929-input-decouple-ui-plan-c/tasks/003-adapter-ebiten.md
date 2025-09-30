# 003 ebiten アダプタ実装

## 目的
- ebiten の入力取得 API をアダプタ層に閉じ込め、`InputSource` 実装を提供する。

## スコープ
- パッケージ: `internal/game/provider/input/ebiten`。
- 実装: `type EbitenInputSource struct { ... }`（`Poll`/`Events`）。
- キー/パッド/マウス → `ControlState`/`InputEvent` 変換。

## 成果物
- 新規コード一式と最小テスト（可能なら `ebiteninput` をインタフェース境界でモック）。
- ワイヤリング箇所（UI サンプルの初期化）での置換。

## 受け入れ基準
- [x] `cmd/ui_sample` がビルド成功。
- [ ] 主要キー（方向/決定/キャンセル）が従来と同じ動作に見える。
- [x] `internal/game/...` から `ebiten` 直接依存が Scene/Usecase から消える。
