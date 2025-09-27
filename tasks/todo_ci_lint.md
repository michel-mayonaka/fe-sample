# TODO — ビルド/CI・Lint (2025-09-27)

- `.golangci.yml` と Makefile を `internal/...` も対象に（`revive/gofumpt/unused` 等を有効化）。
- CI（GitHub Actions）で `make mcp` 実行とキャッシュ、GUI不可環境のスキップ条件整備。
- README に Go 1.25 系の明示と `toolchain`/代替手順を追記。

## 進捗（2025-09-28）
- Lint設定強化: `.golangci.yml` を更新し、`govet/staticcheck/gocritic/ineffassign/errcheck/misspell/unused` を有効化。対象を`./...`に拡張（UI配下も含む）。
- UI配下のLint有効化: 以前の除外を撤廃。`text.BoundString` を全廃し `font.MeasureString` ベースへ置換。`gocritic` 指摘（if-else連鎖）を修正。
- text/v2移行の下準備: `uicore.TextDraw`（`text/v2`ベースの薄いラッパ）を追加し、全描画呼び出しを段階置換。`nolint:staticcheck` をコードから撤去し、`.golangci.yml` の除外設定も削除済み。
- revive強化: `exported/package-comments/var-naming/error-naming/if-return/superfluous-else/early-return/unhandled-error` を有効化し、全体をグリーン化。
- Makefile更新: `lint` を全体対象に、`lint-ci`（厳格）を追加。`mcp` は CI 上で `lint-ci` を使うよう切替。
- CI導入: `.github/workflows/ci.yml` を追加（Go 1.25.x固定、モジュール/ビルドキャッシュ、`make mcp`）。GUI依存のビルドはスキップ条件維持。
- 検証: `golangci-lint run`=0件、`make check-all` ビルドOK、`cmd/ui_sample` ビルドOK。
- CIキャッシュ強化: Actions のキャッシュキーに `**/go.mod` を追加（既存の `**/go.sum` / `.golangci.yml` と併用）し、依存更新検知を安定化。
- 段階導入オプション: `workflow_dispatch` の `extra_linters` 入力を追加し、`EXTRA_LINTERS` 経由で厳格ルール（例: `funlen/gocognit/cyclop`）を手動トグル可能に（Makefile の `lint-ci` と連携）。
- PR1: `Update` のモード別抽出を実施。`updateListMode(mx,my)` と `updateStatusMode(mx,my)` を追加し、`Update` 内の該当処理を移設（挙動不変）。`make mcp`=OK（build/lint 0件）。

## 未了/保留
- text/v2 への移行（`text.Draw` → `text/v2` + `DrawOptions` での描画／色・座標・スケール指定）。
- revive 再有効化（`exported`/`package-comments` など）と `nolint` の段階的撤去。
- 追加ルール（複雑度/長さ）: `funlen/gocognit/cyclop` は段階導入。既定はOFF、CIでは `EXTRA_LINTERS` で随時ON可能。

## 次アクション（提案）
1) text/v2段階移行: 内部ヘルパ（`uicore/textutil` など）を用意し、`Draw(text,color,x,y)` の薄いラッパを先に導入→呼び出し置換→実装を `text/v2` に差し替え。
2) revive強化: コメントの追加を進め、`revive` を再有効化（まず `exported`/`package-comments`）。
3) 追加ルールの本格化: `EXTRA_LINTERS='-E funlen -E gocognit -E cyclop' make lint-ci` をPRで試し、閾値/除外の最小化→恒常ONへ昇格。

## 次セッション着手予定（PR計画）
- PR1: `Update` のモード別抽出（完了）
  - 受入: lint 0件、`make check-all` OK、一覧/ステータス操作の手動確認済み。
- PR2: `Draw` の分割（完了）
  - 追加: `drawList`, `drawStatus`, `drawSimBattle`, `drawInventory`（内部関数）。
  - 差替: `Draw` の `switch` を各 `draw*` 呼び出しに変更（重複描画コードは移設のみ）。
  - 受入: lint 0件、`make check-all` OK、全画面の表示確認。
- PR3: battle描画の小関数化（完了）
  - 分割: `drawBattleHeader`, `drawStartButton`, `drawForecastLeft`, `drawForecastRight`, `drawTerrainLabels` を追加。
  - 互換: `DrawBattleWithTerrain` の外部APIは不変。見た目/挙動も不変。
  - 受入: lint 0件、`make check-all` OK。
- PR4: status描画の小関数化（完了）
  - 分割: `drawStatusHeader`, `drawCoreStats`, `drawWeaponRanks`, `drawMagicRanks`, `drawEquipList` を追加。
  - 互換: `DrawStatus` の外部APIは不変。見た目/挙動も不変。
  - 受入: lint 0件、`make check-all` OK。
- PR2: `Draw` の分割（挙動不変）
  - 追加: `drawList`, `drawStatus`, `drawBattleSim`, `drawInventory`（内部関数）。
  - 差替: `Draw` の `switch` を各 `draw*` 呼び出しに変更（重複描画コードは移設のみ）。
  - 受入: lint 0件、`make check-all` OK、全画面の表示確認。
- PR3: battle描画の小関数化（`internal/ui/screens/battle.go`）
  - 分割: `drawBattleHeader`, `drawForecastLeft`, `drawForecastRight`, `drawTerrainLabels`, `drawStartButton`。
  - 受入: 見た目/挙動不変、lint 0件。
- PR4: status描画の小関数化（`internal/ui/screens/status.go`）
  - 分割: `drawStatusHeader`, `drawCoreStats`, `drawWeaponRanks`, `drawMagicRanks`, `drawEquipList`。
  - 受入: 見た目/挙動不変、lint 0件。

### 以降のフォローアップ（任意）
- repo: `internal/repo/inventory.go` の `Consume` を `isWeaponID/isItemID/consumeUses` で整理（cyclop<=10）。
- popup: `RollLevelUp` をテーブル駆動ループに（cyclop<=10）。
- util: `UnitFromUser` を小関数化し、可能なら adapter 層へ移譲（gocognit低減）。
- CI: 強化ルール（funlen/gocognit/cyclop）を `EXTRA_LINTERS` で段階的にON→閾値を現実値まで引下げ→恒常ON。

### チェックリスト（各PR共通）
- 挙動不変（一覧選択/ステータス/装備/模擬戦の操作手順が変わらない）。
- `golangci-lint run`: 0件。
- `make check-all`: OK（`cmd/ui_sample` ビルド含む）。

---
補足: 「ボイラープレート」と「DrawOptions設計」について
- ボイラープレート: API移行時に各所で毎回書く初期化コード（Faceの用意、座標/色指定、折返し処理など）のこと。重複を避けるため共通ヘルパを作る、という意味。
- DrawOptions設計: text/v2 は `DrawOptions`（位置・色・スケール等の描画オプション）を使うため、
  - 既定（ベースライン基準/左上開始/色/スケール1.0）
  - 行間・折返し・中央寄せなどの拡張
  - UI全体で同じ見た目になるプリセット（タイトル/本文/小サイズ）
  を決め、ヘルパ関数に閉じ込めて呼び出し側をシンプルに保つ方針を指します。

## 進捗詳細（2025-09-28 追記）
- 実装サマリ:
  - PR1 完了: `Update` のモード別抽出（`updateListMode`/`updateStatusMode`）。
  - PR2 完了: `Draw` の分割（`drawList`/`drawStatus`/`drawSimBattle`/`drawInventory`）。
  - PR3 完了: battle 描画の小関数化（`drawBattleHeader`/`drawStartButton`/`drawForecastLeft`/`drawForecastRight`/`drawTerrainLabels`）。
  - PR4 完了: status 描画の小関数化（`drawStatusHeader`/`drawCoreStats`/`drawWeaponRanks`/`drawMagicRanks`/`drawEquipList`）。
- 不具合修正:
  - 症状: `make run` で「ユニット一覧」が表示されない。
  - 原因: PR2 時に `drawList` 内で `ui.DrawCharacterList(...)` の呼び出し漏れ。`drawStatus` でも `ui.DrawStatus(...)` 未呼出。
  - 対応: 両関数に本体描画呼び出しを追加。回帰確認済み。
- 品質確認:
  - `make mcp`: すべて成功（`vet/build/lint` グリーン）。
  - `golangci-lint run`: 0件。
  - UI動作: 一覧→ステータス遷移、在庫タブ（武器/アイテム）、模擬戦フロー（攻撃側/防御側選択→プレビュー→ログ）が従前どおり。

### 回帰確認手順（手動）
1. `make run` を実行。
2. 起動直後に「ユニット一覧」パネルが表示され、各行に「名前」「クラス/Lv」が出ていること。
3. 任意行をクリック→ステータス画面へ。HPバー/基本ステ/武器ランク/魔法ランク/装備が表示されること。
4. ステータス画面で `E`/`I` により在庫タブ（武器/アイテム）が開くこと。
5. 一覧の「模擬戦」ボタン→ポップアップで攻撃側/防御側を選択→バトルプレビューが表示されること。
6. 戻るボタンまたは `Esc`/`X` で前画面に戻れること。

## 次の作業候補（詳細案）

1) 追加Lintルールの段階導入（funlen/gocognit/cyclop）
- 目的: 複雑度抑制と長関数の早期検知。
- 手順:
  - 試行: `EXTRA_LINTERS='-E funlen -E gocognit -E cyclop' make lint-ci` をPR上で実行。
  - 対応: 指摘箇所は抽出関数化（既存の `draw*/update*` 方針を踏襲）。命名は既存規約に準拠。
  - 定着: 実行結果が0件に安定したら `.golangci.yml` に恒常ONで移管。
- 受入条件: CI 0件、描画/挙動不変。
- 想定影響: `internal/repo/inventory.go`, `internal/ui/popup/levelup.go` など。

2) text/v2 への内部移行（`uicore.TextDraw` 差し替え）
- 目的: 古い `text` API 依存を解消し、`DrawOptions` ベースへ移行。
- 方針:
  - 既存の `uicore.TextDraw/DrawWrapped` の内部実装を `text/v2` に差し替え（呼び出し側は変更不要）。
  - プリセット（タイトル/本文/小）と行間・中央寄せ等のオプション設計を固める。
- 手順:
  - スパイク: `uicore/textutil.go` に `v2` 実装を試作 → 視覚差分を確認。
  - 段階移行: まず一部画面で採用→全体へ展開。
- 受入条件: 主要画面で視覚差分が実用上無視できること。lint 0件、ビルドOK。

3) revive ルールの再強化とドキュメント整備
- 目的: 公開APIコメント整備、未使用パラメータの解消、早期return徹底。
- 手順:
  - `exported`/`package-comments` 未対応箇所にGoDocを追加。
  - `unused-parameter` 指摘は関数シグネチャ見直しまたは `_` で明示。
  - `docs/API.md` を更新（新規/変更関数の説明を追記）。
- 受入条件: revive 0件、`docs/API.md` 反映済み。

4) CI 強化（段階）
- 目的: 変更の早期検知と安定化。
- 手順:
  - `make lint-list` の出力で採用候補を見直し、CI で `EXTRA_LINTERS` を週次でONにして把握。
  - キャッシュキーの安定化（`go.sum`/`go.mod`/`.golangci.yml` をキー化）。
  - 将来的に `make verify` を導入し、lint/test/sim を一括実行。
- 受入条件: CI 時間 < 3 分、偽陽性なし。

5) pkg 層の最小テスト追加
- 目的: UIに依存しないロジックの回帰防止。
- 範囲: `pkg/game` の戦闘計算（命中/必殺/三すくみ/追撃条件）。
- 手順: テーブル駆動テストで代表ケース（命中100/0、AS差±3、地形効果）をカバー。
- 受入条件: `go test ./pkg/...` 緑、カバレッジ微増（+5%目安）。

6) 在庫/装備ロジックの整備（軽量）
- 目的: `cyclop`/`gocognit` 指摘の先回りと可読性向上。
- 対象: `internal/repo/inventory.go` の `Consume` 系、`cmd/ui_sample` の装備付け替え。
- 手順: 小関数化（例: `isWeaponID`, `isItemID`, `swapOwnerSlot`）。
- 受入条件: 挙動不変、lint 0件。

7) ドキュメント追補
- 目的: コントリビューションしやすくするためのガイド整備。
- 内容: `README` に「開発ループ（make mcp / run / lint）」「よくあるエラーと回避」「UIキーバインド一覧」を追記。
