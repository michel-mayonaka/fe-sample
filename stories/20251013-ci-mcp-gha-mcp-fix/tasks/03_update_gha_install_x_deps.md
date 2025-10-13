# 03: GitHub Actions に厳格 UI ビルドジョブを追加（依存導入）

ステータス: [未着手]
担当: @tkg-engineer（想定）
開始: 2025-10-13 16:39:05 +0900

## 目的
- `mcp` は非 strict で通しつつ、別ジョブで X11/GL 依存を導入して `cmd/ui_sample` の厳格ビルドを担保する。

## 完了条件（DoD）
- [ ] 新ジョブ `ui-build-strict` が `ubuntu-latest` で必要パッケージ導入後に `go build ./cmd/ui_sample` 成功。
- [ ] `build-and-lint`（既存 mcp ジョブ）は従来通り緑のまま。
- [ ] ワークフローの所要時間が許容範囲（目安 5〜7 分）に収まる。

## 変更案（パッチ概要）
`.github/workflows/ci.yml` にジョブを追加:

```yaml
  ui-build-strict:
    runs-on: ubuntu-latest
    needs: [build-and-lint]
    timeout-minutes: 10
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.25.x'
          cache: true
      - name: Install X11/GL dev packages
        run: |
          sudo apt-get update
          sudo apt-get install -y \
            xorg-dev libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev
      - name: Build UI (strict)
        env:
          MCP_STRICT: "1"
          GOFLAGS: -mod=readonly
        run: |
          make check-ui
```

## 作業手順（概略）
- YAML を編集して新ジョブを追加、キャッシュは既存キーを流用。
- 実行順は `needs` で `build-and-lint` 後に設定（早期失敗のため）。
- 実行後に `check-ui` のログが strict で成功することを確認。

## 成果物リンク
- 変更ファイル: `.github/workflows/ci.yml`

