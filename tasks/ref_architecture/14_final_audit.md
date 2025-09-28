# 14_final_audit — 仕上げと棚卸し

## 目的・背景
- 参照経路/命名/コメント/テストを横断チェックし、移行の最終品質を担保。

## 作業項目（変更点）
- 依存の検査（UI→Port/Provider、Usecase→Repo、UIがRepoへ触れていない）。
- 命名/コメントの一貫性確認、不要コード/コメントの削除。
- 主要ユースケースのテスト不足を TODO として `tasks/known_issues.md` に集約。

## 完了条件
- `make mcp` 成功。
- `go test ./... -race -cover` 成功（環境的に可能な範囲）。
- known issues が整理され、クリティカルは 0。

## 影響範囲
- 全体（ただし実変更は最小）。

## 手順
1) 検索: `rg -n "UserTable\.Update|Save\(" internal/game/scenes`（UI直書きの残骸検知）。
2) 検索: `rg -n "WeaponsTable\(" internal/game/scenes`（Provider統一の漏れ検知）。
3) `make mcp` → `go test ./... -race -cover`。
4) known issues へ記録。

## 品質ゲート
- `make mcp`
- `go test ./... -race -cover`

## 動作確認
- 主要フロー（一覧/ステータス/在庫/戦闘）が通る。

## ロールバック
- 変更は最小・安全のため不要想定。

## Progress / Notes
- YYYY-MM-DD: 着手
- YYYY-MM-DD: 完了

## 関連
- `docs/ARCHITECTURE.md`

