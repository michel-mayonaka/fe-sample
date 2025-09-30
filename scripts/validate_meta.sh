#!/usr/bin/env bash
set -euo pipefail

# Lightweight validator for required meta keys in Story/Discovery files.
# Usage: scripts/validate_meta.sh [paths...]

paths=("$@")
if [[ ${#paths[@]} -eq 0 ]]; then
  mapfile -t paths < <(git ls-files 'stories/**/*.md')
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

