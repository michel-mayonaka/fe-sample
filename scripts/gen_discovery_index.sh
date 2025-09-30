#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "$0")/.." && pwd)"
disc_dir="$root_dir/stories/discovery"
out="$disc_dir/INDEX.md"

mkdir -p "$disc_dir"

ts_now="$(date +'%Y-%m-%d %H:%M:%S %z')"

list_items() {
  local dir="$1" status_label="$2"
  local count=0
  find "$dir" -maxdepth 1 -type f -name '*.md' ! -name 'INDEX.md' ! -name '.gitkeep' -print | while read -r f; do
    title=$(grep -m1 -E '^# ' "$f" | sed -E 's/^# +//; s/^discovery: +//;')
    prio=$(grep -m1 -E '^優先度:' "$f" | awk -F: '{gsub(/ /, "", $2); print $2}')
    created=$(grep -m1 -E '^作成:' "$f" | sed -E 's/^作成:[[:space:]]*//')
    state=$(grep -m1 -E '^(状態|ステータス):' "$f" | sed -E 's/^(状態|ステータス):[[:space:]]*//')
    printf -- "- [%s] %s | %s | %s\n" "${prio:-P?}" "${title:-(no title)}" "${created:-N/A}" "${state:-N/A}"
    count=$((count+1))
  done
}

{
  echo "# Discovery Index"
  echo ""
  echo "生成: ${ts_now}"
  echo ""
  echo "## Open (stories/discovery/*.md)"
  list_items "$disc_dir" "open" || true
  echo ""
  echo "## Promoted (stories/discovery/accepted)"
  list_items "$disc_dir/accepted" "promoted" || true
  echo ""
  echo "## Declined (stories/discovery/declined)"
  list_items "$disc_dir/declined" "declined" || true
} > "$out"

echo "[discovery-index] generated: $out"
