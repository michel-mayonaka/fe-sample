# 02_design_options — 設計案の比較

## 案A: 現状維持（`service.Input` 継続）
- Pros: 既存安定、変更最小。
- Cons: UI補助パッケージの分割方針と不整合、Scenes が `service` 直依存。

## 案B: `internal/game/ui/input` へ移行
- 内容: `Action` と `Reader`（`Press/Down`）を `ui/input` に定義。`Snapshot` 等の入力ポーリングはアプリ側に残す。
- Pros: UI レイヤに責務集約、Scenes の参照が明確。
- Cons: 当面2箇所にコードが分かれる（`service`/`ui`）ためアダプタが必要。

## 案C: `pkg/game/input`（純粋ロジック）+ `internal/game/ui/input`（UI公開）
- 内容: マッピング/状態遷移は `pkg/game/input` に切り出し、Ebiten依存を排除。`ui/input` はUI向けの型公開と適合層。
- Pros: テスト容易・将来の別実装（ゲームパッド等）に強い。
- Cons: 現段階では分割がやや過剰。移行コスト増。

## 推奨（初期）
- Bを第一歩とし、`service.Input`→`ui/input.Reader` のアダプタで橋渡し。将来必要になればCへ拡張。

