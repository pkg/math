package math

import "testing"

var maxUint8Tests = []struct {
	a, b uint8
	want uint8
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
}

func TestMaxUint8(t *testing.T) {
	for i, tt := range maxUint8Tests {
		got := MaxUint8(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxUint8(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}

var minUint8Tests = []struct {
	a, b uint8
	want uint8
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 1},
	{2, 1, 1},
}

func TestMinUint8(t *testing.T) {
	for i, tt := range minUint8Tests {
		got := MinUint8(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MinUint8(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
