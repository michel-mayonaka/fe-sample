# AGENTS/README への運用ルール反映

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-17 00:35:20 +0900

## 目的
- 設計した specs 用メタデータと、その更新フロー（Story/Discovery/Backlog との関係）を既存ドキュメントに反映し、運用をチーム全体で共有できる状態にする。

## 完了条件（DoD）
- [x] `docs/SPECS/AGENTS.md` に、メタキーと状態の意味・更新タイミングが追記されている。
- [x] 必要に応じて `docs/SPECS/README.md` か `docs/KNOWLEDGE/workflows/stories.md` に、spec 状態と Story 運用の関係が1〜2段落で整理されている。
- [x] 代表 spec に付与したメタ情報と、ドキュメントに書かれたルールが矛盾していない。

## 作業手順（概略）
- タスク01/02で決まったメタキーと適用例を確認する。
- `docs/SPECS/AGENTS.md` に「spec の状態管理」の節を追加し、メタキーと状態遷移を簡潔に説明する。
- 必要であれば `docs/SPECS/README.md` や `docs/KNOWLEDGE/workflows/stories.md` に、spec 状態と Story/Backlog の役割分担を追記する。

## 進捗ログ
- 2025-11-17 00:35:20 +0900: タスク作成。
- 2025-11-17 00:38:30 +0900: docs/SPECS/AGENTS.md に「spec の状態メタデータ」節を追加し、メタキーの意味と更新タイミングを追記。docs/SPECS/README.md からもメタデータ運用への導線を追加。

## 依存／ブロッカー
- `stories/20251117-spec-status-metadata/tasks/01_spec_metadata_design.md`
- `stories/20251117-spec-status-metadata/tasks/02_apply_metadata_to_specs.md`

## 成果物リンク
- `docs/SPECS/AGENTS.md`
- `docs/SPECS/README.md`
- `docs/KNOWLEDGE/workflows/stories.md`
