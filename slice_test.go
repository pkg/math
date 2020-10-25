package math

import (
	"testing"
)

var minSliceIndexTests = []struct {
	v    []struct{ n int }
	want int
}{
	{
		v:    []struct{ n int }{{1}},
		want: 0,
	},
	{
		v:    []struct{ n int }{{7}, {2}},
		want: 1,
	},
	{
		v:    []struct{ n int }{{12}, {18}, {17}, {9}},
		want: 3,
	},
}

func TestMinSliceIndex(t *testing.T) {
	for i, tt := range minSliceIndexTests {
		got := MinSliceIndex(tt.v, func(i, j int) bool {
			return tt.v[i].n < tt.v[j].n
		})
		if tt.want != got {
			t.Errorf("%d: MinSliceIndex(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var maxSliceIndexTests = []struct {
	v    []struct{ n int }
	want int
}{
	{
		v:    []struct{ n int }{{1}},
		want: 0,
	},
	{
		v:    []struct{ n int }{{7}, {2}},
		want: 0,
	},
	{
		v:    []struct{ n int }{{12}, {18}, {17}, {19}},
		want: 3,
	},
}

func TestMaxSliceIndex(t *testing.T) {
	for i, tt := range maxSliceIndexTests {
		got := MaxSliceIndex(tt.v, func(i, j int) bool {
			return tt.v[i].n < tt.v[j].n
		})
		if tt.want != got {
			t.Errorf("%d: MaxSliceIndex(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
