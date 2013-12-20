package math

import "testing"

var maxInt32NTests = []struct {
	v    []int32
	want int32
}{
	{[]int32{}, 0},
	{[]int32{0}, 0},
	{[]int32{-1}, -1},
	{[]int32{0, 0}, 0},
	{[]int32{1, 1}, 1},
	{[]int32{1, 2}, 2},
	{[]int32{2, 1}, 2},
	{[]int32{-1, 1}, 1},
	{[]int32{1, -1}, 1},
	{[]int32{-1, 0, 1}, 1},
	{[]int32{100, 42, 42, 3}, 100},
	{[]int32{100, 42, 17, 2, 3}, 100},
}

func TestMaxInt32N(t *testing.T) {
	for i, tt := range maxInt32NTests {
		got := MaxInt32N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxInt32(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minInt32NTests = []struct {
	v    []int32
	want int32
}{
	{[]int32{}, 0},
	{[]int32{0}, 0},
	{[]int32{-1}, -1},
	{[]int32{0, 0}, 0},
	{[]int32{1, 1}, 1},
	{[]int32{1, 2}, 1},
	{[]int32{2, 1}, 1},
	{[]int32{-1, 1}, -1},
	{[]int32{1, -1}, -1},
	{[]int32{-1, 0, 1}, -1},
	{[]int32{100, 42, 17, -3}, -3},
	{[]int32{100, 42, 17, 2, 3}, 2},
}

func TestMinInt32(t *testing.T) {
	for i, tt := range minInt32NTests {
		got := MinInt32N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinInt32(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
