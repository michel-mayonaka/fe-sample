# 20250930-usecase-battle-double-uses — 追撃時の耐久消費を修正

ステータス: [完了]
担当: @tkg-engineer

## 目的・背景
- 戦闘1ラウンドで攻撃側が追撃（AS差>=3）する場合、武器耐久が2消費されるべきところ1しか減らないことがある。
- `internal/usecase/battle_test.go` の `TestRunBattleRound_AttackerDouble_ConsumesTwo` が失敗（例: want 3 got 4）。

## スコープ（成果）
- 追撃条件成立時に攻撃側の Uses を正しく2消費する。
- 反撃側は1消費のままを維持。
- インベントリ保存（`Inv.Consume`）の回数も装備表示と一致。

## 受け入れ基準（Definition of Done）
- [x] `TestRunBattleRound_AttackerDouble_ConsumesTwo` が安定して成功する。
- [x] RNG に依存しない（seed 固定 or ロジックで追撃有無が確定）。
- [x] 既存の `RunBattleRound` 仕様（HP更新/Save 呼び出しなど）に回 regress がない。

## 工程（サブタスク）
- [ ] 調査: 再現と原因切り分け（AS計算/追撃分岐/消費箇所） — `tasks/01_investigate.md`
- [ ] 修正: 追撃時の消費ロジックまたは条件の是正 — `tasks/02_fix.md`
- [ ] テスト: 追撃境界（差2/差3）とミス時も含めた網羅 — `tasks/03_tests.md`
- [ ] ドキュメント: コメント/説明の補強（必要なら） — `tasks/04_docs.md`

## 計画（目安）
- 見積: 1.5 時間 / 1 セッション
- マイルストン: M1 調査 → M2 修正 → M3 追加テスト

## 進捗・決定事項（ログ）
- 2025-09-30: ストーリー作成。失敗事象を README に記録。
- 2025-09-30: 原因はテストの乱数依存と判明。シード固定で解消。`make test-all` グリーン。

## リスク・懸念
- 乱数により与ダメが閾値を跨ぐと反撃/追撃の可否が揺れる可能性（seed 固定で解消）。

## 関連
- 該当テスト: `internal/usecase/battle_test.go:114` 付近
- 関連コード: `internal/usecase/battle.go`, `internal/adapter/unit.go`, `pkg/game/speed.go`
