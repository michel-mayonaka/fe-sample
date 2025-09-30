#!/usr/bin/env bash
set -euo pipefail

# Usage:
#   make consume-discovery FILE=stories/discovery/accepted/....md STORY_DIR=stories/20251001-some-story

file="${FILE:-${1:-}}"
story_dir="${STORY_DIR:-${2:-}}"

if [[ -z "$file" || -z "$story_dir" ]]; then
  echo "Usage: FILE=<accepted.md> STORY_DIR=<stories/YYYYMMDD-slug>" >&2
  exit 1
fi

if [[ ! -f "$file" ]]; then
  echo "file not found: $file" >&2; exit 1
fi
if [[ ! -d "$story_dir" ]]; then
  echo "story dir not found: $story_dir" >&2; exit 1
fi

ts_now="$(date +'%Y-%m-%d %H:%M:%S %z')"

# Update status to consumed
pat='s/^[[:space:]]*ステータス:[[:space:]]*(\[[^]]+\]|[^[:space:]].*)/ステータス: [consumed]/'
if sed -i '' -E "$pat" "$file" 2>/dev/null; then :; else sed -i -E "$pat" "$file"; fi

# Ensure related story link
if grep -q -E '^関連ストーリー:' "$file"; then
  if sed -i '' -E "s|^関連ストーリー:.*$|関連ストーリー: $(basename "$story_dir")|" "$file" 2>/dev/null; then :; else sed -i -E "s|^関連ストーリー:.*$|関連ストーリー: $(basename "$story_dir")|" "$file"; fi
else
  printf "\n関連ストーリー: %s\n" "$(basename "$story_dir")" >> "$file"
fi

# Append log
printf "\n- %s: ストーリー化（%s）\n" "$ts_now" "$(basename "$story_dir")" >> "$file"

# Move to consumed
dest_dir="stories/discovery/consumed"
mkdir -p "$dest_dir"
mv "$file" "$dest_dir/"

echo "[consume_discovery] moved to consumed/: $(basename "$file")"

