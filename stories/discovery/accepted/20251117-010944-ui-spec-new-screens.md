# discovery: 画面仕様の新規作成（タイトル/編成/バトルマップ）

ステータス: [promoted]
担当: @kyosuke_tsubo
開始: 2025-11-17 01:09:44 +0900
優先度: P2
提案元: branch=master
関連ストーリー: N/A

## 目的
- 既存の `docs/specs/ui/status_screen.md` と同じ粒度で、以下 3 画面の仕様を ui spec として追加し、今後の UI 実装/リファクタリングの土台にする。
  - タイトル画面
  - 編成画面（ユニットの入れ替え/出撃メンバー選択）
  - バトルマップ画面（マス目調のシミュレーション画面）

## 背景
- 現状はステータス画面など一部のみ ui spec が存在し、タイトル/編成/バトルマップといったコア画面の仕様が明文化されていない。
- これらの画面は将来的な実装・リファクタリング・差分チェック（Codex workflow）において重要度が高く、先に仕様だけでも揃えておくと後続作業がやりやすくなる。

## DoD候補
- [ ] `docs/specs/ui/_TEMPLATE.md` をもとに、タイトル画面の ui spec が作成されている（例: `docs/specs/ui/title_screen.md`）。
- [ ] 同様に、編成画面の ui spec が作成されている（例: `docs/specs/ui/formation_screen.md`）。
- [ ] 同様に、バトルマップ画面の ui spec が作成されている（例: `docs/specs/ui/battle_map_screen.md`）。
- [ ] 各 spec に「状態/主な実装/最新ストーリー」のメタデータ欄が追加され、現時点では `spec-only` として整理されている。

## 関連
- docs/specs/ui/_TEMPLATE.md
- docs/specs/ui/status_screen.md

## 進捗ログ
- 2025-11-17 01:09:44 +0900: 起票

- 2025-11-17 01:10:24 +0900: Backlog へ昇格（優先度=P2）
