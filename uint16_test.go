package math

import "testing"

var maxUint16Tests = []struct {
	a, b uint16
	want uint16
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
}

func TestMaxUint16(t *testing.T) {
	for i, tt := range maxUint16Tests {
		got := MaxUint16(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxUint(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
