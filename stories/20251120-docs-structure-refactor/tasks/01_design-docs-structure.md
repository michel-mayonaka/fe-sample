# docs 構成のターゲット設計とマッピング案作成

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-20 01:54:51 +0900

## 目的
- docs/ の現状ファイル一覧と役割を棚卸しし、ターゲット構成（architecture/・KNOWLEDGE/・SPECS/ など）へのマッピング案を作る。
- ops-overview（1枚絵）に載せるべき観点（アーキ/テスト戦略/ストーリー運用/Codex 連携）を整理する。

## 完了条件（DoD）
- [ ] 現在の docs/ 配下の主要ファイルが一覧化され、用途（アーキ/ナレッジ/仕様/運用）がタグ付けされている。
- [ ] 提示されたターゲット構成（ディレクトリツリー）が story 本文のイメージを満たしている。
- [ ] 既存 docs をどのディレクトリへ移すかのマッピング表（簡易でよい）が作成されている。

## 作業手順（概略）
- docs/ 配下を一覧し、現状の分類と役割を確認する。
- 提案されたツリー構成（architecture/・KNOWLEDGE/・SPECS/）をベースに、必要なサブディレクトリや命名ルールを補足する。
- 既存ファイルをターゲット構成にマッピングし、迷うものはメモを残す。

## 進捗ログ
- 2025-11-20 01:54:51 +0900: タスク作成。
- 2025-11-19 17:16:00 +0900: docs 現状棚卸しとターゲット構成案を作成。
- 2025-11-19 17:35:00 +0900: 棚卸し表／ターゲットツリー／マッピング表を更新し、DoD を満たしたため設計フェーズ完了。

## 現状棚卸（2025-11-19 時点）
| ドキュメント | 種別タグ | 主な用途/備考 |
| --- | --- | --- |
| docs/architecture/README.md | architecture | UI サンプル全体のレイヤ構成と依存ルール |
| docs/SPECS/reference/api.md | spec/ref | 型や公開 API 梗概（ゲーム内ユースケースの参照元） |
| docs/KNOWLEDGE/engineering/naming.md | knowledge/engineering | 命名規約（コード/データ/アセット） |
| docs/KNOWLEDGE/engineering/comment-style.md | knowledge/engineering | コメント/GoDoc の記法 |
| docs/KNOWLEDGE/data/db-notes.md | knowledge/data | DB/永続化のメモ |
| docs/KNOWLEDGE/ops/ai-operations.md | knowledge/ops | Codex/Vibe 運用メモ |
| docs/KNOWLEDGE/ops/codex-cloud.md | knowledge/ops | Codex Cloud の実行手順 |
| docs/MIGRATION_20251013_CODEX_CLOUD.md | knowledge/ops/migration | Codex Cloud 移行ログ |
| docs/KNOWLEDGE/ops/offline.md | knowledge/ops | オフライン開発プロファイル |
| docs/KNOWLEDGE/meta/docs-structure.md | knowledge/meta | docs 配下の方針と入口整理 |
| docs/SPECS/README.md | specs | 仕様ハブ概要（system/ui の2層） |
| docs/SPECS/AGENTS.md | specs/agents | 仕様の読み方ガイド |
| docs/SPECS/templates/gameplay.md | specs/template | システム仕様テンプレート |
| docs/SPECS/gameplay/battle_basic.md | specs/system | 戦闘基礎シナリオ |
| docs/SPECS/templates/ui.md | specs/template | 画面仕様テンプレート |
| docs/SPECS/ui/status_screen.md | specs/ui | ステータス画面仕様 |
| docs/KNOWLEDGE/workflows/*.md | knowledge/workflow | ストーリー/CI/ローカル開発/ボード運用の手順 |

## ターゲット構成案

```
docs/
  ops-overview.md          # ストーリー駆動/テスト/ナレッジ全体を俯瞰する 1 枚絵
  architecture/
    README.md              # 旧 ARCHITECTURE を移動し、構成方針と入口を明記
    adr/
      ADR-0001-scene-architecture.md  # 代表的 ADR（Scene/Usecase 分離）
  KNOWLEDGE/
    README.md              # 種別タグ・参照優先度
    engineering/
      naming.md
      comment-style.md
    data/
      db-notes.md
    ops/
      ai-operations.md
      codex-cloud.md
      offline.md
      migrations/20251013-codex-cloud.md
    workflows/
      overview.md
      stories.md
      ci.md
      local-dev.md
      vibe-kanban.md
    meta/
      docs-structure.md
  SPECS/
    README.md
    AGENTS.md
    templates/
      gameplay.md
      ui.md
    world/
      (将来: ロア/設定)
    gameplay/
      battle_basic.md
    ui/
      status_screen.md
    reference/
      api.md
```

- `ops-overview.md` に「タスク/ナレッジ/アーキ/テスト戦略」の入口を集約する。
- architecture/README はディレクトリ構成、その下の ADR で意思決定ログを管理。
- KNOWLEDGE はタグ（engineering/data/ops/workflows/meta）で管理し、AGENTS から直接リンク。
- SPECS はカテゴリを world/gameplay/ui/reference に再編し、テンプレート類は `templates/` に集約する。

## 既存 docs → 新構成マッピング
| 旧パス | 新パス | メモ |
| --- | --- | --- |
| docs/architecture/README.md | docs/architecture/README.md | 章立てを維持しつつ README 化 |
| docs/SPECS/reference/api.md | docs/SPECS/reference/api.md | 仕様参照として reference 配下に再配置 |
| docs/KNOWLEDGE/engineering/naming.md | docs/KNOWLEDGE/engineering/naming.md | 命名規約 - engineering カテゴリ |
| docs/KNOWLEDGE/engineering/comment-style.md | docs/KNOWLEDGE/engineering/comment-style.md | コメント規約 |
| docs/KNOWLEDGE/data/db-notes.md | docs/KNOWLEDGE/data/db-notes.md | 永続化メモ |
| docs/KNOWLEDGE/ops/ai-operations.md | docs/KNOWLEDGE/ops/ai-operations.md | Codex/Vibe 運用 |
| docs/KNOWLEDGE/ops/codex-cloud.md | docs/KNOWLEDGE/ops/codex-cloud.md | Codex Cloud 手順 |
| docs/MIGRATION_20251013_CODEX_CLOUD.md | docs/KNOWLEDGE/ops/migrations/20251013-codex-cloud.md | 過去移行ログ |
| docs/KNOWLEDGE/ops/offline.md | docs/KNOWLEDGE/ops/offline.md | オフライン開発 |
| docs/KNOWLEDGE/meta/docs-structure.md | docs/KNOWLEDGE/meta/docs-structure.md | ドキュメント構成ポリシー |
| docs/SPECS/README.md | docs/SPECS/README.md | 仕様ハブ README（新カテゴリ説明へ更新） |
| docs/SPECS/AGENTS.md | docs/SPECS/AGENTS.md | 仕様の読み方ガイド（参照優先度更新） |
| docs/SPECS/templates/gameplay.md | docs/SPECS/templates/gameplay.md | system→gameplay テンプレ内訳 |
| docs/SPECS/gameplay/battle_basic.md | docs/SPECS/gameplay/battle_basic.md | 戦闘仕様 |
| docs/SPECS/templates/ui.md | docs/SPECS/templates/ui.md | UI テンプレ |
| docs/SPECS/ui/status_screen.md | docs/SPECS/ui/status_screen.md | ステータス仕様 |
| docs/KNOWLEDGE/workflows/overview.md | docs/KNOWLEDGE/workflows/overview.md | ワークフロー一覧 |
| docs/KNOWLEDGE/workflows/stories.md | docs/KNOWLEDGE/workflows/stories.md | ストーリー運用 |
| docs/KNOWLEDGE/workflows/ci.md | docs/KNOWLEDGE/workflows/ci.md | CI 方針 |
| docs/KNOWLEDGE/workflows/local-dev.md | docs/KNOWLEDGE/workflows/local-dev.md | ローカル検証 |
| docs/KNOWLEDGE/workflows/vibe-kanban.md | docs/KNOWLEDGE/workflows/vibe-kanban.md | カンバン手順 |

未確定事項:
- `world/` 配下の初回コンテンツは空ディレクトリで用意し、今後のストーリーで埋める。
- `ops-overview.md` のテスト戦略セクションは Task03 で詳細化する（本タスクではアウトラインのみ作成予定）。

## 依存／ブロッカー
- 特になし（ただし stories/20251119-story-work-categories の方針と矛盾しないよう確認する）。

## 成果物リンク
- docs 構成案: `docs/ops-overview.md` の草案、または別メモ。
