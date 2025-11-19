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
- 2025-09-29: 着手。UI 層の直接書込検知→ `internal/game/scenes/status/status.go` の `UserTable.UpdateCharacter` 呼び出しを削除し、`DataPort.PersistUnit` への委譲に統一。
- 2025-09-29: `rg` チェック結果 — UI→Repo 直接書込 0 件、`WeaponsTable()` は `gdata.Provider()` 経由のみを確認。
- 2025-09-29: `make mcp` 実行 — `check-all`/`check-ui` OK。`golangci-lint` は環境依存メッセージ（非必須）。
- 2025-09-29: `/pkg/game` テストを拡充し coverage 87.2% に到達。`internal/usecase` は 42.9%（追加テストを作成）。
- 2025-09-29: `internal/usecase` に `RunBattleRound/ReloadData/PersistUnit/Equip(Weapon|Item)` のテストを追加（反撃射程外/追撃/スロット拡張/エラー伝播を網羅）。

## 関連
- `docs/architecture/README.md`
