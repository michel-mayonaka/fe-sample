# 現況レポート

対象: `pkg/*`, `internal/*`, `cmd/ui_sample`, `docs/*`（2025-09-30 時点）

## コンポーネントと依存
- `pkg/game`: ドメイン（戦闘解決/予測/速度/入力）。テスト充実。UI 非依存。
- `internal/usecase`: アプリケーションサービス（`App`）。Port 実装: Reload/Persist/Equip 等。`data.TableProvider` も実装。
- `internal/repo`: JSON バックエンド（User/Weapons/Inventory）。将来 SQLite あり（`repo/sqlite`）。
- `internal/game`: UI 層（Scenes/UI補助/Provider/Runner）。
  - `app`: ebiten.Game ランナー、Env セットアップ、Provider 注入。
  - `scenes`: 画面群（list/status/inventory/sim 等）。`Env` と `Session` 明確化済み。
  - `service/ui`: UI メトリクス/フォント/画像ロード/描画補助（`widgets`）。
  - `ui/{layout,draw,view,adapter,input}`: 責務分割済み。`ui/input` はドメイン `pkg/game/input` の薄い公開面。
  - `provider/input/ebiten`: 物理入力→ドメイン入力への変換（`Source.Poll()`）。
  - `data`: UI 参照用 Provider（テーブル/ユーザ在庫/ユニット変換）。
  - `util`: 現時点で空ディレクトリ（命名規約上は非推奨名）。
- `internal/model{,/user}`: マスタ/ユーザモデル（純粋型）。
- `internal/infra/userfs`: JSON I/O。
- `cmd/ui_sample`: エントリ。
- `docs`: アーキ/命名/DB/オフライン等の方針が整理されている。

## 境界と責務（観測）
- UI→Usecase の更新は Port 経由で実施（シーンは Repo に直アクセスしない方針）。
- 参照は `internal/game/data.Provider()` に集約。Usecase.App が TableProvider を実装し DI。
- 入力は `pkg/game/input` と Ebiten アダプタで抽象化。`internal/game/ui/input` は薄い再公開。
- `Env`（Port 参照/メタ）と `Session`（UI一時状態）の分離は実装済み。
- Scenes は ebiten/ebitenutil を直接 import（描画責務のため想定範囲内）。

## データフロー（代表）
1) 起動: `app.NewUIAppGame()` で Repo 構築→Usecase.New→`gdata.SetProvider(app)`→Env 構築。
2) 入力: Ebiten→`provider/input/ebiten.Source`→`pkg/game/input`→`ui/input.Reader`→Scenes。
3) 参照: Scenes→`gdata.Provider()`→テーブル/ユーザ在庫取得（必要に応じ UI ユニット変換）。
4) 更新: Scenes→Port（`InventoryPort` 等）→Usecase→Repo→保存。
5) UI メトリクス: 起動時/再読込で `config/uimetrics` を `service/ui` に適用。

## 気づき（要検討ポイント）
- `internal/game/data.TableProvider` が `uicore.Unit` を返す API を含み、UI 型に依存している。
  - 代替案: Provider は純データ（UserTable/UserSnapshot のみ）を返し、UI 型変換は `service/ui` 側へ寄せる。
- `internal/usecase.ItemsTable()` が都度 JSON を直読み。Repo/キャッシュ経由に寄せる余地。
- `internal/game/util` ディレクトリが存在（空）し、命名規約「util 禁止」に抵触。用途特化へ改称/削除が望ましい。
