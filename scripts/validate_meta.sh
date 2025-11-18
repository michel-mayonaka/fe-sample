#!/usr/bin/env bash
set -euo pipefail

# Lightweight validator for required meta keys in Story/Discovery files.
# Usage: scripts/validate_meta.sh [paths...]

paths=("$@")
if [[ ${#paths[@]} -eq 0 ]]; then
  # Git が使える場合は tracked な stories 配下のみを対象にする
  if command -v git >/dev/null 2>&1 && git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    while IFS= read -r p; do
      paths+=("$p")
    done < <(git ls-files 'stories/**/*.md')
  else
    # Git が使えない環境向けのフォールバック（未追跡ファイルも含める）
    while IFS= read -r p; do
      paths+=("$p")
    done < <(find stories -type f -name '*.md')
  fi
fi

missing=0
for f in "${paths[@]}"; do
  [[ -f "$f" ]] || continue
  # Only check Story/Discovery top-level README or discovery files
  base="$(basename "$f")"
  dir="$(dirname "$f")"
  if [[ "$base" == "README.md" || "$dir" == *"/stories/discovery"* || "$dir" == *"/stories/discovery/accepted"* || "$dir" == *"/stories/discovery/declined"* ]]; then
    for key in ステータス 担当 開始; do
      if ! grep -q -E "^[[:space:]]*${key}:" "$f"; then
        echo "[warn] missing '${key}:' in $f" >&2
        missing=1
      fi
    done
  fi
done

exit $missing
