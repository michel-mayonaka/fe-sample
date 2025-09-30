
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

# 現在時刻（秒精度、タイムゾーン付き）
now_ts="$(date +'%Y-%m-%d %H:%M:%S %z')"

cat >"$dir/README.md" <<EOF
# ${dateprefix}-${slug} — Story タイトル

ステータス: [進行中]
担当: @yourname
開始: ${now_ts}

## 目的・背景
- このストーリーで解決する課題/価値を1〜3行で。

## スコープ（成果）
- 具体的なアウトカム（ユーザ視点/開発者視点）。

## 受け入れ基準（Definition of Done）
- [ ] ビルド/テスト基準
- [ ] UI/CLI での確認手順
- [ ] ドキュメント整合

## 工程（サブタスク）
- [ ] タスクA（リンク: \`stories/YYYYMMDD-slug/tasks/01_xxx.md\`）
- [ ] タスクB
- [ ] タスクC

## 計画（目安）
- 見積: X 時間 / セッション
- マイルストン: M1 / M2 / M3

## 進捗・決定事項（ログ）
- ${now_ts}: ストーリー作成

## リスク・懸念
- 例: 依存の変更、CI制約 など

## 関連
- PR: #
- Issue: #
- Docs: \`docs/...\`
EOF

echo "[new_story] created: $dir"
