# 02_再配置/削除方針の検討 — 責務特化サブパッケージ案

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 00:28:21 +0900

## 目的
- `internal/game/util` の中身を「削除するもの」と「責務ごとの専用サブパッケージに移すもの」に分け、具体的な再配置案を作成する。

## 完了条件（DoD）
- [x] 機能ごとに、`rng`/`rect`/`debug` 等のような責務特化サブパッケージ案が整理されている。
- [x] どのシンボルをどこへ移すかの対応表（ラフでよい）が作られている。
- [x] 実際の移動は別ストーリー/Backlog で行う前提として、そのための前提条件や注意点（循環依存など）がメモされている。

## 作業手順（概略）
- `01_現状調査` の結果をもとに、役割ごとにグルーピングする。
- docs/KNOWLEDGE/engineering/naming.md と照らし合わせて、妥当なパッケージ名を検討する。
- 再配置案と削除候補をテキストでまとめる。

## 進捗ログ
- 2025-11-19 00:28:21 +0900: タスク作成。

## 再配置/削除方針メモ（2025-11-18）
- `internal/game/util` の実体はすでに消滅し、唯一残っていた乱数ラッパは `internal/game/rng` として独立済み。よって「撤去」自体は完了扱いとし、本タスクでは再発防止の命名指針と `rng` 活用方針を整理する。
- 今後 util 的役割が必要になった場合は、以下のように責務ごとのサブパッケージへ直接追加する。
  - 乱数: `internal/game/rng`（既存）。`math/rand` の薄いラッパに加え、seed 注入やテスト用 deterministic generator を集約する。
  - 幾何判定／矩形: `internal/game/geom/rect` などの専用パッケージを新設する。現状該当機能なし。
  - UI 補助: UI サービス以下（例: `internal/game/service/ui/textdraw`）のようにより具体的な名称で分割。既存の `service/ui/util.go` は adapter との橋渡しコードのみなので `uibridge.go` 等への rename を後続 backlog で検討。

### 対応表（ラフ）
| 旧 util 機能 | 現状 | 方針 |
| --- | --- | --- |
| `Rand`/`NewRand` | `internal/game/rng.Rand`, `frame_context.Ctx.Rand` | `rng` に集約済。Ctx で `*rng.Rand` を expose し、利用側は `math/rand` 直接参照を避けるよう段階移行する。 |
| （該当なし） | - | 今後 util 的構造を作らない。必要が生じた場合はドメイン名 + パッケージで起票し直す。 |

### 注意点・前提
- `internal/game/rng` は現在 `Ctx` でのみ掴まれており、Scene 側は `math/rand` を使っている。循環依存を避けるため、`scenes.Env.RNG` を `*rng.Rand` に置き換える場合は `scenes` → `game` の import を新たに増やさないよう、`rng` を最下層 package（`internal/game/rng`）として保つ。
- `internal/usecase` 層は `pkg/game` と `math/rand` を直接扱っているため、`internal/game/rng` を共有化するかは別ストーリーの設計判断に委ねる（アプリケーション層に閉じた util であるため、強制置換は不要）。
- util 名を再導入しないよう docs/KNOWLEDGE/engineering/naming.md で今回の整理内容を明文化し、ストーリー完了後に Backlog から項目を除去しておく。

## 依存／ブロッカー
- `01_現状調査` の完了。

## 成果物リンク
- 再配置/削除方針メモ
