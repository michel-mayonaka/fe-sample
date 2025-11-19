# ops-overview — 開発/ドキュメント全体の 1 枚絵

このファイルは「どの情報をどこで探せばよいか」を 1 ページで把握できることを目的としています。詳細は各リンク先を参照してください。

## 1. ストーリー駆動フロー
1. 課題・アイデアを Discovery (`stories/discovery/`) に起票し、`make discovery-index` で整理する。
2. 受け入れ済み Backlog からストーリーを選定し、`make new-story` で骨子を作成する。
3. タスクを 1 ファイル 1 作業で分割しつつ、`stories/YYYYMMDD-slug/tasks/*.md` に DoD を明記する。
4. 作業開始指示後に実装へ着手し、`make mcp` を通してから PR/レビューへ進む。
5. 完了したストーリーは `make finish-story` で `stories/finish/` にアーカイブする。
   - 詳細: `docs/KNOWLEDGE/workflows/stories.md`

## 2. ドキュメントマップ
| テーマ | 入口 | 主な内容 |
| --- | --- | --- |
| アーキテクチャ | `docs/architecture/README.md` | レイヤ構成、依存原則、Scene/Usecase の責務、`architecture/adr/` の決定履歴 |
| 仕様ハブ | `docs/SPECS/README.md` / `docs/SPECS/AGENTS.md` | world / gameplay / ui / reference 各カテゴリとテンプレート、仕様の読み方 |
| ナレッジ/規約 | `docs/KNOWLEDGE/README.md` | engineering（命名/コメント）、ops（Codex/オフライン）、data（DB メモ）、workflows、meta（本ファイル含む） |
| Codex 連携 | `AGENTS.md` / `docs/KNOWLEDGE/ops/ai-operations.md` | エージェント運用ルール、Codex Cloud 手順、AI 連携の前提 |
| Ops & テスト | `docs/KNOWLEDGE/workflows/local-dev.md`, `docs/KNOWLEDGE/workflows/ci.md` | ローカル検証 (`make smoke`, `make mcp`)、CI 構成と厳格 UI ビルド |

## 3. コード/仕様対応表
| 層 | コード位置 | 仕様/ドキュメント |
| --- | --- | --- |
| UI/Scene | `internal/game/scenes`, `internal/game/ui/*` | `docs/SPECS/ui/*.md`, ADR-0001（Scene/Usecase 分離） |
| Usecase | `internal/usecase/*` | `docs/SPECS/gameplay/*.md`, `docs/SPECS/reference/api.md` |
| Repo/Infra | `internal/infra/*`, `internal/repo/*` | `docs/KNOWLEDGE/data/db-notes.md` |
| モデル/データ | `internal/model`, `db/*` | `docs/SPECS/world/`（今後拡充）, `docs/KNOWLEDGE/data/db-notes.md` |

## 4. テストと CI の導線
- 日常検証: `make smoke`（最短）, `make mcp`（vet + lint + unit）。
- UI 厳格検証: `MCP_STRICT=1 make check-ui`（Linux では依存導入が必須）。
- CI: `smoke-offline` → `build-and-lint` → `ui-build-strict`。詳しくは `docs/KNOWLEDGE/workflows/ci.md`。
- Story/Discovery のメタ整合は `make validate-meta`。CI では `build-and-lint` で警告ログ化。

## 5. Codex 連携とナレッジ反映
- 作業開始前に `AGENTS.md` のルールとストーリー DoD を確認する。
- 仕様追加・変更時は `docs/SPECS/` を先に更新し、必要に応じて `docs/KNOWLEDGE/meta/docs-structure.md` へ方針を追記する。
- Codex Cloud や AI 補助を使う場合は `docs/KNOWLEDGE/ops/codex-cloud.md` を参照し、オフラインモードは `docs/KNOWLEDGE/ops/offline.md`。

## 6. 更新ポリシー
- 本ファイルはストーリー開始時の導線確認に用いる。新カテゴリを追加する場合は `docs/KNOWLEDGE/meta/docs-structure.md` のポリシー更新とセットで改訂する。
- `rg "docs/" docs -n` で旧パスが残っていないかを確認し、リンク切れを防止する。
