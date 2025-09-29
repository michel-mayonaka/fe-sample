// Package adapter は、UI 表示用データ（View-Model）の生成や
// 既存テーブル/ユーザデータからの変換を担います。
//
// 依存方向:
// - adapter → model/user（定義・ユーザ保存物）
// - adapter → ui/view（表示用構造体）
// - adapter は scenes や draw/layout に依存しません。
package adapter
