#!/usr/bin/env bash
set -euo pipefail
echo "[MCP] fmt → mcp（vet/build/lint/test）を実行します"
make fmt
make mcp
echo "[MCP] 完了しました"
