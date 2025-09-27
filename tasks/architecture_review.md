# アーキテクチャレビュー — 現状評価と改善提案 (2025-09-27)

本ドキュメントは、現状リポジトリ構成のアーキテクチャレビュー結果と改善提案をまとめたものです。
対象: `/cmd`, `/internal`, `/pkg`, `/db`, `/assets`, `/tasks` 一式。

## 現状サマリ
- エントリ: `cmd/ui_sample/main.go` が状態管理（一覧/詳細/戦闘/模擬戦）、入力、永続化を包括。
- UI層: `internal/ui/{core,widgets,screens,popup}` で描画・レイアウトを分割し、`internal/ui/api.go` で外向けAPI提供。
- ドメイン層: `pkg/game` に戦闘ロジック（予測/解決/地形/三すくみ）。単体テストあり。
- データ層: マスタ `internal/model`（mst_）とユーザ `internal/user`（usr_）で明確分離。
- 資産: 画像はキャッシュ（`internal/assets.ImageCache`）経由でファイル読み込み。`embed` 未使用。
- ツール: Makefile に vet/build/test。`pkg/...` を最小対象にしている。

## 良い点
- ドメインロジックの `pkg/game` 隔離とテスト整備は明確で拡張しやすい。
- UIの関心分離（core/widgets/screens/popup）と API ファサードで呼び出しが簡潔。
- mst_/usr_ の責務分離と命名規約が一貫。
- RNG 注入可能（模擬戦/解決系）で再現性のある検証が可能。
- Makefile に環境依存ビルドの緩和があり、CI化の足掛かりがある。

## 主な懸念
- 型の重複と変換散在: `model/user/ui/pkg(game)` に同義型が並立し、`toGameUnit` が複数箇所に重複。
- ランタイムI/O過多: `model.LoadWeaponsJSON` 等を戦闘/模擬戦の都度呼び出し（複数ファイル）→毎回ファイルI/O。
- `main` の肥大化: 入力/状態遷移/ビジネスロジック/永続化が同居し、保守・テストが困難。
- リソース管理: ポートレートやJSONの再読込にキャッシュがなく、将来スケール時のパフォーマンス懸念。
- 画面スケーリング: 論理解像度固定でレイアウトも固定座標。リサイズ対応は未最適化。
- ドキュメント齟齬: `README.md` が旧ファイルパス参照（`internal/ui/ui.go`）。
- 三すくみ表記ロジックの分散: ラベル/色決め（UI側）と補正（`pkg/game`）が分かれている。
- Go ツールチェーン: `go 1.25` 指定により開発環境差異で失敗リスク（toolchain指示やCI固定が未整備）。

## 改善提案（領域別）
- レイヤリング/依存: `internal/app`（ユースケース/状態遷移）と `internal/repo`（mst/usr/武器読み込み）を導入。`cmd` は初期化のみに縮小。
- 依存注入: `App` に `WeaponsRepo`, `UserRepo`, `RNG`, `Config` を渡し、UIは `App` 経由で操作。
- データ管理/キャッシュ: JSONテーブルのキャッシュ（初回ロード→メモリ提供、Backspaceでリロード）。画像は `internal/assets` の `ImageCache` に一本化し、将来 `embed` へ移行可能に。
- パス集中管理: `internal/config` に `DBRoot`, `UserPath`, `MasterPaths` を集約。
- 型/変換の一本化: `internal/adapter` に `user/model → ui/pkg(game)` 変換を集約し、重複 `toGameUnit` を排除（単体テスト化）。
- UIスケーリング: `uicore.Metrics` で scale 算出し、`ListMargin/LineH*` 等に適用。Rect計算をレイアウト関数へ集約。
- ドメイン集約: 地形プリセット/三すくみラベル/反撃可否・射程判定を `pkg/game` で提供し、UIは結果表示のみ。
- 永続化/エラー: 保存は `UserRepo.Save` に集約し、UIから直接I/Oを排除。失敗時は非パニックでUI通知。
- テスト/品質: Adapter/Repo キャッシュ/変換テスト、`ForecastAtExplain` 恒等式検証、2RN境界テスト。UIは純関数を抽出してスナップショット/ゴールデンテスト。
- ビルド/ツール: `toolchain` 指示または CI でGoバージョン固定。`golangci-lint` ルール拡充＆ `internal/...` も対象に。
- ドキュメント: `README.md` と `docs/API.md` を現行構成に同期。`DB_NOTES.md` に Repo/Adapter 介在の移行手順を追記。

## 優先度（目安）
- 高: Repo/Adapter 導入と JSON/画像キャッシュ、`main` の責務削減、`toGameUnit` 重複解消、ドキュメント齟齬修正。
- 中: スケーリング対応、保存/通知の集約、三すくみ・地形のドメイン側集約。
- 低: `embed` 化、lint/CI 強化、UIテスト拡充、`text/v2` への移行。

## 進捗ログ
- 2025-09-27: 初版作成（現状評価と改善提案を反映）。
- 2025-09-27: ステータス=作成完了。次アクション案: Repo/Adapter 雛形の追加と `main` の薄型化。
- 2025-09-27: Repo/Adapter 雛形を追加（`internal/repo`, `internal/adapter`）。`internal/app` を新設し戦闘解決を `App.RunBattleRound` へ集約。`main.go` から戦闘ロジック/保存処理を委譲し薄型化。
- 2025-09-27: レベルアップ保存処理を `App.PersistUnit` へ委譲。`internal/ui/screens/battle.go` の変換を `adapter.UIToGame` に統一。画像キャッシュ `internal/assets` を導入し `ui/core` の読み込みを置換。
- 2025-09-27: 三すくみ関係/地形名を `pkg/game` に昇格（`TriangleRelationOf`, `TerrainPresetName`）。UIは表示のみ担当。Backspaceで Repo リロード＋画像キャッシュクリアを App API 経由に変更。
- 2025-09-27: スケーリング基盤（`uicore.Metrics`）を追加。`一覧/ボタン/戦闘/模擬戦` の主要レイアウトに適用し、ウィンドウサイズから毎フレーム更新。
- 2025-09-27: ステータス画面・レベルアップポップアップのスケール適用を追加（portrait/HPバー/行高/余白）。バトル開始ボタン/模擬戦サイド枠も拡張。
- 2025-09-27: フォントのスケール連動を実装（`uicore.MaybeUpdateFontFaces`）。ウィンドウサイズに応じて FaceTitle/Main/Small を再生成。
