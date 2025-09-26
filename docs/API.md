# API ドキュメント（Markdown）

本ドキュメントは、このサンプルの型・関数・定数の振る舞いを簡潔に説明します。実装コメントは最小限とし、詳細は本ファイルを更新してください。

## internal/game
- `const LevelUpExp = 100`: レベルアップに必要な経験値。
- `const LevelCap = 20`: レベル上限。

## internal/model（マスタデータ / mst_）
- `type Character`
  - 初期値のみ保持（`Name`, `Class`, `Portrait`, `Stats`, `Growth`, `Weapon`, `Magic`, `Equip(max)` など）。
- `func LoadFromJSON(path string) (*Table, error)`: JSON 読込。
- `func (*Table) Find(id string) (Character, bool)`: 取得。
備考: マスタは初期値のみ。レベルごとの能力は保持しない（成長率に依存し可変）。ファイル名は `mst_*.json`。パスは `db/master/`。

- `type ClassCaps`, `LoadClassCapsJSON`: クラス能力上限（`mst_class_caps.json`）
- `type Weapon`, `LoadWeaponsJSON`: 武器性能（`mst_weapons.json`）
- `type Item`, `LoadItemsJSON`: アイテム性能（`mst_items.json`）

## internal/user（ユーザデータ / usr_）
- `type Character`
  - 現在値を保持（`Level`, `Exp`, `HP`, `HPMax`, `Stats`, `Growth`, `Weapon`, `Magic`, `Equip(uses/max)` など）。
- `func LoadFromJSON(path string) (*Table, error)`: JSON 読込。
- `func (*Table) Find(id string) (Character, bool)`: 取得。
備考: ファイル名は `usr_*.json`。パスは `db/user/`。

## internal/ui（UI描画とデータモデル）
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
- `func SampleUnit() Unit`
  - ユーザテーブル（`db/user/usr_*.json`）のみから読み込み、UI用データに変換。
- `func DrawStatus(dst *ebiten.Image, u Unit)`
  - 1920×1080 を想定。左にポートレート、中央に基本情報/HP/能力値+成長率、右に「武器レベル」「魔法レベル」、下部に装備（耐久）を描画。
- 戦闘関連:
  - `DrawBattle(atk, def)`: 簡易プレビューを描画
  - `RollLevelUp/ApplyGains`: 成長抽選/反映（クラス上限でクランプ）
- ヘルパー（内部利用）
  - `drawPanel`: 影付きパネル。
  - `drawFramedRect`: 金縁の矩形。
  - `drawPortraitPlaceholder`: ポートレート未設定時の表示。
  - `drawPortrait`: 等比・線形補間で画像を枠内に表示。
  - `drawHPBar`: HP割合で色が変化するバー。
  - `drawStatLine`, `drawStatLineWithGrowth`: 能力値（成長率付き）。
  - `drawRankLine`: ランク表示。

フォント: Ebiten examples の M+ 1p Regular を OpenType で初期化（失敗時は basicfont にフォールバック）。

## cmd/ui_sample（エントリ）
- `const screenW = 1920`, `const screenH = 1080`: 論理解像度。
- `type Game`
  - `showHelp bool`: ヘルプ表示フラグ。
  - `unit ui.Unit`: 表示対象ユニット。
- `func NewGame() *Game`: ゲーム状態を初期化。
- `func (*Game) Update() error`: 入力処理（`H`/`Esc`/`Backspace`）。
- `func (*Game) Draw(*ebiten.Image)`: 画面描画（UI呼び出し）。
- `func (*Game) Layout(int,int) (int,int)`: 論理解像度を返す。
- `func main()`: ウィンドウ作成とゲームループ開始。

## 変更時の指針
- 新規/変更した公開要素（型/関数/定数/フィールド）は本ファイルを必ず更新。
- UIレイアウト変更時は、座標・行間・フォントサイズの意図を1行で追記。
- 画像/フォント追加時は、読み込み先・依存・ライセンス（必要なら）を明記。
