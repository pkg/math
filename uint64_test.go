package math

import "testing"

var maxUint64Tests = []struct {
	a, b uint64
	want uint64
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
}

func TestMaxUint64(t *testing.T) {
	for i, tt := range maxUint64Tests {
		got := MaxUint64(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxUint(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}

var minUint64Tests = []struct {
	a, b uint64
	want uint64
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 1},
	{2, 1, 1},
}

func TestMinUint64(t *testing.T) {
	for i, tt := range minUint64Tests {
		got := MinUint64(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MinUint(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
