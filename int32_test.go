package math

import "testing"

var maxInt32Tests = []struct {
	a, b int32
	want int32
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
	{-1, 1, 1},
	{1, -1, 1},
}

func TestMaxInt32(t *testing.T) {
	for i, tt := range maxInt32Tests {
		got := MaxInt32(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxInt(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
