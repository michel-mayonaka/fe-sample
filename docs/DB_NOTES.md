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
