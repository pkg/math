package math

import (
	"math/big"
	"testing"
)

func i(i int64) *big.Int { return big.NewInt(i) }

var equalBigIntTests = []struct {
	a, b *big.Int
	want bool
}{
	{i(-1), i(-1), true},
	{i(-1), i(0), false},
	{i(0), i(0), true},
	{i(0), i(-1), false},
	{i(1), i(-1), false},
	{i(-1), i(1), false},
	{i(1), i(1), true},
}

func TestEqualBigInt(t *testing.T) {
	for i, tt := range equalBigIntTests {
		got := EqualBigInt(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("%d: EqualBigInt(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}

var maxBigIntTests = []struct {
	a, b *big.Int
	want *big.Int
}{
	{i(0), i(0), i(0)},
	{i(1), i(1), i(1)},
	{i(1), i(2), i(2)},
	{i(2), i(1), i(2)},
	{i(-2), i(1), i(1)},
	{i(-1), i(1), i(1)},
	{i(1), i(-1), i(1)},
}

func TestMaxBigInt(t *testing.T) {
	for i, tt := range maxBigIntTests {
		got := MaxBigInt(tt.a, tt.b)
		if !EqualBigInt(tt.want, got) {
			t.Errorf("%d: MaxBigInt(%v, %v) = %v, want %v", i+1, tt.a, tt.b, got, tt.want)
		}
	}
}
