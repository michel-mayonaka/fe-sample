# Stories Backlog（将来のストーリー候補）

本ファイルは、作業中に見つかった「別ストーリー化が望ましい改善点/アイデア」を一時的に蓄える場所です。必要に応じて `make new-story SLUG=<slug>` で昇格し、詳細は新規ストーリー配下に展開します。

運用ルール（簡易）
- 1エントリ=1セクション。短く要点を記述（目的/背景/DoD/関連）
- 重要度が高いものは先頭へ。完了したらエントリは削除
- 週次/セッション終わりに見直し、不要になったものは削除

エントリ雛形
```
## YYYY-MM-DD: タイトル（簡潔）
- 目的: 〜〜を改善/実現する
- 背景: 〜〜が散在/重複/不明確 など
- DoD: 〜〜が達成（例: mcp グリーン、rg 残存 0）
- 参考/関連: ファイル/PR/ストーリーの相互リンク
```

初期エントリ
## [P1] 2025-11-19: ストーリー運用について再整理
- 目的: ストーリー駆動開発と Codex 連携を前提に、実際の運用ルールとツール群を整理し、開発体験を安定・高速化する。
- 背景: 現状のストーリー運用には、ナレッジやスクリプトに落とし切れていない暗黙のフローが残っており、Codex と人間が協調する際に手作業や判断の揺れが発生している。
- DoD: `stories/AGENTS.md` や `docs/workflows/stories.md` などに最新の運用ルール/パターンが反映され、主要フロー（ストーリー起票〜完了、Discovery 取り込み、Codex との分担など）がスクリプト/テンプレートで再現可能になっている（`make story-index` 等のツールも更新済み、`make mcp` グリーン）。
- 参考/関連: docs/workflows/stories.md, AGENTS.md, docs/CODEX_CLOUD.md, stories/BACKLOG.md

## [P2] 2025-11-19: claudecodeのsubagents的な機能をcodexで実現のためのナレッジ追加
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: stories/discovery/20251119-003515-codex-subagents-knowledge.md, Story:  N/A

## [P2] 2025-11-17: ストーリー管理をGitHub Issueに移行するか検討
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: stories/discovery/20251117-022029-story-gh-issues.md, Story:  N/A

## [P2] 2025-11-17: ストーリー/BACKLOG用の作業分類タグの追加
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: stories/discovery/20251117-022029-story-tags.md, Story:  N/A

## [P2] 2025-11-17: 画面仕様の新規作成（タイトル/編成/バトルマップ）
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: stories/discovery/20251117-010944-ui-spec-new-screens.md, Story:  N/A

## [P1] 2025-11-17: workflowからスクリプト化候補を洗い出す workflow
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: stories/discovery/20251117-010253-workflow-script-suggestion.md, Story:  N/A

## [P1] 2025-11-17: アーキテクチャ設計と実装の差分チェック workflow
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: stories/discovery/20251117-010035-arch-impl-diff-workflow.md, Story:  N/A

## [P1] 2025-11-17: specと実装の差分チェック workflow
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: stories/discovery/20251117-010034-spec-impl-diff-workflow.md, Story:  N/A

## [P2] 2025-09-30: 入力レイアウト設定の外部化（config 駆動）
- キー/マウス割当を設定ファイルで差し替え可能にし、`provider/input/ebiten` のレイアウト注入を標準化する。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-30-migrated-06.md)）
- デフォルトレイアウトがコードに固定。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-30-migrated-06.md)）
- `config/input_layout.json` を追加、`app.NewUIAppGame` で読み込み→ `ginput.Layout` へ変換。既定不在時は従来デフォルトを採用、`make mcp` グリーン。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-30-migrated-06.md)）

## [P2] 2025-09-29: FPS デバッグ表示
- フレームパフォーマンスの可視化\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-09.md)）
- 画面増加に伴う描画コストの監視が必要\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-09.md)）
- トグル可能な FPS 表示(HUD)、秒間平均の表示、負荷小、`make mcp` グリーン\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-09.md)）

## [P2] 2025-09-29: マウス座標デバッグ表示
- マウスカーソル位置のゲーム座標/スクリーン座標を HUD に表示して調整を容易化\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-08.md)）
- レイアウト調整やクリック判定の検証負荷が高い\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-08.md)）
- トグル可能なデバッグ HUD 実装(例: F2)、座標オーバレイ表示、UI との重なり確認、パフォーマンス影響軽微\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-08.md)）

## [P2] 2025-09-27: SQLite 準備（将来）
- JSON→SQLite へ段階移行できる下地作り。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-14.md)）
- 将来の性能/整合性要件に備える。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-14.md)）
- `internal/repo` に SQLite 実装の雛形を追加（`modernc.org/sqlite`）、ビルドタグで切替。`docs/DB_NOTES.md` に移行手順とトランザクション方針を追記。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-14.md)）

## [P2] 2025-09-27: UI スケーリング調整
- 表示スケールに応じた見易さ改善。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-16.md)）
- 固定オフセットや折返し幅のハードコードが残存。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-16.md)）
- `uicore.S` の適用範囲拡大、折返し幅/余白の自動計算、HPバーの微調整。`make mcp` グリーン。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-16.md)）

## [P2] 2025-09-27: テスト拡充
- `usecase.App` と UI 純関数のテスト強化。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-15.md)）
- カバレッジと境界テストが不足。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-15.md)）
- `adapter.UIToGame`、Repo キャッシュ/Reload、`ForecastAtExplain` の整合性、UI純関数（Rect/折返し幅/スケール）テストを追加。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-15.md)）

## [P2] 2025-09-27: ドキュメント更新の追従（README/API/DB_NOTES）
- 構成変更と API の最新化をドキュメントへ反映。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-12.md)）
- README/API/DB_NOTES に旧記述が残存。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-12.md)）
- README/`docs/API.md`/`docs/DB_NOTES.md` を現行構成に同期、リンク整合、`make mcp` グリーン。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-12.md)）

## [P2] 2025-09-27: ドメインロジックの整理（UI色/文言生成）
- 予測/ログの文言生成を集約し、I18N を見据えた分離を行う。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-13.md)）
- 文言/色マッピングが散在。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-13.md)）
- 文言生成/色マッピングの統一ユーティリティを追加し、利用箇所を置換。\n
  （[→ Discovery](stories/discovery/accepted/2025-09-27-migrated-13.md)）

## [P3] 2025-09-29: FE風バトル マップ画面の叩き追加
- FE風シミュレーションバトルのマップ画面を試作し、キャラクター一覧からの遷移と配置ロジックの骨組みを固める\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-10.md)）
- 現状は一覧表示のみで戦闘画面が未実装のため、体験の流れが断絶している\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-10.md)）
- キャラクター一覧右上にマップ画面遷移ボタン追加、平地/森/砦レイヤーを持つマップJSONマスタを新設、`db/user/usr_characters.json` の味方1体・敵2体を所定の位置へ初期配置\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-10.md)）

## [P3] 2025-09-29: Portrait のロード責務の最終整理
- 画像読込（I/O）を描画層に寄せ、view-model はパスなど純データに限定\n- `ui/adapter` や `service/levelup` のテストを継続的に検証\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-17.md)）
- 現状 `ui/adapter` で PortraitLoader を注入しているが、用途により描画側での遅延読込が有利な場面がある\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-17.md)）
- view-model を `PortraitPath` に寄せる設計案の比較と方針確定（採用/見送りの記録）\n- CI で `make test-all-ui` を追加ジョブとして実行\n
  （[→ Discovery](stories/discovery/accepted/2025-09-29-migrated-17.md)）


## [P2] 2025-11-16: UI adapter/bridge 後続クリーンアップ
- 目的: `uicore` のフォールバック撤去、`ui/widgets` 互換APIの整理、回帰テスト拡充。
  （[→ Discovery](stories/discovery/accepted/2025-11-16-ui-adapter-cleanups.md)）
- DoD: 旧API削除・テスト追加・`make mcp` グリーン。

## [P2] 2025-11-16: specs/system ディレクトリの細分化
- 目的: `docs/specs/system/` をユースケース/ドメイン単位などで整理し、システム仕様を見通しよく分割する。
- 背景: system 配下に仕様が増えた際に、一箇所に集約しすぎると参照性が下がるため。
- DoD: `docs/specs/system/` のサブディレクトリ/命名方針が決まり、最初のいくつかの仕様が新構成に沿って配置されている（`make mcp` グリーン）。
- 参考/関連: docs/specs/README.md, docs/DOCS_STRUCTURE.md

## [P2] 2025-11-16: specs/ui ディレクトリの細分化
- 目的: `docs/specs/ui/` を画面種別やフロー単位で整理し、画面仕様を見つけやすくする。
- 背景: UI 画面数や状態が増えたときに、単一階層だとどの spec がどの画面か判別しづらくなるため。
- DoD: `docs/specs/ui/` のサブディレクトリ/命名方針が決まり、ステータス画面など代表的な画面仕様が新構成に沿って配置されている（`make mcp` グリーン）。
- 参考/関連: docs/specs/README.md, docs/DOCS_STRUCTURE.md
