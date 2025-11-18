# 命名規約（NAMING）

ステータス: Draft（詳細化）
更新日: 2025-09-29

## 目的・適用範囲
- 本ドキュメントは、リポジトリ内の命名に関する統一基準を定めます。
- 適用対象を包括的にカバーします（詳細はチェックリスト参照）。

## 網羅チェックリスト
- [ ] ファイル名 / ディレクトリ名
- [ ] パッケージ名（Go）
- [ ] 型: 構造体 / インタフェース（クラス相当）
- [ ] 関数 / メソッド / レシーバ
- [ ] 変数 / フィールド / 定数
- [ ] エラー（`ErrXxx` / ラップ方針）
- [ ] テスト名（`*_test.go` / `TestXxx`）
- [ ] DB: テーブル / カラム（`mst_` / `usr_` / `snake_case`）
- [ ] JSON: ルートキー / フィールドキー（`snake_case`）
- [ ] アセット（画像/音源/フォント）
- [ ] スクリプト / Make ターゲット
- [ ] コミット / PR タイトル（Conventional Commits 整合）

## 原則
- 一貫性: 既存規約と衝突しない（README / ARCHITECTURE / COMMENT_STYLE に準拠）。
- 可読性: 省略は一般的な初期ismのみ（例: `ID`, `HTTP`, `URL`）。
- 正規化: コード（CamelCase）とデータ（snake_case）を明確に分離。
- 互換性: データ移行を考慮（JSON→SQLite 予定）。

## ファイル・ディレクトリ
- 形式: 小文字スネークを基本。区切りは `_`、拡張子前の環境/ビルドタグは `_{tag}` を末尾に付与。
- コードファイル: `status_panel.go`, `button_test.go`, `inventory_port.go`。
- ディレクトリ（コード）: パッケージ名に一致（例: `ui`, `gamecore`）。複合語は省略せず明確に。
- ストーリー: `stories/YYYYMMDD-slug/`（例: `20250929-docs-naming`）。
- 禁止/注意:
  - NG: `StatusPanel.go`（大文字混在）, `statusPanel.go`（camel混在）, `status-panel.go`（`-` 使用）。
  - NG: 意味のない略語（`cfgs`, `misc`）。OK: `cfg`（設定の通称として許容、プロジェクト既存踏襲）。

## Go コード
- パッケージ: 小文字単語連結（例: `ui`, `gamecore`）。「役割+領域」を意識（例: `usecase`, `provider`）。
- 型名: UpperCamelCase。インタフェースは「能力名+er」が基本（例: `Renderer`）。架橋層は既存方針に合わせ `DataPort`, `BattlePort`, `InventoryPort`。
- 関数/メソッド: 動詞から開始（例: `LoadUnits`, `RunBattleRound`）。副作用のない取得は `Get` を避け名詞で（例: `Unit()` より `Unit` フィールド/プロパティ、または `FindUnit`）。
- パラメータ: `ctx context.Context` は先頭に `ctx`。時間は `now time.Time`。数は `n`, `count`。インデックスは `i`, `idx`。
- 初期ism: `ID`, `HTTP`, `URL` を全大文字で維持（例: `UserID`, `HTTPServer`）。
- 変数/フィールド: 文脈重視で短め（例: `hp`, `exp`）。否定のブールは避ける（NG: `disable`、OK: `enabled`）。
- 定数: UpperCamelCase。列挙は型付き `iota` を用い、接頭辞でグループ化（例: `ElementFire`/`ElementWater`）。
- エラー: パッケージ変数は `ErrXxx`。文脈付与は `fmt.Errorf("...: %w", err)`、複数は `errors.Join`。
- レシーバ: 1〜2 文字（型名の頭字、例: `u *Unit`）。ミュータブルである旨は命名に含めない。
- テスト: `TestXxx`。サブテストは `t.Run("case/条件", ...)` でスラッシュ区切り。

良い例 / 悪い例:
```go
// OK: インタフェースは能力名+er
type Renderer interface { Render() error }

// OK: Port命名
type DataPort interface { PersistUnit(ctx context.Context, u Unit) error }

// NG: I を付けない（C#風の `IUnit` は不可）
type IRenderer interface { Render() error } // NG

// OK: 初期ism
func (a *App) FindByID(ctx context.Context, id string) (*Unit, error) { /* ... */ return nil, nil }

// NG: 否定bool
func NewOverlay(disable bool) { /* ... */ } // NG -> NewOverlay(enabled bool)
```

## データ（DB / JSON）
- テーブル接頭辞: マスタ `mst_*`、ユーザ `usr_*` を厳守。
- テーブル名: 複数形を基本（例: `mst_characters`, `usr_characters`）。
- カラム名: `snake_case`。主キーは `id`、FK は `xxx_id`（例: `weapon_id`）。
- タイムスタンプ: `created_at`, `updated_at`。論理削除は `deleted_at`（NULL 非削除）。
- 真偽値: `is_` または `has_`（例: `is_active`）。
- JSONキー: `snake_case` を維持し、Go 側はタグでマッピング（`json:"user_id"`）。
- 互換対策: 既存 JSON キー改名は当面避ける。必要時は旧キーも許容するローダで移行（読み込み時に別名対応）。

良い例 / 悪い例:
```sql
-- OK
CREATE TABLE mst_characters (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  hp INTEGER NOT NULL,
  exp INTEGER NOT NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

-- NG: キャメル/省略・接頭辞なし
CREATE TABLE Characters (CharacterId TEXT PRIMARY KEY, CreatedAt TEXT);
```

## アセット
- ディレクトリ: 種別別に分離（`assets/ui/`, `assets/sfx/`, `assets/fonts/`）。
- 命名順序: `用途_部位_状態_バリアント@倍率.拡張子`
  - 例: `btn_primary_idle@2x.png`, `icon_sword_dark@3x.png`
  - 状態（例: `idle`, `hover`, `active`）/ バリアント（例: `dark`, `disabled`）は必要時のみ。
- 解像度サフィックス: `@2x`, `@3x` を採用（最終決定）。
  - 理由: 一般的な高DPI表記との互換性が高く、ツール連携（デザイナ資産）での認知度が高い。
  - 複合例: `status_panel_dark@2x.png`（snake_case と `@` の併用を許容）。
- スプライトシート: `*_atlas.png` + `*.json`（例: `battle_effect_atlas.png`）。

## スクリプト / Make
- スクリプト: 小文字スネーク（例: `new_story.sh`, `finish_story.sh`）。動詞から開始。
- Make ターゲット: 既存踏襲（`check-all`, `test-all`, `mcp`）。動作を端的に表す短語。
- 作業単位: ストーリー関連は `new-story` などハイフン区切り可（既存互換に注意）。

## 汎用名の避け方（helpers/util など）
- 禁止/非推奨の例: `util`, `utils`, `helper(s)`, `common`, `misc`, `core`, `shared`。
  - 理由: 役割が曖昧になり、探索性（grep/Go to Definition）が低下するため。
- 置き換え指針（ドメイン/目的語で具体化）:
  - 乱数ヘルパ → `rng`（例: `internal/game/rng`）
  - 幾何矩形判定 → `rect`（例: `rect.go` へ集約。暫定的に複数機能が混在する場合は `rect_helpers.go` とし、後で分割）
  - 画像処理 → `imageproc` / `imaging`
  - 永続化補助 → `loader`, `saver`, `migrator`
  - UI 補助 → `ui/layout`（座標）, `ui/draw`（描画）, `ui/view`（表示用型）, `ui/adapter`（変換）
  - 変換層 → `adapter`（例: `adapter/ui_to_game.go`）
  - 乱用されやすい `Manager`/`Data`/`Info` は具体語に置換（例: `InventoryRepo`, `BattlePort`）。
- 判断フロー:
  1) 単一概念に帰属するか → その概念名（名詞/役割）で命名。
  2) 複数概念が混在 → ファイル/サブパッケージへ分割。
  3) 一時集約が必要 → ドメイン接頭辞 + `_helpers` を暫定許可（要TODOで分割計画）。
- 本リポジトリの例:
  - 変更: `internal/game/util` → `internal/game/rng`（RNG 専用に特化）。
  - 目的特化: `helpers.go` → `rect_helpers.go`（矩形関連へ限定。将来 `rect.go` 等へ再集約）。

### ケーススタディ: `internal/game/util` 撤去（2025-11）
- 2025-09-29 の暫定改称で RNG ラッパを `internal/game/rng` へ切り出し、そのまま util ディレクトリは空のまま残置していた。
- 2025-11-19 ストーリー `20251119-internal-game-util-specialize` で最終確認し、ディレクトリを完全削除。`math/rand` の利用箇所は段階的に `rng` サブパッケージへ寄せる方針を残し、空の util や曖昧名を再導入しないことを明示した。
- 以降、新たな共有ヘルパーが必要になった場合は Discovery → Story を経て、`rng`/`rect`/`debug` 等のように目的語で命名したサブパッケージを起こすこと。暫定集約でも `*_helpers.go` のようにスコープを限定し、完了後すぐに分割計画を残す。

## コミット / PR
- 形式: Conventional Commits。種別例 `feat`, `fix`, `docs`, `refactor`, `test`, `chore`。
- 例（コミット）:
  - `docs(naming): NAMING に Go/DB の詳細規則を追記`
  - `refactor(ui): HP 表示ユーティリティの命名をガイドに整合`
- 例（PR タイトル）:
  - `docs: 命名規約の詳細化（Go/DB/アセット）`
  - `refactor: Renderer 実装のリネームと影響範囲の是正`
- 本文: 目的/影響範囲/検証手順を箇条書きで簡潔に。

## 例（最小）
```go
// Package 名: ui
package ui

// Unit はゲーム内のユニットを表す。
// HP/守備などの表示は View 側へ委譲。
type Unit struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    HP     int    `json:"hp"`
    Exp    int    `json:"exp"`
}

// ReloadData はマスタ/ユーザデータを再読込する。
func (a *App) ReloadData() error { /* ... */ return nil }

var ErrNotFound = errors.New("unit not found")
```

### 列挙定数の細則（決定）
- 型付き `iota` を用い、型名を接頭辞にした UpperCamelCase を採用。
- 先頭に未定義を示す `XxxUnknown` を配置（ゼロ値安全）。
- 文字列表現が必要な場合は `String()` 実装、または `stringer` 等の生成を利用（任意）。
- ALL_CAPS は使用しない。

良い例 / 悪い例:
```go
type Element int

const (
    ElementUnknown Element = iota
    ElementFire
    ElementWater
    ElementWind
)

// OK: UpperCamel + 型名プレフィックス、ゼロ値は Unknown

// NG: ALL_CAPS / プレフィックス無し / 型未指定の iota
const (
    FIRE = iota // NG
    Water       // NG: プレフィックス無し
)
```

## 未決事項（TBD）
- ブランチ命名規則（必要なら別途策定）。

## 関連
- `docs/ARCHITECTURE.md`（命名・スタイル補足）
- `docs/COMMENT_STYLE.md`（GoDoc 記法）
- `docs/DB_NOTES.md`（移行方針）
