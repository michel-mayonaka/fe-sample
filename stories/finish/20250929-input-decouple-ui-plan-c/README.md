# 20250929-input-decouple-ui-plan-c — 入力ロジックのUI非依存化（案C）

ステータス: [完了]
担当: @tkg-engineer

## 目的・背景
- 入力処理（キー/パッド/マウス）のロジックを UI フレームワーク（ebiten）から切り離し、ドメイン層でテスト可能にする。
- 既存 UI 実装に依存した条件分岐が散在し、`pkg/game` のテスト容易性と再利用性を阻害しているため、依存逆転により改善する（案C）。

## 案C（方針・要点）
- `pkg/game/input` にドメインイベント・インタフェース（`InputSource`/`InputEvent`/`ControlState`）を定義。
- UI 依存はアダプタ層に隔離（`internal/game/provider/input/ebiten`）。`ebiten`→ドメインイベントへの変換のみを担当。
- Scene/Usecase は `ebiten` を直接 import しない。`InputSource` 経由で状態取得。
- マッピング（キー→操作）は設定主導にし、ユニットテストで検証可能にする。

## スコープ（成果）
- `pkg/game` から UI 依存の import を除去し、`go vet`/`go test` が単独で通る。
- `internal/game/provider/input/ebiten` 実装により、既存 UI サンプルでの操作性を維持。
- 入力マッピングのテーブル/コードとテストを整備（最小カバレッジ確保）。
- ドキュメント更新（ARCHITECTURE / API / NAMING 整合）。

## 受け入れ基準（Definition of Done）
- [ ] `make mcp` が成功（`vet/build/lint/test`）。
- [x] `pkg/game` と `internal/usecase` のユニットテストが入力マッピングを検証（usecase 側は pkg のマッピングを間接検証）。
- [x] `cmd/ui_sample` の実行で、従来と同等の基本操作（移動/決定/キャンセル）が機能（ビルド確認済み）。
- [x] `docs/architecture/README.md` に依存関係の更新、`docs/SPECS/reference/api.md` に公開 I/F 追記、`docs/KNOWLEDGE/engineering/naming.md` との整合が確認できる。

## 工程（サブタスク）
- [x] 001 設計固め（案Cの具体化と命名）: `tasks/001-design.md`
- [x] 002 ドメイン I/F とイベント定義: `tasks/002-domain-input-types.md`
- [x] 003 ebiten アダプタ実装: `tasks/003-adapter-ebiten.md`
- [x] 004 Scene から UI 依存排除: `tasks/004-refactor-scenes.md`
- [x] 005 テスト整備（マッピング/モック）: `tasks/005-tests.md`
- [x] 006 ドキュメント更新: `tasks/006-docs.md`

## 計画（目安）
- 見積: 6〜10 時間（2〜3 セッション）
- マイルストン: M1=設計/I/F、M2=アダプタ+リファクタ、M3=テスト+ドキュメント

## 進捗・決定事項（ログ）
- 2025-09-29: ストーリー作成・方針草案（案C）

## リスク・懸念
- UI サンプルの入力分岐が多い場合の影響範囲。操作感の回帰リスク。
- キーマッピング仕様が未確定の場合の再作業。

## 関連
- Docs: `docs/architecture/README.md`, `docs/SPECS/reference/api.md`, `docs/KNOWLEDGE/engineering/naming.md`, `AGENTS.md`
