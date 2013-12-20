package math

import "testing"

var maxInt64NTests = []struct {
	v    []int64
	want int64
}{
	{[]int64{}, 0},
	{[]int64{0}, 0},
	{[]int64{-1}, -1},
	{[]int64{0, 0}, 0},
	{[]int64{1, 1}, 1},
	{[]int64{1, 2}, 2},
	{[]int64{2, 1}, 2},
	{[]int64{-1, 1}, 1},
	{[]int64{1, -1}, 1},
	{[]int64{-1, 0, 1}, 1},
	{[]int64{100, 42, 42, 3}, 100},
	{[]int64{100, 42, 17, 2, 3}, 100},
}

func TestMaxInt64N(t *testing.T) {
	for i, tt := range maxInt64NTests {
		got := MaxInt64N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxInt64(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minInt64NTests = []struct {
	v    []int64
	want int64
}{
	{[]int64{}, 0},
	{[]int64{0}, 0},
	{[]int64{-1}, -1},
	{[]int64{0, 0}, 0},
	{[]int64{1, 1}, 1},
	{[]int64{1, 2}, 1},
	{[]int64{2, 1}, 1},
	{[]int64{-1, 1}, -1},
	{[]int64{1, -1}, -1},
	{[]int64{-1, 0, 1}, -1},
	{[]int64{100, 42, 17, -3}, -3},
	{[]int64{100, 42, 17, 2, 3}, 2},
}

func TestMinInt64(t *testing.T) {
	for i, tt := range minInt64NTests {
		got := MinInt64N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinInt64(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
