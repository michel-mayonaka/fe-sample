# UI サンプル理想アーキテクチャ（合意版）

本ドキュメントは、UI サンプル（Ebiten）プロジェクトにおける「現時点の理想アーキテクチャ」を示します。目的は、役割分担の明確化、依存方向の固定、肥大化の抑止（特に Env）、および拡張・移行の手順を共有することです。

最終更新: 2025-09-29

## 1. レイヤと役割

- UI（Scenes）
  - 画面遷移・入力解釈・描画。UI 一時状態（選択・ポップアップ等）は `Session` に保持。
  - ルール: Repo/Assets/Model を直接操作しない。ユースケース（Port）を介してコマンド、`gdata.Provider` を介してクエリ。

### 1.1 UI 補助サブパッケージ（責務分割）

- `internal/game/ui/layout`: レイアウト計算（矩形/サイズ）。副作用なし。
- `internal/game/ui/draw`: 描画関数（見た目）。レイアウトに依存、I/Oはしない。
- `internal/game/ui/view`: 表示用データ構造（行モデル等）。画像は参照（`*ebiten.Image`）に留める。
- `internal/game/ui/adapter`: view-model 生成（テーブル/ユーザデータ→view）。`PortraitLoader` 抽象で画像読込を疎結合化。

- Usecase（アプリケーションサービス）
  - ビジネスルールの調停・副作用の集約（保存・巻き戻し・整合）。
  - Repo を介して永続化へアクセス。UI 資産（画像キャッシュ等）には直接触れない。

- Repo（インフラ層）
  - 抽象化されたリポジトリインターフェースと実装（JSON/将来 SQLite）。
  - 例: `UserRepo`, `WeaponsRepo`, `InventoryRepo`。

  Repo の単位（集約）方針:
  - 目的は「保存境界を安全に扱うこと」。モデルと常に1:1である必要はない。
  - 本プロジェクトでは「在庫」を1つの集約と見なし、`InventoryRepo` が `usr_weapons` と `usr_items` を横断管理する。
  - 理由: 一覧UI/耐久消費/装備付け替えなど“武器とアイテムをまたぐ操作”を同一トランザクション（将来DB）や同一保存タイミングで扱えるため。
  - 代替: 厳密に1:1へ分割する場合は `UserWeaponsRepo` と `UserItemsRepo` に分け、Usecase で協調保存する（Save/Reload の同期コストが増える）。

- Model / User / Adapter
  - Model: マスタ定義（武器・アイテム等）。
  - User: ユーザセーブのスキーマとテーブル（ID 索引）。
  - Adapter: UI と Game/Core の相互変換、表示値の計算補助。

- Game Runtime（Runner/Game）
  - Ebiten の `Game` 実装、SceneStack 管理、グローバルトグル（データ再読み込み、ヘルプ表示など）。
  - UI レイヤに属する副作用（例: 画像キャッシュクリア）を受け持つ。

- Data Provider（`internal/game/data`）
  - 読み取り専用テーブル群のプロバイダ。UI からの「参照」を一本化。

## 2. 依存原則（CQRS っぽい分離）

- コマンド（更新）: Scenes → Usecase（Ports）→ Repo
- クエリ（参照）  : Scenes → `gdata.Provider` → テーブル/在庫（例: `WeaponsTable`/`ItemsTable`/`UserWeapons`/`UserItems`）
- Scenes は Repo/Assets/Model を直接インポートしない（Adapter/Provider 経由のみ）。
- Scenes は必要に応じて `ui/layout|draw|view|adapter` を参照する（UI 内の責務分割）。
- Usecase は UI 資産に触れない（画像キャッシュクリアは UI 側で実施）。

## 3. Port（ユースケース境界）の分割方針

env.go に全ユースケースを集約しない。ドメイン単位で Port を分割し、Scene は必要最小の Port のみを前提にする。

推奨ポート（例）:
```go
// DataPort: データ再読み込み/保存など横断的操作
type DataPort interface {
    ReloadData() error
    PersistUnit(u uicore.Unit) error
}

// BattlePort: 戦闘解決（本番）
type BattlePort interface {
    RunBattleRound(units []uicore.Unit, selIndex int, attT, defT gcore.Terrain) ([]uicore.Unit, []string, bool, error)
}

// InventoryPort: 在庫と装備操作（更新系のみ）
type InventoryPort interface {
    EquipWeapon(unitID string, slot int, userWeaponID string) error
    EquipItem(unitID string, slot int, userItemID string) error
}

// 注: 以前は上記を合成する `UseCases` を段階移行用に併用していましたが、
// 工程10（2025-09-28）で撤去し、現在は各 Port 参照のみに統一しています。
```

導入手順:
- `scenes/ports_*.go` に Port を定義（UI 側の最小依存）
- `usecase/*.go` で実装（`App` が各ポートを満たす）
- Scene は自分が使う Port だけを `Env` から受け取る（または Scene コンストラクタ引数で個別 DI）
- 参照（Query）は `gdata.Provider()` に統一（`WeaponsTable/ItemsTable/UserWeapons/UserItems`）。

## 4. ディレクトリ構成（理想）

```
cmd/
  ui_sample/
    main.go

internal/
  game/
    app/                  # Ebiten Runner/Window/入力の束ね（ユースケースは含めない）
      core.go             # Game 構築（Ports/Provider注入・Env/Session生成）
      game.go             # ebiten.Game 実装（Update/Draw/グローバル操作）
      runner.go           # SceneStack の更新/描画

    data/                 # 読み取り専用テーブル/在庫の Provider（UI 参照用）
      provider.go         # SetProvider/Provider, TableProvider（WeaponsTable/ItemsTable/UserWeapons/UserItems/UserTable/UserUnitByID/EquipKindAt）

    scenes/               # UI シーン群（UI の状態遷移と描画）
      # ports.go         # （廃止）合成 UseCases。現在は各 Port のみ使用。
      ports_data.go       # DataPort
      ports_battle.go     # BattlePort
      ports_inventory.go  # InventoryPort
      env.go              # Env（最小: Port/メタ情報/Session）
      session.go          # Session（UI一時状態: 選択・ポップアップ等）
      common/...          # 共有UI（popup等）
      character_list/...  # 一覧
      status/...          # ステータス
      inventory/...       # 在庫
      sim/...             # 模擬戦

    service/              # 入力・UI描画ユーティリティ
      ui/...              # テキスト/レイアウト/描画/フォント
      ...
    ui/                   # UI補助サブパッケージ群（責務分割）
      layout/             # 座標計算（矩形）
      draw/               # 描画関数（見た目）
      view/               # 表示用データ（VM）
      adapter/            # VM 生成（テーブル/ユーザ→view）

  usecase/                # アプリケーションサービス（Ports 実装）
    facade.go             # struct App, New(), 依存注入・共通メソッド
    data.go               # DataPort 実装（ReloadData/PersistUnit）
    battle.go             # BattlePort 実装（RunBattleRound）
    inventory.go          # InventoryPort 実装（EquipWeapon/EquipItem）

  repo/                   # リポジトリIFと実装（JSON→将来SQLite）
    user.go
    weapons.go
    inventory.go

  model/                  # マスタ定義（武器・アイテム等）+ user モデル（純粋型）
    user/                 # ユーザセーブ定義（純粋型・入出力なし）
  infra/
    userfs/               # ユーザセーブの JSON 入出力（バックエンド）
  adapter/                # UI↔Game の変換/補助
  assets/                 # 画像/音声ローダとキャッシュ
  config/                 # パス/ビルドタグ等

assets/                   # 画像・フォント・音源
db/                       # JSON DB（将来 SQLite へ移行）
docs/                     # ドキュメント
```

命名規約（抜粋）
- Port 定義: `internal/game/scenes/ports_*.go`
- Usecase 実装: `internal/usecase/*.go`（ドメイン単位）
- テーブル参照は `gdata.Provider().XxxTable()` に統一（Port には載せない）

## 5. Env/Session の責務分離

Env（最小）:
- ポートの参照
- `UserTable`（読み書きテーブル参照を 1 箇所に集約）
- `UserPath` `RNG` 等、アプリ起動時に決定されるメタ情報
- `Session`（UI 一時状態）への委譲

Session（UI 状態）:
- `Units` `SelIndex` `CurrentSlot` などの UI 表示用の一時的値
- ポップアップの開閉・直後状態・在庫タブなど
- ドメイン保存/巻き戻しは行わない（Usecase 経由でのみ変更確定）

この分離により、`env.go` 自体の肥大化を防ぎます。

## 6. シーケンス例

### 6.1 武器装備の確定（選択スロットへ割当）
1) ユーザ操作（在庫ポップアップで候補行を Confirm）
2) Scene → `InventoryPort.EquipWeapon(unitID, slot, userWeaponID)` を呼ぶ
3) Usecase:
   - 既オーナーのスロットから外す→元の装備を巻き戻し
   - 対象ユニットの `slot` に `userWeaponID` を設定
   - `UserRepo.Save()` で永続化
4) Scene: 表示同期（`refreshUnitByID` 等）

ポイント:
- UI は Repo を直接書き換えない
- 複数テーブル更新（巻き戻し）をユースケースに集約

### 6.2 データ再読み込み（ホットリロード）
1) グローバルトグル（Backspace 長押し等）で発火
2) `DataPort.ReloadData()` → Repo 側のキャッシュ再読込
3) UI で `assets.Clear()`（画像キャッシュ等の UI 資産クリア）
4) `uicore.LoadUnitsFromUser` で UI 用ユニットを再構築
5) `gdata.SetProvider(usecaseApp)` により参照テーブルを最新化

### 6.3 戦闘解決（本番）
1) 戦闘開始 → Scene から `BattlePort.RunBattleRound(...)`
2) Usecase 内で `adapter.UIToGame` して解決→HP/耐久を反映→`UserRepo/InventoryRepo.Save()`
3) Scene がログを受け取りオーバレイ表示

## 7. クエリ/コマンドの分離（Why Provider?）

- クエリは参照専用・副作用なし。`gdata.Provider()` を用いることで UI からの参照経路を 1 本化できる。
- コマンドは副作用を伴う。Usecase に集約することで UI 直書きのリスク（半更新・整合崩れ）を回避。

### 7.1 図（Provider と Port）
```
      +-------------------+            +-------------------+
      |      Scenes       |            |      Usecase      |
      |  (UI: input/draw) |            |   (Application)   |
      +---------+---------+            +----------+--------+
                |                               |
      Query --> | gdata.Provider()              | <-- Command (Ports)
                v                               v
      +---------+-------------------------------+----------+
      |                  App (TableProvider, Ports)       |
      |     - WeaponsTable/ItemsTable                     |
      |     - UserWeapons/UserItems (snapshot)            |
      |     - InventoryPort(EquipWeapon/EquipItem)       |
      |     - DataPort(ReloadData/PersistUnit)           |
      +---------------------+----------------------------+
                            |
                            v
                  +---------+----------+
                  |       Repo         |
                  | (User/Weapons/Inv) |
                  +--------------------+
```

Provider と Repository の違い（明確化）:
- Provider（読み取り専用）
  - 役割: 参照テーブルとユーザ在庫スナップショットの提供（例: `WeaponsTable()`/`ItemsTable()`/`UserWeapons()`/`UserItems()`）。
  - 特性: 追加・更新・保存は行わない（Query に特化）。
- Repository（更新を含む）
  - 役割: ユーザ状態や在庫などの保存境界を扱う（例: `UserRepo`、`InventoryRepo`）。
  - 特性: 更新（`Update/Consume`）と保存（`Save`）、再読み込み（`Reload`）。将来的に DB トランザクションの境界。

## 8. テスト方針

- Usecase 単体テスト
  - Repo のフェイク（インメモリ実装）で巻き戻し/装備移譲の整合を検証
  - 例: EquipWeapon の「所有者移動＋巻き戻し＋保存」の一貫性

- Scene の軽量結合テスト
  - 入力→Intent→状態遷移（選択/ポップアップ）を中心に
  - テーブル参照は Provider をフェイク化（DI）して検証

## 9. マイグレーション（JSON → SQLite）

- Port/Usecase から見える境界は不変（Repo 実装差し替え）
- テーブル参照（Provider）はバックエンド差し替えによる影響を受けない
- 手順（例）:
  1) `repo/sqlite/*` 実装を追加
  2) `NewUIAppGame()` の注入先を JSON→SQLite に切替
  3) 回帰テスト（Usecase テストが後方互換を担保）

## 10. 命名・スタイル（補足）

詳細な命名規約は `docs/NAMING.md` を参照してください。本節は抜粋と背景説明のみを記載します。

- ファイル名
  - Port 定義: `ports_data.go`, `ports_battle.go`, `ports_inventory.go`
  - Usecase 実装: `data.go`, `battle.go`, `inventory.go`（`facade.go` に `App` 本体と DI）
- 画面: `character_list.go`, `status.go`, `inventory.go` など機能名ベース

- GoDoc
  - エクスポート識別子には役割を一行で（例: `// BattlePort は戦闘解決のユースケース境界です。`）

### 10.x 型参照の暫定許容
- `scenes` から `internal/model` 型を参照するケースは副作用が無い範囲で暫定許容。
- 将来的には DTO 化または Provider 拡張で置換し、`scenes` からの model import を段階削減する。

## 11. 導入ステップ（現在実装との差分を埋める）

1) `scenes/ports_*.go` を追加し、env.go から境界定義を移動
2) Scene 側の依存を必要最小 Port に変更
3) `usecase/*` をドメイン別ファイルに分割（`facade.go` + 各ドメイン）
4) Provider の拡張（必要なら Items 等を追加）
5) Usecase のユニットテスト追加

## 12. 直近の議論要約（設計方針の確認）

### 12.1 env.go の肥大化懸念について
- Env に全ユースケースを積み増すのは避ける。
- 共有 UI 状態は `Session` に集約し、Env は「Port 参照＋メタ情報＋Session」の最小構成に留める。
- 依存の原則を固定：
  - コマンド（更新）は Port（Usecase）経由。
  - クエリ（参照）は `gdata.Provider` 経由。
- ファイル名と責務の整合を明確化：
  - `env.go` = 環境コンテナ、`session.go` = UI 一時状態、`ports_*.go` = 境界定義、`usecase/*.go` = 実装。

### 12.2 Port（境界）に関する合意
- Port は「機能単位のユースケース契約」を小さく分割したもの（例：Data/Battle/Inventory）。
- 「インターフェースは使う側に置く」原則に従い、Port 定義は `scenes` 配下に配置する。
  - 利点：循環依存の回避、UI 要件の近接管理、Discoverability の向上。
- 将来、CLI/サーバなど別フロントエンドでも共通化する必要が出たら、`internal/ports` 等へ昇格を検討。

### 12.3 Scene と Port の関係
- 各 Scene は自分が必要とする Port だけを利用（最小依存）。
- 依存の持ち方は 2 案：
  - A) Scene のコンストラクタに必要 Port を個別 DI（より厳密に最小化）。
  - B) Env に複数 Port を載せ、Scene は必要なものだけ参照（導入容易・段階移行向け）。
- 具体例：
  - Status → `DataPort`（`PersistUnit`）
  - Inventory → `InventoryPort`（`Inventory/EquipWeapon/EquipItem`）
  - Battle 本番 → `BattlePort`（`RunBattleRound`）
  - グローバル再読み込み（Game 側）→ `DataPort`（`ReloadData`）
- クエリは常に Provider（例：`WeaponsTable`）で統一。Port は原則コマンド専用。

### 12.4 今後のアクション（実装に移す際の順序）
1) `scenes/ports_*.go` を追加し、Env から境界定義を切り出す。
2) Scene ごとの依存を必要最小 Port へ縮小（A または B を選択）。
3) `usecase` をドメイン別ファイル（`data.go`/`battle.go`/`inventory.go`）に整理。
4) Provider の拡張（必要なら Items 等を追加）。
5) Usecase のユニットテスト（装備移譲の巻き戻し、保存の一貫性）を追加。

## 13. 旧 ARCHITECTURE.md から引き継ぐ指針

### 13.1 コア API（最小）
```go
// game.Scene
Update(ctx *game.Ctx) (next game.Scene, err error)
Draw(screen *ebiten.Image)

// game.SceneStack
Current() Scene; Push(Scene); Pop() Scene; Replace(Scene); Size() int

// game.Ctx
// DT/Frame/ScreenW/ScreenH と Input/Assets/Audio/Camera/UI/Rand/Debug を保持

// actor.IActor
Update(*game.Ctx) bool; Draw(*ebiten.Image); Layer() int
```

### 13.2 Update 順序契約（全 Scene で遵守）
1) Input（Snapshot 固定）
2) Script/AI（重い処理は分割）
3) Physics/Board（座標・ZOC・当たり）
4) Resolve（コマンド確定・状態更新）
5) Audio（キュー適用）
6) GC/Spawn（死活整理）
7) Draw（レイヤ順）

### 13.3 入力（抽象アクション）と運用
- アクション: `Up/Down/Left/Right/Confirm/Cancel/Menu/Next/Prev` に加え、便宜上 `OpenWeapons/OpenItems/EquipToggle/Slot1..5/Unassign` を定義。
- Press/Down の使い分け:
  - Press: UI トグル/決定/戻る等の瞬間操作（例: Confirm で決定、Cancel で戻る）。
  - Down : 押下継続に意味がある操作（例: Menu=Backspace 長押しでデータリロード）。

### 13.4 データ提供（DI）: TableProvider の原則
- 目的: UI を取得実装（JSON/SQLite/メモリ）から切り離すための読み取り経路統一。
- 実装: `internal/game/data.TableProvider` を App/Usecase が実装し、`data.SetProvider(app)` で注入。
- 利用: Scene は `data.Provider().WeaponsTable()/ItemsTable/UserWeapons/UserItems` など参照専用メソッドを用いる（コマンドは Port）。

## 14. 運用・適用

- 本ドキュメントを唯一の設計指針とする。既存の `docs/architecture.md`（小文字版）は重複のため後続で削除またはリンク集約を検討。
- 運用: 本アーキテクチャに沿って、複数セッションに分けて段階的に現行コードを整備（Port 分割→依存差し替え→Provider 拡張→テスト強化）。
