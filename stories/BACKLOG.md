# Stories Backlog（将来のストーリー候補）

本ファイルは、作業中に見つかった「別ストーリー化が望ましい改善点/アイデア」を一時的に蓄える場所です。必要に応じて `make new-story SLUG=<slug>` で昇格し、詳細は新規ストーリー配下に展開します。

運用ルール（簡易）
- 1エントリ=1セクション。短く要点を記述（目的/背景/DoD/関連）
- 重要度が高いものは先頭へ。完了したらエントリに `[完了]` を付与し、対応ストーリーへのリンクを追記
- 週次/セッション終わりに見直し、不要になったものは削除

エントリ雛形
```
## YYYY-MM-DD: タイトル（簡潔）
- 目的: 〜〜を改善/実現する
- 背景: 〜〜が散在/重複/不明確 など
- DoD: 〜〜が達成（例: mcp グリーン、rg 残存 0）
- 参考/関連: ファイル/PR/ストーリーの相互リンク
```

初期エントリ

## [P1] 2025-09-30: CI にストーリー検証/索引生成を統合
- 目的: PR/Push 時に `make story-index` と最小の Story 検証（テンプレ必須項目の存在チェック）を自動実行し、索引とメタの整合性を常に維持する。
- 背景: 手動更新だと漏れや古い一覧が残る可能性があるため。
- DoD: GitHub Actions（または将来のCI）で `make story-index` を実行し、成果物 `stories/finish/INDEX.md` を更新/検証。検証スクリプト（軽量）で `ステータス/担当/開始` の必須チェックを行い、エラー時に失敗。
- 参考/関連: scripts/gen_story_index.sh, docs/REF_STORIES.md

## [P1] 2025-09-30: PR テンプレ/Story 参照の必須化（軽量）
- 目的: PR/コミットに対象ストーリーの識別子（`YYYYMMDD-slug`）と DoD/検証手順を含め、追跡性を上げる。
- 背景: レビュー時に関連ストーリーの特定に時間がかかる問題。
- DoD: `.github/PULL_REQUEST_TEMPLATE.md` を追加し、`Story: 20xxxxxx-slug` と DoD チェック、検証手順テンプレを含める。将来のCI導入時に軽い静的チェックを追加。
- 参考/関連: docs/REF_STORIES.md, AGENTS.md

## [P0] 2025-09-30: Provider から UI 型依存を分離（Port/Adapter 再整理）
- 目的: `internal/game/data.TableProvider` を純参照Portにし、UI 型依存を排除して移植性/テスト容易性を高める。
- 背景: 現行 `UserUnitByID(id) (uicore.Unit, bool)` が UI 型を返し、Provider が UI に結合している。
- DoD: Provider から `UserUnitByID` を撤去し、`service/ui/adapter` に `User→Unit` 変換を移管。呼び出し箇所を置換し `make mcp` グリーン。
- 参考/関連: internal/game/data/provider.go, internal/game/service/ui/adapter, docs/ARCHITECTURE.md

## [P0] 2025-09-30: ItemsTable 取得の Repo 化・キャッシュ一貫化
- 目的: `usecase.App.ItemsTable()` の直読みを廃止し、Repo/キャッシュを経由する設計に統一する。
- 背景: 参照経路が分散し、オフライン/性能面の振る舞いが読みづらい。
- DoD: `repo.ItemsRepo` を追加（または WeaponsRepo に隣接キャッシュ追加）。Usecase をRepo経由に切替、呼び出し箇所更新、`make mcp` グリーン。
- 参考/関連: internal/usecase/data.go, internal/repo, docs/DB_NOTES.md
- 補足（設計メモ）: `internal/repo` に ClassCaps/Items の Repo 追加、`UserRepo.Reload()` を実装。`App.ReloadAll()` で Weapons/User/Images の再注入を一括化。`internal/ui/core/load.go` の画像パス直書きは `config` 化または `embed` 化を検討（移管元: 旧 `tasks/`）。

## [P1] 2025-09-29: `internal/game/service/ui/apply.go` の分割検討
- 目的: UI メトリクス適用処理を責務単位で分割し、可読性・変更容易性・テスト容易性を向上する。
- 背景: 現在の `apply.go` は複数ドメイン（List/Status/Sim/Popup/Widgets）への一括適用を担っており、変更差分の把握や衝突が起きやすい。
- DoD: 機能等価のまま `apply_list.go`/`apply_status.go`/`apply_sim.go`/`apply_popup.go`/`apply_widgets.go` 等へ分割、外部 API を不変に維持、`make mcp` グリーン。
- 参考/関連: `internal/game/service/ui/apply.go`, `internal/game/service/ui/*.go`, `docs/NAMING.md`, `docs/ARCHITECTURE.md`

## [P2] 2025-09-30: 入力レイアウト設定の外部化（config 駆動）
- 目的: キー/マウス割当を設定ファイルで差し替え可能にし、`provider/input/ebiten` のレイアウト注入を標準化する。
- 背景: デフォルトレイアウトがコードに固定。
- DoD: `config/input_layout.json` を追加、`app.NewUIAppGame` で読み込み→ `ginput.Layout` へ変換。既定不在時は従来デフォルトを採用、`make mcp` グリーン。
- 参考/関連: internal/game/provider/input/ebiten, pkg/game/input, internal/game/app/core.go

## [P2] 2025-09-30: `internal/game/util` の撤去または責務特化サブパッケージ化
- 目的: 命名規約（util/helpers禁止）に沿い、汎用名の温床を除去する。
- 背景: 現状空ディレクトリが存在。
- DoD: 不要なら削除。必要なら用途別に `rng`/`rect`/`debug` 等へ再配置。README へ根拠追記、`make mcp` グリーン。
- 参考/関連: docs/NAMING.md, internal/game/util

## [P2] 2025-09-29: マウス座標デバッグ表示
- 目的: マウスカーソル位置のゲーム座標/スクリーン座標を HUD に表示して調整を容易化
- 背景: レイアウト調整やクリック判定の検証負荷が高い
- DoD: トグル可能なデバッグ HUD 実装(例: F2)、座標オーバレイ表示、UI との重なり確認、パフォーマンス影響軽微
- 参考/関連: internal/game/app/game.go, internal/game/scenes, Ebiten 入力 API

## [P2] 2025-09-29: FPS デバッグ表示
- 目的: フレームパフォーマンスの可視化
- 背景: 画面増加に伴う描画コストの監視が必要
- DoD: トグル可能な FPS 表示(HUD)、秒間平均の表示、負荷小、`make mcp` グリーン
- 参考/関連: internal/game/app/game.go, ebitenutil.DebugPrint など

## [P3] 2025-09-29: FE風バトル マップ画面の叩き追加
- 目的: FE風シミュレーションバトルのマップ画面を試作し、キャラクター一覧からの遷移と配置ロジックの骨組みを固める
- 背景: 現状は一覧表示のみで戦闘画面が未実装のため、体験の流れが断絶している
- DoD: キャラクター一覧右上にマップ画面遷移ボタン追加、平地/森/砦レイヤーを持つマップJSONマスタを新設、`db/user/usr_characters.json` の味方1体・敵2体を所定の位置へ初期配置
- 参考/関連: cmd/ui_sample, internal/game/scenes, db/master, docs/ARCHITECTURE.md, docs/NAMING.md
- 詳細（叩き要件）:
  - 10×8 グリッドを描画（平地=茶、森=緑、砦=黄）。レイヤ優先度は砦>森>平地。
  - ユーザデータから味方1体を下中央、敵2体を上左右に配置（プレースホルダ描画）。
  - Esc/「＜ 一覧へ」で戻る。読込失敗時は画面内エラー表示。
- 実装方針（最小）:
  - 画面追加: `internal/ui/screens/battlemap.go`
  - 遷移ボタン: `internal/ui/screens/list.go` から遷移。矩形/描画ヘルパは `internal/ui/widgets/buttons.go`
  - マップ定義: マスタJSON（レイヤ配列）。

## [P1] 2025-09-27: レイヤリング/依存の整理（main→app 移譲）
- 目的: 入力処理・状態遷移を `internal/app` に集約し、UI 側は描画＋イベント通知に限定。
- 背景: `main` に処理が残存し責務が曖昧。UI からの直接I/Oも点在。
- DoD: `main` の入力/状態遷移ロジックを `internal/app` へ移譲、UI直I/O撤去、`config` にパス統一、`make mcp` グリーン。
- 参考/関連: （移管元: 旧 `tasks/`）

## [P2] 2025-09-27: ドキュメント更新の追従（README/API/DB_NOTES）
- 目的: 構成変更と API の最新化をドキュメントへ反映。
- 背景: README/API/DB_NOTES に旧記述が残存。
- DoD: README/`docs/API.md`/`docs/DB_NOTES.md` を現行構成に同期、リンク整合、`make mcp` グリーン。
- 参考/関連: （移管元: 旧 `tasks/`）

## [P2] 2025-09-27: ドメインロジックの整理（UI色/文言生成）
- 目的: 予測/ログの文言生成を集約し、I18N を見据えた分離を行う。
- 背景: 文言/色マッピングが散在。
- DoD: 文言生成/色マッピングの統一ユーティリティを追加し、利用箇所を置換。
- 参考/関連: （移管元: 旧 `tasks/`）

## [P2] 2025-09-27: SQLite 準備（将来）
- 目的: JSON→SQLite へ段階移行できる下地作り。
- 背景: 将来の性能/整合性要件に備える。
- DoD: `internal/repo` に SQLite 実装の雛形を追加（`modernc.org/sqlite`）、ビルドタグで切替。`docs/DB_NOTES.md` に移行手順とトランザクション方針を追記。
- 参考/関連: （移管元: 旧 `tasks/`）

## [P2] 2025-09-27: テスト拡充
- 目的: `usecase.App` と UI 純関数のテスト強化。
- 背景: カバレッジと境界テストが不足。
- DoD: `adapter.UIToGame`、Repo キャッシュ/Reload、`ForecastAtExplain` の整合性、UI純関数（Rect/折返し幅/スケール）テストを追加。
- 参考/関連: （移管元: 旧 `tasks/`）

## [P2] 2025-09-27: UI スケーリング調整
- 目的: 表示スケールに応じた見易さ改善。
- 背景: 固定オフセットや折返し幅のハードコードが残存。
- DoD: `uicore.S` の適用範囲拡大、折返し幅/余白の自動計算、HPバーの微調整。`make mcp` グリーン。
- 参考/関連: （移管元: 旧 `tasks/`）

## [P3] 2025-09-29: Portrait のロード責務の最終整理
- 目的: 画像読込（I/O）を描画層に寄せ、view-model はパスなど純データに限定
- 背景: 現状 `ui/adapter` で PortraitLoader を注入しているが、用途により描画側での遅延読込が有利な場面がある
- DoD: view-model を `PortraitPath` に寄せる設計案の比較と方針確定（採用/見送りの記録）

## 2025-09-29: [完了] CIで UI パッケージのテストも実行する
- 目的: `ui/adapter` や `service/levelup` のテストを継続的に検証
- DoD: CI で `make test-all-ui` を追加ジョブとして実行
- 関連: Makefile `test-all-ui` 追加（本コミットで実装済み）
