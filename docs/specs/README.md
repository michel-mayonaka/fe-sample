# SPECS — 仕様ハブ

`docs/SPECS/` は「実装に直結する仕様」を集約するディレクトリです。world（設定）、gameplay（ユースケース/ロジック）、ui（画面）、reference（API 等）を分けて管理し、テンプレートを共通化します。

## 1. 目的
- 仕様本文を 1 か所にまとめ、実装・テストの根拠を明確にする。
- world → gameplay → ui の順に情報を辿れるよう、カテゴリ境界を揃える。
- エージェント/人間が同じ情報を参照し、矛盾があればストーリー/Discovery で検出できる状態を保つ。

## 2. ディレクトリ構成
```
docs/SPECS/
  README.md             # 本ファイル
  AGENTS.md             # 仕様の読み方・参照優先度（エージェント向け）
  world/                # ロア・マスターデータ背景（例: 勢力/装備体系）
  gameplay/             # 戦闘/在庫などの振る舞い（旧 system）
  ui/                   # 画面仕様
  reference/            # API などの補助リファレンス
  templates/
    gameplay.md         # gameplay 用テンプレート
    ui.md               # UI 用テンプレート
```

## 3. 仕様の書き方
1. `templates/` から適切なテンプレートをコピーする（world 向けテンプレートは後続ストーリーで追加予定）。
2. ファイル先頭にメタデータを記載する（例: `状態`, `主な実装`, `最新ストーリー`）。
3. gameplay ↔ ui の対応関係を相互リンクし、背景が world にある場合は `world/` へ追記する。
4. 実装前に spec を更新し、DoD へ「対象 spec 更新」を含める。

## 4. 入口と参照優先度
1. `AGENTS.md` → `docs/ops-overview.md` → ストーリー DoD
2. `docs/SPECS/AGENTS.md`: 仕様の読み方と優先順位
3. 各カテゴリの README（例: `world/README.md`）
4. 個別仕様ファイル（`gameplay/battle_basic.md` など）

## 5. テンプレート一覧
- `templates/gameplay.md`: ユースケース/振る舞い/データ/テスト観点。
- `templates/ui.md`: 画面概要/レイアウト/入力と挙動/状態遷移。

## 6. 今後の拡張
- `world/` の充実（勢力図や時間軸の整理）。
- `reference/` への API/テレメトリ仕様の追加。
- `templates/world.md` の追加、および gameplay 補助資料（シーケンス図など）の標準化。
