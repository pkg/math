package math

import "testing"

var maxIntTests = [] struct {
	a, b int
	want int
} {
	{ 0, 0, 0 },
	{ 1, 1, 1 },
	{ 1, 2, 2 },
	{ 2, 1, 2 },
	{ -1, 1, 1},
	{ 1, -1, 1},
}

func TestMaxInt(t *testing.T) {
	for i, tt := range maxIntTests {
		got := MaxInt(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxInt(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
