# discovery: アーキテクチャ設計と実装の差分チェック workflow

ステータス: [promoted]
担当: @kyosuke_tsubo
開始: 2025-11-17 01:00:35 +0900
優先度: P1
提案元: branch=master
関連ストーリー: N/A

## 目的
- Codex（エージェント）が `docs/ARCHITECTURE.md` や関連設計ドキュメントと、実際のコード構成（ディレクトリ/Port/Adapter 境界）を比較し、差分を一覧化する workflow を用意する。
- 「設計上そうなっているはず」と「実装が実際どうなっているか」のギャップを、Story/Discovery に落とし込みやすい形で可視化する。

## 背景
- アーキテクチャ設計は更新頻度が低い一方で、実装側の変更は細かく積み上がるため、設計ドキュメントとコードの乖離が発生しやすい。
- 人間が手で追うには範囲が広いため、Codex による機械的な「構造の照合＋疑わしい差分の列挙」ができると、リファクタリングや設計見直しの入口として便利。

## DoD候補
- [ ] docs/workflows/ 以下に、「アーキテクチャ設計と実装の差分チェック workflow（Codex 実行前提）」の手順が 1 本のドキュメントとしてまとまっている。
- [ ] `docs/ARCHITECTURE.md` に記載されている代表的な境界（例: BattlePort/DataPort/InventoryPort、scenes/usecase/repo の依存関係）について、この workflow を走らせた結果の差分メモや Discovery が 1〜2 件ある。
- [ ] Codex に投げるためのプロンプト雛形（対象セクション/対象ディレクトリ/期待する出力フォーマット）がドキュメント化されている。
- [ ] 差分結果をもとに新しい Story/Discovery を起こす運用（例: 「差分一覧 → 優先度づけ → Story 切り出し」）が stories/workflows のどこかに簡潔に記載されている。

## 関連
- docs/ARCHITECTURE.md
- docs/API.md
- docs/workflows/overview.md
- docs/workflows/stories.md

## 進捗ログ
- 2025-11-17 01:00:35 +0900: 起票

- 2025-11-17 01:01:12 +0900: Backlog へ昇格（優先度=P1）
