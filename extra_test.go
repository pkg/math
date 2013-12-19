package math

import "testing"

var maxTests = []struct {
	a, b int
	want int
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
	{-1, 1, 1},
	{1, -1, 1},
}

func TestMax(t *testing.T) {
	for i, tt := range maxTests {
		got := Max(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxInt(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}

var minTests = []struct {
	a, b int
	want int
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 1},
	{2, 1, 1},
	{-1, 1, -1},
	{1, -1, -1},
}

func TestMin(t *testing.T) {
	for i, tt := range minTests {
		got := Min(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MinInt(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
