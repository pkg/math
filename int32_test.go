package math

import "testing"

var maxInt32Tests = []struct {
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
	{[]int32{100, 42, 17, 2, 3}, 100},
}

func TestMaxInt32(t *testing.T) {
	for i, tt := range maxInt32Tests {
		got := MaxInt32(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxInt32(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minInt32Tests = []struct {
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
	{[]int32{100, 42, 17, 2, 3}, 2},
}

func TestMinInt32(t *testing.T) {
	for i, tt := range minInt32Tests {
		got := MinInt32(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinInt32(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
