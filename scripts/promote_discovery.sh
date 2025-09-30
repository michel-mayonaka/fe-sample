#!/usr/bin/env bash
set -euo pipefail

# Usage: make promote-discovery FILE=stories/discovery/20250930-123456-foo.md [PRIORITY=P2]

file="${FILE:-${1:-}}"
priority="${PRIORITY:-}"

if [[ -z "$file" || ! -f "$file" ]]; then
  echo "Usage: FILE=<discovery_file.md> [PRIORITY=P2]" >&2
  exit 1
fi

ts_now="$(date +'%Y-%m-%d %H:%M:%S %z')"
date_short="$(date +%Y-%m-%d)"

# Extract fields
title="$(grep -m1 -E '^# ' "$file" | sed -E 's/^# +//; s/^discovery: +//;')"
[ -n "$title" ] || title="$(basename "$file" .md)"
prio_file="$(grep -m1 -E '^優先度:' "$file" | awk -F: '{gsub(/ /, "", $2); print $2}')"
prio="${priority:-${prio_file:-P2}}"
story_link="$(grep -m1 -E '^関連ストーリー:' "$file" | sed -E 's/^関連ストーリー:\s*//')"
[ -n "$story_link" ] || story_link="N/A"

entry_tmp="$(mktemp)"
trap 'rm -f "$entry_tmp"' EXIT

cat >"$entry_tmp" <<EOF
## [${prio}] ${date_short}: ${title}
- 目的: （discovery参照）
- 背景: （discovery参照）
- DoD: （discoveryの DoD候補を要約 or 後続で具体化）
- 参考/関連: ${file}, Story: ${story_link}

EOF

backlog="stories/BACKLOG.md"
if [[ ! -f "$backlog" ]]; then
  echo "backlog not found: $backlog" >&2
  exit 1
fi

# Insert after the "初期エントリ" heading if present; otherwise append to end
line_no=$(grep -n '^初期エントリ' "$backlog" | head -n1 | cut -d: -f1 || true)
if [[ -n "${line_no}" ]]; then
  head -n "$line_no" "$backlog" > "$backlog.tmp"
  cat "$entry_tmp" >> "$backlog.tmp"
  tail -n +$((line_no+1)) "$backlog" >> "$backlog.tmp"
  mv "$backlog.tmp" "$backlog"
else
  cat "$entry_tmp" >> "$backlog"
fi

# Update discovery status/priority and move to accepted
pat='s/^[[:space:]]*(状態|ステータス):[[:space:]]*(\[[^]]+\]|[^[:space:]].*)/ステータス: [promoted]/'
if sed -i '' -E "$pat" "$file" 2>/dev/null; then :; else sed -i -E "$pat" "$file"; fi
pat2="s/^優先度:.*/優先度: ${prio}/"
if sed -i '' -E "$pat2" "$file" 2>/dev/null; then :; else sed -i -E "$pat2" "$file"; fi
printf "\n- %s: Backlog へ昇格（優先度=%s）\n" "$ts_now" "$prio" >> "$file"

dest_dir="stories/discovery/accepted"
mkdir -p "$dest_dir"
mv "$file" "$dest_dir/"

echo "[promote_discovery] added entry to BACKLOG and moved to accepted/"
