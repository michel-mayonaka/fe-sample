# ガイド/インデックス/AGENTS の更新と最終チェック

ステータス: [未着手]
担当: @tkg-engineer
開始: 2025-11-20 01:54:51 +0900

## 目的
- 新しい docs 構成を前提に、AGENTS・docs の各 README・ops-overview などガイド類を更新し、開発者と Codex が迷わず参照できる状態にする。

## 完了条件（DoD）
- [ ] `docs/ops-overview.md` に最新の構成・フローが反映されている。
- [ ] `AGENTS.md` や `docs/specs/README.md` など、docs 参照を含むガイドが新構成ベースで更新されている。
- [ ] `make mcp` が成功し、lint/テスト/ビルドに問題がない。
- [ ] 必要に応じて `make discovery-index` や `make story-index` を実行し、メタ情報の整合性が保たれている。

## 作業手順（概略）
- docs 構成変更の影響を受けるガイド・エントリポイント（AGENTS, README, CODEX_CLOUD など）を洗い出す。
- 新構成への導線（どこを見れば何が分かるか）を ops-overview などに明記する。
- 各種インデックスやメタ情報を更新し、CI が通ることを確認する。

## 進捗ログ
- 2025-11-20 01:54:51 +0900: タスク作成。

## 依存／ブロッカー
- Task01/Task02 の完了（構成とファイル移行が前提）。

## 成果物リンク
- 更新された `docs/ops-overview.md`, `docs/specs/README.md`, `AGENTS.md` など。

