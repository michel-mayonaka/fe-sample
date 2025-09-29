package geom

import "testing"

func TestRectContains_Basic(t *testing.T) {
	if !RectContains(5, 6, 0, 0, 10, 10) {
		t.Fatal("expected point (5,6) inside (0,0,10,10)")
	}
	if RectContains(10, 0, 0, 0, 10, 10) {
		t.Fatal("expected x==x+w to be outside (right edge excluded)")
	}
	if RectContains(0, 10, 0, 0, 10, 10) {
		t.Fatal("expected y==y+h to be outside (bottom edge excluded)")
	}
	if !RectContains(0, 0, 0, 0, 10, 10) {
		t.Fatal("expected left-top to be inside")
	}
}

func TestRectContains_NegativeOrigin(t *testing.T) {
	if !RectContains(-2, -2, -5, -5, 4, 4) { // [-5,-1) × [-5,-1)
		t.Fatal("expected (-2,-2) inside (-5,-5,4,4)")
	}
	if RectContains(0, 0, -5, -5, 0, 4) {
		t.Fatal("w==0 must be outside")
	}
}
