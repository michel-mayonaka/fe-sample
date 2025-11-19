# DBメモ（将来のSQLite移行）

- 目的: JSON（マスタ/ユーザ）を SQLite に移行し、検索性・一貫性・トランザクションを向上。
- ドライバ候補:
  - `modernc.org/sqlite`（CGO不要、クロスプラットフォームに有利）
  - `github.com/mattn/go-sqlite3`（実績多数、要CGO）
- 想定スキーマ（概略）
  - `masters_characters(id PK, name, class, portrait, base_level, base_exp, hp_base, stats_json, growth_json, weapon_json, magic_json)`
  - `users_characters(id PK, level, exp, hp, hp_max)`
  - `users_equips(id, slot_idx, name, uses, max, PRIMARY KEY(id, slot_idx))`
- 合成方針: `masters_characters` をベースに `users_*` で上書き（LEFT JOIN）。
- マイグレーション: 初回起動時にテーブル作成→マスタJSONをロード→ユーザJSONをインポート。
- 将来タスク: DAO層（`internal/repo`）追加、トランザクション境界設計、自動保存/スナップショット。

現行 JSON バックエンドでの参照経路メモ:
- マスタ武器/アイテムは `internal/repo.WeaponsRepo` / `internal/repo.ItemsRepo` で読み込み・キャッシュされ、
  UI からは `internal/game/data.TableProvider`（実装: `internal/usecase.App`）の `WeaponsTable()/ItemsTable()` 経由で参照します。

## プラットフォーム別の扱い（デスクトップ / WebGL）

- デスクトップ（Windows/macOS 等）:
  - 現状: マスタ/ユーザともに JSON バックエンド（`db/master/*.json`, `db/user/*.json`）を `JSON*Repo` 経由で読み書き。
  - 将来方針: 上記スキーマに沿って SQLite へ移行し、`SQLiteUserRepo`/`SQLiteMasterRepo` のような実装を `internal/repo` に追加する。
- WebGL（wasm）:
  - 現状: ユーザデータの永続化は行わず、「デモ版」としてサンプルデータのみで起動する（ユーザRepo初期化に失敗しても panic せず、サンプルユニットで動作）。
  - 将来候補:
    - 簡易セーブが欲しくなった場合は、`BrowserUserRepo` のようなインターフェース実装を追加し、`localStorage`/`IndexedDB` 上に JSON を保存する案を検討する。
    - WebGL で SQLite を使う場合は、メモリDB＋ブラウザストレージへのスナップショット保存など、専用 Story/Discovery で別途検討する。
