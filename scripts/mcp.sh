#!/usr/bin/env bash
set -euo pipefail
echo "[MCP] fmt → check → lint を実行します"
make fmt
make check
make lint
echo "[MCP] すべて成功しました"

