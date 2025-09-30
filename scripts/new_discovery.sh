#!/usr/bin/env bash
set -euo pipefail

# Usage: make new-discovery SLUG=<slug> [TITLE="..."] [PRIORITY=P2] [STORY=YYYYMMDD-slug]

slug="${SLUG:-${1:-}}"
title="${TITLE:-}"
priority="${PRIORITY:-P2}"
story="${STORY:-}"
no_suggest="${NO_SUGGEST:-0}"
owner_default="@${GIT_AUTHOR_NAME:-}"
if [[ -z "$owner_default" || "$owner_default" == "@" ]]; then
  owner_default="@$(git config --global user.name 2>/dev/null | tr ' ' '_' || echo owner)"
fi

if [[ -z "$slug" ]]; then
  echo "Usage: SLUG=<slug> [TITLE=...] [PRIORITY=P2] [STORY=YYYYMMDD-slug]" >&2
  exit 1
fi

ts_file="$(date +%Y%m%d-%H%M%S)"
ts_now="$(date +'%Y-%m-%d %H:%M:%S %z')"
branch="$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo main)"

dir="stories/discovery"
mkdir -p "$dir"
path="$dir/${ts_file}-${slug}.md"

title_default="${slug}"
title_final="${title:-$title_default}"

# Suggest similar entries (titles or slugs) to reduce duplicates
if [[ "$no_suggest" != "1" ]]; then
  echo "[new_discovery] similarity suggestions (if any):" >&2
  rg -n --no-heading -i "${slug//-/ }|${title_final}" stories/discovery stories/BACKLOG.md 2>/dev/null | head -n 10 || true
fi

cat >"$path" <<EOF
# discovery: ${title_final}

ステータス: [open]
担当: ${owner_default}
開始: ${ts_now}
優先度: ${priority}
提案元: branch=${branch}
関連ストーリー: ${story:-N/A}

## 目的
- （簡潔に）

## 背景
- （補足があれば）

## DoD候補
- [ ] （採用時に Backlog へ転記）

## 関連
- ファイル/PR/issue など

## 進捗ログ
- ${ts_now}: 起票
EOF

echo "[new_discovery] created: $path"
