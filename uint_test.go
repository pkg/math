package math

import "testing"

var maxUintTests = []struct {
	a, b uint
	want uint
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
}

func TestMaxUint(t *testing.T) {
	for i, tt := range maxUintTests {
		got := MaxUint(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxUint(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
