package geom

// RectContains は点 (px,py) が矩形 (x,y,w,h) 内に含まれるかを判定します。
//
// 半開区間 [x, x+w) × [y, y+h) を採用します。すなわち左上辺は含み、右下辺は含みません。
// 幅/高さが 0 以下の場合は常に false を返します。
func RectContains(px, py, x, y, w, h int) bool {
    if w <= 0 || h <= 0 { return false }
    return px >= x && py >= y && px < x+w && py < y+h
}

