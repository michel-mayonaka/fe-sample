# 20251013-ci-mcp-gha-mcp-fix — GitHub Actions で `make mcp` が失敗する不具合の修正

ステータス: [進行中]
担当: @tkg-engineer（想定）
開始: 2025-10-13 16:39:05 +0900

## 目的・背景
- GitHub Actions 上で `make mcp` が UI ビルド工程（`check-ui`）にて失敗し、ワークフローが落ちる。
- 失敗ログ抜粋（2025-10-13 提示）: `fatal error: X11/Xlib.h: No such file or directory`（Ebiten/GLFW のビルドに必要な X11 開発ヘッダ欠如）。
- 方針: CI では「環境依存の UI ビルド失敗を非 strict でスキップ」しつつ、別ジョブで依存導入した厳格（strict）UI ビルドも実行して品質を担保する。

## スコープ（成果）
- `make mcp` が GitHub Actions で安定して成功（緑）。
- `check-ui` のスキップ条件に X11 ヘッダ欠如などの典型的環境依存エラーを追加（非 strict 時）。
- 厳格 UI ビルド専用の CI ジョブを追加し、必要な X11/GL 開発パッケージを導入して `cmd/ui_sample` のビルドを検証。
- ドキュメント（AGENTS or README）に CI 方針（非 strict と strict の二段構え）を追記。

## 受け入れ基準（Definition of Done）
- [x] GitHub Actions の既存 `mcp` ワークフローが成功（`check-all` 内の `check-ui` は非 strict 条件でスキップ可）。
- [x] 追加した `ui-build-strict` ジョブで OS 依存パッケージ導入後に `cmd/ui_sample` がビルド成功。
- [x] `docs`/`README`/`AGENTS.md` の CI 章に運用方針と環境依存の扱いを追記。
- [x] ローカルでも `MCP_STRICT=1 make check-ui` が失敗時に適切なヘルプメッセージを出す。

## 工程（サブタスク）
- [x] 01: 失敗原因の整理と再現条件の明確化（`X11/Xlib.h` 欠如の検知パターン定義）〔`stories/20251013-ci-mcp-gha-mcp-fix/tasks/01_root_cause.md`〕
- [x] 02: Makefile `check-ui` の環境依存スキップ条件拡張（X11/GL 関連）〔`stories/20251013-ci-mcp-gha-mcp-fix/tasks/02_update_makefile_check_ui_skip_x11.md`〕
- [x] 03: GitHub Actions に厳格 UI ビルドジョブを追加（必要パッケージ導入）〔`stories/20251013-ci-mcp-gha-mcp-fix/tasks/03_update_gha_install_x_deps.md`〕
- [x] 04: CI 運用ドキュメント更新（strict/non-strict、変数 `MCP_STRICT`・`MCP_OFFLINE` の使い分け）〔`stories/20251013-ci-mcp-gha-mcp-fix/tasks/04_docs_ci_ui_requirements.md`〕

## 計画（目安）
- 見積: 2〜3 時間（内訳: 実装 1h、CI 実行待ち 1h、整備 0.5h）
- マイルストン: M1=検知条件実装 / M2=CI 追加 / M3=Docs 反映

## 進捗・決定事項（ログ）
- 2025-10-13 16:39:05 +0900: ストーリー作成（実装開始指示待ち）
- 2025-10-13 21:07:35 +0900: 作業開始。Makefile の `check-ui` 拡張、CI ジョブ追加、Docs 更新を実施。

## リスク・懸念
- 例: 依存の変更、CI制約 など

## 関連
- PR: #
- Issue: #
- Docs: `docs/...`
