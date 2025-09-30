#!/usr/bin/env bash
set -euo pipefail

# Read a legacy BACKLOG.md from stdin and create accepted discovery files
# for each entry header like: '## [P0] 2025-09-30: タイトル'
# Usage: git show <ref>:stories/BACKLOG.md | scripts/migrate_backlog_to_discovery.sh

acc_dir="stories/discovery/accepted"
mkdir -p "$acc_dir"

tmpdir="$(mktemp -d)"; trap 'rm -rf "$tmpdir"' EXIT
src="$tmpdir/src.md"
cat > "$src"

idx=0
title=""; prio=""; date=""; purpose=""; background=""; dod=""; related=""

flush() {
  [[ -z "$title" ]] && return
  [[ -z "$date" ]] && date="$(date +%Y-%m-%d)"
  [[ -z "$prio" ]] && prio="P2"
  # Skip if same title already exists
  if rg -n "^# discovery:\s*$title$" "$acc_dir" -S >/dev/null 2>&1; then
    title=""; prio=""; date=""; purpose=""; background=""; dod=""; related=""; return
  fi
  idx=$((idx+1))
  fn="$acc_dir/${date}-migrated-$(printf '%02d' "$idx").md"
  {
    echo "# discovery: $title"
    echo
    echo "ステータス: [promoted]"
    echo "担当: @tkg-engineer"
    echo "開始: ${date} 00:00:00 +0900"
    echo "優先度: ${prio}"
    echo "提案元: legacy-backlog"
    echo "関連ストーリー: N/A"
    echo
    echo "## 目的"; [[ -n "$purpose" ]] && printf "%s" "$purpose" || echo "- （legacyから移行）"; echo
    echo "## 背景"; [[ -n "$background" ]] && printf "%s" "$background" || echo "- （legacyから移行）"; echo
    echo "## DoD候補"; [[ -n "$dod" ]] && printf "%s" "$dod" || echo "- （legacyから移行）"; echo
    echo "## 関連"; [[ -n "$related" ]] && printf "%s" "$related" || echo "- BACKLOG.md (legacy)"; echo
    echo "## 進捗ログ"; echo "- ${date} 00:00:00 +0900: legacy BACKLOG から移行"; echo
  } > "$fn"
  title=""; prio=""; date=""; purpose=""; background=""; dod=""; related=""
}

while IFS= read -r line; do
  if [[ $line =~ ^##\ \[P([0-9])\]\ ([0-9]{4}-[0-9]{2}-[0-9]{2}):\ (.*)$ ]]; then
    flush
    prio="P${BASH_REMATCH[1]}"; date="${BASH_REMATCH[2]}"; title="${BASH_REMATCH[3]}"
    continue
  fi
  if [[ "$line" =~ ^-\ 目的: ]]; then purpose+="- ${line#*- 目的: }\n"; continue; fi
  if [[ "$line" =~ ^-\ 背景: ]]; then background+="- ${line#*- 背景: }\n"; continue; fi
  if [[ "$line" =~ ^-\ DoD: ]]; then dod+="- ${line#*- DoD: }\n"; continue; fi
  if [[ "$line" =~ ^-\ 参考/関連: ]]; then related+="- ${line#*- 参考/関連: }\n"; continue; fi
done < "$src"

flush

echo "[migrate_backlog] migrated entries into $acc_dir"
