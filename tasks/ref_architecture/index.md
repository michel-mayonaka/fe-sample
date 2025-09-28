# アーキテクチャ再整備ロードマップ（ref_architecture）

本ドキュメントは、新しい ARCHITECTURE.md に準拠して現行コードを段階的に整備するための工程表です。1 工程 = 1 セッション完了可能な粒度で定義します。各工程の詳細タスクは、承認後に本ディレクトリ直下へ Markdown（1 工程 1 ファイル）として作成します。

## ゴール
- Port/Provider/CQRS 方針に沿った依存の明確化と責務分離。
- Env 肥大化の抑止（Session への集約／Port への切り出し）。
- UI からのコマンドは Usecase 経由、クエリは Provider 経由に統一。
- ドキュメント/テストと実装が整合している状態の確立。

## 進め方（セッション運用）
- 1 セッションで 1 工程のみ対応（差分を小さく、ロールバック容易に）。
- 各工程は「目的/作業項目/完了条件/影響範囲/注意点」を明記した個別 md を作成。
- 実装中に新たな論点が出た場合は index.md を先に更新（工程の順序/内容を調整）。

## 作業時の注意・留意事項（セッション横断で共有）
- 構成原則
  - コマンドは Port（Usecase）経由、クエリは `gdata.Provider` 経由に統一。
  - 「使う側が定義する小さな interface」を維持（Port は Scene 側に定義）。
  - Env は最小（Port 参照＋メタ情報＋Session）。UI 状態は Session へ。
- 進捗記録
  - 各タスクの進捗・決定事項・懸念は、そのタスクに紐づく md に時系列で追記（`Progress`/`Notes` セクション）。
- 互換性/段階移行
  - 段階移行中はブリッジ（合成 UseCases や一時的なフィールド並存）を許容。
  - 既存呼び出しの置換は Scene 単位で行い、1 セッションでは 1 画面（または 1 機能）に限定。
- 品質ゲート
  - 各工程の前後で `make mcp` を実行（`check-all`＋`lint`）。
  - 可能なら `go test ./...` も併走（ロジック層中心）。
  - 影響が UI 表示に及ぶ場合は簡易動作確認（起動・対象画面の操作）を実施。
  - 変更がドキュメント/コメント/API に及ぶ場合は `docs/` と `docs/API.md` を同セッションで更新。
- コミット/PR
  - Conventional Commits（日本語サマリ）。1 工程 1 PR を基本。差分は小さく保つ。
  - PR には目的/影響範囲/検証手順を記載し、関連ドキュメントへのリンクを添付。
- 命名/配置
  - Port 定義: `internal/game/scenes/ports_*.go`、Usecase 実装: `internal/usecase/*.go`。
  - Items など新規テーブルは Provider 拡張（Port ではなく Provider に追加）。
- 検索/置換の指針
  - 大域置換は避け、`rg` 等で発生箇所を明示しながら Scene 単位で差し替え。

## 工程一覧（1 工程 = 1 セッション）

1. 01_define_ports（Port 骨子の追加）
   - 目的: `scenes` に `ports_data.go`/`ports_battle.go`/`ports_inventory.go` を追加し、小さな境界を定義（実装なし）。
   - 完了条件: ビルド成功、`usecase.App` が各 Port を満たすコンパイル保証コメントを追加。
   - 影響範囲: 型定義のみ。既存呼び出しの変更は行わない。

2. 02_env_wire_ports（Env に Port 参照を追加）
   - 目的: `Env` に `DataPort/BattlePort/InventoryPort` フィールドを追加し、現状の `UseCases` と並存配線（ブリッジ）。
   - 完了条件: 既存動作を変えずにビルド成功。`NewUIAppGame` で Port を注入。

3. 03_status_use_data_port（Status 画面の依存切替）
   - 目的: `PersistUnit` 呼び出しを `UseCases` から `DataPort` へ切替。
   - 完了条件: ビルド成功、Status の保存動作を手動確認。

4. 04_inventory_use_inventory_port（在庫画面の依存切替）
   - 目的: 在庫参照・装備確定を `InventoryPort` 経由に統一（UI 直書きを排除）。
   - 完了条件: ビルド成功、装備移譲・巻き戻し・保存を手動確認。

5. 05_battle_use_battle_port（戦闘開始フローの導線接続）
   - 目的: 本番戦闘導線があれば `BattlePort.RunBattleRound` を接続（なければスキップ）。
   - 完了条件: ビルド成功、想定ログ・HP/耐久の反映を確認。

6. 06_provider_unify_lookup（Provider 参照の統一チェック）
   - 目的: UI からのテーブル参照が `gdata.Provider` に統一されていることを確認・是正。
   - 完了条件: 直参照が 0 件、ビルド成功。

7. 07_provider_extend_items（ItemsTable の Provider 化）
   - 目的: `TableProvider` に `ItemsTable()` を追加し、アイテム参照経路を統一。
   - 完了条件: ビルド成功、在庫アイテム表示が従来通り機能。

8. 08_session_helpers（UI 状態ヘルパの移管）
   - 目的: `Selected/SetSelected` など UI 状態操作を `Session` に移動し、呼び出しを差し替え。
   - 完了条件: ビルド成功、一覧/ステータスの選択挙動を確認。

9. 09_usecase_tests_inventory（ユースケースの最小テスト追加）
   - 目的: `EquipWeapon/EquipItem` の所有者移動＋保存一貫性テストを追加（フェイクリポジトリ）。
   - 完了条件: `go test ./...` 成功、主要ハッピーパス/境界系の網羅。

10. 10_remove_aggregated_usecases（合成 UseCases の段階的撤去）
    - 目的: Env/Scene から合成 `UseCases` を除去し、Port 参照のみへ移行。
    - 完了条件: ビルド成功、影響画面の動作確認。

11. 11_docs_sync（ドキュメント同期）
    - 目的: `docs/API.md`/`README.md` の更新（Port/Provider 方針、利用例、依存図）。
    - 完了条件: ドキュメントにコンパイル可能なサンプル・参照が掲載されている。

12. 12_cleanup_legacy_docs（重複ドキュメント整理）
    - 目的: `docs/architecture.md`（小文字）の統合/リンク整理。
    - 完了条件: 重複/齟齬のない状態、README からのリンク確認。

13. 13_repo_sqlite_skeleton（将来の SQLite スケルトン）
    - 目的: `repo/sqlite` の骨組み追加（未配線）。
    - 完了条件: ビルド警告なし、TODO/インターフェース整合コメントを残す。

14. 14_final_audit（仕上げと棚卸し）
    - 目的: 参照経路/命名/コメントを横断チェック。`-race -cover` を含む最終検証。
    - 完了条件: `go test ./... -race -cover` 成功、既知課題は `tasks/known_issues.md` へ集約。

15. 15_usecase_split_files（Usecase のファイル分割）
    - 目的: `internal/usecase/app.go` を役割別ファイル（`facade.go`/`data.go`/`battle.go`/`inventory.go`）に分割し、保守性を向上。
    - 完了条件: `make mcp` 成功。機能挙動不変（コンパイル時の型実装も維持）。

16. 16_comment_style_guidelines（コメント記法の統一方針）
    - 目的: GoDoc の書式（日本語一行要約「X は …」/パッケージコメント必須/最小限）を明文化し、コードと docs の役割分担を固定。
    - 完了条件: `docs/COMMENT_STYLE.md` 追加、主要パッケージにパッケージコメントを付与（最低限）。

## 各工程の成果物と完了報告フォーマット
- 各工程の md に以下を記載（テンプレートは最初の工程作成時に提示）
  - 目的・背景
  - 変更点（ファイル/関数単位）
  - 互換性/移行ノート（あれば）
  - 動作確認手順（CLI/GUI）
  - ビルド/テスト結果（実行コマンドと要約）
  - 既知の懸念・次工程への引き継ぎ

## リスクと緩和策（共通）
- 広範囲置換に伴う破壊的変更 → Scene 単位・工程単位で限定し、小さな PR を徹底。
- Provider 追加時の循環依存 → Port/Provider の所在（scenes/usecase/data）を遵守。
- テスト不在による回 regress → ユースケースの核となる操作に最小テストを順次追加。

---

本 index の内容に問題がなければ、工程 01 の詳細 md（01_define_ports.md）から順次作成します。
