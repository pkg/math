package math

import "testing"

var maxInt64Tests = []struct {
	a, b int64
	want int64
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
	{-1, 1, 1},
	{1, -1, 1},
}

func TestMaxInt64(t *testing.T) {
	for i, tt := range maxInt64Tests {
		got := MaxInt64(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxInt(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
