#!/usr/bin/env bash
set -euo pipefail

# Usage: make decline-discovery FILE=stories/discovery/20250930-123456-foo.md [REASON="..."]

file="${FILE:-${1:-}}"
reason="${REASON:-やらない判断（優先度/効果/タイミングの理由などを追記）}"

if [[ -z "$file" || ! -f "$file" ]]; then
  echo "Usage: FILE=<discovery_file.md> [REASON=...]" >&2
  exit 1
fi

ts_now="$(date +'%Y-%m-%d %H:%M:%S %z')"

pat='s/^[[:space:]]*状態:[[:space:]]*(\[[^]]+\]|[^[:space:]].*)/状態: [declined]/'
if sed -i '' -E "$pat" "$file" 2>/dev/null; then :; else sed -i -E "$pat" "$file"; fi
printf "\n- %s: Decline（理由: %s）\n" "$ts_now" "$reason" >> "$file"

dest_dir="stories/discovery/declined"
mkdir -p "$dest_dir"
mv "$file" "$dest_dir/"

echo "[decline_discovery] moved to declined/ with reason."

