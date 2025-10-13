# 01: 失敗原因の整理と再現条件の明確化

ステータス: [未着手]
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
- [ ] 失敗を再現できる最小手順を README に追記（Linux ヘッドレス想定）。
- [ ] エラーパターン（正規表現）を定義: `X11/Xlib.h|GL/gl.h|libX11|xorg|wayland|xcb|GLX|EGL`。
- [ ] スキップ時のユーザ向けメッセージ文言（原因/導入すべきパッケージ）を確定。

## 作業手順（概略）
- ログから典型的エラー語を抽出し、`grep -Ei` で検知できる形に正規表現化。
- Linux（Docker/act など）で `go build ./cmd/ui_sample` を試行し再現。
- 導入が必要なパッケージの候補を整理（例: `xorg-dev libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev`）。

## 成果物リンク
- なし（このタスクでは調査ノートと正規表現定義のみ）。

