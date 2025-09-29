# 02_update_architecture_docs — ARCHITECTURE 更新

ステータス: [完了]

## 目的
- 監査結果を反映し、`docs/ARCHITECTURE.md` を現行に整合させる。

## 作業概要
- 境界・依存の表現を更新（Scene/Service/Provider/Usecase/Model/DataPort）。
- ライフサイクル節を新設・更新（`input→intent→advance→render`）。
- 代表的な画面フロー図（ASCII 可）を追加。
- 関連参照の整合（`README.md`, `docs/NAMING.md`）。
- Provider の拡張を明記（`UserWeapons()`/`UserItems()` を追加し、在庫参照は Provider 統一）。

## 手順（チェックリスト）
- [ ] `artifacts/audit_findings.md` をインプットに更新差分を列挙
- [ ] 文言・用語を NAMING に合わせて統一
- [ ] 例示コード/パスを最新に更新
- [ ] リンク検証（相対パス）

## 成果物
- `docs/ARCHITECTURE.md` 更新 PR（ブランチ/差分）

## DoD
- [x] PR 上でレビュー観点（境界/依存/命名/ライフサイクル）がクリア（ローカル反映済み、レビュー待ち）
- [x] `make mcp` グリーン（check-all OK）
