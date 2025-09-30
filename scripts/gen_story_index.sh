#!/usr/bin/env bash
set -euo pipefail

# Generate an index of finished stories into stories/finish/INDEX.md

root_dir="$(cd "$(dirname "$0")/.." && pwd)"
finish_dir="$root_dir/stories/finish"
out="$finish_dir/INDEX.md"

if [[ ! -d "$finish_dir" ]]; then
  echo "finish directory not found: $finish_dir" >&2
  exit 1
fi

list_dirs() {
  ls -1d "$finish_dir"/* 2>/dev/null | sort -r | awk 'system("test -d \""$0"\"")==0 {print $0}' || true
}

parse_field() {
  local file="$1" key="$2"
  # Extracts a value like: "ステータス: [完了]" or "- ステータス: [完了]"
  grep -m1 -E "^[[:space:]]*(-[[:space:]]*)?${key}:[[:space:]]*" "$file" | sed -E "s/^[[:space:]]*(-[[:space:]]*)?${key}:[[:space:]]*//" || true
}

get_last_update() {
  local file="$1"
  if command -v git >/dev/null 2>&1 && git -C "$root_dir" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    git -C "$root_dir" log -1 --format='%ci' -- "$file" 2>/dev/null || stat -f '%Sm' -t '%Y-%m-%d %H:%M:%S %z' "$file" 2>/dev/null || echo ""
  else
    # BSD stat
    stat -f '%Sm' -t '%Y-%m-%d %H:%M:%S %z' "$file" 2>/dev/null || echo ""
  fi
}

count=0
tmp="$(mktemp)"
trap 'rm -f "$tmp"' EXIT

for d in $(list_dirs); do
  readme="$d/README.md"
  [[ -f "$readme" ]] || continue
  slug="$(basename "$d")"
  title="$(grep -m1 '^# ' "$readme" | sed -E 's/^# //')"
  [[ -n "$title" ]] || title="$slug"
  status="$(parse_field "$readme" 'ステータス')"
  owner="$(parse_field "$readme" '担当')"
  date_line="$(parse_field "$readme" '日付')"
  started="$(parse_field "$readme" '開始')"
  updated="$(get_last_update "$readme")"

  printf -- "- %s — %s | ステータス: %s | 担当: %s | 日付: %s | 開始: %s | 更新: %s\n" \
    "$slug" "$title" "${status:-不明}" "${owner:-不明}" "${date_line:-N/A}" "${started:-N/A}" "${updated:-}" >>"$tmp"
  count=$((count+1))
done

{
  echo "# 完了ストーリー索引"
  echo
  echo "生成: $(date +'%Y-%m-%d %H:%M:%S %z')"
  echo "件数: ${count}"
  echo
  if [[ $count -gt 0 ]]; then
    cat "$tmp"
  else
    echo "(完了ストーリーはまだありません)"
  fi
} >"$out"

echo "[story-index] generated: $out"
