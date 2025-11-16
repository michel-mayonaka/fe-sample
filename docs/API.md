# API ドキュメント（Markdown）

本ドキュメントは、このサンプルの型・関数・定数の振る舞いを簡潔に説明します。実装コメントは最小限とし、詳細は本ファイルを更新してください。

## internal/game
- `const LevelUpExp = 100`: レベルアップに必要な経験値。
- `const LevelCap = 20`: レベル上限。
- `type Ctx`: フレーム共通情報とサービス群。`DT, Frame, ScreenW, ScreenH, Input, Assets, Audio, Camera, UI, Rand, Debug`。
- `type Scene`: `Update(*Ctx) (Scene, error)`, `Draw(*ebiten.Image)`。
- `type SceneStack`: `Current/Push/Pop/Replace/Size` を提供する LIFO スタック。
- `package actor.IActor`: `Update(*game.Ctx) bool`, `Draw(*ebiten.Image)`, `Layer() int`。
- `package service.Input`: 抽象アクション `Up/Down/Left/Right/Confirm/Cancel/Menu/Next/Prev` と、
  便宜アクション `OpenWeapons/OpenItems/EquipToggle/Slot1..5/Unassign`、
  地形切替 `TerrainAtt1..3/TerrainDef1..3` を提供。`Snapshot/Press/Down` に加え、
  テスト用の `SnapshotWith(func(ebiten.Key) bool)` を用意。

- `package ui/input`: UI 層向けの最小API（段階移行の薄いシム）。
  - 役割: `pkg/game/input` の `Action`/`Reader` を再公開し、既存コードの互換を維持。
  - `func WrapDomain(ginput.Reader) Reader` を提供。旧 `WrapService` は互換のため残置。

## internal/model（マスタデータ / mst_）
- `type Character`
  - 初期値のみ保持（`Name`, `Class`, `Portrait`, `Stats`, `Growth`, `Weapon`, `Magic`, `Equip(max)` など）。
- `func LoadFromJSON(path string) (*Table, error)`: JSON 読込。
- `func (*Table) Find(id string) (Character, bool)`: 取得。
備考: マスタは初期値のみ。レベルごとの能力は保持しない（成長率に依存し可変）。ファイル名は `mst_*.json`。パスは `db/master/`。

- `type ClassCaps`, `LoadClassCapsJSON`: クラス能力上限（`mst_class_caps.json`）
- `type Weapon`, `LoadWeaponsJSON`: 武器性能（`mst_weapons.json`）
- `type Item`, `LoadItemsJSON`: アイテム性能（`mst_items.json`）

## internal/model/user（ユーザデータ / usr_ モデル）
- `type Character`/`type Table`
  - 現在値を保持する純粋型とインメモリ索引（入出力なし）。
- `func NewTable(rows []Character) *Table`: 行から索引を構築。
- `func (*Table) Find(id string) (Character, bool)`: 取得。
備考: ファイル名は `usr_*.json`。パスは `db/user/`（入出力は次節）。

## internal/infra/userfs（ユーザJSON I/O）
- `func LoadTableJSON(path string) (*user.Table, error)`
- `func SaveTableJSON(path string, t *user.Table) error`
- `func LoadUserWeaponsJSON(path string) ([]user.OwnWeapon, error)`
- `func SaveUserWeaponsJSON(path string, rows []user.OwnWeapon) error`
- `func LoadUserItemsJSON(path string) ([]user.OwnItem, error)`
- `func SaveUserItemsJSON(path string, rows []user.OwnItem) error`

## internal/game/service/ui（UI描画ユーティリティ）
- `type Unit`
  - `Name`, `Class`, `Level`, `Exp`, `HP`, `HPMax`: 基本情報。
  - `Stats Stats`: 能力値（力/魔力/技/速さ/幸運/守備/魔防/移動）。
  - `Equip []Item`: 装備（耐久制）。
  - `Portrait *ebiten.Image`: ポートレート画像（任意）。
  - `Weapon WeaponRanks`: 物理系武器ランク（剣/槍/斧/弓）。
  - `Magic MagicRanks`: 魔法系ランク（理/光/闇/杖）。
  - `Growth Growth`: 成長率（%）。
- `type Stats`: 能力値セット。
- `type Growth`: 成長率（%）。
- `type Item`: 装備の耐久（`Name`, `Uses` 残り, `Max` 上限）。
- `type WeaponRanks`: 物理武器ランク。
- `type MagicRanks`: 魔法系ランク。
- 代表的関数:
  - `TextDraw`, `DrawPanel`, `DrawFramedRect`, `DrawHPBar`, `DrawPortrait(Placeholder)` など描画ヘルパ。
  - `ListMarginPx`, `S`, `LineHSmallPx` などメトリクス計算。
  - `MaybeUpdateFontFaces` などフォント管理。

注意（非推奨）:
- `UnitFromUser`, `LoadUnitsFromUser` は互換目的で残置。実装は `ui/adapter` へ委譲され、
  `internal/game/ui/adapter.UnitFromUser` / `BuildUnitsFromProvider` をブリッジ経由で呼び出します。
  新規コードは直接 adapter 側の API を利用してください。

補足: 画面描画は各 Scene に配置（例: ステータス/戦闘プレビュー）。本パッケージは汎用ウィジェットを提供します。

## internal/adapter（UI<->ロジック変換）
- `func UIToGame(wt *model.WeaponTable, u ui.Unit) gcore.Unit`: UIユニットから戦闘用`gcore.Unit`へ変換（先頭装備）。
- `func AttackSpeedOf(wt *model.WeaponTable, u ui.Unit) int`: 攻撃速度（武器重量考慮、未設定時は速さ）。
  - `internal/game/ui/adapter.UnitFromUser(c usr.Character, pl PortraitLoader) ui.Unit`: Provider参照に基づくUI化。
  - `internal/game/ui/adapter.BuildUnitsFromProvider(pl PortraitLoader) []ui.Unit`: Provider→一覧生成。

## internal/game/data（テーブル/在庫のDIプロバイダ）
- `type TableProvider interface { WeaponsTable() *model.WeaponTable; ItemsTable() *model.ItemDefTable; UserWeapons() []user.OwnWeapon; UserItems() []user.OwnItem; UserTable() *user.Table; EquipKindAt(unitID string, slot int) (bool,bool) }`
- `func SetProvider(p TableProvider)`: アプリ側実装を注入（推奨ルート）。
- `func Provider() TableProvider`: 現在のプロバイダ取得。

実装メモ:
- 既定実装は `internal/usecase.App` で、`WeaponsTable()/ItemsTable()` はそれぞれ `WeaponsRepo`/`ItemsRepo` のキャッシュを参照します（Scene 層からの JSON 直読みは禁止）。

利用指針（Provider=参照専用）:
- Scene は `data.Provider().WeaponsTable()/ItemsTable()/UserWeapons()/UserItems()` 経由で参照。
- 旧 `scenes.SetWeaponTable/WeaponTable` は廃止。JSON直読みフォールバックは開発ツール用途のみ。
 - UI 用の `uicore.Unit` 変換は Provider では行わず、`internal/game/ui/adapter.UnitFromUser` に集約する。

Provider と Repository の役割の違い:
- Provider: 読み取り専用の参照提供（Query）。追加/更新/保存はしない。
- Repository: 追加・更新・保存（Command）を扱う。例: `UserRepo.Update/Save`, `WeaponsRepo.Reload`, `ItemsRepo.Reload`, `InventoryRepo.Consume/Save/Reload`。

## internal/game/scenes/sim（模擬戦シーン）
- `type Sim`: シーン本体。`NewSim(env, atk, def)`で生成。
- 更新フロー: `Update → scHandleInput → scAdvance → scFlush`。
- 主要フィールド: `simAtk/simDef`（左/右ユニット）, `attTerrain/defTerrain`, `auto`（自動実行）, `logs`（戦闘ログ）。
- 入力（Intent）: `intentBack/intentRunOne/intentToggleAuto/intentSetTerrainAtt/intentSetTerrainDef`。
- 描画: `DrawBattleWithTerrain`（内部）、`ui/widgets` のボタン群を併用。
- ポップアップ: `LogView`（`popup_log.go`）に分離、Confirmで閉じる。
- ロジック: `engine.go` に簡易戦闘 `SimulateBattleCopy(WithTerrain)` を配置（将来 service へ抽出予定）。

フォント: Ebiten examples の M+ 1p Regular を OpenType で初期化（失敗時は basicfont にフォールバック）。

## pkg/game（ロジック層・UI非依存）
- 型
  - `type Stats { HP, Str, Skl, Spd, Lck, Def, Res, Mov int }`
  - `type Weapon { MT, Hit, Crit, Wt, RMin, RMax int; Type string }`
  - `type Unit { ID, Name, Class string; Lv int; S Stats; W Weapon }`
  - `type Terrain { Avoid, Def, Hit, Heal int }`（MVP: Avoid/Def/Hit を使用）
  - `type ForecastResult { HitDisp, Dmg, Crit int }`
- 三すくみ
  - 暫定: `Sword > Axe > Lance > Sword`（命中±10/威力±2）。
- API
  - `Forecast(att, def Unit) ForecastResult`
    - 地形なしの互換API。2RNは表示値のみ（判定はResolve系）。
  - `ForecastAt(att, def Unit, attTile, defTile Terrain) ForecastResult`
    - 仕様: `atk_hit = hit + skl*2 + floor(lck/2) + attTile.Hit`、`def_avo = spd*2 + lck + defTile.Avoid`、
      `hit_disp = clamp(atk_hit - def_avo + triHit, 0..100)`、`dmg = max(1, str+mt+triMt - (def + defTile.Def))`。
  - `ResolveRound(att, def Unit, rng *rand.Rand) (Unit, Unit, string)`
    - 地形なしの互換API。2RN・最小ダメ1・HP下限0。
  - `ResolveRoundAt(att, def Unit, attTile, defTile Terrain, rng *rand.Rand) (Unit, Unit, string)`
    - 予測値に基づき命中/クリティカル計算（0..100でクランプ）。

### pkg/game/input（入力・UI非依存）
- 型/定数
  - `type Action int`（`ActionUnknown` を 0。`ActionUp/Down/Left/Right/Confirm/Cancel/...`）
  - `type ControlState struct`（`Set/Get/Equal`）
  - `type Reader interface { Press(Action) bool; Down(Action) bool }`
  - `type EdgeReader struct`（`Step(ControlState)`）
  - `type EventKind int`, `type Event struct { Kind, Code, Value, Mods }`, `type Modifier`
  - `type Source interface { Poll() ControlState; Events() []Event }`
  - `type Pointer interface { Position() (x, y int) }`
  - `type Layout struct { Keyboard, Mouse map[int]Action }`, `func DefaultLayout() Layout`

### internal/game/provider/input/ebiten（アダプタ）
- `type Source struct`, `func NewSource(layout input.Layout) *Source`
- 実装: `Poll()`（Ebiten キー/マウス→ControlState）, `Position()`（マウス座標）

## cmd/ui_sample（エントリ）
- `const screenW = 1920`, `const screenH = 1080`: 論理解像度。
- `type Game`
  - `showHelp bool`: ヘルプ表示フラグ。
  - `unit ui.Unit`: 表示対象ユニット。
- `func NewGame() *Game`: ゲーム状態を初期化。
- `func (*Game) Update() error`: 入力処理（`H`/`Esc`/`Backspace`）。
- `func (*Game) Draw(*ebiten.Image)`: 画面描画（UI呼び出し）。
- `func (*Game) Layout(int,int) (int,int)`: 論理解像度を返す。

## UI メトリクス（uicore）
- 読込: `internal/config/uimetrics` にて JSON をロード（ユーザ→マスタ→既定）。
- 適用: `internal/game/service/ui.ApplyMetrics` で `uicore` の変数へ反映。
- 代表キー（list.*）例:
  - `margin`, `itemH`, `itemGap`, `titleOffset`, `portraitSize`
  - 固定オフセット: `headerTopGap`, `itemsTopGap`, `panelInnerPaddingX`, `titleXOffset`, `headerBaseX`,
    `rowTextOffsetX`, `rowTextOffsetY`, `rowBorderPad`, `rowRightIconSize`, `rowRightIconGap`
  - 列配列: `headerColumnsItems`, `headerColumnsWeapons`, `rowColumnsItems`, `rowColumnsWeapons`
 - 代表キー（status.*）例:
   - `panelPad`, `portraitSize`, `textGapX`, `nameOffsetY`
   - `classGapFromName`, `levelGapFromName`, `hpGapFromName`, `hpBarGapFromName`, `hpBarW`, `hpBarH`
   - `statsTopGap`, `statsLineH`, `statsColGap`, `weaponRanksXExtra`, `rankLineH`, `magicRanksTopExtra`
   - `equipTitleGapY`, `equipLineH`, `equipRect{W,H,YOffset}`, `equipLabelGapX`, `equipUsesX`
 - 代表キー（sim.*）例:
   - `startBtnW`, `startBtnH`, `autoRunGap`, `titleYOffset`, `titleXOffsetFromCenter`
   - `terrain.*`: `buttonW/H`, `baseYFromBottom`, `leftBaseXOffset`, `rightBaseXInset`, `buttonGap`, `labelLeftXOffset`, `labelYOffsetFromBottom`
- `func main()`: ウィンドウ作成とゲームループ開始。

## 参考
- 全体設計と移行計画は `docs/ARCHITECTURE.md` を参照。

## 変更時の指針
- 新規/変更した公開要素（型/関数/定数/フィールド）は本ファイルを必ず更新。
- UIレイアウト変更時は、座標・行間・フォントサイズの意図を1行で追記。
- 画像/フォント追加時は、読み込み先・依存・ライセンス（必要なら）を明記。

---

## internal/game/ui/input（抽象入力API）

- 目的: Scene から「入力の意味」だけを参照するための最小API。
- 構成:
  - `type Action int`: 抽象アクション列挙（`Confirm/Cancel/Menu/...` など）。
  - `type Reader interface { Press(Action) bool; Down(Action) bool }`: フレームスナップショットの読み取り。
  - `type ServiceAdapter struct{ S *service.Input }` / `func WrapService(*service.Input) Reader`: 既存実装を適合。
- 列挙（代表）:
  - 基本: `Up, Down, Left, Right, Confirm, Cancel, Menu, Next, Prev`
  - 便宜: `OpenWeapons, OpenItems, EquipToggle, Slot1..5, Unassign`
  - 地形: `TerrainAtt1..3, TerrainDef1..3`
- 利用例:
```
// ctx.Input は uinput.Reader
if ctx.Input.Press(uinput.Confirm) { /* 決定 */ }
if ctx.Input.Down(uinput.Menu)     { /* 長押し検出等 */ }
```
- 備考:
  - `game.Ctx.Input` は `uinput.Reader`。アプリ層で `WrapService(gamesvc.NewInput())` を供給。
  - 取得/マッピング（`Snapshot` 等）はアプリ層で継続実装し、UI からは隠蔽。
