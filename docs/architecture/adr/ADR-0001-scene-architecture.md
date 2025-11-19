# ADR-0001 — Scene / Usecase 分離と Port 設計

- **日付**: 2025-09-29
- **状態**: Accepted
- **背景**: 初期バージョンでは `internal/game/env.go` にすべてのユースケース（データ再読込、戦闘、在庫更新など）が集約され、Scene から直接 Repo や Model を触る箇所が散見された。この構造では責務肥大と依存方向の崩れが起き、UI リファクタ時の影響範囲が制御できない。
- **決定**:
  1. Scene は UI 入力/描画に専念し、更新系の操作は Port（`DataPort/BattlePort/InventoryPort` 等）を介して Usecase へ委譲する。
  2. クエリは `internal/game/data.Provider` に統一し、Scene から Repo やマスタデータを直接参照しない。
  3. Usecase 層は UI 資産（画像キャッシュなど）に触れず、副作用はインフラ層へ閉じ込める。
- **根拠**:
  - UI スタックの肥大化を防ぎ、Scene のテスト容易性を確保するため。
  - Codex/人間の両方が Port 単位で責務を把握でき、Discovery を切り出しやすくするため。
  - 今後のプラットフォーム追加（例: CLI/別 UI）時に、Usecase を流用できる形に保つため。
- **影響**:
  - `internal/game/scenes` では `Env` から必要な Port のみを受け取るよう実装を修正。
  - `internal/usecase/*` で Port 実装を分割し、`internal/game/data` で Provider を再編。
  - `docs/architecture/README.md` の「レイヤと役割」「依存原則」に本決定を明記し、以後の ADR でも決定履歴を管理する。
- **フォローアップ**:
  - Port 増加時はこの ADR を更新せず、新しい ADR を作成して判断を記録する。
  - Scene ごとの Port 利用状況を `stories/discovery/` で定期棚卸しし、肥大し始めたら対応する。
