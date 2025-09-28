package scenes

// 共通ヘルパ（公開）
func PointIn(px, py, x, y, w, h int) bool { return px >= x && py >= y && px < x+w && py < y+h }
