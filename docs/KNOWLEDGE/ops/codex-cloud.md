# codex-cloud 実行ガイド（ワークスペース限定・ネットワーク制限）

本書は codex-cloud 環境（ワークスペース書き込みのみ、ネットワークは原則不可、重要コマンドは承認フロー）で、開発・検証を安定して行うための最小手順を示します。

## 前提
- 依存は `vendor/` に固定済み（未整備の場合はローカル/オンラインで `make vendor-sync` を実行してコミット）。
- UI ビルドは環境依存のため、既定では「検出したらスキップ（非strict）」で運用します。
- 生成物は `out/` 配下に集約されます（ログは `out/logs/`）。

## すぐに試す（最短）
- スモーク（論理層のみ / オフライン強制）
```sh
make smoke
```
- 包括的なオフライン検証（`make mcp` をオフライン前提で実行）
```sh
MCP_OFFLINE=1 make offline
```

## 環境変数の意味
- `MCP_OFFLINE=1`: Go を `-mod=vendor` で動作させ、`GOPROXY=off` を前提にします。
- `MCP_STRICT=1`: UI のビルド失敗をスキップせず失敗扱いにします（Linux は X11/GL 開発ヘッダ導入が必要）。

## UI ビルドに関して
- 既定（`MCP_STRICT=0`）では `make check-ui` が環境依存の既知エラーを検知した際に自動スキップします。
- 厳格に検証したい場合:
```sh
# 例: Debian/Ubuntu で必要パッケージ導入後
MCP_STRICT=1 make check-ui
```
- 直近の UI ビルドエラーは `out/logs/ui_build_err.log` に保存されます。

## 出力/クリーン
- バイナリ: `out/bin/`
- ログ: `out/logs/`
- カバレッジ等: `out/coverage/`
- 片付け: `make clean`

## トラブルシュート
- vendor が無い/不足している: オンライン環境で `make vendor-sync` → コミット。
- Lint ツールが無い: `make lint` は自動スキップします（CI も非必須扱い）。
- UI がビルドできない: 非strictではスキップされます。厳格に検証する場合は Linux に X11/GL 依存を導入してください。

## 参考
- 詳細: `docs/OFFLINE.md`
- CI 方針と整合: README の CI セクション
- Make ターゲット: `make smoke`, `make offline`, `make mcp`, `make clean`
