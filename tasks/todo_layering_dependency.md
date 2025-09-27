# TODO — レイヤリング/依存 (2025-09-27)

- `main` の入力処理・状態遷移を `internal/app` に追加委譲（UIは描画＋イベント通知のみ）。
- UI側のフォールバックI/O撤去（`screens/battle.go: forecastBoth` の直接JSON読込を削除し注入テーブルのみ使用）。
- `config` にパス等の直書きを統一（`Update()` 内のユーザJSON）。

