
#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 || -z "$1" ]]; then
  echo "Usage: $0 <slug> [YYYYMMDD]" >&2
  exit 1
fi

slug="$1"
dateprefix="${2:-$(date +%Y%m%d)}"
dir="stories/${dateprefix}-${slug}"

if [[ -e "$dir" ]]; then
  echo "[new_story] already exists: $dir" >&2
  exit 1
fi

mkdir -p "$dir"
cat >"$dir/README.md" <<'EOF'
# YYYYMMDD-slug — Story タイトル

ステータス: [進行中]
担当: @yourname

## 目的・背景
- このストーリーで解決する課題/価値を1〜3行で。

## スコープ（成果）
- 具体的なアウトカム（ユーザ視点/開発者視点）。

## 受け入れ基準（Definition of Done）
- [ ] ビルド/テスト基準
- [ ] UI/CLI での確認手順
- [ ] ドキュメント整合

## 工程（サブタスク）
- [ ] タスクA（リンク: `tasks/...`）
- [ ] タスクB
- [ ] タスクC

## 計画（目安）
- 見積: X 時間 / セッション
- マイルストン: M1 / M2 / M3

## 進捗・決定事項（ログ）
- YYYY-MM-DD: 着手/方針決定
- YYYY-MM-DD: 実装/検証

## リスク・懸念
- 例: 依存の変更、CI制約 など

## 関連
- PR: #
- Issue: #
- Docs: `docs/...`
EOF

echo "[new_story] created: $dir"
