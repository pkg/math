package math

import (
	"math/big"
	"testing"
)

func r(n,d int64) *big.Rat { return big.NewRat(n,d) }

var equalBigRatTests = []struct {
	a, b *big.Rat
	want bool
}{
	{ r(1,2), r(1,2), true },
	{ r(1,1), r(1,1), true },
	{ r(1,2), r(2,4), true },
	{ r(1,2), r(2,3), false },
}

func TestEqualBigRat(t *testing.T) {
	for i, tt := range equalBigRatTests {
		got := EqualBigRat(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("%d: EqualBigRat(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}

var maxBigRatTests = []struct {
	a, b *big.Rat
	want *big.Rat
}{
	{ r(1,2), r(1,3), r(1,2) },
	{ r(1,3), r(1,2), r(1,2) },
	{ r(1,2), r(1,4), r(1,2) },
	{ r(1,4), r(1,2), r(1,2) },
}

func TestMaxBigRat(t *testing.T) {
	for i, tt := range maxBigRatTests {
		got := MaxBigRat(tt.a, tt.b)
		if !EqualBigRat(tt.want, got) {
			t.Errorf("%d: MaxBigRat(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}

var minBigRatTests = []struct {
	a, b *big.Rat
	want *big.Rat
}{
	{ r(1,2), r(1,3), r(1,3) },
	{ r(1,3), r(1,2), r(1,3) },
	{ r(1,2), r(1,4), r(1,4) },
	{ r(1,4), r(1,2), r(1,4) },
}

func TestMinBigRat(t *testing.T) {
	for i, tt := range minBigRatTests {
		got := MinBigRat(tt.a, tt.b)
		if !EqualBigRat(tt.want, got) {
			t.Errorf("%d: MinBigRat(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
