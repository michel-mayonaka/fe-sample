# 04_spike_adapter — アダプタのスパイク

## 目的
既存 `service.Input` をそのまま使いながら、Scenes からは `ui/input.Reader` を見る構成を検証。

## 概要
- `type Reader = *service.Input` ではなく、Scenes の import を `ui/input` に寄せ、`Action` も `ui/input` 提供に切替。
- 一時的に `ui/input` 側で `type Action = service.Action` の型エイリアスを使う選択肢もあり（段階移行用）。

## 成功条件
- 代表1画面が `ui/input` への参照でビルド通過。

