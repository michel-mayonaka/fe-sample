# KNOWLEDGE — ナレッジと規約の入口

`docs/KNOWLEDGE/` は、仕様そのものではないが日々の開発で参照するルール・運用・メモをカテゴリ別に集約します。まずはここから必要なサブディレクトリへ辿ってください。

## カテゴリ一覧
| カテゴリ | パス | 主な内容 |
| --- | --- | --- |
| engineering | `docs/KNOWLEDGE/engineering/` | 命名規約（`naming.md`）、コメントスタイル（`comment-style.md`）など、コード品質に直結するルール |
| data | `docs/KNOWLEDGE/data/` | `db-notes.md`（マスタ/ユーザデータの扱い、将来の DB 移行方針） |
| ops | `docs/KNOWLEDGE/ops/` | Codex/Vibe 連携（`ai-operations.md`）、Codex Cloud 手順（`codex-cloud.md`）、オフライン開発（`offline.md`）、移行ログ（`migrations/`） |
| workflows | `docs/KNOWLEDGE/workflows/` | ストーリー運用、ローカル検証、CI、Vibe-kanban などの具体的な手順 |
| meta | `docs/KNOWLEDGE/meta/` | `docs-structure.md` など、ドキュメントポリシーや索引に関するメタ情報 |

## 関連エントリ
- 開発全体の 1 枚絵: `docs/ops-overview.md`
- 仕様ハブ（SPECS ディレクトリ）の入口: `docs/SPECS/README.md`
- Codex/開発者共通の最初の読み物: `AGENTS.md`

## 更新時の注意
1. 新しいカテゴリやドキュメントを追加する場合は、本 README と `docs/KNOWLEDGE/meta/docs-structure.md` を更新し、参照経路を明確にする。
2. 運用ルールの変更はストーリーで追跡し、DoD に「AGENTS/ops-overview/関連ナレッジ更新」を含める。
3. 既存ルールと衝突する場合は理由を記し、必要なら `stories/discovery/` にフォローアップを起票する。
