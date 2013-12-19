package math

import "testing"

var maxInt8Tests = []struct {
	v    []int8
	want int8
}{
	{[]int8{}, 0},
	{[]int8{0}, 0},
	{[]int8{-1}, -1},
	{[]int8{0, 0}, 0},
	{[]int8{1, 1}, 1},
	{[]int8{1, 2}, 2},
	{[]int8{2, 1}, 2},
	{[]int8{-1, 1}, 1},
	{[]int8{1, -1}, 1},
	{[]int8{-1, 0, 1}, 1},
	{[]int8{100, 42, 17, 2, 3}, 100},
}

func TestMaxInt8(t *testing.T) {
	for i, tt := range maxInt8Tests {
		got := MaxInt8(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxInt8(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minInt8Tests = []struct {
	v    []int8
	want int8
}{
	{[]int8{}, 0},
	{[]int8{0}, 0},
	{[]int8{-1}, -1},
	{[]int8{0, 0}, 0},
	{[]int8{1, 1}, 1},
	{[]int8{1, 2}, 1},
	{[]int8{2, 1}, 1},
	{[]int8{-1, 1}, -1},
	{[]int8{1, -1}, -1},
	{[]int8{-1, 0, 1}, -1},
	{[]int8{100, 42, 17, 2, 3}, 2},
}

func TestMinInt8(t *testing.T) {
	for i, tt := range minInt8Tests {
		got := MinInt8(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinInt8(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
