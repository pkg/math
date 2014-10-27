package math

import "math/big"

// MaxBigRat returns the larger of the two *big.Rats
func MaxBigRat(a, b *big.Rat) *big.Rat {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

// MinBigRat returns the smaller of the two *big.Rats
func MinBigRat(a, b *big.Rat) *big.Rat {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

// EqualBigRat returns true if both *big.Rats are equal
func EqualBigRat(a, b *big.Rat) bool { return a.Cmp(b) == 0 }
