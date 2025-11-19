# 02 — スキーマ設計とサンプル定義

ステータス: [完了]

## 目的
- UI メトリクスを JSON スキーマ化し、基準解像度・スケール・各画面用の矩形/余白/行高を定義する。

## 設計方針
- 既定の基準解像度は 1920x1080（`uicore.SetBaseResolution` に整合）。
- キーは `snake_case` ではなくコンフィグは `lowerCamel` も可だが、命名規約に従い明確な名前を採用（docs/KNOWLEDGE/engineering/naming.md）。
- ファイル配置案: `db/master/mst_ui_metrics.json`（既定）／`db/user/usr_ui_metrics.json`（上書き）。

## 型スケッチ（Go 側）
```go
type Metrics struct {
  Base struct{ W, H int } `json:"base"`
  List struct {
    Margin       int `json:"margin"`
    ItemH        int `json:"itemH"`
    ItemGap      int `json:"itemGap"`
    TitleOffset  int `json:"titleOffset"`
    PortraitSize int `json:"portraitSize"`
  } `json:"list"`
}
```

## サンプル（`mst_ui_metrics.json`）
```json
{
  "base": { "w": 1920, "h": 1080 },
  "list": {
    "margin": 24,
    "itemH": 100,
    "itemGap": 12,
    "titleOffset": 44,
    "portraitSize": 80
  }
}
```

## 成果物
- スキーマ定義（上記型/命名の確定）
- サンプル JSON を `db/master/mst_ui_metrics.json` に追加（済）
- `internal/config/config.go` にデフォルトパスを追加（`DefaultUIMetricsPath`, `DefaultUserUIMetricsPath` 済）

## 進捗ログ
- 2025-09-29: 雛形作成
- 2025-09-29: スキーマ確定・サンプルファイル追加
