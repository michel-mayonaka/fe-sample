# 004 Scene から UI 依存排除

## 目的
- Scene/Usecase 層が `ebiten` を直接 import せず、`InputSource` による依存逆転を達成する。

## スコープ
- 対象: `internal/game/scene/...`（存在する場合）と関連 Usecase。
- 入力参照箇所を `InputSource` / `ControlState` 参照に置換。
- 命名/可視性の調整（export は最小限）。

## 成果物
- リファクタ済みコード（diff 小さめを心がける）。
- コンパイル/簡易動作確認。

## 受け入れ基準
- [x] 入力取得での `ebiten` 直接呼び出し（`CursorPosition` 等）が Scene/Usecase から 0 件。
- [x] `make check-all` が成功。
