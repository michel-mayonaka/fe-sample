# 03_ドキュメント更新と後続タスク（実装ストーリー/Backlog）の整理

ステータス: [完了]
担当: @tkg-engineer
開始: 2025-11-19 00:28:21 +0900

## 目的
- util 名を避ける方針と今回の整理の位置付けをドキュメントに反映し、実装用ストーリーや Backlog を整理する。

## 完了条件（DoD）
- [x] docs/KNOWLEDGE/engineering/naming.md や必要に応じて docs/architecture/README.md に、`internal/game/util` の扱いと今後の方針が追記されている（少なくともメモレベル）。
- [x] 実装フェーズ用のストーリーまたは Backlog エントリが、必要な単位（例: rng/rect/debug ごとなど）で整理されている。
- [x] 本ストーリーの DoD と discovery の状態（consumed）が整合していることを確認している。

## 作業手順（概略）
- `02_再配置/削除方針の検討` の結果をもとに、命名/アーキテクチャの方針をドキュメントへ反映する。
- 実装段階で扱うべき単位ごとに、Story/Backlog を起票または更新する。
- 必要であれば、今回の検討内容を短いノートとして残す。

## 進捗ログ
- 2025-11-19 00:28:21 +0900: タスク作成。
- 2025-11-18 23:05 JST: docs/KNOWLEDGE/engineering/naming.md へケーススタディを追記し、Backlog 項目を "完了/P2" 扱いで整理。discovery（consumed）との整合を確認。

## 実施メモ（2025-11-18）
- docs/KNOWLEDGE/engineering/naming.md の「汎用名の避け方」に `internal/game/util` 撤去のケーススタディ節を追加し、再発防止とサブパッケージ命名の判断基準を具体例付きで示した。
- stories/BACKLOG.md の該当エントリを `[完了/P2]` とし、今回のストーリーで util 撤去を確認済であること、今後は Discovery → Story を経て責務特化パッケージを起こすことを明文化した。
- discovery: `stories/discovery/consumed/2025-09-30-migrated-07.md` はすでに consumed 状態で、本ストーリーのログにもストーリー連携済みの記録があるため追加作業不要。
- 追加の実装ストーリーは現時点で不要。`math/rand` → `internal/game/rng` への統合など実装判断が必要な場合は別 Discovery を起票する前提で合意。

## 依存／ブロッカー
- `02_再配置/削除方針の検討` の完了。

## 成果物リンク
- 更新した docs/KNOWLEDGE/engineering/naming.md / docs/architecture/README.md、関連 Story/Backlog
