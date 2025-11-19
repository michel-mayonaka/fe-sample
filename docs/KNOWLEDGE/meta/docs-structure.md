# DOCS_STRUCTURE — ドキュメント構成ポリシー

`docs/` 配下の役割と配置方針をまとめたガイドです。仕様/アーキ/ナレッジが混在しないよう、カテゴリと入口を明確に定義します。

## 1. 目的
- 開発者とエージェントが、参照すべき文書へ最短で辿り着けるようにする。
- 「設計思想」「仕様本文」「運用ルール」「ナレッジ」を切り分け、重複やリンク切れを防ぐ。

## 2. ルート構成（2025-11-20 時点）
```
docs/
  ops-overview.md        # ストーリー/ナレッジ/テスト/AI連携を俯瞰する 1 枚絵
  architecture/          # 本 README (理想アーキ) と ADR 群
  KNOWLEDGE/             # 規約・運用メモ・ワークフロー・メタ情報
  SPECS/                 # 実装に直結する仕様ハブ
```

### 2.1 architecture/
- `docs/architecture/README.md`: レイヤ構成・依存原則・導入ステップを定義（本番アーキの基準点）。
- `docs/architecture/adr/*.md`: 主要な意思決定ログ。更新時は README からリンクする。

### 2.2 KNOWLEDGE/
- `engineering/`: 命名（`naming.md`）、コメントスタイル（`comment-style.md`）など。
- `data/`: 永続化メモ（`db-notes.md`）。
- `ops/`: Codex/Vibe 運用（`ai-operations.md`）、Codex Cloud（`codex-cloud.md`）、オフライン（`offline.md`）、移行ログ。
- `workflows/`: ストーリー運用・ローカル検証・CI・vibe-kanban など。
- `meta/`: 本ファイルや索引ポリシー。
- 入口: `docs/KNOWLEDGE/README.md` でカテゴリ別に導線を案内。

### 2.3 SPECS/
- 目的: 実装に直結する仕様（Specification）を world / gameplay / ui / reference に整理し、テンプレートを共通化。
- サブディレクトリ:
  - `world/`: ロア・設定・マスターデータの背景。
  - `gameplay/`: ユースケースや戦闘ルールなどロジック中心の仕様。
  - `ui/`: 画面仕様。
  - `reference/`: API や汎用リファレンス。
  - `templates/`: gameplay/ui 共通テンプレート。
- 入口: `docs/SPECS/README.md`（構成とルール）、`docs/SPECS/AGENTS.md`（仕様の読み方）。

## 3. SPECS の書き方
- 仕様更新はコード変更より先に行う。`docs/SPECS/templates/*.md` をコピーしてメタ情報（状態/主な実装/最新ストーリー）を記入する。
- world → gameplay → ui の順に参照すると背景→ロジック→画面で文脈が揃う。
- 仕様に無い挙動を実装する場合は、ストーリーで仕様追加を伴う。

## 4. 入口と優先順位
1. `AGENTS.md`: リポジトリ全体のクイックリンクと運用ルール。
2. `docs/ops-overview.md`: ストーリー/ナレッジ/テスト導線の 1 枚絵。
3. `docs/SPECS/README.md` + `docs/SPECS/AGENTS.md`: 仕様本文の探し方。
4. `docs/architecture/README.md`: 理想アーキと依存原則。
5. `docs/KNOWLEDGE/README.md`: 規約・ワークフロー・運用メモ。

## 5. 運用ルール
- 仕様優先: 振る舞いを変える場合は `docs/SPECS/` の該当ファイルを更新し、DoD に spec 反映を含める。
- リンク整合: `rg "docs/" docs -n` で旧パスを検出し、構成変更後は `docs/ops-overview.md` と本ファイルを必ず更新する。
- ストーリー駆動: ドキュメント構成変更は Story を切って行い、タスクに `make mcp` とリンク確認を含める。

## 6. 今後の拡張
- world カテゴリの充実（勢力設定や装備体系）。
- `docs/SPECS/reference/` に API やテレメトリ仕様を追記。
- `docs/KNOWLEDGE/ops/migrations/` を用いた移行ログの整理。
- 目次/索引生成の自動化（`make docs-index` など）を検討。
