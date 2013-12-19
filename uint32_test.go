package math

import "testing"

var maxUint32Tests = []struct {
	v    []uint32
	want uint32
}{
	{[]uint32{}, 0},
	{[]uint32{0}, 0},
	{[]uint32{0, 0}, 0},
	{[]uint32{1, 1}, 1},
	{[]uint32{1, 2}, 2},
	{[]uint32{2, 1}, 2},
	{[]uint32{1, 0, 2}, 2},
	{[]uint32{100, 200, 42, 17, 2, 3}, 200},
}

func TestMaxUint32(t *testing.T) {
	for i, tt := range maxUint32Tests {
		got := MaxUint32(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxUint32(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minUint32Tests = []struct {
	v    []uint32
	want uint32
}{
	{[]uint32{}, 0},
	{[]uint32{0}, 0},
	{[]uint32{0, 0}, 0},
	{[]uint32{1, 1}, 1},
	{[]uint32{1, 2}, 1},
	{[]uint32{2, 1}, 1},
	{[]uint32{1, 0, 2}, 0},
	{[]uint32{100, 200, 42, 17, 2, 3}, 2},
}

func TestMinUint32(t *testing.T) {
	for i, tt := range minUint32Tests {
		got := MinUint32(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinUint32(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
