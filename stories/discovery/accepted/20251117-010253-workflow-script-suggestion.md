# discovery: workflowからスクリプト化候補を洗い出す workflow

ステータス: [promoted]
担当: @kyosuke_tsubo
開始: 2025-11-17 01:02:53 +0900
優先度: P1
提案元: branch=master
関連ストーリー: N/A

## 目的
- 既存の workflow ドキュメント（`docs/workflows/*` や Story テンプレのフロー）を Codex が読み取り、「人手で毎回やるよりスクリプト化した方が良いステップ」を候補として洗い出すための標準 workflow を用意する。
- 「スクリプト化すべきかどうか」の判断材料と、実際に作るべきスクリプトの TODO リストを、Discovery/Backlog として管理しやすい形で出力できるようにする。

## 背景
- 手順が増えると、workflow ドキュメントだけでは「何を自動化すべきか」が埋もれがちで、結果として手作業が多くなりミスや漏れが発生しやすい。
- Codex に対して ad-hoc に「どれスクリプト化した方がいい？」と聞くのではなく、「workflow を入力としてスクリプト候補を列挙する」という共通の枠組みを用意しておくと、定期的な改善サイクルに乗せやすい。

## DoD候補
- [ ] docs/workflows/ 以下に「workflow からスクリプト化候補を洗い出す workflow（Codex 実行前提）」の手順が 1 本のドキュメントとしてまとまっている。
- [ ] 最低 1 つ以上の既存 workflow（例: stories/workflows, discovery/backlog 運用, specs 状態メタデータ運用など）に対してこの workflow を適用し、「スクリプト化候補リスト」を Discovery or Story の形で残している。
- [ ] Codex に投げるためのプロンプト雛形（対象 workflow ファイル、評価観点、出力フォーマット）がドキュメント化されている。
- [ ] 新しいスクリプト候補を Backlog/Discovery に流すルート（例: 「候補リスト → 重要度順に Discovery 起票」）が stories/workflows に簡潔に書かれている。

## 関連
- docs/workflows/*
- stories/BACKLOG.md
- scripts/*.sh

## 進捗ログ
- 2025-11-17 01:02:53 +0900: 起票

- 2025-11-17 01:03:27 +0900: Backlog へ昇格（優先度=P1）
