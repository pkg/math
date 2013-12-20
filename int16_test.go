package math

import "testing"

var maxInt16NTests = []struct {
	v    []int16
	want int16
}{
	{[]int16{}, 0},
	{[]int16{0}, 0},
	{[]int16{-1}, -1},
	{[]int16{0, 0}, 0},
	{[]int16{1, 1}, 1},
	{[]int16{1, 2}, 2},
	{[]int16{2, 1}, 2},
	{[]int16{-1, 1}, 1},
	{[]int16{1, -1}, 1},
	{[]int16{-1, 0, 1}, 1},
	{[]int16{100, 42, 42, 3}, 100},
	{[]int16{100, 42, 17, 2, 3}, 100},
}

func TestMaxInt16N(t *testing.T) {
	for i, tt := range maxInt16NTests {
		got := MaxInt16N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxInt16(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minInt16NTests = []struct {
	v    []int16
	want int16
}{
	{[]int16{}, 0},
	{[]int16{0}, 0},
	{[]int16{-1}, -1},
	{[]int16{0, 0}, 0},
	{[]int16{1, 1}, 1},
	{[]int16{1, 2}, 1},
	{[]int16{2, 1}, 1},
	{[]int16{-1, 1}, -1},
	{[]int16{1, -1}, -1},
	{[]int16{-1, 0, 1}, -1},
	{[]int16{100, 42, 17, -3}, -3},
	{[]int16{100, 42, 17, 2, 3}, 2},
}

func TestMinInt16(t *testing.T) {
	for i, tt := range minInt16NTests {
		got := MinInt16N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinInt16(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
