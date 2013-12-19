package math

import "testing"

var maxTests = []struct {
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

func TestMax(t *testing.T) {
	for i, tt := range maxTests {
		got := Max(tt.v...)
		if tt.want != got {
			t.Errorf("%d: Max(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minTests = []struct {
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

func TestMin(t *testing.T) {
	for i, tt := range minTests {
		got := Min(tt.v...)
		if tt.want != got {
			t.Errorf("%d: Min(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
