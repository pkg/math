package math

import "testing"

var maxUintNTests = []struct {
	v    []uint
	want uint
}{
	{[]uint{}, 0},
	{[]uint{0}, 0},
	{[]uint{0, 0}, 0},
	{[]uint{1, 1}, 1},
	{[]uint{1, 2}, 2},
	{[]uint{2, 1}, 2},
	{[]uint{1, 0, 2}, 2},
	{[]uint{100, 200, 42, 17, 2, 3}, 200},
}

func TestMaxUintN(t *testing.T) {
	for i, tt := range maxUintNTests {
		got := MaxUintN(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxUint(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minUintNTests = []struct {
	v    []uint
	want uint
}{
	{[]uint{}, 0},
	{[]uint{0}, 0},
	{[]uint{0, 0}, 0},
	{[]uint{1, 1}, 1},
	{[]uint{1, 2}, 1},
	{[]uint{2, 1}, 1},
	{[]uint{1, 0, 2}, 0},
	{[]uint{100, 200, 42, 17, 2, 3}, 2},
}

func TestMinUint(t *testing.T) {
	for i, tt := range minUintNTests {
		got := MinUintN(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinUint(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
