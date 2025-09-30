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

## [P3] 2025-09-29: Portrait のロード責務の最終整理
- 目的: 画像読込（I/O）を描画層に寄せ、view-model はパスなど純データに限定
- 背景: 現状 `ui/adapter` で PortraitLoader を注入しているが、用途により描画側での遅延読込が有利な場面がある
- DoD: view-model を `PortraitPath` に寄せる設計案の比較と方針確定（採用/見送りの記録）

## 2025-09-29: [完了] CIで UI パッケージのテストも実行する
- 目的: `ui/adapter` や `service/levelup` のテストを継続的に検証
- DoD: CI で `make test-all-ui` を追加ジョブとして実行
- 関連: Makefile `test-all-ui` 追加（本コミットで実装済み）
