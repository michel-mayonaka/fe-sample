#!/usr/bin/env bash
set -euo pipefail

# Generate stories/BACKLOG.md from stories/discovery/accepted/*.md

root_dir="$(cd "$(dirname "$0")/.." && pwd)"
acc_dir="$root_dir/stories/discovery/accepted"
backlog="$root_dir/stories/BACKLOG.md"

mkdir -p "$acc_dir"

ts_now="$(date +'%Y-%m-%d %H:%M:%S %z')"

header_tmp="$(mktemp)"; body_tmp="$(mktemp)"; out_tmp="$(mktemp)"
trap 'rm -f "$header_tmp" "$body_tmp" "$out_tmp"' EXIT

# Split existing BACKLOG at the first occurrence of "初期エントリ" to keep the header as-is
if [[ -f "$backlog" ]]; then
  line_no=$(grep -n '^初期エントリ' "$backlog" | head -n1 | cut -d: -f1 || true)
  if [[ -n "$line_no" ]]; then
    head -n "$line_no" "$backlog" > "$header_tmp"
  else
    # No marker; copy whole file as header then add the marker
    cat "$backlog" > "$header_tmp"
    echo "" >> "$header_tmp"
    echo "初期エントリ" >> "$header_tmp"
  fi
else
  # Minimal header if BACKLOG does not exist
  cat >"$header_tmp" <<EOF
# Stories Backlog（将来のストーリー候補）

本ファイルは、自動生成されます（編集しないでください）。

初期エントリ
EOF
fi

# Collect accepted discoveries
tmp_list="$(mktemp)"; trap 'rm -f "$tmp_list"' EXIT

find "$acc_dir" -maxdepth 1 -type f -name '*.md' -print | while read -r f; do
  title=$(grep -m1 -E '^# ' "$f" | sed -E 's/^# +//; s/^discovery: +//;')
  prio=$(grep -m1 -E '^優先度:' "$f" | awk -F: '{gsub(/ /, "", $2); print $2}')
  created=$(grep -m1 -E '^(作成|開始):' "$f" | sed -E 's/^(作成|開始):[[:space:]]*//')
  date_short=$(echo "$created" | cut -d' ' -f1)
  story=$(grep -m1 -E '^関連ストーリー:' "$f" | sed -E 's/^関連ストーリー:[[:space:]]*//')
  prio_sort=9
  if [[ "$prio" =~ ^P([0-9]+)$ ]]; then
    prio_sort="${BASH_REMATCH[1]}"
  fi
  printf "%s\t%s\t%s\t%s\t%s\n" "${prio_sort}" "${date_short}" "${prio}" "${title}" "${story}" >> "$tmp_list"
done

# Sort by priority (asc) then date (desc)
sort -t $'\t' -k1,1n -k2,2r "$tmp_list" | while IFS=$'\t' read -r prio_s d prio title story; do
  printf "## [%s] %s: %s\n" "$prio" "$d" "$title" >> "$body_tmp"
  printf -- "- 目的: （discovery参照）\n" >> "$body_tmp"
  printf -- "- 背景: （discovery参照）\n" >> "$body_tmp"
  printf -- "- DoD: （discoveryの DoD候補を要約 or 後続で具体化）\n" >> "$body_tmp"
  printf -- "- 参考/関連: Story: %s\n\n" "${story:-N/A}" >> "$body_tmp"
done

{
  cat "$header_tmp"
  echo ""
  cat "$body_tmp"
} > "$out_tmp"

mv "$out_tmp" "$backlog"
echo "[backlog-index] generated: $backlog at $ts_now"
