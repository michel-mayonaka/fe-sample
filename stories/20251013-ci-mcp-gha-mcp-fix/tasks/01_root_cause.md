# 01: 失敗原因の整理と再現条件の明確化

ステータス: [完了]
担当: @tkg-engineer（想定）
開始: 2025-10-13 16:39:05 +0900

## 目的
- GitHub Actions で `make mcp` が失敗する直接原因を明確化し、`check-ui` の環境依存スキップ条件に落とし込む。

## 背景（ログ要約）
- `make check-all` → `make check-ui` で失敗。
- エラー抜粋:
  - `fatal error: X11/Xlib.h: No such file or directory`
  - ebiten/GLFW の Linux ビルドで X11 ヘッダが必要。`ubuntu-latest` には未導入。

## 完了条件（DoD）
- [x] 失敗を再現できる最小手順を README に追記（Linux ヘッドレス想定）。
- [x] エラーパターン（正規表現）を定義: `X11/Xlib.h|GL/gl.h|libX11|xorg|wayland|xcb|GLX|EGL`。
- [x] スキップ時のユーザ向けメッセージ文言（原因/導入すべきパッケージ）を確定。

## 作業手順（概略）
- ログから典型的エラー語を抽出し、`grep -Ei` で検知できる形に正規表現化。
- Linux（Docker/act など）で `go build ./cmd/ui_sample` を試行し再現。
- 導入が必要なパッケージの候補を整理（例: `xorg-dev libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev`）。

## 作業ログ
- 2025-10-13 21:05:00 +0900: 失敗ログを分析し、X11/GL 系の欠如を特定（IOP=++）
- 2025-10-13 21:10:00 +0900: 検知パターンと回避策（apt 依存導入）を策定（IOP=++）

## 成果物リンク
- README.md（CI 章の追記）
- Makefile（`check-ui` の検知・メッセージ）
