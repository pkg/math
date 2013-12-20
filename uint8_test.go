package math

import "testing"

var maxUint8NTests = []struct {
	v    []uint8
	want uint8
}{
	{[]uint8{}, 0},
	{[]uint8{0}, 0},
	{[]uint8{0, 0}, 0},
	{[]uint8{1, 1}, 1},
	{[]uint8{1, 2}, 2},
	{[]uint8{2, 1}, 2},
	{[]uint8{1, 0, 2}, 2},
	{[]uint8{100, 200, 42, 17, 2, 3}, 200},
}

func TestMaxUint8N(t *testing.T) {
	for i, tt := range maxUint8NTests {
		got := MaxUint8N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxUint8(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minUint8NTests = []struct {
	v    []uint8
	want uint8
}{
	{[]uint8{}, 0},
	{[]uint8{0}, 0},
	{[]uint8{0, 0}, 0},
	{[]uint8{1, 1}, 1},
	{[]uint8{1, 2}, 1},
	{[]uint8{2, 1}, 1},
	{[]uint8{1, 0, 2}, 0},
	{[]uint8{100, 200, 42, 17, 2, 3}, 2},
}

func TestMinUint8(t *testing.T) {
	for i, tt := range minUint8NTests {
		got := MinUint8N(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinUint8(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
