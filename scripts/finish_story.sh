#!/usr/bin/env bash
set -euo pipefail

# Usage:
#   make finish-story SLUG=ref-architecture
#   make finish-story PATH=stories/20250928-ref-architecture

slug="${1:-}"  # from make: first arg is SLUG (may be empty)
pth="${2:-}"   # from make: second arg is PATH (optional)

if [[ -n "$pth" ]]; then
  src="$pth"
elif [[ -n "$slug" ]]; then
  # find story dir by slug (exclude _TEMPLATE and finish)
  matches=( $(ls -d stories/*-"$slug" 2>/dev/null | grep -vE 'stories/_TEMPLATE|stories/finish' || true) )
  if [[ ${#matches[@]} -eq 0 ]]; then
    echo "[finish_story] not found for slug: $slug" >&2
    exit 1
  fi
  if [[ ${#matches[@]} -gt 1 ]]; then
    printf "[finish_story] multiple candidates for slug '%s':\n" "$slug" >&2
    printf '  %s\n' "${matches[@]}" >&2
    exit 1
  fi
  src="${matches[0]}"
else
  echo "Usage: make finish-story SLUG=<slug> | PATH=<path>" >&2
  exit 1
fi

if [[ ! -d "$src" ]]; then
  echo "[finish_story] source dir not found: $src" >&2
  exit 1
fi

dest_dir="stories/finish"
mkdir -p "$dest_dir"
base="$(basename "$src")"
dest="$dest_dir/$base"
if [[ -e "$dest" ]]; then
  echo "[finish_story] destination already exists: $dest" >&2
  exit 1
fi

# Update status in README.md to [完了]
readme="$src/README.md"
if [[ -f "$readme" ]]; then
  # try BSD sed (macOS)
  if sed -i '' -E 's/^ステータス: \[[^]]+\]/ステータス: [完了]/' "$readme" 2>/dev/null; then :; else
    # fallback to GNU sed
    sed -i -E 's/^ステータス: \[[^]]+\]/ステータス: [完了]/' "$readme"
  fi
fi

# Move directory (use git mv when possible)
if command -v git >/dev/null 2>&1 && git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  git mv "$src" "$dest"
else
  mv "$src" "$dest"
fi

echo "[finish_story] archived to: $dest"
