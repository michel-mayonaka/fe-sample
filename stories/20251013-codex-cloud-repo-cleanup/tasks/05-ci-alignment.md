# 05 CI 方針との整合

## 目的
既存 CI（`MCP_STRICT=0 make mcp`）と衝突せず、codex-cloud 向け設計と両立させる。

## 手順（ドラフト）
1) 既存ワークフローの参照ジョブ確認。
2) スモークチェックを別ジョブ/ステップとして提案。
3) UI 依存は別ジョブ（strict）に分離したまま維持。

## 成功条件（DoD）
- CI のフロー図（簡易）が README/Docs に反映される。

## 実装メモ（結果）
- GitHub Actions に `smoke-offline` ジョブを追加（`MCP_OFFLINE=1 make smoke`）。
- 既存 `build-and-lint` は `smoke-offline` 依存に変更。`ui-build-strict` は従来どおり後続で厳格検証。
- README に CI フロー（依存関係）を追記。
