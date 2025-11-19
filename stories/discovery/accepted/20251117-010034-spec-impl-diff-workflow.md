# discovery: specと実装の差分チェック workflow

ステータス: [promoted]
担当: @kyosuke_tsubo
開始: 2025-11-17 01:00:34 +0900
優先度: P1
提案元: branch=master
関連ストーリー: N/A

## 目的
- Codex（エージェント）が `docs/SPECS/` とコード（pkg/internal）を横断的に読み、仕様と実装の差分を「Story/Discovery として起こしやすい粒度」で棚卸しできる半自動 workflow を用意する。
- 人間が毎回フルスキャンしなくても、「どの spec がどこまで反映されていないか」のざっくりした見取り図を定期的に得られるようにする。

## 背景
- specs 駆動開発を回し始めると、仕様と実装のズレ（未実装/仕様の更新漏れ/テスト不足）が時間とともに溜まりがち。
- Codex に対して毎回 ad-hoc な指示を出すのではなく、「spec 対応状況を確認するための標準プロンプト＋手順」を決めておくと、定期棚卸しや Story 切り出しが楽になる。

## DoD候補
- [ ] docs/KNOWLEDGE/workflows/ 以下に、「spec と実装の差分チェック workflow（Codex 実行前提）」の手順が 1 本のドキュメントとしてまとまっている。
- [ ] 代表的な spec（例: `docs/SPECS/ui/status_screen.md`, `docs/SPECS/gameplay/battle_basic.md`）に対して、この workflow を実際に走らせた結果のサンプル（差分メモや Discovery）が 1〜2 件ある。
- [ ] Codex に投げるためのプロンプト雛形（例: 「対象 spec / 対象ディレクトリ / 出力フォーマット」を指定したテンプレ）が `docs/SPECS/AGENTS.md` か workflow ドキュメントから参照できる。
- [ ] 手動での運用に依存せず、将来的に CI などへ統合しても破綻しない形（「人が実行する」「Codex が補佐する」が明示された構成）になっている。

## 関連
- docs/SPECS/AGENTS.md
- docs/SPECS/gameplay/*
- docs/SPECS/ui/*
- docs/KNOWLEDGE/workflows/stories.md

## 進捗ログ
- 2025-11-17 01:00:34 +0900: 起票

- 2025-11-17 01:01:12 +0900: Backlog へ昇格（優先度=P1）
