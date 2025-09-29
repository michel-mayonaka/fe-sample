package scenes

// PointIn は点 (px,py) が矩形 (x,y,w,h) 内に含まれるかを判定します。
func PointIn(px, py, x, y, w, h int) bool { return px >= x && py >= y && px < x+w && py < y+h }
