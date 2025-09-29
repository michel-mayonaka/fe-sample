# 01 — 現状調査と対象画面決定

ステータス: [完了]

## 目的
- 座標/サイズのハードコード箇所を洗い出し、初回適用対象画面を決める。

## 対象候補
- `internal/game/service/ui/layout.go`（定数群）
- `internal/game/service/ui/metrics.go`（スケール算出/ユーティリティ）
- `internal/game/ui/layout/*.go`（行レイアウトなど）
- `internal/game/ui/draw/*.go`（一部のオフセット/余白）

## 調査手順
- `rg -n "S\(|List.*Px|Offset|Margin|Gap|Sz|Rect\(" internal/game | sort`
- 代表画面は「インベントリ一覧」を第一候補とする（一覧/パネル/タイトル/行の一通りが揃うため）。

## 成果物
- 対象キー一覧（一次候補）
  - list.margin（既存: 24）
  - list.itemH（100）
  - list.itemGap（12）
  - list.portraitSize（80）
  - list.titleOffset（44）
  - line.main（26）/ line.small（22）
- 代表画面の初回適用対象: インベントリ一覧（`internal/game/ui/layout/list.go` / `internal/game/ui/draw/inventory_list.go`）
- 参考: 画面内の固定オフセット（次段で検討）
  - `S(16), S(32), S(8)`（パネル内余白・タイトル下マージン）
  - ヘッダ列X: `+560, +680, +760, +840, +920, +1000, +1080, +1160`（必要なら `list.columns` に昇格）

## 進捗ログ
- 2025-09-29: 雛形作成
- 2025-09-29: 調査完了・初回適用対象をインベントリ一覧に決定
