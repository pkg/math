package math

import "math/big"

// MaxBigInt returns the larger of the two *big.Ints
func MaxBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

// MinBigInt returns the smaller of the two *big.Ints
func MinBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

// EqualBigInt returns true if both *big.Ints are equal
func EqualBigInt(a, b *big.Int) bool { return a.Cmp(b) == 0 }
