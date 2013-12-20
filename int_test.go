package math

import "testing"

var maxIntNTests = []struct {
	v    []int
	want int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{-1}, -1},
	{[]int{0, 0}, 0},
	{[]int{1, 1}, 1},
	{[]int{1, 2}, 2},
	{[]int{2, 1}, 2},
	{[]int{-1, 1}, 1},
	{[]int{1, -1}, 1},
	{[]int{-1, 0, 1}, 1},
	{[]int{100, 42, 42, 3}, 100},
	{[]int{100, 42, 17, 2, 3}, 100},
}

func TestMaxIntN(t *testing.T) {
	for i, tt := range maxIntNTests {
		got := MaxIntN(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxInt(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minIntNTests = []struct {
	v    []int
	want int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{-1}, -1},
	{[]int{0, 0}, 0},
	{[]int{1, 1}, 1},
	{[]int{1, 2}, 1},
	{[]int{2, 1}, 1},
	{[]int{-1, 1}, -1},
	{[]int{1, -1}, -1},
	{[]int{-1, 0, 1}, -1},
	{[]int{100, 42, 17, -3}, -3},
	{[]int{100, 42, 17, 2, 3}, 2},
}

func TestMinInt(t *testing.T) {
	for i, tt := range minIntNTests {
		got := MinIntN(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinInt(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
