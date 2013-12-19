package math

import "testing"

var maxInt8Tests = []struct {
	a, b int8
	want int8
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
	{-1, 1, 1},
	{1, -1, 1},
}

func TestMaxInt8(t *testing.T) {
	for i, tt := range maxInt8Tests {
		got := MaxInt8(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MaxInt8(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}

var minInt8Tests = []struct {
	a, b int8
	want int8
}{
	{0, 0, 0},
	{1, 1, 1},
	{1, 2, 1},
	{2, 1, 1},
	{-1, 1, -1},
	{1, -1, -1},
}

func TestMinInt8(t *testing.T) {
	for i, tt := range minInt8Tests {
		got := MinInt8(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%d: MinInt8(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
