# 01_現状調査 — `internal/game/util` 配下と利用箇所の棚卸し

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 00:28:21 +0900

## 目的
- `internal/game/util` 配下にどのようなファイル/関数が存在し、どこからどのように利用されているかを把握し、削除や再配置の判断材料を揃える。

## 完了条件（DoD）
- [x] `internal/game/util` 配下のファイル一覧と、それぞれの役割の簡単な説明がメモされている。
- [x] `rg` などで洗い出した利用箇所の一覧があり、「どの機能がどのレイヤから使われているか」が分かる。
- [x] 「明らかに不要なもの」「用途が不明なもの」「再配置候補」のラフな分類が行われている。

## 作業手順（概略）
- `ls internal/game/util` と簡単なコードリーディングで中身を確認する。
- `rg \"internal/game/util\"` やシンボル名検索で利用箇所を列挙する。
- 結果を短いメモにまとめ、次タスクのインプットとする。

## 現状調査メモ（2025-11-18）
- `ls internal/game` の結果、`util` ディレクトリは存在しない。`rng` や `service` など主要ディレクトリのみが並び、Git 管理下にも空ディレクトリは残っていない。
- `rg "internal/game/util" -g "*.go" -g "*.md"` のヒット先は本ストーリー関連の Markdown のみで、Go コードからの import/参照はゼロ。実装上の依存はすでに解消されている。
- `find internal -type d -name '*util*'` もヒットなしで、utility 系ディレクトリそのものが消滅している。`internal/game/service/ui/util.go` などファイル単位の util 名はあるが、本タスクの対象外とする。
- 旧 util 配下の乱数ラッパは `internal/game/rng`（`doc.go`/`rng.go`）へ移設済。現時点で `g_rng` を import しているのは `internal/game/frame_context.go` のみで、`Ctx.Rand *g_rng.Rand` として公開しているが実利用は未実装。
- `math/rand` の直接利用箇所は `internal/game/app/core.go`（Env 初期化）、`internal/game/scenes/env.go`（Env.RNG 型）、`internal/game/scenes/sim/engine.go`（シミュレーション）、さらに `internal/usecase/*` の battle/facade 等に散在する。`internal/game/rng` は現状孤立状態。

### ラフ分類
- 明らかに不要: `internal/game/util` ディレクトリ（削除済）。
- 用途不明: 該当なし。
- 再配置候補: 旧 util → `internal/game/rng` へ既に分離済。今後 `math/rand` 直使用を `rng` に寄せるかは後続ストーリーで扱う。

### コマンド抜粋
```
$ ls internal/game
actor app consts.go data frame_context.go provider render rng scene.go scenes service ui
$ rg -n "internal/game/util" -g"*.go" -g"*.md"
stories/20251119-internal-game-util-specialize/README.md:1: ...
$ find internal -type d -name '*util*'
(出力なし)
```

## 進捗ログ
- 2025-11-19 00:28:21 +0900: タスク作成。
- 2025-11-18 22:55 JST: 現状調査を実施し、`internal/game/util` が削除済みであることと依存解消状況を確認。

## 依存／ブロッカー
- レポジトリ全体に対する検索が必要。

## 成果物リンク
- 現状調査メモ
