package math

import "testing"

var maxUint64Tests = []struct {
	v    []uint64
	want uint64
}{
	{[]uint64{}, 0},
	{[]uint64{0}, 0},
	{[]uint64{0, 0}, 0},
	{[]uint64{1, 1}, 1},
	{[]uint64{1, 2}, 2},
	{[]uint64{2, 1}, 2},
	{[]uint64{1, 0, 2}, 2},
	{[]uint64{100, 200, 42, 17, 2, 3}, 200},
}

func TestMaxUint64(t *testing.T) {
	for i, tt := range maxUint64Tests {
		got := MaxUint64(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxUint64(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minUint64Tests = []struct {
	v    []uint64
	want uint64
}{
	{[]uint64{}, 0},
	{[]uint64{0}, 0},
	{[]uint64{0, 0}, 0},
	{[]uint64{1, 1}, 1},
	{[]uint64{1, 2}, 1},
	{[]uint64{2, 1}, 1},
	{[]uint64{1, 0, 2}, 0},
	{[]uint64{100, 200, 42, 17, 2, 3}, 2},
}

func TestMinUint64(t *testing.T) {
	for i, tt := range minUint64Tests {
		got := MinUint64(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinUint64(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
